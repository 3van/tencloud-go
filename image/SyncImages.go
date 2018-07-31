package tcapi

import (
	"fmt"
)

type SyncImagesRequest struct {
	ImageIds           []string `json:",omitempty" url:",omitempty,dotnumbered"`
	DestinationRegions []string `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) SyncImages(req *SyncImagesRequest) error {
	_, err := c.Do("image", "SyncImages", req)
	if err != nil {
		return fmt.Errorf("[image:SyncImages] request failed: %s", err)
	}

	return nil
}
