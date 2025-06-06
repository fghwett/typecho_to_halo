/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package user_sdk

import (
	"encoding/json"
)

// checks if the MarkSpecifiedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkSpecifiedRequest{}

// MarkSpecifiedRequest struct for MarkSpecifiedRequest
type MarkSpecifiedRequest struct {
	Names []string `json:"names,omitempty"`
}

// NewMarkSpecifiedRequest instantiates a new MarkSpecifiedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkSpecifiedRequest() *MarkSpecifiedRequest {
	this := MarkSpecifiedRequest{}
	return &this
}

// NewMarkSpecifiedRequestWithDefaults instantiates a new MarkSpecifiedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkSpecifiedRequestWithDefaults() *MarkSpecifiedRequest {
	this := MarkSpecifiedRequest{}
	return &this
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *MarkSpecifiedRequest) GetNames() []string {
	if o == nil || IsNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkSpecifiedRequest) GetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *MarkSpecifiedRequest) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *MarkSpecifiedRequest) SetNames(v []string) {
	o.Names = v
}

func (o MarkSpecifiedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkSpecifiedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	return toSerialize, nil
}

type NullableMarkSpecifiedRequest struct {
	value *MarkSpecifiedRequest
	isSet bool
}

func (v NullableMarkSpecifiedRequest) Get() *MarkSpecifiedRequest {
	return v.value
}

func (v *NullableMarkSpecifiedRequest) Set(val *MarkSpecifiedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkSpecifiedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkSpecifiedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkSpecifiedRequest(val *MarkSpecifiedRequest) *NullableMarkSpecifiedRequest {
	return &NullableMarkSpecifiedRequest{value: val, isSet: true}
}

func (v NullableMarkSpecifiedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkSpecifiedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


