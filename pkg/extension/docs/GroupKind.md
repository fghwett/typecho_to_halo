# GroupKind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupKind

`func NewGroupKind() *GroupKind`

NewGroupKind instantiates a new GroupKind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupKindWithDefaults

`func NewGroupKindWithDefaults() *GroupKind`

NewGroupKindWithDefaults instantiates a new GroupKind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GroupKind) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GroupKind) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GroupKind) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GroupKind) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *GroupKind) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GroupKind) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GroupKind) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GroupKind) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


