/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package console_sdk

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ListedSnapshotSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListedSnapshotSpec{}

// ListedSnapshotSpec struct for ListedSnapshotSpec
type ListedSnapshotSpec struct {
	ModifyTime *time.Time `json:"modifyTime,omitempty"`
	Owner string `json:"owner"`
}

type _ListedSnapshotSpec ListedSnapshotSpec

// NewListedSnapshotSpec instantiates a new ListedSnapshotSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListedSnapshotSpec(owner string) *ListedSnapshotSpec {
	this := ListedSnapshotSpec{}
	this.Owner = owner
	return &this
}

// NewListedSnapshotSpecWithDefaults instantiates a new ListedSnapshotSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListedSnapshotSpecWithDefaults() *ListedSnapshotSpec {
	this := ListedSnapshotSpec{}
	return &this
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *ListedSnapshotSpec) GetModifyTime() time.Time {
	if o == nil || IsNil(o.ModifyTime) {
		var ret time.Time
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedSnapshotSpec) GetModifyTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifyTime) {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *ListedSnapshotSpec) HasModifyTime() bool {
	if o != nil && !IsNil(o.ModifyTime) {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given time.Time and assigns it to the ModifyTime field.
func (o *ListedSnapshotSpec) SetModifyTime(v time.Time) {
	o.ModifyTime = &v
}

// GetOwner returns the Owner field value
func (o *ListedSnapshotSpec) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *ListedSnapshotSpec) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *ListedSnapshotSpec) SetOwner(v string) {
	o.Owner = v
}

func (o ListedSnapshotSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListedSnapshotSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModifyTime) {
		toSerialize["modifyTime"] = o.ModifyTime
	}
	toSerialize["owner"] = o.Owner
	return toSerialize, nil
}

func (o *ListedSnapshotSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"owner",
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

	varListedSnapshotSpec := _ListedSnapshotSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListedSnapshotSpec)

	if err != nil {
		return err
	}

	*o = ListedSnapshotSpec(varListedSnapshotSpec)

	return err
}

type NullableListedSnapshotSpec struct {
	value *ListedSnapshotSpec
	isSet bool
}

func (v NullableListedSnapshotSpec) Get() *ListedSnapshotSpec {
	return v.value
}

func (v *NullableListedSnapshotSpec) Set(val *ListedSnapshotSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableListedSnapshotSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableListedSnapshotSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListedSnapshotSpec(val *ListedSnapshotSpec) *NullableListedSnapshotSpec {
	return &NullableListedSnapshotSpec{value: val, isSet: true}
}

func (v NullableListedSnapshotSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListedSnapshotSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


