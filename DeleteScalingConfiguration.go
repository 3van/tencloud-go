package tcapi

import (
	"fmt"
)

type DeleteScalingConfigurationRequest struct {
	ScalingConfigurationId string `json:"scalingConfigurationId" url:"scalingConfigurationId"`
}

func (c *Client) DeleteScalingConfiguration(req *DeleteScalingConfigurationRequest) error {
	_, err := c.Do("scaling", "DeleteScalingConfiguration", req)
	if err != nil {
		return fmt.Errorf("[scaling:DeleteScalingConfiguration] request failed: %s", err)
	}

	return nil
}
