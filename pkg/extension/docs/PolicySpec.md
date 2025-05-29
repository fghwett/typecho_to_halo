# PolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapName** | Pointer to **string** | Reference name of ConfigMap extension | [optional] 
**DisplayName** | **string** | Display name of policy | 
**TemplateName** | **string** | Reference name of PolicyTemplate | 

## Methods

### NewPolicySpec

`func NewPolicySpec(displayName string, templateName string, ) *PolicySpec`

NewPolicySpec instantiates a new PolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySpecWithDefaults

`func NewPolicySpecWithDefaults() *PolicySpec`

NewPolicySpecWithDefaults instantiates a new PolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMapName

`func (o *PolicySpec) GetConfigMapName() string`

GetConfigMapName returns the ConfigMapName field if non-nil, zero value otherwise.

### GetConfigMapNameOk

`func (o *PolicySpec) GetConfigMapNameOk() (*string, bool)`

GetConfigMapNameOk returns a tuple with the ConfigMapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapName

`func (o *PolicySpec) SetConfigMapName(v string)`

SetConfigMapName sets ConfigMapName field to given value.

### HasConfigMapName

`func (o *PolicySpec) HasConfigMapName() bool`

HasConfigMapName returns a boolean if a field has been set.

### GetDisplayName

`func (o *PolicySpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PolicySpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PolicySpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetTemplateName

`func (o *PolicySpec) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *PolicySpec) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *PolicySpec) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


