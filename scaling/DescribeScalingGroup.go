package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeScalingGroupRequest struct {
	ScalingGroupIds        []string `json:"scalingGroupIds,omitempty" url:"scalingGroupIds,omitempty,dotnumbered"`
	ScalingGroupName       string   `json:"scalingGroupName,omitempty" url:"scalingGroupName,omitempty"`
	ScalingConfigurationId string   `json:"scalingConfigurationId,omitempty" url:"scalingConfigurationId,omitempty"`
	Offset                 int      `json:"offset,omitempty" url:"offset,omitempty"`
	Limit                  int      `json:"limit,omitempty" url:"limit,omitempty"`
	VpcId                  string   `json:"vpcId,omitempty" url:"vpcId,omitempty"`
	ProjectId              int      `json:"projectId,omitempty" url:"projectId,omitempty"`
}

type DescribeScalingGroupResponse struct {
	Data DescribeScalingGroupResponseData
}

type DescribeScalingGroupResponseData struct {
	TotalCount      int
	ScalingGroupSet []ScalingGroup `url:",dotnumbered"`
}

func (c *Client) DescribeScalingGroup(request *DescribeScalingGroupRequest) (*DescribeScalingGroupResponseData, error) {
	ret := new(DescribeScalingGroupResponse)
	resp, err := c.Do("scaling", "DescribeScalingGroup", request)
	if err != nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingGroup] request failed: %s", err)
	}

	err = json.Unmarshal(*resp, ret)

	if err != nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingGroup] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingGroup] response unmarshaled to nil")
	}

	return &ret.Data, nil
}
