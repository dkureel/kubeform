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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LoggingBillingAccountSinkLister helps list LoggingBillingAccountSinks.
type LoggingBillingAccountSinkLister interface {
	// List lists all LoggingBillingAccountSinks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountSink, err error)
	// LoggingBillingAccountSinks returns an object that can list and get LoggingBillingAccountSinks.
	LoggingBillingAccountSinks(namespace string) LoggingBillingAccountSinkNamespaceLister
	LoggingBillingAccountSinkListerExpansion
}

// loggingBillingAccountSinkLister implements the LoggingBillingAccountSinkLister interface.
type loggingBillingAccountSinkLister struct {
	indexer cache.Indexer
}

// NewLoggingBillingAccountSinkLister returns a new LoggingBillingAccountSinkLister.
func NewLoggingBillingAccountSinkLister(indexer cache.Indexer) LoggingBillingAccountSinkLister {
	return &loggingBillingAccountSinkLister{indexer: indexer}
}

// List lists all LoggingBillingAccountSinks in the indexer.
func (s *loggingBillingAccountSinkLister) List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountSink, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoggingBillingAccountSink))
	})
	return ret, err
}

// LoggingBillingAccountSinks returns an object that can list and get LoggingBillingAccountSinks.
func (s *loggingBillingAccountSinkLister) LoggingBillingAccountSinks(namespace string) LoggingBillingAccountSinkNamespaceLister {
	return loggingBillingAccountSinkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LoggingBillingAccountSinkNamespaceLister helps list and get LoggingBillingAccountSinks.
type LoggingBillingAccountSinkNamespaceLister interface {
	// List lists all LoggingBillingAccountSinks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountSink, err error)
	// Get retrieves the LoggingBillingAccountSink from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LoggingBillingAccountSink, error)
	LoggingBillingAccountSinkNamespaceListerExpansion
}

// loggingBillingAccountSinkNamespaceLister implements the LoggingBillingAccountSinkNamespaceLister
// interface.
type loggingBillingAccountSinkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LoggingBillingAccountSinks in the indexer for a given namespace.
func (s loggingBillingAccountSinkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountSink, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoggingBillingAccountSink))
	})
	return ret, err
}

// Get retrieves the LoggingBillingAccountSink from the indexer for a given namespace and name.
func (s loggingBillingAccountSinkNamespaceLister) Get(name string) (*v1alpha1.LoggingBillingAccountSink, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("loggingbillingaccountsink"), name)
	}
	return obj.(*v1alpha1.LoggingBillingAccountSink), nil
}
