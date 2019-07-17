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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// DbInstancesGetter has a method to return a DbInstanceInterface.
// A group's client should implement this interface.
type DbInstancesGetter interface {
	DbInstances(namespace string) DbInstanceInterface
}

// DbInstanceInterface has methods to work with DbInstance resources.
type DbInstanceInterface interface {
	Create(*v1alpha1.DbInstance) (*v1alpha1.DbInstance, error)
	Update(*v1alpha1.DbInstance) (*v1alpha1.DbInstance, error)
	UpdateStatus(*v1alpha1.DbInstance) (*v1alpha1.DbInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DbInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.DbInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbInstance, err error)
	DbInstanceExpansion
}

// dbInstances implements DbInstanceInterface
type dbInstances struct {
	client rest.Interface
	ns     string
}

// newDbInstances returns a DbInstances
func newDbInstances(c *AwsV1alpha1Client, namespace string) *dbInstances {
	return &dbInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dbInstance, and returns the corresponding dbInstance object, and an error if there is any.
func (c *dbInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.DbInstance, err error) {
	result = &v1alpha1.DbInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DbInstances that match those selectors.
func (c *dbInstances) List(opts v1.ListOptions) (result *v1alpha1.DbInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DbInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dbInstances.
func (c *dbInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dbinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dbInstance and creates it.  Returns the server's representation of the dbInstance, and an error, if there is any.
func (c *dbInstances) Create(dbInstance *v1alpha1.DbInstance) (result *v1alpha1.DbInstance, err error) {
	result = &v1alpha1.DbInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dbinstances").
		Body(dbInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dbInstance and updates it. Returns the server's representation of the dbInstance, and an error, if there is any.
func (c *dbInstances) Update(dbInstance *v1alpha1.DbInstance) (result *v1alpha1.DbInstance, err error) {
	result = &v1alpha1.DbInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbinstances").
		Name(dbInstance.Name).
		Body(dbInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dbInstances) UpdateStatus(dbInstance *v1alpha1.DbInstance) (result *v1alpha1.DbInstance, err error) {
	result = &v1alpha1.DbInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbinstances").
		Name(dbInstance.Name).
		SubResource("status").
		Body(dbInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the dbInstance and deletes it. Returns an error if one occurs.
func (c *dbInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dbInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dbInstance.
func (c *dbInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbInstance, err error) {
	result = &v1alpha1.DbInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dbinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
