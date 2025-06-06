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

// checks if the DeviceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceStatus{}

// DeviceStatus struct for DeviceStatus
type DeviceStatus struct {
	Browser *string `json:"browser,omitempty"`
	Os *string `json:"os,omitempty"`
}

// NewDeviceStatus instantiates a new DeviceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceStatus() *DeviceStatus {
	this := DeviceStatus{}
	return &this
}

// NewDeviceStatusWithDefaults instantiates a new DeviceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceStatusWithDefaults() *DeviceStatus {
	this := DeviceStatus{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *DeviceStatus) GetBrowser() string {
	if o == nil || IsNil(o.Browser) {
		var ret string
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStatus) GetBrowserOk() (*string, bool) {
	if o == nil || IsNil(o.Browser) {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *DeviceStatus) HasBrowser() bool {
	if o != nil && !IsNil(o.Browser) {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given string and assigns it to the Browser field.
func (o *DeviceStatus) SetBrowser(v string) {
	o.Browser = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *DeviceStatus) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStatus) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *DeviceStatus) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *DeviceStatus) SetOs(v string) {
	o.Os = &v
}

func (o DeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Browser) {
		toSerialize["browser"] = o.Browser
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	return toSerialize, nil
}

type NullableDeviceStatus struct {
	value *DeviceStatus
	isSet bool
}

func (v NullableDeviceStatus) Get() *DeviceStatus {
	return v.value
}

func (v *NullableDeviceStatus) Set(val *DeviceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceStatus(val *DeviceStatus) *NullableDeviceStatus {
	return &NullableDeviceStatus{value: val, isSet: true}
}

func (v NullableDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


