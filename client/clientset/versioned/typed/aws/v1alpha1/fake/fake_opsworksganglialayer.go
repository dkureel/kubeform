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

// FakeOpsworksGangliaLayers implements OpsworksGangliaLayerInterface
type FakeOpsworksGangliaLayers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworksganglialayersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksganglialayers"}

var opsworksganglialayersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksGangliaLayer"}

// Get takes name of the opsworksGangliaLayer, and returns the corresponding opsworksGangliaLayer object, and an error if there is any.
func (c *FakeOpsworksGangliaLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksGangliaLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworksganglialayersResource, c.ns, name), &v1alpha1.OpsworksGangliaLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksGangliaLayer), err
}

// List takes label and field selectors, and returns the list of OpsworksGangliaLayers that match those selectors.
func (c *FakeOpsworksGangliaLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksGangliaLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworksganglialayersResource, opsworksganglialayersKind, c.ns, opts), &v1alpha1.OpsworksGangliaLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksGangliaLayerList{ListMeta: obj.(*v1alpha1.OpsworksGangliaLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksGangliaLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksGangliaLayers.
func (c *FakeOpsworksGangliaLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworksganglialayersResource, c.ns, opts))

}

// Create takes the representation of a opsworksGangliaLayer and creates it.  Returns the server's representation of the opsworksGangliaLayer, and an error, if there is any.
func (c *FakeOpsworksGangliaLayers) Create(opsworksGangliaLayer *v1alpha1.OpsworksGangliaLayer) (result *v1alpha1.OpsworksGangliaLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworksganglialayersResource, c.ns, opsworksGangliaLayer), &v1alpha1.OpsworksGangliaLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksGangliaLayer), err
}

// Update takes the representation of a opsworksGangliaLayer and updates it. Returns the server's representation of the opsworksGangliaLayer, and an error, if there is any.
func (c *FakeOpsworksGangliaLayers) Update(opsworksGangliaLayer *v1alpha1.OpsworksGangliaLayer) (result *v1alpha1.OpsworksGangliaLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworksganglialayersResource, c.ns, opsworksGangliaLayer), &v1alpha1.OpsworksGangliaLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksGangliaLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksGangliaLayers) UpdateStatus(opsworksGangliaLayer *v1alpha1.OpsworksGangliaLayer) (*v1alpha1.OpsworksGangliaLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworksganglialayersResource, "status", c.ns, opsworksGangliaLayer), &v1alpha1.OpsworksGangliaLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksGangliaLayer), err
}

// Delete takes name of the opsworksGangliaLayer and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksGangliaLayers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworksganglialayersResource, c.ns, name), &v1alpha1.OpsworksGangliaLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksGangliaLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworksganglialayersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksGangliaLayerList{})
	return err
}

// Patch applies the patch and returns the patched opsworksGangliaLayer.
func (c *FakeOpsworksGangliaLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksGangliaLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworksganglialayersResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksGangliaLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksGangliaLayer), err
}
