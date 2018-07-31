package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeScalingInstanceRequest struct {
	ScalingGroupId string   `json:"scalingGroupId,omitempty" url:"scalingGroupId,omitempty"`
	InstanceIds    []string `json:"instanceIds,omitempty" url:"instanceIds,omitempty,dotnumbered"`
	CreationType   string   `json:"creationType,omitempty" url:"creationType,omitempty"`
	HealthStatus   string   `json:"healthStatus,omitempty" url:"healthStatus,omitempty"`
	Offset         int      `json:"offset,omitempty" url:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty" url:"limit,omitempty"`
}

type DescribeScalingInstanceResponse struct {
	Data DescribeScalingInstanceResponseData
}

type DescribeScalingInstanceResponseData struct {
	TotalCount          int
	ScalingInstancesSet []ScalingInstance `url:",dotnumbered"`
}

func (c *Client) DescribeScalingInstance(req *DescribeScalingInstanceRequest) (*DescribeScalingInstanceResponseData, error) {
	ret := new(DescribeScalingInstanceResponse)
	resp, err := c.Do("scaling", "DescribeScalingInstance", req)
	if err != nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingInstance] request failed: %s", err)
	}

	err = json.Unmarshal(*resp, ret)

	if err != nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingInstance] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingInstance] response unmarshaled to nil")
	}

	return &ret.Data, nil
}
