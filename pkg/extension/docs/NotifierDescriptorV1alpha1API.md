# \NotifierDescriptorV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotifierDescriptor**](NotifierDescriptorV1alpha1API.md#CreateNotifierDescriptor) | **Post** /apis/notification.halo.run/v1alpha1/notifierDescriptors | 
[**DeleteNotifierDescriptor**](NotifierDescriptorV1alpha1API.md#DeleteNotifierDescriptor) | **Delete** /apis/notification.halo.run/v1alpha1/notifierDescriptors/{name} | 
[**GetNotifierDescriptor**](NotifierDescriptorV1alpha1API.md#GetNotifierDescriptor) | **Get** /apis/notification.halo.run/v1alpha1/notifierDescriptors/{name} | 
[**ListNotifierDescriptor**](NotifierDescriptorV1alpha1API.md#ListNotifierDescriptor) | **Get** /apis/notification.halo.run/v1alpha1/notifierDescriptors | 
[**PatchNotifierDescriptor**](NotifierDescriptorV1alpha1API.md#PatchNotifierDescriptor) | **Patch** /apis/notification.halo.run/v1alpha1/notifierDescriptors/{name} | 
[**UpdateNotifierDescriptor**](NotifierDescriptorV1alpha1API.md#UpdateNotifierDescriptor) | **Put** /apis/notification.halo.run/v1alpha1/notifierDescriptors/{name} | 



## CreateNotifierDescriptor

> NotifierDescriptor CreateNotifierDescriptor(ctx).NotifierDescriptor(notifierDescriptor).Execute()





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
	notifierDescriptor := *openapiclient.NewNotifierDescriptor("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // NotifierDescriptor | Fresh notifierDescriptor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifierDescriptorV1alpha1API.CreateNotifierDescriptor(context.Background()).NotifierDescriptor(notifierDescriptor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierDescriptorV1alpha1API.CreateNotifierDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotifierDescriptor`: NotifierDescriptor
	fmt.Fprintf(os.Stdout, "Response from `NotifierDescriptorV1alpha1API.CreateNotifierDescriptor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotifierDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifierDescriptor** | [**NotifierDescriptor**](NotifierDescriptor.md) | Fresh notifierDescriptor | 

### Return type

[**NotifierDescriptor**](NotifierDescriptor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotifierDescriptor

> DeleteNotifierDescriptor(ctx, name).Execute()





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
	name := "name_example" // string | Name of notifierDescriptor

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotifierDescriptorV1alpha1API.DeleteNotifierDescriptor(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierDescriptorV1alpha1API.DeleteNotifierDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notifierDescriptor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotifierDescriptorRequest struct via the builder pattern


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


## GetNotifierDescriptor

> NotifierDescriptor GetNotifierDescriptor(ctx, name).Execute()





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
	name := "name_example" // string | Name of notifierDescriptor

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifierDescriptorV1alpha1API.GetNotifierDescriptor(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierDescriptorV1alpha1API.GetNotifierDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifierDescriptor`: NotifierDescriptor
	fmt.Fprintf(os.Stdout, "Response from `NotifierDescriptorV1alpha1API.GetNotifierDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notifierDescriptor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotifierDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotifierDescriptor**](NotifierDescriptor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotifierDescriptor

> NotifierDescriptorList ListNotifierDescriptor(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.NotifierDescriptorV1alpha1API.ListNotifierDescriptor(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierDescriptorV1alpha1API.ListNotifierDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotifierDescriptor`: NotifierDescriptorList
	fmt.Fprintf(os.Stdout, "Response from `NotifierDescriptorV1alpha1API.ListNotifierDescriptor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotifierDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**NotifierDescriptorList**](NotifierDescriptorList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNotifierDescriptor

> NotifierDescriptor PatchNotifierDescriptor(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of notifierDescriptor
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifierDescriptorV1alpha1API.PatchNotifierDescriptor(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierDescriptorV1alpha1API.PatchNotifierDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNotifierDescriptor`: NotifierDescriptor
	fmt.Fprintf(os.Stdout, "Response from `NotifierDescriptorV1alpha1API.PatchNotifierDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notifierDescriptor | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNotifierDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**NotifierDescriptor**](NotifierDescriptor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotifierDescriptor

> NotifierDescriptor UpdateNotifierDescriptor(ctx, name).NotifierDescriptor(notifierDescriptor).Execute()





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
	name := "name_example" // string | Name of notifierDescriptor
	notifierDescriptor := *openapiclient.NewNotifierDescriptor("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // NotifierDescriptor | Updated notifierDescriptor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifierDescriptorV1alpha1API.UpdateNotifierDescriptor(context.Background(), name).NotifierDescriptor(notifierDescriptor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierDescriptorV1alpha1API.UpdateNotifierDescriptor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotifierDescriptor`: NotifierDescriptor
	fmt.Fprintf(os.Stdout, "Response from `NotifierDescriptorV1alpha1API.UpdateNotifierDescriptor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notifierDescriptor | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotifierDescriptorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notifierDescriptor** | [**NotifierDescriptor**](NotifierDescriptor.md) | Updated notifierDescriptor | 

### Return type

[**NotifierDescriptor**](NotifierDescriptor.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

