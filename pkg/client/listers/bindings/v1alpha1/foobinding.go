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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mattmoor/foo-binding/pkg/apis/bindings/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FooBindingLister helps list FooBindings.
type FooBindingLister interface {
	// List lists all FooBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FooBinding, err error)
	// FooBindings returns an object that can list and get FooBindings.
	FooBindings(namespace string) FooBindingNamespaceLister
	FooBindingListerExpansion
}

// fooBindingLister implements the FooBindingLister interface.
type fooBindingLister struct {
	indexer cache.Indexer
}

// NewFooBindingLister returns a new FooBindingLister.
func NewFooBindingLister(indexer cache.Indexer) FooBindingLister {
	return &fooBindingLister{indexer: indexer}
}

// List lists all FooBindings in the indexer.
func (s *fooBindingLister) List(selector labels.Selector) (ret []*v1alpha1.FooBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FooBinding))
	})
	return ret, err
}

// FooBindings returns an object that can list and get FooBindings.
func (s *fooBindingLister) FooBindings(namespace string) FooBindingNamespaceLister {
	return fooBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FooBindingNamespaceLister helps list and get FooBindings.
type FooBindingNamespaceLister interface {
	// List lists all FooBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FooBinding, err error)
	// Get retrieves the FooBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FooBinding, error)
	FooBindingNamespaceListerExpansion
}

// fooBindingNamespaceLister implements the FooBindingNamespaceLister
// interface.
type fooBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FooBindings in the indexer for a given namespace.
func (s fooBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FooBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FooBinding))
	})
	return ret, err
}

// Get retrieves the FooBinding from the indexer for a given namespace and name.
func (s fooBindingNamespaceLister) Get(name string) (*v1alpha1.FooBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("foobinding"), name)
	}
	return obj.(*v1alpha1.FooBinding), nil
}
