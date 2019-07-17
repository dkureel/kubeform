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

// RedshiftSubnetGroupLister helps list RedshiftSubnetGroups.
type RedshiftSubnetGroupLister interface {
	// List lists all RedshiftSubnetGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftSubnetGroup, err error)
	// RedshiftSubnetGroups returns an object that can list and get RedshiftSubnetGroups.
	RedshiftSubnetGroups(namespace string) RedshiftSubnetGroupNamespaceLister
	RedshiftSubnetGroupListerExpansion
}

// redshiftSubnetGroupLister implements the RedshiftSubnetGroupLister interface.
type redshiftSubnetGroupLister struct {
	indexer cache.Indexer
}

// NewRedshiftSubnetGroupLister returns a new RedshiftSubnetGroupLister.
func NewRedshiftSubnetGroupLister(indexer cache.Indexer) RedshiftSubnetGroupLister {
	return &redshiftSubnetGroupLister{indexer: indexer}
}

// List lists all RedshiftSubnetGroups in the indexer.
func (s *redshiftSubnetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftSubnetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftSubnetGroup))
	})
	return ret, err
}

// RedshiftSubnetGroups returns an object that can list and get RedshiftSubnetGroups.
func (s *redshiftSubnetGroupLister) RedshiftSubnetGroups(namespace string) RedshiftSubnetGroupNamespaceLister {
	return redshiftSubnetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedshiftSubnetGroupNamespaceLister helps list and get RedshiftSubnetGroups.
type RedshiftSubnetGroupNamespaceLister interface {
	// List lists all RedshiftSubnetGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftSubnetGroup, err error)
	// Get retrieves the RedshiftSubnetGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RedshiftSubnetGroup, error)
	RedshiftSubnetGroupNamespaceListerExpansion
}

// redshiftSubnetGroupNamespaceLister implements the RedshiftSubnetGroupNamespaceLister
// interface.
type redshiftSubnetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedshiftSubnetGroups in the indexer for a given namespace.
func (s redshiftSubnetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftSubnetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftSubnetGroup))
	})
	return ret, err
}

// Get retrieves the RedshiftSubnetGroup from the indexer for a given namespace and name.
func (s redshiftSubnetGroupNamespaceLister) Get(name string) (*v1alpha1.RedshiftSubnetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("redshiftsubnetgroup"), name)
	}
	return obj.(*v1alpha1.RedshiftSubnetGroup), nil
}
