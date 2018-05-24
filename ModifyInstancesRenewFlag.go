package tcapi

import (
	"fmt"
)

type ModifyInstancesRenewFlagRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	RenewFlag   string   `json:",omitempty" url:",omitempty"`
}

func (c *Client) ModifyInstancesRenewFlag(req *ModifyInstancesRenewFlagRequest) error {
	_, err := c.Do("cvm", "ModifyInstancesRenewFlag", req)
	if err != nil {
		return fmt.Errorf("[cvm:ModifyInstancesRenewFlag] request failed: %s", err)
	}

	return nil
}
