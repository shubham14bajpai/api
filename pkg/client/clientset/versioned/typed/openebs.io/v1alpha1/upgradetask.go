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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/openebs/api/v2/pkg/apis/openebs.io/v1alpha1"
	scheme "github.com/openebs/api/v2/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UpgradeTasksGetter has a method to return a UpgradeTaskInterface.
// A group's client should implement this interface.
type UpgradeTasksGetter interface {
	UpgradeTasks(namespace string) UpgradeTaskInterface
}

// UpgradeTaskInterface has methods to work with UpgradeTask resources.
type UpgradeTaskInterface interface {
	Create(*v1alpha1.UpgradeTask) (*v1alpha1.UpgradeTask, error)
	Update(*v1alpha1.UpgradeTask) (*v1alpha1.UpgradeTask, error)
	UpdateStatus(*v1alpha1.UpgradeTask) (*v1alpha1.UpgradeTask, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.UpgradeTask, error)
	List(opts v1.ListOptions) (*v1alpha1.UpgradeTaskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UpgradeTask, err error)
	UpgradeTaskExpansion
}

// upgradeTasks implements UpgradeTaskInterface
type upgradeTasks struct {
	client rest.Interface
	ns     string
}

// newUpgradeTasks returns a UpgradeTasks
func newUpgradeTasks(c *OpenebsV1alpha1Client, namespace string) *upgradeTasks {
	return &upgradeTasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the upgradeTask, and returns the corresponding upgradeTask object, and an error if there is any.
func (c *upgradeTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.UpgradeTask, err error) {
	result = &v1alpha1.UpgradeTask{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upgradetasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UpgradeTasks that match those selectors.
func (c *upgradeTasks) List(opts v1.ListOptions) (result *v1alpha1.UpgradeTaskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UpgradeTaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upgradetasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested upgradeTasks.
func (c *upgradeTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("upgradetasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a upgradeTask and creates it.  Returns the server's representation of the upgradeTask, and an error, if there is any.
func (c *upgradeTasks) Create(upgradeTask *v1alpha1.UpgradeTask) (result *v1alpha1.UpgradeTask, err error) {
	result = &v1alpha1.UpgradeTask{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("upgradetasks").
		Body(upgradeTask).
		Do().
		Into(result)
	return
}

// Update takes the representation of a upgradeTask and updates it. Returns the server's representation of the upgradeTask, and an error, if there is any.
func (c *upgradeTasks) Update(upgradeTask *v1alpha1.UpgradeTask) (result *v1alpha1.UpgradeTask, err error) {
	result = &v1alpha1.UpgradeTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upgradetasks").
		Name(upgradeTask.Name).
		Body(upgradeTask).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *upgradeTasks) UpdateStatus(upgradeTask *v1alpha1.UpgradeTask) (result *v1alpha1.UpgradeTask, err error) {
	result = &v1alpha1.UpgradeTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upgradetasks").
		Name(upgradeTask.Name).
		SubResource("status").
		Body(upgradeTask).
		Do().
		Into(result)
	return
}

// Delete takes name of the upgradeTask and deletes it. Returns an error if one occurs.
func (c *upgradeTasks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upgradetasks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *upgradeTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upgradetasks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched upgradeTask.
func (c *upgradeTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UpgradeTask, err error) {
	result = &v1alpha1.UpgradeTask{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("upgradetasks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
