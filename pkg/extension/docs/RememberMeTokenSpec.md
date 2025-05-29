# RememberMeTokenSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUsed** | Pointer to **time.Time** |  | [optional] 
**Series** | **string** |  | 
**TokenValue** | **string** |  | 
**Username** | **string** |  | 

## Methods

### NewRememberMeTokenSpec

`func NewRememberMeTokenSpec(series string, tokenValue string, username string, ) *RememberMeTokenSpec`

NewRememberMeTokenSpec instantiates a new RememberMeTokenSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRememberMeTokenSpecWithDefaults

`func NewRememberMeTokenSpecWithDefaults() *RememberMeTokenSpec`

NewRememberMeTokenSpecWithDefaults instantiates a new RememberMeTokenSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUsed

`func (o *RememberMeTokenSpec) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *RememberMeTokenSpec) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *RememberMeTokenSpec) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *RememberMeTokenSpec) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetSeries

`func (o *RememberMeTokenSpec) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *RememberMeTokenSpec) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *RememberMeTokenSpec) SetSeries(v string)`

SetSeries sets Series field to given value.


### GetTokenValue

`func (o *RememberMeTokenSpec) GetTokenValue() string`

GetTokenValue returns the TokenValue field if non-nil, zero value otherwise.

### GetTokenValueOk

`func (o *RememberMeTokenSpec) GetTokenValueOk() (*string, bool)`

GetTokenValueOk returns a tuple with the TokenValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValue

`func (o *RememberMeTokenSpec) SetTokenValue(v string)`

SetTokenValue sets TokenValue field to given value.


### GetUsername

`func (o *RememberMeTokenSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RememberMeTokenSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RememberMeTokenSpec) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


