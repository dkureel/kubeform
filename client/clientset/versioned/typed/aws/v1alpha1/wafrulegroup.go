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

// WafRuleGroupsGetter has a method to return a WafRuleGroupInterface.
// A group's client should implement this interface.
type WafRuleGroupsGetter interface {
	WafRuleGroups(namespace string) WafRuleGroupInterface
}

// WafRuleGroupInterface has methods to work with WafRuleGroup resources.
type WafRuleGroupInterface interface {
	Create(*v1alpha1.WafRuleGroup) (*v1alpha1.WafRuleGroup, error)
	Update(*v1alpha1.WafRuleGroup) (*v1alpha1.WafRuleGroup, error)
	UpdateStatus(*v1alpha1.WafRuleGroup) (*v1alpha1.WafRuleGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafRuleGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.WafRuleGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafRuleGroup, err error)
	WafRuleGroupExpansion
}

// wafRuleGroups implements WafRuleGroupInterface
type wafRuleGroups struct {
	client rest.Interface
	ns     string
}

// newWafRuleGroups returns a WafRuleGroups
func newWafRuleGroups(c *AwsV1alpha1Client, namespace string) *wafRuleGroups {
	return &wafRuleGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafRuleGroup, and returns the corresponding wafRuleGroup object, and an error if there is any.
func (c *wafRuleGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.WafRuleGroup, err error) {
	result = &v1alpha1.WafRuleGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafrulegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafRuleGroups that match those selectors.
func (c *wafRuleGroups) List(opts v1.ListOptions) (result *v1alpha1.WafRuleGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafRuleGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafRuleGroups.
func (c *wafRuleGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafRuleGroup and creates it.  Returns the server's representation of the wafRuleGroup, and an error, if there is any.
func (c *wafRuleGroups) Create(wafRuleGroup *v1alpha1.WafRuleGroup) (result *v1alpha1.WafRuleGroup, err error) {
	result = &v1alpha1.WafRuleGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafrulegroups").
		Body(wafRuleGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafRuleGroup and updates it. Returns the server's representation of the wafRuleGroup, and an error, if there is any.
func (c *wafRuleGroups) Update(wafRuleGroup *v1alpha1.WafRuleGroup) (result *v1alpha1.WafRuleGroup, err error) {
	result = &v1alpha1.WafRuleGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafrulegroups").
		Name(wafRuleGroup.Name).
		Body(wafRuleGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafRuleGroups) UpdateStatus(wafRuleGroup *v1alpha1.WafRuleGroup) (result *v1alpha1.WafRuleGroup, err error) {
	result = &v1alpha1.WafRuleGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafrulegroups").
		Name(wafRuleGroup.Name).
		SubResource("status").
		Body(wafRuleGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafRuleGroup and deletes it. Returns an error if one occurs.
func (c *wafRuleGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafrulegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafRuleGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafrulegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafRuleGroup.
func (c *wafRuleGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafRuleGroup, err error) {
	result = &v1alpha1.WafRuleGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafrulegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
