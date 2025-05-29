# \PolicyTemplateV1alpha1API

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyTemplate**](PolicyTemplateV1alpha1API.md#CreatePolicyTemplate) | **Post** /apis/storage.halo.run/v1alpha1/policytemplates | 
[**DeletePolicyTemplate**](PolicyTemplateV1alpha1API.md#DeletePolicyTemplate) | **Delete** /apis/storage.halo.run/v1alpha1/policytemplates/{name} | 
[**GetPolicyTemplate**](PolicyTemplateV1alpha1API.md#GetPolicyTemplate) | **Get** /apis/storage.halo.run/v1alpha1/policytemplates/{name} | 
[**ListPolicyTemplate**](PolicyTemplateV1alpha1API.md#ListPolicyTemplate) | **Get** /apis/storage.halo.run/v1alpha1/policytemplates | 
[**PatchPolicyTemplate**](PolicyTemplateV1alpha1API.md#PatchPolicyTemplate) | **Patch** /apis/storage.halo.run/v1alpha1/policytemplates/{name} | 
[**UpdatePolicyTemplate**](PolicyTemplateV1alpha1API.md#UpdatePolicyTemplate) | **Put** /apis/storage.halo.run/v1alpha1/policytemplates/{name} | 



## CreatePolicyTemplate

> PolicyTemplate CreatePolicyTemplate(ctx).PolicyTemplate(policyTemplate).Execute()





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
	policyTemplate := *openapiclient.NewPolicyTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // PolicyTemplate | Fresh policytemplate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyTemplateV1alpha1API.CreatePolicyTemplate(context.Background()).PolicyTemplate(policyTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyTemplateV1alpha1API.CreatePolicyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicyTemplate`: PolicyTemplate
	fmt.Fprintf(os.Stdout, "Response from `PolicyTemplateV1alpha1API.CreatePolicyTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyTemplate** | [**PolicyTemplate**](PolicyTemplate.md) | Fresh policytemplate | 

### Return type

[**PolicyTemplate**](PolicyTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyTemplate

> DeletePolicyTemplate(ctx, name).Execute()





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
	name := "name_example" // string | Name of policytemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyTemplateV1alpha1API.DeletePolicyTemplate(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyTemplateV1alpha1API.DeletePolicyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of policytemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyTemplateRequest struct via the builder pattern


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


## GetPolicyTemplate

> PolicyTemplate GetPolicyTemplate(ctx, name).Execute()





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
	name := "name_example" // string | Name of policytemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyTemplateV1alpha1API.GetPolicyTemplate(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyTemplateV1alpha1API.GetPolicyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyTemplate`: PolicyTemplate
	fmt.Fprintf(os.Stdout, "Response from `PolicyTemplateV1alpha1API.GetPolicyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of policytemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyTemplate**](PolicyTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyTemplate

> PolicyTemplateList ListPolicyTemplate(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()





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
	resp, r, err := apiClient.PolicyTemplateV1alpha1API.ListPolicyTemplate(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyTemplateV1alpha1API.ListPolicyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicyTemplate`: PolicyTemplateList
	fmt.Fprintf(os.Stdout, "Response from `PolicyTemplateV1alpha1API.ListPolicyTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**PolicyTemplateList**](PolicyTemplateList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPolicyTemplate

> PolicyTemplate PatchPolicyTemplate(ctx, name).JsonPatchInner(jsonPatchInner).Execute()





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
	name := "name_example" // string | Name of policytemplate
	jsonPatchInner := []openapiclient.JsonPatchInner{openapiclient.JsonPatch_inner{AddOperation: openapiclient.NewAddOperation("Op_example", "/a/b/c", interface{}(123))}} // []JsonPatchInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyTemplateV1alpha1API.PatchPolicyTemplate(context.Background(), name).JsonPatchInner(jsonPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyTemplateV1alpha1API.PatchPolicyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPolicyTemplate`: PolicyTemplate
	fmt.Fprintf(os.Stdout, "Response from `PolicyTemplateV1alpha1API.PatchPolicyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of policytemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchInner** | [**[]JsonPatchInner**](JsonPatchInner.md) |  | 

### Return type

[**PolicyTemplate**](PolicyTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyTemplate

> PolicyTemplate UpdatePolicyTemplate(ctx, name).PolicyTemplate(policyTemplate).Execute()





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
	name := "name_example" // string | Name of policytemplate
	policyTemplate := *openapiclient.NewPolicyTemplate("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // PolicyTemplate | Updated policytemplate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyTemplateV1alpha1API.UpdatePolicyTemplate(context.Background(), name).PolicyTemplate(policyTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyTemplateV1alpha1API.UpdatePolicyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicyTemplate`: PolicyTemplate
	fmt.Fprintf(os.Stdout, "Response from `PolicyTemplateV1alpha1API.UpdatePolicyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of policytemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyTemplate** | [**PolicyTemplate**](PolicyTemplate.md) | Updated policytemplate | 

### Return type

[**PolicyTemplate**](PolicyTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

