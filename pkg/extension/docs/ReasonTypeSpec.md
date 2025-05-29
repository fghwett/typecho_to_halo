# ReasonTypeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**DisplayName** | **string** |  | 
**Properties** | Pointer to [**[]ReasonProperty**](ReasonProperty.md) |  | [optional] 

## Methods

### NewReasonTypeSpec

`func NewReasonTypeSpec(description string, displayName string, ) *ReasonTypeSpec`

NewReasonTypeSpec instantiates a new ReasonTypeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReasonTypeSpecWithDefaults

`func NewReasonTypeSpecWithDefaults() *ReasonTypeSpec`

NewReasonTypeSpecWithDefaults instantiates a new ReasonTypeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReasonTypeSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReasonTypeSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReasonTypeSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *ReasonTypeSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ReasonTypeSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ReasonTypeSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetProperties

`func (o *ReasonTypeSpec) GetProperties() []ReasonProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReasonTypeSpec) GetPropertiesOk() (*[]ReasonProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReasonTypeSpec) SetProperties(v []ReasonProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ReasonTypeSpec) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


