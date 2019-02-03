/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	titanclusterv1 "github.com/jlerche/titan-operator/pkg/apis/titancluster/v1"
	versioned "github.com/jlerche/titan-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jlerche/titan-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jlerche/titan-operator/pkg/client/listers/titancluster/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TitanClusterInformer provides access to a shared informer and lister for
// TitanClusters.
type TitanClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TitanClusterLister
}

type titanClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTitanClusterInformer constructs a new informer for TitanCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTitanClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTitanClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTitanClusterInformer constructs a new informer for TitanCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTitanClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FooV1().TitanClusters(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FooV1().TitanClusters(namespace).Watch(options)
			},
		},
		&titanclusterv1.TitanCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *titanClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTitanClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *titanClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&titanclusterv1.TitanCluster{}, f.defaultInformer)
}

func (f *titanClusterInformer) Lister() v1.TitanClusterLister {
	return v1.NewTitanClusterLister(f.Informer().GetIndexer())
}