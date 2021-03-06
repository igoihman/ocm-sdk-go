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

// OpenIDIdentityProviderBuilder contains the data and logic needed to build 'open_ID_identity_provider' objects.
//
// Details for `openid` identity providers.
type OpenIDIdentityProviderBuilder struct {
	ca                       *string
	urls                     *OpenIDURLsBuilder
	claims                   *OpenIDClaimsBuilder
	clientID                 *string
	clientSecret             *string
	extraAuthorizeParameters map[string]string
	extraScopes              []string
}

// NewOpenIDIdentityProvider creates a new builder of 'open_ID_identity_provider' objects.
func NewOpenIDIdentityProvider() *OpenIDIdentityProviderBuilder {
	return new(OpenIDIdentityProviderBuilder)
}

// CA sets the value of the 'CA' attribute
// to the given value.
//
//
func (b *OpenIDIdentityProviderBuilder) CA(value string) *OpenIDIdentityProviderBuilder {
	b.ca = &value
	return b
}

// URLS sets the value of the 'URLS' attribute
// to the given value.
//
// _OpenID_ identity provider URLs.
func (b *OpenIDIdentityProviderBuilder) URLS(value *OpenIDURLsBuilder) *OpenIDIdentityProviderBuilder {
	b.urls = value
	return b
}

// Claims sets the value of the 'claims' attribute
// to the given value.
//
// _OpenID_ identity provider claims.
func (b *OpenIDIdentityProviderBuilder) Claims(value *OpenIDClaimsBuilder) *OpenIDIdentityProviderBuilder {
	b.claims = value
	return b
}

// ClientID sets the value of the 'client_ID' attribute
// to the given value.
//
//
func (b *OpenIDIdentityProviderBuilder) ClientID(value string) *OpenIDIdentityProviderBuilder {
	b.clientID = &value
	return b
}

// ClientSecret sets the value of the 'client_secret' attribute
// to the given value.
//
//
func (b *OpenIDIdentityProviderBuilder) ClientSecret(value string) *OpenIDIdentityProviderBuilder {
	b.clientSecret = &value
	return b
}

// ExtraAuthorizeParameters sets the value of the 'extra_authorize_parameters' attribute
// to the given value.
//
//
func (b *OpenIDIdentityProviderBuilder) ExtraAuthorizeParameters(value map[string]string) *OpenIDIdentityProviderBuilder {
	b.extraAuthorizeParameters = value
	return b
}

// ExtraScopes sets the value of the 'extra_scopes' attribute
// to the given values.
//
//
func (b *OpenIDIdentityProviderBuilder) ExtraScopes(values ...string) *OpenIDIdentityProviderBuilder {
	b.extraScopes = make([]string, len(values))
	copy(b.extraScopes, values)
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *OpenIDIdentityProviderBuilder) Copy(object *OpenIDIdentityProvider) *OpenIDIdentityProviderBuilder {
	if object == nil {
		return b
	}
	b.ca = object.ca
	if object.urls != nil {
		b.urls = NewOpenIDURLs().Copy(object.urls)
	} else {
		b.urls = nil
	}
	if object.claims != nil {
		b.claims = NewOpenIDClaims().Copy(object.claims)
	} else {
		b.claims = nil
	}
	b.clientID = object.clientID
	b.clientSecret = object.clientSecret
	if len(object.extraAuthorizeParameters) > 0 {
		b.extraAuthorizeParameters = make(map[string]string)
		for key, value := range object.extraAuthorizeParameters {
			b.extraAuthorizeParameters[key] = value
		}
	} else {
		b.extraAuthorizeParameters = nil
	}
	if len(object.extraScopes) > 0 {
		b.extraScopes = make([]string, len(object.extraScopes))
		copy(b.extraScopes, object.extraScopes)
	} else {
		b.extraScopes = nil
	}
	return b
}

// Build creates a 'open_ID_identity_provider' object using the configuration stored in the builder.
func (b *OpenIDIdentityProviderBuilder) Build() (object *OpenIDIdentityProvider, err error) {
	object = new(OpenIDIdentityProvider)
	if b.ca != nil {
		object.ca = b.ca
	}
	if b.urls != nil {
		object.urls, err = b.urls.Build()
		if err != nil {
			return
		}
	}
	if b.claims != nil {
		object.claims, err = b.claims.Build()
		if err != nil {
			return
		}
	}
	if b.clientID != nil {
		object.clientID = b.clientID
	}
	if b.clientSecret != nil {
		object.clientSecret = b.clientSecret
	}
	if b.extraAuthorizeParameters != nil {
		object.extraAuthorizeParameters = make(map[string]string)
		for key, value := range b.extraAuthorizeParameters {
			object.extraAuthorizeParameters[key] = value
		}
	}
	if b.extraScopes != nil {
		object.extraScopes = make([]string, len(b.extraScopes))
		copy(object.extraScopes, b.extraScopes)
	}
	return
}
