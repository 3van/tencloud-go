package tcapi

import (
	"encoding/json"
	"fmt"
)

type InquiryPriceResizeInstanceDisksResponse struct {
	RequestId string `json:",omitempty" url:",omitempty"`
	Price     Price  `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) InquiryPriceResizeInstanceDisks(req *ResizeInstanceDisksRequest) (*InquiryPriceResizeInstanceDisksResponse, error) {
	resp, err := c.Do("cvm", "InquiryPriceResizeInstanceDisks", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResizeInstanceDisks] request failed: %s", err)
	}

	ret := new(InquiryPriceResizeInstanceDisksResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResizeInstanceDisks] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResizeInstanceDisks] response unmarshaled to nil")
	}

	return ret, nil
}
