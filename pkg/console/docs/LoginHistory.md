# LoginHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginAt** | **time.Time** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**SourceIp** | **string** |  | 
**Successful** | **bool** |  | 
**UserAgent** | **string** |  | 

## Methods

### NewLoginHistory

`func NewLoginHistory(loginAt time.Time, sourceIp string, successful bool, userAgent string, ) *LoginHistory`

NewLoginHistory instantiates a new LoginHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginHistoryWithDefaults

`func NewLoginHistoryWithDefaults() *LoginHistory`

NewLoginHistoryWithDefaults instantiates a new LoginHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginAt

`func (o *LoginHistory) GetLoginAt() time.Time`

GetLoginAt returns the LoginAt field if non-nil, zero value otherwise.

### GetLoginAtOk

`func (o *LoginHistory) GetLoginAtOk() (*time.Time, bool)`

GetLoginAtOk returns a tuple with the LoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAt

`func (o *LoginHistory) SetLoginAt(v time.Time)`

SetLoginAt sets LoginAt field to given value.


### GetReason

`func (o *LoginHistory) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LoginHistory) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LoginHistory) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LoginHistory) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSourceIp

`func (o *LoginHistory) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *LoginHistory) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *LoginHistory) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.


### GetSuccessful

`func (o *LoginHistory) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *LoginHistory) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *LoginHistory) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.


### GetUserAgent

`func (o *LoginHistory) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *LoginHistory) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *LoginHistory) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


