package titancluster

import "github.com/jlerche/titan-operator/pkg/apis/titancluster/v1"

type StatusUpdaterInterface interface {
	UpdateTitanClusterStatus(*v1.TitanCluster, *v1.TitanClusterStatus) error
}