/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// NetworkBuilder contains the data and logic needed to build 'network' objects.
//
// Network configuration of a cluster.
type NetworkBuilder struct {
	machineCIDR *string
	podCIDR     *string
	serviceCIDR *string
}

// NewNetwork creates a new builder of 'network' objects.
func NewNetwork() *NetworkBuilder {
	return new(NetworkBuilder)
}

// MachineCIDR sets the value of the 'machine_CIDR' attribute
// to the given value.
//
//
func (b *NetworkBuilder) MachineCIDR(value string) *NetworkBuilder {
	b.machineCIDR = &value
	return b
}

// PodCIDR sets the value of the 'pod_CIDR' attribute
// to the given value.
//
//
func (b *NetworkBuilder) PodCIDR(value string) *NetworkBuilder {
	b.podCIDR = &value
	return b
}

// ServiceCIDR sets the value of the 'service_CIDR' attribute
// to the given value.
//
//
func (b *NetworkBuilder) ServiceCIDR(value string) *NetworkBuilder {
	b.serviceCIDR = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *NetworkBuilder) Copy(object *Network) *NetworkBuilder {
	if object == nil {
		return b
	}
	b.machineCIDR = object.machineCIDR
	b.podCIDR = object.podCIDR
	b.serviceCIDR = object.serviceCIDR
	return b
}

// Build creates a 'network' object using the configuration stored in the builder.
func (b *NetworkBuilder) Build() (object *Network, err error) {
	object = new(Network)
	if b.machineCIDR != nil {
		object.machineCIDR = b.machineCIDR
	}
	if b.podCIDR != nil {
		object.podCIDR = b.podCIDR
	}
	if b.serviceCIDR != nil {
		object.serviceCIDR = b.serviceCIDR
	}
	return
}
