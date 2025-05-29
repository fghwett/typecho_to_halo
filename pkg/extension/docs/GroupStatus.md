# GroupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAttachments** | Pointer to **int64** | Total of attachments under the current group | [optional] 
**UpdateTimestamp** | Pointer to **time.Time** | Update timestamp of the group | [optional] 

## Methods

### NewGroupStatus

`func NewGroupStatus() *GroupStatus`

NewGroupStatus instantiates a new GroupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupStatusWithDefaults

`func NewGroupStatusWithDefaults() *GroupStatus`

NewGroupStatusWithDefaults instantiates a new GroupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAttachments

`func (o *GroupStatus) GetTotalAttachments() int64`

GetTotalAttachments returns the TotalAttachments field if non-nil, zero value otherwise.

### GetTotalAttachmentsOk

`func (o *GroupStatus) GetTotalAttachmentsOk() (*int64, bool)`

GetTotalAttachmentsOk returns a tuple with the TotalAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAttachments

`func (o *GroupStatus) SetTotalAttachments(v int64)`

SetTotalAttachments sets TotalAttachments field to given value.

### HasTotalAttachments

`func (o *GroupStatus) HasTotalAttachments() bool`

HasTotalAttachments returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *GroupStatus) GetUpdateTimestamp() time.Time`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *GroupStatus) GetUpdateTimestampOk() (*time.Time, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *GroupStatus) SetUpdateTimestamp(v time.Time)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *GroupStatus) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


