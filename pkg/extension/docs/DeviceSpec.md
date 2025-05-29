# DeviceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** |  | 
**LastAccessedTime** | Pointer to **time.Time** |  | [optional] 
**LastAuthenticatedTime** | Pointer to **time.Time** |  | [optional] 
**PrincipalName** | **string** |  | 
**RememberMeSeriesId** | Pointer to **string** |  | [optional] 
**SessionId** | **string** |  | 
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceSpec

`func NewDeviceSpec(ipAddress string, principalName string, sessionId string, ) *DeviceSpec`

NewDeviceSpec instantiates a new DeviceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSpecWithDefaults

`func NewDeviceSpecWithDefaults() *DeviceSpec`

NewDeviceSpecWithDefaults instantiates a new DeviceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *DeviceSpec) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DeviceSpec) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DeviceSpec) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetLastAccessedTime

`func (o *DeviceSpec) GetLastAccessedTime() time.Time`

GetLastAccessedTime returns the LastAccessedTime field if non-nil, zero value otherwise.

### GetLastAccessedTimeOk

`func (o *DeviceSpec) GetLastAccessedTimeOk() (*time.Time, bool)`

GetLastAccessedTimeOk returns a tuple with the LastAccessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedTime

`func (o *DeviceSpec) SetLastAccessedTime(v time.Time)`

SetLastAccessedTime sets LastAccessedTime field to given value.

### HasLastAccessedTime

`func (o *DeviceSpec) HasLastAccessedTime() bool`

HasLastAccessedTime returns a boolean if a field has been set.

### GetLastAuthenticatedTime

`func (o *DeviceSpec) GetLastAuthenticatedTime() time.Time`

GetLastAuthenticatedTime returns the LastAuthenticatedTime field if non-nil, zero value otherwise.

### GetLastAuthenticatedTimeOk

`func (o *DeviceSpec) GetLastAuthenticatedTimeOk() (*time.Time, bool)`

GetLastAuthenticatedTimeOk returns a tuple with the LastAuthenticatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthenticatedTime

`func (o *DeviceSpec) SetLastAuthenticatedTime(v time.Time)`

SetLastAuthenticatedTime sets LastAuthenticatedTime field to given value.

### HasLastAuthenticatedTime

`func (o *DeviceSpec) HasLastAuthenticatedTime() bool`

HasLastAuthenticatedTime returns a boolean if a field has been set.

### GetPrincipalName

`func (o *DeviceSpec) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *DeviceSpec) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *DeviceSpec) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.


### GetRememberMeSeriesId

`func (o *DeviceSpec) GetRememberMeSeriesId() string`

GetRememberMeSeriesId returns the RememberMeSeriesId field if non-nil, zero value otherwise.

### GetRememberMeSeriesIdOk

`func (o *DeviceSpec) GetRememberMeSeriesIdOk() (*string, bool)`

GetRememberMeSeriesIdOk returns a tuple with the RememberMeSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMeSeriesId

`func (o *DeviceSpec) SetRememberMeSeriesId(v string)`

SetRememberMeSeriesId sets RememberMeSeriesId field to given value.

### HasRememberMeSeriesId

`func (o *DeviceSpec) HasRememberMeSeriesId() bool`

HasRememberMeSeriesId returns a boolean if a field has been set.

### GetSessionId

`func (o *DeviceSpec) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DeviceSpec) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DeviceSpec) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetUserAgent

`func (o *DeviceSpec) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *DeviceSpec) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *DeviceSpec) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *DeviceSpec) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


