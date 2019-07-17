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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// ComputeVPNGatewayLister helps list ComputeVPNGateways.
type ComputeVPNGatewayLister interface {
	// List lists all ComputeVPNGateways in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNGateway, err error)
	// ComputeVPNGateways returns an object that can list and get ComputeVPNGateways.
	ComputeVPNGateways(namespace string) ComputeVPNGatewayNamespaceLister
	ComputeVPNGatewayListerExpansion
}

// computeVPNGatewayLister implements the ComputeVPNGatewayLister interface.
type computeVPNGatewayLister struct {
	indexer cache.Indexer
}

// NewComputeVPNGatewayLister returns a new ComputeVPNGatewayLister.
func NewComputeVPNGatewayLister(indexer cache.Indexer) ComputeVPNGatewayLister {
	return &computeVPNGatewayLister{indexer: indexer}
}

// List lists all ComputeVPNGateways in the indexer.
func (s *computeVPNGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeVPNGateway))
	})
	return ret, err
}

// ComputeVPNGateways returns an object that can list and get ComputeVPNGateways.
func (s *computeVPNGatewayLister) ComputeVPNGateways(namespace string) ComputeVPNGatewayNamespaceLister {
	return computeVPNGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeVPNGatewayNamespaceLister helps list and get ComputeVPNGateways.
type ComputeVPNGatewayNamespaceLister interface {
	// List lists all ComputeVPNGateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNGateway, err error)
	// Get retrieves the ComputeVPNGateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeVPNGateway, error)
	ComputeVPNGatewayNamespaceListerExpansion
}

// computeVPNGatewayNamespaceLister implements the ComputeVPNGatewayNamespaceLister
// interface.
type computeVPNGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeVPNGateways in the indexer for a given namespace.
func (s computeVPNGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeVPNGateway))
	})
	return ret, err
}

// Get retrieves the ComputeVPNGateway from the indexer for a given namespace and name.
func (s computeVPNGatewayNamespaceLister) Get(name string) (*v1alpha1.ComputeVPNGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computevpngateway"), name)
	}
	return obj.(*v1alpha1.ComputeVPNGateway), nil
}
