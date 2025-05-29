# ListedAuthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** |  | [optional] 
**AuthenticationUrl** | Pointer to **string** |  | [optional] 
**BindingUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**HelpPage** | Pointer to **string** |  | [optional] 
**IsBound** | Pointer to **bool** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Priority** | Pointer to **int32** |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**SupportsBinding** | Pointer to **bool** |  | [optional] 
**UnbindingUrl** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewListedAuthProvider

`func NewListedAuthProvider(displayName string, name string, ) *ListedAuthProvider`

NewListedAuthProvider instantiates a new ListedAuthProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedAuthProviderWithDefaults

`func NewListedAuthProviderWithDefaults() *ListedAuthProvider`

NewListedAuthProviderWithDefaults instantiates a new ListedAuthProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *ListedAuthProvider) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ListedAuthProvider) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ListedAuthProvider) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ListedAuthProvider) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthenticationUrl

`func (o *ListedAuthProvider) GetAuthenticationUrl() string`

GetAuthenticationUrl returns the AuthenticationUrl field if non-nil, zero value otherwise.

### GetAuthenticationUrlOk

`func (o *ListedAuthProvider) GetAuthenticationUrlOk() (*string, bool)`

GetAuthenticationUrlOk returns a tuple with the AuthenticationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationUrl

`func (o *ListedAuthProvider) SetAuthenticationUrl(v string)`

SetAuthenticationUrl sets AuthenticationUrl field to given value.

### HasAuthenticationUrl

`func (o *ListedAuthProvider) HasAuthenticationUrl() bool`

HasAuthenticationUrl returns a boolean if a field has been set.

### GetBindingUrl

`func (o *ListedAuthProvider) GetBindingUrl() string`

GetBindingUrl returns the BindingUrl field if non-nil, zero value otherwise.

### GetBindingUrlOk

`func (o *ListedAuthProvider) GetBindingUrlOk() (*string, bool)`

GetBindingUrlOk returns a tuple with the BindingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUrl

`func (o *ListedAuthProvider) SetBindingUrl(v string)`

SetBindingUrl sets BindingUrl field to given value.

### HasBindingUrl

`func (o *ListedAuthProvider) HasBindingUrl() bool`

HasBindingUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ListedAuthProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListedAuthProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListedAuthProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListedAuthProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListedAuthProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListedAuthProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListedAuthProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnabled

`func (o *ListedAuthProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListedAuthProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListedAuthProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListedAuthProvider) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHelpPage

`func (o *ListedAuthProvider) GetHelpPage() string`

GetHelpPage returns the HelpPage field if non-nil, zero value otherwise.

### GetHelpPageOk

`func (o *ListedAuthProvider) GetHelpPageOk() (*string, bool)`

GetHelpPageOk returns a tuple with the HelpPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpPage

`func (o *ListedAuthProvider) SetHelpPage(v string)`

SetHelpPage sets HelpPage field to given value.

### HasHelpPage

`func (o *ListedAuthProvider) HasHelpPage() bool`

HasHelpPage returns a boolean if a field has been set.

### GetIsBound

`func (o *ListedAuthProvider) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *ListedAuthProvider) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *ListedAuthProvider) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.

### HasIsBound

`func (o *ListedAuthProvider) HasIsBound() bool`

HasIsBound returns a boolean if a field has been set.

### GetLogo

`func (o *ListedAuthProvider) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ListedAuthProvider) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ListedAuthProvider) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ListedAuthProvider) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *ListedAuthProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListedAuthProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListedAuthProvider) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *ListedAuthProvider) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListedAuthProvider) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListedAuthProvider) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ListedAuthProvider) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrivileged

`func (o *ListedAuthProvider) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *ListedAuthProvider) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *ListedAuthProvider) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *ListedAuthProvider) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetSupportsBinding

`func (o *ListedAuthProvider) GetSupportsBinding() bool`

GetSupportsBinding returns the SupportsBinding field if non-nil, zero value otherwise.

### GetSupportsBindingOk

`func (o *ListedAuthProvider) GetSupportsBindingOk() (*bool, bool)`

GetSupportsBindingOk returns a tuple with the SupportsBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsBinding

`func (o *ListedAuthProvider) SetSupportsBinding(v bool)`

SetSupportsBinding sets SupportsBinding field to given value.

### HasSupportsBinding

`func (o *ListedAuthProvider) HasSupportsBinding() bool`

HasSupportsBinding returns a boolean if a field has been set.

### GetUnbindingUrl

`func (o *ListedAuthProvider) GetUnbindingUrl() string`

GetUnbindingUrl returns the UnbindingUrl field if non-nil, zero value otherwise.

### GetUnbindingUrlOk

`func (o *ListedAuthProvider) GetUnbindingUrlOk() (*string, bool)`

GetUnbindingUrlOk returns a tuple with the UnbindingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbindingUrl

`func (o *ListedAuthProvider) SetUnbindingUrl(v string)`

SetUnbindingUrl sets UnbindingUrl field to given value.

### HasUnbindingUrl

`func (o *ListedAuthProvider) HasUnbindingUrl() bool`

HasUnbindingUrl returns a boolean if a field has been set.

### GetWebsite

`func (o *ListedAuthProvider) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ListedAuthProvider) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ListedAuthProvider) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ListedAuthProvider) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


