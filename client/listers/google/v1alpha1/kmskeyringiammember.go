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

// KmsKeyRingIamMemberLister helps list KmsKeyRingIamMembers.
type KmsKeyRingIamMemberLister interface {
	// List lists all KmsKeyRingIamMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KmsKeyRingIamMember, err error)
	// KmsKeyRingIamMembers returns an object that can list and get KmsKeyRingIamMembers.
	KmsKeyRingIamMembers(namespace string) KmsKeyRingIamMemberNamespaceLister
	KmsKeyRingIamMemberListerExpansion
}

// kmsKeyRingIamMemberLister implements the KmsKeyRingIamMemberLister interface.
type kmsKeyRingIamMemberLister struct {
	indexer cache.Indexer
}

// NewKmsKeyRingIamMemberLister returns a new KmsKeyRingIamMemberLister.
func NewKmsKeyRingIamMemberLister(indexer cache.Indexer) KmsKeyRingIamMemberLister {
	return &kmsKeyRingIamMemberLister{indexer: indexer}
}

// List lists all KmsKeyRingIamMembers in the indexer.
func (s *kmsKeyRingIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.KmsKeyRingIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsKeyRingIamMember))
	})
	return ret, err
}

// KmsKeyRingIamMembers returns an object that can list and get KmsKeyRingIamMembers.
func (s *kmsKeyRingIamMemberLister) KmsKeyRingIamMembers(namespace string) KmsKeyRingIamMemberNamespaceLister {
	return kmsKeyRingIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KmsKeyRingIamMemberNamespaceLister helps list and get KmsKeyRingIamMembers.
type KmsKeyRingIamMemberNamespaceLister interface {
	// List lists all KmsKeyRingIamMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KmsKeyRingIamMember, err error)
	// Get retrieves the KmsKeyRingIamMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KmsKeyRingIamMember, error)
	KmsKeyRingIamMemberNamespaceListerExpansion
}

// kmsKeyRingIamMemberNamespaceLister implements the KmsKeyRingIamMemberNamespaceLister
// interface.
type kmsKeyRingIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KmsKeyRingIamMembers in the indexer for a given namespace.
func (s kmsKeyRingIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KmsKeyRingIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsKeyRingIamMember))
	})
	return ret, err
}

// Get retrieves the KmsKeyRingIamMember from the indexer for a given namespace and name.
func (s kmsKeyRingIamMemberNamespaceLister) Get(name string) (*v1alpha1.KmsKeyRingIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kmskeyringiammember"), name)
	}
	return obj.(*v1alpha1.KmsKeyRingIamMember), nil
}
