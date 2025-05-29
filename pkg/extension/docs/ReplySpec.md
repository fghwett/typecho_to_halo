# ReplySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowNotification** | **bool** |  | [default to true]
**Approved** | **bool** |  | [default to false]
**ApprovedTime** | Pointer to **time.Time** |  | [optional] 
**CommentName** | **string** |  | 
**Content** | **string** |  | 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Hidden** | **bool** |  | [default to false]
**IpAddress** | Pointer to **string** |  | [optional] 
**Owner** | [**CommentOwner**](CommentOwner.md) |  | 
**Priority** | **int32** |  | [default to 0]
**QuoteReply** | Pointer to **string** |  | [optional] 
**Raw** | **string** |  | 
**Top** | **bool** |  | [default to false]
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewReplySpec

`func NewReplySpec(allowNotification bool, approved bool, commentName string, content string, hidden bool, owner CommentOwner, priority int32, raw string, top bool, ) *ReplySpec`

NewReplySpec instantiates a new ReplySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplySpecWithDefaults

`func NewReplySpecWithDefaults() *ReplySpec`

NewReplySpecWithDefaults instantiates a new ReplySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowNotification

`func (o *ReplySpec) GetAllowNotification() bool`

GetAllowNotification returns the AllowNotification field if non-nil, zero value otherwise.

### GetAllowNotificationOk

`func (o *ReplySpec) GetAllowNotificationOk() (*bool, bool)`

GetAllowNotificationOk returns a tuple with the AllowNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNotification

`func (o *ReplySpec) SetAllowNotification(v bool)`

SetAllowNotification sets AllowNotification field to given value.


### GetApproved

`func (o *ReplySpec) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ReplySpec) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ReplySpec) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetApprovedTime

`func (o *ReplySpec) GetApprovedTime() time.Time`

GetApprovedTime returns the ApprovedTime field if non-nil, zero value otherwise.

### GetApprovedTimeOk

`func (o *ReplySpec) GetApprovedTimeOk() (*time.Time, bool)`

GetApprovedTimeOk returns a tuple with the ApprovedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedTime

`func (o *ReplySpec) SetApprovedTime(v time.Time)`

SetApprovedTime sets ApprovedTime field to given value.

### HasApprovedTime

`func (o *ReplySpec) HasApprovedTime() bool`

HasApprovedTime returns a boolean if a field has been set.

### GetCommentName

`func (o *ReplySpec) GetCommentName() string`

GetCommentName returns the CommentName field if non-nil, zero value otherwise.

### GetCommentNameOk

`func (o *ReplySpec) GetCommentNameOk() (*string, bool)`

GetCommentNameOk returns a tuple with the CommentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentName

`func (o *ReplySpec) SetCommentName(v string)`

SetCommentName sets CommentName field to given value.


### GetContent

`func (o *ReplySpec) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReplySpec) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReplySpec) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreationTime

`func (o *ReplySpec) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ReplySpec) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ReplySpec) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ReplySpec) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetHidden

`func (o *ReplySpec) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ReplySpec) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ReplySpec) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetIpAddress

`func (o *ReplySpec) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ReplySpec) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ReplySpec) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ReplySpec) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetOwner

`func (o *ReplySpec) GetOwner() CommentOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ReplySpec) GetOwnerOk() (*CommentOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ReplySpec) SetOwner(v CommentOwner)`

SetOwner sets Owner field to given value.


### GetPriority

`func (o *ReplySpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ReplySpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ReplySpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetQuoteReply

`func (o *ReplySpec) GetQuoteReply() string`

GetQuoteReply returns the QuoteReply field if non-nil, zero value otherwise.

### GetQuoteReplyOk

`func (o *ReplySpec) GetQuoteReplyOk() (*string, bool)`

GetQuoteReplyOk returns a tuple with the QuoteReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteReply

`func (o *ReplySpec) SetQuoteReply(v string)`

SetQuoteReply sets QuoteReply field to given value.

### HasQuoteReply

`func (o *ReplySpec) HasQuoteReply() bool`

HasQuoteReply returns a boolean if a field has been set.

### GetRaw

`func (o *ReplySpec) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *ReplySpec) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *ReplySpec) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetTop

`func (o *ReplySpec) GetTop() bool`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *ReplySpec) GetTopOk() (*bool, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *ReplySpec) SetTop(v bool)`

SetTop sets Top field to given value.


### GetUserAgent

`func (o *ReplySpec) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ReplySpec) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ReplySpec) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ReplySpec) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


