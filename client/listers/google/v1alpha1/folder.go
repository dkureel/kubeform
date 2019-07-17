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

// FolderLister helps list Folders.
type FolderLister interface {
	// List lists all Folders in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Folder, err error)
	// Folders returns an object that can list and get Folders.
	Folders(namespace string) FolderNamespaceLister
	FolderListerExpansion
}

// folderLister implements the FolderLister interface.
type folderLister struct {
	indexer cache.Indexer
}

// NewFolderLister returns a new FolderLister.
func NewFolderLister(indexer cache.Indexer) FolderLister {
	return &folderLister{indexer: indexer}
}

// List lists all Folders in the indexer.
func (s *folderLister) List(selector labels.Selector) (ret []*v1alpha1.Folder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Folder))
	})
	return ret, err
}

// Folders returns an object that can list and get Folders.
func (s *folderLister) Folders(namespace string) FolderNamespaceLister {
	return folderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FolderNamespaceLister helps list and get Folders.
type FolderNamespaceLister interface {
	// List lists all Folders in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Folder, err error)
	// Get retrieves the Folder from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Folder, error)
	FolderNamespaceListerExpansion
}

// folderNamespaceLister implements the FolderNamespaceLister
// interface.
type folderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Folders in the indexer for a given namespace.
func (s folderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Folder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Folder))
	})
	return ret, err
}

// Get retrieves the Folder from the indexer for a given namespace and name.
func (s folderNamespaceLister) Get(name string) (*v1alpha1.Folder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("folder"), name)
	}
	return obj.(*v1alpha1.Folder), nil
}
