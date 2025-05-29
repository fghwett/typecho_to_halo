# AuthProviderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** |  | 
**AuthenticationUrl** | **string** | Authentication url of the auth provider | 
**BindingUrl** | Pointer to **string** |  | [optional] 
**ConfigMapRef** | Pointer to [**ConfigMapRef**](ConfigMapRef.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** | Display name of the auth provider | 
**HelpPage** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**RememberMeSupport** | Pointer to **bool** |  | [optional] 
**SettingRef** | Pointer to [**SettingRef**](SettingRef.md) |  | [optional] 
**UnbindUrl** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthProviderSpec

`func NewAuthProviderSpec(authType string, authenticationUrl string, displayName string, ) *AuthProviderSpec`

NewAuthProviderSpec instantiates a new AuthProviderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthProviderSpecWithDefaults

`func NewAuthProviderSpecWithDefaults() *AuthProviderSpec`

NewAuthProviderSpecWithDefaults instantiates a new AuthProviderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AuthProviderSpec) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthProviderSpec) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthProviderSpec) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetAuthenticationUrl

`func (o *AuthProviderSpec) GetAuthenticationUrl() string`

GetAuthenticationUrl returns the AuthenticationUrl field if non-nil, zero value otherwise.

### GetAuthenticationUrlOk

`func (o *AuthProviderSpec) GetAuthenticationUrlOk() (*string, bool)`

GetAuthenticationUrlOk returns a tuple with the AuthenticationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationUrl

`func (o *AuthProviderSpec) SetAuthenticationUrl(v string)`

SetAuthenticationUrl sets AuthenticationUrl field to given value.


### GetBindingUrl

`func (o *AuthProviderSpec) GetBindingUrl() string`

GetBindingUrl returns the BindingUrl field if non-nil, zero value otherwise.

### GetBindingUrlOk

`func (o *AuthProviderSpec) GetBindingUrlOk() (*string, bool)`

GetBindingUrlOk returns a tuple with the BindingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUrl

`func (o *AuthProviderSpec) SetBindingUrl(v string)`

SetBindingUrl sets BindingUrl field to given value.

### HasBindingUrl

`func (o *AuthProviderSpec) HasBindingUrl() bool`

HasBindingUrl returns a boolean if a field has been set.

### GetConfigMapRef

`func (o *AuthProviderSpec) GetConfigMapRef() ConfigMapRef`

GetConfigMapRef returns the ConfigMapRef field if non-nil, zero value otherwise.

### GetConfigMapRefOk

`func (o *AuthProviderSpec) GetConfigMapRefOk() (*ConfigMapRef, bool)`

GetConfigMapRefOk returns a tuple with the ConfigMapRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapRef

`func (o *AuthProviderSpec) SetConfigMapRef(v ConfigMapRef)`

SetConfigMapRef sets ConfigMapRef field to given value.

### HasConfigMapRef

`func (o *AuthProviderSpec) HasConfigMapRef() bool`

HasConfigMapRef returns a boolean if a field has been set.

### GetDescription

`func (o *AuthProviderSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthProviderSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthProviderSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthProviderSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AuthProviderSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthProviderSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthProviderSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHelpPage

`func (o *AuthProviderSpec) GetHelpPage() string`

GetHelpPage returns the HelpPage field if non-nil, zero value otherwise.

### GetHelpPageOk

`func (o *AuthProviderSpec) GetHelpPageOk() (*string, bool)`

GetHelpPageOk returns a tuple with the HelpPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpPage

`func (o *AuthProviderSpec) SetHelpPage(v string)`

SetHelpPage sets HelpPage field to given value.

### HasHelpPage

`func (o *AuthProviderSpec) HasHelpPage() bool`

HasHelpPage returns a boolean if a field has been set.

### GetLogo

`func (o *AuthProviderSpec) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AuthProviderSpec) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AuthProviderSpec) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AuthProviderSpec) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMethod

`func (o *AuthProviderSpec) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthProviderSpec) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthProviderSpec) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuthProviderSpec) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRememberMeSupport

`func (o *AuthProviderSpec) GetRememberMeSupport() bool`

GetRememberMeSupport returns the RememberMeSupport field if non-nil, zero value otherwise.

### GetRememberMeSupportOk

`func (o *AuthProviderSpec) GetRememberMeSupportOk() (*bool, bool)`

GetRememberMeSupportOk returns a tuple with the RememberMeSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMeSupport

`func (o *AuthProviderSpec) SetRememberMeSupport(v bool)`

SetRememberMeSupport sets RememberMeSupport field to given value.

### HasRememberMeSupport

`func (o *AuthProviderSpec) HasRememberMeSupport() bool`

HasRememberMeSupport returns a boolean if a field has been set.

### GetSettingRef

`func (o *AuthProviderSpec) GetSettingRef() SettingRef`

GetSettingRef returns the SettingRef field if non-nil, zero value otherwise.

### GetSettingRefOk

`func (o *AuthProviderSpec) GetSettingRefOk() (*SettingRef, bool)`

GetSettingRefOk returns a tuple with the SettingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingRef

`func (o *AuthProviderSpec) SetSettingRef(v SettingRef)`

SetSettingRef sets SettingRef field to given value.

### HasSettingRef

`func (o *AuthProviderSpec) HasSettingRef() bool`

HasSettingRef returns a boolean if a field has been set.

### GetUnbindUrl

`func (o *AuthProviderSpec) GetUnbindUrl() string`

GetUnbindUrl returns the UnbindUrl field if non-nil, zero value otherwise.

### GetUnbindUrlOk

`func (o *AuthProviderSpec) GetUnbindUrlOk() (*string, bool)`

GetUnbindUrlOk returns a tuple with the UnbindUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbindUrl

`func (o *AuthProviderSpec) SetUnbindUrl(v string)`

SetUnbindUrl sets UnbindUrl field to given value.

### HasUnbindUrl

`func (o *AuthProviderSpec) HasUnbindUrl() bool`

HasUnbindUrl returns a boolean if a field has been set.

### GetWebsite

`func (o *AuthProviderSpec) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AuthProviderSpec) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AuthProviderSpec) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AuthProviderSpec) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


