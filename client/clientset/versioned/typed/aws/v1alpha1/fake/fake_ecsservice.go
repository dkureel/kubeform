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

// FakeEcsServices implements EcsServiceInterface
type FakeEcsServices struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ecsservicesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ecsservices"}

var ecsservicesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EcsService"}

// Get takes name of the ecsService, and returns the corresponding ecsService object, and an error if there is any.
func (c *FakeEcsServices) Get(name string, options v1.GetOptions) (result *v1alpha1.EcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ecsservicesResource, c.ns, name), &v1alpha1.EcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsService), err
}

// List takes label and field selectors, and returns the list of EcsServices that match those selectors.
func (c *FakeEcsServices) List(opts v1.ListOptions) (result *v1alpha1.EcsServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ecsservicesResource, ecsservicesKind, c.ns, opts), &v1alpha1.EcsServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EcsServiceList{ListMeta: obj.(*v1alpha1.EcsServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.EcsServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ecsServices.
func (c *FakeEcsServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ecsservicesResource, c.ns, opts))

}

// Create takes the representation of a ecsService and creates it.  Returns the server's representation of the ecsService, and an error, if there is any.
func (c *FakeEcsServices) Create(ecsService *v1alpha1.EcsService) (result *v1alpha1.EcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ecsservicesResource, c.ns, ecsService), &v1alpha1.EcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsService), err
}

// Update takes the representation of a ecsService and updates it. Returns the server's representation of the ecsService, and an error, if there is any.
func (c *FakeEcsServices) Update(ecsService *v1alpha1.EcsService) (result *v1alpha1.EcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ecsservicesResource, c.ns, ecsService), &v1alpha1.EcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEcsServices) UpdateStatus(ecsService *v1alpha1.EcsService) (*v1alpha1.EcsService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ecsservicesResource, "status", c.ns, ecsService), &v1alpha1.EcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsService), err
}

// Delete takes name of the ecsService and deletes it. Returns an error if one occurs.
func (c *FakeEcsServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ecsservicesResource, c.ns, name), &v1alpha1.EcsService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEcsServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ecsservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EcsServiceList{})
	return err
}

// Patch applies the patch and returns the patched ecsService.
func (c *FakeEcsServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ecsservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EcsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsService), err
}
