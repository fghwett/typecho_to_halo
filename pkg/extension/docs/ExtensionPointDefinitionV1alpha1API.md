# \ExtensionPointDefinitionV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtensionPointDefinition**](ExtensionPointDefinitionV1alpha1API.md#CreateExtensionPointDefinition) | **Post** /apis/plugin.halo.run/v1alpha1/extensionpointdefinitions | 
[**DeleteExtensionPointDefinition**](ExtensionPointDefinitionV1alpha1API.md#DeleteExtensionPointDefinition) | **Delete** /apis/plugin.halo.run/v1alpha1/extensionpointdefinitions/{name} | 
[**GetExtensionPointDefinition**](ExtensionPointDefinitionV1alpha1API.md#GetExtensionPointDefinition) | **Get** /apis/plugin.halo.run/v1alpha1/extensionpointdefinitions/{name} | 
[**ListExtensionPointDefinition**](ExtensionPointDefinitionV1alpha1API.md#ListExtensionPointDefinition) | **Get** /apis/plugin.halo.run/v1alpha1/extensionpointdefinitions | 
[**PatchExtensionPointDefinition**](ExtensionPointDefinitionV1alpha1API.md#PatchExtensionPointDefinition) | **Patch** /apis/plugin.halo.run/v1alpha1/extensionpointdefinitions/{name} | 
[**UpdateExtensionPointDefinition**](ExtensionPointDefinitionV1alpha1API.md#UpdateExtensionPointDefinition) | **Put** /apis/plugin.halo.run/v1alpha1/extensionpointdefinitions/{name} | 



## CreateExtensionPointDefinition

> ExtensionPointDefinition CreateExtensionPointDefinition(ctx).ExtensionPointDefinition(extensionPointDefinition).Execute()





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
	extensionPointDefinition := *openapiclient.NewExtensionPointDefinition("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewExtensionPointSpec("ClassName_example", "DisplayName_example", "Type_example")) // ExtensionPointDefinition | Fresh extensionpointdefinition (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionPointDefinitionV1alpha1API.CreateExtensionPointDefinition(context.Background()).ExtensionPointDefinition(extensionPointDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionPointDefinitionV1alpha1API.CreateExtensionPointDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExtensionPointDefinition`: ExtensionPointDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionPointDefinitionV1alpha1API.CreateExtensionPointDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExtensionPointDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extensionPointDefinition** | [**ExtensionPointDefinition**](ExtensionPointDefinition.md) | Fresh extensionpointdefinition | 

### Return type

[**ExtensionPointDefinition**](ExtensionPointDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtensionPointDefinition

> DeleteExtensionPointDefinition(ctx, name).Execute()





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
	name := "name_example" // string | Name of extensionpointdefinition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExtensionPointDefinitionV1alpha1API.DeleteExtensionPointDefinition(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionPointDefinitionV1alpha1API.DeleteExtensionPointDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensionpointdefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionPointDefinitionRequest struct via the builder pattern


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


## GetExtensionPointDefinition

> ExtensionPointDefinition GetExtensionPointDefinition(ctx, name).Execute()





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
	name := "name_example" // string | Name of extensionpointdefinition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionPointDefinitionV1alpha1API.GetExtensionPointDefinition(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionPointDefinitionV1alpha1API.GetExtensionPointDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensionPointDefinition`: ExtensionPointDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionPointDefinitionV1alpha1API.GetExtensionPointDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensionpointdefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionPointDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionPointDefinition**](ExtensionPointDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensionPointDefinition

> ExtensionPointDefinitionList ListExtensionPointDefinition(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.ExtensionPointDefinitionV1alpha1API.ListExtensionPointDefinition(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionPointDefinitionV1alpha1API.ListExtensionPointDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExtensionPointDefinition`: ExtensionPointDefinitionList
	fmt.Fprintf(os.Stdout, "Response from `ExtensionPointDefinitionV1alpha1API.ListExtensionPointDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionPointDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ExtensionPointDefinitionList**](ExtensionPointDefinitionList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExtensionPointDefinition

> ExtensionPointDefinition PatchExtensionPointDefinition(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of extensionpointdefinition
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionPointDefinitionV1alpha1API.PatchExtensionPointDefinition(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionPointDefinitionV1alpha1API.PatchExtensionPointDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExtensionPointDefinition`: ExtensionPointDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionPointDefinitionV1alpha1API.PatchExtensionPointDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensionpointdefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExtensionPointDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**ExtensionPointDefinition**](ExtensionPointDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtensionPointDefinition

> ExtensionPointDefinition UpdateExtensionPointDefinition(ctx, name).ExtensionPointDefinition(extensionPointDefinition).Execute()





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
	name := "name_example" // string | Name of extensionpointdefinition
	extensionPointDefinition := *openapiclient.NewExtensionPointDefinition("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewExtensionPointSpec("ClassName_example", "DisplayName_example", "Type_example")) // ExtensionPointDefinition | Updated extensionpointdefinition (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExtensionPointDefinitionV1alpha1API.UpdateExtensionPointDefinition(context.Background(), name).ExtensionPointDefinition(extensionPointDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensionPointDefinitionV1alpha1API.UpdateExtensionPointDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtensionPointDefinition`: ExtensionPointDefinition
	fmt.Fprintf(os.Stdout, "Response from `ExtensionPointDefinitionV1alpha1API.UpdateExtensionPointDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of extensionpointdefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionPointDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionPointDefinition** | [**ExtensionPointDefinition**](ExtensionPointDefinition.md) | Updated extensionpointdefinition | 

### Return type

[**ExtensionPointDefinition**](ExtensionPointDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

