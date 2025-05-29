# \AttachmentV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalTransferAttachment**](AttachmentV1alpha1ConsoleAPI.md#ExternalTransferAttachment) | **Post** /apis/api.console.halo.run/v1alpha1/attachments/-/upload-from-url | 
[**SearchAttachments**](AttachmentV1alpha1ConsoleAPI.md#SearchAttachments) | **Get** /apis/api.console.halo.run/v1alpha1/attachments | 
[**UploadAttachment**](AttachmentV1alpha1ConsoleAPI.md#UploadAttachment) | **Post** /apis/api.console.halo.run/v1alpha1/attachments/upload | 



## ExternalTransferAttachment

> Attachment ExternalTransferAttachment(ctx).UploadFromUrlRequest(uploadFromUrlRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-console"
)

func main() {
	uploadFromUrlRequest := *openapiclient.NewUploadFromUrlRequest("PolicyName_example", "Url_example") // UploadFromUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentV1alpha1ConsoleAPI.ExternalTransferAttachment(context.Background()).UploadFromUrlRequest(uploadFromUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentV1alpha1ConsoleAPI.ExternalTransferAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalTransferAttachment`: Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentV1alpha1ConsoleAPI.ExternalTransferAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalTransferAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadFromUrlRequest** | [**UploadFromUrlRequest**](UploadFromUrlRequest.md) |  | 

### Return type

[**Attachment**](Attachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAttachments

> AttachmentList SearchAttachments(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Ungrouped(ungrouped).Keyword(keyword).Accepts(accepts).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-console"
)

func main() {
	page := int32(56) // int32 | Page number. Default is 0. (optional)
	size := int32(56) // int32 | Size number. Default is 0. (optional)
	labelSelector := []string{"Inner_example"} // []string | Label selector. e.g.: hidden!=true (optional)
	fieldSelector := []string{"Inner_example"} // []string | Field selector. e.g.: metadata.name==halo (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	ungrouped := true // bool | Filter attachments without group. This parameter will ignore group parameter. (optional)
	keyword := "keyword_example" // string | Keyword for searching. (optional)
	accepts := []string{"Inner_example"} // []string | Acceptable media types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentV1alpha1ConsoleAPI.SearchAttachments(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Ungrouped(ungrouped).Keyword(keyword).Accepts(accepts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentV1alpha1ConsoleAPI.SearchAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAttachments`: AttachmentList
	fmt.Fprintf(os.Stdout, "Response from `AttachmentV1alpha1ConsoleAPI.SearchAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **ungrouped** | **bool** | Filter attachments without group. This parameter will ignore group parameter. | 
 **keyword** | **string** | Keyword for searching. | 
 **accepts** | **[]string** | Acceptable media types. | 

### Return type

[**AttachmentList**](AttachmentList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAttachment

> Attachment UploadAttachment(ctx).File(file).PolicyName(policyName).GroupName(groupName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-console"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | 
	policyName := "policyName_example" // string | Storage policy name
	groupName := "groupName_example" // string | The name of the group to which the attachment belongs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentV1alpha1ConsoleAPI.UploadAttachment(context.Background()).File(file).PolicyName(policyName).GroupName(groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentV1alpha1ConsoleAPI.UploadAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadAttachment`: Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentV1alpha1ConsoleAPI.UploadAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **policyName** | **string** | Storage policy name | 
 **groupName** | **string** | The name of the group to which the attachment belongs | 

### Return type

[**Attachment**](Attachment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

