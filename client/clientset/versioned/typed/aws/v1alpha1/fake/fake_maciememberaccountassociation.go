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

// FakeMacieMemberAccountAssociations implements MacieMemberAccountAssociationInterface
type FakeMacieMemberAccountAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var maciememberaccountassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "maciememberaccountassociations"}

var maciememberaccountassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "MacieMemberAccountAssociation"}

// Get takes name of the macieMemberAccountAssociation, and returns the corresponding macieMemberAccountAssociation object, and an error if there is any.
func (c *FakeMacieMemberAccountAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.MacieMemberAccountAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(maciememberaccountassociationsResource, c.ns, name), &v1alpha1.MacieMemberAccountAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieMemberAccountAssociation), err
}

// List takes label and field selectors, and returns the list of MacieMemberAccountAssociations that match those selectors.
func (c *FakeMacieMemberAccountAssociations) List(opts v1.ListOptions) (result *v1alpha1.MacieMemberAccountAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(maciememberaccountassociationsResource, maciememberaccountassociationsKind, c.ns, opts), &v1alpha1.MacieMemberAccountAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MacieMemberAccountAssociationList{ListMeta: obj.(*v1alpha1.MacieMemberAccountAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.MacieMemberAccountAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested macieMemberAccountAssociations.
func (c *FakeMacieMemberAccountAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(maciememberaccountassociationsResource, c.ns, opts))

}

// Create takes the representation of a macieMemberAccountAssociation and creates it.  Returns the server's representation of the macieMemberAccountAssociation, and an error, if there is any.
func (c *FakeMacieMemberAccountAssociations) Create(macieMemberAccountAssociation *v1alpha1.MacieMemberAccountAssociation) (result *v1alpha1.MacieMemberAccountAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(maciememberaccountassociationsResource, c.ns, macieMemberAccountAssociation), &v1alpha1.MacieMemberAccountAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieMemberAccountAssociation), err
}

// Update takes the representation of a macieMemberAccountAssociation and updates it. Returns the server's representation of the macieMemberAccountAssociation, and an error, if there is any.
func (c *FakeMacieMemberAccountAssociations) Update(macieMemberAccountAssociation *v1alpha1.MacieMemberAccountAssociation) (result *v1alpha1.MacieMemberAccountAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(maciememberaccountassociationsResource, c.ns, macieMemberAccountAssociation), &v1alpha1.MacieMemberAccountAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieMemberAccountAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMacieMemberAccountAssociations) UpdateStatus(macieMemberAccountAssociation *v1alpha1.MacieMemberAccountAssociation) (*v1alpha1.MacieMemberAccountAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(maciememberaccountassociationsResource, "status", c.ns, macieMemberAccountAssociation), &v1alpha1.MacieMemberAccountAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieMemberAccountAssociation), err
}

// Delete takes name of the macieMemberAccountAssociation and deletes it. Returns an error if one occurs.
func (c *FakeMacieMemberAccountAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(maciememberaccountassociationsResource, c.ns, name), &v1alpha1.MacieMemberAccountAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMacieMemberAccountAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(maciememberaccountassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MacieMemberAccountAssociationList{})
	return err
}

// Patch applies the patch and returns the patched macieMemberAccountAssociation.
func (c *FakeMacieMemberAccountAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MacieMemberAccountAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(maciememberaccountassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MacieMemberAccountAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MacieMemberAccountAssociation), err
}
