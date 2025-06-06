/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package extension_sdk

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ReplySpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplySpec{}

// ReplySpec struct for ReplySpec
type ReplySpec struct {
	AllowNotification bool `json:"allowNotification"`
	Approved bool `json:"approved"`
	ApprovedTime *time.Time `json:"approvedTime,omitempty"`
	CommentName string `json:"commentName"`
	Content string `json:"content"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	Hidden bool `json:"hidden"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Owner CommentOwner `json:"owner"`
	Priority int32 `json:"priority"`
	QuoteReply *string `json:"quoteReply,omitempty"`
	Raw string `json:"raw"`
	Top bool `json:"top"`
	UserAgent *string `json:"userAgent,omitempty"`
}

type _ReplySpec ReplySpec

// NewReplySpec instantiates a new ReplySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplySpec(allowNotification bool, approved bool, commentName string, content string, hidden bool, owner CommentOwner, priority int32, raw string, top bool) *ReplySpec {
	this := ReplySpec{}
	this.AllowNotification = allowNotification
	this.Approved = approved
	this.CommentName = commentName
	this.Content = content
	this.Hidden = hidden
	this.Owner = owner
	this.Priority = priority
	this.Raw = raw
	this.Top = top
	return &this
}

// NewReplySpecWithDefaults instantiates a new ReplySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplySpecWithDefaults() *ReplySpec {
	this := ReplySpec{}
	var allowNotification bool = true
	this.AllowNotification = allowNotification
	var approved bool = false
	this.Approved = approved
	var hidden bool = false
	this.Hidden = hidden
	var priority int32 = 0
	this.Priority = priority
	var top bool = false
	this.Top = top
	return &this
}

// GetAllowNotification returns the AllowNotification field value
func (o *ReplySpec) GetAllowNotification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowNotification
}

// GetAllowNotificationOk returns a tuple with the AllowNotification field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetAllowNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowNotification, true
}

// SetAllowNotification sets field value
func (o *ReplySpec) SetAllowNotification(v bool) {
	o.AllowNotification = v
}

// GetApproved returns the Approved field value
func (o *ReplySpec) GetApproved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetApprovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approved, true
}

// SetApproved sets field value
func (o *ReplySpec) SetApproved(v bool) {
	o.Approved = v
}

// GetApprovedTime returns the ApprovedTime field value if set, zero value otherwise.
func (o *ReplySpec) GetApprovedTime() time.Time {
	if o == nil || IsNil(o.ApprovedTime) {
		var ret time.Time
		return ret
	}
	return *o.ApprovedTime
}

// GetApprovedTimeOk returns a tuple with the ApprovedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetApprovedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ApprovedTime) {
		return nil, false
	}
	return o.ApprovedTime, true
}

// HasApprovedTime returns a boolean if a field has been set.
func (o *ReplySpec) HasApprovedTime() bool {
	if o != nil && !IsNil(o.ApprovedTime) {
		return true
	}

	return false
}

// SetApprovedTime gets a reference to the given time.Time and assigns it to the ApprovedTime field.
func (o *ReplySpec) SetApprovedTime(v time.Time) {
	o.ApprovedTime = &v
}

// GetCommentName returns the CommentName field value
func (o *ReplySpec) GetCommentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommentName
}

// GetCommentNameOk returns a tuple with the CommentName field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetCommentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentName, true
}

// SetCommentName sets field value
func (o *ReplySpec) SetCommentName(v string) {
	o.CommentName = v
}

// GetContent returns the Content field value
func (o *ReplySpec) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ReplySpec) SetContent(v string) {
	o.Content = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *ReplySpec) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ReplySpec) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *ReplySpec) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetHidden returns the Hidden field value
func (o *ReplySpec) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *ReplySpec) SetHidden(v bool) {
	o.Hidden = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ReplySpec) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ReplySpec) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ReplySpec) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetOwner returns the Owner field value
func (o *ReplySpec) GetOwner() CommentOwner {
	if o == nil {
		var ret CommentOwner
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetOwnerOk() (*CommentOwner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *ReplySpec) SetOwner(v CommentOwner) {
	o.Owner = v
}

// GetPriority returns the Priority field value
func (o *ReplySpec) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ReplySpec) SetPriority(v int32) {
	o.Priority = v
}

// GetQuoteReply returns the QuoteReply field value if set, zero value otherwise.
func (o *ReplySpec) GetQuoteReply() string {
	if o == nil || IsNil(o.QuoteReply) {
		var ret string
		return ret
	}
	return *o.QuoteReply
}

// GetQuoteReplyOk returns a tuple with the QuoteReply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetQuoteReplyOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteReply) {
		return nil, false
	}
	return o.QuoteReply, true
}

// HasQuoteReply returns a boolean if a field has been set.
func (o *ReplySpec) HasQuoteReply() bool {
	if o != nil && !IsNil(o.QuoteReply) {
		return true
	}

	return false
}

// SetQuoteReply gets a reference to the given string and assigns it to the QuoteReply field.
func (o *ReplySpec) SetQuoteReply(v string) {
	o.QuoteReply = &v
}

// GetRaw returns the Raw field value
func (o *ReplySpec) GetRaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raw, true
}

// SetRaw sets field value
func (o *ReplySpec) SetRaw(v string) {
	o.Raw = v
}

// GetTop returns the Top field value
func (o *ReplySpec) GetTop() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Top
}

// GetTopOk returns a tuple with the Top field value
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetTopOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Top, true
}

// SetTop sets field value
func (o *ReplySpec) SetTop(v bool) {
	o.Top = v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *ReplySpec) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplySpec) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *ReplySpec) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *ReplySpec) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o ReplySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowNotification"] = o.AllowNotification
	toSerialize["approved"] = o.Approved
	if !IsNil(o.ApprovedTime) {
		toSerialize["approvedTime"] = o.ApprovedTime
	}
	toSerialize["commentName"] = o.CommentName
	toSerialize["content"] = o.Content
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	toSerialize["hidden"] = o.Hidden
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	toSerialize["owner"] = o.Owner
	toSerialize["priority"] = o.Priority
	if !IsNil(o.QuoteReply) {
		toSerialize["quoteReply"] = o.QuoteReply
	}
	toSerialize["raw"] = o.Raw
	toSerialize["top"] = o.Top
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	return toSerialize, nil
}

func (o *ReplySpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowNotification",
		"approved",
		"commentName",
		"content",
		"hidden",
		"owner",
		"priority",
		"raw",
		"top",
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

	varReplySpec := _ReplySpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplySpec)

	if err != nil {
		return err
	}

	*o = ReplySpec(varReplySpec)

	return err
}

type NullableReplySpec struct {
	value *ReplySpec
	isSet bool
}

func (v NullableReplySpec) Get() *ReplySpec {
	return v.value
}

func (v *NullableReplySpec) Set(val *ReplySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReplySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReplySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplySpec(val *ReplySpec) *NullableReplySpec {
	return &NullableReplySpec{value: val, isSet: true}
}

func (v NullableReplySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


