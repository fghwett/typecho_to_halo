# NotifierDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | Pointer to [**NotifierDescriptorSpec**](NotifierDescriptorSpec.md) |  | [optional] 

## Methods

### NewNotifierDescriptor

`func NewNotifierDescriptor(apiVersion string, kind string, metadata Metadata, ) *NotifierDescriptor`

NewNotifierDescriptor instantiates a new NotifierDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifierDescriptorWithDefaults

`func NewNotifierDescriptorWithDefaults() *NotifierDescriptor`

NewNotifierDescriptorWithDefaults instantiates a new NotifierDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotifierDescriptor) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotifierDescriptor) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotifierDescriptor) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NotifierDescriptor) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotifierDescriptor) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotifierDescriptor) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NotifierDescriptor) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotifierDescriptor) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotifierDescriptor) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *NotifierDescriptor) GetSpec() NotifierDescriptorSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NotifierDescriptor) GetSpecOk() (*NotifierDescriptorSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NotifierDescriptor) SetSpec(v NotifierDescriptorSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NotifierDescriptor) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


