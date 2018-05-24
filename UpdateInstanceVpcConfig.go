package tcapi

import (
	"fmt"
)

type UpdateInstanceVpcConfigRequest struct {
	InstanceId          string              `json:",omitempty" url:",omitempty"`
	VirtualPrivateCloud VirtualPrivateCloud `json:",omitempty" url:",omitempty,dotnumbered"`
	ForceStop           bool                `json:",omitempty" url:",omitempty"`
}

func (c *Client) UpdateInstanceVpcConfig(req *UpdateInstanceVpcConfigRequest) error {
	_, err := c.Do("cvm", "UpdateInstanceVpcConfig", req)
	if err != nil {
		return fmt.Errorf("[cvm:UpdateInstanceVpcConfig] request failed: %s", err)
	}

	return nil
}
