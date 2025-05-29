# UserDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**CurrentDevice** | **bool** |  | 
**Device** | [**Device**](Device.md) |  | 

## Methods

### NewUserDevice

`func NewUserDevice(active bool, currentDevice bool, device Device, ) *UserDevice`

NewUserDevice instantiates a new UserDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDeviceWithDefaults

`func NewUserDeviceWithDefaults() *UserDevice`

NewUserDeviceWithDefaults instantiates a new UserDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UserDevice) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserDevice) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserDevice) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCurrentDevice

`func (o *UserDevice) GetCurrentDevice() bool`

GetCurrentDevice returns the CurrentDevice field if non-nil, zero value otherwise.

### GetCurrentDeviceOk

`func (o *UserDevice) GetCurrentDeviceOk() (*bool, bool)`

GetCurrentDeviceOk returns a tuple with the CurrentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDevice

`func (o *UserDevice) SetCurrentDevice(v bool)`

SetCurrentDevice sets CurrentDevice field to given value.


### GetDevice

`func (o *UserDevice) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *UserDevice) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *UserDevice) SetDevice(v Device)`

SetDevice sets Device field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


