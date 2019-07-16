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

// SsmMaintenanceWindowTargetLister helps list SsmMaintenanceWindowTargets.
type SsmMaintenanceWindowTargetLister interface {
	// List lists all SsmMaintenanceWindowTargets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SsmMaintenanceWindowTarget, err error)
	// Get retrieves the SsmMaintenanceWindowTarget from the index for a given name.
	Get(name string) (*v1alpha1.SsmMaintenanceWindowTarget, error)
	SsmMaintenanceWindowTargetListerExpansion
}

// ssmMaintenanceWindowTargetLister implements the SsmMaintenanceWindowTargetLister interface.
type ssmMaintenanceWindowTargetLister struct {
	indexer cache.Indexer
}

// NewSsmMaintenanceWindowTargetLister returns a new SsmMaintenanceWindowTargetLister.
func NewSsmMaintenanceWindowTargetLister(indexer cache.Indexer) SsmMaintenanceWindowTargetLister {
	return &ssmMaintenanceWindowTargetLister{indexer: indexer}
}

// List lists all SsmMaintenanceWindowTargets in the indexer.
func (s *ssmMaintenanceWindowTargetLister) List(selector labels.Selector) (ret []*v1alpha1.SsmMaintenanceWindowTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SsmMaintenanceWindowTarget))
	})
	return ret, err
}

// Get retrieves the SsmMaintenanceWindowTarget from the index for a given name.
func (s *ssmMaintenanceWindowTargetLister) Get(name string) (*v1alpha1.SsmMaintenanceWindowTarget, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ssmmaintenancewindowtarget"), name)
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTarget), nil
}