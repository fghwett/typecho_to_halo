# CategoryVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**PostCount** | Pointer to **int32** |  | [optional] 
**Spec** | Pointer to [**CategorySpec**](CategorySpec.md) |  | [optional] 
**Status** | Pointer to [**CategoryStatus**](CategoryStatus.md) |  | [optional] 

## Methods

### NewCategoryVo

`func NewCategoryVo(metadata Metadata, ) *CategoryVo`

NewCategoryVo instantiates a new CategoryVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryVoWithDefaults

`func NewCategoryVoWithDefaults() *CategoryVo`

NewCategoryVoWithDefaults instantiates a new CategoryVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CategoryVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CategoryVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CategoryVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetPostCount

`func (o *CategoryVo) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *CategoryVo) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *CategoryVo) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.

### HasPostCount

`func (o *CategoryVo) HasPostCount() bool`

HasPostCount returns a boolean if a field has been set.

### GetSpec

`func (o *CategoryVo) GetSpec() CategorySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CategoryVo) GetSpecOk() (*CategorySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CategoryVo) SetSpec(v CategorySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CategoryVo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CategoryVo) GetStatus() CategoryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CategoryVo) GetStatusOk() (*CategoryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CategoryVo) SetStatus(v CategoryStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CategoryVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


