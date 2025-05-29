# SinglePageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowComment** | **bool** |  | [default to true]
**BaseSnapshot** | Pointer to **string** |  | [optional] 
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
**Template** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Visible** | **string** |  | [default to "PUBLIC"]

## Methods

### NewSinglePageSpec

`func NewSinglePageSpec(allowComment bool, deleted bool, excerpt Excerpt, pinned bool, priority int32, publish bool, slug string, title string, visible string, ) *SinglePageSpec`

NewSinglePageSpec instantiates a new SinglePageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglePageSpecWithDefaults

`func NewSinglePageSpecWithDefaults() *SinglePageSpec`

NewSinglePageSpecWithDefaults instantiates a new SinglePageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowComment

`func (o *SinglePageSpec) GetAllowComment() bool`

GetAllowComment returns the AllowComment field if non-nil, zero value otherwise.

### GetAllowCommentOk

`func (o *SinglePageSpec) GetAllowCommentOk() (*bool, bool)`

GetAllowCommentOk returns a tuple with the AllowComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowComment

`func (o *SinglePageSpec) SetAllowComment(v bool)`

SetAllowComment sets AllowComment field to given value.


### GetBaseSnapshot

`func (o *SinglePageSpec) GetBaseSnapshot() string`

GetBaseSnapshot returns the BaseSnapshot field if non-nil, zero value otherwise.

### GetBaseSnapshotOk

`func (o *SinglePageSpec) GetBaseSnapshotOk() (*string, bool)`

GetBaseSnapshotOk returns a tuple with the BaseSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshot

`func (o *SinglePageSpec) SetBaseSnapshot(v string)`

SetBaseSnapshot sets BaseSnapshot field to given value.

### HasBaseSnapshot

`func (o *SinglePageSpec) HasBaseSnapshot() bool`

HasBaseSnapshot returns a boolean if a field has been set.

### GetCover

`func (o *SinglePageSpec) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *SinglePageSpec) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *SinglePageSpec) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *SinglePageSpec) HasCover() bool`

HasCover returns a boolean if a field has been set.

### GetDeleted

`func (o *SinglePageSpec) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SinglePageSpec) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SinglePageSpec) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetExcerpt

`func (o *SinglePageSpec) GetExcerpt() Excerpt`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *SinglePageSpec) GetExcerptOk() (*Excerpt, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *SinglePageSpec) SetExcerpt(v Excerpt)`

SetExcerpt sets Excerpt field to given value.


### GetHeadSnapshot

`func (o *SinglePageSpec) GetHeadSnapshot() string`

GetHeadSnapshot returns the HeadSnapshot field if non-nil, zero value otherwise.

### GetHeadSnapshotOk

`func (o *SinglePageSpec) GetHeadSnapshotOk() (*string, bool)`

GetHeadSnapshotOk returns a tuple with the HeadSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSnapshot

`func (o *SinglePageSpec) SetHeadSnapshot(v string)`

SetHeadSnapshot sets HeadSnapshot field to given value.

### HasHeadSnapshot

`func (o *SinglePageSpec) HasHeadSnapshot() bool`

HasHeadSnapshot returns a boolean if a field has been set.

### GetHtmlMetas

`func (o *SinglePageSpec) GetHtmlMetas() []map[string]string`

GetHtmlMetas returns the HtmlMetas field if non-nil, zero value otherwise.

### GetHtmlMetasOk

`func (o *SinglePageSpec) GetHtmlMetasOk() (*[]map[string]string, bool)`

GetHtmlMetasOk returns a tuple with the HtmlMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlMetas

`func (o *SinglePageSpec) SetHtmlMetas(v []map[string]string)`

SetHtmlMetas sets HtmlMetas field to given value.

### HasHtmlMetas

`func (o *SinglePageSpec) HasHtmlMetas() bool`

HasHtmlMetas returns a boolean if a field has been set.

### GetOwner

`func (o *SinglePageSpec) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SinglePageSpec) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SinglePageSpec) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SinglePageSpec) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPinned

`func (o *SinglePageSpec) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *SinglePageSpec) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *SinglePageSpec) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetPriority

`func (o *SinglePageSpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SinglePageSpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SinglePageSpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPublish

`func (o *SinglePageSpec) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *SinglePageSpec) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *SinglePageSpec) SetPublish(v bool)`

SetPublish sets Publish field to given value.


### GetPublishTime

`func (o *SinglePageSpec) GetPublishTime() time.Time`

GetPublishTime returns the PublishTime field if non-nil, zero value otherwise.

### GetPublishTimeOk

`func (o *SinglePageSpec) GetPublishTimeOk() (*time.Time, bool)`

GetPublishTimeOk returns a tuple with the PublishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTime

`func (o *SinglePageSpec) SetPublishTime(v time.Time)`

SetPublishTime sets PublishTime field to given value.

### HasPublishTime

`func (o *SinglePageSpec) HasPublishTime() bool`

HasPublishTime returns a boolean if a field has been set.

### GetReleaseSnapshot

`func (o *SinglePageSpec) GetReleaseSnapshot() string`

GetReleaseSnapshot returns the ReleaseSnapshot field if non-nil, zero value otherwise.

### GetReleaseSnapshotOk

`func (o *SinglePageSpec) GetReleaseSnapshotOk() (*string, bool)`

GetReleaseSnapshotOk returns a tuple with the ReleaseSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSnapshot

`func (o *SinglePageSpec) SetReleaseSnapshot(v string)`

SetReleaseSnapshot sets ReleaseSnapshot field to given value.

### HasReleaseSnapshot

`func (o *SinglePageSpec) HasReleaseSnapshot() bool`

HasReleaseSnapshot returns a boolean if a field has been set.

### GetSlug

`func (o *SinglePageSpec) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SinglePageSpec) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SinglePageSpec) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTemplate

`func (o *SinglePageSpec) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SinglePageSpec) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SinglePageSpec) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SinglePageSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTitle

`func (o *SinglePageSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SinglePageSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SinglePageSpec) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVisible

`func (o *SinglePageSpec) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *SinglePageSpec) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *SinglePageSpec) SetVisible(v string)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


