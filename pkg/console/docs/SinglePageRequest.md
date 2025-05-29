# SinglePageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**ContentUpdateParam**](ContentUpdateParam.md) |  | 
**Page** | [**SinglePage**](SinglePage.md) |  | 

## Methods

### NewSinglePageRequest

`func NewSinglePageRequest(content ContentUpdateParam, page SinglePage, ) *SinglePageRequest`

NewSinglePageRequest instantiates a new SinglePageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglePageRequestWithDefaults

`func NewSinglePageRequestWithDefaults() *SinglePageRequest`

NewSinglePageRequestWithDefaults instantiates a new SinglePageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SinglePageRequest) GetContent() ContentUpdateParam`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SinglePageRequest) GetContentOk() (*ContentUpdateParam, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SinglePageRequest) SetContent(v ContentUpdateParam)`

SetContent sets Content field to given value.


### GetPage

`func (o *SinglePageRequest) GetPage() SinglePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SinglePageRequest) GetPageOk() (*SinglePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SinglePageRequest) SetPage(v SinglePage)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


