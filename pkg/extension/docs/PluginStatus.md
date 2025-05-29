# PluginStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to  |  | [optional] 
**Entry** | Pointer to **string** |  | [optional] 
**LastProbeState** | Pointer to **string** |  | [optional] 
**LastStartTime** | Pointer to **time.Time** |  | [optional] 
**LoadLocation** | Pointer to **string** | Load location of the plugin, often a path. | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Stylesheet** | Pointer to **string** |  | [optional] 

## Methods

### NewPluginStatus

`func NewPluginStatus() *PluginStatus`

NewPluginStatus instantiates a new PluginStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginStatusWithDefaults

`func NewPluginStatusWithDefaults() *PluginStatus`

NewPluginStatusWithDefaults instantiates a new PluginStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *PluginStatus) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PluginStatus) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PluginStatus) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PluginStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetEntry

`func (o *PluginStatus) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *PluginStatus) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *PluginStatus) SetEntry(v string)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *PluginStatus) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetLastProbeState

`func (o *PluginStatus) GetLastProbeState() string`

GetLastProbeState returns the LastProbeState field if non-nil, zero value otherwise.

### GetLastProbeStateOk

`func (o *PluginStatus) GetLastProbeStateOk() (*string, bool)`

GetLastProbeStateOk returns a tuple with the LastProbeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbeState

`func (o *PluginStatus) SetLastProbeState(v string)`

SetLastProbeState sets LastProbeState field to given value.

### HasLastProbeState

`func (o *PluginStatus) HasLastProbeState() bool`

HasLastProbeState returns a boolean if a field has been set.

### GetLastStartTime

`func (o *PluginStatus) GetLastStartTime() time.Time`

GetLastStartTime returns the LastStartTime field if non-nil, zero value otherwise.

### GetLastStartTimeOk

`func (o *PluginStatus) GetLastStartTimeOk() (*time.Time, bool)`

GetLastStartTimeOk returns a tuple with the LastStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStartTime

`func (o *PluginStatus) SetLastStartTime(v time.Time)`

SetLastStartTime sets LastStartTime field to given value.

### HasLastStartTime

`func (o *PluginStatus) HasLastStartTime() bool`

HasLastStartTime returns a boolean if a field has been set.

### GetLoadLocation

`func (o *PluginStatus) GetLoadLocation() string`

GetLoadLocation returns the LoadLocation field if non-nil, zero value otherwise.

### GetLoadLocationOk

`func (o *PluginStatus) GetLoadLocationOk() (*string, bool)`

GetLoadLocationOk returns a tuple with the LoadLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLocation

`func (o *PluginStatus) SetLoadLocation(v string)`

SetLoadLocation sets LoadLocation field to given value.

### HasLoadLocation

`func (o *PluginStatus) HasLoadLocation() bool`

HasLoadLocation returns a boolean if a field has been set.

### GetLogo

`func (o *PluginStatus) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PluginStatus) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PluginStatus) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PluginStatus) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPhase

`func (o *PluginStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PluginStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PluginStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PluginStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetStylesheet

`func (o *PluginStatus) GetStylesheet() string`

GetStylesheet returns the Stylesheet field if non-nil, zero value otherwise.

### GetStylesheetOk

`func (o *PluginStatus) GetStylesheetOk() (*string, bool)`

GetStylesheetOk returns a tuple with the Stylesheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStylesheet

`func (o *PluginStatus) SetStylesheet(v string)`

SetStylesheet sets Stylesheet field to given value.

### HasStylesheet

`func (o *PluginStatus) HasStylesheet() bool`

HasStylesheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


