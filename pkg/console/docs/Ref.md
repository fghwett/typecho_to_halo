# Ref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Extension group | [optional] 
**Kind** | Pointer to **string** | Extension kind | [optional] 
**Name** | **string** | Extension name. This field is mandatory | 
**Version** | Pointer to **string** | Extension version | [optional] 

## Methods

### NewRef

`func NewRef(name string, ) *Ref`

NewRef instantiates a new Ref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefWithDefaults

`func NewRefWithDefaults() *Ref`

NewRefWithDefaults instantiates a new Ref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *Ref) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Ref) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Ref) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Ref) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *Ref) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Ref) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Ref) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Ref) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *Ref) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ref) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ref) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Ref) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Ref) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Ref) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Ref) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


