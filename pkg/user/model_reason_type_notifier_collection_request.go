/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package user_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReasonTypeNotifierCollectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReasonTypeNotifierCollectionRequest{}

// ReasonTypeNotifierCollectionRequest struct for ReasonTypeNotifierCollectionRequest
type ReasonTypeNotifierCollectionRequest struct {
	ReasonTypeNotifiers []ReasonTypeNotifierRequest `json:"reasonTypeNotifiers"`
}

type _ReasonTypeNotifierCollectionRequest ReasonTypeNotifierCollectionRequest

// NewReasonTypeNotifierCollectionRequest instantiates a new ReasonTypeNotifierCollectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReasonTypeNotifierCollectionRequest(reasonTypeNotifiers []ReasonTypeNotifierRequest) *ReasonTypeNotifierCollectionRequest {
	this := ReasonTypeNotifierCollectionRequest{}
	this.ReasonTypeNotifiers = reasonTypeNotifiers
	return &this
}

// NewReasonTypeNotifierCollectionRequestWithDefaults instantiates a new ReasonTypeNotifierCollectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonTypeNotifierCollectionRequestWithDefaults() *ReasonTypeNotifierCollectionRequest {
	this := ReasonTypeNotifierCollectionRequest{}
	return &this
}

// GetReasonTypeNotifiers returns the ReasonTypeNotifiers field value
func (o *ReasonTypeNotifierCollectionRequest) GetReasonTypeNotifiers() []ReasonTypeNotifierRequest {
	if o == nil {
		var ret []ReasonTypeNotifierRequest
		return ret
	}

	return o.ReasonTypeNotifiers
}

// GetReasonTypeNotifiersOk returns a tuple with the ReasonTypeNotifiers field value
// and a boolean to check if the value has been set.
func (o *ReasonTypeNotifierCollectionRequest) GetReasonTypeNotifiersOk() ([]ReasonTypeNotifierRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasonTypeNotifiers, true
}

// SetReasonTypeNotifiers sets field value
func (o *ReasonTypeNotifierCollectionRequest) SetReasonTypeNotifiers(v []ReasonTypeNotifierRequest) {
	o.ReasonTypeNotifiers = v
}

func (o ReasonTypeNotifierCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReasonTypeNotifierCollectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reasonTypeNotifiers"] = o.ReasonTypeNotifiers
	return toSerialize, nil
}

func (o *ReasonTypeNotifierCollectionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reasonTypeNotifiers",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReasonTypeNotifierCollectionRequest := _ReasonTypeNotifierCollectionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReasonTypeNotifierCollectionRequest)

	if err != nil {
		return err
	}

	*o = ReasonTypeNotifierCollectionRequest(varReasonTypeNotifierCollectionRequest)

	return err
}

type NullableReasonTypeNotifierCollectionRequest struct {
	value *ReasonTypeNotifierCollectionRequest
	isSet bool
}

func (v NullableReasonTypeNotifierCollectionRequest) Get() *ReasonTypeNotifierCollectionRequest {
	return v.value
}

func (v *NullableReasonTypeNotifierCollectionRequest) Set(val *ReasonTypeNotifierCollectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonTypeNotifierCollectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonTypeNotifierCollectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonTypeNotifierCollectionRequest(val *ReasonTypeNotifierCollectionRequest) *NullableReasonTypeNotifierCollectionRequest {
	return &NullableReasonTypeNotifierCollectionRequest{value: val, isSet: true}
}

func (v NullableReasonTypeNotifierCollectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReasonTypeNotifierCollectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


