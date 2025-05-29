# TotpAuthLinkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthLink** | Pointer to **string** |  | [optional] 
**RawSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewTotpAuthLinkResponse

`func NewTotpAuthLinkResponse() *TotpAuthLinkResponse`

NewTotpAuthLinkResponse instantiates a new TotpAuthLinkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotpAuthLinkResponseWithDefaults

`func NewTotpAuthLinkResponseWithDefaults() *TotpAuthLinkResponse`

NewTotpAuthLinkResponseWithDefaults instantiates a new TotpAuthLinkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthLink

`func (o *TotpAuthLinkResponse) GetAuthLink() string`

GetAuthLink returns the AuthLink field if non-nil, zero value otherwise.

### GetAuthLinkOk

`func (o *TotpAuthLinkResponse) GetAuthLinkOk() (*string, bool)`

GetAuthLinkOk returns a tuple with the AuthLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthLink

`func (o *TotpAuthLinkResponse) SetAuthLink(v string)`

SetAuthLink sets AuthLink field to given value.

### HasAuthLink

`func (o *TotpAuthLinkResponse) HasAuthLink() bool`

HasAuthLink returns a boolean if a field has been set.

### GetRawSecret

`func (o *TotpAuthLinkResponse) GetRawSecret() string`

GetRawSecret returns the RawSecret field if non-nil, zero value otherwise.

### GetRawSecretOk

`func (o *TotpAuthLinkResponse) GetRawSecretOk() (*string, bool)`

GetRawSecretOk returns a tuple with the RawSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSecret

`func (o *TotpAuthLinkResponse) SetRawSecret(v string)`

SetRawSecret sets RawSecret field to given value.

### HasRawSecret

`func (o *TotpAuthLinkResponse) HasRawSecret() bool`

HasRawSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


