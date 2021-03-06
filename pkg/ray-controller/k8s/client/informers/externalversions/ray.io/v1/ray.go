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

	rayiov1 "github.com/ray-operator/pkg/ray-controller/k8s/apis/ray.io/v1"
	versioned "github.com/ray-operator/pkg/ray-controller/k8s/client/clientset/versioned"
	internalinterfaces "github.com/ray-operator/pkg/ray-controller/k8s/client/informers/externalversions/internalinterfaces"
	v1 "github.com/ray-operator/pkg/ray-controller/k8s/client/listers/ray.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RayInformer provides access to a shared informer and lister for
// Rays.
type RayInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RayLister
}

type rayInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRayInformer constructs a new informer for Ray type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRayInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRayInformer constructs a new informer for Ray type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RayV1().Rays(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RayV1().Rays(namespace).Watch(options)
			},
		},
		&rayiov1.Ray{},
		resyncPeriod,
		indexers,
	)
}

func (f *rayInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRayInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rayInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rayiov1.Ray{}, f.defaultInformer)
}

func (f *rayInformer) Lister() v1.RayLister {
	return v1.NewRayLister(f.Informer().GetIndexer())
}
