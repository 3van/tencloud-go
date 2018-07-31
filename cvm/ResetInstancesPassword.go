package tcapi

import (
	"fmt"
)

type ResetInstancesPasswordRequest struct {
	InstanceIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	Password    string   `json:",omitempty" url:",omitempty"`
	UserName    string   `json:",omitempty" url:",omitempty"`
	ForceStop   bool     `json:",omitempty" url:",omitempty"`
}

func (c *Client) ResetInstancesPassword(req *ResetInstancesPasswordRequest) error {
	_, err := c.Do("cvm", "ResetInstancesPassword", req)
	if err != nil {
		return fmt.Errorf("[cvm:ResetInstancesPassword] request failed: %s", err)
	}

	return nil
}
