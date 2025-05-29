# \PluginV1alpha1PublicAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryPluginAvailableByName**](PluginV1alpha1PublicAPI.md#QueryPluginAvailableByName) | **Get** /apis/api.plugin.halo.run/v1alpha1/plugins/{name}/available | 



## QueryPluginAvailableByName

> bool QueryPluginAvailableByName(ctx, name).Execute()





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
	name := "name_example" // string | Plugin name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1PublicAPI.QueryPluginAvailableByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1PublicAPI.QueryPluginAvailableByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPluginAvailableByName`: bool
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1PublicAPI.QueryPluginAvailableByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Plugin name | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryPluginAvailableByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

