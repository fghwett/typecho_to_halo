# ReplyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowNotification** | Pointer to **bool** |  | [optional] [default to false]
**Content** | **string** |  | 
**Owner** | Pointer to [**CommentEmailOwner**](CommentEmailOwner.md) |  | [optional] 
**QuoteReply** | Pointer to **string** |  | [optional] 
**Raw** | **string** |  | 

## Methods

### NewReplyRequest

`func NewReplyRequest(content string, raw string, ) *ReplyRequest`

NewReplyRequest instantiates a new ReplyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyRequestWithDefaults

`func NewReplyRequestWithDefaults() *ReplyRequest`

NewReplyRequestWithDefaults instantiates a new ReplyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowNotification

`func (o *ReplyRequest) GetAllowNotification() bool`

GetAllowNotification returns the AllowNotification field if non-nil, zero value otherwise.

### GetAllowNotificationOk

`func (o *ReplyRequest) GetAllowNotificationOk() (*bool, bool)`

GetAllowNotificationOk returns a tuple with the AllowNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNotification

`func (o *ReplyRequest) SetAllowNotification(v bool)`

SetAllowNotification sets AllowNotification field to given value.

### HasAllowNotification

`func (o *ReplyRequest) HasAllowNotification() bool`

HasAllowNotification returns a boolean if a field has been set.

### GetContent

`func (o *ReplyRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReplyRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReplyRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetOwner

`func (o *ReplyRequest) GetOwner() CommentEmailOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ReplyRequest) GetOwnerOk() (*CommentEmailOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ReplyRequest) SetOwner(v CommentEmailOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ReplyRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetQuoteReply

`func (o *ReplyRequest) GetQuoteReply() string`

GetQuoteReply returns the QuoteReply field if non-nil, zero value otherwise.

### GetQuoteReplyOk

`func (o *ReplyRequest) GetQuoteReplyOk() (*string, bool)`

GetQuoteReplyOk returns a tuple with the QuoteReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteReply

`func (o *ReplyRequest) SetQuoteReply(v string)`

SetQuoteReply sets QuoteReply field to given value.

### HasQuoteReply

`func (o *ReplyRequest) HasQuoteReply() bool`

HasQuoteReply returns a boolean if a field has been set.

### GetRaw

`func (o *ReplyRequest) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *ReplyRequest) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *ReplyRequest) SetRaw(v string)`

SetRaw sets Raw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


