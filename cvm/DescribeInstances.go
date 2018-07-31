package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeInstancesRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	Filters     []Filter `json:",omitempty" url:",omitempty,dotnumbered"`
	Offset      int      `json:",omitempty" url:",omitempty"`
	Limit       int      `json:",omitempty" url:",omitempty"`
}

type DescribeInstancesResponse struct {
	RequestId   string     `json:",omitempty" url:",omitempty"`
	TotalCount  int        `json:",omitempty" url:",omitempty"`
	InstanceSet []Instance `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) DescribeInstances(req *DescribeInstancesRequest) (*DescribeInstancesResponse, error) {
	resp, err := c.Do("cvm", "DescribeInstances", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeInstances] request failed: %s", err)
	}

	ret := new(DescribeInstancesResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeInstances] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:DescribeInstances] response unmarshaled to nil")
	}

	return ret, nil
}
