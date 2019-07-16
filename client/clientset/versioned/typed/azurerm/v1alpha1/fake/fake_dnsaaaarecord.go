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

// FakeDnsAaaaRecords implements DnsAaaaRecordInterface
type FakeDnsAaaaRecords struct {
	Fake *FakeAzurermV1alpha1
}

var dnsaaaarecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "dnsaaaarecords"}

var dnsaaaarecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DnsAaaaRecord"}

// Get takes name of the dnsAaaaRecord, and returns the corresponding dnsAaaaRecord object, and an error if there is any.
func (c *FakeDnsAaaaRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsAaaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dnsaaaarecordsResource, name), &v1alpha1.DnsAaaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsAaaaRecord), err
}

// List takes label and field selectors, and returns the list of DnsAaaaRecords that match those selectors.
func (c *FakeDnsAaaaRecords) List(opts v1.ListOptions) (result *v1alpha1.DnsAaaaRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dnsaaaarecordsResource, dnsaaaarecordsKind, opts), &v1alpha1.DnsAaaaRecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DnsAaaaRecordList{ListMeta: obj.(*v1alpha1.DnsAaaaRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.DnsAaaaRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnsAaaaRecords.
func (c *FakeDnsAaaaRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dnsaaaarecordsResource, opts))
}

// Create takes the representation of a dnsAaaaRecord and creates it.  Returns the server's representation of the dnsAaaaRecord, and an error, if there is any.
func (c *FakeDnsAaaaRecords) Create(dnsAaaaRecord *v1alpha1.DnsAaaaRecord) (result *v1alpha1.DnsAaaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dnsaaaarecordsResource, dnsAaaaRecord), &v1alpha1.DnsAaaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsAaaaRecord), err
}

// Update takes the representation of a dnsAaaaRecord and updates it. Returns the server's representation of the dnsAaaaRecord, and an error, if there is any.
func (c *FakeDnsAaaaRecords) Update(dnsAaaaRecord *v1alpha1.DnsAaaaRecord) (result *v1alpha1.DnsAaaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dnsaaaarecordsResource, dnsAaaaRecord), &v1alpha1.DnsAaaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsAaaaRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnsAaaaRecords) UpdateStatus(dnsAaaaRecord *v1alpha1.DnsAaaaRecord) (*v1alpha1.DnsAaaaRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dnsaaaarecordsResource, "status", dnsAaaaRecord), &v1alpha1.DnsAaaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsAaaaRecord), err
}

// Delete takes name of the dnsAaaaRecord and deletes it. Returns an error if one occurs.
func (c *FakeDnsAaaaRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dnsaaaarecordsResource, name), &v1alpha1.DnsAaaaRecord{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnsAaaaRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dnsaaaarecordsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DnsAaaaRecordList{})
	return err
}

// Patch applies the patch and returns the patched dnsAaaaRecord.
func (c *FakeDnsAaaaRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsAaaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dnsaaaarecordsResource, name, pt, data, subresources...), &v1alpha1.DnsAaaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsAaaaRecord), err
}