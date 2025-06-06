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

// checks if the UcUploadRequestFormData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UcUploadRequestFormData{}

// UcUploadRequestFormData struct for UcUploadRequestFormData
type UcUploadRequestFormData struct {
	All map[string]map[string]interface{} `json:"all,omitempty"`
	Empty *bool `json:"empty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UcUploadRequestFormData UcUploadRequestFormData

// NewUcUploadRequestFormData instantiates a new UcUploadRequestFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcUploadRequestFormData() *UcUploadRequestFormData {
	this := UcUploadRequestFormData{}
	return &this
}

// NewUcUploadRequestFormDataWithDefaults instantiates a new UcUploadRequestFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcUploadRequestFormDataWithDefaults() *UcUploadRequestFormData {
	this := UcUploadRequestFormData{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *UcUploadRequestFormData) GetAll() map[string]map[string]interface{} {
	if o == nil || IsNil(o.All) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcUploadRequestFormData) GetAllOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.All) {
		return map[string]map[string]interface{}{}, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *UcUploadRequestFormData) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given map[string]map[string]interface{} and assigns it to the All field.
func (o *UcUploadRequestFormData) SetAll(v map[string]map[string]interface{}) {
	o.All = v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *UcUploadRequestFormData) GetEmpty() bool {
	if o == nil || IsNil(o.Empty) {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcUploadRequestFormData) GetEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.Empty) {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *UcUploadRequestFormData) HasEmpty() bool {
	if o != nil && !IsNil(o.Empty) {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *UcUploadRequestFormData) SetEmpty(v bool) {
	o.Empty = &v
}

func (o UcUploadRequestFormData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UcUploadRequestFormData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	if !IsNil(o.Empty) {
		toSerialize["empty"] = o.Empty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UcUploadRequestFormData) UnmarshalJSON(data []byte) (err error) {
	varUcUploadRequestFormData := _UcUploadRequestFormData{}

	err = json.Unmarshal(data, &varUcUploadRequestFormData)

	if err != nil {
		return err
	}

	*o = UcUploadRequestFormData(varUcUploadRequestFormData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "all")
		delete(additionalProperties, "empty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUcUploadRequestFormData struct {
	value *UcUploadRequestFormData
	isSet bool
}

func (v NullableUcUploadRequestFormData) Get() *UcUploadRequestFormData {
	return v.value
}

func (v *NullableUcUploadRequestFormData) Set(val *UcUploadRequestFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableUcUploadRequestFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableUcUploadRequestFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcUploadRequestFormData(val *UcUploadRequestFormData) *NullableUcUploadRequestFormData {
	return &NullableUcUploadRequestFormData{value: val, isSet: true}
}

func (v NullableUcUploadRequestFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcUploadRequestFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


