# NotificationTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonSelector** | Pointer to [**ReasonSelector**](ReasonSelector.md) |  | [optional] 
**Template** | Pointer to [**TemplateContent**](TemplateContent.md) |  | [optional] 

## Methods

### NewNotificationTemplateSpec

`func NewNotificationTemplateSpec() *NotificationTemplateSpec`

NewNotificationTemplateSpec instantiates a new NotificationTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTemplateSpecWithDefaults

`func NewNotificationTemplateSpecWithDefaults() *NotificationTemplateSpec`

NewNotificationTemplateSpecWithDefaults instantiates a new NotificationTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonSelector

`func (o *NotificationTemplateSpec) GetReasonSelector() ReasonSelector`

GetReasonSelector returns the ReasonSelector field if non-nil, zero value otherwise.

### GetReasonSelectorOk

`func (o *NotificationTemplateSpec) GetReasonSelectorOk() (*ReasonSelector, bool)`

GetReasonSelectorOk returns a tuple with the ReasonSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSelector

`func (o *NotificationTemplateSpec) SetReasonSelector(v ReasonSelector)`

SetReasonSelector sets ReasonSelector field to given value.

### HasReasonSelector

`func (o *NotificationTemplateSpec) HasReasonSelector() bool`

HasReasonSelector returns a boolean if a field has been set.

### GetTemplate

`func (o *NotificationTemplateSpec) GetTemplate() TemplateContent`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *NotificationTemplateSpec) GetTemplateOk() (*TemplateContent, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *NotificationTemplateSpec) SetTemplate(v TemplateContent)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *NotificationTemplateSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


