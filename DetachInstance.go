package tcapi

import (
	"fmt"
)

type DetachInstanceRequest struct {
	ScalingGroupId string   `json:"scalingGroupId,omitempty" url:"scalingGroupId,omitempty"`
	InstanceIds    []string `json:"instanceIds,omitempty" url:"instanceIds,omitempty,dotnumbered"`
}

func (c *Client) DetachInstance(req *DetachInstanceRequest) error {
	_, err := c.Do("scaling", "DetachInstance", req)
	if err != nil {
		return fmt.Errorf("[scaling:DetachInstance] request failed: %s", err)
	}

	return nil
}
