# SubscriptionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Perhaps users need to unsubscribe and interact without receiving notifications again | [optional] 
**Reason** | [**InterestReason**](InterestReason.md) |  | 
**Subscriber** | [**SubscriptionSubscriber**](SubscriptionSubscriber.md) |  | 
**UnsubscribeToken** | **string** | The token to unsubscribe | 

## Methods

### NewSubscriptionSpec

`func NewSubscriptionSpec(reason InterestReason, subscriber SubscriptionSubscriber, unsubscribeToken string, ) *SubscriptionSpec`

NewSubscriptionSpec instantiates a new SubscriptionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionSpecWithDefaults

`func NewSubscriptionSpecWithDefaults() *SubscriptionSpec`

NewSubscriptionSpecWithDefaults instantiates a new SubscriptionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *SubscriptionSpec) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SubscriptionSpec) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SubscriptionSpec) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SubscriptionSpec) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetReason

`func (o *SubscriptionSpec) GetReason() InterestReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubscriptionSpec) GetReasonOk() (*InterestReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubscriptionSpec) SetReason(v InterestReason)`

SetReason sets Reason field to given value.


### GetSubscriber

`func (o *SubscriptionSpec) GetSubscriber() SubscriptionSubscriber`

GetSubscriber returns the Subscriber field if non-nil, zero value otherwise.

### GetSubscriberOk

`func (o *SubscriptionSpec) GetSubscriberOk() (*SubscriptionSubscriber, bool)`

GetSubscriberOk returns a tuple with the Subscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriber

`func (o *SubscriptionSpec) SetSubscriber(v SubscriptionSubscriber)`

SetSubscriber sets Subscriber field to given value.


### GetUnsubscribeToken

`func (o *SubscriptionSpec) GetUnsubscribeToken() string`

GetUnsubscribeToken returns the UnsubscribeToken field if non-nil, zero value otherwise.

### GetUnsubscribeTokenOk

`func (o *SubscriptionSpec) GetUnsubscribeTokenOk() (*string, bool)`

GetUnsubscribeTokenOk returns a tuple with the UnsubscribeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeToken

`func (o *SubscriptionSpec) SetUnsubscribeToken(v string)`

SetUnsubscribeToken sets UnsubscribeToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


