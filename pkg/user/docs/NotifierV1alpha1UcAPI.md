# \NotifierV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchReceiverConfig**](NotifierV1alpha1UcAPI.md#FetchReceiverConfig) | **Get** /apis/api.notification.halo.run/v1alpha1/notifiers/{name}/receiver-config | 
[**SaveReceiverConfig**](NotifierV1alpha1UcAPI.md#SaveReceiverConfig) | **Post** /apis/api.notification.halo.run/v1alpha1/notifiers/{name}/receiver-config | 



## FetchReceiverConfig

> map[string]interface{} FetchReceiverConfig(ctx, name).Execute()





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
	name := "name_example" // string | Notifier name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifierV1alpha1UcAPI.FetchReceiverConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierV1alpha1UcAPI.FetchReceiverConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchReceiverConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotifierV1alpha1UcAPI.FetchReceiverConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Notifier name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchReceiverConfigRequest struct via the builder pattern


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


## SaveReceiverConfig

> SaveReceiverConfig(ctx, name).Body(body).Execute()





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
	name := "name_example" // string | Notifier name
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotifierV1alpha1UcAPI.SaveReceiverConfig(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifierV1alpha1UcAPI.SaveReceiverConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSaveReceiverConfigRequest struct via the builder pattern


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

