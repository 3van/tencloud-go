package tcapi

import (
	"fmt"
)

type RebootInstancesRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	ForceReboot bool     `json:",omitempty" url:",omitempty"`
}

func (c *Client) RebootInstances(req *RebootInstancesRequest) error {
	_, err := c.Do("cvm", "RebootInstances", req)
	if err != nil {
		return fmt.Errorf("[cvm:RebootInstances] request failed: %s", err)
	}

	return nil
}
