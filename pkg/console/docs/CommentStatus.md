# CommentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNewReply** | Pointer to **bool** |  | [optional] 
**LastReplyTime** | Pointer to **time.Time** |  | [optional] 
**ObservedVersion** | Pointer to **int64** |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**UnreadReplyCount** | Pointer to **int32** |  | [optional] 
**VisibleReplyCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCommentStatus

`func NewCommentStatus() *CommentStatus`

NewCommentStatus instantiates a new CommentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentStatusWithDefaults

`func NewCommentStatusWithDefaults() *CommentStatus`

NewCommentStatusWithDefaults instantiates a new CommentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNewReply

`func (o *CommentStatus) GetHasNewReply() bool`

GetHasNewReply returns the HasNewReply field if non-nil, zero value otherwise.

### GetHasNewReplyOk

`func (o *CommentStatus) GetHasNewReplyOk() (*bool, bool)`

GetHasNewReplyOk returns a tuple with the HasNewReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNewReply

`func (o *CommentStatus) SetHasNewReply(v bool)`

SetHasNewReply sets HasNewReply field to given value.

### HasHasNewReply

`func (o *CommentStatus) HasHasNewReply() bool`

HasHasNewReply returns a boolean if a field has been set.

### GetLastReplyTime

`func (o *CommentStatus) GetLastReplyTime() time.Time`

GetLastReplyTime returns the LastReplyTime field if non-nil, zero value otherwise.

### GetLastReplyTimeOk

`func (o *CommentStatus) GetLastReplyTimeOk() (*time.Time, bool)`

GetLastReplyTimeOk returns a tuple with the LastReplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyTime

`func (o *CommentStatus) SetLastReplyTime(v time.Time)`

SetLastReplyTime sets LastReplyTime field to given value.

### HasLastReplyTime

`func (o *CommentStatus) HasLastReplyTime() bool`

HasLastReplyTime returns a boolean if a field has been set.

### GetObservedVersion

`func (o *CommentStatus) GetObservedVersion() int64`

GetObservedVersion returns the ObservedVersion field if non-nil, zero value otherwise.

### GetObservedVersionOk

`func (o *CommentStatus) GetObservedVersionOk() (*int64, bool)`

GetObservedVersionOk returns a tuple with the ObservedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedVersion

`func (o *CommentStatus) SetObservedVersion(v int64)`

SetObservedVersion sets ObservedVersion field to given value.

### HasObservedVersion

`func (o *CommentStatus) HasObservedVersion() bool`

HasObservedVersion returns a boolean if a field has been set.

### GetReplyCount

`func (o *CommentStatus) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *CommentStatus) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *CommentStatus) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *CommentStatus) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetUnreadReplyCount

`func (o *CommentStatus) GetUnreadReplyCount() int32`

GetUnreadReplyCount returns the UnreadReplyCount field if non-nil, zero value otherwise.

### GetUnreadReplyCountOk

`func (o *CommentStatus) GetUnreadReplyCountOk() (*int32, bool)`

GetUnreadReplyCountOk returns a tuple with the UnreadReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadReplyCount

`func (o *CommentStatus) SetUnreadReplyCount(v int32)`

SetUnreadReplyCount sets UnreadReplyCount field to given value.

### HasUnreadReplyCount

`func (o *CommentStatus) HasUnreadReplyCount() bool`

HasUnreadReplyCount returns a boolean if a field has been set.

### GetVisibleReplyCount

`func (o *CommentStatus) GetVisibleReplyCount() int32`

GetVisibleReplyCount returns the VisibleReplyCount field if non-nil, zero value otherwise.

### GetVisibleReplyCountOk

`func (o *CommentStatus) GetVisibleReplyCountOk() (*int32, bool)`

GetVisibleReplyCountOk returns a tuple with the VisibleReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleReplyCount

`func (o *CommentStatus) SetVisibleReplyCount(v int32)`

SetVisibleReplyCount sets VisibleReplyCount field to given value.

### HasVisibleReplyCount

`func (o *CommentStatus) HasVisibleReplyCount() bool`

HasVisibleReplyCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


