package v1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TitanCluster
type TitanCluster struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec TitanClusterSpec `json:"spec"`
	Status TitanClusterStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TitanClusterList
type TitanClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items []TitanCluster `json:"items"`
}

type TitanClusterSpec struct {
	PD PDSpec `json:"pd,omitempty"`
	Titan TitanSpec `json:"titan,omitempty"`
	TiKV TiKVSpec `json:"tikv,omitempty"`
	Service string `json:"service,omitempty"`
	Config map[string]string `json:"config,omitempty"`
	Monitor *MonitorSpec `json:"monitor,omitempty"`
	// privileged
	Services []Service `json:"services,omitempty"`
	ConfigMap string `json:"configMap,omitempty"`
	Paused bool `json:"paused,omitempty"`
	State ClusterState `json:"state,omitempty"`
	RetentionDuration string `json:"retentionDuration,omitempty"`
}

type TitanClusterStatus struct {
	PDStatus PDStatus `json:"pdStatus,omitempty"`
	TiKVStatus TiKVStatus `json:"tikvStatus,omitempty"`
	TitanStatus TitanStatus `json:"titanStatus,omitempty"`
}

type NormalStatus struct {
	ReallyNormal bool `json:"reallyNormal,omitempty"`
}

type GracefulDeletedStatus struct {
	ReallyGracefulDeleted bool `json:"reallyGracefulDeleted,omitempty"`
	RetentionTime *metav1.Time `json:"retentionTime,omitempty"`
}

type RestoreStatus struct {
	ReallyRestoring bool `json:"reallyRestoring,omitempty"`
	BeginTime *metav1.Time `json:"beginTime,omitempty"`
	EndTime *metav1.Time `json:"endTime,omitempty"`
}

type ClusterState string

const (
	StateNormal ClusterState = ""
	StateGracefulDeleted = "GracefulDeleted"
	StateRestore = "Restore"
)

type PDSpec struct {
	ContainerSpec
	Size int32 `json:"size"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	NodeSelectorRequired bool `json:"nodeSelectorRequired,omitempty"`
}

type TitanSpec struct {
	ContainerSpec
	Size int32 `json:"size"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	NodeSelectorRequired bool `json:"nodeSelectorRequired,omitempty"`
}

type TiKVSpec struct {
	ContainerSpec
	Size int32 `json:"size"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	NodeSelectorRequired bool `json:"nodeSelectorRequired,omitempty"`
}

// privileged spec TBI

type MonitorSpec struct {
	Prometheus ContainerSpec `json:"prometheus,omitempty"`
	RetentionDays int32 `json:"retentionDays,omitempty"`
	Grafana *ContainerSpec `json:"grafana,omitempty"`
	// Dashboard installer probably not necessary here
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	NodeSelectorRequired bool `json:"nodeSelectorRequired,omitempty"`
}

type ContainerSpec struct {
	Image string `json:"image"`
	Requests *ResourceRequirement `json:"requests,omitempty"`
	Limits *ResourceRequirement `json:"limits,omitempty"`
}

type Service struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type ResourceRequirement struct {
	CPU string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Storage string `json:"storage,omitempty"`
}

const (
	AnnotationStorageSize string = "storage.titancluster/size"
	TitanVolumeName string = "titan-volume-hostpath"
	TitanSetResourcePlural string = "titansets"
	TitanReviewResourcePlural string = "titanreviews"
	TiKVStateUp string = "Up"
)

type ContainerType string

const (
	PDContainerType ContainerType = "pd"
	TitanContainerType ContainerType = "titan"
	TiKVContainerType ContainerType = "tikv"
	// TiDB has a Prometheus pushgateway definition here
	UnknownContainerType ContainerType = "unknown"
)

func (ct ContainerType) String() string {
	return string(ct)
}

type PDStatus struct {
	StatefulSet PDStatefulSetStatus `json:"statefulSet,omitempty"`
	Members map[string]PDMember `json:"members,omitempty"`
	Upgrading bool `json:"upgrading,omitempty"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

type PDStatefulSetStatus struct {
	Name string `json:"name,omitempty"`
}

type PDMember struct {
	Name string `json:"name"`
	ID string `json:"id"`
	IP string `json:"ip"`
}

type TitanStatus struct {
	Members map[string]TitanMember `json:"members,omitempty"`
	Upgrading bool `json:"upgrading,omitempty"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

type TitanMember struct {
	IP string `json:"ip"`
}

type TiKVStatus struct {
	Stores map[string]TiKVStores `json:"stores,omitempty"`
	Upgrading bool `json:"upgrading,omitempty"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

type TiKVStores struct {
	ID string `json:"id"`
	IP string `json:"ip"`
	State string `json:"state"`
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime"`
}

func (tc *TitanCluster) GetConfigMapName() string {
	if tc.Spec.ConfigMap != "" {
		return tc.Spec.ConfigMap
	}

	return fmt.Sprintf("%s-config", tc.GetName())
}