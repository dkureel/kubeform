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

// DxLagsGetter has a method to return a DxLagInterface.
// A group's client should implement this interface.
type DxLagsGetter interface {
	DxLags(namespace string) DxLagInterface
}

// DxLagInterface has methods to work with DxLag resources.
type DxLagInterface interface {
	Create(*v1alpha1.DxLag) (*v1alpha1.DxLag, error)
	Update(*v1alpha1.DxLag) (*v1alpha1.DxLag, error)
	UpdateStatus(*v1alpha1.DxLag) (*v1alpha1.DxLag, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DxLag, error)
	List(opts v1.ListOptions) (*v1alpha1.DxLagList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxLag, err error)
	DxLagExpansion
}

// dxLags implements DxLagInterface
type dxLags struct {
	client rest.Interface
	ns     string
}

// newDxLags returns a DxLags
func newDxLags(c *AwsV1alpha1Client, namespace string) *dxLags {
	return &dxLags{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dxLag, and returns the corresponding dxLag object, and an error if there is any.
func (c *dxLags) Get(name string, options v1.GetOptions) (result *v1alpha1.DxLag, err error) {
	result = &v1alpha1.DxLag{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxlags").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DxLags that match those selectors.
func (c *dxLags) List(opts v1.ListOptions) (result *v1alpha1.DxLagList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DxLagList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxlags").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dxLags.
func (c *dxLags) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dxlags").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dxLag and creates it.  Returns the server's representation of the dxLag, and an error, if there is any.
func (c *dxLags) Create(dxLag *v1alpha1.DxLag) (result *v1alpha1.DxLag, err error) {
	result = &v1alpha1.DxLag{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dxlags").
		Body(dxLag).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dxLag and updates it. Returns the server's representation of the dxLag, and an error, if there is any.
func (c *dxLags) Update(dxLag *v1alpha1.DxLag) (result *v1alpha1.DxLag, err error) {
	result = &v1alpha1.DxLag{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxlags").
		Name(dxLag.Name).
		Body(dxLag).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dxLags) UpdateStatus(dxLag *v1alpha1.DxLag) (result *v1alpha1.DxLag, err error) {
	result = &v1alpha1.DxLag{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxlags").
		Name(dxLag.Name).
		SubResource("status").
		Body(dxLag).
		Do().
		Into(result)
	return
}

// Delete takes name of the dxLag and deletes it. Returns an error if one occurs.
func (c *dxLags) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxlags").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dxLags) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxlags").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dxLag.
func (c *dxLags) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxLag, err error) {
	result = &v1alpha1.DxLag{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dxlags").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
