# PatSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**LastUsed** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Revoked** | Pointer to **bool** |  | [optional] 
**RevokesAt** | Pointer to **time.Time** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**TokenId** | **string** |  | 
**Username** | **string** |  | 

## Methods

### NewPatSpec

`func NewPatSpec(name string, tokenId string, username string, ) *PatSpec`

NewPatSpec instantiates a new PatSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatSpecWithDefaults

`func NewPatSpecWithDefaults() *PatSpec`

NewPatSpecWithDefaults instantiates a new PatSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PatSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PatSpec) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PatSpec) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PatSpec) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PatSpec) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLastUsed

`func (o *PatSpec) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *PatSpec) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *PatSpec) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *PatSpec) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetName

`func (o *PatSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatSpec) SetName(v string)`

SetName sets Name field to given value.


### GetRevoked

`func (o *PatSpec) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *PatSpec) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *PatSpec) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *PatSpec) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetRevokesAt

`func (o *PatSpec) GetRevokesAt() time.Time`

GetRevokesAt returns the RevokesAt field if non-nil, zero value otherwise.

### GetRevokesAtOk

`func (o *PatSpec) GetRevokesAtOk() (*time.Time, bool)`

GetRevokesAtOk returns a tuple with the RevokesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokesAt

`func (o *PatSpec) SetRevokesAt(v time.Time)`

SetRevokesAt sets RevokesAt field to given value.

### HasRevokesAt

`func (o *PatSpec) HasRevokesAt() bool`

HasRevokesAt returns a boolean if a field has been set.

### GetRoles

`func (o *PatSpec) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PatSpec) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PatSpec) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PatSpec) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScopes

`func (o *PatSpec) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PatSpec) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PatSpec) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PatSpec) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenId

`func (o *PatSpec) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PatSpec) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PatSpec) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetUsername

`func (o *PatSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatSpec) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


