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

// FakeRoute53DelegationSets implements Route53DelegationSetInterface
type FakeRoute53DelegationSets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var route53delegationsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "route53delegationsets"}

var route53delegationsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Route53DelegationSet"}

// Get takes name of the route53DelegationSet, and returns the corresponding route53DelegationSet object, and an error if there is any.
func (c *FakeRoute53DelegationSets) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53DelegationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(route53delegationsetsResource, c.ns, name), &v1alpha1.Route53DelegationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53DelegationSet), err
}

// List takes label and field selectors, and returns the list of Route53DelegationSets that match those selectors.
func (c *FakeRoute53DelegationSets) List(opts v1.ListOptions) (result *v1alpha1.Route53DelegationSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(route53delegationsetsResource, route53delegationsetsKind, c.ns, opts), &v1alpha1.Route53DelegationSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Route53DelegationSetList{ListMeta: obj.(*v1alpha1.Route53DelegationSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.Route53DelegationSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested route53DelegationSets.
func (c *FakeRoute53DelegationSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(route53delegationsetsResource, c.ns, opts))

}

// Create takes the representation of a route53DelegationSet and creates it.  Returns the server's representation of the route53DelegationSet, and an error, if there is any.
func (c *FakeRoute53DelegationSets) Create(route53DelegationSet *v1alpha1.Route53DelegationSet) (result *v1alpha1.Route53DelegationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(route53delegationsetsResource, c.ns, route53DelegationSet), &v1alpha1.Route53DelegationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53DelegationSet), err
}

// Update takes the representation of a route53DelegationSet and updates it. Returns the server's representation of the route53DelegationSet, and an error, if there is any.
func (c *FakeRoute53DelegationSets) Update(route53DelegationSet *v1alpha1.Route53DelegationSet) (result *v1alpha1.Route53DelegationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(route53delegationsetsResource, c.ns, route53DelegationSet), &v1alpha1.Route53DelegationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53DelegationSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoute53DelegationSets) UpdateStatus(route53DelegationSet *v1alpha1.Route53DelegationSet) (*v1alpha1.Route53DelegationSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(route53delegationsetsResource, "status", c.ns, route53DelegationSet), &v1alpha1.Route53DelegationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53DelegationSet), err
}

// Delete takes name of the route53DelegationSet and deletes it. Returns an error if one occurs.
func (c *FakeRoute53DelegationSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(route53delegationsetsResource, c.ns, name), &v1alpha1.Route53DelegationSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoute53DelegationSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(route53delegationsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.Route53DelegationSetList{})
	return err
}

// Patch applies the patch and returns the patched route53DelegationSet.
func (c *FakeRoute53DelegationSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53DelegationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(route53delegationsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Route53DelegationSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53DelegationSet), err
}
