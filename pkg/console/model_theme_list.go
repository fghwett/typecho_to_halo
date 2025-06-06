/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package console_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ThemeList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemeList{}

// ThemeList struct for ThemeList
type ThemeList struct {
	// Indicates whether current page is the first page.
	First bool `json:"first"`
	// Indicates whether current page has previous page.
	HasNext bool `json:"hasNext"`
	// Indicates whether current page has previous page.
	HasPrevious bool `json:"hasPrevious"`
	// A chunk of items.
	Items []Theme `json:"items"`
	// Indicates whether current page is the last page.
	Last bool `json:"last"`
	// Page number, starts from 1. If not set or equal to 0, it means no pagination.
	Page int32 `json:"page"`
	// Size of each page. If not set or equal to 0, it means no pagination.
	Size int32 `json:"size"`
	// Total elements.
	Total int64 `json:"total"`
	// Indicates total pages.
	TotalPages int64 `json:"totalPages"`
}

type _ThemeList ThemeList

// NewThemeList instantiates a new ThemeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeList(first bool, hasNext bool, hasPrevious bool, items []Theme, last bool, page int32, size int32, total int64, totalPages int64) *ThemeList {
	this := ThemeList{}
	this.First = first
	this.HasNext = hasNext
	this.HasPrevious = hasPrevious
	this.Items = items
	this.Last = last
	this.Page = page
	this.Size = size
	this.Total = total
	this.TotalPages = totalPages
	return &this
}

// NewThemeListWithDefaults instantiates a new ThemeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeListWithDefaults() *ThemeList {
	this := ThemeList{}
	return &this
}

// GetFirst returns the First field value
func (o *ThemeList) GetFirst() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetFirstOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value
func (o *ThemeList) SetFirst(v bool) {
	o.First = v
}

// GetHasNext returns the HasNext field value
func (o *ThemeList) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetHasNextOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *ThemeList) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrevious returns the HasPrevious field value
func (o *ThemeList) GetHasPrevious() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrevious
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetHasPreviousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPrevious, true
}

// SetHasPrevious sets field value
func (o *ThemeList) SetHasPrevious(v bool) {
	o.HasPrevious = v
}

// GetItems returns the Items field value
func (o *ThemeList) GetItems() []Theme {
	if o == nil {
		var ret []Theme
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetItemsOk() ([]Theme, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ThemeList) SetItems(v []Theme) {
	o.Items = v
}

// GetLast returns the Last field value
func (o *ThemeList) GetLast() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Last
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetLastOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last, true
}

// SetLast sets field value
func (o *ThemeList) SetLast(v bool) {
	o.Last = v
}

// GetPage returns the Page field value
func (o *ThemeList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ThemeList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *ThemeList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ThemeList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *ThemeList) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ThemeList) SetTotal(v int64) {
	o.Total = v
}

// GetTotalPages returns the TotalPages field value
func (o *ThemeList) GetTotalPages() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *ThemeList) GetTotalPagesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *ThemeList) SetTotalPages(v int64) {
	o.TotalPages = v
}

func (o ThemeList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThemeList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["first"] = o.First
	toSerialize["hasNext"] = o.HasNext
	toSerialize["hasPrevious"] = o.HasPrevious
	toSerialize["items"] = o.Items
	toSerialize["last"] = o.Last
	toSerialize["page"] = o.Page
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["totalPages"] = o.TotalPages
	return toSerialize, nil
}

func (o *ThemeList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first",
		"hasNext",
		"hasPrevious",
		"items",
		"last",
		"page",
		"size",
		"total",
		"totalPages",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varThemeList := _ThemeList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThemeList)

	if err != nil {
		return err
	}

	*o = ThemeList(varThemeList)

	return err
}

type NullableThemeList struct {
	value *ThemeList
	isSet bool
}

func (v NullableThemeList) Get() *ThemeList {
	return v.value
}

func (v *NullableThemeList) Set(val *ThemeList) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeList) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeList(val *ThemeList) *NullableThemeList {
	return &NullableThemeList{value: val, isSet: true}
}

func (v NullableThemeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThemeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


