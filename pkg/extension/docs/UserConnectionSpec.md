# UserConnectionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderUserId** | **string** |  | 
**RegistrationId** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewUserConnectionSpec

`func NewUserConnectionSpec(providerUserId string, registrationId string, username string, ) *UserConnectionSpec`

NewUserConnectionSpec instantiates a new UserConnectionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConnectionSpecWithDefaults

`func NewUserConnectionSpecWithDefaults() *UserConnectionSpec`

NewUserConnectionSpecWithDefaults instantiates a new UserConnectionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderUserId

`func (o *UserConnectionSpec) GetProviderUserId() string`

GetProviderUserId returns the ProviderUserId field if non-nil, zero value otherwise.

### GetProviderUserIdOk

`func (o *UserConnectionSpec) GetProviderUserIdOk() (*string, bool)`

GetProviderUserIdOk returns a tuple with the ProviderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUserId

`func (o *UserConnectionSpec) SetProviderUserId(v string)`

SetProviderUserId sets ProviderUserId field to given value.


### GetRegistrationId

`func (o *UserConnectionSpec) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *UserConnectionSpec) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *UserConnectionSpec) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.


### GetUpdatedAt

`func (o *UserConnectionSpec) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserConnectionSpec) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserConnectionSpec) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserConnectionSpec) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *UserConnectionSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserConnectionSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserConnectionSpec) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


