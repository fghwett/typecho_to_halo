# ListedSnapshotDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Spec** | [**ListedSnapshotSpec**](ListedSnapshotSpec.md) |  | 

## Methods

### NewListedSnapshotDto

`func NewListedSnapshotDto(metadata Metadata, spec ListedSnapshotSpec, ) *ListedSnapshotDto`

NewListedSnapshotDto instantiates a new ListedSnapshotDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedSnapshotDtoWithDefaults

`func NewListedSnapshotDtoWithDefaults() *ListedSnapshotDto`

NewListedSnapshotDtoWithDefaults instantiates a new ListedSnapshotDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ListedSnapshotDto) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListedSnapshotDto) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListedSnapshotDto) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ListedSnapshotDto) GetSpec() ListedSnapshotSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ListedSnapshotDto) GetSpecOk() (*ListedSnapshotSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ListedSnapshotDto) SetSpec(v ListedSnapshotSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


