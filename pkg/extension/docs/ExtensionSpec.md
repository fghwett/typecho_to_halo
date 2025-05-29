# ExtensionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**ExtensionPointName** | **string** |  | 
**Icon** | Pointer to **string** |  | [optional] 

## Methods

### NewExtensionSpec

`func NewExtensionSpec(className string, displayName string, extensionPointName string, ) *ExtensionSpec`

NewExtensionSpec instantiates a new ExtensionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionSpecWithDefaults

`func NewExtensionSpecWithDefaults() *ExtensionSpec`

NewExtensionSpecWithDefaults instantiates a new ExtensionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *ExtensionSpec) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ExtensionSpec) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ExtensionSpec) SetClassName(v string)`

SetClassName sets ClassName field to given value.


### GetDescription

`func (o *ExtensionSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtensionSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtensionSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtensionSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ExtensionSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ExtensionSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ExtensionSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExtensionPointName

`func (o *ExtensionSpec) GetExtensionPointName() string`

GetExtensionPointName returns the ExtensionPointName field if non-nil, zero value otherwise.

### GetExtensionPointNameOk

`func (o *ExtensionSpec) GetExtensionPointNameOk() (*string, bool)`

GetExtensionPointNameOk returns a tuple with the ExtensionPointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionPointName

`func (o *ExtensionSpec) SetExtensionPointName(v string)`

SetExtensionPointName sets ExtensionPointName field to given value.


### GetIcon

`func (o *ExtensionSpec) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ExtensionSpec) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ExtensionSpec) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ExtensionSpec) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


