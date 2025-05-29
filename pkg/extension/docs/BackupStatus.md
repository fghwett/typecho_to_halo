# BackupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTimestamp** | Pointer to **time.Time** |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**StartTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBackupStatus

`func NewBackupStatus() *BackupStatus`

NewBackupStatus instantiates a new BackupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStatusWithDefaults

`func NewBackupStatusWithDefaults() *BackupStatus`

NewBackupStatusWithDefaults instantiates a new BackupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTimestamp

`func (o *BackupStatus) GetCompletionTimestamp() time.Time`

GetCompletionTimestamp returns the CompletionTimestamp field if non-nil, zero value otherwise.

### GetCompletionTimestampOk

`func (o *BackupStatus) GetCompletionTimestampOk() (*time.Time, bool)`

GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTimestamp

`func (o *BackupStatus) SetCompletionTimestamp(v time.Time)`

SetCompletionTimestamp sets CompletionTimestamp field to given value.

### HasCompletionTimestamp

`func (o *BackupStatus) HasCompletionTimestamp() bool`

HasCompletionTimestamp returns a boolean if a field has been set.

### GetFailureMessage

`func (o *BackupStatus) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *BackupStatus) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *BackupStatus) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *BackupStatus) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetFailureReason

`func (o *BackupStatus) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *BackupStatus) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *BackupStatus) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *BackupStatus) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFilename

`func (o *BackupStatus) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *BackupStatus) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *BackupStatus) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *BackupStatus) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPhase

`func (o *BackupStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *BackupStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *BackupStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *BackupStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetSize

`func (o *BackupStatus) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BackupStatus) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BackupStatus) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BackupStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *BackupStatus) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *BackupStatus) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *BackupStatus) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *BackupStatus) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


