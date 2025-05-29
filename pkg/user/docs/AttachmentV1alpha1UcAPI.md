# \AttachmentV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttachmentForPost**](AttachmentV1alpha1UcAPI.md#CreateAttachmentForPost) | **Post** /apis/uc.api.storage.halo.run/v1alpha1/attachments | 
[**ExternalTransferAttachment1**](AttachmentV1alpha1UcAPI.md#ExternalTransferAttachment1) | **Post** /apis/uc.api.storage.halo.run/v1alpha1/attachments/-/upload-from-url | 
[**ListMyAttachments**](AttachmentV1alpha1UcAPI.md#ListMyAttachments) | **Get** /apis/uc.api.storage.halo.run/v1alpha1/attachments | 
[**UploadUcAttachment**](AttachmentV1alpha1UcAPI.md#UploadUcAttachment) | **Post** /apis/uc.api.storage.halo.run/v1alpha1/attachments/-/upload | 



## CreateAttachmentForPost

> Attachment CreateAttachmentForPost(ctx).File(file).WaitForPermalink(waitForPermalink).PostName(postName).SinglePageName(singlePageName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | 
	waitForPermalink := true // bool | Wait for permalink. (optional)
	postName := "postName_example" // string | Post name. (optional)
	singlePageName := "singlePageName_example" // string | Single page name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentV1alpha1UcAPI.CreateAttachmentForPost(context.Background()).File(file).WaitForPermalink(waitForPermalink).PostName(postName).SinglePageName(singlePageName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentV1alpha1UcAPI.CreateAttachmentForPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAttachmentForPost`: Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentV1alpha1UcAPI.CreateAttachmentForPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttachmentForPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **waitForPermalink** | **bool** | Wait for permalink. | 
 **postName** | **string** | Post name. | 
 **singlePageName** | **string** | Single page name. | 

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


## ExternalTransferAttachment1

> Attachment ExternalTransferAttachment1(ctx).UploadFromUrlRequest(uploadFromUrlRequest).WaitForPermalink(waitForPermalink).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	uploadFromUrlRequest := *openapiclient.NewUploadFromUrlRequest("Url_example") // UploadFromUrlRequest | 
	waitForPermalink := true // bool | Wait for permalink. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentV1alpha1UcAPI.ExternalTransferAttachment1(context.Background()).UploadFromUrlRequest(uploadFromUrlRequest).WaitForPermalink(waitForPermalink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentV1alpha1UcAPI.ExternalTransferAttachment1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalTransferAttachment1`: Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentV1alpha1UcAPI.ExternalTransferAttachment1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalTransferAttachment1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadFromUrlRequest** | [**UploadFromUrlRequest**](UploadFromUrlRequest.md) |  | 
 **waitForPermalink** | **bool** | Wait for permalink. | 

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


## ListMyAttachments

> AttachmentList ListMyAttachments(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Ungrouped(ungrouped).Keyword(keyword).Accepts(accepts).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
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
	resp, r, err := apiClient.AttachmentV1alpha1UcAPI.ListMyAttachments(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Ungrouped(ungrouped).Keyword(keyword).Accepts(accepts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentV1alpha1UcAPI.ListMyAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyAttachments`: AttachmentList
	fmt.Fprintf(os.Stdout, "Response from `AttachmentV1alpha1UcAPI.ListMyAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyAttachmentsRequest struct via the builder pattern


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


## UploadUcAttachment

> Attachment UploadUcAttachment(ctx).File(file).FormData(formData).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | 
	formData := *openapiclient.NewUcUploadRequestFormData() // UcUploadRequestFormData |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentV1alpha1UcAPI.UploadUcAttachment(context.Background()).File(file).FormData(formData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentV1alpha1UcAPI.UploadUcAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadUcAttachment`: Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentV1alpha1UcAPI.UploadUcAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadUcAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **formData** | [**UcUploadRequestFormData**](UcUploadRequestFormData.md) |  | 

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

