package tcapi

import (
	"fmt"
)

type ResetInstanceRequest struct {
	InstanceId      string          `json:",omitempty" url:",omitempty"`
	ImageId         string          `json:",omitempty" url:",omitempty"`
	SystemDisk      SystemDisk      `json:",omitempty" url:",omitempty,dotnumbered"`
	LoginSettings   LoginSettings   `json:",omitempty" url:",omitempty,dotnumbered"`
	EnhancedService EnhancedService `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) ResetInstance(req *ResetInstanceRequest) error {
	_, err := c.Do("cvm", "ResetInstance", req)
	if err != nil {
		return fmt.Errorf("[cvm:ResetInstance] request failed: %s", err)
	}

	return nil
}
