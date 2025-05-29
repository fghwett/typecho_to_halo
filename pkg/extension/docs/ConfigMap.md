# ConfigMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Data** | Pointer to **map[string]string** |  | [optional] 
**Kind** | **string** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewConfigMap

`func NewConfigMap(apiVersion string, kind string, metadata Metadata, ) *ConfigMap`

NewConfigMap instantiates a new ConfigMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMapWithDefaults

`func NewConfigMapWithDefaults() *ConfigMap`

NewConfigMapWithDefaults instantiates a new ConfigMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ConfigMap) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConfigMap) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConfigMap) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetData

`func (o *ConfigMap) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigMap) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigMap) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *ConfigMap) HasData() bool`

HasData returns a boolean if a field has been set.

### GetKind

`func (o *ConfigMap) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConfigMap) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConfigMap) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ConfigMap) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConfigMap) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConfigMap) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


