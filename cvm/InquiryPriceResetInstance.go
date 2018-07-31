package tcapi

import (
	"encoding/json"
	"fmt"
)

type InquiryPriceResetInstanceResponse struct {
	RequestId string `json:",omitempty" url:",omitempty"`
	Price     Price  `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) InquiryPriceResetInstance(req *ResetInstanceRequest) (*InquiryPriceResetInstanceResponse, error) {
	resp, err := c.Do("cvm", "InquiryPriceResetInstance", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResetInstance] request failed: %s", err)
	}

	ret := new(InquiryPriceResetInstanceResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResetInstance] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResetInstance] response unmarshaled to nil")
	}

	return ret, nil
}
