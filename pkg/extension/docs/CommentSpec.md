# CommentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowNotification** | **bool** |  | [default to true]
**Approved** | **bool** |  | [default to false]
**ApprovedTime** | Pointer to **time.Time** |  | [optional] 
**Content** | **string** |  | 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Hidden** | **bool** |  | [default to false]
**IpAddress** | Pointer to **string** |  | [optional] 
**LastReadTime** | Pointer to **time.Time** |  | [optional] 
**Owner** | [**CommentOwner**](CommentOwner.md) |  | 
**Priority** | **int32** |  | [default to 0]
**Raw** | **string** |  | 
**SubjectRef** | [**Ref**](Ref.md) |  | 
**Top** | **bool** |  | [default to false]
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewCommentSpec

`func NewCommentSpec(allowNotification bool, approved bool, content string, hidden bool, owner CommentOwner, priority int32, raw string, subjectRef Ref, top bool, ) *CommentSpec`

NewCommentSpec instantiates a new CommentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentSpecWithDefaults

`func NewCommentSpecWithDefaults() *CommentSpec`

NewCommentSpecWithDefaults instantiates a new CommentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowNotification

`func (o *CommentSpec) GetAllowNotification() bool`

GetAllowNotification returns the AllowNotification field if non-nil, zero value otherwise.

### GetAllowNotificationOk

`func (o *CommentSpec) GetAllowNotificationOk() (*bool, bool)`

GetAllowNotificationOk returns a tuple with the AllowNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNotification

`func (o *CommentSpec) SetAllowNotification(v bool)`

SetAllowNotification sets AllowNotification field to given value.


### GetApproved

`func (o *CommentSpec) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *CommentSpec) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *CommentSpec) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetApprovedTime

`func (o *CommentSpec) GetApprovedTime() time.Time`

GetApprovedTime returns the ApprovedTime field if non-nil, zero value otherwise.

### GetApprovedTimeOk

`func (o *CommentSpec) GetApprovedTimeOk() (*time.Time, bool)`

GetApprovedTimeOk returns a tuple with the ApprovedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedTime

`func (o *CommentSpec) SetApprovedTime(v time.Time)`

SetApprovedTime sets ApprovedTime field to given value.

### HasApprovedTime

`func (o *CommentSpec) HasApprovedTime() bool`

HasApprovedTime returns a boolean if a field has been set.

### GetContent

`func (o *CommentSpec) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CommentSpec) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CommentSpec) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreationTime

`func (o *CommentSpec) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CommentSpec) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CommentSpec) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CommentSpec) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetHidden

`func (o *CommentSpec) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *CommentSpec) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *CommentSpec) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetIpAddress

`func (o *CommentSpec) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CommentSpec) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CommentSpec) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CommentSpec) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastReadTime

`func (o *CommentSpec) GetLastReadTime() time.Time`

GetLastReadTime returns the LastReadTime field if non-nil, zero value otherwise.

### GetLastReadTimeOk

`func (o *CommentSpec) GetLastReadTimeOk() (*time.Time, bool)`

GetLastReadTimeOk returns a tuple with the LastReadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadTime

`func (o *CommentSpec) SetLastReadTime(v time.Time)`

SetLastReadTime sets LastReadTime field to given value.

### HasLastReadTime

`func (o *CommentSpec) HasLastReadTime() bool`

HasLastReadTime returns a boolean if a field has been set.

### GetOwner

`func (o *CommentSpec) GetOwner() CommentOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CommentSpec) GetOwnerOk() (*CommentOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CommentSpec) SetOwner(v CommentOwner)`

SetOwner sets Owner field to given value.


### GetPriority

`func (o *CommentSpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CommentSpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CommentSpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetRaw

`func (o *CommentSpec) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *CommentSpec) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *CommentSpec) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetSubjectRef

`func (o *CommentSpec) GetSubjectRef() Ref`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *CommentSpec) GetSubjectRefOk() (*Ref, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *CommentSpec) SetSubjectRef(v Ref)`

SetSubjectRef sets SubjectRef field to given value.


### GetTop

`func (o *CommentSpec) GetTop() bool`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *CommentSpec) GetTopOk() (*bool, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *CommentSpec) SetTop(v bool)`

SetTop sets Top field to given value.


### GetUserAgent

`func (o *CommentSpec) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CommentSpec) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CommentSpec) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CommentSpec) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


