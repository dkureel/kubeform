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

// FakeElastictranscoderPresets implements ElastictranscoderPresetInterface
type FakeElastictranscoderPresets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var elastictranscoderpresetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "elastictranscoderpresets"}

var elastictranscoderpresetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ElastictranscoderPreset"}

// Get takes name of the elastictranscoderPreset, and returns the corresponding elastictranscoderPreset object, and an error if there is any.
func (c *FakeElastictranscoderPresets) Get(name string, options v1.GetOptions) (result *v1alpha1.ElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elastictranscoderpresetsResource, c.ns, name), &v1alpha1.ElastictranscoderPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastictranscoderPreset), err
}

// List takes label and field selectors, and returns the list of ElastictranscoderPresets that match those selectors.
func (c *FakeElastictranscoderPresets) List(opts v1.ListOptions) (result *v1alpha1.ElastictranscoderPresetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elastictranscoderpresetsResource, elastictranscoderpresetsKind, c.ns, opts), &v1alpha1.ElastictranscoderPresetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElastictranscoderPresetList{ListMeta: obj.(*v1alpha1.ElastictranscoderPresetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElastictranscoderPresetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elastictranscoderPresets.
func (c *FakeElastictranscoderPresets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elastictranscoderpresetsResource, c.ns, opts))

}

// Create takes the representation of a elastictranscoderPreset and creates it.  Returns the server's representation of the elastictranscoderPreset, and an error, if there is any.
func (c *FakeElastictranscoderPresets) Create(elastictranscoderPreset *v1alpha1.ElastictranscoderPreset) (result *v1alpha1.ElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elastictranscoderpresetsResource, c.ns, elastictranscoderPreset), &v1alpha1.ElastictranscoderPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastictranscoderPreset), err
}

// Update takes the representation of a elastictranscoderPreset and updates it. Returns the server's representation of the elastictranscoderPreset, and an error, if there is any.
func (c *FakeElastictranscoderPresets) Update(elastictranscoderPreset *v1alpha1.ElastictranscoderPreset) (result *v1alpha1.ElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elastictranscoderpresetsResource, c.ns, elastictranscoderPreset), &v1alpha1.ElastictranscoderPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastictranscoderPreset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElastictranscoderPresets) UpdateStatus(elastictranscoderPreset *v1alpha1.ElastictranscoderPreset) (*v1alpha1.ElastictranscoderPreset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elastictranscoderpresetsResource, "status", c.ns, elastictranscoderPreset), &v1alpha1.ElastictranscoderPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastictranscoderPreset), err
}

// Delete takes name of the elastictranscoderPreset and deletes it. Returns an error if one occurs.
func (c *FakeElastictranscoderPresets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(elastictranscoderpresetsResource, c.ns, name), &v1alpha1.ElastictranscoderPreset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElastictranscoderPresets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elastictranscoderpresetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElastictranscoderPresetList{})
	return err
}

// Patch applies the patch and returns the patched elastictranscoderPreset.
func (c *FakeElastictranscoderPresets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElastictranscoderPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elastictranscoderpresetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ElastictranscoderPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastictranscoderPreset), err
}
