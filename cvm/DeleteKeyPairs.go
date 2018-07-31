package tcapi

import (
	"fmt"
)

type DeleteKeyPairsRequest struct {
	KeyIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) DeleteKeyPairs(req *DeleteKeyPairsRequest) error {
	_, err := c.Do("cvm", "DeleteKeyPairs", req)
	if err != nil {
		return fmt.Errorf("[cvm:DeleteKeyPairs] request failed: %s", err)
	}

	return nil
}
