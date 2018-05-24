package tcapi

import (
	"fmt"
)

type TerminateInstancesRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) TerminateInstances(req *TerminateInstancesRequest) error {
	_, err := c.Do("cvm", "TerminateInstances", req)
	if err != nil {
		return fmt.Errorf("[cvm:TerminateInstances] request failed: %s", err)
	}

	return nil
}
