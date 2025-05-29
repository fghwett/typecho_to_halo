# CommentVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Owner** | [**OwnerInfo**](OwnerInfo.md) |  | 
**Spec** | [**CommentSpec**](CommentSpec.md) |  | 
**Stats** | [**CommentStatsVo**](CommentStatsVo.md) |  | 
**Status** | Pointer to [**CommentStatus**](CommentStatus.md) |  | [optional] 

## Methods

### NewCommentVo

`func NewCommentVo(metadata Metadata, owner OwnerInfo, spec CommentSpec, stats CommentStatsVo, ) *CommentVo`

NewCommentVo instantiates a new CommentVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentVoWithDefaults

`func NewCommentVoWithDefaults() *CommentVo`

NewCommentVoWithDefaults instantiates a new CommentVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CommentVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommentVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommentVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetOwner

`func (o *CommentVo) GetOwner() OwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CommentVo) GetOwnerOk() (*OwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CommentVo) SetOwner(v OwnerInfo)`

SetOwner sets Owner field to given value.


### GetSpec

`func (o *CommentVo) GetSpec() CommentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CommentVo) GetSpecOk() (*CommentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CommentVo) SetSpec(v CommentSpec)`

SetSpec sets Spec field to given value.


### GetStats

`func (o *CommentVo) GetStats() CommentStatsVo`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CommentVo) GetStatsOk() (*CommentStatsVo, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CommentVo) SetStats(v CommentStatsVo)`

SetStats sets Stats field to given value.


### GetStatus

`func (o *CommentVo) GetStatus() CommentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommentVo) GetStatusOk() (*CommentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommentVo) SetStatus(v CommentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommentVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


