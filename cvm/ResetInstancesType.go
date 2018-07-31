package tcapi

import (
	"fmt"
)

type ResetInstancesTypeRequest struct {
	InstanceIds  []string `json:",omitempty" url:",omitempty,dotnumbered"`
	InstanceType string   `json:",omitempty" url:",omitempty"`
	ForceStop    bool     `json:",omitempty" url:",omitempty"`
}

func (c *Client) ResetInstancesType(req *ResetInstancesTypeRequest) error {
	_, err := c.Do("cvm", "ResetInstancesType", req)
	if err != nil {
		return fmt.Errorf("[cvm:ResetInstancesType] request failed: %s", err)
	}

	return nil
}
