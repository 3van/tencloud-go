package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeImagesRequest struct {
	ImageIds []string `json:",omitempty" url:",omitempty,dotnumbered"`
	Filters  []Filter `json:",omitempty" url:",omitempty,dotnumbered"`
	Offset   int      `json:",omitempty" url:",omitempty"`
	Limit    int      `json:",omitempty" url:",omitempty"`
}

type DescribeImagesResponse struct {
	RequestId  string  `json:",omitempty" url:",omitempty"`
	ImageSet   []Image `json:",omitempty" url:",omitempty,dotnumbered"`
	TotalCount int     `json:",omitempty" url:",omitempty"`
}

func (c *Client) DescribeImages(req *DescribeImagesRequest) (*DescribeImagesResponse, error) {
	resp, err := c.Do("image", "DescribeImages", req)
	if err != nil {
		return nil, fmt.Errorf("[image:DescribeImages] request failed: %s", err)
	}

	ret := new(DescribeImagesResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[image:DescribeImages] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[image:DescribeImages] response unmarshaled to nil")
	}

	return ret, nil
}
