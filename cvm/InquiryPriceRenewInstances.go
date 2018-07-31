package tcapi

import (
	"encoding/json"
	"fmt"
)

type InquiryPriceRenewInstancesResponse struct {
	RequestId string `json:",omitempty" url:",omitempty"`
	Price     Price  `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) InquiryPriceRenewInstances(req *RenewInstancesRequest) (*InquiryPriceRenewInstancesResponse, error) {
	resp, err := c.Do("cvm", "InquiryPriceRenewInstances", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceRenewInstances] request failed: %s", err)
	}

	ret := new(InquiryPriceRenewInstancesResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceRenewInstances] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceRenewInstances] response unmarshaled to nil")
	}

	return ret, nil
}
