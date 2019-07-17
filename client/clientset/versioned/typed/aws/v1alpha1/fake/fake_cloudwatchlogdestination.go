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

// FakeCloudwatchLogDestinations implements CloudwatchLogDestinationInterface
type FakeCloudwatchLogDestinations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudwatchlogdestinationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudwatchlogdestinations"}

var cloudwatchlogdestinationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudwatchLogDestination"}

// Get takes name of the cloudwatchLogDestination, and returns the corresponding cloudwatchLogDestination object, and an error if there is any.
func (c *FakeCloudwatchLogDestinations) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudwatchlogdestinationsResource, c.ns, name), &v1alpha1.CloudwatchLogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestination), err
}

// List takes label and field selectors, and returns the list of CloudwatchLogDestinations that match those selectors.
func (c *FakeCloudwatchLogDestinations) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogDestinationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudwatchlogdestinationsResource, cloudwatchlogdestinationsKind, c.ns, opts), &v1alpha1.CloudwatchLogDestinationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudwatchLogDestinationList{ListMeta: obj.(*v1alpha1.CloudwatchLogDestinationList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudwatchLogDestinationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogDestinations.
func (c *FakeCloudwatchLogDestinations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudwatchlogdestinationsResource, c.ns, opts))

}

// Create takes the representation of a cloudwatchLogDestination and creates it.  Returns the server's representation of the cloudwatchLogDestination, and an error, if there is any.
func (c *FakeCloudwatchLogDestinations) Create(cloudwatchLogDestination *v1alpha1.CloudwatchLogDestination) (result *v1alpha1.CloudwatchLogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudwatchlogdestinationsResource, c.ns, cloudwatchLogDestination), &v1alpha1.CloudwatchLogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestination), err
}

// Update takes the representation of a cloudwatchLogDestination and updates it. Returns the server's representation of the cloudwatchLogDestination, and an error, if there is any.
func (c *FakeCloudwatchLogDestinations) Update(cloudwatchLogDestination *v1alpha1.CloudwatchLogDestination) (result *v1alpha1.CloudwatchLogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudwatchlogdestinationsResource, c.ns, cloudwatchLogDestination), &v1alpha1.CloudwatchLogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestination), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudwatchLogDestinations) UpdateStatus(cloudwatchLogDestination *v1alpha1.CloudwatchLogDestination) (*v1alpha1.CloudwatchLogDestination, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudwatchlogdestinationsResource, "status", c.ns, cloudwatchLogDestination), &v1alpha1.CloudwatchLogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestination), err
}

// Delete takes name of the cloudwatchLogDestination and deletes it. Returns an error if one occurs.
func (c *FakeCloudwatchLogDestinations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudwatchlogdestinationsResource, c.ns, name), &v1alpha1.CloudwatchLogDestination{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudwatchLogDestinations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudwatchlogdestinationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudwatchLogDestinationList{})
	return err
}

// Patch applies the patch and returns the patched cloudwatchLogDestination.
func (c *FakeCloudwatchLogDestinations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudwatchlogdestinationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudwatchLogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestination), err
}
