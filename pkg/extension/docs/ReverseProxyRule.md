# ReverseProxyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**FileReverseProxyProvider**](FileReverseProxyProvider.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewReverseProxyRule

`func NewReverseProxyRule() *ReverseProxyRule`

NewReverseProxyRule instantiates a new ReverseProxyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseProxyRuleWithDefaults

`func NewReverseProxyRuleWithDefaults() *ReverseProxyRule`

NewReverseProxyRuleWithDefaults instantiates a new ReverseProxyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *ReverseProxyRule) GetFile() FileReverseProxyProvider`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ReverseProxyRule) GetFileOk() (*FileReverseProxyProvider, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ReverseProxyRule) SetFile(v FileReverseProxyProvider)`

SetFile sets File field to given value.

### HasFile

`func (o *ReverseProxyRule) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetPath

`func (o *ReverseProxyRule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReverseProxyRule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReverseProxyRule) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReverseProxyRule) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


