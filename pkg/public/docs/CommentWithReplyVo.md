# CommentWithReplyVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Owner** | [**OwnerInfo**](OwnerInfo.md) |  | 
**Replies** | Pointer to [**ListResultReplyVo**](ListResultReplyVo.md) |  | [optional] 
**Spec** | [**CommentSpec**](CommentSpec.md) |  | 
**Stats** | [**CommentStatsVo**](CommentStatsVo.md) |  | 
**Status** | Pointer to [**CommentStatus**](CommentStatus.md) |  | [optional] 

## Methods

### NewCommentWithReplyVo

`func NewCommentWithReplyVo(metadata Metadata, owner OwnerInfo, spec CommentSpec, stats CommentStatsVo, ) *CommentWithReplyVo`

NewCommentWithReplyVo instantiates a new CommentWithReplyVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithReplyVoWithDefaults

`func NewCommentWithReplyVoWithDefaults() *CommentWithReplyVo`

NewCommentWithReplyVoWithDefaults instantiates a new CommentWithReplyVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CommentWithReplyVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommentWithReplyVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommentWithReplyVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetOwner

`func (o *CommentWithReplyVo) GetOwner() OwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CommentWithReplyVo) GetOwnerOk() (*OwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CommentWithReplyVo) SetOwner(v OwnerInfo)`

SetOwner sets Owner field to given value.


### GetReplies

`func (o *CommentWithReplyVo) GetReplies() ListResultReplyVo`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *CommentWithReplyVo) GetRepliesOk() (*ListResultReplyVo, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *CommentWithReplyVo) SetReplies(v ListResultReplyVo)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *CommentWithReplyVo) HasReplies() bool`

HasReplies returns a boolean if a field has been set.

### GetSpec

`func (o *CommentWithReplyVo) GetSpec() CommentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CommentWithReplyVo) GetSpecOk() (*CommentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CommentWithReplyVo) SetSpec(v CommentSpec)`

SetSpec sets Spec field to given value.


### GetStats

`func (o *CommentWithReplyVo) GetStats() CommentStatsVo`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CommentWithReplyVo) GetStatsOk() (*CommentStatsVo, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CommentWithReplyVo) SetStats(v CommentStatsVo)`

SetStats sets Stats field to given value.


### GetStatus

`func (o *CommentWithReplyVo) GetStatus() CommentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommentWithReplyVo) GetStatusOk() (*CommentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommentWithReplyVo) SetStatus(v CommentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommentWithReplyVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


