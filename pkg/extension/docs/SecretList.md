# SecretList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | **bool** | Indicates whether current page is the first page. | 
**HasNext** | **bool** | Indicates whether current page has previous page. | 
**HasPrevious** | **bool** | Indicates whether current page has previous page. | 
**Items** | [**[]Secret**](Secret.md) | A chunk of items. | 
**Last** | **bool** | Indicates whether current page is the last page. | 
**Page** | **int32** | Page number, starts from 1. If not set or equal to 0, it means no pagination. | 
**Size** | **int32** | Size of each page. If not set or equal to 0, it means no pagination. | 
**Total** | **int64** | Total elements. | 
**TotalPages** | **int64** | Indicates total pages. | 

## Methods

### NewSecretList

`func NewSecretList(first bool, hasNext bool, hasPrevious bool, items []Secret, last bool, page int32, size int32, total int64, totalPages int64, ) *SecretList`

NewSecretList instantiates a new SecretList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretListWithDefaults

`func NewSecretListWithDefaults() *SecretList`

NewSecretListWithDefaults instantiates a new SecretList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *SecretList) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *SecretList) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *SecretList) SetFirst(v bool)`

SetFirst sets First field to given value.


### GetHasNext

`func (o *SecretList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *SecretList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *SecretList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *SecretList) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *SecretList) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *SecretList) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetItems

`func (o *SecretList) GetItems() []Secret`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecretList) GetItemsOk() (*[]Secret, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecretList) SetItems(v []Secret)`

SetItems sets Items field to given value.


### GetLast

`func (o *SecretList) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *SecretList) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *SecretList) SetLast(v bool)`

SetLast sets Last field to given value.


### GetPage

`func (o *SecretList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SecretList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SecretList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *SecretList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SecretList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SecretList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *SecretList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SecretList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SecretList) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *SecretList) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *SecretList) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *SecretList) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


