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

// EbsDefaultKmsKeyLister helps list EbsDefaultKmsKeys.
type EbsDefaultKmsKeyLister interface {
	// List lists all EbsDefaultKmsKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EbsDefaultKmsKey, err error)
	// EbsDefaultKmsKeys returns an object that can list and get EbsDefaultKmsKeys.
	EbsDefaultKmsKeys(namespace string) EbsDefaultKmsKeyNamespaceLister
	EbsDefaultKmsKeyListerExpansion
}

// ebsDefaultKmsKeyLister implements the EbsDefaultKmsKeyLister interface.
type ebsDefaultKmsKeyLister struct {
	indexer cache.Indexer
}

// NewEbsDefaultKmsKeyLister returns a new EbsDefaultKmsKeyLister.
func NewEbsDefaultKmsKeyLister(indexer cache.Indexer) EbsDefaultKmsKeyLister {
	return &ebsDefaultKmsKeyLister{indexer: indexer}
}

// List lists all EbsDefaultKmsKeys in the indexer.
func (s *ebsDefaultKmsKeyLister) List(selector labels.Selector) (ret []*v1alpha1.EbsDefaultKmsKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EbsDefaultKmsKey))
	})
	return ret, err
}

// EbsDefaultKmsKeys returns an object that can list and get EbsDefaultKmsKeys.
func (s *ebsDefaultKmsKeyLister) EbsDefaultKmsKeys(namespace string) EbsDefaultKmsKeyNamespaceLister {
	return ebsDefaultKmsKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EbsDefaultKmsKeyNamespaceLister helps list and get EbsDefaultKmsKeys.
type EbsDefaultKmsKeyNamespaceLister interface {
	// List lists all EbsDefaultKmsKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EbsDefaultKmsKey, err error)
	// Get retrieves the EbsDefaultKmsKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EbsDefaultKmsKey, error)
	EbsDefaultKmsKeyNamespaceListerExpansion
}

// ebsDefaultKmsKeyNamespaceLister implements the EbsDefaultKmsKeyNamespaceLister
// interface.
type ebsDefaultKmsKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EbsDefaultKmsKeys in the indexer for a given namespace.
func (s ebsDefaultKmsKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EbsDefaultKmsKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EbsDefaultKmsKey))
	})
	return ret, err
}

// Get retrieves the EbsDefaultKmsKey from the indexer for a given namespace and name.
func (s ebsDefaultKmsKeyNamespaceLister) Get(name string) (*v1alpha1.EbsDefaultKmsKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ebsdefaultkmskey"), name)
	}
	return obj.(*v1alpha1.EbsDefaultKmsKey), nil
}
