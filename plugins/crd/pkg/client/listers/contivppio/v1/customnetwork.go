// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Shivkb/contiv-vpp/plugins/crd/pkg/apis/contivppio/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomNetworkLister helps list CustomNetworks.
type CustomNetworkLister interface {
	// List lists all CustomNetworks in the indexer.
	List(selector labels.Selector) (ret []*v1.CustomNetwork, err error)
	// CustomNetworks returns an object that can list and get CustomNetworks.
	CustomNetworks(namespace string) CustomNetworkNamespaceLister
	CustomNetworkListerExpansion
}

// customNetworkLister implements the CustomNetworkLister interface.
type customNetworkLister struct {
	indexer cache.Indexer
}

// NewCustomNetworkLister returns a new CustomNetworkLister.
func NewCustomNetworkLister(indexer cache.Indexer) CustomNetworkLister {
	return &customNetworkLister{indexer: indexer}
}

// List lists all CustomNetworks in the indexer.
func (s *customNetworkLister) List(selector labels.Selector) (ret []*v1.CustomNetwork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CustomNetwork))
	})
	return ret, err
}

// CustomNetworks returns an object that can list and get CustomNetworks.
func (s *customNetworkLister) CustomNetworks(namespace string) CustomNetworkNamespaceLister {
	return customNetworkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CustomNetworkNamespaceLister helps list and get CustomNetworks.
type CustomNetworkNamespaceLister interface {
	// List lists all CustomNetworks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CustomNetwork, err error)
	// Get retrieves the CustomNetwork from the indexer for a given namespace and name.
	Get(name string) (*v1.CustomNetwork, error)
	CustomNetworkNamespaceListerExpansion
}

// customNetworkNamespaceLister implements the CustomNetworkNamespaceLister
// interface.
type customNetworkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CustomNetworks in the indexer for a given namespace.
func (s customNetworkNamespaceLister) List(selector labels.Selector) (ret []*v1.CustomNetwork, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CustomNetwork))
	})
	return ret, err
}

// Get retrieves the CustomNetwork from the indexer for a given namespace and name.
func (s customNetworkNamespaceLister) Get(name string) (*v1.CustomNetwork, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("customnetwork"), name)
	}
	return obj.(*v1.CustomNetwork), nil
}
