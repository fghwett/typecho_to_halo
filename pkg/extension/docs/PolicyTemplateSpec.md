# PolicyTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**SettingName** | **string** |  | 

## Methods

### NewPolicyTemplateSpec

`func NewPolicyTemplateSpec(settingName string, ) *PolicyTemplateSpec`

NewPolicyTemplateSpec instantiates a new PolicyTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateSpecWithDefaults

`func NewPolicyTemplateSpecWithDefaults() *PolicyTemplateSpec`

NewPolicyTemplateSpecWithDefaults instantiates a new PolicyTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PolicyTemplateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PolicyTemplateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PolicyTemplateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PolicyTemplateSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSettingName

`func (o *PolicyTemplateSpec) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *PolicyTemplateSpec) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *PolicyTemplateSpec) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


