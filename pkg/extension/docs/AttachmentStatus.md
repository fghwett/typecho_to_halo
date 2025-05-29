# AttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permalink** | Pointer to **string** | Permalink of attachment. If it is in local storage, the public URL will be set. If it is in s3 storage, the Object URL will be set.  | [optional] 
**Thumbnails** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAttachmentStatus

`func NewAttachmentStatus() *AttachmentStatus`

NewAttachmentStatus instantiates a new AttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentStatusWithDefaults

`func NewAttachmentStatusWithDefaults() *AttachmentStatus`

NewAttachmentStatusWithDefaults instantiates a new AttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermalink

`func (o *AttachmentStatus) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *AttachmentStatus) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *AttachmentStatus) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *AttachmentStatus) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetThumbnails

`func (o *AttachmentStatus) GetThumbnails() map[string]string`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *AttachmentStatus) GetThumbnailsOk() (*map[string]string, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *AttachmentStatus) SetThumbnails(v map[string]string)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *AttachmentStatus) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


