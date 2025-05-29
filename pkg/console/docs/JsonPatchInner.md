# JsonPatchInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Path** | **string** | A JSON Pointer path pointing to the location to move/copy from. | 
**Value** | **interface{}** | Value can be any JSON value | 
**From** | **string** | A JSON Pointer path pointing to the location to move/copy from. | 

## Methods

### NewJsonPatchInner

`func NewJsonPatchInner(op string, path string, value interface{}, from string, ) *JsonPatchInner`

NewJsonPatchInner instantiates a new JsonPatchInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPatchInnerWithDefaults

`func NewJsonPatchInnerWithDefaults() *JsonPatchInner`

NewJsonPatchInnerWithDefaults instantiates a new JsonPatchInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *JsonPatchInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JsonPatchInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JsonPatchInner) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *JsonPatchInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JsonPatchInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JsonPatchInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *JsonPatchInner) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JsonPatchInner) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JsonPatchInner) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *JsonPatchInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JsonPatchInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetFrom

`func (o *JsonPatchInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *JsonPatchInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *JsonPatchInner) SetFrom(v string)`

SetFrom sets From field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


