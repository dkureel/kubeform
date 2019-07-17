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

// DatasyncAgentLister helps list DatasyncAgents.
type DatasyncAgentLister interface {
	// List lists all DatasyncAgents in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DatasyncAgent, err error)
	// DatasyncAgents returns an object that can list and get DatasyncAgents.
	DatasyncAgents(namespace string) DatasyncAgentNamespaceLister
	DatasyncAgentListerExpansion
}

// datasyncAgentLister implements the DatasyncAgentLister interface.
type datasyncAgentLister struct {
	indexer cache.Indexer
}

// NewDatasyncAgentLister returns a new DatasyncAgentLister.
func NewDatasyncAgentLister(indexer cache.Indexer) DatasyncAgentLister {
	return &datasyncAgentLister{indexer: indexer}
}

// List lists all DatasyncAgents in the indexer.
func (s *datasyncAgentLister) List(selector labels.Selector) (ret []*v1alpha1.DatasyncAgent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasyncAgent))
	})
	return ret, err
}

// DatasyncAgents returns an object that can list and get DatasyncAgents.
func (s *datasyncAgentLister) DatasyncAgents(namespace string) DatasyncAgentNamespaceLister {
	return datasyncAgentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatasyncAgentNamespaceLister helps list and get DatasyncAgents.
type DatasyncAgentNamespaceLister interface {
	// List lists all DatasyncAgents in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DatasyncAgent, err error)
	// Get retrieves the DatasyncAgent from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DatasyncAgent, error)
	DatasyncAgentNamespaceListerExpansion
}

// datasyncAgentNamespaceLister implements the DatasyncAgentNamespaceLister
// interface.
type datasyncAgentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DatasyncAgents in the indexer for a given namespace.
func (s datasyncAgentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DatasyncAgent, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasyncAgent))
	})
	return ret, err
}

// Get retrieves the DatasyncAgent from the indexer for a given namespace and name.
func (s datasyncAgentNamespaceLister) Get(name string) (*v1alpha1.DatasyncAgent, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datasyncagent"), name)
	}
	return obj.(*v1alpha1.DatasyncAgent), nil
}
