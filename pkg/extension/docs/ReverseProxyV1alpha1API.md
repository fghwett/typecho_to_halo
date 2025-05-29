# \ReverseProxyV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReverseProxy**](ReverseProxyV1alpha1API.md#CreateReverseProxy) | **Post** /apis/plugin.halo.run/v1alpha1/reverseproxies | 
[**DeleteReverseProxy**](ReverseProxyV1alpha1API.md#DeleteReverseProxy) | **Delete** /apis/plugin.halo.run/v1alpha1/reverseproxies/{name} | 
[**GetReverseProxy**](ReverseProxyV1alpha1API.md#GetReverseProxy) | **Get** /apis/plugin.halo.run/v1alpha1/reverseproxies/{name} | 
[**ListReverseProxy**](ReverseProxyV1alpha1API.md#ListReverseProxy) | **Get** /apis/plugin.halo.run/v1alpha1/reverseproxies | 
[**PatchReverseProxy**](ReverseProxyV1alpha1API.md#PatchReverseProxy) | **Patch** /apis/plugin.halo.run/v1alpha1/reverseproxies/{name} | 
[**UpdateReverseProxy**](ReverseProxyV1alpha1API.md#UpdateReverseProxy) | **Put** /apis/plugin.halo.run/v1alpha1/reverseproxies/{name} | 



## CreateReverseProxy

> ReverseProxy CreateReverseProxy(ctx).ReverseProxy(reverseProxy).Execute()





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
	reverseProxy := *openapiclient.NewReverseProxy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // ReverseProxy | Fresh reverseproxy (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReverseProxyV1alpha1API.CreateReverseProxy(context.Background()).ReverseProxy(reverseProxy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReverseProxyV1alpha1API.CreateReverseProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReverseProxy`: ReverseProxy
	fmt.Fprintf(os.Stdout, "Response from `ReverseProxyV1alpha1API.CreateReverseProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReverseProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reverseProxy** | [**ReverseProxy**](ReverseProxy.md) | Fresh reverseproxy | 

### Return type

[**ReverseProxy**](ReverseProxy.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReverseProxy

> DeleteReverseProxy(ctx, name).Execute()





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
	name := "name_example" // string | Name of reverseproxy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReverseProxyV1alpha1API.DeleteReverseProxy(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReverseProxyV1alpha1API.DeleteReverseProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of reverseproxy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReverseProxyRequest struct via the builder pattern


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


## GetReverseProxy

> ReverseProxy GetReverseProxy(ctx, name).Execute()





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
	name := "name_example" // string | Name of reverseproxy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReverseProxyV1alpha1API.GetReverseProxy(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReverseProxyV1alpha1API.GetReverseProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReverseProxy`: ReverseProxy
	fmt.Fprintf(os.Stdout, "Response from `ReverseProxyV1alpha1API.GetReverseProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of reverseproxy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReverseProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReverseProxy**](ReverseProxy.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReverseProxy

> ReverseProxyList ListReverseProxy(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.ReverseProxyV1alpha1API.ListReverseProxy(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReverseProxyV1alpha1API.ListReverseProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReverseProxy`: ReverseProxyList
	fmt.Fprintf(os.Stdout, "Response from `ReverseProxyV1alpha1API.ListReverseProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReverseProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ReverseProxyList**](ReverseProxyList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchReverseProxy

> ReverseProxy PatchReverseProxy(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of reverseproxy
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReverseProxyV1alpha1API.PatchReverseProxy(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReverseProxyV1alpha1API.PatchReverseProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchReverseProxy`: ReverseProxy
	fmt.Fprintf(os.Stdout, "Response from `ReverseProxyV1alpha1API.PatchReverseProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of reverseproxy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchReverseProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**ReverseProxy**](ReverseProxy.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReverseProxy

> ReverseProxy UpdateReverseProxy(ctx, name).ReverseProxy(reverseProxy).Execute()





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
	name := "name_example" // string | Name of reverseproxy
	reverseProxy := *openapiclient.NewReverseProxy("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // ReverseProxy | Updated reverseproxy (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReverseProxyV1alpha1API.UpdateReverseProxy(context.Background(), name).ReverseProxy(reverseProxy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReverseProxyV1alpha1API.UpdateReverseProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReverseProxy`: ReverseProxy
	fmt.Fprintf(os.Stdout, "Response from `ReverseProxyV1alpha1API.UpdateReverseProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of reverseproxy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReverseProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reverseProxy** | [**ReverseProxy**](ReverseProxy.md) | Updated reverseproxy | 

### Return type

[**ReverseProxy**](ReverseProxy.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

