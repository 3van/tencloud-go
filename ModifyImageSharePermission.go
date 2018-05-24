package tcapi

import (
	"fmt"
)

type ModifyImageSharePermissionRequest struct {
	ImageId    string   `json:",omitempty" url:",omitempty"`
	AccountIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	Permission string   `json:",omitempty" url:",omitempty"`
}

func (c *Client) ModifyImageSharePermission(req *ModifyImageSharePermissionRequest) error {
	_, err := c.Do("image", "ModifyImageSharePermission", req)
	if err != nil {
		return fmt.Errorf("[image:ModifyImageSharePermission] request failed: %s", err)
	}

	return nil
}
