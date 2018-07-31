package tcapi

import (
	"encoding/json"
	"fmt"
)

// it is *profoundly* annoying how similar-but-sufficiently-different
// this is to an Instance struct
type RunInstancesRequest struct {
	InstanceChargeType    string                `json:",omitempty" url:",omitempty"`
	InstanceChargePrepaid InstanceChargePrepaid `json:",omitempty" url:",omitempty,dotnumbered"`
	Placement             Placement             `json:",omitempty" url:",omitempty,dotnumbered"`
	InstanceType          string                `json:",omitempty" url:",omitempty"`
	ImageId               string                `json:",omitempty" url:",omitempty"`
	SystemDisk            SystemDisk            `json:",omitempty" url:",omitempty,dotnumbered"`
	DataDisks             []DataDisk            `json:",omitempty" url:",omitempty,dotnumbered"`
	VirtualPrivateCloud   VirtualPrivateCloud   `json:",omitempty" url:",omitempty,dotnumbered"`
	InternetAccessible    InternetAccessible    `json:",omitempty" url:",omitempty,dotnumbered"`
	InstanceCount         int                   `json:",omitempty" url:",omitempty"`
	InstanceName          string                `json:",omitempty" url:",omitempty"`
	LoginSettings         LoginSettings         `json:",omitempty" url:",omitempty,dotnumbered"`
	SecurityGroupIds      []string              `json:",omitempty" url:",omitempty,dotnumbered"`
	EnhancedService       EnhancedService       `json:",omitempty" url:",omitempty,dotnumbered"`
	ClientToken           string                `json:",omitempty" url:",omitempty"`
	UserData              string                `json:",omitempty" url:",omitempty"`
}

type RunInstancesResponse struct {
	RequestId     string   `json:",omitempty" url:",omitempty"`
	InstanceIdSet []string `json:",omitempty" url:",omitempty,dotnumbered"`
}

func (c *Client) RunInstances(req *RunInstancesRequest) (*RunInstancesResponse, error) {
	resp, err := c.Do("cvm", "RunInstances", req)
	if err != nil {
		return nil, fmt.Errorf("[cvm:RunInstances] request failed: %s", err)
	}

	ret := new(RunInstancesResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[cvm:RunInstances] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[cvm:RunInstances] response unmarshaled to nil")
	}

	return ret, nil
}
