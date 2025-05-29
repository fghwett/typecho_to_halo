# PostStatus

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

### NewPostStatus

`func NewPostStatus() *PostStatus`

NewPostStatus instantiates a new PostStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostStatusWithDefaults

`func NewPostStatusWithDefaults() *PostStatus`

NewPostStatusWithDefaults instantiates a new PostStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentsCount

`func (o *PostStatus) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *PostStatus) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *PostStatus) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *PostStatus) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetConditions

`func (o *PostStatus) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PostStatus) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PostStatus) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PostStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContributors

`func (o *PostStatus) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *PostStatus) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *PostStatus) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *PostStatus) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetExcerpt

`func (o *PostStatus) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *PostStatus) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *PostStatus) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *PostStatus) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetHideFromList

`func (o *PostStatus) GetHideFromList() bool`

GetHideFromList returns the HideFromList field if non-nil, zero value otherwise.

### GetHideFromListOk

`func (o *PostStatus) GetHideFromListOk() (*bool, bool)`

GetHideFromListOk returns a tuple with the HideFromList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromList

`func (o *PostStatus) SetHideFromList(v bool)`

SetHideFromList sets HideFromList field to given value.

### HasHideFromList

`func (o *PostStatus) HasHideFromList() bool`

HasHideFromList returns a boolean if a field has been set.

### GetInProgress

`func (o *PostStatus) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *PostStatus) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *PostStatus) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *PostStatus) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetLastModifyTime

`func (o *PostStatus) GetLastModifyTime() time.Time`

GetLastModifyTime returns the LastModifyTime field if non-nil, zero value otherwise.

### GetLastModifyTimeOk

`func (o *PostStatus) GetLastModifyTimeOk() (*time.Time, bool)`

GetLastModifyTimeOk returns a tuple with the LastModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifyTime

`func (o *PostStatus) SetLastModifyTime(v time.Time)`

SetLastModifyTime sets LastModifyTime field to given value.

### HasLastModifyTime

`func (o *PostStatus) HasLastModifyTime() bool`

HasLastModifyTime returns a boolean if a field has been set.

### GetObservedVersion

`func (o *PostStatus) GetObservedVersion() int64`

GetObservedVersion returns the ObservedVersion field if non-nil, zero value otherwise.

### GetObservedVersionOk

`func (o *PostStatus) GetObservedVersionOk() (*int64, bool)`

GetObservedVersionOk returns a tuple with the ObservedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedVersion

`func (o *PostStatus) SetObservedVersion(v int64)`

SetObservedVersion sets ObservedVersion field to given value.

### HasObservedVersion

`func (o *PostStatus) HasObservedVersion() bool`

HasObservedVersion returns a boolean if a field has been set.

### GetPermalink

`func (o *PostStatus) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *PostStatus) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *PostStatus) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *PostStatus) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetPhase

`func (o *PostStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PostStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PostStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PostStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


