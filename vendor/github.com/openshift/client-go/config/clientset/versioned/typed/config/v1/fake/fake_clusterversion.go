// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterVersions implements ClusterVersionInterface
type FakeClusterVersions struct {
	Fake *FakeConfigV1
}

var clusterversionsResource = v1.SchemeGroupVersion.WithResource("clusterversions")

var clusterversionsKind = v1.SchemeGroupVersion.WithKind("ClusterVersion")

// Get takes name of the clusterVersion, and returns the corresponding clusterVersion object, and an error if there is any.
func (c *FakeClusterVersions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterversionsResource, name), &v1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterVersion), err
}

// List takes label and field selectors, and returns the list of ClusterVersions that match those selectors.
func (c *FakeClusterVersions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterversionsResource, clusterversionsKind, opts), &v1.ClusterVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterVersionList{ListMeta: obj.(*v1.ClusterVersionList).ListMeta}
	for _, item := range obj.(*v1.ClusterVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterVersions.
func (c *FakeClusterVersions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterversionsResource, opts))
}

// Create takes the representation of a clusterVersion and creates it.  Returns the server's representation of the clusterVersion, and an error, if there is any.
func (c *FakeClusterVersions) Create(ctx context.Context, clusterVersion *v1.ClusterVersion, opts metav1.CreateOptions) (result *v1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterversionsResource, clusterVersion), &v1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterVersion), err
}

// Update takes the representation of a clusterVersion and updates it. Returns the server's representation of the clusterVersion, and an error, if there is any.
func (c *FakeClusterVersions) Update(ctx context.Context, clusterVersion *v1.ClusterVersion, opts metav1.UpdateOptions) (result *v1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterversionsResource, clusterVersion), &v1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterVersions) UpdateStatus(ctx context.Context, clusterVersion *v1.ClusterVersion, opts metav1.UpdateOptions) (*v1.ClusterVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterversionsResource, "status", clusterVersion), &v1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterVersion), err
}

// Delete takes name of the clusterVersion and deletes it. Returns an error if one occurs.
func (c *FakeClusterVersions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterversionsResource, name, opts), &v1.ClusterVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterVersions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterversionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterVersionList{})
	return err
}

// Patch applies the patch and returns the patched clusterVersion.
func (c *FakeClusterVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterversionsResource, name, pt, data, subresources...), &v1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterVersion), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterVersion.
func (c *FakeClusterVersions) Apply(ctx context.Context, clusterVersion *configv1.ClusterVersionApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterVersion, err error) {
	if clusterVersion == nil {
		return nil, fmt.Errorf("clusterVersion provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterVersion)
	if err != nil {
		return nil, err
	}
	name := clusterVersion.Name
	if name == nil {
		return nil, fmt.Errorf("clusterVersion.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterversionsResource, *name, types.ApplyPatchType, data), &v1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterVersion), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeClusterVersions) ApplyStatus(ctx context.Context, clusterVersion *configv1.ClusterVersionApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterVersion, err error) {
	if clusterVersion == nil {
		return nil, fmt.Errorf("clusterVersion provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterVersion)
	if err != nil {
		return nil, err
	}
	name := clusterVersion.Name
	if name == nil {
		return nil, fmt.Errorf("clusterVersion.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterversionsResource, *name, types.ApplyPatchType, data, "status"), &v1.ClusterVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterVersion), err
}
