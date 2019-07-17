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

// DevicefarmProjectLister helps list DevicefarmProjects.
type DevicefarmProjectLister interface {
	// List lists all DevicefarmProjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DevicefarmProject, err error)
	// DevicefarmProjects returns an object that can list and get DevicefarmProjects.
	DevicefarmProjects(namespace string) DevicefarmProjectNamespaceLister
	DevicefarmProjectListerExpansion
}

// devicefarmProjectLister implements the DevicefarmProjectLister interface.
type devicefarmProjectLister struct {
	indexer cache.Indexer
}

// NewDevicefarmProjectLister returns a new DevicefarmProjectLister.
func NewDevicefarmProjectLister(indexer cache.Indexer) DevicefarmProjectLister {
	return &devicefarmProjectLister{indexer: indexer}
}

// List lists all DevicefarmProjects in the indexer.
func (s *devicefarmProjectLister) List(selector labels.Selector) (ret []*v1alpha1.DevicefarmProject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevicefarmProject))
	})
	return ret, err
}

// DevicefarmProjects returns an object that can list and get DevicefarmProjects.
func (s *devicefarmProjectLister) DevicefarmProjects(namespace string) DevicefarmProjectNamespaceLister {
	return devicefarmProjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DevicefarmProjectNamespaceLister helps list and get DevicefarmProjects.
type DevicefarmProjectNamespaceLister interface {
	// List lists all DevicefarmProjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DevicefarmProject, err error)
	// Get retrieves the DevicefarmProject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DevicefarmProject, error)
	DevicefarmProjectNamespaceListerExpansion
}

// devicefarmProjectNamespaceLister implements the DevicefarmProjectNamespaceLister
// interface.
type devicefarmProjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DevicefarmProjects in the indexer for a given namespace.
func (s devicefarmProjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DevicefarmProject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevicefarmProject))
	})
	return ret, err
}

// Get retrieves the DevicefarmProject from the indexer for a given namespace and name.
func (s devicefarmProjectNamespaceLister) Get(name string) (*v1alpha1.DevicefarmProject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("devicefarmproject"), name)
	}
	return obj.(*v1alpha1.DevicefarmProject), nil
}
