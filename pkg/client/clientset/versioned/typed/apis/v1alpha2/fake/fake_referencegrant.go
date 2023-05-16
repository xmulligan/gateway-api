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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// FakeReferenceGrants implements ReferenceGrantInterface
type FakeReferenceGrants struct {
	Fake *FakeGatewayV1alpha2
	ns   string
}

var referencegrantsResource = v1alpha2.SchemeGroupVersion.WithResource("referencegrants")

var referencegrantsKind = v1alpha2.SchemeGroupVersion.WithKind("ReferenceGrant")

// Get takes name of the referenceGrant, and returns the corresponding referenceGrant object, and an error if there is any.
func (c *FakeReferenceGrants) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ReferenceGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(referencegrantsResource, c.ns, name), &v1alpha2.ReferenceGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReferenceGrant), err
}

// List takes label and field selectors, and returns the list of ReferenceGrants that match those selectors.
func (c *FakeReferenceGrants) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ReferenceGrantList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(referencegrantsResource, referencegrantsKind, c.ns, opts), &v1alpha2.ReferenceGrantList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ReferenceGrantList{ListMeta: obj.(*v1alpha2.ReferenceGrantList).ListMeta}
	for _, item := range obj.(*v1alpha2.ReferenceGrantList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested referenceGrants.
func (c *FakeReferenceGrants) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(referencegrantsResource, c.ns, opts))

}

// Create takes the representation of a referenceGrant and creates it.  Returns the server's representation of the referenceGrant, and an error, if there is any.
func (c *FakeReferenceGrants) Create(ctx context.Context, referenceGrant *v1alpha2.ReferenceGrant, opts v1.CreateOptions) (result *v1alpha2.ReferenceGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(referencegrantsResource, c.ns, referenceGrant), &v1alpha2.ReferenceGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReferenceGrant), err
}

// Update takes the representation of a referenceGrant and updates it. Returns the server's representation of the referenceGrant, and an error, if there is any.
func (c *FakeReferenceGrants) Update(ctx context.Context, referenceGrant *v1alpha2.ReferenceGrant, opts v1.UpdateOptions) (result *v1alpha2.ReferenceGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(referencegrantsResource, c.ns, referenceGrant), &v1alpha2.ReferenceGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReferenceGrant), err
}

// Delete takes name of the referenceGrant and deletes it. Returns an error if one occurs.
func (c *FakeReferenceGrants) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(referencegrantsResource, c.ns, name, opts), &v1alpha2.ReferenceGrant{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReferenceGrants) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(referencegrantsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ReferenceGrantList{})
	return err
}

// Patch applies the patch and returns the patched referenceGrant.
func (c *FakeReferenceGrants) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ReferenceGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(referencegrantsResource, c.ns, name, pt, data, subresources...), &v1alpha2.ReferenceGrant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReferenceGrant), err
}
