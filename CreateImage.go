package tcapi

import (
	"fmt"
)

type CreateImageRequest struct {
	InstanceId       string `json:",omitempty" url:",omitempty"`
	ImageName        string `json:",omitempty" url:",omitempty"`
	ImageDescription string `json:",omitempty" url:",omitempty,encoded"`
}

func (c *Client) CreateImage(req *CreateImageRequest) error {
	_, err := c.Do("image", "CreateImage", req)
	if err != nil {
		return fmt.Errorf("[image:CreateImage] request failed: %s", err)
	}

	return nil
}
