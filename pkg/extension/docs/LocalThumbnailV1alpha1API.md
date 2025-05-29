# \LocalThumbnailV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocalThumbnail**](LocalThumbnailV1alpha1API.md#CreateLocalThumbnail) | **Post** /apis/storage.halo.run/v1alpha1/localthumbnails | 
[**DeleteLocalThumbnail**](LocalThumbnailV1alpha1API.md#DeleteLocalThumbnail) | **Delete** /apis/storage.halo.run/v1alpha1/localthumbnails/{name} | 
[**GetLocalThumbnail**](LocalThumbnailV1alpha1API.md#GetLocalThumbnail) | **Get** /apis/storage.halo.run/v1alpha1/localthumbnails/{name} | 
[**ListLocalThumbnail**](LocalThumbnailV1alpha1API.md#ListLocalThumbnail) | **Get** /apis/storage.halo.run/v1alpha1/localthumbnails | 
[**PatchLocalThumbnail**](LocalThumbnailV1alpha1API.md#PatchLocalThumbnail) | **Patch** /apis/storage.halo.run/v1alpha1/localthumbnails/{name} | 
[**UpdateLocalThumbnail**](LocalThumbnailV1alpha1API.md#UpdateLocalThumbnail) | **Put** /apis/storage.halo.run/v1alpha1/localthumbnails/{name} | 



## CreateLocalThumbnail

> LocalThumbnail CreateLocalThumbnail(ctx).LocalThumbnail(localThumbnail).Execute()





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
	localThumbnail := *openapiclient.NewLocalThumbnail("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewLocalThumbnailSpec("FilePath_example", "ImageSignature_example", "ImageUri_example", "Size_example", "ThumbSignature_example", "ThumbnailUri_example")) // LocalThumbnail | Fresh localthumbnail (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalThumbnailV1alpha1API.CreateLocalThumbnail(context.Background()).LocalThumbnail(localThumbnail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalThumbnailV1alpha1API.CreateLocalThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLocalThumbnail`: LocalThumbnail
	fmt.Fprintf(os.Stdout, "Response from `LocalThumbnailV1alpha1API.CreateLocalThumbnail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **localThumbnail** | [**LocalThumbnail**](LocalThumbnail.md) | Fresh localthumbnail | 

### Return type

[**LocalThumbnail**](LocalThumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocalThumbnail

> DeleteLocalThumbnail(ctx, name).Execute()





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
	name := "name_example" // string | Name of localthumbnail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LocalThumbnailV1alpha1API.DeleteLocalThumbnail(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalThumbnailV1alpha1API.DeleteLocalThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of localthumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalThumbnailRequest struct via the builder pattern


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


## GetLocalThumbnail

> LocalThumbnail GetLocalThumbnail(ctx, name).Execute()





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
	name := "name_example" // string | Name of localthumbnail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalThumbnailV1alpha1API.GetLocalThumbnail(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalThumbnailV1alpha1API.GetLocalThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalThumbnail`: LocalThumbnail
	fmt.Fprintf(os.Stdout, "Response from `LocalThumbnailV1alpha1API.GetLocalThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of localthumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocalThumbnail**](LocalThumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocalThumbnail

> LocalThumbnailList ListLocalThumbnail(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.LocalThumbnailV1alpha1API.ListLocalThumbnail(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalThumbnailV1alpha1API.ListLocalThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLocalThumbnail`: LocalThumbnailList
	fmt.Fprintf(os.Stdout, "Response from `LocalThumbnailV1alpha1API.ListLocalThumbnail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLocalThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**LocalThumbnailList**](LocalThumbnailList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLocalThumbnail

> LocalThumbnail PatchLocalThumbnail(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of localthumbnail
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalThumbnailV1alpha1API.PatchLocalThumbnail(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalThumbnailV1alpha1API.PatchLocalThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchLocalThumbnail`: LocalThumbnail
	fmt.Fprintf(os.Stdout, "Response from `LocalThumbnailV1alpha1API.PatchLocalThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of localthumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLocalThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**LocalThumbnail**](LocalThumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalThumbnail

> LocalThumbnail UpdateLocalThumbnail(ctx, name).LocalThumbnail(localThumbnail).Execute()





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
	name := "name_example" // string | Name of localthumbnail
	localThumbnail := *openapiclient.NewLocalThumbnail("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewLocalThumbnailSpec("FilePath_example", "ImageSignature_example", "ImageUri_example", "Size_example", "ThumbSignature_example", "ThumbnailUri_example")) // LocalThumbnail | Updated localthumbnail (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalThumbnailV1alpha1API.UpdateLocalThumbnail(context.Background(), name).LocalThumbnail(localThumbnail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalThumbnailV1alpha1API.UpdateLocalThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLocalThumbnail`: LocalThumbnail
	fmt.Fprintf(os.Stdout, "Response from `LocalThumbnailV1alpha1API.UpdateLocalThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of localthumbnail | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **localThumbnail** | [**LocalThumbnail**](LocalThumbnail.md) | Updated localthumbnail | 

### Return type

[**LocalThumbnail**](LocalThumbnail.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

