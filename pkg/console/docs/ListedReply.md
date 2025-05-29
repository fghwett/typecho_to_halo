# ListedReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | [**OwnerInfo**](OwnerInfo.md) |  | 
**Reply** | [**Reply**](Reply.md) |  | 
**Stats** | [**CommentStats**](CommentStats.md) |  | 

## Methods

### NewListedReply

`func NewListedReply(owner OwnerInfo, reply Reply, stats CommentStats, ) *ListedReply`

NewListedReply instantiates a new ListedReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedReplyWithDefaults

`func NewListedReplyWithDefaults() *ListedReply`

NewListedReplyWithDefaults instantiates a new ListedReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ListedReply) GetOwner() OwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListedReply) GetOwnerOk() (*OwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListedReply) SetOwner(v OwnerInfo)`

SetOwner sets Owner field to given value.


### GetReply

`func (o *ListedReply) GetReply() Reply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *ListedReply) GetReplyOk() (*Reply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *ListedReply) SetReply(v Reply)`

SetReply sets Reply field to given value.


### GetStats

`func (o *ListedReply) GetStats() CommentStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListedReply) GetStatsOk() (*CommentStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListedReply) SetStats(v CommentStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


