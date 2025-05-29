# ListedPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | [**[]Category**](Category.md) |  | 
**Contributors** | [**[]Contributor**](Contributor.md) |  | 
**Owner** | [**Contributor**](Contributor.md) |  | 
**Post** | [**Post**](Post.md) |  | 
**Stats** | [**Stats**](Stats.md) |  | 
**Tags** | [**[]Tag**](Tag.md) |  | 

## Methods

### NewListedPost

`func NewListedPost(categories []Category, contributors []Contributor, owner Contributor, post Post, stats Stats, tags []Tag, ) *ListedPost`

NewListedPost instantiates a new ListedPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedPostWithDefaults

`func NewListedPostWithDefaults() *ListedPost`

NewListedPostWithDefaults instantiates a new ListedPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ListedPost) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ListedPost) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ListedPost) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetContributors

`func (o *ListedPost) GetContributors() []Contributor`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ListedPost) GetContributorsOk() (*[]Contributor, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ListedPost) SetContributors(v []Contributor)`

SetContributors sets Contributors field to given value.


### GetOwner

`func (o *ListedPost) GetOwner() Contributor`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListedPost) GetOwnerOk() (*Contributor, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListedPost) SetOwner(v Contributor)`

SetOwner sets Owner field to given value.


### GetPost

`func (o *ListedPost) GetPost() Post`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *ListedPost) GetPostOk() (*Post, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *ListedPost) SetPost(v Post)`

SetPost sets Post field to given value.


### GetStats

`func (o *ListedPost) GetStats() Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListedPost) GetStatsOk() (*Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListedPost) SetStats(v Stats)`

SetStats sets Stats field to given value.


### GetTags

`func (o *ListedPost) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListedPost) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListedPost) SetTags(v []Tag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


