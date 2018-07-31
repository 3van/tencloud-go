package tcapi

import (
	"fmt"
)

type StartInstancesRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) StartInstances(req *StartInstancesRequest) error {
	_, err := c.Do("cvm", "StartInstances", req)
	if err != nil {
		return fmt.Errorf("[cvm:StartInstances] request failed: %s", err)
	}

	return nil
}
