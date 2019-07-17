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

// FakeSesIdentityPolicies implements SesIdentityPolicyInterface
type FakeSesIdentityPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var sesidentitypoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "sesidentitypolicies"}

var sesidentitypoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SesIdentityPolicy"}

// Get takes name of the sesIdentityPolicy, and returns the corresponding sesIdentityPolicy object, and an error if there is any.
func (c *FakeSesIdentityPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.SesIdentityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sesidentitypoliciesResource, c.ns, name), &v1alpha1.SesIdentityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesIdentityPolicy), err
}

// List takes label and field selectors, and returns the list of SesIdentityPolicies that match those selectors.
func (c *FakeSesIdentityPolicies) List(opts v1.ListOptions) (result *v1alpha1.SesIdentityPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sesidentitypoliciesResource, sesidentitypoliciesKind, c.ns, opts), &v1alpha1.SesIdentityPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SesIdentityPolicyList{ListMeta: obj.(*v1alpha1.SesIdentityPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.SesIdentityPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sesIdentityPolicies.
func (c *FakeSesIdentityPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sesidentitypoliciesResource, c.ns, opts))

}

// Create takes the representation of a sesIdentityPolicy and creates it.  Returns the server's representation of the sesIdentityPolicy, and an error, if there is any.
func (c *FakeSesIdentityPolicies) Create(sesIdentityPolicy *v1alpha1.SesIdentityPolicy) (result *v1alpha1.SesIdentityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sesidentitypoliciesResource, c.ns, sesIdentityPolicy), &v1alpha1.SesIdentityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesIdentityPolicy), err
}

// Update takes the representation of a sesIdentityPolicy and updates it. Returns the server's representation of the sesIdentityPolicy, and an error, if there is any.
func (c *FakeSesIdentityPolicies) Update(sesIdentityPolicy *v1alpha1.SesIdentityPolicy) (result *v1alpha1.SesIdentityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sesidentitypoliciesResource, c.ns, sesIdentityPolicy), &v1alpha1.SesIdentityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesIdentityPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSesIdentityPolicies) UpdateStatus(sesIdentityPolicy *v1alpha1.SesIdentityPolicy) (*v1alpha1.SesIdentityPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sesidentitypoliciesResource, "status", c.ns, sesIdentityPolicy), &v1alpha1.SesIdentityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesIdentityPolicy), err
}

// Delete takes name of the sesIdentityPolicy and deletes it. Returns an error if one occurs.
func (c *FakeSesIdentityPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sesidentitypoliciesResource, c.ns, name), &v1alpha1.SesIdentityPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSesIdentityPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sesidentitypoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SesIdentityPolicyList{})
	return err
}

// Patch applies the patch and returns the patched sesIdentityPolicy.
func (c *FakeSesIdentityPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesIdentityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sesidentitypoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SesIdentityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesIdentityPolicy), err
}
