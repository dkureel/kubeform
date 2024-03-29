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

// IamPolicyAttachmentsGetter has a method to return a IamPolicyAttachmentInterface.
// A group's client should implement this interface.
type IamPolicyAttachmentsGetter interface {
	IamPolicyAttachments(namespace string) IamPolicyAttachmentInterface
}

// IamPolicyAttachmentInterface has methods to work with IamPolicyAttachment resources.
type IamPolicyAttachmentInterface interface {
	Create(*v1alpha1.IamPolicyAttachment) (*v1alpha1.IamPolicyAttachment, error)
	Update(*v1alpha1.IamPolicyAttachment) (*v1alpha1.IamPolicyAttachment, error)
	UpdateStatus(*v1alpha1.IamPolicyAttachment) (*v1alpha1.IamPolicyAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamPolicyAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.IamPolicyAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamPolicyAttachment, err error)
	IamPolicyAttachmentExpansion
}

// iamPolicyAttachments implements IamPolicyAttachmentInterface
type iamPolicyAttachments struct {
	client rest.Interface
	ns     string
}

// newIamPolicyAttachments returns a IamPolicyAttachments
func newIamPolicyAttachments(c *AwsV1alpha1Client, namespace string) *iamPolicyAttachments {
	return &iamPolicyAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamPolicyAttachment, and returns the corresponding iamPolicyAttachment object, and an error if there is any.
func (c *iamPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.IamPolicyAttachment, err error) {
	result = &v1alpha1.IamPolicyAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamPolicyAttachments that match those selectors.
func (c *iamPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.IamPolicyAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamPolicyAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamPolicyAttachments.
func (c *iamPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamPolicyAttachment and creates it.  Returns the server's representation of the iamPolicyAttachment, and an error, if there is any.
func (c *iamPolicyAttachments) Create(iamPolicyAttachment *v1alpha1.IamPolicyAttachment) (result *v1alpha1.IamPolicyAttachment, err error) {
	result = &v1alpha1.IamPolicyAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		Body(iamPolicyAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamPolicyAttachment and updates it. Returns the server's representation of the iamPolicyAttachment, and an error, if there is any.
func (c *iamPolicyAttachments) Update(iamPolicyAttachment *v1alpha1.IamPolicyAttachment) (result *v1alpha1.IamPolicyAttachment, err error) {
	result = &v1alpha1.IamPolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		Name(iamPolicyAttachment.Name).
		Body(iamPolicyAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamPolicyAttachments) UpdateStatus(iamPolicyAttachment *v1alpha1.IamPolicyAttachment) (result *v1alpha1.IamPolicyAttachment, err error) {
	result = &v1alpha1.IamPolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		Name(iamPolicyAttachment.Name).
		SubResource("status").
		Body(iamPolicyAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *iamPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iampolicyattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamPolicyAttachment.
func (c *iamPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamPolicyAttachment, err error) {
	result = &v1alpha1.IamPolicyAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iampolicyattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
