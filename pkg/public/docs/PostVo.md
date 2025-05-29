# PostVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]CategoryVo**](CategoryVo.md) |  | [optional] 
**Content** | Pointer to [**ContentVo**](ContentVo.md) |  | [optional] 
**Contributors** | Pointer to [**[]ContributorVo**](ContributorVo.md) |  | [optional] 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Owner** | Pointer to [**ContributorVo**](ContributorVo.md) |  | [optional] 
**Spec** | Pointer to [**PostSpec**](PostSpec.md) |  | [optional] 
**Stats** | Pointer to [**StatsVo**](StatsVo.md) |  | [optional] 
**Status** | Pointer to [**PostStatus**](PostStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]TagVo**](TagVo.md) |  | [optional] 

## Methods

### NewPostVo

`func NewPostVo(metadata Metadata, ) *PostVo`

NewPostVo instantiates a new PostVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostVoWithDefaults

`func NewPostVoWithDefaults() *PostVo`

NewPostVoWithDefaults instantiates a new PostVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *PostVo) GetCategories() []CategoryVo`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PostVo) GetCategoriesOk() (*[]CategoryVo, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PostVo) SetCategories(v []CategoryVo)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PostVo) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetContent

`func (o *PostVo) GetContent() ContentVo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PostVo) GetContentOk() (*ContentVo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PostVo) SetContent(v ContentVo)`

SetContent sets Content field to given value.

### HasContent

`func (o *PostVo) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContributors

`func (o *PostVo) GetContributors() []ContributorVo`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *PostVo) GetContributorsOk() (*[]ContributorVo, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *PostVo) SetContributors(v []ContributorVo)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *PostVo) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetMetadata

`func (o *PostVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetOwner

`func (o *PostVo) GetOwner() ContributorVo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PostVo) GetOwnerOk() (*ContributorVo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PostVo) SetOwner(v ContributorVo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PostVo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSpec

`func (o *PostVo) GetSpec() PostSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PostVo) GetSpecOk() (*PostSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PostVo) SetSpec(v PostSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PostVo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStats

`func (o *PostVo) GetStats() StatsVo`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *PostVo) GetStatsOk() (*StatsVo, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *PostVo) SetStats(v StatsVo)`

SetStats sets Stats field to given value.

### HasStats

`func (o *PostVo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *PostVo) GetStatus() PostStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostVo) GetStatusOk() (*PostStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostVo) SetStatus(v PostStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PostVo) GetTags() []TagVo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostVo) GetTagsOk() (*[]TagVo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostVo) SetTags(v []TagVo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostVo) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


