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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageBuckets implements StorageBucketInterface
type FakeStorageBuckets struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var storagebucketsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagebuckets"}

var storagebucketsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageBucket"}

// Get takes name of the storageBucket, and returns the corresponding storageBucket object, and an error if there is any.
func (c *FakeStorageBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagebucketsResource, c.ns, name), &v1alpha1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucket), err
}

// List takes label and field selectors, and returns the list of StorageBuckets that match those selectors.
func (c *FakeStorageBuckets) List(opts v1.ListOptions) (result *v1alpha1.StorageBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagebucketsResource, storagebucketsKind, c.ns, opts), &v1alpha1.StorageBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageBucketList{ListMeta: obj.(*v1alpha1.StorageBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageBuckets.
func (c *FakeStorageBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagebucketsResource, c.ns, opts))

}

// Create takes the representation of a storageBucket and creates it.  Returns the server's representation of the storageBucket, and an error, if there is any.
func (c *FakeStorageBuckets) Create(storageBucket *v1alpha1.StorageBucket) (result *v1alpha1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagebucketsResource, c.ns, storageBucket), &v1alpha1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucket), err
}

// Update takes the representation of a storageBucket and updates it. Returns the server's representation of the storageBucket, and an error, if there is any.
func (c *FakeStorageBuckets) Update(storageBucket *v1alpha1.StorageBucket) (result *v1alpha1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagebucketsResource, c.ns, storageBucket), &v1alpha1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageBuckets) UpdateStatus(storageBucket *v1alpha1.StorageBucket) (*v1alpha1.StorageBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagebucketsResource, "status", c.ns, storageBucket), &v1alpha1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucket), err
}

// Delete takes name of the storageBucket and deletes it. Returns an error if one occurs.
func (c *FakeStorageBuckets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagebucketsResource, c.ns, name), &v1alpha1.StorageBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagebucketsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageBucketList{})
	return err
}

// Patch applies the patch and returns the patched storageBucket.
func (c *FakeStorageBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagebucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucket), err
}
