package tcapi

import (
	"fmt"
)

type AssociateInstancesKeyPairsRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	KeyIds      []string `json:",omitempty" url:",omitempty,dotnumbered"`
	ForceStop   bool     `json:",omitempty" url:",omitempty"`
}

func (c *Client) AssociateInstancesKeyPairs(req *AssociateInstancesKeyPairsRequest) error {
	_, err := c.Do("cvm", "AssociateInstancesKeyPairs", req)
	if err != nil {
		return fmt.Errorf("[cvm:AssociateInstancesKeyPairs] request failed: %s", err)
	}

	return nil
}
