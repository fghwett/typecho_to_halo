# ReasonSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**ReasonAttributes**](ReasonAttributes.md) |  | [optional] 
**Author** | **string** |  | 
**ReasonType** | **string** |  | 
**Subject** | [**ReasonSubject**](ReasonSubject.md) |  | 

## Methods

### NewReasonSpec

`func NewReasonSpec(author string, reasonType string, subject ReasonSubject, ) *ReasonSpec`

NewReasonSpec instantiates a new ReasonSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReasonSpecWithDefaults

`func NewReasonSpecWithDefaults() *ReasonSpec`

NewReasonSpecWithDefaults instantiates a new ReasonSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ReasonSpec) GetAttributes() ReasonAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReasonSpec) GetAttributesOk() (*ReasonAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReasonSpec) SetAttributes(v ReasonAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ReasonSpec) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAuthor

`func (o *ReasonSpec) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReasonSpec) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReasonSpec) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetReasonType

`func (o *ReasonSpec) GetReasonType() string`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *ReasonSpec) GetReasonTypeOk() (*string, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *ReasonSpec) SetReasonType(v string)`

SetReasonType sets ReasonType field to given value.


### GetSubject

`func (o *ReasonSpec) GetSubject() ReasonSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReasonSpec) GetSubjectOk() (*ReasonSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReasonSpec) SetSubject(v ReasonSubject)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


