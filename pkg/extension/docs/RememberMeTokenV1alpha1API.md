# \RememberMeTokenV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRememberMeToken**](RememberMeTokenV1alpha1API.md#CreateRememberMeToken) | **Post** /apis/security.halo.run/v1alpha1/remembermetokens | 
[**DeleteRememberMeToken**](RememberMeTokenV1alpha1API.md#DeleteRememberMeToken) | **Delete** /apis/security.halo.run/v1alpha1/remembermetokens/{name} | 
[**GetRememberMeToken**](RememberMeTokenV1alpha1API.md#GetRememberMeToken) | **Get** /apis/security.halo.run/v1alpha1/remembermetokens/{name} | 
[**ListRememberMeToken**](RememberMeTokenV1alpha1API.md#ListRememberMeToken) | **Get** /apis/security.halo.run/v1alpha1/remembermetokens | 
[**PatchRememberMeToken**](RememberMeTokenV1alpha1API.md#PatchRememberMeToken) | **Patch** /apis/security.halo.run/v1alpha1/remembermetokens/{name} | 
[**UpdateRememberMeToken**](RememberMeTokenV1alpha1API.md#UpdateRememberMeToken) | **Put** /apis/security.halo.run/v1alpha1/remembermetokens/{name} | 



## CreateRememberMeToken

> RememberMeToken CreateRememberMeToken(ctx).RememberMeToken(rememberMeToken).Execute()





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
	rememberMeToken := *openapiclient.NewRememberMeToken("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewRememberMeTokenSpec("Series_example", "TokenValue_example", "Username_example")) // RememberMeToken | Fresh remembermetoken (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RememberMeTokenV1alpha1API.CreateRememberMeToken(context.Background()).RememberMeToken(rememberMeToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RememberMeTokenV1alpha1API.CreateRememberMeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRememberMeToken`: RememberMeToken
	fmt.Fprintf(os.Stdout, "Response from `RememberMeTokenV1alpha1API.CreateRememberMeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRememberMeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rememberMeToken** | [**RememberMeToken**](RememberMeToken.md) | Fresh remembermetoken | 

### Return type

[**RememberMeToken**](RememberMeToken.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRememberMeToken

> DeleteRememberMeToken(ctx, name).Execute()





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
	name := "name_example" // string | Name of remembermetoken

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RememberMeTokenV1alpha1API.DeleteRememberMeToken(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RememberMeTokenV1alpha1API.DeleteRememberMeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of remembermetoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRememberMeTokenRequest struct via the builder pattern


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


## GetRememberMeToken

> RememberMeToken GetRememberMeToken(ctx, name).Execute()





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
	name := "name_example" // string | Name of remembermetoken

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RememberMeTokenV1alpha1API.GetRememberMeToken(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RememberMeTokenV1alpha1API.GetRememberMeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRememberMeToken`: RememberMeToken
	fmt.Fprintf(os.Stdout, "Response from `RememberMeTokenV1alpha1API.GetRememberMeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of remembermetoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRememberMeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RememberMeToken**](RememberMeToken.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRememberMeToken

> RememberMeTokenList ListRememberMeToken(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.RememberMeTokenV1alpha1API.ListRememberMeToken(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RememberMeTokenV1alpha1API.ListRememberMeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRememberMeToken`: RememberMeTokenList
	fmt.Fprintf(os.Stdout, "Response from `RememberMeTokenV1alpha1API.ListRememberMeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRememberMeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**RememberMeTokenList**](RememberMeTokenList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRememberMeToken

> RememberMeToken PatchRememberMeToken(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of remembermetoken
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RememberMeTokenV1alpha1API.PatchRememberMeToken(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RememberMeTokenV1alpha1API.PatchRememberMeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRememberMeToken`: RememberMeToken
	fmt.Fprintf(os.Stdout, "Response from `RememberMeTokenV1alpha1API.PatchRememberMeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of remembermetoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRememberMeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**RememberMeToken**](RememberMeToken.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRememberMeToken

> RememberMeToken UpdateRememberMeToken(ctx, name).RememberMeToken(rememberMeToken).Execute()





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
	name := "name_example" // string | Name of remembermetoken
	rememberMeToken := *openapiclient.NewRememberMeToken("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewRememberMeTokenSpec("Series_example", "TokenValue_example", "Username_example")) // RememberMeToken | Updated remembermetoken (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RememberMeTokenV1alpha1API.UpdateRememberMeToken(context.Background(), name).RememberMeToken(rememberMeToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RememberMeTokenV1alpha1API.UpdateRememberMeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRememberMeToken`: RememberMeToken
	fmt.Fprintf(os.Stdout, "Response from `RememberMeTokenV1alpha1API.UpdateRememberMeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of remembermetoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRememberMeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rememberMeToken** | [**RememberMeToken**](RememberMeToken.md) | Updated remembermetoken | 

### Return type

[**RememberMeToken**](RememberMeToken.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

