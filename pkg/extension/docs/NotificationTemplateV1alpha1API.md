# \NotificationTemplateV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationTemplate**](NotificationTemplateV1alpha1API.md#CreateNotificationTemplate) | **Post** /apis/notification.halo.run/v1alpha1/notificationtemplates | 
[**DeleteNotificationTemplate**](NotificationTemplateV1alpha1API.md#DeleteNotificationTemplate) | **Delete** /apis/notification.halo.run/v1alpha1/notificationtemplates/{name} | 
[**GetNotificationTemplate**](NotificationTemplateV1alpha1API.md#GetNotificationTemplate) | **Get** /apis/notification.halo.run/v1alpha1/notificationtemplates/{name} | 
[**ListNotificationTemplate**](NotificationTemplateV1alpha1API.md#ListNotificationTemplate) | **Get** /apis/notification.halo.run/v1alpha1/notificationtemplates | 
[**PatchNotificationTemplate**](NotificationTemplateV1alpha1API.md#PatchNotificationTemplate) | **Patch** /apis/notification.halo.run/v1alpha1/notificationtemplates/{name} | 
[**UpdateNotificationTemplate**](NotificationTemplateV1alpha1API.md#UpdateNotificationTemplate) | **Put** /apis/notification.halo.run/v1alpha1/notificationtemplates/{name} | 



## CreateNotificationTemplate

> NotificationTemplate CreateNotificationTemplate(ctx).NotificationTemplate(notificationTemplate).Execute()





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
	notificationTemplate := *openapiclient.NewNotificationTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // NotificationTemplate | Fresh notificationtemplate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationTemplateV1alpha1API.CreateNotificationTemplate(context.Background()).NotificationTemplate(notificationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplateV1alpha1API.CreateNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationTemplate`: NotificationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NotificationTemplateV1alpha1API.CreateNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationTemplate** | [**NotificationTemplate**](NotificationTemplate.md) | Fresh notificationtemplate | 

### Return type

[**NotificationTemplate**](NotificationTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationTemplate

> DeleteNotificationTemplate(ctx, name).Execute()





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
	name := "name_example" // string | Name of notificationtemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationTemplateV1alpha1API.DeleteNotificationTemplate(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplateV1alpha1API.DeleteNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notificationtemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationTemplateRequest struct via the builder pattern


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


## GetNotificationTemplate

> NotificationTemplate GetNotificationTemplate(ctx, name).Execute()





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
	name := "name_example" // string | Name of notificationtemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationTemplateV1alpha1API.GetNotificationTemplate(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplateV1alpha1API.GetNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationTemplate`: NotificationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NotificationTemplateV1alpha1API.GetNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notificationtemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationTemplate**](NotificationTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationTemplate

> NotificationTemplateList ListNotificationTemplate(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.NotificationTemplateV1alpha1API.ListNotificationTemplate(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplateV1alpha1API.ListNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotificationTemplate`: NotificationTemplateList
	fmt.Fprintf(os.Stdout, "Response from `NotificationTemplateV1alpha1API.ListNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**NotificationTemplateList**](NotificationTemplateList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNotificationTemplate

> NotificationTemplate PatchNotificationTemplate(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of notificationtemplate
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationTemplateV1alpha1API.PatchNotificationTemplate(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplateV1alpha1API.PatchNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNotificationTemplate`: NotificationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NotificationTemplateV1alpha1API.PatchNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notificationtemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**NotificationTemplate**](NotificationTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationTemplate

> NotificationTemplate UpdateNotificationTemplate(ctx, name).NotificationTemplate(notificationTemplate).Execute()





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
	name := "name_example" // string | Name of notificationtemplate
	notificationTemplate := *openapiclient.NewNotificationTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // NotificationTemplate | Updated notificationtemplate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationTemplateV1alpha1API.UpdateNotificationTemplate(context.Background(), name).NotificationTemplate(notificationTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationTemplateV1alpha1API.UpdateNotificationTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationTemplate`: NotificationTemplate
	fmt.Fprintf(os.Stdout, "Response from `NotificationTemplateV1alpha1API.UpdateNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of notificationtemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationTemplate** | [**NotificationTemplate**](NotificationTemplate.md) | Updated notificationtemplate | 

### Return type

[**NotificationTemplate**](NotificationTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

