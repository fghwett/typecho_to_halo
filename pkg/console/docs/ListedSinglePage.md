# ListedSinglePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contributors** | [**[]Contributor**](Contributor.md) |  | 
**Owner** | [**Contributor**](Contributor.md) |  | 
**Page** | [**SinglePage**](SinglePage.md) |  | 
**Stats** | [**Stats**](Stats.md) |  | 

## Methods

### NewListedSinglePage

`func NewListedSinglePage(contributors []Contributor, owner Contributor, page SinglePage, stats Stats, ) *ListedSinglePage`

NewListedSinglePage instantiates a new ListedSinglePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedSinglePageWithDefaults

`func NewListedSinglePageWithDefaults() *ListedSinglePage`

NewListedSinglePageWithDefaults instantiates a new ListedSinglePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributors

`func (o *ListedSinglePage) GetContributors() []Contributor`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ListedSinglePage) GetContributorsOk() (*[]Contributor, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ListedSinglePage) SetContributors(v []Contributor)`

SetContributors sets Contributors field to given value.


### GetOwner

`func (o *ListedSinglePage) GetOwner() Contributor`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListedSinglePage) GetOwnerOk() (*Contributor, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListedSinglePage) SetOwner(v Contributor)`

SetOwner sets Owner field to given value.


### GetPage

`func (o *ListedSinglePage) GetPage() SinglePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListedSinglePage) GetPageOk() (*SinglePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListedSinglePage) SetPage(v SinglePage)`

SetPage sets Page field to given value.


### GetStats

`func (o *ListedSinglePage) GetStats() Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListedSinglePage) GetStatsOk() (*Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListedSinglePage) SetStats(v Stats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


