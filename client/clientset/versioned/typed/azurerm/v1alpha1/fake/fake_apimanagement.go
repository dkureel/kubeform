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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiManagements implements ApiManagementInterface
type FakeApiManagements struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagements"}

var apimanagementsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagement"}

// Get takes name of the apiManagement, and returns the corresponding apiManagement object, and an error if there is any.
func (c *FakeApiManagements) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementsResource, c.ns, name), &v1alpha1.ApiManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagement), err
}

// List takes label and field selectors, and returns the list of ApiManagements that match those selectors.
func (c *FakeApiManagements) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementsResource, apimanagementsKind, c.ns, opts), &v1alpha1.ApiManagementList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementList{ListMeta: obj.(*v1alpha1.ApiManagementList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagements.
func (c *FakeApiManagements) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementsResource, c.ns, opts))

}

// Create takes the representation of a apiManagement and creates it.  Returns the server's representation of the apiManagement, and an error, if there is any.
func (c *FakeApiManagements) Create(apiManagement *v1alpha1.ApiManagement) (result *v1alpha1.ApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementsResource, c.ns, apiManagement), &v1alpha1.ApiManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagement), err
}

// Update takes the representation of a apiManagement and updates it. Returns the server's representation of the apiManagement, and an error, if there is any.
func (c *FakeApiManagements) Update(apiManagement *v1alpha1.ApiManagement) (result *v1alpha1.ApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementsResource, c.ns, apiManagement), &v1alpha1.ApiManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagement), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagements) UpdateStatus(apiManagement *v1alpha1.ApiManagement) (*v1alpha1.ApiManagement, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementsResource, "status", c.ns, apiManagement), &v1alpha1.ApiManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagement), err
}

// Delete takes name of the apiManagement and deletes it. Returns an error if one occurs.
func (c *FakeApiManagements) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementsResource, c.ns, name), &v1alpha1.ApiManagement{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagements) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementList{})
	return err
}

// Patch applies the patch and returns the patched apiManagement.
func (c *FakeApiManagements) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagement), err
}
