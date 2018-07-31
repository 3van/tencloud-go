package tcapi

import (
	"fmt"
)

type DeleteImagesRequest struct {
	ImageIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) DeleteImages(req *DeleteImagesRequest) error {
	_, err := c.Do("image", "DeleteImages", req)
	if err != nil {
		return fmt.Errorf("[image:DeleteImages] request failed: %s", err)
	}

	return nil
}
