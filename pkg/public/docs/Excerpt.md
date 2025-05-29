# Excerpt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoGenerate** | **bool** |  | [default to true]
**Raw** | Pointer to **string** |  | [optional] 

## Methods

### NewExcerpt

`func NewExcerpt(autoGenerate bool, ) *Excerpt`

NewExcerpt instantiates a new Excerpt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcerptWithDefaults

`func NewExcerptWithDefaults() *Excerpt`

NewExcerptWithDefaults instantiates a new Excerpt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoGenerate

`func (o *Excerpt) GetAutoGenerate() bool`

GetAutoGenerate returns the AutoGenerate field if non-nil, zero value otherwise.

### GetAutoGenerateOk

`func (o *Excerpt) GetAutoGenerateOk() (*bool, bool)`

GetAutoGenerateOk returns a tuple with the AutoGenerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGenerate

`func (o *Excerpt) SetAutoGenerate(v bool)`

SetAutoGenerate sets AutoGenerate field to given value.


### GetRaw

`func (o *Excerpt) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *Excerpt) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *Excerpt) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *Excerpt) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


