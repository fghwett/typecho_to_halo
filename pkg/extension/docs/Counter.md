# Counter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**ApprovedComment** | Pointer to **int32** |  | [optional] 
**Downvote** | Pointer to **int32** |  | [optional] 
**Kind** | **string** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**TotalComment** | Pointer to **int32** |  | [optional] 
**Upvote** | Pointer to **int32** |  | [optional] 
**Visit** | Pointer to **int32** |  | [optional] 

## Methods

### NewCounter

`func NewCounter(apiVersion string, kind string, metadata Metadata, ) *Counter`

NewCounter instantiates a new Counter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterWithDefaults

`func NewCounterWithDefaults() *Counter`

NewCounterWithDefaults instantiates a new Counter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Counter) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Counter) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Counter) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetApprovedComment

`func (o *Counter) GetApprovedComment() int32`

GetApprovedComment returns the ApprovedComment field if non-nil, zero value otherwise.

### GetApprovedCommentOk

`func (o *Counter) GetApprovedCommentOk() (*int32, bool)`

GetApprovedCommentOk returns a tuple with the ApprovedComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedComment

`func (o *Counter) SetApprovedComment(v int32)`

SetApprovedComment sets ApprovedComment field to given value.

### HasApprovedComment

`func (o *Counter) HasApprovedComment() bool`

HasApprovedComment returns a boolean if a field has been set.

### GetDownvote

`func (o *Counter) GetDownvote() int32`

GetDownvote returns the Downvote field if non-nil, zero value otherwise.

### GetDownvoteOk

`func (o *Counter) GetDownvoteOk() (*int32, bool)`

GetDownvoteOk returns a tuple with the Downvote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownvote

`func (o *Counter) SetDownvote(v int32)`

SetDownvote sets Downvote field to given value.

### HasDownvote

`func (o *Counter) HasDownvote() bool`

HasDownvote returns a boolean if a field has been set.

### GetKind

`func (o *Counter) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Counter) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Counter) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Counter) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Counter) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Counter) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetTotalComment

`func (o *Counter) GetTotalComment() int32`

GetTotalComment returns the TotalComment field if non-nil, zero value otherwise.

### GetTotalCommentOk

`func (o *Counter) GetTotalCommentOk() (*int32, bool)`

GetTotalCommentOk returns a tuple with the TotalComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalComment

`func (o *Counter) SetTotalComment(v int32)`

SetTotalComment sets TotalComment field to given value.

### HasTotalComment

`func (o *Counter) HasTotalComment() bool`

HasTotalComment returns a boolean if a field has been set.

### GetUpvote

`func (o *Counter) GetUpvote() int32`

GetUpvote returns the Upvote field if non-nil, zero value otherwise.

### GetUpvoteOk

`func (o *Counter) GetUpvoteOk() (*int32, bool)`

GetUpvoteOk returns a tuple with the Upvote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpvote

`func (o *Counter) SetUpvote(v int32)`

SetUpvote sets Upvote field to given value.

### HasUpvote

`func (o *Counter) HasUpvote() bool`

HasUpvote returns a boolean if a field has been set.

### GetVisit

`func (o *Counter) GetVisit() int32`

GetVisit returns the Visit field if non-nil, zero value otherwise.

### GetVisitOk

`func (o *Counter) GetVisitOk() (*int32, bool)`

GetVisitOk returns a tuple with the Visit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisit

`func (o *Counter) SetVisit(v int32)`

SetVisit sets Visit field to given value.

### HasVisit

`func (o *Counter) HasVisit() bool`

HasVisit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


