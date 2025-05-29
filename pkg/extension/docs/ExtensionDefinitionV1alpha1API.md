# \ExtensionDefinitionV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtensionDefinition**](ExtensionDefinitionV1alpha1API.md#CreateExtensionDefinition) | **Post** /apis/plugin.halo.run/v1alpha1/extensiondefinitions | 
[**DeleteExtensionDefinition**](ExtensionDefinitionV1alpha1API.md#DeleteExtensionDefinition) | **Delete** /apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name} | 
[**GetExtensionDefinition**](ExtensionDefinitionV1alpha1API.md#GetExtensionDefinition) | **Get** /apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name} | 
[**ListExtensionDefinition**](ExtensionDefinitionV1alpha1API.md#ListExtensionDefinition) | **Get** /apis/plugin.halo.run/v1alpha1/extensiondefinitions | 
[**PatchExtensionDefinition**](ExtensionDefinitionV1alpha1API.md#PatchExtensionDefinition) | **Patch** /apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name} | 
[**UpdateExtensionDefinition**](ExtensionDefinitionV1alpha1API.md#UpdateExtensionDefinition) | **Put** /apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name} | 



## CreateExtensionDefinition

> ExtensionDefinition CreateExtensionDefinition(ctx).ExtensionDefinition(extensionDefinition).Execute()





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
	extensionDefinition := *openapiclient.NewExtensionDefinition("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewExtensionSpec("ClassName_example", "DisplayName_example", "ExtensionPointName_example")) // ExtensionDefinition | Fresh extensiondefinition (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionDefinitionV1alpha1API.CreateExtensionDefinition(context.Background()).ExtensionDefinition(extensionDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionDefinitionV1alpha1API.CreateExtensionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExtensionDefinition`: ExtensionDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionDefinitionV1alpha1API.CreateExtensionDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExtensionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extensionDefinition** | [**ExtensionDefinition**](ExtensionDefinition.md) | Fresh extensiondefinition | 

### Return type

[**ExtensionDefinition**](ExtensionDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtensionDefinition

> DeleteExtensionDefinition(ctx, name).Execute()





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
	name := "name_example" // string | Name of extensiondefinition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionDefinitionV1alpha1API.DeleteExtensionDefinition(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionDefinitionV1alpha1API.DeleteExtensionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensiondefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionDefinitionRequest struct via the builder pattern


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


## GetExtensionDefinition

> ExtensionDefinition GetExtensionDefinition(ctx, name).Execute()





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
	name := "name_example" // string | Name of extensiondefinition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionDefinitionV1alpha1API.GetExtensionDefinition(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionDefinitionV1alpha1API.GetExtensionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionDefinition`: ExtensionDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionDefinitionV1alpha1API.GetExtensionDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensiondefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionDefinition**](ExtensionDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensionDefinition

> ExtensionDefinitionList ListExtensionDefinition(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.ExtensionDefinitionV1alpha1API.ListExtensionDefinition(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionDefinitionV1alpha1API.ListExtensionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExtensionDefinition`: ExtensionDefinitionList
	fmt.Fprintf(os.Stdout, "Response from `ExtensionDefinitionV1alpha1API.ListExtensionDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ExtensionDefinitionList**](ExtensionDefinitionList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExtensionDefinition

> ExtensionDefinition PatchExtensionDefinition(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of extensiondefinition
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionDefinitionV1alpha1API.PatchExtensionDefinition(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionDefinitionV1alpha1API.PatchExtensionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExtensionDefinition`: ExtensionDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionDefinitionV1alpha1API.PatchExtensionDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensiondefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExtensionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**ExtensionDefinition**](ExtensionDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtensionDefinition

> ExtensionDefinition UpdateExtensionDefinition(ctx, name).ExtensionDefinition(extensionDefinition).Execute()





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
	name := "name_example" // string | Name of extensiondefinition
	extensionDefinition := *openapiclient.NewExtensionDefinition("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewExtensionSpec("ClassName_example", "DisplayName_example", "ExtensionPointName_example")) // ExtensionDefinition | Updated extensiondefinition (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionDefinitionV1alpha1API.UpdateExtensionDefinition(context.Background(), name).ExtensionDefinition(extensionDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionDefinitionV1alpha1API.UpdateExtensionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtensionDefinition`: ExtensionDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionDefinitionV1alpha1API.UpdateExtensionDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensiondefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionDefinition** | [**ExtensionDefinition**](ExtensionDefinition.md) | Updated extensiondefinition | 

### Return type

[**ExtensionDefinition**](ExtensionDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

