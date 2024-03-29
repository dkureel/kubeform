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

// FakeApiManagementGroupUsers implements ApiManagementGroupUserInterface
type FakeApiManagementGroupUsers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementgroupusersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementgroupusers"}

var apimanagementgroupusersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementGroupUser"}

// Get takes name of the apiManagementGroupUser, and returns the corresponding apiManagementGroupUser object, and an error if there is any.
func (c *FakeApiManagementGroupUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementgroupusersResource, c.ns, name), &v1alpha1.ApiManagementGroupUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroupUser), err
}

// List takes label and field selectors, and returns the list of ApiManagementGroupUsers that match those selectors.
func (c *FakeApiManagementGroupUsers) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementGroupUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementgroupusersResource, apimanagementgroupusersKind, c.ns, opts), &v1alpha1.ApiManagementGroupUserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementGroupUserList{ListMeta: obj.(*v1alpha1.ApiManagementGroupUserList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementGroupUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementGroupUsers.
func (c *FakeApiManagementGroupUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementgroupusersResource, c.ns, opts))

}

// Create takes the representation of a apiManagementGroupUser and creates it.  Returns the server's representation of the apiManagementGroupUser, and an error, if there is any.
func (c *FakeApiManagementGroupUsers) Create(apiManagementGroupUser *v1alpha1.ApiManagementGroupUser) (result *v1alpha1.ApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementgroupusersResource, c.ns, apiManagementGroupUser), &v1alpha1.ApiManagementGroupUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroupUser), err
}

// Update takes the representation of a apiManagementGroupUser and updates it. Returns the server's representation of the apiManagementGroupUser, and an error, if there is any.
func (c *FakeApiManagementGroupUsers) Update(apiManagementGroupUser *v1alpha1.ApiManagementGroupUser) (result *v1alpha1.ApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementgroupusersResource, c.ns, apiManagementGroupUser), &v1alpha1.ApiManagementGroupUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroupUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementGroupUsers) UpdateStatus(apiManagementGroupUser *v1alpha1.ApiManagementGroupUser) (*v1alpha1.ApiManagementGroupUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementgroupusersResource, "status", c.ns, apiManagementGroupUser), &v1alpha1.ApiManagementGroupUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroupUser), err
}

// Delete takes name of the apiManagementGroupUser and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementGroupUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementgroupusersResource, c.ns, name), &v1alpha1.ApiManagementGroupUser{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementGroupUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementgroupusersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementGroupUserList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementGroupUser.
func (c *FakeApiManagementGroupUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementgroupusersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementGroupUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroupUser), err
}
