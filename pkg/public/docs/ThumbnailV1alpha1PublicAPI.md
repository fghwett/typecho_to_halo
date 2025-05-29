# \ThumbnailV1alpha1PublicAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetThumbnailByUri**](ThumbnailV1alpha1PublicAPI.md#GetThumbnailByUri) | **Get** /apis/api.storage.halo.run/v1alpha1/thumbnails/-/via-uri | 



## GetThumbnailByUri

> *os.File GetThumbnailByUri(ctx).Uri(uri).Size(size).Execute()





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
	uri := "uri_example" // string | The URI of the image
	size := "size_example" // string | The size of the thumbnail,available values are s,m,l,xl

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThumbnailV1alpha1PublicAPI.GetThumbnailByUri(context.Background()).Uri(uri).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailV1alpha1PublicAPI.GetThumbnailByUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThumbnailByUri`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ThumbnailV1alpha1PublicAPI.GetThumbnailByUri`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailByUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uri** | **string** | The URI of the image | 
 **size** | **string** | The size of the thumbnail,available values are s,m,l,xl | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

