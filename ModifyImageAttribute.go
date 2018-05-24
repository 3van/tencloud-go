package tcapi

import (
	"fmt"
)

type ModifyImageAttributeRequest struct {
	ImageId          string `json:",omitempty" url:",omitempty"`
	ImageName        string `json:",omitempty" url:",omitempty"`
	ImageDescription string `json:",omitempty" url:",omitempty"`
}

func (c *Client) ModifyImageAttribute(req *ModifyImageAttributeRequest) error {
	_, err := c.Do("image", "ModifyImageAttribute", req)
	if err != nil {
		return fmt.Errorf("[image:ModifyImageAttribute] request failed: %s", err)
	}

	return nil
}
