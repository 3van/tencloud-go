package tcapi

import (
	"encoding/json"
)

type ErrorResponse struct {
	Code     string
	CodeDesc string
	Message  string
}

// this is maddening, but i don't want to take the easy way out and interface{} this
type OldErrorResponse struct {
	Code     int
	CodeDesc string
	Message  string
}

type V3ErrorResponse struct {
	Error ErrorResponse `url",dotnumbered"`
}

type V3BaseResponse struct {
	Response *json.RawMessage
}

// from https://cloud.tencent.com/document/api/213/9451

type Placement struct {
	Zone      string
	ProjectId int
	HostIds   []string `url",dotnumbered"`
}

type SystemDisk struct {
	DiskType string
	DiskId   string
	DiskSize int
}

type DataDisk struct {
	DiskType string
	DiskId   string
	DiskSize int
}

type VirtualPrivateCloud struct {
	VpcId              string
	SubnetId           string
	AsVpcGateway       bool
	PrivateIpAddresses []string `url",dotnumbered"`
}

type InternetAccessible struct {
	InternetChargeType      string
	InternetMaxBandwidthOut int
	PublicIpAssigned        string
}

type InstanceChargePrepaid struct {
	Period    int
	RenewFlag string
}

type LoginSettings struct {
	Password       string
	KeyIds         []string `url",dotnumbered"`
	KeepImageLogin string
}

type RunSecurityServiceEnabled struct {
	Enabled string
}

type RunMonitorServiceEnabled struct {
	Enabled string
}

type EnhancedService struct {
	SecurityService RunSecurityServiceEnabled `url",dotnumbered"`
	MonitorService  RunMonitorServiceEnabled  `url",dotnumbered"`
}

type ItemPrice struct {
	UnitPrice     float32
	ChargeUnit    string
	OriginalPrice float32
	DiscountPrice float32
}

type Price struct {
	InstancePrice  ItemPrice `url",dotnumbered"`
	BandwidthPrice ItemPrice `url",dotnumbered"`
}

type Filter struct {
	Name   string
	Values []string `url",dotnumbered"`
}

type InstanceStatus struct {
	InstanceId    string
	InstanceState string
}

type Instance struct {
	Placement           Placement `url",dotnumbered"`
	InstanceId          string
	InstanceType        string
	CPU                 int
	Memory              int
	RestrictState       string
	InstanceName        string
	InstanceChargeType  string
	SystemDisk          SystemDisk          `url",dotnumbered"`
	DataDisks           []DataDisk          `url",dotnumbered"`
	PrivateIpAddresses  []string            `url",dotnumbered"`
	PublicIpAddresses   []string            `url",dotnumbered"`
	InternetAccessible  InternetAccessible  `url",dotnumbered"`
	VirtualPrivateCloud VirtualPrivateCloud `url",dotnumbered"`
	ImageId             string
	RenewFlag           string
	CreatedTime         string
	ExpiredTime         string
}

type InstanceTypeConfig struct {
	Zone              string
	InstanceType      string
	InstanceFamily    string
	GPU               int
	CPU               int
	Memory            int
	CbsSupport        string
	InstanceTypeState string
}

type ImageSharedAccount struct {
	ImageId   string
	AccountId string
}

type Quota struct {
	QuotaId      string
	QuotaCurrent int
	QuotaLimit   int
}

type Image struct {
	ImageId          string
	OsName           string
	ImageSize        string
	ImageType        int
	CreatedTime      string
	ImageState       string
	ImageName        string
	ImageDescription string
	ImageSource      string
	ImageCreator     string
}

type AvailabilityZone struct {
	RegionId  string
	Zone      string
	ZoneName  string
	ZoneState string
}

type KeyPair struct {
	KeyId                 string
	KeyName               string
	ProjectId             string
	Description           string
	PublicKey             string
	PrivateKey            string
	AssociatedInstanceIds []string `url",dotnumbered"`
	CreatedTime           string
}

type KeyPairInstances struct {
	KeyId                   string
	AssociatedInstanceIdSet []string `url",dotnumbered"`
}

type Address struct {
	AddressId        string
	AddressName      string
	AddressState     string
	AddressIp        string
	BindedResourceId string
	CreatedTime      string
}

type InstanceChargeTypeConfig struct {
	InstanceChargeType string
	Description        string
}

type InternetChargeTypeConfig struct {
	InternetChargeType string
	Description        string
}

type InternetBandwidthConfig struct {
	StartTime          string
	EndTime            string
	InternetAccessible InternetAccessible `url",dotnumbered"`
}

type ScalingConfiguration struct {
	ScalingConfigurationId   string              `json:"scalingConfigurationId,omitempty" url:"scalingConfigurationId,omitempty"`
	ScalingConfigurationName string              `json:"scalingConfigurationName" url:"scalingConfigurationName"`
	ScalingGroupIdSet        []ScalingGroupIdSet `json:"scalingGroupIdSet,omitempty" url:"scalingGroupIdSet,dotnumbered,omitempty"`
	CPU                      int                 `json:"cpu" url:"cpu"`
	Memory                   int                 `json:"mem" url:"mem"`
	ImageType                int                 `json:"imageType" url:"imageType"`
	ImageId                  string              `json:"imageId" url:"imageId"`
	StorageType              int                 `json:"storageType" url:"storageType"`
	StorageSize              int                 `json:"storageSize" url:"storageSize"`
	RootSize                 int                 `json:"rootSize" url:"rootSize"`
	BandwidthType            string              `json:"bandwidthType" url:"bandwidthType"`
	Bandwidth                int                 `json:"bandwidth" url:"bandwidth"`
	WanIp                    int                 `json:"wanIp" url:"wanIp"`
	KeyId                    string              `json:"keyId" url:"keyId"`
	Password                 string              `json:"password" url:"password"`
	SgSet                    []SecurityGroupSet  `json:"sgSet" url:"sgSet,dotnumbered"`
	NeedMonitorAgent         int                 `json:"needMonitorAgent" url:"needMonitorAgent"`
	NeedSecurityAgent        int                 `json:"needSecurityAgent" url:"needSecurityAgent"`
	CreateTime               string              `json:"createTime,omitempty" url:"createTime,omitempty"`
	ProjectId                int                 `json:"projectId" url:"projectId"`
	AscStatus                int                 `json:"ascStatus" url:"ascStatus"`
	DataSnapshotId           string              `json:"dataSnapshotId" url:"dataSnapshotId"`
	ItemId                   int                 `json:"itemId" url:"itemId"`
	KeepImageSetting         int                 `json:"keepImageSetting" url:"keepImageSetting"`
	Type                     string              `json:"type" url:"type"`
	UserData                 string              `json:"userdata" url:"userdata"`
}

type ScalingGroupIdSet struct {
	ScalingGroupId   string `json:"scalingGroupId" url:"scalingGroupId"`
	ScalingGroupName string `json:"scalingGroupName" url:"scalingGroupName"`
}

type SecurityGroupSet struct {
	SgId string `json:"sgId" url:"sgId"`
}

type ScalingGroup struct {
	ScalingGroupId           string              `json:"scalingGroupId,omitempty" url:"scalingGroupId,omitempty"`
	ScalingGroupName         string              `json:"scalingGroupName" url:"scalingGroupName"`
	ScalingConfigurationId   string              `json:"scalingConfigurationId" url:"scalingConfigurationId"`
	ScalingConfigurationName string              `json:"scalingConfigurationName,omitempty" url:"scalingConfigurationName,omitempty"`
	MinSize                  int                 `json:"minSize" url:"minSize"`
	MaxSize                  int                 `json:"maxSize" url:"maxSize"`
	CreateTime               string              `json:"createTime,omitempty" url:"createTime,omitempty"`
	InstanceNum              int                 `json:"instanceNum" url:"instanceNum"`
	RemovePolicy             string              `json:"removePolicy" url:"removePolicy"`
	LoadBalancerIdSet        []LoadBalancerIdSet `json:"loadBalancerIdSet" url:"loadBalancerIdSet,dotnumbered"`
	LoadBalancerIds          []string            `json:"loadBalancerIds,omitempty" url:"loadBalancerIds,omitempty,dotnumbered"`
	VpcId                    string              `json:"vpcId" url:"vpcId"`
	SubnetIdSet              []SubnetIdSet       `json:"subnetIdSet" url:"subnetIdSet,dotnumbered"`
	SubnetIds                []string            `json:"subnetIds,omitempty" url:"subnetIds,omitempty,dotnumbered"`
	ZoneIdSet                []ZoneIdSet         `json:"zoneIdSet" url:"zoneIdSet,dotnumbered"`
	ZoneIds                  []string            `json:"zoneIds,omitempty" url:"zoneIds,omitempty,dotnumbered"`
	ProjectId                int                 `json:"projectId" url:"projectId"`
	BInScalingActivity       int                 `json:"bInScalingActivity" url:"bInScalingActivity"`
	DesiredCapacity          int                 `json:"desiredCapacity" url:"desiredCapacity"`
}

type LoadBalancerIdSet struct {
	Status         int    `json:"status" url:"status"`
	LoadBalancerId string `json:"loadBalancerId" url:"loadBalancerId"`
	Owner          string `json:"owner" url:"owner"`
	ZoneId         int    `json:"zoneId" url:"zoneId"`
}

type SubnetIdSet struct {
	Status   int    `json:"status" url:"status"`
	SubnetId string `json:"subnetId" url:"subnetId"`
	Owner    string `json:"owner" url:"owner"`
	ZoneId   int    `json:"zoneId" url:"zoneId"`
}

type ZoneIdSet struct {
	Status int    `json:"status" url:"status"`
	Owner  string `json:"owner" url:"owner"`
	ZoneId int    `json:"zoneId" url:"zoneId"`
}

type ScalingInstance struct {
	InstanceId           string `json:"instanceId" url:"instanceId"`
	HealthStatus         string `json:"healthStatus" url:"healthStatus"`
	CreationType         string `json:"creationType" url:"creationType"`
	LifeCycleState       string `json:"lifeCycleState" url:"lifeCycleState"`
	ProtectedFromScaleIn int    `json:"protectedFromScaleIn" url:"protectedFromScaleIn"`
	AddTime              string `json:"addTime" url:"addTime"`
}

type SharePermission struct {
	CreatedTime string
	Account     string
}
