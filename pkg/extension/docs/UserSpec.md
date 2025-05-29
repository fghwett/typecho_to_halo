# UserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisplayName** | **string** |  | 
**Email** | **string** |  | 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**LoginHistoryLimit** | Pointer to **int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**RegisteredAt** | Pointer to **time.Time** |  | [optional] 
**TotpEncryptedSecret** | Pointer to **string** |  | [optional] 
**TwoFactorAuthEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserSpec

`func NewUserSpec(displayName string, email string, ) *UserSpec`

NewUserSpec instantiates a new UserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSpecWithDefaults

`func NewUserSpecWithDefaults() *UserSpec`

NewUserSpecWithDefaults instantiates a new UserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *UserSpec) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserSpec) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserSpec) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserSpec) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBio

`func (o *UserSpec) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UserSpec) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UserSpec) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UserSpec) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetDisabled

`func (o *UserSpec) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserSpec) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserSpec) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UserSpec) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmail

`func (o *UserSpec) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSpec) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSpec) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailVerified

`func (o *UserSpec) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserSpec) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserSpec) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *UserSpec) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetLoginHistoryLimit

`func (o *UserSpec) GetLoginHistoryLimit() int32`

GetLoginHistoryLimit returns the LoginHistoryLimit field if non-nil, zero value otherwise.

### GetLoginHistoryLimitOk

`func (o *UserSpec) GetLoginHistoryLimitOk() (*int32, bool)`

GetLoginHistoryLimitOk returns a tuple with the LoginHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHistoryLimit

`func (o *UserSpec) SetLoginHistoryLimit(v int32)`

SetLoginHistoryLimit sets LoginHistoryLimit field to given value.

### HasLoginHistoryLimit

`func (o *UserSpec) HasLoginHistoryLimit() bool`

HasLoginHistoryLimit returns a boolean if a field has been set.

### GetPassword

`func (o *UserSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *UserSpec) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserSpec) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserSpec) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserSpec) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetRegisteredAt

`func (o *UserSpec) GetRegisteredAt() time.Time`

GetRegisteredAt returns the RegisteredAt field if non-nil, zero value otherwise.

### GetRegisteredAtOk

`func (o *UserSpec) GetRegisteredAtOk() (*time.Time, bool)`

GetRegisteredAtOk returns a tuple with the RegisteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAt

`func (o *UserSpec) SetRegisteredAt(v time.Time)`

SetRegisteredAt sets RegisteredAt field to given value.

### HasRegisteredAt

`func (o *UserSpec) HasRegisteredAt() bool`

HasRegisteredAt returns a boolean if a field has been set.

### GetTotpEncryptedSecret

`func (o *UserSpec) GetTotpEncryptedSecret() string`

GetTotpEncryptedSecret returns the TotpEncryptedSecret field if non-nil, zero value otherwise.

### GetTotpEncryptedSecretOk

`func (o *UserSpec) GetTotpEncryptedSecretOk() (*string, bool)`

GetTotpEncryptedSecretOk returns a tuple with the TotpEncryptedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEncryptedSecret

`func (o *UserSpec) SetTotpEncryptedSecret(v string)`

SetTotpEncryptedSecret sets TotpEncryptedSecret field to given value.

### HasTotpEncryptedSecret

`func (o *UserSpec) HasTotpEncryptedSecret() bool`

HasTotpEncryptedSecret returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *UserSpec) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *UserSpec) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *UserSpec) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *UserSpec) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


