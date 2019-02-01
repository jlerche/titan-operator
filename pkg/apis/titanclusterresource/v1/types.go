package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TitanCluster

type TitanCluster struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
}

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