# ReplyVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Owner** | [**OwnerInfo**](OwnerInfo.md) |  | 
**Spec** | [**ReplySpec**](ReplySpec.md) |  | 
**Stats** | [**CommentStatsVo**](CommentStatsVo.md) |  | 

## Methods

### NewReplyVo

`func NewReplyVo(metadata Metadata, owner OwnerInfo, spec ReplySpec, stats CommentStatsVo, ) *ReplyVo`

NewReplyVo instantiates a new ReplyVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyVoWithDefaults

`func NewReplyVoWithDefaults() *ReplyVo`

NewReplyVoWithDefaults instantiates a new ReplyVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ReplyVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReplyVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReplyVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetOwner

`func (o *ReplyVo) GetOwner() OwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ReplyVo) GetOwnerOk() (*OwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ReplyVo) SetOwner(v OwnerInfo)`

SetOwner sets Owner field to given value.


### GetSpec

`func (o *ReplyVo) GetSpec() ReplySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ReplyVo) GetSpecOk() (*ReplySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ReplyVo) SetSpec(v ReplySpec)`

SetSpec sets Spec field to given value.


### GetStats

`func (o *ReplyVo) GetStats() CommentStatsVo`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ReplyVo) GetStatsOk() (*CommentStatsVo, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ReplyVo) SetStats(v CommentStatsVo)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


