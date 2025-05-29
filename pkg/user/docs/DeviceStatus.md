# DeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceStatus

`func NewDeviceStatus() *DeviceStatus`

NewDeviceStatus instantiates a new DeviceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStatusWithDefaults

`func NewDeviceStatusWithDefaults() *DeviceStatus`

NewDeviceStatusWithDefaults instantiates a new DeviceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *DeviceStatus) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *DeviceStatus) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *DeviceStatus) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *DeviceStatus) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetOs

`func (o *DeviceStatus) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeviceStatus) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeviceStatus) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *DeviceStatus) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


