# MenuItemSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to **[]string** | Children of this menu item | [optional] 
**DisplayName** | Pointer to **string** | The display name of menu item. | [optional] 
**Href** | Pointer to **string** | The href of this menu item. | [optional] 
**Priority** | Pointer to **int32** | The priority is for ordering. | [optional] 
**Target** | Pointer to **string** | The &lt;a&gt; target attribute of this menu item. | [optional] 
**TargetRef** | Pointer to [**Ref**](Ref.md) |  | [optional] 

## Methods

### NewMenuItemSpec

`func NewMenuItemSpec() *MenuItemSpec`

NewMenuItemSpec instantiates a new MenuItemSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuItemSpecWithDefaults

`func NewMenuItemSpecWithDefaults() *MenuItemSpec`

NewMenuItemSpecWithDefaults instantiates a new MenuItemSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *MenuItemSpec) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MenuItemSpec) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MenuItemSpec) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MenuItemSpec) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDisplayName

`func (o *MenuItemSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MenuItemSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MenuItemSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MenuItemSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHref

`func (o *MenuItemSpec) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MenuItemSpec) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MenuItemSpec) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MenuItemSpec) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetPriority

`func (o *MenuItemSpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MenuItemSpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MenuItemSpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MenuItemSpec) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTarget

`func (o *MenuItemSpec) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MenuItemSpec) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MenuItemSpec) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MenuItemSpec) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetRef

`func (o *MenuItemSpec) GetTargetRef() Ref`

GetTargetRef returns the TargetRef field if non-nil, zero value otherwise.

### GetTargetRefOk

`func (o *MenuItemSpec) GetTargetRefOk() (*Ref, bool)`

GetTargetRefOk returns a tuple with the TargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRef

`func (o *MenuItemSpec) SetTargetRef(v Ref)`

SetTargetRef sets TargetRef field to given value.

### HasTargetRef

`func (o *MenuItemSpec) HasTargetRef() bool`

HasTargetRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


