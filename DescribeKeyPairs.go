package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeKeyPairsRequest struct {
	KeyIds  []string `json:",omitempty" url:",omitempty,dotnumbered"`
	Filters []Filter `json:",omitempty" url:",omitempty,dotnumbered"`
	Offset  int      `json:",omitempty" url:",omitempty"`
	Limit   int      `json:",omitempty" url:",omitempty"`
}

type DescribeKeyPairsResponse struct {
	RequestId  string    `json:",omitempty" url:",omitempty"`
	TotalCount int       `json:",omitempty" url:",omitempty"`
	KeyPairSet []KeyPair `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) DescribeKeyPairs(req *DescribeKeyPairsRequest) (*DescribeKeyPairsResponse, error) {
	resp, err := c.Do("cvm", "DescribeKeyPairs", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeKeyPairs] request failed: %s", err)
	}

	ret := new(DescribeKeyPairsResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:DescribeKeyPairs] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:DescribeKeyPairs] response unmarshaled to nil")
	}

	return ret, nil
}
