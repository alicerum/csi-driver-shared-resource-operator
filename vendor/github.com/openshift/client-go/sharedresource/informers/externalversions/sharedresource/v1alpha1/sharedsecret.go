// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	sharedresourcev1alpha1 "github.com/openshift/api/sharedresource/v1alpha1"
	versioned "github.com/openshift/client-go/sharedresource/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/sharedresource/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/client-go/sharedresource/listers/sharedresource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SharedSecretInformer provides access to a shared informer and lister for
// SharedSecrets.
type SharedSecretInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SharedSecretLister
}

type sharedSecretInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSharedSecretInformer constructs a new informer for SharedSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSharedSecretInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSharedSecretInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSharedSecretInformer constructs a new informer for SharedSecret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSharedSecretInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SharedresourceV1alpha1().SharedSecrets().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SharedresourceV1alpha1().SharedSecrets().Watch(context.TODO(), options)
			},
		},
		&sharedresourcev1alpha1.SharedSecret{},
		resyncPeriod,
		indexers,
	)
}

func (f *sharedSecretInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSharedSecretInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sharedSecretInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sharedresourcev1alpha1.SharedSecret{}, f.defaultInformer)
}

func (f *sharedSecretInformer) Lister() v1alpha1.SharedSecretLister {
	return v1alpha1.NewSharedSecretLister(f.Informer().GetIndexer())
}
