# SinglePageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentsCount** | Pointer to **int32** |  | [optional] 
**Conditions** | Pointer to  |  | [optional] 
**Contributors** | Pointer to **[]string** |  | [optional] 
**Excerpt** | Pointer to **string** |  | [optional] 
**HideFromList** | Pointer to **bool** |  | [optional] 
**InProgress** | Pointer to **bool** |  | [optional] 
**LastModifyTime** | Pointer to **time.Time** |  | [optional] 
**ObservedVersion** | Pointer to **int64** |  | [optional] 
**Permalink** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 

## Methods

### NewSinglePageStatus

`func NewSinglePageStatus() *SinglePageStatus`

NewSinglePageStatus instantiates a new SinglePageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglePageStatusWithDefaults

`func NewSinglePageStatusWithDefaults() *SinglePageStatus`

NewSinglePageStatusWithDefaults instantiates a new SinglePageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentsCount

`func (o *SinglePageStatus) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *SinglePageStatus) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *SinglePageStatus) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *SinglePageStatus) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetConditions

`func (o *SinglePageStatus) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SinglePageStatus) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SinglePageStatus) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SinglePageStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContributors

`func (o *SinglePageStatus) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *SinglePageStatus) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *SinglePageStatus) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *SinglePageStatus) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetExcerpt

`func (o *SinglePageStatus) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *SinglePageStatus) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *SinglePageStatus) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *SinglePageStatus) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetHideFromList

`func (o *SinglePageStatus) GetHideFromList() bool`

GetHideFromList returns the HideFromList field if non-nil, zero value otherwise.

### GetHideFromListOk

`func (o *SinglePageStatus) GetHideFromListOk() (*bool, bool)`

GetHideFromListOk returns a tuple with the HideFromList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromList

`func (o *SinglePageStatus) SetHideFromList(v bool)`

SetHideFromList sets HideFromList field to given value.

### HasHideFromList

`func (o *SinglePageStatus) HasHideFromList() bool`

HasHideFromList returns a boolean if a field has been set.

### GetInProgress

`func (o *SinglePageStatus) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *SinglePageStatus) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *SinglePageStatus) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *SinglePageStatus) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetLastModifyTime

`func (o *SinglePageStatus) GetLastModifyTime() time.Time`

GetLastModifyTime returns the LastModifyTime field if non-nil, zero value otherwise.

### GetLastModifyTimeOk

`func (o *SinglePageStatus) GetLastModifyTimeOk() (*time.Time, bool)`

GetLastModifyTimeOk returns a tuple with the LastModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifyTime

`func (o *SinglePageStatus) SetLastModifyTime(v time.Time)`

SetLastModifyTime sets LastModifyTime field to given value.

### HasLastModifyTime

`func (o *SinglePageStatus) HasLastModifyTime() bool`

HasLastModifyTime returns a boolean if a field has been set.

### GetObservedVersion

`func (o *SinglePageStatus) GetObservedVersion() int64`

GetObservedVersion returns the ObservedVersion field if non-nil, zero value otherwise.

### GetObservedVersionOk

`func (o *SinglePageStatus) GetObservedVersionOk() (*int64, bool)`

GetObservedVersionOk returns a tuple with the ObservedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedVersion

`func (o *SinglePageStatus) SetObservedVersion(v int64)`

SetObservedVersion sets ObservedVersion field to given value.

### HasObservedVersion

`func (o *SinglePageStatus) HasObservedVersion() bool`

HasObservedVersion returns a boolean if a field has been set.

### GetPermalink

`func (o *SinglePageStatus) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *SinglePageStatus) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *SinglePageStatus) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *SinglePageStatus) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetPhase

`func (o *SinglePageStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SinglePageStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SinglePageStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SinglePageStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


