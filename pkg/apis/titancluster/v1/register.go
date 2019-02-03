package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	localSchemeBuilder = &SchemeBuilder
	AddToScheme = SchemeBuilder.AddToScheme
	groupName = "foo.com"
)

var SchemeGroupVersion = schema.GroupVersion{Group: groupName, Version: "v1"}

func init() {
	localSchemeBuilder.Register(addKnownTypes)
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &TitanCluster{}, &TitanClusterList{},)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}