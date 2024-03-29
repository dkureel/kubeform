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

// FakeDxPrivateVirtualInterfaces implements DxPrivateVirtualInterfaceInterface
type FakeDxPrivateVirtualInterfaces struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dxprivatevirtualinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxprivatevirtualinterfaces"}

var dxprivatevirtualinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxPrivateVirtualInterface"}

// Get takes name of the dxPrivateVirtualInterface, and returns the corresponding dxPrivateVirtualInterface object, and an error if there is any.
func (c *FakeDxPrivateVirtualInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.DxPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dxprivatevirtualinterfacesResource, c.ns, name), &v1alpha1.DxPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPrivateVirtualInterface), err
}

// List takes label and field selectors, and returns the list of DxPrivateVirtualInterfaces that match those selectors.
func (c *FakeDxPrivateVirtualInterfaces) List(opts v1.ListOptions) (result *v1alpha1.DxPrivateVirtualInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dxprivatevirtualinterfacesResource, dxprivatevirtualinterfacesKind, c.ns, opts), &v1alpha1.DxPrivateVirtualInterfaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxPrivateVirtualInterfaceList{ListMeta: obj.(*v1alpha1.DxPrivateVirtualInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxPrivateVirtualInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxPrivateVirtualInterfaces.
func (c *FakeDxPrivateVirtualInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dxprivatevirtualinterfacesResource, c.ns, opts))

}

// Create takes the representation of a dxPrivateVirtualInterface and creates it.  Returns the server's representation of the dxPrivateVirtualInterface, and an error, if there is any.
func (c *FakeDxPrivateVirtualInterfaces) Create(dxPrivateVirtualInterface *v1alpha1.DxPrivateVirtualInterface) (result *v1alpha1.DxPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dxprivatevirtualinterfacesResource, c.ns, dxPrivateVirtualInterface), &v1alpha1.DxPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPrivateVirtualInterface), err
}

// Update takes the representation of a dxPrivateVirtualInterface and updates it. Returns the server's representation of the dxPrivateVirtualInterface, and an error, if there is any.
func (c *FakeDxPrivateVirtualInterfaces) Update(dxPrivateVirtualInterface *v1alpha1.DxPrivateVirtualInterface) (result *v1alpha1.DxPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dxprivatevirtualinterfacesResource, c.ns, dxPrivateVirtualInterface), &v1alpha1.DxPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPrivateVirtualInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxPrivateVirtualInterfaces) UpdateStatus(dxPrivateVirtualInterface *v1alpha1.DxPrivateVirtualInterface) (*v1alpha1.DxPrivateVirtualInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dxprivatevirtualinterfacesResource, "status", c.ns, dxPrivateVirtualInterface), &v1alpha1.DxPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPrivateVirtualInterface), err
}

// Delete takes name of the dxPrivateVirtualInterface and deletes it. Returns an error if one occurs.
func (c *FakeDxPrivateVirtualInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dxprivatevirtualinterfacesResource, c.ns, name), &v1alpha1.DxPrivateVirtualInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxPrivateVirtualInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dxprivatevirtualinterfacesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxPrivateVirtualInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched dxPrivateVirtualInterface.
func (c *FakeDxPrivateVirtualInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dxprivatevirtualinterfacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DxPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPrivateVirtualInterface), err
}
