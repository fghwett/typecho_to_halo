# SnapshotList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | **bool** | Indicates whether current page is the first page. | 
**HasNext** | **bool** | Indicates whether current page has previous page. | 
**HasPrevious** | **bool** | Indicates whether current page has previous page. | 
**Items** | [**[]Snapshot**](Snapshot.md) | A chunk of items. | 
**Last** | **bool** | Indicates whether current page is the last page. | 
**Page** | **int32** | Page number, starts from 1. If not set or equal to 0, it means no pagination. | 
**Size** | **int32** | Size of each page. If not set or equal to 0, it means no pagination. | 
**Total** | **int64** | Total elements. | 
**TotalPages** | **int64** | Indicates total pages. | 

## Methods

### NewSnapshotList

`func NewSnapshotList(first bool, hasNext bool, hasPrevious bool, items []Snapshot, last bool, page int32, size int32, total int64, totalPages int64, ) *SnapshotList`

NewSnapshotList instantiates a new SnapshotList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotListWithDefaults

`func NewSnapshotListWithDefaults() *SnapshotList`

NewSnapshotListWithDefaults instantiates a new SnapshotList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *SnapshotList) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *SnapshotList) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *SnapshotList) SetFirst(v bool)`

SetFirst sets First field to given value.


### GetHasNext

`func (o *SnapshotList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *SnapshotList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *SnapshotList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *SnapshotList) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *SnapshotList) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *SnapshotList) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetItems

`func (o *SnapshotList) GetItems() []Snapshot`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SnapshotList) GetItemsOk() (*[]Snapshot, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SnapshotList) SetItems(v []Snapshot)`

SetItems sets Items field to given value.


### GetLast

`func (o *SnapshotList) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *SnapshotList) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *SnapshotList) SetLast(v bool)`

SetLast sets Last field to given value.


### GetPage

`func (o *SnapshotList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SnapshotList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SnapshotList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *SnapshotList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *SnapshotList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SnapshotList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SnapshotList) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *SnapshotList) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *SnapshotList) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *SnapshotList) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


