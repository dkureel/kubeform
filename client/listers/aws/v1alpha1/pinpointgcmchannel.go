/*
Copyright 2019 The Kubeform Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// PinpointGcmChannelLister helps list PinpointGcmChannels.
type PinpointGcmChannelLister interface {
	// List lists all PinpointGcmChannels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointGcmChannel, err error)
	// PinpointGcmChannels returns an object that can list and get PinpointGcmChannels.
	PinpointGcmChannels(namespace string) PinpointGcmChannelNamespaceLister
	PinpointGcmChannelListerExpansion
}

// pinpointGcmChannelLister implements the PinpointGcmChannelLister interface.
type pinpointGcmChannelLister struct {
	indexer cache.Indexer
}

// NewPinpointGcmChannelLister returns a new PinpointGcmChannelLister.
func NewPinpointGcmChannelLister(indexer cache.Indexer) PinpointGcmChannelLister {
	return &pinpointGcmChannelLister{indexer: indexer}
}

// List lists all PinpointGcmChannels in the indexer.
func (s *pinpointGcmChannelLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointGcmChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointGcmChannel))
	})
	return ret, err
}

// PinpointGcmChannels returns an object that can list and get PinpointGcmChannels.
func (s *pinpointGcmChannelLister) PinpointGcmChannels(namespace string) PinpointGcmChannelNamespaceLister {
	return pinpointGcmChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PinpointGcmChannelNamespaceLister helps list and get PinpointGcmChannels.
type PinpointGcmChannelNamespaceLister interface {
	// List lists all PinpointGcmChannels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PinpointGcmChannel, err error)
	// Get retrieves the PinpointGcmChannel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PinpointGcmChannel, error)
	PinpointGcmChannelNamespaceListerExpansion
}

// pinpointGcmChannelNamespaceLister implements the PinpointGcmChannelNamespaceLister
// interface.
type pinpointGcmChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PinpointGcmChannels in the indexer for a given namespace.
func (s pinpointGcmChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PinpointGcmChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PinpointGcmChannel))
	})
	return ret, err
}

// Get retrieves the PinpointGcmChannel from the indexer for a given namespace and name.
func (s pinpointGcmChannelNamespaceLister) Get(name string) (*v1alpha1.PinpointGcmChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pinpointgcmchannel"), name)
	}
	return obj.(*v1alpha1.PinpointGcmChannel), nil
}
