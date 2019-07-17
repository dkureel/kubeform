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

// GuarddutyInviteAccepterLister helps list GuarddutyInviteAccepters.
type GuarddutyInviteAccepterLister interface {
	// List lists all GuarddutyInviteAccepters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GuarddutyInviteAccepter, err error)
	// GuarddutyInviteAccepters returns an object that can list and get GuarddutyInviteAccepters.
	GuarddutyInviteAccepters(namespace string) GuarddutyInviteAccepterNamespaceLister
	GuarddutyInviteAccepterListerExpansion
}

// guarddutyInviteAccepterLister implements the GuarddutyInviteAccepterLister interface.
type guarddutyInviteAccepterLister struct {
	indexer cache.Indexer
}

// NewGuarddutyInviteAccepterLister returns a new GuarddutyInviteAccepterLister.
func NewGuarddutyInviteAccepterLister(indexer cache.Indexer) GuarddutyInviteAccepterLister {
	return &guarddutyInviteAccepterLister{indexer: indexer}
}

// List lists all GuarddutyInviteAccepters in the indexer.
func (s *guarddutyInviteAccepterLister) List(selector labels.Selector) (ret []*v1alpha1.GuarddutyInviteAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GuarddutyInviteAccepter))
	})
	return ret, err
}

// GuarddutyInviteAccepters returns an object that can list and get GuarddutyInviteAccepters.
func (s *guarddutyInviteAccepterLister) GuarddutyInviteAccepters(namespace string) GuarddutyInviteAccepterNamespaceLister {
	return guarddutyInviteAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GuarddutyInviteAccepterNamespaceLister helps list and get GuarddutyInviteAccepters.
type GuarddutyInviteAccepterNamespaceLister interface {
	// List lists all GuarddutyInviteAccepters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GuarddutyInviteAccepter, err error)
	// Get retrieves the GuarddutyInviteAccepter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GuarddutyInviteAccepter, error)
	GuarddutyInviteAccepterNamespaceListerExpansion
}

// guarddutyInviteAccepterNamespaceLister implements the GuarddutyInviteAccepterNamespaceLister
// interface.
type guarddutyInviteAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GuarddutyInviteAccepters in the indexer for a given namespace.
func (s guarddutyInviteAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GuarddutyInviteAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GuarddutyInviteAccepter))
	})
	return ret, err
}

// Get retrieves the GuarddutyInviteAccepter from the indexer for a given namespace and name.
func (s guarddutyInviteAccepterNamespaceLister) Get(name string) (*v1alpha1.GuarddutyInviteAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("guarddutyinviteaccepter"), name)
	}
	return obj.(*v1alpha1.GuarddutyInviteAccepter), nil
}
