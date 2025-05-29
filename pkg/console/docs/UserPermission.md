# UserPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**[]Role**](Role.md) |  | 
**Roles** | [**[]Role**](Role.md) |  | 
**UiPermissions** | **[]string** |  | 

## Methods

### NewUserPermission

`func NewUserPermission(permissions []Role, roles []Role, uiPermissions []string, ) *UserPermission`

NewUserPermission instantiates a new UserPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionWithDefaults

`func NewUserPermissionWithDefaults() *UserPermission`

NewUserPermissionWithDefaults instantiates a new UserPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *UserPermission) GetPermissions() []Role`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserPermission) GetPermissionsOk() (*[]Role, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserPermission) SetPermissions(v []Role)`

SetPermissions sets Permissions field to given value.


### GetRoles

`func (o *UserPermission) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserPermission) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserPermission) SetRoles(v []Role)`

SetRoles sets Roles field to given value.


### GetUiPermissions

`func (o *UserPermission) GetUiPermissions() []string`

GetUiPermissions returns the UiPermissions field if non-nil, zero value otherwise.

### GetUiPermissionsOk

`func (o *UserPermission) GetUiPermissionsOk() (*[]string, bool)`

GetUiPermissionsOk returns a tuple with the UiPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiPermissions

`func (o *UserPermission) SetUiPermissions(v []string)`

SetUiPermissions sets UiPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


