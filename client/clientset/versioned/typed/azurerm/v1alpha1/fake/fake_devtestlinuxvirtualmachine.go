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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeDevTestLinuxVirtualMachines implements DevTestLinuxVirtualMachineInterface
type FakeDevTestLinuxVirtualMachines struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var devtestlinuxvirtualmachinesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "devtestlinuxvirtualmachines"}

var devtestlinuxvirtualmachinesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DevTestLinuxVirtualMachine"}

// Get takes name of the devTestLinuxVirtualMachine, and returns the corresponding devTestLinuxVirtualMachine object, and an error if there is any.
func (c *FakeDevTestLinuxVirtualMachines) Get(name string, options v1.GetOptions) (result *v1alpha1.DevTestLinuxVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devtestlinuxvirtualmachinesResource, c.ns, name), &v1alpha1.DevTestLinuxVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLinuxVirtualMachine), err
}

// List takes label and field selectors, and returns the list of DevTestLinuxVirtualMachines that match those selectors.
func (c *FakeDevTestLinuxVirtualMachines) List(opts v1.ListOptions) (result *v1alpha1.DevTestLinuxVirtualMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devtestlinuxvirtualmachinesResource, devtestlinuxvirtualmachinesKind, c.ns, opts), &v1alpha1.DevTestLinuxVirtualMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevTestLinuxVirtualMachineList{ListMeta: obj.(*v1alpha1.DevTestLinuxVirtualMachineList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevTestLinuxVirtualMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devTestLinuxVirtualMachines.
func (c *FakeDevTestLinuxVirtualMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devtestlinuxvirtualmachinesResource, c.ns, opts))

}

// Create takes the representation of a devTestLinuxVirtualMachine and creates it.  Returns the server's representation of the devTestLinuxVirtualMachine, and an error, if there is any.
func (c *FakeDevTestLinuxVirtualMachines) Create(devTestLinuxVirtualMachine *v1alpha1.DevTestLinuxVirtualMachine) (result *v1alpha1.DevTestLinuxVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devtestlinuxvirtualmachinesResource, c.ns, devTestLinuxVirtualMachine), &v1alpha1.DevTestLinuxVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLinuxVirtualMachine), err
}

// Update takes the representation of a devTestLinuxVirtualMachine and updates it. Returns the server's representation of the devTestLinuxVirtualMachine, and an error, if there is any.
func (c *FakeDevTestLinuxVirtualMachines) Update(devTestLinuxVirtualMachine *v1alpha1.DevTestLinuxVirtualMachine) (result *v1alpha1.DevTestLinuxVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devtestlinuxvirtualmachinesResource, c.ns, devTestLinuxVirtualMachine), &v1alpha1.DevTestLinuxVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLinuxVirtualMachine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevTestLinuxVirtualMachines) UpdateStatus(devTestLinuxVirtualMachine *v1alpha1.DevTestLinuxVirtualMachine) (*v1alpha1.DevTestLinuxVirtualMachine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devtestlinuxvirtualmachinesResource, "status", c.ns, devTestLinuxVirtualMachine), &v1alpha1.DevTestLinuxVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLinuxVirtualMachine), err
}

// Delete takes name of the devTestLinuxVirtualMachine and deletes it. Returns an error if one occurs.
func (c *FakeDevTestLinuxVirtualMachines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devtestlinuxvirtualmachinesResource, c.ns, name), &v1alpha1.DevTestLinuxVirtualMachine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevTestLinuxVirtualMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devtestlinuxvirtualmachinesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevTestLinuxVirtualMachineList{})
	return err
}

// Patch applies the patch and returns the patched devTestLinuxVirtualMachine.
func (c *FakeDevTestLinuxVirtualMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevTestLinuxVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devtestlinuxvirtualmachinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevTestLinuxVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLinuxVirtualMachine), err
}
