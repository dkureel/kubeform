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

// ApiGatewayAccountsGetter has a method to return a ApiGatewayAccountInterface.
// A group's client should implement this interface.
type ApiGatewayAccountsGetter interface {
	ApiGatewayAccounts(namespace string) ApiGatewayAccountInterface
}

// ApiGatewayAccountInterface has methods to work with ApiGatewayAccount resources.
type ApiGatewayAccountInterface interface {
	Create(*v1alpha1.ApiGatewayAccount) (*v1alpha1.ApiGatewayAccount, error)
	Update(*v1alpha1.ApiGatewayAccount) (*v1alpha1.ApiGatewayAccount, error)
	UpdateStatus(*v1alpha1.ApiGatewayAccount) (*v1alpha1.ApiGatewayAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiGatewayAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiGatewayAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayAccount, err error)
	ApiGatewayAccountExpansion
}

// apiGatewayAccounts implements ApiGatewayAccountInterface
type apiGatewayAccounts struct {
	client rest.Interface
	ns     string
}

// newApiGatewayAccounts returns a ApiGatewayAccounts
func newApiGatewayAccounts(c *AwsV1alpha1Client, namespace string) *apiGatewayAccounts {
	return &apiGatewayAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiGatewayAccount, and returns the corresponding apiGatewayAccount object, and an error if there is any.
func (c *apiGatewayAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayAccount, err error) {
	result = &v1alpha1.ApiGatewayAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiGatewayAccounts that match those selectors.
func (c *apiGatewayAccounts) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiGatewayAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiGatewayAccounts.
func (c *apiGatewayAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiGatewayAccount and creates it.  Returns the server's representation of the apiGatewayAccount, and an error, if there is any.
func (c *apiGatewayAccounts) Create(apiGatewayAccount *v1alpha1.ApiGatewayAccount) (result *v1alpha1.ApiGatewayAccount, err error) {
	result = &v1alpha1.ApiGatewayAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		Body(apiGatewayAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiGatewayAccount and updates it. Returns the server's representation of the apiGatewayAccount, and an error, if there is any.
func (c *apiGatewayAccounts) Update(apiGatewayAccount *v1alpha1.ApiGatewayAccount) (result *v1alpha1.ApiGatewayAccount, err error) {
	result = &v1alpha1.ApiGatewayAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		Name(apiGatewayAccount.Name).
		Body(apiGatewayAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiGatewayAccounts) UpdateStatus(apiGatewayAccount *v1alpha1.ApiGatewayAccount) (result *v1alpha1.ApiGatewayAccount, err error) {
	result = &v1alpha1.ApiGatewayAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		Name(apiGatewayAccount.Name).
		SubResource("status").
		Body(apiGatewayAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiGatewayAccount and deletes it. Returns an error if one occurs.
func (c *apiGatewayAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiGatewayAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiGatewayAccount.
func (c *apiGatewayAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayAccount, err error) {
	result = &v1alpha1.ApiGatewayAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apigatewayaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
