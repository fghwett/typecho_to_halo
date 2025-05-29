# MenuSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the menu. | 
**MenuItems** | Pointer to **[]string** | Menu items of this menu. | [optional] 

## Methods

### NewMenuSpec

`func NewMenuSpec(displayName string, ) *MenuSpec`

NewMenuSpec instantiates a new MenuSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuSpecWithDefaults

`func NewMenuSpecWithDefaults() *MenuSpec`

NewMenuSpecWithDefaults instantiates a new MenuSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MenuSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MenuSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MenuSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMenuItems

`func (o *MenuSpec) GetMenuItems() []string`

GetMenuItems returns the MenuItems field if non-nil, zero value otherwise.

### GetMenuItemsOk

`func (o *MenuSpec) GetMenuItemsOk() (*[]string, bool)`

GetMenuItemsOk returns a tuple with the MenuItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuItems

`func (o *MenuSpec) SetMenuItems(v []string)`

SetMenuItems sets MenuItems field to given value.

### HasMenuItems

`func (o *MenuSpec) HasMenuItems() bool`

HasMenuItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


