/*
Copyright The Kubeform Authors.

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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudwatchLogStreamsGetter has a method to return a CloudwatchLogStreamInterface.
// A group's client should implement this interface.
type CloudwatchLogStreamsGetter interface {
	CloudwatchLogStreams(namespace string) CloudwatchLogStreamInterface
}

// CloudwatchLogStreamInterface has methods to work with CloudwatchLogStream resources.
type CloudwatchLogStreamInterface interface {
	Create(*v1alpha1.CloudwatchLogStream) (*v1alpha1.CloudwatchLogStream, error)
	Update(*v1alpha1.CloudwatchLogStream) (*v1alpha1.CloudwatchLogStream, error)
	UpdateStatus(*v1alpha1.CloudwatchLogStream) (*v1alpha1.CloudwatchLogStream, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudwatchLogStream, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudwatchLogStreamList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogStream, err error)
	CloudwatchLogStreamExpansion
}

// cloudwatchLogStreams implements CloudwatchLogStreamInterface
type cloudwatchLogStreams struct {
	client rest.Interface
	ns     string
}

// newCloudwatchLogStreams returns a CloudwatchLogStreams
func newCloudwatchLogStreams(c *AwsV1alpha1Client, namespace string) *cloudwatchLogStreams {
	return &cloudwatchLogStreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudwatchLogStream, and returns the corresponding cloudwatchLogStream object, and an error if there is any.
func (c *cloudwatchLogStreams) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogStream, err error) {
	result = &v1alpha1.CloudwatchLogStream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudwatchLogStreams that match those selectors.
func (c *cloudwatchLogStreams) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogStreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudwatchLogStreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogStreams.
func (c *cloudwatchLogStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudwatchLogStream and creates it.  Returns the server's representation of the cloudwatchLogStream, and an error, if there is any.
func (c *cloudwatchLogStreams) Create(cloudwatchLogStream *v1alpha1.CloudwatchLogStream) (result *v1alpha1.CloudwatchLogStream, err error) {
	result = &v1alpha1.CloudwatchLogStream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		Body(cloudwatchLogStream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudwatchLogStream and updates it. Returns the server's representation of the cloudwatchLogStream, and an error, if there is any.
func (c *cloudwatchLogStreams) Update(cloudwatchLogStream *v1alpha1.CloudwatchLogStream) (result *v1alpha1.CloudwatchLogStream, err error) {
	result = &v1alpha1.CloudwatchLogStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		Name(cloudwatchLogStream.Name).
		Body(cloudwatchLogStream).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudwatchLogStreams) UpdateStatus(cloudwatchLogStream *v1alpha1.CloudwatchLogStream) (result *v1alpha1.CloudwatchLogStream, err error) {
	result = &v1alpha1.CloudwatchLogStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		Name(cloudwatchLogStream.Name).
		SubResource("status").
		Body(cloudwatchLogStream).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudwatchLogStream and deletes it. Returns an error if one occurs.
func (c *cloudwatchLogStreams) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudwatchLogStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudwatchLogStream.
func (c *cloudwatchLogStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogStream, err error) {
	result = &v1alpha1.CloudwatchLogStream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudwatchlogstreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
