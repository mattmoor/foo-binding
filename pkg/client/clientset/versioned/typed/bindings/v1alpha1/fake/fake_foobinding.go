/*
Copyright 2019 The Knative Authors

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
	v1alpha1 "github.com/mattmoor/foo-binding/pkg/apis/bindings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFooBindings implements FooBindingInterface
type FakeFooBindings struct {
	Fake *FakeBindingsV1alpha1
	ns   string
}

var foobindingsResource = schema.GroupVersionResource{Group: "bindings.mattmoor.dev", Version: "v1alpha1", Resource: "foobindings"}

var foobindingsKind = schema.GroupVersionKind{Group: "bindings.mattmoor.dev", Version: "v1alpha1", Kind: "FooBinding"}

// Get takes name of the fooBinding, and returns the corresponding fooBinding object, and an error if there is any.
func (c *FakeFooBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.FooBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(foobindingsResource, c.ns, name), &v1alpha1.FooBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FooBinding), err
}

// List takes label and field selectors, and returns the list of FooBindings that match those selectors.
func (c *FakeFooBindings) List(opts v1.ListOptions) (result *v1alpha1.FooBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(foobindingsResource, foobindingsKind, c.ns, opts), &v1alpha1.FooBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FooBindingList{ListMeta: obj.(*v1alpha1.FooBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.FooBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fooBindings.
func (c *FakeFooBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(foobindingsResource, c.ns, opts))

}

// Create takes the representation of a fooBinding and creates it.  Returns the server's representation of the fooBinding, and an error, if there is any.
func (c *FakeFooBindings) Create(fooBinding *v1alpha1.FooBinding) (result *v1alpha1.FooBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(foobindingsResource, c.ns, fooBinding), &v1alpha1.FooBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FooBinding), err
}

// Update takes the representation of a fooBinding and updates it. Returns the server's representation of the fooBinding, and an error, if there is any.
func (c *FakeFooBindings) Update(fooBinding *v1alpha1.FooBinding) (result *v1alpha1.FooBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(foobindingsResource, c.ns, fooBinding), &v1alpha1.FooBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FooBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFooBindings) UpdateStatus(fooBinding *v1alpha1.FooBinding) (*v1alpha1.FooBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(foobindingsResource, "status", c.ns, fooBinding), &v1alpha1.FooBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FooBinding), err
}

// Delete takes name of the fooBinding and deletes it. Returns an error if one occurs.
func (c *FakeFooBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(foobindingsResource, c.ns, name), &v1alpha1.FooBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFooBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(foobindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FooBindingList{})
	return err
}

// Patch applies the patch and returns the patched fooBinding.
func (c *FakeFooBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FooBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(foobindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FooBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FooBinding), err
}