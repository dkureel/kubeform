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

// FakeElbs implements ElbInterface
type FakeElbs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var elbsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "elbs"}

var elbsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Elb"}

// Get takes name of the elb, and returns the corresponding elb object, and an error if there is any.
func (c *FakeElbs) Get(name string, options v1.GetOptions) (result *v1alpha1.Elb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elbsResource, c.ns, name), &v1alpha1.Elb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elb), err
}

// List takes label and field selectors, and returns the list of Elbs that match those selectors.
func (c *FakeElbs) List(opts v1.ListOptions) (result *v1alpha1.ElbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elbsResource, elbsKind, c.ns, opts), &v1alpha1.ElbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElbList{ListMeta: obj.(*v1alpha1.ElbList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elbs.
func (c *FakeElbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elbsResource, c.ns, opts))

}

// Create takes the representation of a elb and creates it.  Returns the server's representation of the elb, and an error, if there is any.
func (c *FakeElbs) Create(elb *v1alpha1.Elb) (result *v1alpha1.Elb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elbsResource, c.ns, elb), &v1alpha1.Elb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elb), err
}

// Update takes the representation of a elb and updates it. Returns the server's representation of the elb, and an error, if there is any.
func (c *FakeElbs) Update(elb *v1alpha1.Elb) (result *v1alpha1.Elb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elbsResource, c.ns, elb), &v1alpha1.Elb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElbs) UpdateStatus(elb *v1alpha1.Elb) (*v1alpha1.Elb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elbsResource, "status", c.ns, elb), &v1alpha1.Elb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elb), err
}

// Delete takes name of the elb and deletes it. Returns an error if one occurs.
func (c *FakeElbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(elbsResource, c.ns, name), &v1alpha1.Elb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elbsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElbList{})
	return err
}

// Patch applies the patch and returns the patched elb.
func (c *FakeElbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Elb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elbsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Elb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elb), err
}
