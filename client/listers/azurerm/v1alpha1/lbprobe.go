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

// LbProbeLister helps list LbProbes.
type LbProbeLister interface {
	// List lists all LbProbes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LbProbe, err error)
	// LbProbes returns an object that can list and get LbProbes.
	LbProbes(namespace string) LbProbeNamespaceLister
	LbProbeListerExpansion
}

// lbProbeLister implements the LbProbeLister interface.
type lbProbeLister struct {
	indexer cache.Indexer
}

// NewLbProbeLister returns a new LbProbeLister.
func NewLbProbeLister(indexer cache.Indexer) LbProbeLister {
	return &lbProbeLister{indexer: indexer}
}

// List lists all LbProbes in the indexer.
func (s *lbProbeLister) List(selector labels.Selector) (ret []*v1alpha1.LbProbe, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbProbe))
	})
	return ret, err
}

// LbProbes returns an object that can list and get LbProbes.
func (s *lbProbeLister) LbProbes(namespace string) LbProbeNamespaceLister {
	return lbProbeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LbProbeNamespaceLister helps list and get LbProbes.
type LbProbeNamespaceLister interface {
	// List lists all LbProbes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LbProbe, err error)
	// Get retrieves the LbProbe from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LbProbe, error)
	LbProbeNamespaceListerExpansion
}

// lbProbeNamespaceLister implements the LbProbeNamespaceLister
// interface.
type lbProbeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LbProbes in the indexer for a given namespace.
func (s lbProbeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LbProbe, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbProbe))
	})
	return ret, err
}

// Get retrieves the LbProbe from the indexer for a given namespace and name.
func (s lbProbeNamespaceLister) Get(name string) (*v1alpha1.LbProbe, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lbprobe"), name)
	}
	return obj.(*v1alpha1.LbProbe), nil
}
