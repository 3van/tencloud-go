package tcapi

import (
	"fmt"
)

type ResizeInstanceDisksRequest struct {
	InstanceId string     `json:",omitempty" url:",omitempty"`
	DataDisks  []DataDisk `json:",omitempty" url:",omitempty,dotnumbered"`
	ForceStop  bool       `json:",omitempty" url:",omitempty"`
}

func (c *Client) ResizeInstanceDisks(req *ResizeInstanceDisksRequest) error {
	_, err := c.Do("cvm", "ResizeInstanceDisks", req)
	if err != nil {
		return fmt.Errorf("[cvm:ResizeInstanceDisks] request failed: %s", err)
	}

	return nil
}
