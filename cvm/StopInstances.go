package tcapi

import (
	"fmt"
)

type StopInstancesRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	ForceStop   bool     `json:",omitempty" url:",omitempty"`
}

func (c *Client) StopInstances(req *StopInstancesRequest) error {
	_, err := c.Do("cvm", "StopInstances", req)
	if err != nil {
		return fmt.Errorf("[cvm:StopInstances] request failed: %s", err)
	}

	return nil
}
