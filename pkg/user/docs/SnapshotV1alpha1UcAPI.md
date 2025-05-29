# \SnapshotV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnapshotForPost**](SnapshotV1alpha1UcAPI.md#GetSnapshotForPost) | **Get** /apis/uc.api.content.halo.run/v1alpha1/snapshots/{name} | 



## GetSnapshotForPost

> Snapshot GetSnapshotForPost(ctx, name).PostName(postName).Patched(patched).Execute()





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
	name := "name_example" // string | Snapshot name.
	postName := "postName_example" // string | Post name.
	patched := true // bool | Should include patched content and raw or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotV1alpha1UcAPI.GetSnapshotForPost(context.Background(), name).PostName(postName).Patched(patched).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotV1alpha1UcAPI.GetSnapshotForPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotForPost`: Snapshot
	fmt.Fprintf(os.Stdout, "Response from `SnapshotV1alpha1UcAPI.GetSnapshotForPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Snapshot name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotForPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postName** | **string** | Post name. | 
 **patched** | **bool** | Should include patched content and raw or not. | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

