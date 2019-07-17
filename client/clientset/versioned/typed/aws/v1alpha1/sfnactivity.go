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

// SfnActivitiesGetter has a method to return a SfnActivityInterface.
// A group's client should implement this interface.
type SfnActivitiesGetter interface {
	SfnActivities(namespace string) SfnActivityInterface
}

// SfnActivityInterface has methods to work with SfnActivity resources.
type SfnActivityInterface interface {
	Create(*v1alpha1.SfnActivity) (*v1alpha1.SfnActivity, error)
	Update(*v1alpha1.SfnActivity) (*v1alpha1.SfnActivity, error)
	UpdateStatus(*v1alpha1.SfnActivity) (*v1alpha1.SfnActivity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SfnActivity, error)
	List(opts v1.ListOptions) (*v1alpha1.SfnActivityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SfnActivity, err error)
	SfnActivityExpansion
}

// sfnActivities implements SfnActivityInterface
type sfnActivities struct {
	client rest.Interface
	ns     string
}

// newSfnActivities returns a SfnActivities
func newSfnActivities(c *AwsV1alpha1Client, namespace string) *sfnActivities {
	return &sfnActivities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sfnActivity, and returns the corresponding sfnActivity object, and an error if there is any.
func (c *sfnActivities) Get(name string, options v1.GetOptions) (result *v1alpha1.SfnActivity, err error) {
	result = &v1alpha1.SfnActivity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfnactivities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SfnActivities that match those selectors.
func (c *sfnActivities) List(opts v1.ListOptions) (result *v1alpha1.SfnActivityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SfnActivityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfnactivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sfnActivities.
func (c *sfnActivities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sfnactivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sfnActivity and creates it.  Returns the server's representation of the sfnActivity, and an error, if there is any.
func (c *sfnActivities) Create(sfnActivity *v1alpha1.SfnActivity) (result *v1alpha1.SfnActivity, err error) {
	result = &v1alpha1.SfnActivity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sfnactivities").
		Body(sfnActivity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sfnActivity and updates it. Returns the server's representation of the sfnActivity, and an error, if there is any.
func (c *sfnActivities) Update(sfnActivity *v1alpha1.SfnActivity) (result *v1alpha1.SfnActivity, err error) {
	result = &v1alpha1.SfnActivity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfnactivities").
		Name(sfnActivity.Name).
		Body(sfnActivity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sfnActivities) UpdateStatus(sfnActivity *v1alpha1.SfnActivity) (result *v1alpha1.SfnActivity, err error) {
	result = &v1alpha1.SfnActivity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfnactivities").
		Name(sfnActivity.Name).
		SubResource("status").
		Body(sfnActivity).
		Do().
		Into(result)
	return
}

// Delete takes name of the sfnActivity and deletes it. Returns an error if one occurs.
func (c *sfnActivities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfnactivities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sfnActivities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfnactivities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sfnActivity.
func (c *sfnActivities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SfnActivity, err error) {
	result = &v1alpha1.SfnActivity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sfnactivities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
