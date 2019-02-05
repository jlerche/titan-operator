package titancluster

import (
	"github.com/jlerche/titan-operator/pkg/apis/titancluster/v1"
	"github.com/jlerche/titan-operator/pkg/controller"
	apps "k8s.io/api/apps/v1beta1"
	"k8s.io/client-go/tools/record"
	"time"
)

const (
	pdConnTimeout = 2 * time.Second
)

type ControlInterface interface {
	UpdateTitanCluster(*v1.TitanCluster, []*apps.StatefulSet) error
}

func NewDefaultTitanClusterControl(setControl controller.StatefulSetControlInterface, statusUpdater StatusUpdaterInterface, recorder record.EventRecorder) ControlInterface {
	return &defaultTitanClusterClusterControl{setControl, statusUpdater, recorder,}
}

type defaultTitanClusterClusterControl struct {
	setControl controller.StatefulSetControlInterface
	statusUpdater StatusUpdaterInterface
	recorder record.EventRecorder
}

func (tcc *defaultTitanClusterClusterControl) UpdateTitanCluster(tc *v1.TitanCluster, sets []*apps.StatefulSet) error {
	status, err := tcc
}

func (tcc *defaultTitanClusterClusterControl) updateTitanCluster(tc *v1.TitanCluster, sets []*apps.StatefulSet) (*v1.TitanClusterStatus, error) {
	status := &tc.Status
	newPDSet, err := tcc.getNewPDSetFor
}

