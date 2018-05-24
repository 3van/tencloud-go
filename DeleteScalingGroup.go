package tcapi

import (
	"fmt"
)

type DeleteScalingGroupRequest struct {
	ScalingGroupId string `json:"scalingGroupId" url:"scalingGroupId"`
}

func (c *Client) DeleteScalingGroup(req *DeleteScalingGroupRequest) error {
	_, err := c.Do("scaling", "DeleteScalingGroup", req)
	if err != nil {
		return fmt.Errorf("[scaling:DeleteScalingGroup] request failed: %s", err)
	}

	return nil
}
