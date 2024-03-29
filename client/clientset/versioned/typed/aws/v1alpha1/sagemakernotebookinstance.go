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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SagemakerNotebookInstancesGetter has a method to return a SagemakerNotebookInstanceInterface.
// A group's client should implement this interface.
type SagemakerNotebookInstancesGetter interface {
	SagemakerNotebookInstances(namespace string) SagemakerNotebookInstanceInterface
}

// SagemakerNotebookInstanceInterface has methods to work with SagemakerNotebookInstance resources.
type SagemakerNotebookInstanceInterface interface {
	Create(*v1alpha1.SagemakerNotebookInstance) (*v1alpha1.SagemakerNotebookInstance, error)
	Update(*v1alpha1.SagemakerNotebookInstance) (*v1alpha1.SagemakerNotebookInstance, error)
	UpdateStatus(*v1alpha1.SagemakerNotebookInstance) (*v1alpha1.SagemakerNotebookInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SagemakerNotebookInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.SagemakerNotebookInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SagemakerNotebookInstance, err error)
	SagemakerNotebookInstanceExpansion
}

// sagemakerNotebookInstances implements SagemakerNotebookInstanceInterface
type sagemakerNotebookInstances struct {
	client rest.Interface
	ns     string
}

// newSagemakerNotebookInstances returns a SagemakerNotebookInstances
func newSagemakerNotebookInstances(c *AwsV1alpha1Client, namespace string) *sagemakerNotebookInstances {
	return &sagemakerNotebookInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sagemakerNotebookInstance, and returns the corresponding sagemakerNotebookInstance object, and an error if there is any.
func (c *sagemakerNotebookInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.SagemakerNotebookInstance, err error) {
	result = &v1alpha1.SagemakerNotebookInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SagemakerNotebookInstances that match those selectors.
func (c *sagemakerNotebookInstances) List(opts v1.ListOptions) (result *v1alpha1.SagemakerNotebookInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SagemakerNotebookInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sagemakerNotebookInstances.
func (c *sagemakerNotebookInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sagemakerNotebookInstance and creates it.  Returns the server's representation of the sagemakerNotebookInstance, and an error, if there is any.
func (c *sagemakerNotebookInstances) Create(sagemakerNotebookInstance *v1alpha1.SagemakerNotebookInstance) (result *v1alpha1.SagemakerNotebookInstance, err error) {
	result = &v1alpha1.SagemakerNotebookInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		Body(sagemakerNotebookInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sagemakerNotebookInstance and updates it. Returns the server's representation of the sagemakerNotebookInstance, and an error, if there is any.
func (c *sagemakerNotebookInstances) Update(sagemakerNotebookInstance *v1alpha1.SagemakerNotebookInstance) (result *v1alpha1.SagemakerNotebookInstance, err error) {
	result = &v1alpha1.SagemakerNotebookInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		Name(sagemakerNotebookInstance.Name).
		Body(sagemakerNotebookInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sagemakerNotebookInstances) UpdateStatus(sagemakerNotebookInstance *v1alpha1.SagemakerNotebookInstance) (result *v1alpha1.SagemakerNotebookInstance, err error) {
	result = &v1alpha1.SagemakerNotebookInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		Name(sagemakerNotebookInstance.Name).
		SubResource("status").
		Body(sagemakerNotebookInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the sagemakerNotebookInstance and deletes it. Returns an error if one occurs.
func (c *sagemakerNotebookInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sagemakerNotebookInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sagemakerNotebookInstance.
func (c *sagemakerNotebookInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SagemakerNotebookInstance, err error) {
	result = &v1alpha1.SagemakerNotebookInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sagemakernotebookinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
