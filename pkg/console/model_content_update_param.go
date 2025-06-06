/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package console_sdk

import (
	"encoding/json"
)

// checks if the ContentUpdateParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentUpdateParam{}

// ContentUpdateParam struct for ContentUpdateParam
type ContentUpdateParam struct {
	Content *string `json:"content,omitempty"`
	Raw *string `json:"raw,omitempty"`
	RawType *string `json:"rawType,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewContentUpdateParam instantiates a new ContentUpdateParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentUpdateParam() *ContentUpdateParam {
	this := ContentUpdateParam{}
	return &this
}

// NewContentUpdateParamWithDefaults instantiates a new ContentUpdateParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentUpdateParamWithDefaults() *ContentUpdateParam {
	this := ContentUpdateParam{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ContentUpdateParam) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentUpdateParam) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ContentUpdateParam) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ContentUpdateParam) SetContent(v string) {
	o.Content = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *ContentUpdateParam) GetRaw() string {
	if o == nil || IsNil(o.Raw) {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentUpdateParam) GetRawOk() (*string, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *ContentUpdateParam) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *ContentUpdateParam) SetRaw(v string) {
	o.Raw = &v
}

// GetRawType returns the RawType field value if set, zero value otherwise.
func (o *ContentUpdateParam) GetRawType() string {
	if o == nil || IsNil(o.RawType) {
		var ret string
		return ret
	}
	return *o.RawType
}

// GetRawTypeOk returns a tuple with the RawType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentUpdateParam) GetRawTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RawType) {
		return nil, false
	}
	return o.RawType, true
}

// HasRawType returns a boolean if a field has been set.
func (o *ContentUpdateParam) HasRawType() bool {
	if o != nil && !IsNil(o.RawType) {
		return true
	}

	return false
}

// SetRawType gets a reference to the given string and assigns it to the RawType field.
func (o *ContentUpdateParam) SetRawType(v string) {
	o.RawType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ContentUpdateParam) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentUpdateParam) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ContentUpdateParam) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ContentUpdateParam) SetVersion(v int64) {
	o.Version = &v
}

func (o ContentUpdateParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentUpdateParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	if !IsNil(o.RawType) {
		toSerialize["rawType"] = o.RawType
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableContentUpdateParam struct {
	value *ContentUpdateParam
	isSet bool
}

func (v NullableContentUpdateParam) Get() *ContentUpdateParam {
	return v.value
}

func (v *NullableContentUpdateParam) Set(val *ContentUpdateParam) {
	v.value = val
	v.isSet = true
}

func (v NullableContentUpdateParam) IsSet() bool {
	return v.isSet
}

func (v *NullableContentUpdateParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentUpdateParam(val *ContentUpdateParam) *NullableContentUpdateParam {
	return &NullableContentUpdateParam{value: val, isSet: true}
}

func (v NullableContentUpdateParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentUpdateParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


