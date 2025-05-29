# \PersonalAccessTokenV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePersonalAccessToken**](PersonalAccessTokenV1alpha1API.md#CreatePersonalAccessToken) | **Post** /apis/security.halo.run/v1alpha1/personalaccesstokens | 
[**DeletePersonalAccessToken**](PersonalAccessTokenV1alpha1API.md#DeletePersonalAccessToken) | **Delete** /apis/security.halo.run/v1alpha1/personalaccesstokens/{name} | 
[**GetPersonalAccessToken**](PersonalAccessTokenV1alpha1API.md#GetPersonalAccessToken) | **Get** /apis/security.halo.run/v1alpha1/personalaccesstokens/{name} | 
[**ListPersonalAccessToken**](PersonalAccessTokenV1alpha1API.md#ListPersonalAccessToken) | **Get** /apis/security.halo.run/v1alpha1/personalaccesstokens | 
[**PatchPersonalAccessToken**](PersonalAccessTokenV1alpha1API.md#PatchPersonalAccessToken) | **Patch** /apis/security.halo.run/v1alpha1/personalaccesstokens/{name} | 
[**UpdatePersonalAccessToken**](PersonalAccessTokenV1alpha1API.md#UpdatePersonalAccessToken) | **Put** /apis/security.halo.run/v1alpha1/personalaccesstokens/{name} | 



## CreatePersonalAccessToken

> PersonalAccessToken CreatePersonalAccessToken(ctx).PersonalAccessToken(personalAccessToken).Execute()





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
	personalAccessToken := *openapiclient.NewPersonalAccessToken("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // PersonalAccessToken | Fresh personalaccesstoken (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenV1alpha1API.CreatePersonalAccessToken(context.Background()).PersonalAccessToken(personalAccessToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1API.CreatePersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePersonalAccessToken`: PersonalAccessToken
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenV1alpha1API.CreatePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personalAccessToken** | [**PersonalAccessToken**](PersonalAccessToken.md) | Fresh personalaccesstoken | 

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


## DeletePersonalAccessToken

> DeletePersonalAccessToken(ctx, name).Execute()





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
	name := "name_example" // string | Name of personalaccesstoken

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonalAccessTokenV1alpha1API.DeletePersonalAccessToken(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1API.DeletePersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of personalaccesstoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonalAccessTokenRequest struct via the builder pattern


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


## GetPersonalAccessToken

> PersonalAccessToken GetPersonalAccessToken(ctx, name).Execute()





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
	name := "name_example" // string | Name of personalaccesstoken

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenV1alpha1API.GetPersonalAccessToken(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1API.GetPersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonalAccessToken`: PersonalAccessToken
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenV1alpha1API.GetPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of personalaccesstoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListPersonalAccessToken

> PersonalAccessTokenList ListPersonalAccessToken(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.PersonalAccessTokenV1alpha1API.ListPersonalAccessToken(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1API.ListPersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPersonalAccessToken`: PersonalAccessTokenList
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenV1alpha1API.ListPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PersonalAccessTokenList**](PersonalAccessTokenList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPersonalAccessToken

> PersonalAccessToken PatchPersonalAccessToken(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of personalaccesstoken
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenV1alpha1API.PatchPersonalAccessToken(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1API.PatchPersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPersonalAccessToken`: PersonalAccessToken
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenV1alpha1API.PatchPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of personalaccesstoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**PersonalAccessToken**](PersonalAccessToken.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonalAccessToken

> PersonalAccessToken UpdatePersonalAccessToken(ctx, name).PersonalAccessToken(personalAccessToken).Execute()





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
	name := "name_example" // string | Name of personalaccesstoken
	personalAccessToken := *openapiclient.NewPersonalAccessToken("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // PersonalAccessToken | Updated personalaccesstoken (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenV1alpha1API.UpdatePersonalAccessToken(context.Background(), name).PersonalAccessToken(personalAccessToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenV1alpha1API.UpdatePersonalAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePersonalAccessToken`: PersonalAccessToken
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenV1alpha1API.UpdatePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of personalaccesstoken | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **personalAccessToken** | [**PersonalAccessToken**](PersonalAccessToken.md) | Updated personalaccesstoken | 

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

