# PluginSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**PluginAuthor**](PluginAuthor.md) |  | [optional] 
**ConfigMapName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Homepage** | Pointer to **string** |  | [optional] 
**Issues** | Pointer to **string** |  | [optional] 
**License** | Pointer to [**[]License**](License.md) |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**PluginClass** | Pointer to **string** |  | [optional] 
**PluginDependencies** | Pointer to **map[string]string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Requires** | Pointer to **string** |  | [optional] 
**SettingName** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewPluginSpec

`func NewPluginSpec(version string, ) *PluginSpec`

NewPluginSpec instantiates a new PluginSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginSpecWithDefaults

`func NewPluginSpecWithDefaults() *PluginSpec`

NewPluginSpecWithDefaults instantiates a new PluginSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *PluginSpec) GetAuthor() PluginAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PluginSpec) GetAuthorOk() (*PluginAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PluginSpec) SetAuthor(v PluginAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PluginSpec) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetConfigMapName

`func (o *PluginSpec) GetConfigMapName() string`

GetConfigMapName returns the ConfigMapName field if non-nil, zero value otherwise.

### GetConfigMapNameOk

`func (o *PluginSpec) GetConfigMapNameOk() (*string, bool)`

GetConfigMapNameOk returns a tuple with the ConfigMapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapName

`func (o *PluginSpec) SetConfigMapName(v string)`

SetConfigMapName sets ConfigMapName field to given value.

### HasConfigMapName

`func (o *PluginSpec) HasConfigMapName() bool`

HasConfigMapName returns a boolean if a field has been set.

### GetDescription

`func (o *PluginSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PluginSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PluginSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PluginSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PluginSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *PluginSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PluginSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PluginSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PluginSpec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHomepage

`func (o *PluginSpec) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *PluginSpec) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *PluginSpec) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *PluginSpec) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetIssues

`func (o *PluginSpec) GetIssues() string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *PluginSpec) GetIssuesOk() (*string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *PluginSpec) SetIssues(v string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *PluginSpec) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetLicense

`func (o *PluginSpec) GetLicense() []License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PluginSpec) GetLicenseOk() (*[]License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PluginSpec) SetLicense(v []License)`

SetLicense sets License field to given value.

### HasLicense

`func (o *PluginSpec) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLogo

`func (o *PluginSpec) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PluginSpec) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PluginSpec) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PluginSpec) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPluginClass

`func (o *PluginSpec) GetPluginClass() string`

GetPluginClass returns the PluginClass field if non-nil, zero value otherwise.

### GetPluginClassOk

`func (o *PluginSpec) GetPluginClassOk() (*string, bool)`

GetPluginClassOk returns a tuple with the PluginClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginClass

`func (o *PluginSpec) SetPluginClass(v string)`

SetPluginClass sets PluginClass field to given value.

### HasPluginClass

`func (o *PluginSpec) HasPluginClass() bool`

HasPluginClass returns a boolean if a field has been set.

### GetPluginDependencies

`func (o *PluginSpec) GetPluginDependencies() map[string]string`

GetPluginDependencies returns the PluginDependencies field if non-nil, zero value otherwise.

### GetPluginDependenciesOk

`func (o *PluginSpec) GetPluginDependenciesOk() (*map[string]string, bool)`

GetPluginDependenciesOk returns a tuple with the PluginDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDependencies

`func (o *PluginSpec) SetPluginDependencies(v map[string]string)`

SetPluginDependencies sets PluginDependencies field to given value.

### HasPluginDependencies

`func (o *PluginSpec) HasPluginDependencies() bool`

HasPluginDependencies returns a boolean if a field has been set.

### GetRepo

`func (o *PluginSpec) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PluginSpec) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PluginSpec) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PluginSpec) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRequires

`func (o *PluginSpec) GetRequires() string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *PluginSpec) GetRequiresOk() (*string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *PluginSpec) SetRequires(v string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *PluginSpec) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### GetSettingName

`func (o *PluginSpec) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *PluginSpec) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *PluginSpec) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *PluginSpec) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetVersion

`func (o *PluginSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginSpec) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


