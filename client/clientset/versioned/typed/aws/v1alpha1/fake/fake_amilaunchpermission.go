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

// FakeAmiLaunchPermissions implements AmiLaunchPermissionInterface
type FakeAmiLaunchPermissions struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var amilaunchpermissionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "amilaunchpermissions"}

var amilaunchpermissionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AmiLaunchPermission"}

// Get takes name of the amiLaunchPermission, and returns the corresponding amiLaunchPermission object, and an error if there is any.
func (c *FakeAmiLaunchPermissions) Get(name string, options v1.GetOptions) (result *v1alpha1.AmiLaunchPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(amilaunchpermissionsResource, c.ns, name), &v1alpha1.AmiLaunchPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiLaunchPermission), err
}

// List takes label and field selectors, and returns the list of AmiLaunchPermissions that match those selectors.
func (c *FakeAmiLaunchPermissions) List(opts v1.ListOptions) (result *v1alpha1.AmiLaunchPermissionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(amilaunchpermissionsResource, amilaunchpermissionsKind, c.ns, opts), &v1alpha1.AmiLaunchPermissionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AmiLaunchPermissionList{ListMeta: obj.(*v1alpha1.AmiLaunchPermissionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AmiLaunchPermissionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested amiLaunchPermissions.
func (c *FakeAmiLaunchPermissions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(amilaunchpermissionsResource, c.ns, opts))

}

// Create takes the representation of a amiLaunchPermission and creates it.  Returns the server's representation of the amiLaunchPermission, and an error, if there is any.
func (c *FakeAmiLaunchPermissions) Create(amiLaunchPermission *v1alpha1.AmiLaunchPermission) (result *v1alpha1.AmiLaunchPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(amilaunchpermissionsResource, c.ns, amiLaunchPermission), &v1alpha1.AmiLaunchPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiLaunchPermission), err
}

// Update takes the representation of a amiLaunchPermission and updates it. Returns the server's representation of the amiLaunchPermission, and an error, if there is any.
func (c *FakeAmiLaunchPermissions) Update(amiLaunchPermission *v1alpha1.AmiLaunchPermission) (result *v1alpha1.AmiLaunchPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(amilaunchpermissionsResource, c.ns, amiLaunchPermission), &v1alpha1.AmiLaunchPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiLaunchPermission), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAmiLaunchPermissions) UpdateStatus(amiLaunchPermission *v1alpha1.AmiLaunchPermission) (*v1alpha1.AmiLaunchPermission, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(amilaunchpermissionsResource, "status", c.ns, amiLaunchPermission), &v1alpha1.AmiLaunchPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiLaunchPermission), err
}

// Delete takes name of the amiLaunchPermission and deletes it. Returns an error if one occurs.
func (c *FakeAmiLaunchPermissions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(amilaunchpermissionsResource, c.ns, name), &v1alpha1.AmiLaunchPermission{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAmiLaunchPermissions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(amilaunchpermissionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AmiLaunchPermissionList{})
	return err
}

// Patch applies the patch and returns the patched amiLaunchPermission.
func (c *FakeAmiLaunchPermissions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AmiLaunchPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(amilaunchpermissionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AmiLaunchPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiLaunchPermission), err
}
