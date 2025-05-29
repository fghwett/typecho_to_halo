# NotificationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HtmlContent** | **string** |  | 
**LastReadAt** | Pointer to **time.Time** |  | [optional] 
**RawContent** | **string** |  | 
**Reason** | **string** | The name of reason | 
**Recipient** | **string** | The name of user | 
**Title** | **string** |  | 
**Unread** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotificationSpec

`func NewNotificationSpec(htmlContent string, rawContent string, reason string, recipient string, title string, ) *NotificationSpec`

NewNotificationSpec instantiates a new NotificationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSpecWithDefaults

`func NewNotificationSpecWithDefaults() *NotificationSpec`

NewNotificationSpecWithDefaults instantiates a new NotificationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHtmlContent

`func (o *NotificationSpec) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *NotificationSpec) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *NotificationSpec) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.


### GetLastReadAt

`func (o *NotificationSpec) GetLastReadAt() time.Time`

GetLastReadAt returns the LastReadAt field if non-nil, zero value otherwise.

### GetLastReadAtOk

`func (o *NotificationSpec) GetLastReadAtOk() (*time.Time, bool)`

GetLastReadAtOk returns a tuple with the LastReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadAt

`func (o *NotificationSpec) SetLastReadAt(v time.Time)`

SetLastReadAt sets LastReadAt field to given value.

### HasLastReadAt

`func (o *NotificationSpec) HasLastReadAt() bool`

HasLastReadAt returns a boolean if a field has been set.

### GetRawContent

`func (o *NotificationSpec) GetRawContent() string`

GetRawContent returns the RawContent field if non-nil, zero value otherwise.

### GetRawContentOk

`func (o *NotificationSpec) GetRawContentOk() (*string, bool)`

GetRawContentOk returns a tuple with the RawContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawContent

`func (o *NotificationSpec) SetRawContent(v string)`

SetRawContent sets RawContent field to given value.


### GetReason

`func (o *NotificationSpec) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotificationSpec) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotificationSpec) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRecipient

`func (o *NotificationSpec) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NotificationSpec) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NotificationSpec) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetTitle

`func (o *NotificationSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationSpec) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUnread

`func (o *NotificationSpec) GetUnread() bool`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *NotificationSpec) GetUnreadOk() (*bool, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *NotificationSpec) SetUnread(v bool)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *NotificationSpec) HasUnread() bool`

HasUnread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


