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

// ApiManagementApisGetter has a method to return a ApiManagementApiInterface.
// A group's client should implement this interface.
type ApiManagementApisGetter interface {
	ApiManagementApis() ApiManagementApiInterface
}

// ApiManagementApiInterface has methods to work with ApiManagementApi resources.
type ApiManagementApiInterface interface {
	Create(*v1alpha1.ApiManagementApi) (*v1alpha1.ApiManagementApi, error)
	Update(*v1alpha1.ApiManagementApi) (*v1alpha1.ApiManagementApi, error)
	UpdateStatus(*v1alpha1.ApiManagementApi) (*v1alpha1.ApiManagementApi, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementApi, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementApiList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementApi, err error)
	ApiManagementApiExpansion
}

// apiManagementApis implements ApiManagementApiInterface
type apiManagementApis struct {
	client rest.Interface
}

// newApiManagementApis returns a ApiManagementApis
func newApiManagementApis(c *AzurermV1alpha1Client) *apiManagementApis {
	return &apiManagementApis{
		client: c.RESTClient(),
	}
}

// Get takes name of the apiManagementApi, and returns the corresponding apiManagementApi object, and an error if there is any.
func (c *apiManagementApis) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementApi, err error) {
	result = &v1alpha1.ApiManagementApi{}
	err = c.client.Get().
		Resource("apimanagementapis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementApis that match those selectors.
func (c *apiManagementApis) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementApiList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementApiList{}
	err = c.client.Get().
		Resource("apimanagementapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementApis.
func (c *apiManagementApis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apimanagementapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementApi and creates it.  Returns the server's representation of the apiManagementApi, and an error, if there is any.
func (c *apiManagementApis) Create(apiManagementApi *v1alpha1.ApiManagementApi) (result *v1alpha1.ApiManagementApi, err error) {
	result = &v1alpha1.ApiManagementApi{}
	err = c.client.Post().
		Resource("apimanagementapis").
		Body(apiManagementApi).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementApi and updates it. Returns the server's representation of the apiManagementApi, and an error, if there is any.
func (c *apiManagementApis) Update(apiManagementApi *v1alpha1.ApiManagementApi) (result *v1alpha1.ApiManagementApi, err error) {
	result = &v1alpha1.ApiManagementApi{}
	err = c.client.Put().
		Resource("apimanagementapis").
		Name(apiManagementApi.Name).
		Body(apiManagementApi).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementApis) UpdateStatus(apiManagementApi *v1alpha1.ApiManagementApi) (result *v1alpha1.ApiManagementApi, err error) {
	result = &v1alpha1.ApiManagementApi{}
	err = c.client.Put().
		Resource("apimanagementapis").
		Name(apiManagementApi.Name).
		SubResource("status").
		Body(apiManagementApi).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementApi and deletes it. Returns an error if one occurs.
func (c *apiManagementApis) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apimanagementapis").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementApis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apimanagementapis").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementApi.
func (c *apiManagementApis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementApi, err error) {
	result = &v1alpha1.ApiManagementApi{}
	err = c.client.Patch(pt).
		Resource("apimanagementapis").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}