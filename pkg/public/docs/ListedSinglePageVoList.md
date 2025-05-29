# ListedSinglePageVoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | **bool** | Indicates whether current page is the first page. | 
**HasNext** | **bool** | Indicates whether current page has previous page. | 
**HasPrevious** | **bool** | Indicates whether current page has previous page. | 
**Items** | [**[]ListedSinglePageVo**](ListedSinglePageVo.md) | A chunk of items. | 
**Last** | **bool** | Indicates whether current page is the last page. | 
**Page** | **int32** | Page number, starts from 1. If not set or equal to 0, it means no pagination. | 
**Size** | **int32** | Size of each page. If not set or equal to 0, it means no pagination. | 
**Total** | **int64** | Total elements. | 
**TotalPages** | **int64** | Indicates total pages. | 

## Methods

### NewListedSinglePageVoList

`func NewListedSinglePageVoList(first bool, hasNext bool, hasPrevious bool, items []ListedSinglePageVo, last bool, page int32, size int32, total int64, totalPages int64, ) *ListedSinglePageVoList`

NewListedSinglePageVoList instantiates a new ListedSinglePageVoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedSinglePageVoListWithDefaults

`func NewListedSinglePageVoListWithDefaults() *ListedSinglePageVoList`

NewListedSinglePageVoListWithDefaults instantiates a new ListedSinglePageVoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *ListedSinglePageVoList) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ListedSinglePageVoList) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ListedSinglePageVoList) SetFirst(v bool)`

SetFirst sets First field to given value.


### GetHasNext

`func (o *ListedSinglePageVoList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListedSinglePageVoList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListedSinglePageVoList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *ListedSinglePageVoList) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *ListedSinglePageVoList) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *ListedSinglePageVoList) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetItems

`func (o *ListedSinglePageVoList) GetItems() []ListedSinglePageVo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListedSinglePageVoList) GetItemsOk() (*[]ListedSinglePageVo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListedSinglePageVoList) SetItems(v []ListedSinglePageVo)`

SetItems sets Items field to given value.


### GetLast

`func (o *ListedSinglePageVoList) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *ListedSinglePageVoList) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *ListedSinglePageVoList) SetLast(v bool)`

SetLast sets Last field to given value.


### GetPage

`func (o *ListedSinglePageVoList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListedSinglePageVoList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListedSinglePageVoList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *ListedSinglePageVoList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListedSinglePageVoList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListedSinglePageVoList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *ListedSinglePageVoList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListedSinglePageVoList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListedSinglePageVoList) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *ListedSinglePageVoList) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ListedSinglePageVoList) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ListedSinglePageVoList) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


