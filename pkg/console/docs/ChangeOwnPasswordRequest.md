# ChangeOwnPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPassword** | **string** | Old password. | 
**Password** | **string** | New password. | 

## Methods

### NewChangeOwnPasswordRequest

`func NewChangeOwnPasswordRequest(oldPassword string, password string, ) *ChangeOwnPasswordRequest`

NewChangeOwnPasswordRequest instantiates a new ChangeOwnPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeOwnPasswordRequestWithDefaults

`func NewChangeOwnPasswordRequestWithDefaults() *ChangeOwnPasswordRequest`

NewChangeOwnPasswordRequestWithDefaults instantiates a new ChangeOwnPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPassword

`func (o *ChangeOwnPasswordRequest) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *ChangeOwnPasswordRequest) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *ChangeOwnPasswordRequest) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.


### GetPassword

`func (o *ChangeOwnPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ChangeOwnPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ChangeOwnPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


