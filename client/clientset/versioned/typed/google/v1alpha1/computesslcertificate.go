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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeSslCertificatesGetter has a method to return a ComputeSslCertificateInterface.
// A group's client should implement this interface.
type ComputeSslCertificatesGetter interface {
	ComputeSslCertificates(namespace string) ComputeSslCertificateInterface
}

// ComputeSslCertificateInterface has methods to work with ComputeSslCertificate resources.
type ComputeSslCertificateInterface interface {
	Create(*v1alpha1.ComputeSslCertificate) (*v1alpha1.ComputeSslCertificate, error)
	Update(*v1alpha1.ComputeSslCertificate) (*v1alpha1.ComputeSslCertificate, error)
	UpdateStatus(*v1alpha1.ComputeSslCertificate) (*v1alpha1.ComputeSslCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeSslCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeSslCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSslCertificate, err error)
	ComputeSslCertificateExpansion
}

// computeSslCertificates implements ComputeSslCertificateInterface
type computeSslCertificates struct {
	client rest.Interface
	ns     string
}

// newComputeSslCertificates returns a ComputeSslCertificates
func newComputeSslCertificates(c *GoogleV1alpha1Client, namespace string) *computeSslCertificates {
	return &computeSslCertificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeSslCertificate, and returns the corresponding computeSslCertificate object, and an error if there is any.
func (c *computeSslCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSslCertificate, err error) {
	result = &v1alpha1.ComputeSslCertificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeSslCertificates that match those selectors.
func (c *computeSslCertificates) List(opts v1.ListOptions) (result *v1alpha1.ComputeSslCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeSslCertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesslcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeSslCertificates.
func (c *computeSslCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computesslcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeSslCertificate and creates it.  Returns the server's representation of the computeSslCertificate, and an error, if there is any.
func (c *computeSslCertificates) Create(computeSslCertificate *v1alpha1.ComputeSslCertificate) (result *v1alpha1.ComputeSslCertificate, err error) {
	result = &v1alpha1.ComputeSslCertificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Body(computeSslCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeSslCertificate and updates it. Returns the server's representation of the computeSslCertificate, and an error, if there is any.
func (c *computeSslCertificates) Update(computeSslCertificate *v1alpha1.ComputeSslCertificate) (result *v1alpha1.ComputeSslCertificate, err error) {
	result = &v1alpha1.ComputeSslCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(computeSslCertificate.Name).
		Body(computeSslCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeSslCertificates) UpdateStatus(computeSslCertificate *v1alpha1.ComputeSslCertificate) (result *v1alpha1.ComputeSslCertificate, err error) {
	result = &v1alpha1.ComputeSslCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(computeSslCertificate.Name).
		SubResource("status").
		Body(computeSslCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeSslCertificate and deletes it. Returns an error if one occurs.
func (c *computeSslCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeSslCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesslcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeSslCertificate.
func (c *computeSslCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSslCertificate, err error) {
	result = &v1alpha1.ComputeSslCertificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computesslcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
