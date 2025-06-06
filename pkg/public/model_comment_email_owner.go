/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public_sdk

import (
	"encoding/json"
)

// checks if the CommentEmailOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentEmailOwner{}

// CommentEmailOwner struct for CommentEmailOwner
type CommentEmailOwner struct {
	Avatar *string `json:"avatar,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Email *string `json:"email,omitempty"`
	Website *string `json:"website,omitempty"`
}

// NewCommentEmailOwner instantiates a new CommentEmailOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentEmailOwner() *CommentEmailOwner {
	this := CommentEmailOwner{}
	return &this
}

// NewCommentEmailOwnerWithDefaults instantiates a new CommentEmailOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentEmailOwnerWithDefaults() *CommentEmailOwner {
	this := CommentEmailOwner{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *CommentEmailOwner) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentEmailOwner) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *CommentEmailOwner) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *CommentEmailOwner) SetAvatar(v string) {
	o.Avatar = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CommentEmailOwner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentEmailOwner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CommentEmailOwner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CommentEmailOwner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CommentEmailOwner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentEmailOwner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CommentEmailOwner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CommentEmailOwner) SetEmail(v string) {
	o.Email = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *CommentEmailOwner) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentEmailOwner) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *CommentEmailOwner) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *CommentEmailOwner) SetWebsite(v string) {
	o.Website = &v
}

func (o CommentEmailOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentEmailOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

type NullableCommentEmailOwner struct {
	value *CommentEmailOwner
	isSet bool
}

func (v NullableCommentEmailOwner) Get() *CommentEmailOwner {
	return v.value
}

func (v *NullableCommentEmailOwner) Set(val *CommentEmailOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentEmailOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentEmailOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentEmailOwner(val *CommentEmailOwner) *NullableCommentEmailOwner {
	return &NullableCommentEmailOwner{value: val, isSet: true}
}

func (v NullableCommentEmailOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentEmailOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


