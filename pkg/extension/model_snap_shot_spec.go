/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package extension_sdk

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the SnapShotSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapShotSpec{}

// SnapShotSpec struct for SnapShotSpec
type SnapShotSpec struct {
	ContentPatch *string `json:"contentPatch,omitempty"`
	Contributors []string `json:"contributors,omitempty"`
	LastModifyTime *time.Time `json:"lastModifyTime,omitempty"`
	Owner string `json:"owner"`
	ParentSnapshotName *string `json:"parentSnapshotName,omitempty"`
	RawPatch *string `json:"rawPatch,omitempty"`
	RawType string `json:"rawType"`
	SubjectRef Ref `json:"subjectRef"`
}

type _SnapShotSpec SnapShotSpec

// NewSnapShotSpec instantiates a new SnapShotSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapShotSpec(owner string, rawType string, subjectRef Ref) *SnapShotSpec {
	this := SnapShotSpec{}
	this.Owner = owner
	this.RawType = rawType
	this.SubjectRef = subjectRef
	return &this
}

// NewSnapShotSpecWithDefaults instantiates a new SnapShotSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapShotSpecWithDefaults() *SnapShotSpec {
	this := SnapShotSpec{}
	return &this
}

// GetContentPatch returns the ContentPatch field value if set, zero value otherwise.
func (o *SnapShotSpec) GetContentPatch() string {
	if o == nil || IsNil(o.ContentPatch) {
		var ret string
		return ret
	}
	return *o.ContentPatch
}

// GetContentPatchOk returns a tuple with the ContentPatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetContentPatchOk() (*string, bool) {
	if o == nil || IsNil(o.ContentPatch) {
		return nil, false
	}
	return o.ContentPatch, true
}

// HasContentPatch returns a boolean if a field has been set.
func (o *SnapShotSpec) HasContentPatch() bool {
	if o != nil && !IsNil(o.ContentPatch) {
		return true
	}

	return false
}

// SetContentPatch gets a reference to the given string and assigns it to the ContentPatch field.
func (o *SnapShotSpec) SetContentPatch(v string) {
	o.ContentPatch = &v
}

// GetContributors returns the Contributors field value if set, zero value otherwise.
func (o *SnapShotSpec) GetContributors() []string {
	if o == nil || IsNil(o.Contributors) {
		var ret []string
		return ret
	}
	return o.Contributors
}

// GetContributorsOk returns a tuple with the Contributors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetContributorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Contributors) {
		return nil, false
	}
	return o.Contributors, true
}

// HasContributors returns a boolean if a field has been set.
func (o *SnapShotSpec) HasContributors() bool {
	if o != nil && !IsNil(o.Contributors) {
		return true
	}

	return false
}

// SetContributors gets a reference to the given []string and assigns it to the Contributors field.
func (o *SnapShotSpec) SetContributors(v []string) {
	o.Contributors = v
}

// GetLastModifyTime returns the LastModifyTime field value if set, zero value otherwise.
func (o *SnapShotSpec) GetLastModifyTime() time.Time {
	if o == nil || IsNil(o.LastModifyTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifyTime
}

// GetLastModifyTimeOk returns a tuple with the LastModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetLastModifyTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifyTime) {
		return nil, false
	}
	return o.LastModifyTime, true
}

// HasLastModifyTime returns a boolean if a field has been set.
func (o *SnapShotSpec) HasLastModifyTime() bool {
	if o != nil && !IsNil(o.LastModifyTime) {
		return true
	}

	return false
}

// SetLastModifyTime gets a reference to the given time.Time and assigns it to the LastModifyTime field.
func (o *SnapShotSpec) SetLastModifyTime(v time.Time) {
	o.LastModifyTime = &v
}

// GetOwner returns the Owner field value
func (o *SnapShotSpec) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *SnapShotSpec) SetOwner(v string) {
	o.Owner = v
}

// GetParentSnapshotName returns the ParentSnapshotName field value if set, zero value otherwise.
func (o *SnapShotSpec) GetParentSnapshotName() string {
	if o == nil || IsNil(o.ParentSnapshotName) {
		var ret string
		return ret
	}
	return *o.ParentSnapshotName
}

// GetParentSnapshotNameOk returns a tuple with the ParentSnapshotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetParentSnapshotNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentSnapshotName) {
		return nil, false
	}
	return o.ParentSnapshotName, true
}

// HasParentSnapshotName returns a boolean if a field has been set.
func (o *SnapShotSpec) HasParentSnapshotName() bool {
	if o != nil && !IsNil(o.ParentSnapshotName) {
		return true
	}

	return false
}

// SetParentSnapshotName gets a reference to the given string and assigns it to the ParentSnapshotName field.
func (o *SnapShotSpec) SetParentSnapshotName(v string) {
	o.ParentSnapshotName = &v
}

// GetRawPatch returns the RawPatch field value if set, zero value otherwise.
func (o *SnapShotSpec) GetRawPatch() string {
	if o == nil || IsNil(o.RawPatch) {
		var ret string
		return ret
	}
	return *o.RawPatch
}

// GetRawPatchOk returns a tuple with the RawPatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetRawPatchOk() (*string, bool) {
	if o == nil || IsNil(o.RawPatch) {
		return nil, false
	}
	return o.RawPatch, true
}

// HasRawPatch returns a boolean if a field has been set.
func (o *SnapShotSpec) HasRawPatch() bool {
	if o != nil && !IsNil(o.RawPatch) {
		return true
	}

	return false
}

// SetRawPatch gets a reference to the given string and assigns it to the RawPatch field.
func (o *SnapShotSpec) SetRawPatch(v string) {
	o.RawPatch = &v
}

// GetRawType returns the RawType field value
func (o *SnapShotSpec) GetRawType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawType
}

// GetRawTypeOk returns a tuple with the RawType field value
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetRawTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawType, true
}

// SetRawType sets field value
func (o *SnapShotSpec) SetRawType(v string) {
	o.RawType = v
}

// GetSubjectRef returns the SubjectRef field value
func (o *SnapShotSpec) GetSubjectRef() Ref {
	if o == nil {
		var ret Ref
		return ret
	}

	return o.SubjectRef
}

// GetSubjectRefOk returns a tuple with the SubjectRef field value
// and a boolean to check if the value has been set.
func (o *SnapShotSpec) GetSubjectRefOk() (*Ref, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectRef, true
}

// SetSubjectRef sets field value
func (o *SnapShotSpec) SetSubjectRef(v Ref) {
	o.SubjectRef = v
}

func (o SnapShotSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapShotSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentPatch) {
		toSerialize["contentPatch"] = o.ContentPatch
	}
	if !IsNil(o.Contributors) {
		toSerialize["contributors"] = o.Contributors
	}
	if !IsNil(o.LastModifyTime) {
		toSerialize["lastModifyTime"] = o.LastModifyTime
	}
	toSerialize["owner"] = o.Owner
	if !IsNil(o.ParentSnapshotName) {
		toSerialize["parentSnapshotName"] = o.ParentSnapshotName
	}
	if !IsNil(o.RawPatch) {
		toSerialize["rawPatch"] = o.RawPatch
	}
	toSerialize["rawType"] = o.RawType
	toSerialize["subjectRef"] = o.SubjectRef
	return toSerialize, nil
}

func (o *SnapShotSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"owner",
		"rawType",
		"subjectRef",
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

	varSnapShotSpec := _SnapShotSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSnapShotSpec)

	if err != nil {
		return err
	}

	*o = SnapShotSpec(varSnapShotSpec)

	return err
}

type NullableSnapShotSpec struct {
	value *SnapShotSpec
	isSet bool
}

func (v NullableSnapShotSpec) Get() *SnapShotSpec {
	return v.value
}

func (v *NullableSnapShotSpec) Set(val *SnapShotSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapShotSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapShotSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapShotSpec(val *SnapShotSpec) *NullableSnapShotSpec {
	return &NullableSnapShotSpec{value: val, isSet: true}
}

func (v NullableSnapShotSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapShotSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


