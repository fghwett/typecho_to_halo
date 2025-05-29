# UploadFromUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**PolicyName** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewUploadFromUrlRequest

`func NewUploadFromUrlRequest(policyName string, url string, ) *UploadFromUrlRequest`

NewUploadFromUrlRequest instantiates a new UploadFromUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFromUrlRequestWithDefaults

`func NewUploadFromUrlRequestWithDefaults() *UploadFromUrlRequest`

NewUploadFromUrlRequestWithDefaults instantiates a new UploadFromUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *UploadFromUrlRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadFromUrlRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadFromUrlRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *UploadFromUrlRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetGroupName

`func (o *UploadFromUrlRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UploadFromUrlRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UploadFromUrlRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *UploadFromUrlRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetPolicyName

`func (o *UploadFromUrlRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *UploadFromUrlRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *UploadFromUrlRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetUrl

`func (o *UploadFromUrlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFromUrlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFromUrlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


