# ReasonTypeNotifierMatrix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifiers** | Pointer to [**[]NotifierInfo**](NotifierInfo.md) |  | [optional] 
**ReasonTypes** | Pointer to [**[]ReasonTypeInfo**](ReasonTypeInfo.md) |  | [optional] 
**StateMatrix** | Pointer to **[][]bool** |  | [optional] 

## Methods

### NewReasonTypeNotifierMatrix

`func NewReasonTypeNotifierMatrix() *ReasonTypeNotifierMatrix`

NewReasonTypeNotifierMatrix instantiates a new ReasonTypeNotifierMatrix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReasonTypeNotifierMatrixWithDefaults

`func NewReasonTypeNotifierMatrixWithDefaults() *ReasonTypeNotifierMatrix`

NewReasonTypeNotifierMatrixWithDefaults instantiates a new ReasonTypeNotifierMatrix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifiers

`func (o *ReasonTypeNotifierMatrix) GetNotifiers() []NotifierInfo`

GetNotifiers returns the Notifiers field if non-nil, zero value otherwise.

### GetNotifiersOk

`func (o *ReasonTypeNotifierMatrix) GetNotifiersOk() (*[]NotifierInfo, bool)`

GetNotifiersOk returns a tuple with the Notifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiers

`func (o *ReasonTypeNotifierMatrix) SetNotifiers(v []NotifierInfo)`

SetNotifiers sets Notifiers field to given value.

### HasNotifiers

`func (o *ReasonTypeNotifierMatrix) HasNotifiers() bool`

HasNotifiers returns a boolean if a field has been set.

### GetReasonTypes

`func (o *ReasonTypeNotifierMatrix) GetReasonTypes() []ReasonTypeInfo`

GetReasonTypes returns the ReasonTypes field if non-nil, zero value otherwise.

### GetReasonTypesOk

`func (o *ReasonTypeNotifierMatrix) GetReasonTypesOk() (*[]ReasonTypeInfo, bool)`

GetReasonTypesOk returns a tuple with the ReasonTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonTypes

`func (o *ReasonTypeNotifierMatrix) SetReasonTypes(v []ReasonTypeInfo)`

SetReasonTypes sets ReasonTypes field to given value.

### HasReasonTypes

`func (o *ReasonTypeNotifierMatrix) HasReasonTypes() bool`

HasReasonTypes returns a boolean if a field has been set.

### GetStateMatrix

`func (o *ReasonTypeNotifierMatrix) GetStateMatrix() [][]bool`

GetStateMatrix returns the StateMatrix field if non-nil, zero value otherwise.

### GetStateMatrixOk

`func (o *ReasonTypeNotifierMatrix) GetStateMatrixOk() (*[][]bool, bool)`

GetStateMatrixOk returns a tuple with the StateMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMatrix

`func (o *ReasonTypeNotifierMatrix) SetStateMatrix(v [][]bool)`

SetStateMatrix sets StateMatrix field to given value.

### HasStateMatrix

`func (o *ReasonTypeNotifierMatrix) HasStateMatrix() bool`

HasStateMatrix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


