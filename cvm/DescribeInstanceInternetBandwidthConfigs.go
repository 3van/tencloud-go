package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeInstanceInternetBandwidthConfigsRequest struct {
	InstanceId string `json:",omitempty" url:",omitempty"`
}

type DescribeInstanceInternetBandwidthConfigsResponse struct {
	RequestId                  string                    `json:",omitempty" url:",omitempty"`
	InternetBandwidthConfigSet []InternetBandwidthConfig `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) DescribeInstanceInternetBandwidthConfigs(req *DescribeInstanceInternetBandwidthConfigsRequest) (*DescribeInstanceInternetBandwidthConfigsResponse, error) {
	resp, err := c.Do("cvm", "DescribeInstanceInternetBandwidthConfigs", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeInstanceInternetBandwidthConfigs] request failed: %s", err)
	}

	ret := new(DescribeInstanceInternetBandwidthConfigsResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeInstanceInternetBandwidthConfigs] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:DescribeInstanceInternetBandwidthConfigs] response unmarshaled to nil")
	}

	return ret, nil
}
