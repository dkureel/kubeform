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

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStoragegatewayUploadBuffers implements StoragegatewayUploadBufferInterface
type FakeStoragegatewayUploadBuffers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var storagegatewayuploadbuffersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "storagegatewayuploadbuffers"}

var storagegatewayuploadbuffersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "StoragegatewayUploadBuffer"}

// Get takes name of the storagegatewayUploadBuffer, and returns the corresponding storagegatewayUploadBuffer object, and an error if there is any.
func (c *FakeStoragegatewayUploadBuffers) Get(name string, options v1.GetOptions) (result *v1alpha1.StoragegatewayUploadBuffer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagegatewayuploadbuffersResource, c.ns, name), &v1alpha1.StoragegatewayUploadBuffer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayUploadBuffer), err
}

// List takes label and field selectors, and returns the list of StoragegatewayUploadBuffers that match those selectors.
func (c *FakeStoragegatewayUploadBuffers) List(opts v1.ListOptions) (result *v1alpha1.StoragegatewayUploadBufferList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagegatewayuploadbuffersResource, storagegatewayuploadbuffersKind, c.ns, opts), &v1alpha1.StoragegatewayUploadBufferList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StoragegatewayUploadBufferList{ListMeta: obj.(*v1alpha1.StoragegatewayUploadBufferList).ListMeta}
	for _, item := range obj.(*v1alpha1.StoragegatewayUploadBufferList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storagegatewayUploadBuffers.
func (c *FakeStoragegatewayUploadBuffers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagegatewayuploadbuffersResource, c.ns, opts))

}

// Create takes the representation of a storagegatewayUploadBuffer and creates it.  Returns the server's representation of the storagegatewayUploadBuffer, and an error, if there is any.
func (c *FakeStoragegatewayUploadBuffers) Create(storagegatewayUploadBuffer *v1alpha1.StoragegatewayUploadBuffer) (result *v1alpha1.StoragegatewayUploadBuffer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagegatewayuploadbuffersResource, c.ns, storagegatewayUploadBuffer), &v1alpha1.StoragegatewayUploadBuffer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayUploadBuffer), err
}

// Update takes the representation of a storagegatewayUploadBuffer and updates it. Returns the server's representation of the storagegatewayUploadBuffer, and an error, if there is any.
func (c *FakeStoragegatewayUploadBuffers) Update(storagegatewayUploadBuffer *v1alpha1.StoragegatewayUploadBuffer) (result *v1alpha1.StoragegatewayUploadBuffer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagegatewayuploadbuffersResource, c.ns, storagegatewayUploadBuffer), &v1alpha1.StoragegatewayUploadBuffer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayUploadBuffer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStoragegatewayUploadBuffers) UpdateStatus(storagegatewayUploadBuffer *v1alpha1.StoragegatewayUploadBuffer) (*v1alpha1.StoragegatewayUploadBuffer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagegatewayuploadbuffersResource, "status", c.ns, storagegatewayUploadBuffer), &v1alpha1.StoragegatewayUploadBuffer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayUploadBuffer), err
}

// Delete takes name of the storagegatewayUploadBuffer and deletes it. Returns an error if one occurs.
func (c *FakeStoragegatewayUploadBuffers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagegatewayuploadbuffersResource, c.ns, name), &v1alpha1.StoragegatewayUploadBuffer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStoragegatewayUploadBuffers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagegatewayuploadbuffersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StoragegatewayUploadBufferList{})
	return err
}

// Patch applies the patch and returns the patched storagegatewayUploadBuffer.
func (c *FakeStoragegatewayUploadBuffers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StoragegatewayUploadBuffer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagegatewayuploadbuffersResource, c.ns, name, pt, data, subresources...), &v1alpha1.StoragegatewayUploadBuffer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayUploadBuffer), err
}
