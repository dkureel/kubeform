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

// FakeSignalrServices implements SignalrServiceInterface
type FakeSignalrServices struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var signalrservicesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "signalrservices"}

var signalrservicesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SignalrService"}

// Get takes name of the signalrService, and returns the corresponding signalrService object, and an error if there is any.
func (c *FakeSignalrServices) Get(name string, options v1.GetOptions) (result *v1alpha1.SignalrService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(signalrservicesResource, c.ns, name), &v1alpha1.SignalrService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SignalrService), err
}

// List takes label and field selectors, and returns the list of SignalrServices that match those selectors.
func (c *FakeSignalrServices) List(opts v1.ListOptions) (result *v1alpha1.SignalrServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(signalrservicesResource, signalrservicesKind, c.ns, opts), &v1alpha1.SignalrServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SignalrServiceList{ListMeta: obj.(*v1alpha1.SignalrServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SignalrServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested signalrServices.
func (c *FakeSignalrServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(signalrservicesResource, c.ns, opts))

}

// Create takes the representation of a signalrService and creates it.  Returns the server's representation of the signalrService, and an error, if there is any.
func (c *FakeSignalrServices) Create(signalrService *v1alpha1.SignalrService) (result *v1alpha1.SignalrService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(signalrservicesResource, c.ns, signalrService), &v1alpha1.SignalrService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SignalrService), err
}

// Update takes the representation of a signalrService and updates it. Returns the server's representation of the signalrService, and an error, if there is any.
func (c *FakeSignalrServices) Update(signalrService *v1alpha1.SignalrService) (result *v1alpha1.SignalrService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(signalrservicesResource, c.ns, signalrService), &v1alpha1.SignalrService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SignalrService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSignalrServices) UpdateStatus(signalrService *v1alpha1.SignalrService) (*v1alpha1.SignalrService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(signalrservicesResource, "status", c.ns, signalrService), &v1alpha1.SignalrService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SignalrService), err
}

// Delete takes name of the signalrService and deletes it. Returns an error if one occurs.
func (c *FakeSignalrServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(signalrservicesResource, c.ns, name), &v1alpha1.SignalrService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSignalrServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(signalrservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SignalrServiceList{})
	return err
}

// Patch applies the patch and returns the patched signalrService.
func (c *FakeSignalrServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SignalrService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(signalrservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SignalrService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SignalrService), err
}
