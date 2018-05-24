package tcapi

import (
	"fmt"
)

type ModifyScalingGroupRequest struct {
	ScalingGroupId         string `json:"scalingGroupId" url:"scalingGroupId"`
	MinSize                int    `json:"minSize" url:"minSize"`
	MaxSize                int    `json:"maxSize" url:"maxSize"`
	RemovePolicy           string `json:"removePolicy" url:"removePolicy"`
	ScalingGroupName       string `json:"scalingGroupName,omitempty" url:"scalingGroupName,omitempty"`
	ScalingConfigurationId string `json:"scalingConfigurationId,omitempty" url:"scalingConfigurationId,omitempty"`
	DesiredCapacity        int    `json:"desiredCapacity" url:"desiredCapacity"`
}

func (c *Client) ModifyScalingGroup(req *ModifyScalingGroupRequest) error {
	_, err := c.Do("scaling", "ModifyScalingGroup", req)
	if err != nil {
		return fmt.Errorf("[scaling:ModifyScalingGroup] request failed: %s", err)
	}

	return nil
}
