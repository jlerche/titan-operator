package controller

import (
	"fmt"
	"github.com/jlerche/titan-operator/pkg/apis/titancluster/v1"
	apps "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/listers/apps/v1beta1"
	"k8s.io/client-go/tools/record"
	"strings"
)

type StatefulSetControlInterface interface {
	CreateStatefulSet(*v1.TitanCluster, *apps.StatefulSet) error
	UpdateStatefulSet(*v1.TitanCluster, *apps.StatefulSet) error
	DeleteStatefulSet(*v1.TitanCluster, *apps.StatefulSet) error
}

type defaultStatefulSetControl struct {
	kubeCli   kubernetes.Interface
	setLister v1beta1.StatefulSetLister
	recorder  record.EventRecorder
}

func NewRealStatefulSetControl(kubeCli kubernetes.Interface, setLister v1beta1.StatefulSetLister, recorder record.EventRecorder) *defaultStatefulSetControl {
	return &defaultStatefulSetControl{kubeCli, setLister, recorder,}
}

func (dss *defaultStatefulSetControl) CreateStatefulSet(tc *v1.TitanCluster, set *apps.StatefulSet) error {
	set, err := dss.kubeCli.AppsV1beta1().StatefulSets(tc.Namespace).Create(set)
	dss.recordStatefulSetEvent("create", tc, set, err)
	return err
}

func (dss *defaultStatefulSetControl) UpdateStatefulSet(tc *v1.TitanCluster, set *apps.StatefulSet) error {
	return nil
}

func (dss *defaultStatefulSetControl) DeleteStatefulSet(tc *v1.TitanCluster, set *apps.StatefulSet) error {
	return nil
}

func (dss *defaultStatefulSetControl) recordStatefulSetEvent(verb string, tc *v1.TitanCluster, set *apps.StatefulSet, err error) {
	tcName := tc.Name
	setName := set.Name
	if err == nil {
		reason := fmt.Sprintf("Successful%s", strings.Title(verb))
		message := fmt.Sprintf("%s StatefulSet %s in TitanCluster %s successful", strings.ToLower(verb), setName, tcName)
		dss.recorder.Event(tc, corev1.EventTypeNormal, reason, message)
	} else {
		reason := fmt.Sprintf("Failed%s", strings.Title(verb))
		message := fmt.Sprintf("%s StatefulSet %s in TitanCluster %s failed error: %s", strings.ToLower(verb), setName, tcName, err)
		dss.recorder.Event(tc, corev1.EventTypeWarning, reason, message)
	}
}

