package tcapi

import (
	"encoding/json"
	"fmt"
)

type ImportKeyPairRequest struct {
	KeyName   string `json:",omitempty" url:",omitempty"`
	ProjectId int    `json:",omitempty" url:",omitempty"`
	PublicKey string `json:",omitempty" url:",omitempty"`
}

type ImportKeyPairResponse struct {
	RequestId string `json:",omitempty" url:",omitempty"`
	KeyId     string `json:",omitempty" url:",omitempty"`
}

func (c *Client) ImportKeyPair(req *ImportKeyPairRequest) (*ImportKeyPairResponse, error) {
	resp, err := c.Do("cvm", "ImportKeyPair", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:ImportKeyPair] request failed: %s", err)
	}

	ret := new(ImportKeyPairResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:ImportKeyPair] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:ImportKeyPair] response unmarshaled to nil")
	}

	return ret, nil
}
