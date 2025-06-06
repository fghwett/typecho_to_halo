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

// checks if the UserPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPermission{}

// UserPermission struct for UserPermission
type UserPermission struct {
	Permissions []Role `json:"permissions"`
	Roles []Role `json:"roles"`
	UiPermissions []string `json:"uiPermissions"`
}

type _UserPermission UserPermission

// NewUserPermission instantiates a new UserPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPermission(permissions []Role, roles []Role, uiPermissions []string) *UserPermission {
	this := UserPermission{}
	this.Permissions = permissions
	this.Roles = roles
	this.UiPermissions = uiPermissions
	return &this
}

// NewUserPermissionWithDefaults instantiates a new UserPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPermissionWithDefaults() *UserPermission {
	this := UserPermission{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *UserPermission) GetPermissions() []Role {
	if o == nil {
		var ret []Role
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *UserPermission) GetPermissionsOk() ([]Role, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *UserPermission) SetPermissions(v []Role) {
	o.Permissions = v
}

// GetRoles returns the Roles field value
func (o *UserPermission) GetRoles() []Role {
	if o == nil {
		var ret []Role
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserPermission) GetRolesOk() ([]Role, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *UserPermission) SetRoles(v []Role) {
	o.Roles = v
}

// GetUiPermissions returns the UiPermissions field value
func (o *UserPermission) GetUiPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UiPermissions
}

// GetUiPermissionsOk returns a tuple with the UiPermissions field value
// and a boolean to check if the value has been set.
func (o *UserPermission) GetUiPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UiPermissions, true
}

// SetUiPermissions sets field value
func (o *UserPermission) SetUiPermissions(v []string) {
	o.UiPermissions = v
}

func (o UserPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	toSerialize["roles"] = o.Roles
	toSerialize["uiPermissions"] = o.UiPermissions
	return toSerialize, nil
}

func (o *UserPermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permissions",
		"roles",
		"uiPermissions",
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

	varUserPermission := _UserPermission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserPermission)

	if err != nil {
		return err
	}

	*o = UserPermission(varUserPermission)

	return err
}

type NullableUserPermission struct {
	value *UserPermission
	isSet bool
}

func (v NullableUserPermission) Get() *UserPermission {
	return v.value
}

func (v *NullableUserPermission) Set(val *UserPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPermission(val *UserPermission) *NullableUserPermission {
	return &NullableUserPermission{value: val, isSet: true}
}

func (v NullableUserPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


