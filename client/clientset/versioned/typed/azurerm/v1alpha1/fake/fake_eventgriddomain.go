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

// FakeEventgridDomains implements EventgridDomainInterface
type FakeEventgridDomains struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var eventgriddomainsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "eventgriddomains"}

var eventgriddomainsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "EventgridDomain"}

// Get takes name of the eventgridDomain, and returns the corresponding eventgridDomain object, and an error if there is any.
func (c *FakeEventgridDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.EventgridDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventgriddomainsResource, c.ns, name), &v1alpha1.EventgridDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridDomain), err
}

// List takes label and field selectors, and returns the list of EventgridDomains that match those selectors.
func (c *FakeEventgridDomains) List(opts v1.ListOptions) (result *v1alpha1.EventgridDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventgriddomainsResource, eventgriddomainsKind, c.ns, opts), &v1alpha1.EventgridDomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventgridDomainList{ListMeta: obj.(*v1alpha1.EventgridDomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventgridDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventgridDomains.
func (c *FakeEventgridDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventgriddomainsResource, c.ns, opts))

}

// Create takes the representation of a eventgridDomain and creates it.  Returns the server's representation of the eventgridDomain, and an error, if there is any.
func (c *FakeEventgridDomains) Create(eventgridDomain *v1alpha1.EventgridDomain) (result *v1alpha1.EventgridDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventgriddomainsResource, c.ns, eventgridDomain), &v1alpha1.EventgridDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridDomain), err
}

// Update takes the representation of a eventgridDomain and updates it. Returns the server's representation of the eventgridDomain, and an error, if there is any.
func (c *FakeEventgridDomains) Update(eventgridDomain *v1alpha1.EventgridDomain) (result *v1alpha1.EventgridDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventgriddomainsResource, c.ns, eventgridDomain), &v1alpha1.EventgridDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridDomain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventgridDomains) UpdateStatus(eventgridDomain *v1alpha1.EventgridDomain) (*v1alpha1.EventgridDomain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventgriddomainsResource, "status", c.ns, eventgridDomain), &v1alpha1.EventgridDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridDomain), err
}

// Delete takes name of the eventgridDomain and deletes it. Returns an error if one occurs.
func (c *FakeEventgridDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventgriddomainsResource, c.ns, name), &v1alpha1.EventgridDomain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventgridDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventgriddomainsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventgridDomainList{})
	return err
}

// Patch applies the patch and returns the patched eventgridDomain.
func (c *FakeEventgridDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventgridDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventgriddomainsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventgridDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridDomain), err
}
