# \IndexV1alpha1PublicAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndicesSearch**](IndexV1alpha1PublicAPI.md#IndicesSearch) | **Post** /apis/api.halo.run/v1alpha1/indices/-/search | 
[**SearchPost**](IndexV1alpha1PublicAPI.md#SearchPost) | **Get** /apis/api.halo.run/v1alpha1/indices/post | 



## IndicesSearch

> SearchResult IndicesSearch(ctx).SearchOption(searchOption).Execute()





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
	searchOption := *openapiclient.NewSearchOption("Keyword_example") // SearchOption | Please note that the \"filterPublished\", \"filterExposed\" and \"filterRecycled\" fields are ignored in this endpoint. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexV1alpha1PublicAPI.IndicesSearch(context.Background()).SearchOption(searchOption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexV1alpha1PublicAPI.IndicesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IndicesSearch`: SearchResult
	fmt.Fprintf(os.Stdout, "Response from `IndexV1alpha1PublicAPI.IndicesSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndicesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchOption** | [**SearchOption**](SearchOption.md) | Please note that the \&quot;filterPublished\&quot;, \&quot;filterExposed\&quot; and \&quot;filterRecycled\&quot; fields are ignored in this endpoint. | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPost

> SearchResult SearchPost(ctx).Keyword(keyword).Limit(limit).HighlightPreTag(highlightPreTag).HighlightPostTag(highlightPostTag).Execute()





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
	keyword := "keyword_example" // string | Keyword to search
	limit := int32(56) // int32 | Limit of search results (optional)
	highlightPreTag := "highlightPreTag_example" // string | Highlight pre tag (optional)
	highlightPostTag := "highlightPostTag_example" // string | Highlight post tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexV1alpha1PublicAPI.SearchPost(context.Background()).Keyword(keyword).Limit(limit).HighlightPreTag(highlightPreTag).HighlightPostTag(highlightPostTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexV1alpha1PublicAPI.SearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPost`: SearchResult
	fmt.Fprintf(os.Stdout, "Response from `IndexV1alpha1PublicAPI.SearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyword** | **string** | Keyword to search | 
 **limit** | **int32** | Limit of search results | 
 **highlightPreTag** | **string** | Highlight pre tag | 
 **highlightPostTag** | **string** | Highlight post tag | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

