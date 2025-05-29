# SinglePageVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**ContentVo**](ContentVo.md) |  | [optional] 
**Contributors** | Pointer to [**[]ContributorVo**](ContributorVo.md) |  | [optional] 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Owner** | Pointer to [**ContributorVo**](ContributorVo.md) |  | [optional] 
**Spec** | Pointer to [**SinglePageSpec**](SinglePageSpec.md) |  | [optional] 
**Stats** | Pointer to [**StatsVo**](StatsVo.md) |  | [optional] 
**Status** | Pointer to [**SinglePageStatus**](SinglePageStatus.md) |  | [optional] 

## Methods

### NewSinglePageVo

`func NewSinglePageVo(metadata Metadata, ) *SinglePageVo`

NewSinglePageVo instantiates a new SinglePageVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglePageVoWithDefaults

`func NewSinglePageVoWithDefaults() *SinglePageVo`

NewSinglePageVoWithDefaults instantiates a new SinglePageVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SinglePageVo) GetContent() ContentVo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SinglePageVo) GetContentOk() (*ContentVo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SinglePageVo) SetContent(v ContentVo)`

SetContent sets Content field to given value.

### HasContent

`func (o *SinglePageVo) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContributors

`func (o *SinglePageVo) GetContributors() []ContributorVo`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *SinglePageVo) GetContributorsOk() (*[]ContributorVo, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *SinglePageVo) SetContributors(v []ContributorVo)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *SinglePageVo) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetMetadata

`func (o *SinglePageVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SinglePageVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SinglePageVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetOwner

`func (o *SinglePageVo) GetOwner() ContributorVo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SinglePageVo) GetOwnerOk() (*ContributorVo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SinglePageVo) SetOwner(v ContributorVo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SinglePageVo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSpec

`func (o *SinglePageVo) GetSpec() SinglePageSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SinglePageVo) GetSpecOk() (*SinglePageSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SinglePageVo) SetSpec(v SinglePageSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SinglePageVo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStats

`func (o *SinglePageVo) GetStats() StatsVo`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SinglePageVo) GetStatsOk() (*StatsVo, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SinglePageVo) SetStats(v StatsVo)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SinglePageVo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *SinglePageVo) GetStatus() SinglePageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SinglePageVo) GetStatusOk() (*SinglePageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SinglePageVo) SetStatus(v SinglePageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SinglePageVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


