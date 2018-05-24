package tcapi

import (
	"fmt"
)

type ResetInstancesInternetMaxBandwidthRequest struct {
	InstanceIds        []string           `json:",omitempty" url:",omitempty,dotnumbered"`
	InternetAccessible InternetAccessible `json:",omitempty" url:",omitempty,dotnumbered"`
	StartTime          bool               `json:",omitempty" url:",omitempty"`
	EndTime            bool               `json:",omitempty" url:",omitempty"`
}

func (c *Client) ResetInstancesInternetMaxBandwidth(req *ResetInstancesInternetMaxBandwidthRequest) error {
	_, err := c.Do("cvm", "ResetInstancesInternetMaxBandwidth", req)
	if err != nil {
		return fmt.Errorf("[cvm:ResetInstancesInternetMaxBandwidth] request failed: %s", err)
	}

	return nil
}
