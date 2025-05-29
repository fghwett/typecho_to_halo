# SettingForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormSchema** | **[]map[string]interface{}** |  | 
**Group** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewSettingForm

`func NewSettingForm(formSchema []map[string]interface{}, group string, ) *SettingForm`

NewSettingForm instantiates a new SettingForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingFormWithDefaults

`func NewSettingFormWithDefaults() *SettingForm`

NewSettingFormWithDefaults instantiates a new SettingForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormSchema

`func (o *SettingForm) GetFormSchema() []map[string]interface{}`

GetFormSchema returns the FormSchema field if non-nil, zero value otherwise.

### GetFormSchemaOk

`func (o *SettingForm) GetFormSchemaOk() (*[]map[string]interface{}, bool)`

GetFormSchemaOk returns a tuple with the FormSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormSchema

`func (o *SettingForm) SetFormSchema(v []map[string]interface{})`

SetFormSchema sets FormSchema field to given value.


### GetGroup

`func (o *SettingForm) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SettingForm) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SettingForm) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetLabel

`func (o *SettingForm) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SettingForm) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SettingForm) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SettingForm) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


