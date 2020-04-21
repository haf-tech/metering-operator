// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	meteringv1 "github.com/kubernetes-reporting/metering-operator/pkg/apis/metering/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMeteringConfigs implements MeteringConfigInterface
type FakeMeteringConfigs struct {
	Fake *FakeMeteringV1
	ns   string
}

var meteringconfigsResource = schema.GroupVersionResource{Group: "metering.openshift.io", Version: "v1", Resource: "meteringconfigs"}

var meteringconfigsKind = schema.GroupVersionKind{Group: "metering.openshift.io", Version: "v1", Kind: "MeteringConfig"}

// Get takes name of the meteringConfig, and returns the corresponding meteringConfig object, and an error if there is any.
func (c *FakeMeteringConfigs) Get(name string, options v1.GetOptions) (result *meteringv1.MeteringConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(meteringconfigsResource, c.ns, name), &meteringv1.MeteringConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.MeteringConfig), err
}

// List takes label and field selectors, and returns the list of MeteringConfigs that match those selectors.
func (c *FakeMeteringConfigs) List(opts v1.ListOptions) (result *meteringv1.MeteringConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(meteringconfigsResource, meteringconfigsKind, c.ns, opts), &meteringv1.MeteringConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &meteringv1.MeteringConfigList{ListMeta: obj.(*meteringv1.MeteringConfigList).ListMeta}
	for _, item := range obj.(*meteringv1.MeteringConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested meteringConfigs.
func (c *FakeMeteringConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(meteringconfigsResource, c.ns, opts))

}

// Create takes the representation of a meteringConfig and creates it.  Returns the server's representation of the meteringConfig, and an error, if there is any.
func (c *FakeMeteringConfigs) Create(meteringConfig *meteringv1.MeteringConfig) (result *meteringv1.MeteringConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(meteringconfigsResource, c.ns, meteringConfig), &meteringv1.MeteringConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.MeteringConfig), err
}

// Update takes the representation of a meteringConfig and updates it. Returns the server's representation of the meteringConfig, and an error, if there is any.
func (c *FakeMeteringConfigs) Update(meteringConfig *meteringv1.MeteringConfig) (result *meteringv1.MeteringConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(meteringconfigsResource, c.ns, meteringConfig), &meteringv1.MeteringConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.MeteringConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMeteringConfigs) UpdateStatus(meteringConfig *meteringv1.MeteringConfig) (*meteringv1.MeteringConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(meteringconfigsResource, "status", c.ns, meteringConfig), &meteringv1.MeteringConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.MeteringConfig), err
}

// Delete takes name of the meteringConfig and deletes it. Returns an error if one occurs.
func (c *FakeMeteringConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(meteringconfigsResource, c.ns, name), &meteringv1.MeteringConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMeteringConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(meteringconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &meteringv1.MeteringConfigList{})
	return err
}

// Patch applies the patch and returns the patched meteringConfig.
func (c *FakeMeteringConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *meteringv1.MeteringConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(meteringconfigsResource, c.ns, name, pt, data, subresources...), &meteringv1.MeteringConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.MeteringConfig), err
}
