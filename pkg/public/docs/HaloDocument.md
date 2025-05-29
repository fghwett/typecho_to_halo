# HaloDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Content** | **string** |  | 
**CreationTimestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Exposed** | Pointer to **bool** |  | [optional] 
**Id** | **string** |  | 
**MetadataName** | **string** |  | 
**OwnerName** | **string** |  | 
**Permalink** | **string** |  | 
**Published** | Pointer to **bool** |  | [optional] 
**Recycled** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | **string** |  | 
**Type** | **string** |  | 
**UpdateTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewHaloDocument

`func NewHaloDocument(content string, id string, metadataName string, ownerName string, permalink string, title string, type_ string, ) *HaloDocument`

NewHaloDocument instantiates a new HaloDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHaloDocumentWithDefaults

`func NewHaloDocumentWithDefaults() *HaloDocument`

NewHaloDocumentWithDefaults instantiates a new HaloDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *HaloDocument) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *HaloDocument) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *HaloDocument) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *HaloDocument) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCategories

`func (o *HaloDocument) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *HaloDocument) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *HaloDocument) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *HaloDocument) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetContent

`func (o *HaloDocument) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *HaloDocument) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *HaloDocument) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreationTimestamp

`func (o *HaloDocument) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *HaloDocument) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *HaloDocument) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *HaloDocument) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *HaloDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HaloDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HaloDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HaloDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExposed

`func (o *HaloDocument) GetExposed() bool`

GetExposed returns the Exposed field if non-nil, zero value otherwise.

### GetExposedOk

`func (o *HaloDocument) GetExposedOk() (*bool, bool)`

GetExposedOk returns a tuple with the Exposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposed

`func (o *HaloDocument) SetExposed(v bool)`

SetExposed sets Exposed field to given value.

### HasExposed

`func (o *HaloDocument) HasExposed() bool`

HasExposed returns a boolean if a field has been set.

### GetId

`func (o *HaloDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HaloDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HaloDocument) SetId(v string)`

SetId sets Id field to given value.


### GetMetadataName

`func (o *HaloDocument) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *HaloDocument) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *HaloDocument) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.


### GetOwnerName

`func (o *HaloDocument) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *HaloDocument) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *HaloDocument) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetPermalink

`func (o *HaloDocument) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *HaloDocument) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *HaloDocument) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.


### GetPublished

`func (o *HaloDocument) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *HaloDocument) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *HaloDocument) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *HaloDocument) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRecycled

`func (o *HaloDocument) GetRecycled() bool`

GetRecycled returns the Recycled field if non-nil, zero value otherwise.

### GetRecycledOk

`func (o *HaloDocument) GetRecycledOk() (*bool, bool)`

GetRecycledOk returns a tuple with the Recycled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycled

`func (o *HaloDocument) SetRecycled(v bool)`

SetRecycled sets Recycled field to given value.

### HasRecycled

`func (o *HaloDocument) HasRecycled() bool`

HasRecycled returns a boolean if a field has been set.

### GetTags

`func (o *HaloDocument) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HaloDocument) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HaloDocument) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HaloDocument) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *HaloDocument) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HaloDocument) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HaloDocument) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *HaloDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HaloDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HaloDocument) SetType(v string)`

SetType sets Type field to given value.


### GetUpdateTimestamp

`func (o *HaloDocument) GetUpdateTimestamp() time.Time`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *HaloDocument) GetUpdateTimestampOk() (*time.Time, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *HaloDocument) SetUpdateTimestamp(v time.Time)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *HaloDocument) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


