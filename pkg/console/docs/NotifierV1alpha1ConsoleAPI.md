# \NotifierV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSenderConfig**](NotifierV1alpha1ConsoleAPI.md#FetchSenderConfig) | **Get** /apis/api.console.halo.run/v1alpha1/notifiers/{name}/sender-config | 
[**SaveSenderConfig**](NotifierV1alpha1ConsoleAPI.md#SaveSenderConfig) | **Post** /apis/api.console.halo.run/v1alpha1/notifiers/{name}/sender-config | 
[**VerifyEmailSenderConfig**](NotifierV1alpha1ConsoleAPI.md#VerifyEmailSenderConfig) | **Post** /apis/console.api.notification.halo.run/v1alpha1/notifiers/default-email-notifier/verify-connection | 



## FetchSenderConfig

> map[string]interface{} FetchSenderConfig(ctx, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-console"
)

func main() {
	name := "name_example" // string | Notifier name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifierV1alpha1ConsoleAPI.FetchSenderConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierV1alpha1ConsoleAPI.FetchSenderConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSenderConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotifierV1alpha1ConsoleAPI.FetchSenderConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Notifier name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSenderConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveSenderConfig

> SaveSenderConfig(ctx, name).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-console"
)

func main() {
	name := "name_example" // string | Notifier name
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotifierV1alpha1ConsoleAPI.SaveSenderConfig(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierV1alpha1ConsoleAPI.SaveSenderConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Notifier name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveSenderConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmailSenderConfig

> VerifyEmailSenderConfig(ctx).EmailConfigValidationRequest(emailConfigValidationRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-console"
)

func main() {
	emailConfigValidationRequest := *openapiclient.NewEmailConfigValidationRequest() // EmailConfigValidationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotifierV1alpha1ConsoleAPI.VerifyEmailSenderConfig(context.Background()).EmailConfigValidationRequest(emailConfigValidationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierV1alpha1ConsoleAPI.VerifyEmailSenderConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEmailSenderConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailConfigValidationRequest** | [**EmailConfigValidationRequest**](EmailConfigValidationRequest.md) |  | 

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

