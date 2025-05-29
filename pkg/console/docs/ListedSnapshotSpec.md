# ListedSnapshotSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModifyTime** | Pointer to **time.Time** |  | [optional] 
**Owner** | **string** |  | 

## Methods

### NewListedSnapshotSpec

`func NewListedSnapshotSpec(owner string, ) *ListedSnapshotSpec`

NewListedSnapshotSpec instantiates a new ListedSnapshotSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedSnapshotSpecWithDefaults

`func NewListedSnapshotSpecWithDefaults() *ListedSnapshotSpec`

NewListedSnapshotSpecWithDefaults instantiates a new ListedSnapshotSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModifyTime

`func (o *ListedSnapshotSpec) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *ListedSnapshotSpec) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *ListedSnapshotSpec) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *ListedSnapshotSpec) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetOwner

`func (o *ListedSnapshotSpec) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListedSnapshotSpec) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListedSnapshotSpec) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


