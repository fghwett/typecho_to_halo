# \SinglePageV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSinglePage**](SinglePageV1alpha1API.md#CreateSinglePage) | **Post** /apis/content.halo.run/v1alpha1/singlepages | 
[**DeleteSinglePage**](SinglePageV1alpha1API.md#DeleteSinglePage) | **Delete** /apis/content.halo.run/v1alpha1/singlepages/{name} | 
[**GetSinglePage**](SinglePageV1alpha1API.md#GetSinglePage) | **Get** /apis/content.halo.run/v1alpha1/singlepages/{name} | 
[**ListSinglePage**](SinglePageV1alpha1API.md#ListSinglePage) | **Get** /apis/content.halo.run/v1alpha1/singlepages | 
[**PatchSinglePage**](SinglePageV1alpha1API.md#PatchSinglePage) | **Patch** /apis/content.halo.run/v1alpha1/singlepages/{name} | 
[**UpdateSinglePage**](SinglePageV1alpha1API.md#UpdateSinglePage) | **Put** /apis/content.halo.run/v1alpha1/singlepages/{name} | 



## CreateSinglePage

> SinglePage CreateSinglePage(ctx).SinglePage(singlePage).Execute()





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
	singlePage := *openapiclient.NewSinglePage("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewSinglePageSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example")) // SinglePage | Fresh singlepage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1API.CreateSinglePage(context.Background()).SinglePage(singlePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1API.CreateSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSinglePage`: SinglePage
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1API.CreateSinglePage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singlePage** | [**SinglePage**](SinglePage.md) | Fresh singlepage | 

### Return type

[**SinglePage**](SinglePage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSinglePage

> DeleteSinglePage(ctx, name).Execute()





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
	name := "name_example" // string | Name of singlepage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SinglePageV1alpha1API.DeleteSinglePage(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1API.DeleteSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of singlepage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSinglePageRequest struct via the builder pattern


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


## GetSinglePage

> SinglePage GetSinglePage(ctx, name).Execute()





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
	name := "name_example" // string | Name of singlepage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1API.GetSinglePage(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1API.GetSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSinglePage`: SinglePage
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1API.GetSinglePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of singlepage | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SinglePage**](SinglePage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSinglePage

> SinglePageList ListSinglePage(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1API.ListSinglePage(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1API.ListSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSinglePage`: SinglePageList
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1API.ListSinglePage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**SinglePageList**](SinglePageList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSinglePage

> SinglePage PatchSinglePage(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of singlepage
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1API.PatchSinglePage(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1API.PatchSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSinglePage`: SinglePage
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1API.PatchSinglePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of singlepage | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**SinglePage**](SinglePage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSinglePage

> SinglePage UpdateSinglePage(ctx, name).SinglePage(singlePage).Execute()





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
	name := "name_example" // string | Name of singlepage
	singlePage := *openapiclient.NewSinglePage("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewSinglePageSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example")) // SinglePage | Updated singlepage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1API.UpdateSinglePage(context.Background(), name).SinglePage(singlePage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1API.UpdateSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSinglePage`: SinglePage
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1API.UpdateSinglePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of singlepage | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singlePage** | [**SinglePage**](SinglePage.md) | Updated singlepage | 

### Return type

[**SinglePage**](SinglePage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

