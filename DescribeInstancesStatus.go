package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeInstancesStatusRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	Offset      int      `json:",omitempty" url:",omitempty"`
	Limit       int      `json:",omitempty" url:",omitempty"`
}

type DescribeInstancesStatusResponse struct {
	RequestId         string           `json:",omitempty" url:",omitempty"`
	TotalCount        int              `json:",omitempty" url:",omitempty"`
	InstanceStatusSet []InstanceStatus `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) DescribeInstancesStatus(req *DescribeInstancesStatusRequest) (*DescribeInstancesStatusResponse, error) {
	resp, err := c.Do("cvm", "DescribeInstancesStatus", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeInstancesStatus] request failed: %s", err)
	}

	ret := new(DescribeInstancesStatusResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeInstancesStatus] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:DescribeInstancesStatus] response unmarshaled to nil")
	}

	return ret, nil
}
