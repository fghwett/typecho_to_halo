# \ThumbnailV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateThumbnail**](ThumbnailV1alpha1API.md#CreateThumbnail) | **Post** /apis/storage.halo.run/v1alpha1/thumbnails | 
[**DeleteThumbnail**](ThumbnailV1alpha1API.md#DeleteThumbnail) | **Delete** /apis/storage.halo.run/v1alpha1/thumbnails/{name} | 
[**GetThumbnail**](ThumbnailV1alpha1API.md#GetThumbnail) | **Get** /apis/storage.halo.run/v1alpha1/thumbnails/{name} | 
[**ListThumbnail**](ThumbnailV1alpha1API.md#ListThumbnail) | **Get** /apis/storage.halo.run/v1alpha1/thumbnails | 
[**PatchThumbnail**](ThumbnailV1alpha1API.md#PatchThumbnail) | **Patch** /apis/storage.halo.run/v1alpha1/thumbnails/{name} | 
[**UpdateThumbnail**](ThumbnailV1alpha1API.md#UpdateThumbnail) | **Put** /apis/storage.halo.run/v1alpha1/thumbnails/{name} | 



## CreateThumbnail

> Thumbnail CreateThumbnail(ctx).Thumbnail(thumbnail).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-extension"
)

func main() {
	thumbnail := *openapiclient.NewThumbnail("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewThumbnailSpec("ImageSignature_example", "ImageUri_example", "Size_example", "ThumbnailUri_example")) // Thumbnail | Fresh thumbnail (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThumbnailV1alpha1API.CreateThumbnail(context.Background()).Thumbnail(thumbnail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailV1alpha1API.CreateThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThumbnail`: Thumbnail
	fmt.Fprintf(os.Stdout, "Response from `ThumbnailV1alpha1API.CreateThumbnail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thumbnail** | [**Thumbnail**](Thumbnail.md) | Fresh thumbnail | 

### Return type

[**Thumbnail**](Thumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThumbnail

> DeleteThumbnail(ctx, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-extension"
)

func main() {
	name := "name_example" // string | Name of thumbnail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThumbnailV1alpha1API.DeleteThumbnail(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailV1alpha1API.DeleteThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of thumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnail

> Thumbnail GetThumbnail(ctx, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-extension"
)

func main() {
	name := "name_example" // string | Name of thumbnail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThumbnailV1alpha1API.GetThumbnail(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailV1alpha1API.GetThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThumbnail`: Thumbnail
	fmt.Fprintf(os.Stdout, "Response from `ThumbnailV1alpha1API.GetThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of thumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Thumbnail**](Thumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListThumbnail

> ThumbnailList ListThumbnail(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-extension"
)

func main() {
	page := int32(56) // int32 | Page number. Default is 0. (optional)
	size := int32(56) // int32 | Size number. Default is 0. (optional)
	labelSelector := []string{"Inner_example"} // []string | Label selector. e.g.: hidden!=true (optional)
	fieldSelector := []string{"Inner_example"} // []string | Field selector. e.g.: metadata.name==halo (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThumbnailV1alpha1API.ListThumbnail(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailV1alpha1API.ListThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListThumbnail`: ThumbnailList
	fmt.Fprintf(os.Stdout, "Response from `ThumbnailV1alpha1API.ListThumbnail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ThumbnailList**](ThumbnailList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchThumbnail

> Thumbnail PatchThumbnail(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-extension"
)

func main() {
	name := "name_example" // string | Name of thumbnail
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThumbnailV1alpha1API.PatchThumbnail(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailV1alpha1API.PatchThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchThumbnail`: Thumbnail
	fmt.Fprintf(os.Stdout, "Response from `ThumbnailV1alpha1API.PatchThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of thumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**Thumbnail**](Thumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThumbnail

> Thumbnail UpdateThumbnail(ctx, name).Thumbnail(thumbnail).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-extension"
)

func main() {
	name := "name_example" // string | Name of thumbnail
	thumbnail := *openapiclient.NewThumbnail("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewThumbnailSpec("ImageSignature_example", "ImageUri_example", "Size_example", "ThumbnailUri_example")) // Thumbnail | Updated thumbnail (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThumbnailV1alpha1API.UpdateThumbnail(context.Background(), name).Thumbnail(thumbnail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailV1alpha1API.UpdateThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateThumbnail`: Thumbnail
	fmt.Fprintf(os.Stdout, "Response from `ThumbnailV1alpha1API.UpdateThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of thumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbnail** | [**Thumbnail**](Thumbnail.md) | Updated thumbnail | 

### Return type

[**Thumbnail**](Thumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

