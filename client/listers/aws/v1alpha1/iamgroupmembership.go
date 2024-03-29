/*
Copyright The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IamGroupMembershipLister helps list IamGroupMemberships.
type IamGroupMembershipLister interface {
	// List lists all IamGroupMemberships in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamGroupMembership, err error)
	// IamGroupMemberships returns an object that can list and get IamGroupMemberships.
	IamGroupMemberships(namespace string) IamGroupMembershipNamespaceLister
	IamGroupMembershipListerExpansion
}

// iamGroupMembershipLister implements the IamGroupMembershipLister interface.
type iamGroupMembershipLister struct {
	indexer cache.Indexer
}

// NewIamGroupMembershipLister returns a new IamGroupMembershipLister.
func NewIamGroupMembershipLister(indexer cache.Indexer) IamGroupMembershipLister {
	return &iamGroupMembershipLister{indexer: indexer}
}

// List lists all IamGroupMemberships in the indexer.
func (s *iamGroupMembershipLister) List(selector labels.Selector) (ret []*v1alpha1.IamGroupMembership, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamGroupMembership))
	})
	return ret, err
}

// IamGroupMemberships returns an object that can list and get IamGroupMemberships.
func (s *iamGroupMembershipLister) IamGroupMemberships(namespace string) IamGroupMembershipNamespaceLister {
	return iamGroupMembershipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamGroupMembershipNamespaceLister helps list and get IamGroupMemberships.
type IamGroupMembershipNamespaceLister interface {
	// List lists all IamGroupMemberships in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamGroupMembership, err error)
	// Get retrieves the IamGroupMembership from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamGroupMembership, error)
	IamGroupMembershipNamespaceListerExpansion
}

// iamGroupMembershipNamespaceLister implements the IamGroupMembershipNamespaceLister
// interface.
type iamGroupMembershipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamGroupMemberships in the indexer for a given namespace.
func (s iamGroupMembershipNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamGroupMembership, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamGroupMembership))
	})
	return ret, err
}

// Get retrieves the IamGroupMembership from the indexer for a given namespace and name.
func (s iamGroupMembershipNamespaceLister) Get(name string) (*v1alpha1.IamGroupMembership, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamgroupmembership"), name)
	}
	return obj.(*v1alpha1.IamGroupMembership), nil
}
