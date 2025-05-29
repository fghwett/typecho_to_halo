# InterestReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | The expression to be interested in | [optional] 
**ReasonType** | **string** | The name of the reason definition to be interested in | 
**Subject** | [**InterestReasonSubject**](InterestReasonSubject.md) |  | 

## Methods

### NewInterestReason

`func NewInterestReason(reasonType string, subject InterestReasonSubject, ) *InterestReason`

NewInterestReason instantiates a new InterestReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterestReasonWithDefaults

`func NewInterestReasonWithDefaults() *InterestReason`

NewInterestReasonWithDefaults instantiates a new InterestReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *InterestReason) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *InterestReason) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *InterestReason) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *InterestReason) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetReasonType

`func (o *InterestReason) GetReasonType() string`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *InterestReason) GetReasonTypeOk() (*string, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *InterestReason) SetReasonType(v string)`

SetReasonType sets ReasonType field to given value.


### GetSubject

`func (o *InterestReason) GetSubject() InterestReasonSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InterestReason) GetSubjectOk() (*InterestReasonSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InterestReason) SetSubject(v InterestReasonSubject)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


