# TagVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**PostCount** | Pointer to **int32** |  | [optional] 
**Spec** | Pointer to [**TagSpec**](TagSpec.md) |  | [optional] 
**Status** | Pointer to [**TagStatus**](TagStatus.md) |  | [optional] 

## Methods

### NewTagVo

`func NewTagVo(metadata Metadata, ) *TagVo`

NewTagVo instantiates a new TagVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagVoWithDefaults

`func NewTagVoWithDefaults() *TagVo`

NewTagVoWithDefaults instantiates a new TagVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *TagVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TagVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TagVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetPostCount

`func (o *TagVo) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *TagVo) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *TagVo) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.

### HasPostCount

`func (o *TagVo) HasPostCount() bool`

HasPostCount returns a boolean if a field has been set.

### GetSpec

`func (o *TagVo) GetSpec() TagSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TagVo) GetSpecOk() (*TagSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TagVo) SetSpec(v TagSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TagVo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *TagVo) GetStatus() TagStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TagVo) GetStatusOk() (*TagStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TagVo) SetStatus(v TagStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TagVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


