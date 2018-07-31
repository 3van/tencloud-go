package tcapi

import (
	"fmt"
)

type ModifyInstancesAttributeRequest struct {
	InstanceIds  []string `json:",omitempty" url:",omitempty,dotnumbered"`
	InstanceName string   `json:",omitempty" url:",omitempty"`
}

func (c *Client) ModifyInstancesAttribute(req *ModifyInstancesAttributeRequest) error {
	_, err := c.Do("cvm", "ModifyInstancesAttribute", req)
	if err != nil {
		return fmt.Errorf("[cvm:ModifyInstancesAttribute] request failed: %s", err)
	}

	return nil
}
