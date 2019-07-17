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

// CloudformationStackSetInstanceLister helps list CloudformationStackSetInstances.
type CloudformationStackSetInstanceLister interface {
	// List lists all CloudformationStackSetInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSetInstance, err error)
	// CloudformationStackSetInstances returns an object that can list and get CloudformationStackSetInstances.
	CloudformationStackSetInstances(namespace string) CloudformationStackSetInstanceNamespaceLister
	CloudformationStackSetInstanceListerExpansion
}

// cloudformationStackSetInstanceLister implements the CloudformationStackSetInstanceLister interface.
type cloudformationStackSetInstanceLister struct {
	indexer cache.Indexer
}

// NewCloudformationStackSetInstanceLister returns a new CloudformationStackSetInstanceLister.
func NewCloudformationStackSetInstanceLister(indexer cache.Indexer) CloudformationStackSetInstanceLister {
	return &cloudformationStackSetInstanceLister{indexer: indexer}
}

// List lists all CloudformationStackSetInstances in the indexer.
func (s *cloudformationStackSetInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSetInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudformationStackSetInstance))
	})
	return ret, err
}

// CloudformationStackSetInstances returns an object that can list and get CloudformationStackSetInstances.
func (s *cloudformationStackSetInstanceLister) CloudformationStackSetInstances(namespace string) CloudformationStackSetInstanceNamespaceLister {
	return cloudformationStackSetInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudformationStackSetInstanceNamespaceLister helps list and get CloudformationStackSetInstances.
type CloudformationStackSetInstanceNamespaceLister interface {
	// List lists all CloudformationStackSetInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSetInstance, err error)
	// Get retrieves the CloudformationStackSetInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudformationStackSetInstance, error)
	CloudformationStackSetInstanceNamespaceListerExpansion
}

// cloudformationStackSetInstanceNamespaceLister implements the CloudformationStackSetInstanceNamespaceLister
// interface.
type cloudformationStackSetInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudformationStackSetInstances in the indexer for a given namespace.
func (s cloudformationStackSetInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSetInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudformationStackSetInstance))
	})
	return ret, err
}

// Get retrieves the CloudformationStackSetInstance from the indexer for a given namespace and name.
func (s cloudformationStackSetInstanceNamespaceLister) Get(name string) (*v1alpha1.CloudformationStackSetInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudformationstacksetinstance"), name)
	}
	return obj.(*v1alpha1.CloudformationStackSetInstance), nil
}
