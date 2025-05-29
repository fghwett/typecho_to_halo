# ListedSinglePageVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contributors** | Pointer to [**[]ContributorVo**](ContributorVo.md) |  | [optional] 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Owner** | Pointer to [**ContributorVo**](ContributorVo.md) |  | [optional] 
**Spec** | Pointer to [**SinglePageSpec**](SinglePageSpec.md) |  | [optional] 
**Stats** | Pointer to [**StatsVo**](StatsVo.md) |  | [optional] 
**Status** | Pointer to [**SinglePageStatus**](SinglePageStatus.md) |  | [optional] 

## Methods

### NewListedSinglePageVo

`func NewListedSinglePageVo(metadata Metadata, ) *ListedSinglePageVo`

NewListedSinglePageVo instantiates a new ListedSinglePageVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedSinglePageVoWithDefaults

`func NewListedSinglePageVoWithDefaults() *ListedSinglePageVo`

NewListedSinglePageVoWithDefaults instantiates a new ListedSinglePageVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributors

`func (o *ListedSinglePageVo) GetContributors() []ContributorVo`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ListedSinglePageVo) GetContributorsOk() (*[]ContributorVo, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ListedSinglePageVo) SetContributors(v []ContributorVo)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ListedSinglePageVo) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetMetadata

`func (o *ListedSinglePageVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListedSinglePageVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListedSinglePageVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetOwner

`func (o *ListedSinglePageVo) GetOwner() ContributorVo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListedSinglePageVo) GetOwnerOk() (*ContributorVo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListedSinglePageVo) SetOwner(v ContributorVo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListedSinglePageVo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSpec

`func (o *ListedSinglePageVo) GetSpec() SinglePageSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ListedSinglePageVo) GetSpecOk() (*SinglePageSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ListedSinglePageVo) SetSpec(v SinglePageSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ListedSinglePageVo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStats

`func (o *ListedSinglePageVo) GetStats() StatsVo`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListedSinglePageVo) GetStatsOk() (*StatsVo, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListedSinglePageVo) SetStats(v StatsVo)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListedSinglePageVo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ListedSinglePageVo) GetStatus() SinglePageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListedSinglePageVo) GetStatusOk() (*SinglePageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListedSinglePageVo) SetStatus(v SinglePageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListedSinglePageVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


