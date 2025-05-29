# \AnnotationSettingV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnnotationSetting**](AnnotationSettingV1alpha1API.md#CreateAnnotationSetting) | **Post** /api/v1alpha1/annotationsettings | 
[**DeleteAnnotationSetting**](AnnotationSettingV1alpha1API.md#DeleteAnnotationSetting) | **Delete** /api/v1alpha1/annotationsettings/{name} | 
[**GetAnnotationSetting**](AnnotationSettingV1alpha1API.md#GetAnnotationSetting) | **Get** /api/v1alpha1/annotationsettings/{name} | 
[**ListAnnotationSetting**](AnnotationSettingV1alpha1API.md#ListAnnotationSetting) | **Get** /api/v1alpha1/annotationsettings | 
[**PatchAnnotationSetting**](AnnotationSettingV1alpha1API.md#PatchAnnotationSetting) | **Patch** /api/v1alpha1/annotationsettings/{name} | 
[**UpdateAnnotationSetting**](AnnotationSettingV1alpha1API.md#UpdateAnnotationSetting) | **Put** /api/v1alpha1/annotationsettings/{name} | 



## CreateAnnotationSetting

> AnnotationSetting CreateAnnotationSetting(ctx).AnnotationSetting(annotationSetting).Execute()





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
	annotationSetting := *openapiclient.NewAnnotationSetting("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewAnnotationSettingSpec([]map[string]interface{}{map[string]interface{}(123)}, *openapiclient.NewGroupKind())) // AnnotationSetting | Fresh annotationsetting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnotationSettingV1alpha1API.CreateAnnotationSetting(context.Background()).AnnotationSetting(annotationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationSettingV1alpha1API.CreateAnnotationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnnotationSetting`: AnnotationSetting
	fmt.Fprintf(os.Stdout, "Response from `AnnotationSettingV1alpha1API.CreateAnnotationSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnnotationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **annotationSetting** | [**AnnotationSetting**](AnnotationSetting.md) | Fresh annotationsetting | 

### Return type

[**AnnotationSetting**](AnnotationSetting.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnnotationSetting

> DeleteAnnotationSetting(ctx, name).Execute()





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
	name := "name_example" // string | Name of annotationsetting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnotationSettingV1alpha1API.DeleteAnnotationSetting(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationSettingV1alpha1API.DeleteAnnotationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of annotationsetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationSettingRequest struct via the builder pattern


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


## GetAnnotationSetting

> AnnotationSetting GetAnnotationSetting(ctx, name).Execute()





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
	name := "name_example" // string | Name of annotationsetting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnotationSettingV1alpha1API.GetAnnotationSetting(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationSettingV1alpha1API.GetAnnotationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnnotationSetting`: AnnotationSetting
	fmt.Fprintf(os.Stdout, "Response from `AnnotationSettingV1alpha1API.GetAnnotationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of annotationsetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnotationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnnotationSetting**](AnnotationSetting.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnnotationSetting

> AnnotationSettingList ListAnnotationSetting(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.AnnotationSettingV1alpha1API.ListAnnotationSetting(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationSettingV1alpha1API.ListAnnotationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnnotationSetting`: AnnotationSettingList
	fmt.Fprintf(os.Stdout, "Response from `AnnotationSettingV1alpha1API.ListAnnotationSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAnnotationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**AnnotationSettingList**](AnnotationSettingList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAnnotationSetting

> AnnotationSetting PatchAnnotationSetting(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of annotationsetting
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnotationSettingV1alpha1API.PatchAnnotationSetting(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationSettingV1alpha1API.PatchAnnotationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAnnotationSetting`: AnnotationSetting
	fmt.Fprintf(os.Stdout, "Response from `AnnotationSettingV1alpha1API.PatchAnnotationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of annotationsetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAnnotationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**AnnotationSetting**](AnnotationSetting.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnnotationSetting

> AnnotationSetting UpdateAnnotationSetting(ctx, name).AnnotationSetting(annotationSetting).Execute()





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
	name := "name_example" // string | Name of annotationsetting
	annotationSetting := *openapiclient.NewAnnotationSetting("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewAnnotationSettingSpec([]map[string]interface{}{map[string]interface{}(123)}, *openapiclient.NewGroupKind())) // AnnotationSetting | Updated annotationsetting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnotationSettingV1alpha1API.UpdateAnnotationSetting(context.Background(), name).AnnotationSetting(annotationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationSettingV1alpha1API.UpdateAnnotationSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnnotationSetting`: AnnotationSetting
	fmt.Fprintf(os.Stdout, "Response from `AnnotationSettingV1alpha1API.UpdateAnnotationSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of annotationsetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnnotationSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **annotationSetting** | [**AnnotationSetting**](AnnotationSetting.md) | Updated annotationsetting | 

### Return type

[**AnnotationSetting**](AnnotationSetting.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

