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

// IamServiceLinkedRoleLister helps list IamServiceLinkedRoles.
type IamServiceLinkedRoleLister interface {
	// List lists all IamServiceLinkedRoles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamServiceLinkedRole, err error)
	// IamServiceLinkedRoles returns an object that can list and get IamServiceLinkedRoles.
	IamServiceLinkedRoles(namespace string) IamServiceLinkedRoleNamespaceLister
	IamServiceLinkedRoleListerExpansion
}

// iamServiceLinkedRoleLister implements the IamServiceLinkedRoleLister interface.
type iamServiceLinkedRoleLister struct {
	indexer cache.Indexer
}

// NewIamServiceLinkedRoleLister returns a new IamServiceLinkedRoleLister.
func NewIamServiceLinkedRoleLister(indexer cache.Indexer) IamServiceLinkedRoleLister {
	return &iamServiceLinkedRoleLister{indexer: indexer}
}

// List lists all IamServiceLinkedRoles in the indexer.
func (s *iamServiceLinkedRoleLister) List(selector labels.Selector) (ret []*v1alpha1.IamServiceLinkedRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamServiceLinkedRole))
	})
	return ret, err
}

// IamServiceLinkedRoles returns an object that can list and get IamServiceLinkedRoles.
func (s *iamServiceLinkedRoleLister) IamServiceLinkedRoles(namespace string) IamServiceLinkedRoleNamespaceLister {
	return iamServiceLinkedRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamServiceLinkedRoleNamespaceLister helps list and get IamServiceLinkedRoles.
type IamServiceLinkedRoleNamespaceLister interface {
	// List lists all IamServiceLinkedRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamServiceLinkedRole, err error)
	// Get retrieves the IamServiceLinkedRole from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamServiceLinkedRole, error)
	IamServiceLinkedRoleNamespaceListerExpansion
}

// iamServiceLinkedRoleNamespaceLister implements the IamServiceLinkedRoleNamespaceLister
// interface.
type iamServiceLinkedRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamServiceLinkedRoles in the indexer for a given namespace.
func (s iamServiceLinkedRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamServiceLinkedRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamServiceLinkedRole))
	})
	return ret, err
}

// Get retrieves the IamServiceLinkedRole from the indexer for a given namespace and name.
func (s iamServiceLinkedRoleNamespaceLister) Get(name string) (*v1alpha1.IamServiceLinkedRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamservicelinkedrole"), name)
	}
	return obj.(*v1alpha1.IamServiceLinkedRole), nil
}
