package tcapi

import (
	"encoding/json"
	"fmt"
)

type InquiryPriceRunInstancesResponse struct {
	RequestId string `json:",omitempty" url:",omitempty"`
	Price     Price  `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) InquiryPriceRunInstances(req *RunInstancesRequest) (*InquiryPriceRunInstancesResponse, error) {
	resp, err := c.Do("cvm", "InquiryPriceRunInstances", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceRunInstances] request failed: %s", err)
	}

	ret := new(InquiryPriceRunInstancesResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceRunInstances] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:InquiryPriceRunInstances] response unmarshaled to nil")
	}

	return ret, nil
}
