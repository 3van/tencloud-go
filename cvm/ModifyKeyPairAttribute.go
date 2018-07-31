package tcapi

import (
	"fmt"
)

type ModifyKeyPairAttributeRequest struct {
	KeyId       string `json:",omitempty" url:",omitempty"`
	KeyName     string `json:",omitempty" url:",omitempty"`
	Description string `json:",omitempty" url:",omitempty"`
}

func (c *Client) ModifyKeyPairAttribute(req *ModifyKeyPairAttributeRequest) error {
	_, err := c.Do("cvm", "ModifyKeyPairAttribute", req)
	if err != nil {
		return fmt.Errorf("[cvm:ModifyKeyPairAttribute] request failed: %s", err)
	}

	return nil
}
