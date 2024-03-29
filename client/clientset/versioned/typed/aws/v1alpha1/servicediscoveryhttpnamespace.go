/*
Copyright The Kubeform Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceDiscoveryHTTPNamespacesGetter has a method to return a ServiceDiscoveryHTTPNamespaceInterface.
// A group's client should implement this interface.
type ServiceDiscoveryHTTPNamespacesGetter interface {
	ServiceDiscoveryHTTPNamespaces(namespace string) ServiceDiscoveryHTTPNamespaceInterface
}

// ServiceDiscoveryHTTPNamespaceInterface has methods to work with ServiceDiscoveryHTTPNamespace resources.
type ServiceDiscoveryHTTPNamespaceInterface interface {
	Create(*v1alpha1.ServiceDiscoveryHTTPNamespace) (*v1alpha1.ServiceDiscoveryHTTPNamespace, error)
	Update(*v1alpha1.ServiceDiscoveryHTTPNamespace) (*v1alpha1.ServiceDiscoveryHTTPNamespace, error)
	UpdateStatus(*v1alpha1.ServiceDiscoveryHTTPNamespace) (*v1alpha1.ServiceDiscoveryHTTPNamespace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceDiscoveryHTTPNamespace, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceDiscoveryHTTPNamespaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDiscoveryHTTPNamespace, err error)
	ServiceDiscoveryHTTPNamespaceExpansion
}

// serviceDiscoveryHTTPNamespaces implements ServiceDiscoveryHTTPNamespaceInterface
type serviceDiscoveryHTTPNamespaces struct {
	client rest.Interface
	ns     string
}

// newServiceDiscoveryHTTPNamespaces returns a ServiceDiscoveryHTTPNamespaces
func newServiceDiscoveryHTTPNamespaces(c *AwsV1alpha1Client, namespace string) *serviceDiscoveryHTTPNamespaces {
	return &serviceDiscoveryHTTPNamespaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceDiscoveryHTTPNamespace, and returns the corresponding serviceDiscoveryHTTPNamespace object, and an error if there is any.
func (c *serviceDiscoveryHTTPNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceDiscoveryHTTPNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryHTTPNamespace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceDiscoveryHTTPNamespaces that match those selectors.
func (c *serviceDiscoveryHTTPNamespaces) List(opts v1.ListOptions) (result *v1alpha1.ServiceDiscoveryHTTPNamespaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceDiscoveryHTTPNamespaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceDiscoveryHTTPNamespaces.
func (c *serviceDiscoveryHTTPNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a serviceDiscoveryHTTPNamespace and creates it.  Returns the server's representation of the serviceDiscoveryHTTPNamespace, and an error, if there is any.
func (c *serviceDiscoveryHTTPNamespaces) Create(serviceDiscoveryHTTPNamespace *v1alpha1.ServiceDiscoveryHTTPNamespace) (result *v1alpha1.ServiceDiscoveryHTTPNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryHTTPNamespace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		Body(serviceDiscoveryHTTPNamespace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceDiscoveryHTTPNamespace and updates it. Returns the server's representation of the serviceDiscoveryHTTPNamespace, and an error, if there is any.
func (c *serviceDiscoveryHTTPNamespaces) Update(serviceDiscoveryHTTPNamespace *v1alpha1.ServiceDiscoveryHTTPNamespace) (result *v1alpha1.ServiceDiscoveryHTTPNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryHTTPNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		Name(serviceDiscoveryHTTPNamespace.Name).
		Body(serviceDiscoveryHTTPNamespace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *serviceDiscoveryHTTPNamespaces) UpdateStatus(serviceDiscoveryHTTPNamespace *v1alpha1.ServiceDiscoveryHTTPNamespace) (result *v1alpha1.ServiceDiscoveryHTTPNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryHTTPNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		Name(serviceDiscoveryHTTPNamespace.Name).
		SubResource("status").
		Body(serviceDiscoveryHTTPNamespace).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceDiscoveryHTTPNamespace and deletes it. Returns an error if one occurs.
func (c *serviceDiscoveryHTTPNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceDiscoveryHTTPNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceDiscoveryHTTPNamespace.
func (c *serviceDiscoveryHTTPNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDiscoveryHTTPNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryHTTPNamespace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicediscoveryhttpnamespaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
