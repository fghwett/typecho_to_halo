# ExtensionPointDefinitionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | **bool** | Indicates whether current page is the first page. | 
**HasNext** | **bool** | Indicates whether current page has previous page. | 
**HasPrevious** | **bool** | Indicates whether current page has previous page. | 
**Items** | [**[]ExtensionPointDefinition**](ExtensionPointDefinition.md) | A chunk of items. | 
**Last** | **bool** | Indicates whether current page is the last page. | 
**Page** | **int32** | Page number, starts from 1. If not set or equal to 0, it means no pagination. | 
**Size** | **int32** | Size of each page. If not set or equal to 0, it means no pagination. | 
**Total** | **int64** | Total elements. | 
**TotalPages** | **int64** | Indicates total pages. | 

## Methods

### NewExtensionPointDefinitionList

`func NewExtensionPointDefinitionList(first bool, hasNext bool, hasPrevious bool, items []ExtensionPointDefinition, last bool, page int32, size int32, total int64, totalPages int64, ) *ExtensionPointDefinitionList`

NewExtensionPointDefinitionList instantiates a new ExtensionPointDefinitionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionPointDefinitionListWithDefaults

`func NewExtensionPointDefinitionListWithDefaults() *ExtensionPointDefinitionList`

NewExtensionPointDefinitionListWithDefaults instantiates a new ExtensionPointDefinitionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *ExtensionPointDefinitionList) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ExtensionPointDefinitionList) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ExtensionPointDefinitionList) SetFirst(v bool)`

SetFirst sets First field to given value.


### GetHasNext

`func (o *ExtensionPointDefinitionList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ExtensionPointDefinitionList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ExtensionPointDefinitionList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *ExtensionPointDefinitionList) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *ExtensionPointDefinitionList) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *ExtensionPointDefinitionList) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetItems

`func (o *ExtensionPointDefinitionList) GetItems() []ExtensionPointDefinition`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExtensionPointDefinitionList) GetItemsOk() (*[]ExtensionPointDefinition, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExtensionPointDefinitionList) SetItems(v []ExtensionPointDefinition)`

SetItems sets Items field to given value.


### GetLast

`func (o *ExtensionPointDefinitionList) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *ExtensionPointDefinitionList) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *ExtensionPointDefinitionList) SetLast(v bool)`

SetLast sets Last field to given value.


### GetPage

`func (o *ExtensionPointDefinitionList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ExtensionPointDefinitionList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ExtensionPointDefinitionList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *ExtensionPointDefinitionList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ExtensionPointDefinitionList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ExtensionPointDefinitionList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *ExtensionPointDefinitionList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExtensionPointDefinitionList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExtensionPointDefinitionList) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *ExtensionPointDefinitionList) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ExtensionPointDefinitionList) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ExtensionPointDefinitionList) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


