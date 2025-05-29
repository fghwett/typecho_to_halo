# UserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLoginAt** | Pointer to **time.Time** |  | [optional] 
**LoginHistories** | Pointer to [**[]LoginHistory**](LoginHistory.md) |  | [optional] 
**Permalink** | Pointer to **string** |  | [optional] 

## Methods

### NewUserStatus

`func NewUserStatus() *UserStatus`

NewUserStatus instantiates a new UserStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusWithDefaults

`func NewUserStatusWithDefaults() *UserStatus`

NewUserStatusWithDefaults instantiates a new UserStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLoginAt

`func (o *UserStatus) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UserStatus) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UserStatus) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UserStatus) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLoginHistories

`func (o *UserStatus) GetLoginHistories() []LoginHistory`

GetLoginHistories returns the LoginHistories field if non-nil, zero value otherwise.

### GetLoginHistoriesOk

`func (o *UserStatus) GetLoginHistoriesOk() (*[]LoginHistory, bool)`

GetLoginHistoriesOk returns a tuple with the LoginHistories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHistories

`func (o *UserStatus) SetLoginHistories(v []LoginHistory)`

SetLoginHistories sets LoginHistories field to given value.

### HasLoginHistories

`func (o *UserStatus) HasLoginHistories() bool`

HasLoginHistories returns a boolean if a field has been set.

### GetPermalink

`func (o *UserStatus) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *UserStatus) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *UserStatus) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *UserStatus) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


