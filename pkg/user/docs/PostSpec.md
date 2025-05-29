# PostSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowComment** | **bool** |  | [default to true]
**BaseSnapshot** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Cover** | Pointer to **string** |  | [optional] 
**Deleted** | **bool** |  | [default to false]
**Excerpt** | [**Excerpt**](Excerpt.md) |  | 
**HeadSnapshot** | Pointer to **string** |  | [optional] 
**HtmlMetas** | Pointer to **[]map[string]string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Pinned** | **bool** |  | [default to false]
**Priority** | **int32** |  | [default to 0]
**Publish** | **bool** |  | [default to false]
**PublishTime** | Pointer to **time.Time** |  | [optional] 
**ReleaseSnapshot** | Pointer to **string** |  | [optional] 
**Slug** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Visible** | **string** |  | [default to "PUBLIC"]

## Methods

### NewPostSpec

`func NewPostSpec(allowComment bool, deleted bool, excerpt Excerpt, pinned bool, priority int32, publish bool, slug string, title string, visible string, ) *PostSpec`

NewPostSpec instantiates a new PostSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSpecWithDefaults

`func NewPostSpecWithDefaults() *PostSpec`

NewPostSpecWithDefaults instantiates a new PostSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowComment

`func (o *PostSpec) GetAllowComment() bool`

GetAllowComment returns the AllowComment field if non-nil, zero value otherwise.

### GetAllowCommentOk

`func (o *PostSpec) GetAllowCommentOk() (*bool, bool)`

GetAllowCommentOk returns a tuple with the AllowComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowComment

`func (o *PostSpec) SetAllowComment(v bool)`

SetAllowComment sets AllowComment field to given value.


### GetBaseSnapshot

`func (o *PostSpec) GetBaseSnapshot() string`

GetBaseSnapshot returns the BaseSnapshot field if non-nil, zero value otherwise.

### GetBaseSnapshotOk

`func (o *PostSpec) GetBaseSnapshotOk() (*string, bool)`

GetBaseSnapshotOk returns a tuple with the BaseSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshot

`func (o *PostSpec) SetBaseSnapshot(v string)`

SetBaseSnapshot sets BaseSnapshot field to given value.

### HasBaseSnapshot

`func (o *PostSpec) HasBaseSnapshot() bool`

HasBaseSnapshot returns a boolean if a field has been set.

### GetCategories

`func (o *PostSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PostSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PostSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PostSpec) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCover

`func (o *PostSpec) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *PostSpec) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *PostSpec) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *PostSpec) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetDeleted

`func (o *PostSpec) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PostSpec) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PostSpec) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetExcerpt

`func (o *PostSpec) GetExcerpt() Excerpt`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *PostSpec) GetExcerptOk() (*Excerpt, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *PostSpec) SetExcerpt(v Excerpt)`

SetExcerpt sets Excerpt field to given value.


### GetHeadSnapshot

`func (o *PostSpec) GetHeadSnapshot() string`

GetHeadSnapshot returns the HeadSnapshot field if non-nil, zero value otherwise.

### GetHeadSnapshotOk

`func (o *PostSpec) GetHeadSnapshotOk() (*string, bool)`

GetHeadSnapshotOk returns a tuple with the HeadSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSnapshot

`func (o *PostSpec) SetHeadSnapshot(v string)`

SetHeadSnapshot sets HeadSnapshot field to given value.

### HasHeadSnapshot

`func (o *PostSpec) HasHeadSnapshot() bool`

HasHeadSnapshot returns a boolean if a field has been set.

### GetHtmlMetas

`func (o *PostSpec) GetHtmlMetas() []map[string]string`

GetHtmlMetas returns the HtmlMetas field if non-nil, zero value otherwise.

### GetHtmlMetasOk

`func (o *PostSpec) GetHtmlMetasOk() (*[]map[string]string, bool)`

GetHtmlMetasOk returns a tuple with the HtmlMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlMetas

`func (o *PostSpec) SetHtmlMetas(v []map[string]string)`

SetHtmlMetas sets HtmlMetas field to given value.

### HasHtmlMetas

`func (o *PostSpec) HasHtmlMetas() bool`

HasHtmlMetas returns a boolean if a field has been set.

### GetOwner

`func (o *PostSpec) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PostSpec) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PostSpec) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PostSpec) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPinned

`func (o *PostSpec) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *PostSpec) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *PostSpec) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetPriority

`func (o *PostSpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PostSpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PostSpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPublish

`func (o *PostSpec) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *PostSpec) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *PostSpec) SetPublish(v bool)`

SetPublish sets Publish field to given value.


### GetPublishTime

`func (o *PostSpec) GetPublishTime() time.Time`

GetPublishTime returns the PublishTime field if non-nil, zero value otherwise.

### GetPublishTimeOk

`func (o *PostSpec) GetPublishTimeOk() (*time.Time, bool)`

GetPublishTimeOk returns a tuple with the PublishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTime

`func (o *PostSpec) SetPublishTime(v time.Time)`

SetPublishTime sets PublishTime field to given value.

### HasPublishTime

`func (o *PostSpec) HasPublishTime() bool`

HasPublishTime returns a boolean if a field has been set.

### GetReleaseSnapshot

`func (o *PostSpec) GetReleaseSnapshot() string`

GetReleaseSnapshot returns the ReleaseSnapshot field if non-nil, zero value otherwise.

### GetReleaseSnapshotOk

`func (o *PostSpec) GetReleaseSnapshotOk() (*string, bool)`

GetReleaseSnapshotOk returns a tuple with the ReleaseSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSnapshot

`func (o *PostSpec) SetReleaseSnapshot(v string)`

SetReleaseSnapshot sets ReleaseSnapshot field to given value.

### HasReleaseSnapshot

`func (o *PostSpec) HasReleaseSnapshot() bool`

HasReleaseSnapshot returns a boolean if a field has been set.

### GetSlug

`func (o *PostSpec) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PostSpec) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PostSpec) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTags

`func (o *PostSpec) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostSpec) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostSpec) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *PostSpec) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PostSpec) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PostSpec) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PostSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *PostSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostSpec) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVisible

`func (o *PostSpec) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *PostSpec) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *PostSpec) SetVisible(v string)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


