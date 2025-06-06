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

// checks if the ReasonSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReasonSubject{}

// ReasonSubject struct for ReasonSubject
type ReasonSubject struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Name string `json:"name"`
	Title string `json:"title"`
	Url *string `json:"url,omitempty"`
}

type _ReasonSubject ReasonSubject

// NewReasonSubject instantiates a new ReasonSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReasonSubject(apiVersion string, kind string, name string, title string) *ReasonSubject {
	this := ReasonSubject{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Name = name
	this.Title = title
	return &this
}

// NewReasonSubjectWithDefaults instantiates a new ReasonSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonSubjectWithDefaults() *ReasonSubject {
	this := ReasonSubject{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ReasonSubject) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ReasonSubject) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ReasonSubject) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *ReasonSubject) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ReasonSubject) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ReasonSubject) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *ReasonSubject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReasonSubject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReasonSubject) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *ReasonSubject) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ReasonSubject) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ReasonSubject) SetTitle(v string) {
	o.Title = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ReasonSubject) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonSubject) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ReasonSubject) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ReasonSubject) SetUrl(v string) {
	o.Url = &v
}

func (o ReasonSubject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReasonSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

func (o *ReasonSubject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
		"name",
		"title",
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

	varReasonSubject := _ReasonSubject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReasonSubject)

	if err != nil {
		return err
	}

	*o = ReasonSubject(varReasonSubject)

	return err
}

type NullableReasonSubject struct {
	value *ReasonSubject
	isSet bool
}

func (v NullableReasonSubject) Get() *ReasonSubject {
	return v.value
}

func (v *NullableReasonSubject) Set(val *ReasonSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonSubject(val *ReasonSubject) *NullableReasonSubject {
	return &NullableReasonSubject{value: val, isSet: true}
}

func (v NullableReasonSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReasonSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


