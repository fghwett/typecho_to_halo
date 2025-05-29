# PersonalAccessTokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | **bool** | Indicates whether current page is the first page. | 
**HasNext** | **bool** | Indicates whether current page has previous page. | 
**HasPrevious** | **bool** | Indicates whether current page has previous page. | 
**Items** | [**[]PersonalAccessToken**](PersonalAccessToken.md) | A chunk of items. | 
**Last** | **bool** | Indicates whether current page is the last page. | 
**Page** | **int32** | Page number, starts from 1. If not set or equal to 0, it means no pagination. | 
**Size** | **int32** | Size of each page. If not set or equal to 0, it means no pagination. | 
**Total** | **int64** | Total elements. | 
**TotalPages** | **int64** | Indicates total pages. | 

## Methods

### NewPersonalAccessTokenList

`func NewPersonalAccessTokenList(first bool, hasNext bool, hasPrevious bool, items []PersonalAccessToken, last bool, page int32, size int32, total int64, totalPages int64, ) *PersonalAccessTokenList`

NewPersonalAccessTokenList instantiates a new PersonalAccessTokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalAccessTokenListWithDefaults

`func NewPersonalAccessTokenListWithDefaults() *PersonalAccessTokenList`

NewPersonalAccessTokenListWithDefaults instantiates a new PersonalAccessTokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *PersonalAccessTokenList) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PersonalAccessTokenList) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PersonalAccessTokenList) SetFirst(v bool)`

SetFirst sets First field to given value.


### GetHasNext

`func (o *PersonalAccessTokenList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PersonalAccessTokenList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PersonalAccessTokenList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *PersonalAccessTokenList) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *PersonalAccessTokenList) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *PersonalAccessTokenList) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetItems

`func (o *PersonalAccessTokenList) GetItems() []PersonalAccessToken`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PersonalAccessTokenList) GetItemsOk() (*[]PersonalAccessToken, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PersonalAccessTokenList) SetItems(v []PersonalAccessToken)`

SetItems sets Items field to given value.


### GetLast

`func (o *PersonalAccessTokenList) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PersonalAccessTokenList) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PersonalAccessTokenList) SetLast(v bool)`

SetLast sets Last field to given value.


### GetPage

`func (o *PersonalAccessTokenList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PersonalAccessTokenList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PersonalAccessTokenList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *PersonalAccessTokenList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PersonalAccessTokenList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PersonalAccessTokenList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *PersonalAccessTokenList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PersonalAccessTokenList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PersonalAccessTokenList) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *PersonalAccessTokenList) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PersonalAccessTokenList) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PersonalAccessTokenList) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


