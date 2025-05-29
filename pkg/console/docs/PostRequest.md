# PostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**ContentUpdateParam**](ContentUpdateParam.md) |  | [optional] 
**Post** | [**Post**](Post.md) |  | 

## Methods

### NewPostRequest

`func NewPostRequest(post Post, ) *PostRequest`

NewPostRequest instantiates a new PostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRequestWithDefaults

`func NewPostRequestWithDefaults() *PostRequest`

NewPostRequestWithDefaults instantiates a new PostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PostRequest) GetContent() ContentUpdateParam`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PostRequest) GetContentOk() (*ContentUpdateParam, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PostRequest) SetContent(v ContentUpdateParam)`

SetContent sets Content field to given value.

### HasContent

`func (o *PostRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPost

`func (o *PostRequest) GetPost() Post`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *PostRequest) GetPostOk() (*Post, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *PostRequest) SetPost(v Post)`

SetPost sets Post field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


