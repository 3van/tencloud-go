package tcapi

import (
	"encoding/json"
	"fmt"
)

type DescribeScalingConfigurationRequest struct {
	ScalingConfigurationIds  []string `json:"scalingConfigurationIds,omitempty" url:"scalingConfigurationIds,omitempty,dotnumbered"`
	ScalingConfigurationName string   `json:"scalingConfigurationName,omitempty" url:"scalingConfigurationName,omitempty"`
	Offset                   int      `json:"offset,omitempty" url:"offset,omitempty"`
	Limit                    int      `json:"limit,omitempty" url:"limit,omitempty"`
	ProjectId                int      `json:"projectId,omitempty" url:"projectId,omitempty"`
}

type DescribeScalingConfigurationResponse struct {
	Data DescribeScalingConfigurationResponseData
}

type DescribeScalingConfigurationResponseData struct {
	TotalCount              int
	ScalingConfigurationSet []ScalingConfiguration `url:",dotnumbered"`
}

func (c *Client) DescribeScalingConfiguration(req *DescribeScalingConfigurationRequest) (*DescribeScalingConfigurationResponseData, error) {
	ret := new(DescribeScalingConfigurationResponse)
	resp, err := c.Do("scaling", "DescribeScalingConfiguration", req)
	if err != nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingConfiguration] request failed: %s", err)
	}

	err = json.Unmarshal(*resp, ret)

	if err != nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingConfiguration] response unmarshal failed: %s", err)
	}
	if ret == nil {
		return nil, fmt.Errorf("[scaling:DescribeScalingConfiguration] response unmarshaled to nil")
	}

	return &ret.Data, nil
}
