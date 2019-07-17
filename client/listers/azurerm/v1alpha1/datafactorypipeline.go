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

// DataFactoryPipelineLister helps list DataFactoryPipelines.
type DataFactoryPipelineLister interface {
	// List lists all DataFactoryPipelines in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryPipeline, err error)
	// DataFactoryPipelines returns an object that can list and get DataFactoryPipelines.
	DataFactoryPipelines(namespace string) DataFactoryPipelineNamespaceLister
	DataFactoryPipelineListerExpansion
}

// dataFactoryPipelineLister implements the DataFactoryPipelineLister interface.
type dataFactoryPipelineLister struct {
	indexer cache.Indexer
}

// NewDataFactoryPipelineLister returns a new DataFactoryPipelineLister.
func NewDataFactoryPipelineLister(indexer cache.Indexer) DataFactoryPipelineLister {
	return &dataFactoryPipelineLister{indexer: indexer}
}

// List lists all DataFactoryPipelines in the indexer.
func (s *dataFactoryPipelineLister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryPipeline, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryPipeline))
	})
	return ret, err
}

// DataFactoryPipelines returns an object that can list and get DataFactoryPipelines.
func (s *dataFactoryPipelineLister) DataFactoryPipelines(namespace string) DataFactoryPipelineNamespaceLister {
	return dataFactoryPipelineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataFactoryPipelineNamespaceLister helps list and get DataFactoryPipelines.
type DataFactoryPipelineNamespaceLister interface {
	// List lists all DataFactoryPipelines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryPipeline, err error)
	// Get retrieves the DataFactoryPipeline from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataFactoryPipeline, error)
	DataFactoryPipelineNamespaceListerExpansion
}

// dataFactoryPipelineNamespaceLister implements the DataFactoryPipelineNamespaceLister
// interface.
type dataFactoryPipelineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataFactoryPipelines in the indexer for a given namespace.
func (s dataFactoryPipelineNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryPipeline, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryPipeline))
	})
	return ret, err
}

// Get retrieves the DataFactoryPipeline from the indexer for a given namespace and name.
func (s dataFactoryPipelineNamespaceLister) Get(name string) (*v1alpha1.DataFactoryPipeline, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datafactorypipeline"), name)
	}
	return obj.(*v1alpha1.DataFactoryPipeline), nil
}
