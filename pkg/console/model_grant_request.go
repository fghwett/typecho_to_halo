/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package console_sdk

import (
	"encoding/json"
)

// checks if the GrantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantRequest{}

// GrantRequest struct for GrantRequest
type GrantRequest struct {
	Roles []string `json:"roles,omitempty"`
}

// NewGrantRequest instantiates a new GrantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantRequest() *GrantRequest {
	this := GrantRequest{}
	return &this
}

// NewGrantRequestWithDefaults instantiates a new GrantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantRequestWithDefaults() *GrantRequest {
	this := GrantRequest{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *GrantRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantRequest) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GrantRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GrantRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o GrantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableGrantRequest struct {
	value *GrantRequest
	isSet bool
}

func (v NullableGrantRequest) Get() *GrantRequest {
	return v.value
}

func (v *NullableGrantRequest) Set(val *GrantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantRequest(val *GrantRequest) *NullableGrantRequest {
	return &NullableGrantRequest{value: val, isSet: true}
}

func (v NullableGrantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


