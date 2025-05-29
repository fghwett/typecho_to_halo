# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**CreationTimestamp** | Pointer to **NullableTime** |  | [optional] 
**DeletionTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Finalizers** | Pointer to **[]string** |  | [optional] 
**GenerateName** | Pointer to **string** | The name field will be generated automatically according to the given generateName field | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** | Metadata name | 
**Version** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata(name string, ) *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *Metadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Metadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Metadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Metadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *Metadata) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *Metadata) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *Metadata) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *Metadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### SetCreationTimestampNil

`func (o *Metadata) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *Metadata) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil
### GetDeletionTimestamp

`func (o *Metadata) GetDeletionTimestamp() time.Time`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *Metadata) GetDeletionTimestampOk() (*time.Time, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *Metadata) SetDeletionTimestamp(v time.Time)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *Metadata) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### SetDeletionTimestampNil

`func (o *Metadata) SetDeletionTimestampNil(b bool)`

 SetDeletionTimestampNil sets the value for DeletionTimestamp to be an explicit nil

### UnsetDeletionTimestamp
`func (o *Metadata) UnsetDeletionTimestamp()`

UnsetDeletionTimestamp ensures that no value is present for DeletionTimestamp, not even an explicit nil
### GetFinalizers

`func (o *Metadata) GetFinalizers() []*string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *Metadata) GetFinalizersOk() (*[]*string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *Metadata) SetFinalizers(v []*string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *Metadata) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### SetFinalizersNil

`func (o *Metadata) SetFinalizersNil(b bool)`

 SetFinalizersNil sets the value for Finalizers to be an explicit nil

### UnsetFinalizers
`func (o *Metadata) UnsetFinalizers()`

UnsetFinalizers ensures that no value is present for Finalizers, not even an explicit nil
### GetGenerateName

`func (o *Metadata) GetGenerateName() string`

GetGenerateName returns the GenerateName field if non-nil, zero value otherwise.

### GetGenerateNameOk

`func (o *Metadata) GetGenerateNameOk() (*string, bool)`

GetGenerateNameOk returns a tuple with the GenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateName

`func (o *Metadata) SetGenerateName(v string)`

SetGenerateName sets GenerateName field to given value.

### HasGenerateName

`func (o *Metadata) HasGenerateName() bool`

HasGenerateName returns a boolean if a field has been set.

### GetLabels

`func (o *Metadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Metadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Metadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Metadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *Metadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metadata) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Metadata) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Metadata) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Metadata) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Metadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Metadata) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Metadata) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


