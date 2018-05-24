package tcapi

import (
	"encoding/json"
	"fmt"
)

// it is *profoundly* annoying how similar-but-sufficiently-different
// this is to a ScalingConfiguration struct
type CreateScalingConfigurationRequest struct {
	ScalingConfigurationName string `json:"scalingConfigurationName" url:"scalingConfigurationName"`
	ImageId                  string `json:"imageId" url:"imageId"`
	CPU                      int    `json:"cpu" url:"cpu"`
	Memory                   int    `json:"mem" url:"mem"`
	StorageType              int    `json:"storageType" url:"storageType"`
	StorageSize              int    `json:"storageSize" url:"storageSize"`
	BandwidthType            string `json:"bandwidthType" url:"bandwidthType"`
	Bandwidth                int    `json:"bandwidth" url:"bandwidth"`
	ImageType                int    `json:"imageType" url:"imageType"`
	RootSize                 int    `json:"rootSize" url:"rootSize"`
	KeyId                    string `json:"keyId,omitempty" url:"keyId,omitempty"`
	Password                 string `json:"password,omitempty" url:"password,omitempty"`
	NeedMonitorAgent         int    `json:"needMonitorAgent" url:"needMonitorAgent"`
	NeedSecurityAgent        int    `json:"needSecurityAgent" url:"needSecurityAgent"`
	WanIp                    int    `json:"wanIp" url:"wanIp"`
	SgId                     string `json:"sgId,omitempty" url:"sgId,omitempty"`
	ProjectId                int    `json:"projectId,omitempty" url:"projectId,omitempty"`
	DataSnapshotId           string `json:"dataSnapshotId,omitempty" url:"dataSnapshotId,omitempty"`
	CvmType                  int    `json:"cvmType,omitempty" url:"cvmType,omitempty"`
	UserData                 string `json:"userdata,omitempty" url:"userdata,omitempty"`
}

type CreateScalingConfigurationResponse struct {
	Data CreateScalingConfigurationResponseData `json:"data" url:"data"`
}

type CreateScalingConfigurationResponseData struct {
	ScalingConfigurationIdSet []string `json:"scalingConfigurationIdSet" url:"scalingConfigurationIdSet"`
}

func (c *Client) CreateScalingConfiguration(req *CreateScalingConfigurationRequest) (*[]string, error) {
	resp, err := c.Do("scaling", "CreateScalingConfiguration", req)
	if err != nil {
		return nil, fmt.Errorf("[scaling:CreateScalingConfiguration] request failed: %s", err)
	}

	ret := new(CreateScalingConfigurationResponse)
	err = json.Unmarshal(*resp, ret)
	if err != nil {
		return nil, fmt.Errorf("[scaling:CreateScalingConfiguration] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[scaling:CreateScalingConfiguration] response unmarshaled to nil")
	}

	return &ret.Data.ScalingConfigurationIdSet, nil
}
