package tcapi

import (
	"fmt"
)

type ModifyInstancesProjectRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	ProjectId   int      `json:",omitempty" url:",omitempty"`
}

func (c *Client) ModifyInstancesProject(req *ModifyInstancesProjectRequest) error {
	_, err := c.Do("cvm", "ModifyInstancesProject", req)
	if err != nil {
		return fmt.Errorf("[cvm:ModifyInstancesProject] request failed: %s", err)
	}

	return nil
}
