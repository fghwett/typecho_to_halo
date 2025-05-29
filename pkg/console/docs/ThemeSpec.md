# ThemeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**Author**](Author.md) |  | 
**ConfigMapName** | Pointer to **string** |  | [optional] 
**CustomTemplates** | Pointer to [**CustomTemplates**](CustomTemplates.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**Homepage** | Pointer to **string** |  | [optional] 
**Issues** | Pointer to **string** |  | [optional] 
**License** | Pointer to [**[]License**](License.md) |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Require** | Pointer to **string** | Deprecated, use &#x60;requires&#x60; instead. | [optional] 
**Requires** | Pointer to **string** |  | [optional] 
**SettingName** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewThemeSpec

`func NewThemeSpec(author Author, displayName string, version string, ) *ThemeSpec`

NewThemeSpec instantiates a new ThemeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThemeSpecWithDefaults

`func NewThemeSpecWithDefaults() *ThemeSpec`

NewThemeSpecWithDefaults instantiates a new ThemeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ThemeSpec) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ThemeSpec) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ThemeSpec) SetAuthor(v Author)`

SetAuthor sets Author field to given value.


### GetConfigMapName

`func (o *ThemeSpec) GetConfigMapName() string`

GetConfigMapName returns the ConfigMapName field if non-nil, zero value otherwise.

### GetConfigMapNameOk

`func (o *ThemeSpec) GetConfigMapNameOk() (*string, bool)`

GetConfigMapNameOk returns a tuple with the ConfigMapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapName

`func (o *ThemeSpec) SetConfigMapName(v string)`

SetConfigMapName sets ConfigMapName field to given value.

### HasConfigMapName

`func (o *ThemeSpec) HasConfigMapName() bool`

HasConfigMapName returns a boolean if a field has been set.

### GetCustomTemplates

`func (o *ThemeSpec) GetCustomTemplates() CustomTemplates`

GetCustomTemplates returns the CustomTemplates field if non-nil, zero value otherwise.

### GetCustomTemplatesOk

`func (o *ThemeSpec) GetCustomTemplatesOk() (*CustomTemplates, bool)`

GetCustomTemplatesOk returns a tuple with the CustomTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTemplates

`func (o *ThemeSpec) SetCustomTemplates(v CustomTemplates)`

SetCustomTemplates sets CustomTemplates field to given value.

### HasCustomTemplates

`func (o *ThemeSpec) HasCustomTemplates() bool`

HasCustomTemplates returns a boolean if a field has been set.

### GetDescription

`func (o *ThemeSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThemeSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThemeSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThemeSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ThemeSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ThemeSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ThemeSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHomepage

`func (o *ThemeSpec) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *ThemeSpec) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *ThemeSpec) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *ThemeSpec) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetIssues

`func (o *ThemeSpec) GetIssues() string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ThemeSpec) GetIssuesOk() (*string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ThemeSpec) SetIssues(v string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ThemeSpec) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetLicense

`func (o *ThemeSpec) GetLicense() []License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ThemeSpec) GetLicenseOk() (*[]License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ThemeSpec) SetLicense(v []License)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ThemeSpec) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLogo

`func (o *ThemeSpec) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ThemeSpec) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ThemeSpec) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ThemeSpec) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetRepo

`func (o *ThemeSpec) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ThemeSpec) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ThemeSpec) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ThemeSpec) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRequire

`func (o *ThemeSpec) GetRequire() string`

GetRequire returns the Require field if non-nil, zero value otherwise.

### GetRequireOk

`func (o *ThemeSpec) GetRequireOk() (*string, bool)`

GetRequireOk returns a tuple with the Require field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequire

`func (o *ThemeSpec) SetRequire(v string)`

SetRequire sets Require field to given value.

### HasRequire

`func (o *ThemeSpec) HasRequire() bool`

HasRequire returns a boolean if a field has been set.

### GetRequires

`func (o *ThemeSpec) GetRequires() string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *ThemeSpec) GetRequiresOk() (*string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *ThemeSpec) SetRequires(v string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *ThemeSpec) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### GetSettingName

`func (o *ThemeSpec) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *ThemeSpec) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *ThemeSpec) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *ThemeSpec) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetVersion

`func (o *ThemeSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThemeSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThemeSpec) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetWebsite

`func (o *ThemeSpec) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ThemeSpec) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ThemeSpec) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ThemeSpec) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


