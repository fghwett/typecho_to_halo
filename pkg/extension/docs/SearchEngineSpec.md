# SearchEngineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**Logo** | Pointer to **string** |  | [optional] 
**PostSearchImpl** | Pointer to **string** |  | [optional] 
**SettingRef** | Pointer to [**Ref**](Ref.md) |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchEngineSpec

`func NewSearchEngineSpec(displayName string, ) *SearchEngineSpec`

NewSearchEngineSpec instantiates a new SearchEngineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEngineSpecWithDefaults

`func NewSearchEngineSpecWithDefaults() *SearchEngineSpec`

NewSearchEngineSpecWithDefaults instantiates a new SearchEngineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SearchEngineSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchEngineSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchEngineSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchEngineSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *SearchEngineSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SearchEngineSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SearchEngineSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLogo

`func (o *SearchEngineSpec) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SearchEngineSpec) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SearchEngineSpec) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SearchEngineSpec) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPostSearchImpl

`func (o *SearchEngineSpec) GetPostSearchImpl() string`

GetPostSearchImpl returns the PostSearchImpl field if non-nil, zero value otherwise.

### GetPostSearchImplOk

`func (o *SearchEngineSpec) GetPostSearchImplOk() (*string, bool)`

GetPostSearchImplOk returns a tuple with the PostSearchImpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSearchImpl

`func (o *SearchEngineSpec) SetPostSearchImpl(v string)`

SetPostSearchImpl sets PostSearchImpl field to given value.

### HasPostSearchImpl

`func (o *SearchEngineSpec) HasPostSearchImpl() bool`

HasPostSearchImpl returns a boolean if a field has been set.

### GetSettingRef

`func (o *SearchEngineSpec) GetSettingRef() Ref`

GetSettingRef returns the SettingRef field if non-nil, zero value otherwise.

### GetSettingRefOk

`func (o *SearchEngineSpec) GetSettingRefOk() (*Ref, bool)`

GetSettingRefOk returns a tuple with the SettingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingRef

`func (o *SearchEngineSpec) SetSettingRef(v Ref)`

SetSettingRef sets SettingRef field to given value.

### HasSettingRef

`func (o *SearchEngineSpec) HasSettingRef() bool`

HasSettingRef returns a boolean if a field has been set.

### GetWebsite

`func (o *SearchEngineSpec) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *SearchEngineSpec) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *SearchEngineSpec) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *SearchEngineSpec) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


