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

// checks if the AuthProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthProvider{}

// AuthProvider struct for AuthProvider
type AuthProvider struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec AuthProviderSpec `json:"spec"`
}

type _AuthProvider AuthProvider

// NewAuthProvider instantiates a new AuthProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthProvider(apiVersion string, kind string, metadata Metadata, spec AuthProviderSpec) *AuthProvider {
	this := AuthProvider{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewAuthProviderWithDefaults instantiates a new AuthProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthProviderWithDefaults() *AuthProvider {
	this := AuthProvider{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *AuthProvider) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *AuthProvider) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *AuthProvider) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *AuthProvider) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AuthProvider) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AuthProvider) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *AuthProvider) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *AuthProvider) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *AuthProvider) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *AuthProvider) GetSpec() AuthProviderSpec {
	if o == nil {
		var ret AuthProviderSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *AuthProvider) GetSpecOk() (*AuthProviderSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *AuthProvider) SetSpec(v AuthProviderSpec) {
	o.Spec = v
}

func (o AuthProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	toSerialize["spec"] = o.Spec
	return toSerialize, nil
}

func (o *AuthProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
		"metadata",
		"spec",
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

	varAuthProvider := _AuthProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthProvider)

	if err != nil {
		return err
	}

	*o = AuthProvider(varAuthProvider)

	return err
}

type NullableAuthProvider struct {
	value *AuthProvider
	isSet bool
}

func (v NullableAuthProvider) Get() *AuthProvider {
	return v.value
}

func (v *NullableAuthProvider) Set(val *AuthProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthProvider(val *AuthProvider) *NullableAuthProvider {
	return &NullableAuthProvider{value: val, isSet: true}
}

func (v NullableAuthProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


