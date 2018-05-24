package tcapi

import (
	"encoding/json"
	"fmt"
)

// it is *profoundly* annoying how similar-but-sufficiently-different
// this is to a ScalingGroup struct
type CreateScalingGroupRequest struct {
	ScalingConfigurationId string   `json:"scalingConfigurationId" url:"scalingConfigurationId"`
	ScalingGroupName       string   `json:"scalingGroupName" url:"scalingGroupName"`
	MinSize                int      `json:"minSize" url:"minSize"`
	MaxSize                int      `json:"maxSize" url:"maxSize"`
	RemovePolicy           string   `json:"removePolicy" url:"removePolicy"`
	VpcId                  string   `json:"vpcId" url:"vpcId,omitempty"`
	ZoneIds                []string `json:"zoneIds" url:"zoneIds,omitempty,dotnumbered"`
	LoadBalancerIds        []string `json:"loadBalancerIds" url:"loadBalancerIds,omitempty,dotnumbered"`
	SubnetIds              []string `json:"subnetIds" url:"subnetIds,omitempty,dotnumbered"`
	ProjectId              int      `json:"projectId" url:"projectId,omitempty"`
	DesiredCapacity        int      `json:"desiredCapacity" url:"desiredCapacity"`
}

type CreateScalingGroupResponse struct {
	Data CreateScalingGroupResponseData `json:"data" url:"data"`
}

type CreateScalingGroupResponseData struct {
	ScalingGroupIdSet []string `json:"scalingGroupIdSet" url:"scalingGroupIdSet,dotnumbered"`
}

func (c *Client) CreateScalingGroup(req *CreateScalingGroupRequest) (*[]string, error) {
	resp, err := c.Do("scaling", "CreateScalingGroup", req)
	if err != nil {
		return nil, fmt.Errorf("[scaling:CreateScalingGroup] request failed: %s", err)
	}

	ret := new(CreateScalingGroupResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[scaling:CreateScalingGroup] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[scaling:CreateScalingGroup] response unmarshaled to nil")
	}

	return &ret.Data.ScalingGroupIdSet, nil
}
