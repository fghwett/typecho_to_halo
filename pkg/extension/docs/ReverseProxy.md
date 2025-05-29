# ReverseProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Rules** | Pointer to [**[]ReverseProxyRule**](ReverseProxyRule.md) |  | [optional] 

## Methods

### NewReverseProxy

`func NewReverseProxy(apiVersion string, kind string, metadata Metadata, ) *ReverseProxy`

NewReverseProxy instantiates a new ReverseProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseProxyWithDefaults

`func NewReverseProxyWithDefaults() *ReverseProxy`

NewReverseProxyWithDefaults instantiates a new ReverseProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ReverseProxy) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ReverseProxy) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ReverseProxy) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ReverseProxy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReverseProxy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReverseProxy) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ReverseProxy) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReverseProxy) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReverseProxy) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetRules

`func (o *ReverseProxy) GetRules() []ReverseProxyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ReverseProxy) GetRulesOk() (*[]ReverseProxyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ReverseProxy) SetRules(v []ReverseProxyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ReverseProxy) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


