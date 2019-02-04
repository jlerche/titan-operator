package titancluster

import (
	"github.com/jlerche/titan-operator/pkg/apis/titancluster/v1"
	"github.com/jlerche/titan-operator/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
)

var controllerKind = v1.SchemeGroupVersion.WithKind("TitanCluster")

type Controller struct {
	kubeClient kubernetes.Interface
	cli versioned.Interface
}