# MenuItemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Calculated Display name of menu item. | [optional] 
**Href** | Pointer to **string** | Calculated href of manu item. | [optional] 

## Methods

### NewMenuItemStatus

`func NewMenuItemStatus() *MenuItemStatus`

NewMenuItemStatus instantiates a new MenuItemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuItemStatusWithDefaults

`func NewMenuItemStatusWithDefaults() *MenuItemStatus`

NewMenuItemStatusWithDefaults instantiates a new MenuItemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MenuItemStatus) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MenuItemStatus) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MenuItemStatus) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MenuItemStatus) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHref

`func (o *MenuItemStatus) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MenuItemStatus) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MenuItemStatus) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MenuItemStatus) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


