# \SystemConfigV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemConfigByGroup**](SystemConfigV1alpha1ConsoleAPI.md#GetSystemConfigByGroup) | **Get** /apis/console.api.halo.run/v1alpha1/systemconfigs/{group} | 
[**UpdateSystemConfigByGroup**](SystemConfigV1alpha1ConsoleAPI.md#UpdateSystemConfigByGroup) | **Put** /apis/console.api.halo.run/v1alpha1/systemconfigs/{group} | 



## GetSystemConfigByGroup

> map[string]interface{} GetSystemConfigByGroup(ctx, group).Execute()





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
	group := "group_example" // string | Group of the system config

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemConfigV1alpha1ConsoleAPI.GetSystemConfigByGroup(context.Background(), group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigV1alpha1ConsoleAPI.GetSystemConfigByGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemConfigByGroup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SystemConfigV1alpha1ConsoleAPI.GetSystemConfigByGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**group** | **string** | Group of the system config | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemConfigByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSystemConfigByGroup

> UpdateSystemConfigByGroup(ctx, group).Body(body).Execute()





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
	group := "group_example" // string | Group of the system config
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemConfigV1alpha1ConsoleAPI.UpdateSystemConfigByGroup(context.Background(), group).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemConfigV1alpha1ConsoleAPI.UpdateSystemConfigByGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**group** | **string** | Group of the system config | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemConfigByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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

