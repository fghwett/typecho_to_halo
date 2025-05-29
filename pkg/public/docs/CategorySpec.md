# CategorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to **[]string** |  | [optional] 
**Cover** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**HideFromList** | Pointer to **bool** |  | [optional] 
**PostTemplate** | Pointer to **string** |  | [optional] 
**PreventParentPostCascadeQuery** | Pointer to **bool** |  | [optional] 
**Priority** | **int32** |  | [default to 0]
**Slug** | **string** |  | 
**Template** | Pointer to **string** |  | [optional] 

## Methods

### NewCategorySpec

`func NewCategorySpec(displayName string, priority int32, slug string, ) *CategorySpec`

NewCategorySpec instantiates a new CategorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySpecWithDefaults

`func NewCategorySpecWithDefaults() *CategorySpec`

NewCategorySpecWithDefaults instantiates a new CategorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *CategorySpec) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CategorySpec) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CategorySpec) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CategorySpec) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCover

`func (o *CategorySpec) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *CategorySpec) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *CategorySpec) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *CategorySpec) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetDescription

`func (o *CategorySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategorySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategorySpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategorySpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CategorySpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CategorySpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CategorySpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHideFromList

`func (o *CategorySpec) GetHideFromList() bool`

GetHideFromList returns the HideFromList field if non-nil, zero value otherwise.

### GetHideFromListOk

`func (o *CategorySpec) GetHideFromListOk() (*bool, bool)`

GetHideFromListOk returns a tuple with the HideFromList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromList

`func (o *CategorySpec) SetHideFromList(v bool)`

SetHideFromList sets HideFromList field to given value.

### HasHideFromList

`func (o *CategorySpec) HasHideFromList() bool`

HasHideFromList returns a boolean if a field has been set.

### GetPostTemplate

`func (o *CategorySpec) GetPostTemplate() string`

GetPostTemplate returns the PostTemplate field if non-nil, zero value otherwise.

### GetPostTemplateOk

`func (o *CategorySpec) GetPostTemplateOk() (*string, bool)`

GetPostTemplateOk returns a tuple with the PostTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTemplate

`func (o *CategorySpec) SetPostTemplate(v string)`

SetPostTemplate sets PostTemplate field to given value.

### HasPostTemplate

`func (o *CategorySpec) HasPostTemplate() bool`

HasPostTemplate returns a boolean if a field has been set.

### GetPreventParentPostCascadeQuery

`func (o *CategorySpec) GetPreventParentPostCascadeQuery() bool`

GetPreventParentPostCascadeQuery returns the PreventParentPostCascadeQuery field if non-nil, zero value otherwise.

### GetPreventParentPostCascadeQueryOk

`func (o *CategorySpec) GetPreventParentPostCascadeQueryOk() (*bool, bool)`

GetPreventParentPostCascadeQueryOk returns a tuple with the PreventParentPostCascadeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventParentPostCascadeQuery

`func (o *CategorySpec) SetPreventParentPostCascadeQuery(v bool)`

SetPreventParentPostCascadeQuery sets PreventParentPostCascadeQuery field to given value.

### HasPreventParentPostCascadeQuery

`func (o *CategorySpec) HasPreventParentPostCascadeQuery() bool`

HasPreventParentPostCascadeQuery returns a boolean if a field has been set.

### GetPriority

`func (o *CategorySpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CategorySpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CategorySpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSlug

`func (o *CategorySpec) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CategorySpec) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CategorySpec) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTemplate

`func (o *CategorySpec) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CategorySpec) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CategorySpec) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CategorySpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


