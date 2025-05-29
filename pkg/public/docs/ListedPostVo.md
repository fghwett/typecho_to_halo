# ListedPostVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]CategoryVo**](CategoryVo.md) |  | [optional] 
**Contributors** | Pointer to [**[]ContributorVo**](ContributorVo.md) |  | [optional] 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Owner** | Pointer to [**ContributorVo**](ContributorVo.md) |  | [optional] 
**Spec** | Pointer to [**PostSpec**](PostSpec.md) |  | [optional] 
**Stats** | Pointer to [**StatsVo**](StatsVo.md) |  | [optional] 
**Status** | Pointer to [**PostStatus**](PostStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]TagVo**](TagVo.md) |  | [optional] 

## Methods

### NewListedPostVo

`func NewListedPostVo(metadata Metadata, ) *ListedPostVo`

NewListedPostVo instantiates a new ListedPostVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedPostVoWithDefaults

`func NewListedPostVoWithDefaults() *ListedPostVo`

NewListedPostVoWithDefaults instantiates a new ListedPostVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ListedPostVo) GetCategories() []CategoryVo`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ListedPostVo) GetCategoriesOk() (*[]CategoryVo, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ListedPostVo) SetCategories(v []CategoryVo)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ListedPostVo) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetContributors

`func (o *ListedPostVo) GetContributors() []ContributorVo`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ListedPostVo) GetContributorsOk() (*[]ContributorVo, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ListedPostVo) SetContributors(v []ContributorVo)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ListedPostVo) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetMetadata

`func (o *ListedPostVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListedPostVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListedPostVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetOwner

`func (o *ListedPostVo) GetOwner() ContributorVo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListedPostVo) GetOwnerOk() (*ContributorVo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListedPostVo) SetOwner(v ContributorVo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListedPostVo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSpec

`func (o *ListedPostVo) GetSpec() PostSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ListedPostVo) GetSpecOk() (*PostSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ListedPostVo) SetSpec(v PostSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ListedPostVo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStats

`func (o *ListedPostVo) GetStats() StatsVo`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListedPostVo) GetStatsOk() (*StatsVo, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListedPostVo) SetStats(v StatsVo)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListedPostVo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ListedPostVo) GetStatus() PostStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListedPostVo) GetStatusOk() (*PostStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListedPostVo) SetStatus(v PostStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListedPostVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ListedPostVo) GetTags() []TagVo`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListedPostVo) GetTagsOk() (*[]TagVo, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListedPostVo) SetTags(v []TagVo)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListedPostVo) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


