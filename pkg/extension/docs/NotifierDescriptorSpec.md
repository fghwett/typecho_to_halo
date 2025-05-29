# NotifierDescriptorSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**NotifierExtName** | **string** |  | 
**ReceiverSettingRef** | Pointer to [**NotifierSettingRef**](NotifierSettingRef.md) |  | [optional] 
**SenderSettingRef** | Pointer to [**NotifierSettingRef**](NotifierSettingRef.md) |  | [optional] 

## Methods

### NewNotifierDescriptorSpec

`func NewNotifierDescriptorSpec(displayName string, notifierExtName string, ) *NotifierDescriptorSpec`

NewNotifierDescriptorSpec instantiates a new NotifierDescriptorSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifierDescriptorSpecWithDefaults

`func NewNotifierDescriptorSpecWithDefaults() *NotifierDescriptorSpec`

NewNotifierDescriptorSpecWithDefaults instantiates a new NotifierDescriptorSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NotifierDescriptorSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotifierDescriptorSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotifierDescriptorSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotifierDescriptorSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *NotifierDescriptorSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NotifierDescriptorSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NotifierDescriptorSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetNotifierExtName

`func (o *NotifierDescriptorSpec) GetNotifierExtName() string`

GetNotifierExtName returns the NotifierExtName field if non-nil, zero value otherwise.

### GetNotifierExtNameOk

`func (o *NotifierDescriptorSpec) GetNotifierExtNameOk() (*string, bool)`

GetNotifierExtNameOk returns a tuple with the NotifierExtName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierExtName

`func (o *NotifierDescriptorSpec) SetNotifierExtName(v string)`

SetNotifierExtName sets NotifierExtName field to given value.


### GetReceiverSettingRef

`func (o *NotifierDescriptorSpec) GetReceiverSettingRef() NotifierSettingRef`

GetReceiverSettingRef returns the ReceiverSettingRef field if non-nil, zero value otherwise.

### GetReceiverSettingRefOk

`func (o *NotifierDescriptorSpec) GetReceiverSettingRefOk() (*NotifierSettingRef, bool)`

GetReceiverSettingRefOk returns a tuple with the ReceiverSettingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverSettingRef

`func (o *NotifierDescriptorSpec) SetReceiverSettingRef(v NotifierSettingRef)`

SetReceiverSettingRef sets ReceiverSettingRef field to given value.

### HasReceiverSettingRef

`func (o *NotifierDescriptorSpec) HasReceiverSettingRef() bool`

HasReceiverSettingRef returns a boolean if a field has been set.

### GetSenderSettingRef

`func (o *NotifierDescriptorSpec) GetSenderSettingRef() NotifierSettingRef`

GetSenderSettingRef returns the SenderSettingRef field if non-nil, zero value otherwise.

### GetSenderSettingRefOk

`func (o *NotifierDescriptorSpec) GetSenderSettingRefOk() (*NotifierSettingRef, bool)`

GetSenderSettingRefOk returns a tuple with the SenderSettingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSettingRef

`func (o *NotifierDescriptorSpec) SetSenderSettingRef(v NotifierSettingRef)`

SetSenderSettingRef sets SenderSettingRef field to given value.

### HasSenderSettingRef

`func (o *NotifierDescriptorSpec) HasSenderSettingRef() bool`

HasSenderSettingRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


