# AnnotationSettingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormSchema** | **[]map[string]interface{}** |  | 
**TargetRef** | [**GroupKind**](GroupKind.md) |  | 

## Methods

### NewAnnotationSettingSpec

`func NewAnnotationSettingSpec(formSchema []map[string]interface{}, targetRef GroupKind, ) *AnnotationSettingSpec`

NewAnnotationSettingSpec instantiates a new AnnotationSettingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationSettingSpecWithDefaults

`func NewAnnotationSettingSpecWithDefaults() *AnnotationSettingSpec`

NewAnnotationSettingSpecWithDefaults instantiates a new AnnotationSettingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormSchema

`func (o *AnnotationSettingSpec) GetFormSchema() []map[string]interface{}`

GetFormSchema returns the FormSchema field if non-nil, zero value otherwise.

### GetFormSchemaOk

`func (o *AnnotationSettingSpec) GetFormSchemaOk() (*[]map[string]interface{}, bool)`

GetFormSchemaOk returns a tuple with the FormSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormSchema

`func (o *AnnotationSettingSpec) SetFormSchema(v []map[string]interface{})`

SetFormSchema sets FormSchema field to given value.


### GetTargetRef

`func (o *AnnotationSettingSpec) GetTargetRef() GroupKind`

GetTargetRef returns the TargetRef field if non-nil, zero value otherwise.

### GetTargetRefOk

`func (o *AnnotationSettingSpec) GetTargetRefOk() (*GroupKind, bool)`

GetTargetRefOk returns a tuple with the TargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRef

`func (o *AnnotationSettingSpec) SetTargetRef(v GroupKind)`

SetTargetRef sets TargetRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


