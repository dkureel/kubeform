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

// WafSizeConstraintSetsGetter has a method to return a WafSizeConstraintSetInterface.
// A group's client should implement this interface.
type WafSizeConstraintSetsGetter interface {
	WafSizeConstraintSets(namespace string) WafSizeConstraintSetInterface
}

// WafSizeConstraintSetInterface has methods to work with WafSizeConstraintSet resources.
type WafSizeConstraintSetInterface interface {
	Create(*v1alpha1.WafSizeConstraintSet) (*v1alpha1.WafSizeConstraintSet, error)
	Update(*v1alpha1.WafSizeConstraintSet) (*v1alpha1.WafSizeConstraintSet, error)
	UpdateStatus(*v1alpha1.WafSizeConstraintSet) (*v1alpha1.WafSizeConstraintSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafSizeConstraintSet, error)
	List(opts v1.ListOptions) (*v1alpha1.WafSizeConstraintSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafSizeConstraintSet, err error)
	WafSizeConstraintSetExpansion
}

// wafSizeConstraintSets implements WafSizeConstraintSetInterface
type wafSizeConstraintSets struct {
	client rest.Interface
	ns     string
}

// newWafSizeConstraintSets returns a WafSizeConstraintSets
func newWafSizeConstraintSets(c *AwsV1alpha1Client, namespace string) *wafSizeConstraintSets {
	return &wafSizeConstraintSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafSizeConstraintSet, and returns the corresponding wafSizeConstraintSet object, and an error if there is any.
func (c *wafSizeConstraintSets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafSizeConstraintSet, err error) {
	result = &v1alpha1.WafSizeConstraintSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafSizeConstraintSets that match those selectors.
func (c *wafSizeConstraintSets) List(opts v1.ListOptions) (result *v1alpha1.WafSizeConstraintSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafSizeConstraintSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafSizeConstraintSets.
func (c *wafSizeConstraintSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafSizeConstraintSet and creates it.  Returns the server's representation of the wafSizeConstraintSet, and an error, if there is any.
func (c *wafSizeConstraintSets) Create(wafSizeConstraintSet *v1alpha1.WafSizeConstraintSet) (result *v1alpha1.WafSizeConstraintSet, err error) {
	result = &v1alpha1.WafSizeConstraintSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		Body(wafSizeConstraintSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafSizeConstraintSet and updates it. Returns the server's representation of the wafSizeConstraintSet, and an error, if there is any.
func (c *wafSizeConstraintSets) Update(wafSizeConstraintSet *v1alpha1.WafSizeConstraintSet) (result *v1alpha1.WafSizeConstraintSet, err error) {
	result = &v1alpha1.WafSizeConstraintSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		Name(wafSizeConstraintSet.Name).
		Body(wafSizeConstraintSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafSizeConstraintSets) UpdateStatus(wafSizeConstraintSet *v1alpha1.WafSizeConstraintSet) (result *v1alpha1.WafSizeConstraintSet, err error) {
	result = &v1alpha1.WafSizeConstraintSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		Name(wafSizeConstraintSet.Name).
		SubResource("status").
		Body(wafSizeConstraintSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafSizeConstraintSet and deletes it. Returns an error if one occurs.
func (c *wafSizeConstraintSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafSizeConstraintSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafSizeConstraintSet.
func (c *wafSizeConstraintSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafSizeConstraintSet, err error) {
	result = &v1alpha1.WafSizeConstraintSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafsizeconstraintsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
