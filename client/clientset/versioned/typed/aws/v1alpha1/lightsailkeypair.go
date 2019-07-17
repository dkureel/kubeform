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

// LightsailKeyPairsGetter has a method to return a LightsailKeyPairInterface.
// A group's client should implement this interface.
type LightsailKeyPairsGetter interface {
	LightsailKeyPairs(namespace string) LightsailKeyPairInterface
}

// LightsailKeyPairInterface has methods to work with LightsailKeyPair resources.
type LightsailKeyPairInterface interface {
	Create(*v1alpha1.LightsailKeyPair) (*v1alpha1.LightsailKeyPair, error)
	Update(*v1alpha1.LightsailKeyPair) (*v1alpha1.LightsailKeyPair, error)
	UpdateStatus(*v1alpha1.LightsailKeyPair) (*v1alpha1.LightsailKeyPair, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LightsailKeyPair, error)
	List(opts v1.ListOptions) (*v1alpha1.LightsailKeyPairList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailKeyPair, err error)
	LightsailKeyPairExpansion
}

// lightsailKeyPairs implements LightsailKeyPairInterface
type lightsailKeyPairs struct {
	client rest.Interface
	ns     string
}

// newLightsailKeyPairs returns a LightsailKeyPairs
func newLightsailKeyPairs(c *AwsV1alpha1Client, namespace string) *lightsailKeyPairs {
	return &lightsailKeyPairs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lightsailKeyPair, and returns the corresponding lightsailKeyPair object, and an error if there is any.
func (c *lightsailKeyPairs) Get(name string, options v1.GetOptions) (result *v1alpha1.LightsailKeyPair, err error) {
	result = &v1alpha1.LightsailKeyPair{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LightsailKeyPairs that match those selectors.
func (c *lightsailKeyPairs) List(opts v1.ListOptions) (result *v1alpha1.LightsailKeyPairList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LightsailKeyPairList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lightsailKeyPairs.
func (c *lightsailKeyPairs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lightsailKeyPair and creates it.  Returns the server's representation of the lightsailKeyPair, and an error, if there is any.
func (c *lightsailKeyPairs) Create(lightsailKeyPair *v1alpha1.LightsailKeyPair) (result *v1alpha1.LightsailKeyPair, err error) {
	result = &v1alpha1.LightsailKeyPair{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		Body(lightsailKeyPair).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lightsailKeyPair and updates it. Returns the server's representation of the lightsailKeyPair, and an error, if there is any.
func (c *lightsailKeyPairs) Update(lightsailKeyPair *v1alpha1.LightsailKeyPair) (result *v1alpha1.LightsailKeyPair, err error) {
	result = &v1alpha1.LightsailKeyPair{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		Name(lightsailKeyPair.Name).
		Body(lightsailKeyPair).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lightsailKeyPairs) UpdateStatus(lightsailKeyPair *v1alpha1.LightsailKeyPair) (result *v1alpha1.LightsailKeyPair, err error) {
	result = &v1alpha1.LightsailKeyPair{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		Name(lightsailKeyPair.Name).
		SubResource("status").
		Body(lightsailKeyPair).
		Do().
		Into(result)
	return
}

// Delete takes name of the lightsailKeyPair and deletes it. Returns an error if one occurs.
func (c *lightsailKeyPairs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lightsailKeyPairs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lightsailKeyPair.
func (c *lightsailKeyPairs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailKeyPair, err error) {
	result = &v1alpha1.LightsailKeyPair{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lightsailkeypairs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
