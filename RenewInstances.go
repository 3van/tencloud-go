package tcapi

import (
	"fmt"
)

type RenewInstancesRequest struct {
	InstanceIds           []string              `json:",omitempty" url:",omitempty,dotnumbered"`
	InstanceChargePrepaid InstanceChargePrepaid `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) RenewInstances(req *RenewInstancesRequest) error {
	_, err := c.Do("cvm", "RenewInstances", req)
	if err != nil {
		return fmt.Errorf("[cvm:RenewInstances] request failed: %s", err)
	}

	return nil
}
