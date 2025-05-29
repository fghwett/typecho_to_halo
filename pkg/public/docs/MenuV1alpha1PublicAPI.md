# \MenuV1alpha1PublicAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryMenuByName**](MenuV1alpha1PublicAPI.md#QueryMenuByName) | **Get** /apis/api.halo.run/v1alpha1/menus/{name} | 
[**QueryPrimaryMenu**](MenuV1alpha1PublicAPI.md#QueryPrimaryMenu) | **Get** /apis/api.halo.run/v1alpha1/menus/- | 



## QueryMenuByName

> MenuVo QueryMenuByName(ctx, name).Execute()





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
	name := "name_example" // string | Menu name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MenuV1alpha1PublicAPI.QueryMenuByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuV1alpha1PublicAPI.QueryMenuByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMenuByName`: MenuVo
	fmt.Fprintf(os.Stdout, "Response from `MenuV1alpha1PublicAPI.QueryMenuByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Menu name | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryMenuByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MenuVo**](MenuVo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPrimaryMenu

> MenuVo QueryPrimaryMenu(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MenuV1alpha1PublicAPI.QueryPrimaryMenu(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuV1alpha1PublicAPI.QueryPrimaryMenu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPrimaryMenu`: MenuVo
	fmt.Fprintf(os.Stdout, "Response from `MenuV1alpha1PublicAPI.QueryPrimaryMenu`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQueryPrimaryMenuRequest struct via the builder pattern


### Return type

[**MenuVo**](MenuVo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

