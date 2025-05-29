# SearchOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**FilterExposed** | Pointer to **bool** |  | [optional] 
**FilterPublished** | Pointer to **bool** |  | [optional] 
**FilterRecycled** | Pointer to **bool** |  | [optional] 
**HighlightPostTag** | Pointer to **string** |  | [optional] 
**HighlightPreTag** | Pointer to **string** |  | [optional] 
**IncludeCategoryNames** | Pointer to **[]string** |  | [optional] 
**IncludeOwnerNames** | Pointer to **[]string** |  | [optional] 
**IncludeTagNames** | Pointer to **[]string** |  | [optional] 
**IncludeTypes** | Pointer to **[]string** |  | [optional] 
**Keyword** | **string** |  | 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchOption

`func NewSearchOption(keyword string, ) *SearchOption`

NewSearchOption instantiates a new SearchOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchOptionWithDefaults

`func NewSearchOptionWithDefaults() *SearchOption`

NewSearchOptionWithDefaults instantiates a new SearchOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *SearchOption) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SearchOption) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SearchOption) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SearchOption) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFilterExposed

`func (o *SearchOption) GetFilterExposed() bool`

GetFilterExposed returns the FilterExposed field if non-nil, zero value otherwise.

### GetFilterExposedOk

`func (o *SearchOption) GetFilterExposedOk() (*bool, bool)`

GetFilterExposedOk returns a tuple with the FilterExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExposed

`func (o *SearchOption) SetFilterExposed(v bool)`

SetFilterExposed sets FilterExposed field to given value.

### HasFilterExposed

`func (o *SearchOption) HasFilterExposed() bool`

HasFilterExposed returns a boolean if a field has been set.

### GetFilterPublished

`func (o *SearchOption) GetFilterPublished() bool`

GetFilterPublished returns the FilterPublished field if non-nil, zero value otherwise.

### GetFilterPublishedOk

`func (o *SearchOption) GetFilterPublishedOk() (*bool, bool)`

GetFilterPublishedOk returns a tuple with the FilterPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPublished

`func (o *SearchOption) SetFilterPublished(v bool)`

SetFilterPublished sets FilterPublished field to given value.

### HasFilterPublished

`func (o *SearchOption) HasFilterPublished() bool`

HasFilterPublished returns a boolean if a field has been set.

### GetFilterRecycled

`func (o *SearchOption) GetFilterRecycled() bool`

GetFilterRecycled returns the FilterRecycled field if non-nil, zero value otherwise.

### GetFilterRecycledOk

`func (o *SearchOption) GetFilterRecycledOk() (*bool, bool)`

GetFilterRecycledOk returns a tuple with the FilterRecycled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterRecycled

`func (o *SearchOption) SetFilterRecycled(v bool)`

SetFilterRecycled sets FilterRecycled field to given value.

### HasFilterRecycled

`func (o *SearchOption) HasFilterRecycled() bool`

HasFilterRecycled returns a boolean if a field has been set.

### GetHighlightPostTag

`func (o *SearchOption) GetHighlightPostTag() string`

GetHighlightPostTag returns the HighlightPostTag field if non-nil, zero value otherwise.

### GetHighlightPostTagOk

`func (o *SearchOption) GetHighlightPostTagOk() (*string, bool)`

GetHighlightPostTagOk returns a tuple with the HighlightPostTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightPostTag

`func (o *SearchOption) SetHighlightPostTag(v string)`

SetHighlightPostTag sets HighlightPostTag field to given value.

### HasHighlightPostTag

`func (o *SearchOption) HasHighlightPostTag() bool`

HasHighlightPostTag returns a boolean if a field has been set.

### GetHighlightPreTag

`func (o *SearchOption) GetHighlightPreTag() string`

GetHighlightPreTag returns the HighlightPreTag field if non-nil, zero value otherwise.

### GetHighlightPreTagOk

`func (o *SearchOption) GetHighlightPreTagOk() (*string, bool)`

GetHighlightPreTagOk returns a tuple with the HighlightPreTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightPreTag

`func (o *SearchOption) SetHighlightPreTag(v string)`

SetHighlightPreTag sets HighlightPreTag field to given value.

### HasHighlightPreTag

`func (o *SearchOption) HasHighlightPreTag() bool`

HasHighlightPreTag returns a boolean if a field has been set.

### GetIncludeCategoryNames

`func (o *SearchOption) GetIncludeCategoryNames() []string`

GetIncludeCategoryNames returns the IncludeCategoryNames field if non-nil, zero value otherwise.

### GetIncludeCategoryNamesOk

`func (o *SearchOption) GetIncludeCategoryNamesOk() (*[]string, bool)`

GetIncludeCategoryNamesOk returns a tuple with the IncludeCategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCategoryNames

`func (o *SearchOption) SetIncludeCategoryNames(v []string)`

SetIncludeCategoryNames sets IncludeCategoryNames field to given value.

### HasIncludeCategoryNames

`func (o *SearchOption) HasIncludeCategoryNames() bool`

HasIncludeCategoryNames returns a boolean if a field has been set.

### GetIncludeOwnerNames

`func (o *SearchOption) GetIncludeOwnerNames() []string`

GetIncludeOwnerNames returns the IncludeOwnerNames field if non-nil, zero value otherwise.

### GetIncludeOwnerNamesOk

`func (o *SearchOption) GetIncludeOwnerNamesOk() (*[]string, bool)`

GetIncludeOwnerNamesOk returns a tuple with the IncludeOwnerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOwnerNames

`func (o *SearchOption) SetIncludeOwnerNames(v []string)`

SetIncludeOwnerNames sets IncludeOwnerNames field to given value.

### HasIncludeOwnerNames

`func (o *SearchOption) HasIncludeOwnerNames() bool`

HasIncludeOwnerNames returns a boolean if a field has been set.

### GetIncludeTagNames

`func (o *SearchOption) GetIncludeTagNames() []string`

GetIncludeTagNames returns the IncludeTagNames field if non-nil, zero value otherwise.

### GetIncludeTagNamesOk

`func (o *SearchOption) GetIncludeTagNamesOk() (*[]string, bool)`

GetIncludeTagNamesOk returns a tuple with the IncludeTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagNames

`func (o *SearchOption) SetIncludeTagNames(v []string)`

SetIncludeTagNames sets IncludeTagNames field to given value.

### HasIncludeTagNames

`func (o *SearchOption) HasIncludeTagNames() bool`

HasIncludeTagNames returns a boolean if a field has been set.

### GetIncludeTypes

`func (o *SearchOption) GetIncludeTypes() []string`

GetIncludeTypes returns the IncludeTypes field if non-nil, zero value otherwise.

### GetIncludeTypesOk

`func (o *SearchOption) GetIncludeTypesOk() (*[]string, bool)`

GetIncludeTypesOk returns a tuple with the IncludeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTypes

`func (o *SearchOption) SetIncludeTypes(v []string)`

SetIncludeTypes sets IncludeTypes field to given value.

### HasIncludeTypes

`func (o *SearchOption) HasIncludeTypes() bool`

HasIncludeTypes returns a boolean if a field has been set.

### GetKeyword

`func (o *SearchOption) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *SearchOption) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *SearchOption) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.


### GetLimit

`func (o *SearchOption) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchOption) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchOption) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchOption) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


