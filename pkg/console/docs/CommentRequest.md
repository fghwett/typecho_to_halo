# CommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowNotification** | Pointer to **bool** |  | [optional] [default to false]
**Content** | **string** |  | 
**Owner** | Pointer to [**CommentEmailOwner**](CommentEmailOwner.md) |  | [optional] 
**Raw** | **string** |  | 
**SubjectRef** | [**Ref**](Ref.md) |  | 

## Methods

### NewCommentRequest

`func NewCommentRequest(content string, raw string, subjectRef Ref, ) *CommentRequest`

NewCommentRequest instantiates a new CommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentRequestWithDefaults

`func NewCommentRequestWithDefaults() *CommentRequest`

NewCommentRequestWithDefaults instantiates a new CommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowNotification

`func (o *CommentRequest) GetAllowNotification() bool`

GetAllowNotification returns the AllowNotification field if non-nil, zero value otherwise.

### GetAllowNotificationOk

`func (o *CommentRequest) GetAllowNotificationOk() (*bool, bool)`

GetAllowNotificationOk returns a tuple with the AllowNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNotification

`func (o *CommentRequest) SetAllowNotification(v bool)`

SetAllowNotification sets AllowNotification field to given value.

### HasAllowNotification

`func (o *CommentRequest) HasAllowNotification() bool`

HasAllowNotification returns a boolean if a field has been set.

### GetContent

`func (o *CommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CommentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetOwner

`func (o *CommentRequest) GetOwner() CommentEmailOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CommentRequest) GetOwnerOk() (*CommentEmailOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CommentRequest) SetOwner(v CommentEmailOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CommentRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRaw

`func (o *CommentRequest) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *CommentRequest) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *CommentRequest) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetSubjectRef

`func (o *CommentRequest) GetSubjectRef() Ref`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *CommentRequest) GetSubjectRefOk() (*Ref, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *CommentRequest) SetSubjectRef(v Ref)`

SetSubjectRef sets SubjectRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


