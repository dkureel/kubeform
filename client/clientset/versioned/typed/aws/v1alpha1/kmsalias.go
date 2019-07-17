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

// KmsAliasesGetter has a method to return a KmsAliasInterface.
// A group's client should implement this interface.
type KmsAliasesGetter interface {
	KmsAliases(namespace string) KmsAliasInterface
}

// KmsAliasInterface has methods to work with KmsAlias resources.
type KmsAliasInterface interface {
	Create(*v1alpha1.KmsAlias) (*v1alpha1.KmsAlias, error)
	Update(*v1alpha1.KmsAlias) (*v1alpha1.KmsAlias, error)
	UpdateStatus(*v1alpha1.KmsAlias) (*v1alpha1.KmsAlias, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KmsAlias, error)
	List(opts v1.ListOptions) (*v1alpha1.KmsAliasList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsAlias, err error)
	KmsAliasExpansion
}

// kmsAliases implements KmsAliasInterface
type kmsAliases struct {
	client rest.Interface
	ns     string
}

// newKmsAliases returns a KmsAliases
func newKmsAliases(c *AwsV1alpha1Client, namespace string) *kmsAliases {
	return &kmsAliases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kmsAlias, and returns the corresponding kmsAlias object, and an error if there is any.
func (c *kmsAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsAlias, err error) {
	result = &v1alpha1.KmsAlias{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmsaliases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KmsAliases that match those selectors.
func (c *kmsAliases) List(opts v1.ListOptions) (result *v1alpha1.KmsAliasList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KmsAliasList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmsaliases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kmsAliases.
func (c *kmsAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kmsaliases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kmsAlias and creates it.  Returns the server's representation of the kmsAlias, and an error, if there is any.
func (c *kmsAliases) Create(kmsAlias *v1alpha1.KmsAlias) (result *v1alpha1.KmsAlias, err error) {
	result = &v1alpha1.KmsAlias{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kmsaliases").
		Body(kmsAlias).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kmsAlias and updates it. Returns the server's representation of the kmsAlias, and an error, if there is any.
func (c *kmsAliases) Update(kmsAlias *v1alpha1.KmsAlias) (result *v1alpha1.KmsAlias, err error) {
	result = &v1alpha1.KmsAlias{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmsaliases").
		Name(kmsAlias.Name).
		Body(kmsAlias).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kmsAliases) UpdateStatus(kmsAlias *v1alpha1.KmsAlias) (result *v1alpha1.KmsAlias, err error) {
	result = &v1alpha1.KmsAlias{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmsaliases").
		Name(kmsAlias.Name).
		SubResource("status").
		Body(kmsAlias).
		Do().
		Into(result)
	return
}

// Delete takes name of the kmsAlias and deletes it. Returns an error if one occurs.
func (c *kmsAliases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmsaliases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kmsAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmsaliases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kmsAlias.
func (c *kmsAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsAlias, err error) {
	result = &v1alpha1.KmsAlias{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kmsaliases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
