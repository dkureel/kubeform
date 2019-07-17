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

// KmsExternalKeyLister helps list KmsExternalKeys.
type KmsExternalKeyLister interface {
	// List lists all KmsExternalKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KmsExternalKey, err error)
	// KmsExternalKeys returns an object that can list and get KmsExternalKeys.
	KmsExternalKeys(namespace string) KmsExternalKeyNamespaceLister
	KmsExternalKeyListerExpansion
}

// kmsExternalKeyLister implements the KmsExternalKeyLister interface.
type kmsExternalKeyLister struct {
	indexer cache.Indexer
}

// NewKmsExternalKeyLister returns a new KmsExternalKeyLister.
func NewKmsExternalKeyLister(indexer cache.Indexer) KmsExternalKeyLister {
	return &kmsExternalKeyLister{indexer: indexer}
}

// List lists all KmsExternalKeys in the indexer.
func (s *kmsExternalKeyLister) List(selector labels.Selector) (ret []*v1alpha1.KmsExternalKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsExternalKey))
	})
	return ret, err
}

// KmsExternalKeys returns an object that can list and get KmsExternalKeys.
func (s *kmsExternalKeyLister) KmsExternalKeys(namespace string) KmsExternalKeyNamespaceLister {
	return kmsExternalKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KmsExternalKeyNamespaceLister helps list and get KmsExternalKeys.
type KmsExternalKeyNamespaceLister interface {
	// List lists all KmsExternalKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KmsExternalKey, err error)
	// Get retrieves the KmsExternalKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KmsExternalKey, error)
	KmsExternalKeyNamespaceListerExpansion
}

// kmsExternalKeyNamespaceLister implements the KmsExternalKeyNamespaceLister
// interface.
type kmsExternalKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KmsExternalKeys in the indexer for a given namespace.
func (s kmsExternalKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KmsExternalKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsExternalKey))
	})
	return ret, err
}

// Get retrieves the KmsExternalKey from the indexer for a given namespace and name.
func (s kmsExternalKeyNamespaceLister) Get(name string) (*v1alpha1.KmsExternalKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kmsexternalkey"), name)
	}
	return obj.(*v1alpha1.KmsExternalKey), nil
}
