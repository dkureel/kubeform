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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeResourcegroupsGroups implements ResourcegroupsGroupInterface
type FakeResourcegroupsGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var resourcegroupsgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "resourcegroupsgroups"}

var resourcegroupsgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ResourcegroupsGroup"}

// Get takes name of the resourcegroupsGroup, and returns the corresponding resourcegroupsGroup object, and an error if there is any.
func (c *FakeResourcegroupsGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourcegroupsGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcegroupsgroupsResource, c.ns, name), &v1alpha1.ResourcegroupsGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcegroupsGroup), err
}

// List takes label and field selectors, and returns the list of ResourcegroupsGroups that match those selectors.
func (c *FakeResourcegroupsGroups) List(opts v1.ListOptions) (result *v1alpha1.ResourcegroupsGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcegroupsgroupsResource, resourcegroupsgroupsKind, c.ns, opts), &v1alpha1.ResourcegroupsGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourcegroupsGroupList{ListMeta: obj.(*v1alpha1.ResourcegroupsGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourcegroupsGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourcegroupsGroups.
func (c *FakeResourcegroupsGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcegroupsgroupsResource, c.ns, opts))

}

// Create takes the representation of a resourcegroupsGroup and creates it.  Returns the server's representation of the resourcegroupsGroup, and an error, if there is any.
func (c *FakeResourcegroupsGroups) Create(resourcegroupsGroup *v1alpha1.ResourcegroupsGroup) (result *v1alpha1.ResourcegroupsGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcegroupsgroupsResource, c.ns, resourcegroupsGroup), &v1alpha1.ResourcegroupsGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcegroupsGroup), err
}

// Update takes the representation of a resourcegroupsGroup and updates it. Returns the server's representation of the resourcegroupsGroup, and an error, if there is any.
func (c *FakeResourcegroupsGroups) Update(resourcegroupsGroup *v1alpha1.ResourcegroupsGroup) (result *v1alpha1.ResourcegroupsGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcegroupsgroupsResource, c.ns, resourcegroupsGroup), &v1alpha1.ResourcegroupsGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcegroupsGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourcegroupsGroups) UpdateStatus(resourcegroupsGroup *v1alpha1.ResourcegroupsGroup) (*v1alpha1.ResourcegroupsGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcegroupsgroupsResource, "status", c.ns, resourcegroupsGroup), &v1alpha1.ResourcegroupsGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcegroupsGroup), err
}

// Delete takes name of the resourcegroupsGroup and deletes it. Returns an error if one occurs.
func (c *FakeResourcegroupsGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourcegroupsgroupsResource, c.ns, name), &v1alpha1.ResourcegroupsGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourcegroupsGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcegroupsgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourcegroupsGroupList{})
	return err
}

// Patch applies the patch and returns the patched resourcegroupsGroup.
func (c *FakeResourcegroupsGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourcegroupsGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcegroupsgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourcegroupsGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourcegroupsGroup), err
}
