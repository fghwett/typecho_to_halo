# LocalThumbnail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**LocalThumbnailSpec**](LocalThumbnailSpec.md) |  | 
**Status** | Pointer to [**LocalThumbnailStatus**](LocalThumbnailStatus.md) |  | [optional] 

## Methods

### NewLocalThumbnail

`func NewLocalThumbnail(apiVersion string, kind string, metadata Metadata, spec LocalThumbnailSpec, ) *LocalThumbnail`

NewLocalThumbnail instantiates a new LocalThumbnail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalThumbnailWithDefaults

`func NewLocalThumbnailWithDefaults() *LocalThumbnail`

NewLocalThumbnailWithDefaults instantiates a new LocalThumbnail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *LocalThumbnail) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *LocalThumbnail) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *LocalThumbnail) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *LocalThumbnail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LocalThumbnail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LocalThumbnail) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *LocalThumbnail) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LocalThumbnail) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LocalThumbnail) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *LocalThumbnail) GetSpec() LocalThumbnailSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *LocalThumbnail) GetSpecOk() (*LocalThumbnailSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *LocalThumbnail) SetSpec(v LocalThumbnailSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *LocalThumbnail) GetStatus() LocalThumbnailStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LocalThumbnail) GetStatusOk() (*LocalThumbnailStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LocalThumbnail) SetStatus(v LocalThumbnailStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LocalThumbnail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


