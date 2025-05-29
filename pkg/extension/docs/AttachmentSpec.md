# AttachmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of attachment | [optional] 
**GroupName** | Pointer to **string** | Group name | [optional] 
**MediaType** | Pointer to **string** | Media type of attachment | [optional] 
**OwnerName** | Pointer to **string** | Name of User who uploads the attachment | [optional] 
**PolicyName** | Pointer to **string** | Policy name | [optional] 
**Size** | Pointer to **int64** | Size of attachment. Unit is Byte | [optional] 
**Tags** | Pointer to **[]string** | Tags of attachment | [optional] 

## Methods

### NewAttachmentSpec

`func NewAttachmentSpec() *AttachmentSpec`

NewAttachmentSpec instantiates a new AttachmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentSpecWithDefaults

`func NewAttachmentSpecWithDefaults() *AttachmentSpec`

NewAttachmentSpecWithDefaults instantiates a new AttachmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AttachmentSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AttachmentSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AttachmentSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AttachmentSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroupName

`func (o *AttachmentSpec) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AttachmentSpec) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AttachmentSpec) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *AttachmentSpec) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetMediaType

`func (o *AttachmentSpec) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AttachmentSpec) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AttachmentSpec) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *AttachmentSpec) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetOwnerName

`func (o *AttachmentSpec) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *AttachmentSpec) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *AttachmentSpec) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *AttachmentSpec) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetPolicyName

`func (o *AttachmentSpec) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AttachmentSpec) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AttachmentSpec) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *AttachmentSpec) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetSize

`func (o *AttachmentSpec) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AttachmentSpec) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AttachmentSpec) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AttachmentSpec) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTags

`func (o *AttachmentSpec) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AttachmentSpec) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AttachmentSpec) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AttachmentSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


