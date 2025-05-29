# SnapShotSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentPatch** | Pointer to **string** |  | [optional] 
**Contributors** | Pointer to **[]string** |  | [optional] 
**LastModifyTime** | Pointer to **time.Time** |  | [optional] 
**Owner** | **string** |  | 
**ParentSnapshotName** | Pointer to **string** |  | [optional] 
**RawPatch** | Pointer to **string** |  | [optional] 
**RawType** | **string** |  | 
**SubjectRef** | [**Ref**](Ref.md) |  | 

## Methods

### NewSnapShotSpec

`func NewSnapShotSpec(owner string, rawType string, subjectRef Ref, ) *SnapShotSpec`

NewSnapShotSpec instantiates a new SnapShotSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapShotSpecWithDefaults

`func NewSnapShotSpecWithDefaults() *SnapShotSpec`

NewSnapShotSpecWithDefaults instantiates a new SnapShotSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentPatch

`func (o *SnapShotSpec) GetContentPatch() string`

GetContentPatch returns the ContentPatch field if non-nil, zero value otherwise.

### GetContentPatchOk

`func (o *SnapShotSpec) GetContentPatchOk() (*string, bool)`

GetContentPatchOk returns a tuple with the ContentPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPatch

`func (o *SnapShotSpec) SetContentPatch(v string)`

SetContentPatch sets ContentPatch field to given value.

### HasContentPatch

`func (o *SnapShotSpec) HasContentPatch() bool`

HasContentPatch returns a boolean if a field has been set.

### GetContributors

`func (o *SnapShotSpec) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *SnapShotSpec) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *SnapShotSpec) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *SnapShotSpec) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetLastModifyTime

`func (o *SnapShotSpec) GetLastModifyTime() time.Time`

GetLastModifyTime returns the LastModifyTime field if non-nil, zero value otherwise.

### GetLastModifyTimeOk

`func (o *SnapShotSpec) GetLastModifyTimeOk() (*time.Time, bool)`

GetLastModifyTimeOk returns a tuple with the LastModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifyTime

`func (o *SnapShotSpec) SetLastModifyTime(v time.Time)`

SetLastModifyTime sets LastModifyTime field to given value.

### HasLastModifyTime

`func (o *SnapShotSpec) HasLastModifyTime() bool`

HasLastModifyTime returns a boolean if a field has been set.

### GetOwner

`func (o *SnapShotSpec) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SnapShotSpec) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SnapShotSpec) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParentSnapshotName

`func (o *SnapShotSpec) GetParentSnapshotName() string`

GetParentSnapshotName returns the ParentSnapshotName field if non-nil, zero value otherwise.

### GetParentSnapshotNameOk

`func (o *SnapShotSpec) GetParentSnapshotNameOk() (*string, bool)`

GetParentSnapshotNameOk returns a tuple with the ParentSnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshotName

`func (o *SnapShotSpec) SetParentSnapshotName(v string)`

SetParentSnapshotName sets ParentSnapshotName field to given value.

### HasParentSnapshotName

`func (o *SnapShotSpec) HasParentSnapshotName() bool`

HasParentSnapshotName returns a boolean if a field has been set.

### GetRawPatch

`func (o *SnapShotSpec) GetRawPatch() string`

GetRawPatch returns the RawPatch field if non-nil, zero value otherwise.

### GetRawPatchOk

`func (o *SnapShotSpec) GetRawPatchOk() (*string, bool)`

GetRawPatchOk returns a tuple with the RawPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPatch

`func (o *SnapShotSpec) SetRawPatch(v string)`

SetRawPatch sets RawPatch field to given value.

### HasRawPatch

`func (o *SnapShotSpec) HasRawPatch() bool`

HasRawPatch returns a boolean if a field has been set.

### GetRawType

`func (o *SnapShotSpec) GetRawType() string`

GetRawType returns the RawType field if non-nil, zero value otherwise.

### GetRawTypeOk

`func (o *SnapShotSpec) GetRawTypeOk() (*string, bool)`

GetRawTypeOk returns a tuple with the RawType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawType

`func (o *SnapShotSpec) SetRawType(v string)`

SetRawType sets RawType field to given value.


### GetSubjectRef

`func (o *SnapShotSpec) GetSubjectRef() Ref`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *SnapShotSpec) GetSubjectRefOk() (*Ref, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *SnapShotSpec) SetSubjectRef(v Ref)`

SetSubjectRef sets SubjectRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


