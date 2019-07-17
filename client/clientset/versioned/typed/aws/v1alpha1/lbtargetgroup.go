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

// LbTargetGroupsGetter has a method to return a LbTargetGroupInterface.
// A group's client should implement this interface.
type LbTargetGroupsGetter interface {
	LbTargetGroups(namespace string) LbTargetGroupInterface
}

// LbTargetGroupInterface has methods to work with LbTargetGroup resources.
type LbTargetGroupInterface interface {
	Create(*v1alpha1.LbTargetGroup) (*v1alpha1.LbTargetGroup, error)
	Update(*v1alpha1.LbTargetGroup) (*v1alpha1.LbTargetGroup, error)
	UpdateStatus(*v1alpha1.LbTargetGroup) (*v1alpha1.LbTargetGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LbTargetGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.LbTargetGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbTargetGroup, err error)
	LbTargetGroupExpansion
}

// lbTargetGroups implements LbTargetGroupInterface
type lbTargetGroups struct {
	client rest.Interface
	ns     string
}

// newLbTargetGroups returns a LbTargetGroups
func newLbTargetGroups(c *AwsV1alpha1Client, namespace string) *lbTargetGroups {
	return &lbTargetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lbTargetGroup, and returns the corresponding lbTargetGroup object, and an error if there is any.
func (c *lbTargetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.LbTargetGroup, err error) {
	result = &v1alpha1.LbTargetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LbTargetGroups that match those selectors.
func (c *lbTargetGroups) List(opts v1.ListOptions) (result *v1alpha1.LbTargetGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LbTargetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lbTargetGroups.
func (c *lbTargetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lbTargetGroup and creates it.  Returns the server's representation of the lbTargetGroup, and an error, if there is any.
func (c *lbTargetGroups) Create(lbTargetGroup *v1alpha1.LbTargetGroup) (result *v1alpha1.LbTargetGroup, err error) {
	result = &v1alpha1.LbTargetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		Body(lbTargetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lbTargetGroup and updates it. Returns the server's representation of the lbTargetGroup, and an error, if there is any.
func (c *lbTargetGroups) Update(lbTargetGroup *v1alpha1.LbTargetGroup) (result *v1alpha1.LbTargetGroup, err error) {
	result = &v1alpha1.LbTargetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		Name(lbTargetGroup.Name).
		Body(lbTargetGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lbTargetGroups) UpdateStatus(lbTargetGroup *v1alpha1.LbTargetGroup) (result *v1alpha1.LbTargetGroup, err error) {
	result = &v1alpha1.LbTargetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		Name(lbTargetGroup.Name).
		SubResource("status").
		Body(lbTargetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the lbTargetGroup and deletes it. Returns an error if one occurs.
func (c *lbTargetGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lbTargetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbtargetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lbTargetGroup.
func (c *lbTargetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbTargetGroup, err error) {
	result = &v1alpha1.LbTargetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lbtargetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
