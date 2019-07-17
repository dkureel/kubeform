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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// LbBackendAddressPoolLister helps list LbBackendAddressPools.
type LbBackendAddressPoolLister interface {
	// List lists all LbBackendAddressPools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LbBackendAddressPool, err error)
	// LbBackendAddressPools returns an object that can list and get LbBackendAddressPools.
	LbBackendAddressPools(namespace string) LbBackendAddressPoolNamespaceLister
	LbBackendAddressPoolListerExpansion
}

// lbBackendAddressPoolLister implements the LbBackendAddressPoolLister interface.
type lbBackendAddressPoolLister struct {
	indexer cache.Indexer
}

// NewLbBackendAddressPoolLister returns a new LbBackendAddressPoolLister.
func NewLbBackendAddressPoolLister(indexer cache.Indexer) LbBackendAddressPoolLister {
	return &lbBackendAddressPoolLister{indexer: indexer}
}

// List lists all LbBackendAddressPools in the indexer.
func (s *lbBackendAddressPoolLister) List(selector labels.Selector) (ret []*v1alpha1.LbBackendAddressPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbBackendAddressPool))
	})
	return ret, err
}

// LbBackendAddressPools returns an object that can list and get LbBackendAddressPools.
func (s *lbBackendAddressPoolLister) LbBackendAddressPools(namespace string) LbBackendAddressPoolNamespaceLister {
	return lbBackendAddressPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LbBackendAddressPoolNamespaceLister helps list and get LbBackendAddressPools.
type LbBackendAddressPoolNamespaceLister interface {
	// List lists all LbBackendAddressPools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LbBackendAddressPool, err error)
	// Get retrieves the LbBackendAddressPool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LbBackendAddressPool, error)
	LbBackendAddressPoolNamespaceListerExpansion
}

// lbBackendAddressPoolNamespaceLister implements the LbBackendAddressPoolNamespaceLister
// interface.
type lbBackendAddressPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LbBackendAddressPools in the indexer for a given namespace.
func (s lbBackendAddressPoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LbBackendAddressPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbBackendAddressPool))
	})
	return ret, err
}

// Get retrieves the LbBackendAddressPool from the indexer for a given namespace and name.
func (s lbBackendAddressPoolNamespaceLister) Get(name string) (*v1alpha1.LbBackendAddressPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lbbackendaddresspool"), name)
	}
	return obj.(*v1alpha1.LbBackendAddressPool), nil
}
