# ListedComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | [**Comment**](Comment.md) |  | 
**Owner** | [**OwnerInfo**](OwnerInfo.md) |  | 
**Stats** | [**CommentStats**](CommentStats.md) |  | 
**Subject** | Pointer to [**Extension**](Extension.md) |  | [optional] 

## Methods

### NewListedComment

`func NewListedComment(comment Comment, owner OwnerInfo, stats CommentStats, ) *ListedComment`

NewListedComment instantiates a new ListedComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedCommentWithDefaults

`func NewListedCommentWithDefaults() *ListedComment`

NewListedCommentWithDefaults instantiates a new ListedComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ListedComment) GetComment() Comment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ListedComment) GetCommentOk() (*Comment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ListedComment) SetComment(v Comment)`

SetComment sets Comment field to given value.


### GetOwner

`func (o *ListedComment) GetOwner() OwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListedComment) GetOwnerOk() (*OwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListedComment) SetOwner(v OwnerInfo)`

SetOwner sets Owner field to given value.


### GetStats

`func (o *ListedComment) GetStats() CommentStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListedComment) GetStatsOk() (*CommentStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListedComment) SetStats(v CommentStats)`

SetStats sets Stats field to given value.


### GetSubject

`func (o *ListedComment) GetSubject() Extension`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ListedComment) GetSubjectOk() (*Extension, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ListedComment) SetSubject(v Extension)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ListedComment) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


