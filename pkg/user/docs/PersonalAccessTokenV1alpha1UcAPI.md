# \PersonalAccessTokenV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePat**](PersonalAccessTokenV1alpha1UcAPI.md#DeletePat) | **Delete** /apis/uc.api.security.halo.run/v1alpha1/personalaccesstokens/{name} | 
[**GeneratePat**](PersonalAccessTokenV1alpha1UcAPI.md#GeneratePat) | **Post** /apis/uc.api.security.halo.run/v1alpha1/personalaccesstokens | 
[**ObtainPat**](PersonalAccessTokenV1alpha1UcAPI.md#ObtainPat) | **Get** /apis/uc.api.security.halo.run/v1alpha1/personalaccesstokens/{name} | 
[**ObtainPats**](PersonalAccessTokenV1alpha1UcAPI.md#ObtainPats) | **Get** /apis/uc.api.security.halo.run/v1alpha1/personalaccesstokens | 
[**RestorePat**](PersonalAccessTokenV1alpha1UcAPI.md#RestorePat) | **Put** /apis/uc.api.security.halo.run/v1alpha1/personalaccesstokens/{name}/actions/restoration | 
[**RevokePat**](PersonalAccessTokenV1alpha1UcAPI.md#RevokePat) | **Put** /apis/uc.api.security.halo.run/v1alpha1/personalaccesstokens/{name}/actions/revocation | 



## DeletePat

> DeletePat(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonalAccessTokenV1alpha1UcAPI.DeletePat(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1UcAPI.DeletePat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePatRequest struct via the builder pattern


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


## GeneratePat

> PersonalAccessToken GeneratePat(ctx).PersonalAccessToken(personalAccessToken).Execute()





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
	personalAccessToken := *openapiclient.NewPersonalAccessToken("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // PersonalAccessToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenV1alpha1UcAPI.GeneratePat(context.Background()).PersonalAccessToken(personalAccessToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1UcAPI.GeneratePat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeneratePat`: PersonalAccessToken
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenV1alpha1UcAPI.GeneratePat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personalAccessToken** | [**PersonalAccessToken**](PersonalAccessToken.md) |  | 

### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObtainPat

> ObtainPat(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonalAccessTokenV1alpha1UcAPI.ObtainPat(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1UcAPI.ObtainPat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiObtainPatRequest struct via the builder pattern


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


## ObtainPats

> []PersonalAccessToken ObtainPats(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenV1alpha1UcAPI.ObtainPats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1UcAPI.ObtainPats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObtainPats`: []PersonalAccessToken
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenV1alpha1UcAPI.ObtainPats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiObtainPatsRequest struct via the builder pattern


### Return type

[**[]PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePat

> RestorePat(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonalAccessTokenV1alpha1UcAPI.RestorePat(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1UcAPI.RestorePat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePatRequest struct via the builder pattern


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


## RevokePat

> RevokePat(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonalAccessTokenV1alpha1UcAPI.RevokePat(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1UcAPI.RevokePat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePatRequest struct via the builder pattern


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

