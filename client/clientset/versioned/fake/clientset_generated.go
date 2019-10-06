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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "kubeform.dev/kubeform/client/clientset/versioned"
	awsv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/aws/v1alpha1"
	fakeawsv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/aws/v1alpha1/fake"
	azurermv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/azurerm/v1alpha1"
	fakeazurermv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/azurerm/v1alpha1/fake"
	basev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/base/v1alpha1"
	fakebasev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/base/v1alpha1/fake"
	digitaloceanv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/digitalocean/v1alpha1"
	fakedigitaloceanv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/digitalocean/v1alpha1/fake"
	googlev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/google/v1alpha1"
	fakegooglev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/google/v1alpha1/fake"
	linodev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/linode/v1alpha1"
	fakelinodev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/linode/v1alpha1/fake"
	modulesv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/modules/v1alpha1"
	fakemodulesv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/modules/v1alpha1/fake"
	basev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/register.go/v1alpha1"
	fakebasev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/register.go/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// AwsV1alpha1 retrieves the AwsV1alpha1Client
func (c *Clientset) AwsV1alpha1() awsv1alpha1.AwsV1alpha1Interface {
	return &fakeawsv1alpha1.FakeAwsV1alpha1{Fake: &c.Fake}
}

// AzurermV1alpha1 retrieves the AzurermV1alpha1Client
func (c *Clientset) AzurermV1alpha1() azurermv1alpha1.AzurermV1alpha1Interface {
	return &fakeazurermv1alpha1.FakeAzurermV1alpha1{Fake: &c.Fake}
}

// BaseV1alpha1 retrieves the BaseV1alpha1Client
func (c *Clientset) BaseV1alpha1() basev1alpha1.BaseV1alpha1Interface {
	return &fakebasev1alpha1.FakeBaseV1alpha1{Fake: &c.Fake}
}

// DigitaloceanV1alpha1 retrieves the DigitaloceanV1alpha1Client
func (c *Clientset) DigitaloceanV1alpha1() digitaloceanv1alpha1.DigitaloceanV1alpha1Interface {
	return &fakedigitaloceanv1alpha1.FakeDigitaloceanV1alpha1{Fake: &c.Fake}
}

// GoogleV1alpha1 retrieves the GoogleV1alpha1Client
func (c *Clientset) GoogleV1alpha1() googlev1alpha1.GoogleV1alpha1Interface {
	return &fakegooglev1alpha1.FakeGoogleV1alpha1{Fake: &c.Fake}
}

// LinodeV1alpha1 retrieves the LinodeV1alpha1Client
func (c *Clientset) LinodeV1alpha1() linodev1alpha1.LinodeV1alpha1Interface {
	return &fakelinodev1alpha1.FakeLinodeV1alpha1{Fake: &c.Fake}
}

// ModulesV1alpha1 retrieves the ModulesV1alpha1Client
func (c *Clientset) ModulesV1alpha1() modulesv1alpha1.ModulesV1alpha1Interface {
	return &fakemodulesv1alpha1.FakeModulesV1alpha1{Fake: &c.Fake}
}

// BaseV1alpha1 retrieves the BaseV1alpha1Client
func (c *Clientset) BaseV1alpha1() basev1alpha1.BaseV1alpha1Interface {
	return &fakebasev1alpha1.FakeBaseV1alpha1{Fake: &c.Fake}
}
