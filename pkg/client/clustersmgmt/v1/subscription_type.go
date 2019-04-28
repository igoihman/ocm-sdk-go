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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// SubscriptionKind is the name of the type used to represent objects
// of type 'subscription'.
const SubscriptionKind = "Subscription"

// SubscriptionLinkKind is the name of the type used to represent links
// to objects of type 'subscription'.
const SubscriptionLinkKind = "SubscriptionLink"

// SubscriptionNilKind is the name of the type used to nil references
// to objects of type 'subscription'.
const SubscriptionNilKind = "SubscriptionNil"

// Subscription represents the values of the 'subscription' type.
//
// Definition of a subscription.
type Subscription struct {
	id   *string
	href *string
	link bool
}

// Kind returns the name of the type of the object.
func (o *Subscription) Kind() string {
	if o == nil {
		return SubscriptionNilKind
	}
	if o.link {
		return SubscriptionLinkKind
	}
	return SubscriptionKind
}

// ID returns the identifier of the object.
func (o *Subscription) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Subscription) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Subscription) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Subscription) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Subscription) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// SubscriptionListKind is the name of the type used to represent list of
// objects of type 'subscription'.
const SubscriptionListKind = "SubscriptionList"

// SubscriptionListLinkKind is the name of the type used to represent links
// to list of objects of type 'subscription'.
const SubscriptionListLinkKind = "SubscriptionListLink"

// SubscriptionNilKind is the name of the type used to nil lists of
// objects of type 'subscription'.
const SubscriptionListNilKind = "SubscriptionListNil"

// SubscriptionList is a list of values of the 'subscription' type.
type SubscriptionList struct {
	href  *string
	link  bool
	items []*Subscription
}

// Kind returns the name of the type of the object.
func (l *SubscriptionList) Kind() string {
	if l == nil {
		return SubscriptionListNilKind
	}
	if l.link {
		return SubscriptionListLinkKind
	}
	return SubscriptionListKind
}

// Link returns true iif this is a link.
func (l *SubscriptionList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *SubscriptionList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *SubscriptionList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *SubscriptionList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *SubscriptionList) Slice() []*Subscription {
	var slice []*Subscription
	if l == nil {
		slice = make([]*Subscription, 0)
	} else {
		slice = make([]*Subscription, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *SubscriptionList) Each(f func(item *Subscription) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *SubscriptionList) Range(f func(index int, item *Subscription) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
