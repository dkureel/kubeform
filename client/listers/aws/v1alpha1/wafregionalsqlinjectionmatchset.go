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

// WafregionalSQLInjectionMatchSetLister helps list WafregionalSQLInjectionMatchSets.
type WafregionalSQLInjectionMatchSetLister interface {
	// List lists all WafregionalSQLInjectionMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalSQLInjectionMatchSet, err error)
	// WafregionalSQLInjectionMatchSets returns an object that can list and get WafregionalSQLInjectionMatchSets.
	WafregionalSQLInjectionMatchSets(namespace string) WafregionalSQLInjectionMatchSetNamespaceLister
	WafregionalSQLInjectionMatchSetListerExpansion
}

// wafregionalSQLInjectionMatchSetLister implements the WafregionalSQLInjectionMatchSetLister interface.
type wafregionalSQLInjectionMatchSetLister struct {
	indexer cache.Indexer
}

// NewWafregionalSQLInjectionMatchSetLister returns a new WafregionalSQLInjectionMatchSetLister.
func NewWafregionalSQLInjectionMatchSetLister(indexer cache.Indexer) WafregionalSQLInjectionMatchSetLister {
	return &wafregionalSQLInjectionMatchSetLister{indexer: indexer}
}

// List lists all WafregionalSQLInjectionMatchSets in the indexer.
func (s *wafregionalSQLInjectionMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalSQLInjectionMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalSQLInjectionMatchSet))
	})
	return ret, err
}

// WafregionalSQLInjectionMatchSets returns an object that can list and get WafregionalSQLInjectionMatchSets.
func (s *wafregionalSQLInjectionMatchSetLister) WafregionalSQLInjectionMatchSets(namespace string) WafregionalSQLInjectionMatchSetNamespaceLister {
	return wafregionalSQLInjectionMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafregionalSQLInjectionMatchSetNamespaceLister helps list and get WafregionalSQLInjectionMatchSets.
type WafregionalSQLInjectionMatchSetNamespaceLister interface {
	// List lists all WafregionalSQLInjectionMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalSQLInjectionMatchSet, err error)
	// Get retrieves the WafregionalSQLInjectionMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafregionalSQLInjectionMatchSet, error)
	WafregionalSQLInjectionMatchSetNamespaceListerExpansion
}

// wafregionalSQLInjectionMatchSetNamespaceLister implements the WafregionalSQLInjectionMatchSetNamespaceLister
// interface.
type wafregionalSQLInjectionMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafregionalSQLInjectionMatchSets in the indexer for a given namespace.
func (s wafregionalSQLInjectionMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalSQLInjectionMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalSQLInjectionMatchSet))
	})
	return ret, err
}

// Get retrieves the WafregionalSQLInjectionMatchSet from the indexer for a given namespace and name.
func (s wafregionalSQLInjectionMatchSetNamespaceLister) Get(name string) (*v1alpha1.WafregionalSQLInjectionMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafregionalsqlinjectionmatchset"), name)
	}
	return obj.(*v1alpha1.WafregionalSQLInjectionMatchSet), nil
}
