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

// WafByteMatchSetLister helps list WafByteMatchSets.
type WafByteMatchSetLister interface {
	// List lists all WafByteMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafByteMatchSet, err error)
	// Get retrieves the WafByteMatchSet from the index for a given name.
	Get(name string) (*v1alpha1.WafByteMatchSet, error)
	WafByteMatchSetListerExpansion
}

// wafByteMatchSetLister implements the WafByteMatchSetLister interface.
type wafByteMatchSetLister struct {
	indexer cache.Indexer
}

// NewWafByteMatchSetLister returns a new WafByteMatchSetLister.
func NewWafByteMatchSetLister(indexer cache.Indexer) WafByteMatchSetLister {
	return &wafByteMatchSetLister{indexer: indexer}
}

// List lists all WafByteMatchSets in the indexer.
func (s *wafByteMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.WafByteMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafByteMatchSet))
	})
	return ret, err
}

// Get retrieves the WafByteMatchSet from the index for a given name.
func (s *wafByteMatchSetLister) Get(name string) (*v1alpha1.WafByteMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafbytematchset"), name)
	}
	return obj.(*v1alpha1.WafByteMatchSet), nil
}