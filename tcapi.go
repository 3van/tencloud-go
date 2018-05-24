package tcapi

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/3van/go-querystring/query"
)

const (
	version        = "0.4.1"
	userAgent      = "qcloud-go/" + version
	apiPath        = "/v2/index.php"
	apiBase        = "api.qcloud.com"
	apiProto       = "https"
	defaultTimeout = time.Second * 15
)

var moduleApiVersion = map[string]string{
	"cvm":     "2017-03-12",
	"dfw":     "",
	"eip":     "",
	"image":   "2017-03-12",
	"scaling": "",
	"vpc":     "",
}

// Client instances contain configuration and context information for the API.
type Client struct {
	client *http.Client // the HTTP client instance used to talk to the API

	// API secret identifier
	SecretId string
	// API secret
	Secret string
	// Region to use in API operations
	Region string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New returns a new instance of the client.
// region defaults to "hk" if null. If you wish to use your own http.Client
// instance (to provide caching, et al.), you may provide one; if nil, the
// http.DefaultClient will be used.
//
// if environment variables are present, they will override any values set
// statically here:
// 		TENCENT_REGION     -> region
//		TENCENT_API_KEY_ID -> secret_id
//		TENCENT_API_KEY    -> secret
func New(secret_id, secret, region string, httpClient *http.Client) *Client {
	if os.Getenv("TENCENT_REGION") != "" {
		region = os.Getenv("TENCENT_REGION")
	}
	if os.Getenv("TENCENT_API_KEY_ID") != "" {
		secret_id = os.Getenv("TENCENT_API_KEY_ID")
	}
	if os.Getenv("TENCENT_API_KEY") != "" {
		secret = os.Getenv("TENCENT_API_KEY")
	}

	if region == "" {
		region = "hk"
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	httpClient.Timeout = defaultTimeout

	return &Client{
		client:   httpClient,
		Secret:   secret,
		SecretId: secret_id,
		Region:   region,
	}
}

func (c *Client) Do(module, request string, params interface{}) (*json.RawMessage, error) {
	module = strings.ToLower(module)
	version := ""
	if _, ok := moduleApiVersion[module]; ok {
		version = moduleApiVersion[module]
	}

	reqURL, err := url.Parse(strings.Join([]string{apiProto, "://", module, ".", apiBase, apiPath}, ""))
	if err != nil {
		return nil, fmt.Errorf("could not parse baseURL: %s", err)
	}

	if err != nil {
		return nil, fmt.Errorf("error encoding params struct: %s", err)
	}

	finalParams, err := query.Values(params)
	if err != nil {
		return nil, fmt.Errorf("could not parse parameters: %s", err)
	}

	finalParams.Set("Action", request)
	finalParams.Set("Region", c.Region)
	finalParams.Set("SecretId", c.SecretId)
	finalParams.Set("Timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	finalParams.Set("Nonce", strconv.Itoa(rand.Int()))
	if moduleApiVersion[module] != "" {
		finalParams.Set("Version", version)
	}

	err = c.signRequestParams("GET", *reqURL, &finalParams)
	if err != nil {
		return nil, fmt.Errorf("could not sign request: %s", err)
	}

	reqURL.RawQuery = finalParams.Encode()
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create new request: %s", err)
	}

	req.Header.Set("User-Agent", userAgent)

	//fmt.Printf("[tcapi-debug] *client.Do(): QueryString: %s\n", reqURL.String())

	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request to API: %s", err)
	}

	defer func() {
		io.CopyN(ioutil.Discard, httpResp.Body, 512)
		httpResp.Body.Close()
	}()

	resp, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	} else if err == nil && resp == nil {
		return nil, fmt.Errorf("received empty response body")
	}

	switch version {
	case "2017-03-12":
		v3Resp := new(V3BaseResponse)
		err = json.Unmarshal(resp, v3Resp)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response from API to calling type: %s", err)
		}
		if v3Resp == nil {
			return nil, fmt.Errorf("API response was unmarshalled to nil")
		}

		errResp := new(V3ErrorResponse)
		err = json.Unmarshal(*v3Resp.Response, errResp)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response from API to ErrorResponse: %s", err)
		}
		if errResp.Error.Code != "" {
			return nil, fmt.Errorf("API returned an error (%v - %s): %s", errResp.Error.Code, errResp.Error.CodeDesc, errResp.Error.Message)
		}

		return v3Resp.Response, nil
	default:
		oldResp := new(OldErrorResponse)
		err = json.Unmarshal(resp, oldResp)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response to old-style error format: %s", err)
		}
		if oldResp.Code != 0 {
			return nil, fmt.Errorf("API returned an error (%v - %s): %s", oldResp.Code, oldResp.CodeDesc, oldResp.Message)
		}

		retResp := json.RawMessage(resp)
		return &retResp, nil
	}
}

func (c *Client) signRequestParams(method string, req url.URL, params *url.Values) error {
	method = strings.ToUpper(method)
	req.RawQuery = params.Encode()

	message := fmt.Sprintf("%s%s%s", method, req.Hostname(), req.RequestURI())
	sig := hmac.New(sha1.New, []byte(c.Secret))
	_, err := sig.Write([]byte(message))
	if err != nil {
		return err
	}

	// hashtag VapeLyfeeeeeee
	eSig := base64.StdEncoding.EncodeToString(sig.Sum(nil))
	params.Set("Signature", eSig)

	return nil
}
