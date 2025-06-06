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

// checks if the EmailConfigValidationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailConfigValidationRequest{}

// EmailConfigValidationRequest struct for EmailConfigValidationRequest
type EmailConfigValidationRequest struct {
	DisplayName *string `json:"displayName,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	Encryption *string `json:"encryption,omitempty"`
	Host *string `json:"host,omitempty"`
	Password *string `json:"password,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Sender *string `json:"sender,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewEmailConfigValidationRequest instantiates a new EmailConfigValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailConfigValidationRequest() *EmailConfigValidationRequest {
	this := EmailConfigValidationRequest{}
	return &this
}

// NewEmailConfigValidationRequestWithDefaults instantiates a new EmailConfigValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailConfigValidationRequestWithDefaults() *EmailConfigValidationRequest {
	this := EmailConfigValidationRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EmailConfigValidationRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *EmailConfigValidationRequest) SetEnable(v bool) {
	o.Enable = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetEncryption() string {
	if o == nil || IsNil(o.Encryption) {
		var ret string
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetEncryptionOk() (*string, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given string and assigns it to the Encryption field.
func (o *EmailConfigValidationRequest) SetEncryption(v string) {
	o.Encryption = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EmailConfigValidationRequest) SetHost(v string) {
	o.Host = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EmailConfigValidationRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *EmailConfigValidationRequest) SetPort(v int32) {
	o.Port = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *EmailConfigValidationRequest) SetSender(v string) {
	o.Sender = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EmailConfigValidationRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfigValidationRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EmailConfigValidationRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EmailConfigValidationRequest) SetUsername(v string) {
	o.Username = &v
}

func (o EmailConfigValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailConfigValidationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableEmailConfigValidationRequest struct {
	value *EmailConfigValidationRequest
	isSet bool
}

func (v NullableEmailConfigValidationRequest) Get() *EmailConfigValidationRequest {
	return v.value
}

func (v *NullableEmailConfigValidationRequest) Set(val *EmailConfigValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailConfigValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailConfigValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailConfigValidationRequest(val *EmailConfigValidationRequest) *NullableEmailConfigValidationRequest {
	return &NullableEmailConfigValidationRequest{value: val, isSet: true}
}

func (v NullableEmailConfigValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailConfigValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


