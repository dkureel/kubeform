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

// FakeLogAnalyticsWorkspaces implements LogAnalyticsWorkspaceInterface
type FakeLogAnalyticsWorkspaces struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var loganalyticsworkspacesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "loganalyticsworkspaces"}

var loganalyticsworkspacesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogAnalyticsWorkspace"}

// Get takes name of the logAnalyticsWorkspace, and returns the corresponding logAnalyticsWorkspace object, and an error if there is any.
func (c *FakeLogAnalyticsWorkspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.LogAnalyticsWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loganalyticsworkspacesResource, c.ns, name), &v1alpha1.LogAnalyticsWorkspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspace), err
}

// List takes label and field selectors, and returns the list of LogAnalyticsWorkspaces that match those selectors.
func (c *FakeLogAnalyticsWorkspaces) List(opts v1.ListOptions) (result *v1alpha1.LogAnalyticsWorkspaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loganalyticsworkspacesResource, loganalyticsworkspacesKind, c.ns, opts), &v1alpha1.LogAnalyticsWorkspaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogAnalyticsWorkspaceList{ListMeta: obj.(*v1alpha1.LogAnalyticsWorkspaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogAnalyticsWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logAnalyticsWorkspaces.
func (c *FakeLogAnalyticsWorkspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loganalyticsworkspacesResource, c.ns, opts))

}

// Create takes the representation of a logAnalyticsWorkspace and creates it.  Returns the server's representation of the logAnalyticsWorkspace, and an error, if there is any.
func (c *FakeLogAnalyticsWorkspaces) Create(logAnalyticsWorkspace *v1alpha1.LogAnalyticsWorkspace) (result *v1alpha1.LogAnalyticsWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loganalyticsworkspacesResource, c.ns, logAnalyticsWorkspace), &v1alpha1.LogAnalyticsWorkspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspace), err
}

// Update takes the representation of a logAnalyticsWorkspace and updates it. Returns the server's representation of the logAnalyticsWorkspace, and an error, if there is any.
func (c *FakeLogAnalyticsWorkspaces) Update(logAnalyticsWorkspace *v1alpha1.LogAnalyticsWorkspace) (result *v1alpha1.LogAnalyticsWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loganalyticsworkspacesResource, c.ns, logAnalyticsWorkspace), &v1alpha1.LogAnalyticsWorkspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogAnalyticsWorkspaces) UpdateStatus(logAnalyticsWorkspace *v1alpha1.LogAnalyticsWorkspace) (*v1alpha1.LogAnalyticsWorkspace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loganalyticsworkspacesResource, "status", c.ns, logAnalyticsWorkspace), &v1alpha1.LogAnalyticsWorkspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspace), err
}

// Delete takes name of the logAnalyticsWorkspace and deletes it. Returns an error if one occurs.
func (c *FakeLogAnalyticsWorkspaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loganalyticsworkspacesResource, c.ns, name), &v1alpha1.LogAnalyticsWorkspace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogAnalyticsWorkspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loganalyticsworkspacesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogAnalyticsWorkspaceList{})
	return err
}

// Patch applies the patch and returns the patched logAnalyticsWorkspace.
func (c *FakeLogAnalyticsWorkspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogAnalyticsWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loganalyticsworkspacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogAnalyticsWorkspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsWorkspace), err
}
