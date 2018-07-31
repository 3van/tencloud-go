package tcapi

import (
	"encoding/json"
	"fmt"
)

type CreateKeyPairRequest struct {
	KeyName   string `json:",omitempty" url:",omitempty"`
	ProjectId int    `json:",omitempty" url:",omitempty"`
}

type CreateKeyPairResponse struct {
	RequestId string  `json:",omitempty" url:",omitempty"`
	KeyPair   KeyPair `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) CreateKeyPair(req *CreateKeyPairRequest) (*CreateKeyPairResponse, error) {
	resp, err := c.Do("cvm", "CreateKeyPair", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:CreateKeyPair] request failed: %s", err)
	}

	ret := new(CreateKeyPairResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:CreateKeyPair] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:CreateKeyPair] response unmarshaled to nil")
	}

	return ret, nil
}
