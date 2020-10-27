/*
Copyright 2020 The OpenEBS Authors

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
	cstorv1 "github.com/openebs/api/v2/pkg/apis/cstor/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCStorPoolClusters implements CStorPoolClusterInterface
type FakeCStorPoolClusters struct {
	Fake *FakeCstorV1
	ns   string
}

var cstorpoolclustersResource = schema.GroupVersionResource{Group: "cstor.openebs.io", Version: "v1", Resource: "cstorpoolclusters"}

var cstorpoolclustersKind = schema.GroupVersionKind{Group: "cstor.openebs.io", Version: "v1", Kind: "CStorPoolCluster"}

// Get takes name of the cStorPoolCluster, and returns the corresponding cStorPoolCluster object, and an error if there is any.
func (c *FakeCStorPoolClusters) Get(name string, options v1.GetOptions) (result *cstorv1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cstorpoolclustersResource, c.ns, name), &cstorv1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolCluster), err
}

// List takes label and field selectors, and returns the list of CStorPoolClusters that match those selectors.
func (c *FakeCStorPoolClusters) List(opts v1.ListOptions) (result *cstorv1.CStorPoolClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cstorpoolclustersResource, cstorpoolclustersKind, c.ns, opts), &cstorv1.CStorPoolClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cstorv1.CStorPoolClusterList{ListMeta: obj.(*cstorv1.CStorPoolClusterList).ListMeta}
	for _, item := range obj.(*cstorv1.CStorPoolClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cStorPoolClusters.
func (c *FakeCStorPoolClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cstorpoolclustersResource, c.ns, opts))

}

// Create takes the representation of a cStorPoolCluster and creates it.  Returns the server's representation of the cStorPoolCluster, and an error, if there is any.
func (c *FakeCStorPoolClusters) Create(cStorPoolCluster *cstorv1.CStorPoolCluster) (result *cstorv1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cstorpoolclustersResource, c.ns, cStorPoolCluster), &cstorv1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolCluster), err
}

// Update takes the representation of a cStorPoolCluster and updates it. Returns the server's representation of the cStorPoolCluster, and an error, if there is any.
func (c *FakeCStorPoolClusters) Update(cStorPoolCluster *cstorv1.CStorPoolCluster) (result *cstorv1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cstorpoolclustersResource, c.ns, cStorPoolCluster), &cstorv1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolCluster), err
}

// Delete takes name of the cStorPoolCluster and deletes it. Returns an error if one occurs.
func (c *FakeCStorPoolClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cstorpoolclustersResource, c.ns, name), &cstorv1.CStorPoolCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCStorPoolClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cstorpoolclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cstorv1.CStorPoolClusterList{})
	return err
}

// Patch applies the patch and returns the patched cStorPoolCluster.
func (c *FakeCStorPoolClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cstorv1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cstorpoolclustersResource, c.ns, name, pt, data, subresources...), &cstorv1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cstorv1.CStorPoolCluster), err
}
