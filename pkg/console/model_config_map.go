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

// checks if the ConfigMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigMap{}

// ConfigMap struct for ConfigMap
type ConfigMap struct {
	ApiVersion string `json:"apiVersion"`
	Data *map[string]string `json:"data,omitempty"`
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
}

type _ConfigMap ConfigMap

// NewConfigMap instantiates a new ConfigMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigMap(apiVersion string, kind string, metadata Metadata) *ConfigMap {
	this := ConfigMap{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewConfigMapWithDefaults instantiates a new ConfigMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMapWithDefaults() *ConfigMap {
	this := ConfigMap{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ConfigMap) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ConfigMap) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ConfigMap) GetData() map[string]string {
	if o == nil || IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ConfigMap) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *ConfigMap) SetData(v map[string]string) {
	o.Data = &v
}

// GetKind returns the Kind field value
func (o *ConfigMap) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConfigMap) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ConfigMap) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ConfigMap) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ConfigMap) SetMetadata(v Metadata) {
	o.Metadata = v
}

func (o ConfigMap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *ConfigMap) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
		"metadata",
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

	varConfigMap := _ConfigMap{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigMap)

	if err != nil {
		return err
	}

	*o = ConfigMap(varConfigMap)

	return err
}

type NullableConfigMap struct {
	value *ConfigMap
	isSet bool
}

func (v NullableConfigMap) Get() *ConfigMap {
	return v.value
}

func (v *NullableConfigMap) Set(val *ConfigMap) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigMap) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigMap(val *ConfigMap) *NullableConfigMap {
	return &NullableConfigMap{value: val, isSet: true}
}

func (v NullableConfigMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


