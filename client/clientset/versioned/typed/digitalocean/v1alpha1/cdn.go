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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// CdnsGetter has a method to return a CdnInterface.
// A group's client should implement this interface.
type CdnsGetter interface {
	Cdns(namespace string) CdnInterface
}

// CdnInterface has methods to work with Cdn resources.
type CdnInterface interface {
	Create(*v1alpha1.Cdn) (*v1alpha1.Cdn, error)
	Update(*v1alpha1.Cdn) (*v1alpha1.Cdn, error)
	UpdateStatus(*v1alpha1.Cdn) (*v1alpha1.Cdn, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Cdn, error)
	List(opts v1.ListOptions) (*v1alpha1.CdnList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Cdn, err error)
	CdnExpansion
}

// cdns implements CdnInterface
type cdns struct {
	client rest.Interface
	ns     string
}

// newCdns returns a Cdns
func newCdns(c *DigitaloceanV1alpha1Client, namespace string) *cdns {
	return &cdns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cdn, and returns the corresponding cdn object, and an error if there is any.
func (c *cdns) Get(name string, options v1.GetOptions) (result *v1alpha1.Cdn, err error) {
	result = &v1alpha1.Cdn{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Cdns that match those selectors.
func (c *cdns) List(opts v1.ListOptions) (result *v1alpha1.CdnList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CdnList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cdns.
func (c *cdns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cdns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cdn and creates it.  Returns the server's representation of the cdn, and an error, if there is any.
func (c *cdns) Create(cdn *v1alpha1.Cdn) (result *v1alpha1.Cdn, err error) {
	result = &v1alpha1.Cdn{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cdns").
		Body(cdn).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cdn and updates it. Returns the server's representation of the cdn, and an error, if there is any.
func (c *cdns) Update(cdn *v1alpha1.Cdn) (result *v1alpha1.Cdn, err error) {
	result = &v1alpha1.Cdn{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdns").
		Name(cdn.Name).
		Body(cdn).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cdns) UpdateStatus(cdn *v1alpha1.Cdn) (result *v1alpha1.Cdn, err error) {
	result = &v1alpha1.Cdn{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdns").
		Name(cdn.Name).
		SubResource("status").
		Body(cdn).
		Do().
		Into(result)
	return
}

// Delete takes name of the cdn and deletes it. Returns an error if one occurs.
func (c *cdns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cdns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cdn.
func (c *cdns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Cdn, err error) {
	result = &v1alpha1.Cdn{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cdns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
