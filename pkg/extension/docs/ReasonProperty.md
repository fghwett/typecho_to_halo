# ReasonProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Optional** | Pointer to **bool** |  | [optional] [default to false]
**Type** | **string** |  | 

## Methods

### NewReasonProperty

`func NewReasonProperty(name string, type_ string, ) *ReasonProperty`

NewReasonProperty instantiates a new ReasonProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReasonPropertyWithDefaults

`func NewReasonPropertyWithDefaults() *ReasonProperty`

NewReasonPropertyWithDefaults instantiates a new ReasonProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReasonProperty) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReasonProperty) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReasonProperty) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReasonProperty) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ReasonProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReasonProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReasonProperty) SetName(v string)`

SetName sets Name field to given value.


### GetOptional

`func (o *ReasonProperty) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ReasonProperty) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ReasonProperty) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ReasonProperty) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetType

`func (o *ReasonProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReasonProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReasonProperty) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


