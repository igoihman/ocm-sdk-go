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

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

import (
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// selfAccessReviewRequestData is the data structure used internally to marshal and unmarshal
// objects of type 'self_access_review_request'.
type selfAccessReviewRequestData struct {
	Action         *string "json:\"action,omitempty\""
	ClusterID      *string "json:\"cluster_id,omitempty\""
	ClusterUUID    *string "json:\"cluster_uuid,omitempty\""
	OrganizationID *string "json:\"organization_id,omitempty\""
	ResourceType   *string "json:\"resource_type,omitempty\""
	SubscriptionID *string "json:\"subscription_id,omitempty\""
}

// MarshalSelfAccessReviewRequest writes a value of the 'self_access_review_request' to the given target,
// which can be a writer or a JSON encoder.
func MarshalSelfAccessReviewRequest(object *SelfAccessReviewRequest, target interface{}) error {
	encoder, err := helpers.NewEncoder(target)
	if err != nil {
		return err
	}
	data, err := object.wrap()
	if err != nil {
		return err
	}
	return encoder.Encode(data)
}

// wrap is the method used internally to convert a value of the 'self_access_review_request'
// value to a JSON document.
func (o *SelfAccessReviewRequest) wrap() (data *selfAccessReviewRequestData, err error) {
	if o == nil {
		return
	}
	data = new(selfAccessReviewRequestData)
	data.Action = o.action
	data.ClusterID = o.clusterID
	data.ClusterUUID = o.clusterUUID
	data.OrganizationID = o.organizationID
	data.ResourceType = o.resourceType
	data.SubscriptionID = o.subscriptionID
	return
}

// UnmarshalSelfAccessReviewRequest reads a value of the 'self_access_review_request' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalSelfAccessReviewRequest(source interface{}) (object *SelfAccessReviewRequest, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(selfAccessReviewRequestData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'self_access_review_request' type.
func (d *selfAccessReviewRequestData) unwrap() (object *SelfAccessReviewRequest, err error) {
	if d == nil {
		return
	}
	object = new(SelfAccessReviewRequest)
	object.action = d.Action
	object.clusterID = d.ClusterID
	object.clusterUUID = d.ClusterUUID
	object.organizationID = d.OrganizationID
	object.resourceType = d.ResourceType
	object.subscriptionID = d.SubscriptionID
	return
}
