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

// FakeApiGatewayRestApis implements ApiGatewayRestApiInterface
type FakeApiGatewayRestApis struct {
	Fake *FakeAwsV1alpha1
}

var apigatewayrestapisResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewayrestapis"}

var apigatewayrestapisKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayRestApi"}

// Get takes name of the apiGatewayRestApi, and returns the corresponding apiGatewayRestApi object, and an error if there is any.
func (c *FakeApiGatewayRestApis) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apigatewayrestapisResource, name), &v1alpha1.ApiGatewayRestApi{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestApi), err
}

// List takes label and field selectors, and returns the list of ApiGatewayRestApis that match those selectors.
func (c *FakeApiGatewayRestApis) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayRestApiList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apigatewayrestapisResource, apigatewayrestapisKind, opts), &v1alpha1.ApiGatewayRestApiList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayRestApiList{ListMeta: obj.(*v1alpha1.ApiGatewayRestApiList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayRestApiList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayRestApis.
func (c *FakeApiGatewayRestApis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apigatewayrestapisResource, opts))
}

// Create takes the representation of a apiGatewayRestApi and creates it.  Returns the server's representation of the apiGatewayRestApi, and an error, if there is any.
func (c *FakeApiGatewayRestApis) Create(apiGatewayRestApi *v1alpha1.ApiGatewayRestApi) (result *v1alpha1.ApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apigatewayrestapisResource, apiGatewayRestApi), &v1alpha1.ApiGatewayRestApi{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestApi), err
}

// Update takes the representation of a apiGatewayRestApi and updates it. Returns the server's representation of the apiGatewayRestApi, and an error, if there is any.
func (c *FakeApiGatewayRestApis) Update(apiGatewayRestApi *v1alpha1.ApiGatewayRestApi) (result *v1alpha1.ApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apigatewayrestapisResource, apiGatewayRestApi), &v1alpha1.ApiGatewayRestApi{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestApi), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayRestApis) UpdateStatus(apiGatewayRestApi *v1alpha1.ApiGatewayRestApi) (*v1alpha1.ApiGatewayRestApi, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apigatewayrestapisResource, "status", apiGatewayRestApi), &v1alpha1.ApiGatewayRestApi{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestApi), err
}

// Delete takes name of the apiGatewayRestApi and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayRestApis) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apigatewayrestapisResource, name), &v1alpha1.ApiGatewayRestApi{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayRestApis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apigatewayrestapisResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayRestApiList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayRestApi.
func (c *FakeApiGatewayRestApis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayRestApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apigatewayrestapisResource, name, pt, data, subresources...), &v1alpha1.ApiGatewayRestApi{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayRestApi), err
}