# TagSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** |  | [optional] 
**Cover** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**Slug** | **string** |  | 

## Methods

### NewTagSpec

`func NewTagSpec(displayName string, slug string, ) *TagSpec`

NewTagSpec instantiates a new TagSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSpecWithDefaults

`func NewTagSpecWithDefaults() *TagSpec`

NewTagSpecWithDefaults instantiates a new TagSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *TagSpec) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TagSpec) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TagSpec) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TagSpec) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCover

`func (o *TagSpec) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *TagSpec) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *TagSpec) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *TagSpec) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetDisplayName

`func (o *TagSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TagSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TagSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSlug

`func (o *TagSpec) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TagSpec) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TagSpec) SetSlug(v string)`

SetSlug sets Slug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


