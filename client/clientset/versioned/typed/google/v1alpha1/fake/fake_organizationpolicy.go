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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeOrganizationPolicies implements OrganizationPolicyInterface
type FakeOrganizationPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var organizationpoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "organizationpolicies"}

var organizationpoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "OrganizationPolicy"}

// Get takes name of the organizationPolicy, and returns the corresponding organizationPolicy object, and an error if there is any.
func (c *FakeOrganizationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.OrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(organizationpoliciesResource, c.ns, name), &v1alpha1.OrganizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationPolicy), err
}

// List takes label and field selectors, and returns the list of OrganizationPolicies that match those selectors.
func (c *FakeOrganizationPolicies) List(opts v1.ListOptions) (result *v1alpha1.OrganizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(organizationpoliciesResource, organizationpoliciesKind, c.ns, opts), &v1alpha1.OrganizationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationPolicyList{ListMeta: obj.(*v1alpha1.OrganizationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationPolicies.
func (c *FakeOrganizationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(organizationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a organizationPolicy and creates it.  Returns the server's representation of the organizationPolicy, and an error, if there is any.
func (c *FakeOrganizationPolicies) Create(organizationPolicy *v1alpha1.OrganizationPolicy) (result *v1alpha1.OrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(organizationpoliciesResource, c.ns, organizationPolicy), &v1alpha1.OrganizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationPolicy), err
}

// Update takes the representation of a organizationPolicy and updates it. Returns the server's representation of the organizationPolicy, and an error, if there is any.
func (c *FakeOrganizationPolicies) Update(organizationPolicy *v1alpha1.OrganizationPolicy) (result *v1alpha1.OrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(organizationpoliciesResource, c.ns, organizationPolicy), &v1alpha1.OrganizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationPolicies) UpdateStatus(organizationPolicy *v1alpha1.OrganizationPolicy) (*v1alpha1.OrganizationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(organizationpoliciesResource, "status", c.ns, organizationPolicy), &v1alpha1.OrganizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationPolicy), err
}

// Delete takes name of the organizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(organizationpoliciesResource, c.ns, name), &v1alpha1.OrganizationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(organizationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched organizationPolicy.
func (c *FakeOrganizationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(organizationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrganizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationPolicy), err
}
