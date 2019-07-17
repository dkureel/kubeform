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

// FakeSqlElasticpools implements SqlElasticpoolInterface
type FakeSqlElasticpools struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var sqlelasticpoolsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "sqlelasticpools"}

var sqlelasticpoolsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SqlElasticpool"}

// Get takes name of the sqlElasticpool, and returns the corresponding sqlElasticpool object, and an error if there is any.
func (c *FakeSqlElasticpools) Get(name string, options v1.GetOptions) (result *v1alpha1.SqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqlelasticpoolsResource, c.ns, name), &v1alpha1.SqlElasticpool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlElasticpool), err
}

// List takes label and field selectors, and returns the list of SqlElasticpools that match those selectors.
func (c *FakeSqlElasticpools) List(opts v1.ListOptions) (result *v1alpha1.SqlElasticpoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqlelasticpoolsResource, sqlelasticpoolsKind, c.ns, opts), &v1alpha1.SqlElasticpoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqlElasticpoolList{ListMeta: obj.(*v1alpha1.SqlElasticpoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqlElasticpoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlElasticpools.
func (c *FakeSqlElasticpools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqlelasticpoolsResource, c.ns, opts))

}

// Create takes the representation of a sqlElasticpool and creates it.  Returns the server's representation of the sqlElasticpool, and an error, if there is any.
func (c *FakeSqlElasticpools) Create(sqlElasticpool *v1alpha1.SqlElasticpool) (result *v1alpha1.SqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqlelasticpoolsResource, c.ns, sqlElasticpool), &v1alpha1.SqlElasticpool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlElasticpool), err
}

// Update takes the representation of a sqlElasticpool and updates it. Returns the server's representation of the sqlElasticpool, and an error, if there is any.
func (c *FakeSqlElasticpools) Update(sqlElasticpool *v1alpha1.SqlElasticpool) (result *v1alpha1.SqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqlelasticpoolsResource, c.ns, sqlElasticpool), &v1alpha1.SqlElasticpool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlElasticpool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqlElasticpools) UpdateStatus(sqlElasticpool *v1alpha1.SqlElasticpool) (*v1alpha1.SqlElasticpool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqlelasticpoolsResource, "status", c.ns, sqlElasticpool), &v1alpha1.SqlElasticpool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlElasticpool), err
}

// Delete takes name of the sqlElasticpool and deletes it. Returns an error if one occurs.
func (c *FakeSqlElasticpools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqlelasticpoolsResource, c.ns, name), &v1alpha1.SqlElasticpool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlElasticpools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqlelasticpoolsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqlElasticpoolList{})
	return err
}

// Patch applies the patch and returns the patched sqlElasticpool.
func (c *FakeSqlElasticpools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlElasticpool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqlelasticpoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SqlElasticpool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlElasticpool), err
}
