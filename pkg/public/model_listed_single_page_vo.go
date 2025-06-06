/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListedSinglePageVo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListedSinglePageVo{}

// ListedSinglePageVo A chunk of items.
type ListedSinglePageVo struct {
	Contributors []ContributorVo `json:"contributors,omitempty"`
	Metadata Metadata `json:"metadata"`
	Owner *ContributorVo `json:"owner,omitempty"`
	Spec *SinglePageSpec `json:"spec,omitempty"`
	Stats *StatsVo `json:"stats,omitempty"`
	Status *SinglePageStatus `json:"status,omitempty"`
}

type _ListedSinglePageVo ListedSinglePageVo

// NewListedSinglePageVo instantiates a new ListedSinglePageVo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListedSinglePageVo(metadata Metadata) *ListedSinglePageVo {
	this := ListedSinglePageVo{}
	this.Metadata = metadata
	return &this
}

// NewListedSinglePageVoWithDefaults instantiates a new ListedSinglePageVo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListedSinglePageVoWithDefaults() *ListedSinglePageVo {
	this := ListedSinglePageVo{}
	return &this
}

// GetContributors returns the Contributors field value if set, zero value otherwise.
func (o *ListedSinglePageVo) GetContributors() []ContributorVo {
	if o == nil || IsNil(o.Contributors) {
		var ret []ContributorVo
		return ret
	}
	return o.Contributors
}

// GetContributorsOk returns a tuple with the Contributors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedSinglePageVo) GetContributorsOk() ([]ContributorVo, bool) {
	if o == nil || IsNil(o.Contributors) {
		return nil, false
	}
	return o.Contributors, true
}

// HasContributors returns a boolean if a field has been set.
func (o *ListedSinglePageVo) HasContributors() bool {
	if o != nil && !IsNil(o.Contributors) {
		return true
	}

	return false
}

// SetContributors gets a reference to the given []ContributorVo and assigns it to the Contributors field.
func (o *ListedSinglePageVo) SetContributors(v []ContributorVo) {
	o.Contributors = v
}

// GetMetadata returns the Metadata field value
func (o *ListedSinglePageVo) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ListedSinglePageVo) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ListedSinglePageVo) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ListedSinglePageVo) GetOwner() ContributorVo {
	if o == nil || IsNil(o.Owner) {
		var ret ContributorVo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedSinglePageVo) GetOwnerOk() (*ContributorVo, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ListedSinglePageVo) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ContributorVo and assigns it to the Owner field.
func (o *ListedSinglePageVo) SetOwner(v ContributorVo) {
	o.Owner = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ListedSinglePageVo) GetSpec() SinglePageSpec {
	if o == nil || IsNil(o.Spec) {
		var ret SinglePageSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedSinglePageVo) GetSpecOk() (*SinglePageSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ListedSinglePageVo) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given SinglePageSpec and assigns it to the Spec field.
func (o *ListedSinglePageVo) SetSpec(v SinglePageSpec) {
	o.Spec = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ListedSinglePageVo) GetStats() StatsVo {
	if o == nil || IsNil(o.Stats) {
		var ret StatsVo
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedSinglePageVo) GetStatsOk() (*StatsVo, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *ListedSinglePageVo) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given StatsVo and assigns it to the Stats field.
func (o *ListedSinglePageVo) SetStats(v StatsVo) {
	o.Stats = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListedSinglePageVo) GetStatus() SinglePageStatus {
	if o == nil || IsNil(o.Status) {
		var ret SinglePageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListedSinglePageVo) GetStatusOk() (*SinglePageStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ListedSinglePageVo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SinglePageStatus and assigns it to the Status field.
func (o *ListedSinglePageVo) SetStatus(v SinglePageStatus) {
	o.Status = &v
}

func (o ListedSinglePageVo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListedSinglePageVo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contributors) {
		toSerialize["contributors"] = o.Contributors
	}
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *ListedSinglePageVo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata",
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

	varListedSinglePageVo := _ListedSinglePageVo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListedSinglePageVo)

	if err != nil {
		return err
	}

	*o = ListedSinglePageVo(varListedSinglePageVo)

	return err
}

type NullableListedSinglePageVo struct {
	value *ListedSinglePageVo
	isSet bool
}

func (v NullableListedSinglePageVo) Get() *ListedSinglePageVo {
	return v.value
}

func (v *NullableListedSinglePageVo) Set(val *ListedSinglePageVo) {
	v.value = val
	v.isSet = true
}

func (v NullableListedSinglePageVo) IsSet() bool {
	return v.isSet
}

func (v *NullableListedSinglePageVo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListedSinglePageVo(val *ListedSinglePageVo) *NullableListedSinglePageVo {
	return &NullableListedSinglePageVo{value: val, isSet: true}
}

func (v NullableListedSinglePageVo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListedSinglePageVo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


