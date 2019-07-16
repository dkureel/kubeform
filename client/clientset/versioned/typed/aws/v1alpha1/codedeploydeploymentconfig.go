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

// CodedeployDeploymentConfigsGetter has a method to return a CodedeployDeploymentConfigInterface.
// A group's client should implement this interface.
type CodedeployDeploymentConfigsGetter interface {
	CodedeployDeploymentConfigs() CodedeployDeploymentConfigInterface
}

// CodedeployDeploymentConfigInterface has methods to work with CodedeployDeploymentConfig resources.
type CodedeployDeploymentConfigInterface interface {
	Create(*v1alpha1.CodedeployDeploymentConfig) (*v1alpha1.CodedeployDeploymentConfig, error)
	Update(*v1alpha1.CodedeployDeploymentConfig) (*v1alpha1.CodedeployDeploymentConfig, error)
	UpdateStatus(*v1alpha1.CodedeployDeploymentConfig) (*v1alpha1.CodedeployDeploymentConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CodedeployDeploymentConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.CodedeployDeploymentConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodedeployDeploymentConfig, err error)
	CodedeployDeploymentConfigExpansion
}

// codedeployDeploymentConfigs implements CodedeployDeploymentConfigInterface
type codedeployDeploymentConfigs struct {
	client rest.Interface
}

// newCodedeployDeploymentConfigs returns a CodedeployDeploymentConfigs
func newCodedeployDeploymentConfigs(c *AwsV1alpha1Client) *codedeployDeploymentConfigs {
	return &codedeployDeploymentConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the codedeployDeploymentConfig, and returns the corresponding codedeployDeploymentConfig object, and an error if there is any.
func (c *codedeployDeploymentConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	result = &v1alpha1.CodedeployDeploymentConfig{}
	err = c.client.Get().
		Resource("codedeploydeploymentconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CodedeployDeploymentConfigs that match those selectors.
func (c *codedeployDeploymentConfigs) List(opts v1.ListOptions) (result *v1alpha1.CodedeployDeploymentConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CodedeployDeploymentConfigList{}
	err = c.client.Get().
		Resource("codedeploydeploymentconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested codedeployDeploymentConfigs.
func (c *codedeployDeploymentConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("codedeploydeploymentconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a codedeployDeploymentConfig and creates it.  Returns the server's representation of the codedeployDeploymentConfig, and an error, if there is any.
func (c *codedeployDeploymentConfigs) Create(codedeployDeploymentConfig *v1alpha1.CodedeployDeploymentConfig) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	result = &v1alpha1.CodedeployDeploymentConfig{}
	err = c.client.Post().
		Resource("codedeploydeploymentconfigs").
		Body(codedeployDeploymentConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a codedeployDeploymentConfig and updates it. Returns the server's representation of the codedeployDeploymentConfig, and an error, if there is any.
func (c *codedeployDeploymentConfigs) Update(codedeployDeploymentConfig *v1alpha1.CodedeployDeploymentConfig) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	result = &v1alpha1.CodedeployDeploymentConfig{}
	err = c.client.Put().
		Resource("codedeploydeploymentconfigs").
		Name(codedeployDeploymentConfig.Name).
		Body(codedeployDeploymentConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *codedeployDeploymentConfigs) UpdateStatus(codedeployDeploymentConfig *v1alpha1.CodedeployDeploymentConfig) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	result = &v1alpha1.CodedeployDeploymentConfig{}
	err = c.client.Put().
		Resource("codedeploydeploymentconfigs").
		Name(codedeployDeploymentConfig.Name).
		SubResource("status").
		Body(codedeployDeploymentConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the codedeployDeploymentConfig and deletes it. Returns an error if one occurs.
func (c *codedeployDeploymentConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("codedeploydeploymentconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *codedeployDeploymentConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("codedeploydeploymentconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched codedeployDeploymentConfig.
func (c *codedeployDeploymentConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	result = &v1alpha1.CodedeployDeploymentConfig{}
	err = c.client.Patch(pt).
		Resource("codedeploydeploymentconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}