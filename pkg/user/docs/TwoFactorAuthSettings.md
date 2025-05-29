# TwoFactorAuthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TotpConfigured** | Pointer to **bool** |  | [optional] 

## Methods

### NewTwoFactorAuthSettings

`func NewTwoFactorAuthSettings() *TwoFactorAuthSettings`

NewTwoFactorAuthSettings instantiates a new TwoFactorAuthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorAuthSettingsWithDefaults

`func NewTwoFactorAuthSettingsWithDefaults() *TwoFactorAuthSettings`

NewTwoFactorAuthSettingsWithDefaults instantiates a new TwoFactorAuthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *TwoFactorAuthSettings) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TwoFactorAuthSettings) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TwoFactorAuthSettings) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *TwoFactorAuthSettings) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetEmailVerified

`func (o *TwoFactorAuthSettings) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *TwoFactorAuthSettings) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *TwoFactorAuthSettings) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *TwoFactorAuthSettings) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetEnabled

`func (o *TwoFactorAuthSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TwoFactorAuthSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TwoFactorAuthSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TwoFactorAuthSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTotpConfigured

`func (o *TwoFactorAuthSettings) GetTotpConfigured() bool`

GetTotpConfigured returns the TotpConfigured field if non-nil, zero value otherwise.

### GetTotpConfiguredOk

`func (o *TwoFactorAuthSettings) GetTotpConfiguredOk() (*bool, bool)`

GetTotpConfiguredOk returns a tuple with the TotpConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpConfigured

`func (o *TwoFactorAuthSettings) SetTotpConfigured(v bool)`

SetTotpConfigured sets TotpConfigured field to given value.

### HasTotpConfigured

`func (o *TwoFactorAuthSettings) HasTotpConfigured() bool`

HasTotpConfigured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


