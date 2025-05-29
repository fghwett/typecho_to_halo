# \NotificationV1alpha1PublicAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Unsubscribe**](NotificationV1alpha1PublicAPI.md#Unsubscribe) | **Get** /apis/api.notification.halo.run/v1alpha1/subscriptions/{name}/unsubscribe | 



## Unsubscribe

> string Unsubscribe(ctx, name).Token(token).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "api.halo.run/apis/openapi-go-public"
)

func main() {
	name := "name_example" // string | Subscription name
	token := "token_example" // string | Unsubscribe token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationV1alpha1PublicAPI.Unsubscribe(context.Background(), name).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationV1alpha1PublicAPI.Unsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Unsubscribe`: string
	fmt.Fprintf(os.Stdout, "Response from `NotificationV1alpha1PublicAPI.Unsubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Subscription name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** | Unsubscribe token | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

