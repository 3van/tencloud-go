package tcapi

import (
	"fmt"
)

type DisassociateInstancesKeyPairsRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	KeyIds      []string `json:",omitempty" url:",omitempty,dotnumbered"`
	ForceStop   bool     `json:",omitempty" url:",omitempty"`
}

func (c *Client) DisassociateInstancesKeyPairs(req *DisassociateInstancesKeyPairsRequest) error {
	_, err := c.Do("cvm", "DisassociateInstancesKeyPairs", req)
	if err != nil {
		return fmt.Errorf("[cvm:DisassociateInstancesKeyPairs] request failed: %s", err)
	}

	return nil
}
