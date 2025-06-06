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

// checks if the ListedAuthProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListedAuthProvider{}

// ListedAuthProvider struct for ListedAuthProvider
type ListedAuthProvider struct {
	AuthType *string `json:"authType,omitempty"`
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`
	BindingUrl *string `json:"bindingUrl,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName string `json:"displayName"`
	Enabled *bool `json:"enabled,omitempty"`
	HelpPage *string `json:"helpPage,omitempty"`
	IsBound *bool `json:"isBound,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Name string `json:"name"`
	Priority *int32 `json:"priority,omitempty"`
	Privileged *bool `json:"privileged,omitempty"`
	SupportsBinding *bool `json:"supportsBinding,omitempty"`
	UnbindingUrl *string `json:"unbindingUrl,omitempty"`
	Website *string `json:"website,omitempty"`
}

type _ListedAuthProvider ListedAuthProvider

// NewListedAuthProvider instantiates a new ListedAuthProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListedAuthProvider(displayName string, name string) *ListedAuthProvider {
	this := ListedAuthProvider{}
	this.DisplayName = displayName
	this.Name = name
	return &this
}

// NewListedAuthProviderWithDefaults instantiates a new ListedAuthProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListedAuthProviderWithDefaults() *ListedAuthProvider {
	this := ListedAuthProvider{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *ListedAuthProvider) SetAuthType(v string) {
	o.AuthType = &v
}

// GetAuthenticationUrl returns the AuthenticationUrl field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetAuthenticationUrl() string {
	if o == nil || IsNil(o.AuthenticationUrl) {
		var ret string
		return ret
	}
	return *o.AuthenticationUrl
}

// GetAuthenticationUrlOk returns a tuple with the AuthenticationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetAuthenticationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationUrl) {
		return nil, false
	}
	return o.AuthenticationUrl, true
}

// HasAuthenticationUrl returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasAuthenticationUrl() bool {
	if o != nil && !IsNil(o.AuthenticationUrl) {
		return true
	}

	return false
}

// SetAuthenticationUrl gets a reference to the given string and assigns it to the AuthenticationUrl field.
func (o *ListedAuthProvider) SetAuthenticationUrl(v string) {
	o.AuthenticationUrl = &v
}

// GetBindingUrl returns the BindingUrl field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetBindingUrl() string {
	if o == nil || IsNil(o.BindingUrl) {
		var ret string
		return ret
	}
	return *o.BindingUrl
}

// GetBindingUrlOk returns a tuple with the BindingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetBindingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BindingUrl) {
		return nil, false
	}
	return o.BindingUrl, true
}

// HasBindingUrl returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasBindingUrl() bool {
	if o != nil && !IsNil(o.BindingUrl) {
		return true
	}

	return false
}

// SetBindingUrl gets a reference to the given string and assigns it to the BindingUrl field.
func (o *ListedAuthProvider) SetBindingUrl(v string) {
	o.BindingUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListedAuthProvider) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value
func (o *ListedAuthProvider) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ListedAuthProvider) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ListedAuthProvider) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHelpPage returns the HelpPage field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetHelpPage() string {
	if o == nil || IsNil(o.HelpPage) {
		var ret string
		return ret
	}
	return *o.HelpPage
}

// GetHelpPageOk returns a tuple with the HelpPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetHelpPageOk() (*string, bool) {
	if o == nil || IsNil(o.HelpPage) {
		return nil, false
	}
	return o.HelpPage, true
}

// HasHelpPage returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasHelpPage() bool {
	if o != nil && !IsNil(o.HelpPage) {
		return true
	}

	return false
}

// SetHelpPage gets a reference to the given string and assigns it to the HelpPage field.
func (o *ListedAuthProvider) SetHelpPage(v string) {
	o.HelpPage = &v
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *ListedAuthProvider) SetIsBound(v bool) {
	o.IsBound = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ListedAuthProvider) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value
func (o *ListedAuthProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListedAuthProvider) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ListedAuthProvider) SetPriority(v int32) {
	o.Priority = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetPrivileged() bool {
	if o == nil || IsNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetPrivilegedOk() (*bool, bool) {
	if o == nil || IsNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasPrivileged() bool {
	if o != nil && !IsNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *ListedAuthProvider) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetSupportsBinding returns the SupportsBinding field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetSupportsBinding() bool {
	if o == nil || IsNil(o.SupportsBinding) {
		var ret bool
		return ret
	}
	return *o.SupportsBinding
}

// GetSupportsBindingOk returns a tuple with the SupportsBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetSupportsBindingOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsBinding) {
		return nil, false
	}
	return o.SupportsBinding, true
}

// HasSupportsBinding returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasSupportsBinding() bool {
	if o != nil && !IsNil(o.SupportsBinding) {
		return true
	}

	return false
}

// SetSupportsBinding gets a reference to the given bool and assigns it to the SupportsBinding field.
func (o *ListedAuthProvider) SetSupportsBinding(v bool) {
	o.SupportsBinding = &v
}

// GetUnbindingUrl returns the UnbindingUrl field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetUnbindingUrl() string {
	if o == nil || IsNil(o.UnbindingUrl) {
		var ret string
		return ret
	}
	return *o.UnbindingUrl
}

// GetUnbindingUrlOk returns a tuple with the UnbindingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetUnbindingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UnbindingUrl) {
		return nil, false
	}
	return o.UnbindingUrl, true
}

// HasUnbindingUrl returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasUnbindingUrl() bool {
	if o != nil && !IsNil(o.UnbindingUrl) {
		return true
	}

	return false
}

// SetUnbindingUrl gets a reference to the given string and assigns it to the UnbindingUrl field.
func (o *ListedAuthProvider) SetUnbindingUrl(v string) {
	o.UnbindingUrl = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ListedAuthProvider) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedAuthProvider) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ListedAuthProvider) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *ListedAuthProvider) SetWebsite(v string) {
	o.Website = &v
}

func (o ListedAuthProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListedAuthProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	if !IsNil(o.AuthenticationUrl) {
		toSerialize["authenticationUrl"] = o.AuthenticationUrl
	}
	if !IsNil(o.BindingUrl) {
		toSerialize["bindingUrl"] = o.BindingUrl
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.HelpPage) {
		toSerialize["helpPage"] = o.HelpPage
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !IsNil(o.SupportsBinding) {
		toSerialize["supportsBinding"] = o.SupportsBinding
	}
	if !IsNil(o.UnbindingUrl) {
		toSerialize["unbindingUrl"] = o.UnbindingUrl
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

func (o *ListedAuthProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"name",
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

	varListedAuthProvider := _ListedAuthProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListedAuthProvider)

	if err != nil {
		return err
	}

	*o = ListedAuthProvider(varListedAuthProvider)

	return err
}

type NullableListedAuthProvider struct {
	value *ListedAuthProvider
	isSet bool
}

func (v NullableListedAuthProvider) Get() *ListedAuthProvider {
	return v.value
}

func (v *NullableListedAuthProvider) Set(val *ListedAuthProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableListedAuthProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableListedAuthProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListedAuthProvider(val *ListedAuthProvider) *NullableListedAuthProvider {
	return &NullableListedAuthProvider{value: val, isSet: true}
}

func (v NullableListedAuthProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListedAuthProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


