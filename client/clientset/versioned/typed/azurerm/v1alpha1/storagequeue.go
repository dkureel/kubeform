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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// StorageQueuesGetter has a method to return a StorageQueueInterface.
// A group's client should implement this interface.
type StorageQueuesGetter interface {
	StorageQueues(namespace string) StorageQueueInterface
}

// StorageQueueInterface has methods to work with StorageQueue resources.
type StorageQueueInterface interface {
	Create(*v1alpha1.StorageQueue) (*v1alpha1.StorageQueue, error)
	Update(*v1alpha1.StorageQueue) (*v1alpha1.StorageQueue, error)
	UpdateStatus(*v1alpha1.StorageQueue) (*v1alpha1.StorageQueue, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageQueue, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageQueueList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageQueue, err error)
	StorageQueueExpansion
}

// storageQueues implements StorageQueueInterface
type storageQueues struct {
	client rest.Interface
	ns     string
}

// newStorageQueues returns a StorageQueues
func newStorageQueues(c *AzurermV1alpha1Client, namespace string) *storageQueues {
	return &storageQueues{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storageQueue, and returns the corresponding storageQueue object, and an error if there is any.
func (c *storageQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageQueue, err error) {
	result = &v1alpha1.StorageQueue{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagequeues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageQueues that match those selectors.
func (c *storageQueues) List(opts v1.ListOptions) (result *v1alpha1.StorageQueueList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageQueueList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagequeues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageQueues.
func (c *storageQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storagequeues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storageQueue and creates it.  Returns the server's representation of the storageQueue, and an error, if there is any.
func (c *storageQueues) Create(storageQueue *v1alpha1.StorageQueue) (result *v1alpha1.StorageQueue, err error) {
	result = &v1alpha1.StorageQueue{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storagequeues").
		Body(storageQueue).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageQueue and updates it. Returns the server's representation of the storageQueue, and an error, if there is any.
func (c *storageQueues) Update(storageQueue *v1alpha1.StorageQueue) (result *v1alpha1.StorageQueue, err error) {
	result = &v1alpha1.StorageQueue{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagequeues").
		Name(storageQueue.Name).
		Body(storageQueue).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *storageQueues) UpdateStatus(storageQueue *v1alpha1.StorageQueue) (result *v1alpha1.StorageQueue, err error) {
	result = &v1alpha1.StorageQueue{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagequeues").
		Name(storageQueue.Name).
		SubResource("status").
		Body(storageQueue).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageQueue and deletes it. Returns an error if one occurs.
func (c *storageQueues) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagequeues").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagequeues").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageQueue.
func (c *storageQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageQueue, err error) {
	result = &v1alpha1.StorageQueue{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storagequeues").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
