# AnnotationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**AnnotationSettingSpec**](AnnotationSettingSpec.md) |  | 

## Methods

### NewAnnotationSetting

`func NewAnnotationSetting(apiVersion string, kind string, metadata Metadata, spec AnnotationSettingSpec, ) *AnnotationSetting`

NewAnnotationSetting instantiates a new AnnotationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationSettingWithDefaults

`func NewAnnotationSettingWithDefaults() *AnnotationSetting`

NewAnnotationSettingWithDefaults instantiates a new AnnotationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AnnotationSetting) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AnnotationSetting) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AnnotationSetting) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *AnnotationSetting) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AnnotationSetting) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AnnotationSetting) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *AnnotationSetting) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AnnotationSetting) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AnnotationSetting) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *AnnotationSetting) GetSpec() AnnotationSettingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AnnotationSetting) GetSpecOk() (*AnnotationSettingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AnnotationSetting) SetSpec(v AnnotationSettingSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


