package tcapi

import (
	"encoding/json"
	"fmt"
)

type InquiryPriceResetInstancesTypeResponse struct {
	RequestId string `json:",omitempty" url:",omitempty"`
	Price     Price  `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) InquiryPriceResetInstancesType(req *ResetInstancesTypeRequest) (*InquiryPriceResetInstancesTypeResponse, error) {
	resp, err := c.Do("cvm", "InquiryPriceResetInstancesType", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResetInstancesType] request failed: %s", err)
	}

	ret := new(InquiryPriceResetInstancesTypeResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResetInstancesType] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceResetInstancesType] response unmarshaled to nil")
	}

	return ret, nil
}
