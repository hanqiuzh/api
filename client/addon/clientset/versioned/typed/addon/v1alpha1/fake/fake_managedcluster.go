// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/open-cluster-management/api/addon/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedClusters implements ManagedClusterInterface
type FakeManagedClusters struct {
	Fake *FakeAddonV1alpha1
}

var managedclustersResource = schema.GroupVersionResource{Group: "addon.open-cluster-management.io", Version: "v1alpha1", Resource: "managedclusters"}

var managedclustersKind = schema.GroupVersionKind{Group: "addon.open-cluster-management.io", Version: "v1alpha1", Kind: "ManagedCluster"}

// Get takes name of the managedCluster, and returns the corresponding managedCluster object, and an error if there is any.
func (c *FakeManagedClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(managedclustersResource, name), &v1alpha1.ManagedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedCluster), err
}

// List takes label and field selectors, and returns the list of ManagedClusters that match those selectors.
func (c *FakeManagedClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(managedclustersResource, managedclustersKind, opts), &v1alpha1.ManagedClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedClusterList{ListMeta: obj.(*v1alpha1.ManagedClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedClusters.
func (c *FakeManagedClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(managedclustersResource, opts))
}

// Create takes the representation of a managedCluster and creates it.  Returns the server's representation of the managedCluster, and an error, if there is any.
func (c *FakeManagedClusters) Create(ctx context.Context, managedCluster *v1alpha1.ManagedCluster, opts v1.CreateOptions) (result *v1alpha1.ManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(managedclustersResource, managedCluster), &v1alpha1.ManagedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedCluster), err
}

// Update takes the representation of a managedCluster and updates it. Returns the server's representation of the managedCluster, and an error, if there is any.
func (c *FakeManagedClusters) Update(ctx context.Context, managedCluster *v1alpha1.ManagedCluster, opts v1.UpdateOptions) (result *v1alpha1.ManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(managedclustersResource, managedCluster), &v1alpha1.ManagedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedClusters) UpdateStatus(ctx context.Context, managedCluster *v1alpha1.ManagedCluster, opts v1.UpdateOptions) (*v1alpha1.ManagedCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(managedclustersResource, "status", managedCluster), &v1alpha1.ManagedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedCluster), err
}

// Delete takes name of the managedCluster and deletes it. Returns an error if one occurs.
func (c *FakeManagedClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(managedclustersResource, name), &v1alpha1.ManagedCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(managedclustersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedClusterList{})
	return err
}

// Patch applies the patch and returns the patched managedCluster.
func (c *FakeManagedClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(managedclustersResource, name, pt, data, subresources...), &v1alpha1.ManagedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedCluster), err
}
