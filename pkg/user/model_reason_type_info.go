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

// checks if the ReasonTypeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReasonTypeInfo{}

// ReasonTypeInfo struct for ReasonTypeInfo
type ReasonTypeInfo struct {
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Name *string `json:"name,omitempty"`
	UiPermissions []string `json:"uiPermissions,omitempty"`
}

// NewReasonTypeInfo instantiates a new ReasonTypeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReasonTypeInfo() *ReasonTypeInfo {
	this := ReasonTypeInfo{}
	return &this
}

// NewReasonTypeInfoWithDefaults instantiates a new ReasonTypeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonTypeInfoWithDefaults() *ReasonTypeInfo {
	this := ReasonTypeInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReasonTypeInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonTypeInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReasonTypeInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReasonTypeInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ReasonTypeInfo) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonTypeInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ReasonTypeInfo) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ReasonTypeInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReasonTypeInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonTypeInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReasonTypeInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReasonTypeInfo) SetName(v string) {
	o.Name = &v
}

// GetUiPermissions returns the UiPermissions field value if set, zero value otherwise.
func (o *ReasonTypeInfo) GetUiPermissions() []string {
	if o == nil || IsNil(o.UiPermissions) {
		var ret []string
		return ret
	}
	return o.UiPermissions
}

// GetUiPermissionsOk returns a tuple with the UiPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonTypeInfo) GetUiPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.UiPermissions) {
		return nil, false
	}
	return o.UiPermissions, true
}

// HasUiPermissions returns a boolean if a field has been set.
func (o *ReasonTypeInfo) HasUiPermissions() bool {
	if o != nil && !IsNil(o.UiPermissions) {
		return true
	}

	return false
}

// SetUiPermissions gets a reference to the given []string and assigns it to the UiPermissions field.
func (o *ReasonTypeInfo) SetUiPermissions(v []string) {
	o.UiPermissions = v
}

func (o ReasonTypeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReasonTypeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UiPermissions) {
		toSerialize["uiPermissions"] = o.UiPermissions
	}
	return toSerialize, nil
}

type NullableReasonTypeInfo struct {
	value *ReasonTypeInfo
	isSet bool
}

func (v NullableReasonTypeInfo) Get() *ReasonTypeInfo {
	return v.value
}

func (v *NullableReasonTypeInfo) Set(val *ReasonTypeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonTypeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonTypeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonTypeInfo(val *ReasonTypeInfo) *NullableReasonTypeInfo {
	return &NullableReasonTypeInfo{value: val, isSet: true}
}

func (v NullableReasonTypeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReasonTypeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


