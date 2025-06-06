/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package extension_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NotifierDescriptorSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifierDescriptorSpec{}

// NotifierDescriptorSpec struct for NotifierDescriptorSpec
type NotifierDescriptorSpec struct {
	Description *string `json:"description,omitempty"`
	DisplayName string `json:"displayName"`
	NotifierExtName string `json:"notifierExtName"`
	ReceiverSettingRef *NotifierSettingRef `json:"receiverSettingRef,omitempty"`
	SenderSettingRef *NotifierSettingRef `json:"senderSettingRef,omitempty"`
}

type _NotifierDescriptorSpec NotifierDescriptorSpec

// NewNotifierDescriptorSpec instantiates a new NotifierDescriptorSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifierDescriptorSpec(displayName string, notifierExtName string) *NotifierDescriptorSpec {
	this := NotifierDescriptorSpec{}
	this.DisplayName = displayName
	this.NotifierExtName = notifierExtName
	return &this
}

// NewNotifierDescriptorSpecWithDefaults instantiates a new NotifierDescriptorSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifierDescriptorSpecWithDefaults() *NotifierDescriptorSpec {
	this := NotifierDescriptorSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotifierDescriptorSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifierDescriptorSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotifierDescriptorSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotifierDescriptorSpec) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value
func (o *NotifierDescriptorSpec) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *NotifierDescriptorSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *NotifierDescriptorSpec) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetNotifierExtName returns the NotifierExtName field value
func (o *NotifierDescriptorSpec) GetNotifierExtName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifierExtName
}

// GetNotifierExtNameOk returns a tuple with the NotifierExtName field value
// and a boolean to check if the value has been set.
func (o *NotifierDescriptorSpec) GetNotifierExtNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifierExtName, true
}

// SetNotifierExtName sets field value
func (o *NotifierDescriptorSpec) SetNotifierExtName(v string) {
	o.NotifierExtName = v
}

// GetReceiverSettingRef returns the ReceiverSettingRef field value if set, zero value otherwise.
func (o *NotifierDescriptorSpec) GetReceiverSettingRef() NotifierSettingRef {
	if o == nil || IsNil(o.ReceiverSettingRef) {
		var ret NotifierSettingRef
		return ret
	}
	return *o.ReceiverSettingRef
}

// GetReceiverSettingRefOk returns a tuple with the ReceiverSettingRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifierDescriptorSpec) GetReceiverSettingRefOk() (*NotifierSettingRef, bool) {
	if o == nil || IsNil(o.ReceiverSettingRef) {
		return nil, false
	}
	return o.ReceiverSettingRef, true
}

// HasReceiverSettingRef returns a boolean if a field has been set.
func (o *NotifierDescriptorSpec) HasReceiverSettingRef() bool {
	if o != nil && !IsNil(o.ReceiverSettingRef) {
		return true
	}

	return false
}

// SetReceiverSettingRef gets a reference to the given NotifierSettingRef and assigns it to the ReceiverSettingRef field.
func (o *NotifierDescriptorSpec) SetReceiverSettingRef(v NotifierSettingRef) {
	o.ReceiverSettingRef = &v
}

// GetSenderSettingRef returns the SenderSettingRef field value if set, zero value otherwise.
func (o *NotifierDescriptorSpec) GetSenderSettingRef() NotifierSettingRef {
	if o == nil || IsNil(o.SenderSettingRef) {
		var ret NotifierSettingRef
		return ret
	}
	return *o.SenderSettingRef
}

// GetSenderSettingRefOk returns a tuple with the SenderSettingRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifierDescriptorSpec) GetSenderSettingRefOk() (*NotifierSettingRef, bool) {
	if o == nil || IsNil(o.SenderSettingRef) {
		return nil, false
	}
	return o.SenderSettingRef, true
}

// HasSenderSettingRef returns a boolean if a field has been set.
func (o *NotifierDescriptorSpec) HasSenderSettingRef() bool {
	if o != nil && !IsNil(o.SenderSettingRef) {
		return true
	}

	return false
}

// SetSenderSettingRef gets a reference to the given NotifierSettingRef and assigns it to the SenderSettingRef field.
func (o *NotifierDescriptorSpec) SetSenderSettingRef(v NotifierSettingRef) {
	o.SenderSettingRef = &v
}

func (o NotifierDescriptorSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifierDescriptorSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["notifierExtName"] = o.NotifierExtName
	if !IsNil(o.ReceiverSettingRef) {
		toSerialize["receiverSettingRef"] = o.ReceiverSettingRef
	}
	if !IsNil(o.SenderSettingRef) {
		toSerialize["senderSettingRef"] = o.SenderSettingRef
	}
	return toSerialize, nil
}

func (o *NotifierDescriptorSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"notifierExtName",
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

	varNotifierDescriptorSpec := _NotifierDescriptorSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotifierDescriptorSpec)

	if err != nil {
		return err
	}

	*o = NotifierDescriptorSpec(varNotifierDescriptorSpec)

	return err
}

type NullableNotifierDescriptorSpec struct {
	value *NotifierDescriptorSpec
	isSet bool
}

func (v NullableNotifierDescriptorSpec) Get() *NotifierDescriptorSpec {
	return v.value
}

func (v *NullableNotifierDescriptorSpec) Set(val *NotifierDescriptorSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifierDescriptorSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifierDescriptorSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifierDescriptorSpec(val *NotifierDescriptorSpec) *NullableNotifierDescriptorSpec {
	return &NullableNotifierDescriptorSpec{value: val, isSet: true}
}

func (v NullableNotifierDescriptorSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifierDescriptorSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


