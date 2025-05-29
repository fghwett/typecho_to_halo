# \TagV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPostTags**](TagV1alpha1ConsoleAPI.md#ListPostTags) | **Get** /apis/api.console.halo.run/v1alpha1/tags | 



## ListPostTags

> TagList ListPostTags(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Keyword(keyword).Execute()





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
	page := int32(56) // int32 | Page number. Default is 0. (optional)
	size := int32(56) // int32 | Size number. Default is 0. (optional)
	labelSelector := []string{"Inner_example"} // []string | Label selector. e.g.: hidden!=true (optional)
	fieldSelector := []string{"Inner_example"} // []string | Field selector. e.g.: metadata.name==halo (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	keyword := "keyword_example" // string | Post tags filtered by keyword. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagV1alpha1ConsoleAPI.ListPostTags(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Keyword(keyword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagV1alpha1ConsoleAPI.ListPostTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostTags`: TagList
	fmt.Fprintf(os.Stdout, "Response from `TagV1alpha1ConsoleAPI.ListPostTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **keyword** | **string** | Post tags filtered by keyword. | 

### Return type

[**TagList**](TagList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

