# \NotificationV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSpecifiedNotification**](NotificationV1alpha1UcAPI.md#DeleteSpecifiedNotification) | **Delete** /apis/api.notification.halo.run/v1alpha1/userspaces/{username}/notifications/{name} | 
[**ListUserNotificationPreferences**](NotificationV1alpha1UcAPI.md#ListUserNotificationPreferences) | **Get** /apis/api.notification.halo.run/v1alpha1/userspaces/{username}/notification-preferences | 
[**ListUserNotifications**](NotificationV1alpha1UcAPI.md#ListUserNotifications) | **Get** /apis/api.notification.halo.run/v1alpha1/userspaces/{username}/notifications | 
[**MarkNotificationAsRead**](NotificationV1alpha1UcAPI.md#MarkNotificationAsRead) | **Put** /apis/api.notification.halo.run/v1alpha1/userspaces/{username}/notifications/{name}/mark-as-read | 
[**MarkNotificationsAsRead**](NotificationV1alpha1UcAPI.md#MarkNotificationsAsRead) | **Put** /apis/api.notification.halo.run/v1alpha1/userspaces/{username}/notifications/-/mark-specified-as-read | 
[**SaveUserNotificationPreferences**](NotificationV1alpha1UcAPI.md#SaveUserNotificationPreferences) | **Post** /apis/api.notification.halo.run/v1alpha1/userspaces/{username}/notification-preferences | 



## DeleteSpecifiedNotification

> Notification DeleteSpecifiedNotification(ctx, username, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	username := "username_example" // string | Username
	name := "name_example" // string | Notification name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationV1alpha1UcAPI.DeleteSpecifiedNotification(context.Background(), username, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationV1alpha1UcAPI.DeleteSpecifiedNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSpecifiedNotification`: Notification
	fmt.Fprintf(os.Stdout, "Response from `NotificationV1alpha1UcAPI.DeleteSpecifiedNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 
**name** | **string** | Notification name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpecifiedNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Notification**](Notification.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserNotificationPreferences

> ReasonTypeNotifierMatrix ListUserNotificationPreferences(ctx, username).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	username := "username_example" // string | Username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationV1alpha1UcAPI.ListUserNotificationPreferences(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationV1alpha1UcAPI.ListUserNotificationPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserNotificationPreferences`: ReasonTypeNotifierMatrix
	fmt.Fprintf(os.Stdout, "Response from `NotificationV1alpha1UcAPI.ListUserNotificationPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserNotificationPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReasonTypeNotifierMatrix**](ReasonTypeNotifierMatrix.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserNotifications

> NotificationList ListUserNotifications(ctx, username).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	username := "username_example" // string | Username
	page := int32(56) // int32 | Page number. Default is 0. (optional)
	size := int32(56) // int32 | Size number. Default is 0. (optional)
	labelSelector := []string{"Inner_example"} // []string | Label selector. e.g.: hidden!=true (optional)
	fieldSelector := []string{"Inner_example"} // []string | Field selector. e.g.: metadata.name==halo (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationV1alpha1UcAPI.ListUserNotifications(context.Background(), username).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationV1alpha1UcAPI.ListUserNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserNotifications`: NotificationList
	fmt.Fprintf(os.Stdout, "Response from `NotificationV1alpha1UcAPI.ListUserNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**NotificationList**](NotificationList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationAsRead

> Notification MarkNotificationAsRead(ctx, username, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	username := "username_example" // string | Username
	name := "name_example" // string | Notification name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationV1alpha1UcAPI.MarkNotificationAsRead(context.Background(), username, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationV1alpha1UcAPI.MarkNotificationAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNotificationAsRead`: Notification
	fmt.Fprintf(os.Stdout, "Response from `NotificationV1alpha1UcAPI.MarkNotificationAsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 
**name** | **string** | Notification name | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Notification**](Notification.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationsAsRead

> []string MarkNotificationsAsRead(ctx, username).MarkSpecifiedRequest(markSpecifiedRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	username := "username_example" // string | Username
	markSpecifiedRequest := *openapiclient.NewMarkSpecifiedRequest() // MarkSpecifiedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationV1alpha1UcAPI.MarkNotificationsAsRead(context.Background(), username).MarkSpecifiedRequest(markSpecifiedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationV1alpha1UcAPI.MarkNotificationsAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNotificationsAsRead`: []string
	fmt.Fprintf(os.Stdout, "Response from `NotificationV1alpha1UcAPI.MarkNotificationsAsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationsAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **markSpecifiedRequest** | [**MarkSpecifiedRequest**](MarkSpecifiedRequest.md) |  | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveUserNotificationPreferences

> ReasonTypeNotifierMatrix SaveUserNotificationPreferences(ctx, username).ReasonTypeNotifierCollectionRequest(reasonTypeNotifierCollectionRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-user"
)

func main() {
	username := "username_example" // string | Username
	reasonTypeNotifierCollectionRequest := *openapiclient.NewReasonTypeNotifierCollectionRequest([]openapiclient.ReasonTypeNotifierRequest{*openapiclient.NewReasonTypeNotifierRequest()}) // ReasonTypeNotifierCollectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationV1alpha1UcAPI.SaveUserNotificationPreferences(context.Background(), username).ReasonTypeNotifierCollectionRequest(reasonTypeNotifierCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationV1alpha1UcAPI.SaveUserNotificationPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveUserNotificationPreferences`: ReasonTypeNotifierMatrix
	fmt.Fprintf(os.Stdout, "Response from `NotificationV1alpha1UcAPI.SaveUserNotificationPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveUserNotificationPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reasonTypeNotifierCollectionRequest** | [**ReasonTypeNotifierCollectionRequest**](ReasonTypeNotifierCollectionRequest.md) |  | 

### Return type

[**ReasonTypeNotifierMatrix**](ReasonTypeNotifierMatrix.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

