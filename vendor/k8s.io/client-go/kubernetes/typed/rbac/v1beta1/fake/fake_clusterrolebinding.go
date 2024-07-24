/*
Copyright The Kubernetes Authors.

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
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "k8s.io/api/rbac/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rbacv1beta1 "k8s.io/client-go/applyconfigurations/rbac/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRoleBindings implements ClusterRoleBindingInterface
type FakeClusterRoleBindings struct {
	Fake *FakeRbacV1beta1
}

var clusterrolebindingsResource = v1beta1.SchemeGroupVersion.WithResource("clusterrolebindings")

var clusterrolebindingsKind = v1beta1.SchemeGroupVersion.WithKind("ClusterRoleBinding")

// Get takes name of the clusterRoleBinding, and returns the corresponding clusterRoleBinding object, and an error if there is any.
func (c *FakeClusterRoleBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterrolebindingsResource, name), &v1beta1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterRoleBinding), err
}

// List takes label and field selectors, and returns the list of ClusterRoleBindings that match those selectors.
func (c *FakeClusterRoleBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterRoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterrolebindingsResource, clusterrolebindingsKind, opts), &v1beta1.ClusterRoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterRoleBindingList{ListMeta: obj.(*v1beta1.ClusterRoleBindingList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoleBindings.
func (c *FakeClusterRoleBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterrolebindingsResource, opts))
}

// Create takes the representation of a clusterRoleBinding and creates it.  Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *FakeClusterRoleBindings) Create(ctx context.Context, clusterRoleBinding *v1beta1.ClusterRoleBinding, opts v1.CreateOptions) (result *v1beta1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterrolebindingsResource, clusterRoleBinding), &v1beta1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterRoleBinding), err
}

// Update takes the representation of a clusterRoleBinding and updates it. Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *FakeClusterRoleBindings) Update(ctx context.Context, clusterRoleBinding *v1beta1.ClusterRoleBinding, opts v1.UpdateOptions) (result *v1beta1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterrolebindingsResource, clusterRoleBinding), &v1beta1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterRoleBinding), err
}

// Delete takes name of the clusterRoleBinding and deletes it. Returns an error if one occurs.
func (c *FakeClusterRoleBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterrolebindingsResource, name, opts), &v1beta1.ClusterRoleBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRoleBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterrolebindingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ClusterRoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched clusterRoleBinding.
func (c *FakeClusterRoleBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrolebindingsResource, name, pt, data, subresources...), &v1beta1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterRoleBinding), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterRoleBinding.
func (c *FakeClusterRoleBindings) Apply(ctx context.Context, clusterRoleBinding *rbacv1beta1.ClusterRoleBindingApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.ClusterRoleBinding, err error) {
	if clusterRoleBinding == nil {
		return nil, fmt.Errorf("clusterRoleBinding provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterRoleBinding)
	if err != nil {
		return nil, err
	}
	name := clusterRoleBinding.Name
	if name == nil {
		return nil, fmt.Errorf("clusterRoleBinding.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrolebindingsResource, *name, types.ApplyPatchType, data), &v1beta1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterRoleBinding), err
}
