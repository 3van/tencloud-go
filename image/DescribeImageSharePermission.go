package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeImageSharePermissionRequest struct {
	ImageId string `json:",omitempty" url:",omitempty"`
}

type DescribeImageSharePermissionResponse struct {
	RequestId          string            `json:",omitempty" url:",omitempty"`
	SharePermissionSet []SharePermission `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) DescribeImageSharePermission(req *DescribeImageSharePermissionRequest) (*DescribeImageSharePermissionResponse, error) {
	resp, err := c.Do("image", "DescribeImageSharePermission", req)
	if err != nil {
		return nil, fmt.Errorf("[image:DescribeImageSharePermission] request failed: %s", err)
	}

	ret := new(DescribeImageSharePermissionResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[image:DescribeImageSharePermission] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[image:DescribeImageSharePermission] response unmarshaled to nil")
	}

	return ret, nil
}
