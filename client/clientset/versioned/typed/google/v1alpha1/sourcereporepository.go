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

// SourcerepoRepositoriesGetter has a method to return a SourcerepoRepositoryInterface.
// A group's client should implement this interface.
type SourcerepoRepositoriesGetter interface {
	SourcerepoRepositories(namespace string) SourcerepoRepositoryInterface
}

// SourcerepoRepositoryInterface has methods to work with SourcerepoRepository resources.
type SourcerepoRepositoryInterface interface {
	Create(*v1alpha1.SourcerepoRepository) (*v1alpha1.SourcerepoRepository, error)
	Update(*v1alpha1.SourcerepoRepository) (*v1alpha1.SourcerepoRepository, error)
	UpdateStatus(*v1alpha1.SourcerepoRepository) (*v1alpha1.SourcerepoRepository, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SourcerepoRepository, error)
	List(opts v1.ListOptions) (*v1alpha1.SourcerepoRepositoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SourcerepoRepository, err error)
	SourcerepoRepositoryExpansion
}

// sourcerepoRepositories implements SourcerepoRepositoryInterface
type sourcerepoRepositories struct {
	client rest.Interface
	ns     string
}

// newSourcerepoRepositories returns a SourcerepoRepositories
func newSourcerepoRepositories(c *GoogleV1alpha1Client, namespace string) *sourcerepoRepositories {
	return &sourcerepoRepositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sourcerepoRepository, and returns the corresponding sourcerepoRepository object, and an error if there is any.
func (c *sourcerepoRepositories) Get(name string, options v1.GetOptions) (result *v1alpha1.SourcerepoRepository, err error) {
	result = &v1alpha1.SourcerepoRepository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SourcerepoRepositories that match those selectors.
func (c *sourcerepoRepositories) List(opts v1.ListOptions) (result *v1alpha1.SourcerepoRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SourcerepoRepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sourcerepoRepositories.
func (c *sourcerepoRepositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sourcerepoRepository and creates it.  Returns the server's representation of the sourcerepoRepository, and an error, if there is any.
func (c *sourcerepoRepositories) Create(sourcerepoRepository *v1alpha1.SourcerepoRepository) (result *v1alpha1.SourcerepoRepository, err error) {
	result = &v1alpha1.SourcerepoRepository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		Body(sourcerepoRepository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sourcerepoRepository and updates it. Returns the server's representation of the sourcerepoRepository, and an error, if there is any.
func (c *sourcerepoRepositories) Update(sourcerepoRepository *v1alpha1.SourcerepoRepository) (result *v1alpha1.SourcerepoRepository, err error) {
	result = &v1alpha1.SourcerepoRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		Name(sourcerepoRepository.Name).
		Body(sourcerepoRepository).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sourcerepoRepositories) UpdateStatus(sourcerepoRepository *v1alpha1.SourcerepoRepository) (result *v1alpha1.SourcerepoRepository, err error) {
	result = &v1alpha1.SourcerepoRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		Name(sourcerepoRepository.Name).
		SubResource("status").
		Body(sourcerepoRepository).
		Do().
		Into(result)
	return
}

// Delete takes name of the sourcerepoRepository and deletes it. Returns an error if one occurs.
func (c *sourcerepoRepositories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sourcerepoRepositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcereporepositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sourcerepoRepository.
func (c *sourcerepoRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SourcerepoRepository, err error) {
	result = &v1alpha1.SourcerepoRepository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sourcereporepositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
