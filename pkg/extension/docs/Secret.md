# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Data** | Pointer to **map[string]string** |  | [optional] 
**Kind** | **string** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**StringData** | Pointer to **map[string]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSecret

`func NewSecret(apiVersion string, kind string, metadata Metadata, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Secret) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Secret) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Secret) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetData

`func (o *Secret) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Secret) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Secret) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *Secret) HasData() bool`

HasData returns a boolean if a field has been set.

### GetKind

`func (o *Secret) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Secret) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Secret) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Secret) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Secret) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Secret) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetStringData

`func (o *Secret) GetStringData() map[string]string`

GetStringData returns the StringData field if non-nil, zero value otherwise.

### GetStringDataOk

`func (o *Secret) GetStringDataOk() (*map[string]string, bool)`

GetStringDataOk returns a tuple with the StringData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringData

`func (o *Secret) SetStringData(v map[string]string)`

SetStringData sets StringData field to given value.

### HasStringData

`func (o *Secret) HasStringData() bool`

HasStringData returns a boolean if a field has been set.

### GetType

`func (o *Secret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Secret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Secret) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Secret) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


