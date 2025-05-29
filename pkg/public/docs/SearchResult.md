# SearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]HaloDocument**](HaloDocument.md) |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**ProcessingTimeMillis** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewSearchResult

`func NewSearchResult() *SearchResult`

NewSearchResult instantiates a new SearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWithDefaults

`func NewSearchResultWithDefaults() *SearchResult`

NewSearchResultWithDefaults instantiates a new SearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *SearchResult) GetHits() []HaloDocument`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchResult) GetHitsOk() (*[]HaloDocument, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchResult) SetHits(v []HaloDocument)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchResult) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetKeyword

`func (o *SearchResult) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *SearchResult) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *SearchResult) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *SearchResult) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetLimit

`func (o *SearchResult) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchResult) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchResult) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchResult) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetProcessingTimeMillis

`func (o *SearchResult) GetProcessingTimeMillis() int64`

GetProcessingTimeMillis returns the ProcessingTimeMillis field if non-nil, zero value otherwise.

### GetProcessingTimeMillisOk

`func (o *SearchResult) GetProcessingTimeMillisOk() (*int64, bool)`

GetProcessingTimeMillisOk returns a tuple with the ProcessingTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMillis

`func (o *SearchResult) SetProcessingTimeMillis(v int64)`

SetProcessingTimeMillis sets ProcessingTimeMillis field to given value.

### HasProcessingTimeMillis

`func (o *SearchResult) HasProcessingTimeMillis() bool`

HasProcessingTimeMillis returns a boolean if a field has been set.

### GetTotal

`func (o *SearchResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


