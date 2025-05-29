# MenuItemVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Metadata** | [**Metadata**](Metadata.md) |  | 
**ParentName** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**MenuItemSpec**](MenuItemSpec.md) |  | [optional] 
**Status** | Pointer to [**MenuItemStatus**](MenuItemStatus.md) |  | [optional] 

## Methods

### NewMenuItemVo

`func NewMenuItemVo(metadata Metadata, ) *MenuItemVo`

NewMenuItemVo instantiates a new MenuItemVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuItemVoWithDefaults

`func NewMenuItemVoWithDefaults() *MenuItemVo`

NewMenuItemVoWithDefaults instantiates a new MenuItemVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MenuItemVo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MenuItemVo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MenuItemVo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MenuItemVo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMetadata

`func (o *MenuItemVo) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MenuItemVo) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MenuItemVo) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetParentName

`func (o *MenuItemVo) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *MenuItemVo) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *MenuItemVo) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *MenuItemVo) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetSpec

`func (o *MenuItemVo) GetSpec() MenuItemSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MenuItemVo) GetSpecOk() (*MenuItemSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MenuItemVo) SetSpec(v MenuItemSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MenuItemVo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MenuItemVo) GetStatus() MenuItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MenuItemVo) GetStatusOk() (*MenuItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MenuItemVo) SetStatus(v MenuItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MenuItemVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


