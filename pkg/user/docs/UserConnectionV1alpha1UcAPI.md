# \UserConnectionV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisconnectMyConnection**](UserConnectionV1alpha1UcAPI.md#DisconnectMyConnection) | **Put** /apis/uc.api.auth.halo.run/v1alpha1/user-connections/{registerId}/disconnect | 



## DisconnectMyConnection

> []UserConnection DisconnectMyConnection(ctx, registerId).Execute()





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
	registerId := "registerId_example" // string | The registration ID of the third-party platform.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserConnectionV1alpha1UcAPI.DisconnectMyConnection(context.Background(), registerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionV1alpha1UcAPI.DisconnectMyConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectMyConnection`: []UserConnection
	fmt.Fprintf(os.Stdout, "Response from `UserConnectionV1alpha1UcAPI.DisconnectMyConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registerId** | **string** | The registration ID of the third-party platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectMyConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserConnection**](UserConnection.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

