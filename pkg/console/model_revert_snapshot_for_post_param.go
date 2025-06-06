/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package console_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RevertSnapshotForPostParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevertSnapshotForPostParam{}

// RevertSnapshotForPostParam struct for RevertSnapshotForPostParam
type RevertSnapshotForPostParam struct {
	SnapshotName string `json:"snapshotName"`
}

type _RevertSnapshotForPostParam RevertSnapshotForPostParam

// NewRevertSnapshotForPostParam instantiates a new RevertSnapshotForPostParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevertSnapshotForPostParam(snapshotName string) *RevertSnapshotForPostParam {
	this := RevertSnapshotForPostParam{}
	this.SnapshotName = snapshotName
	return &this
}

// NewRevertSnapshotForPostParamWithDefaults instantiates a new RevertSnapshotForPostParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevertSnapshotForPostParamWithDefaults() *RevertSnapshotForPostParam {
	this := RevertSnapshotForPostParam{}
	return &this
}

// GetSnapshotName returns the SnapshotName field value
func (o *RevertSnapshotForPostParam) GetSnapshotName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotName
}

// GetSnapshotNameOk returns a tuple with the SnapshotName field value
// and a boolean to check if the value has been set.
func (o *RevertSnapshotForPostParam) GetSnapshotNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotName, true
}

// SetSnapshotName sets field value
func (o *RevertSnapshotForPostParam) SetSnapshotName(v string) {
	o.SnapshotName = v
}

func (o RevertSnapshotForPostParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevertSnapshotForPostParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snapshotName"] = o.SnapshotName
	return toSerialize, nil
}

func (o *RevertSnapshotForPostParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"snapshotName",
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

	varRevertSnapshotForPostParam := _RevertSnapshotForPostParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRevertSnapshotForPostParam)

	if err != nil {
		return err
	}

	*o = RevertSnapshotForPostParam(varRevertSnapshotForPostParam)

	return err
}

type NullableRevertSnapshotForPostParam struct {
	value *RevertSnapshotForPostParam
	isSet bool
}

func (v NullableRevertSnapshotForPostParam) Get() *RevertSnapshotForPostParam {
	return v.value
}

func (v *NullableRevertSnapshotForPostParam) Set(val *RevertSnapshotForPostParam) {
	v.value = val
	v.isSet = true
}

func (v NullableRevertSnapshotForPostParam) IsSet() bool {
	return v.isSet
}

func (v *NullableRevertSnapshotForPostParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevertSnapshotForPostParam(val *RevertSnapshotForPostParam) *NullableRevertSnapshotForPostParam {
	return &NullableRevertSnapshotForPostParam{value: val, isSet: true}
}

func (v NullableRevertSnapshotForPostParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevertSnapshotForPostParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


