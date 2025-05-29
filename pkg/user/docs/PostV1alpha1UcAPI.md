# \PostV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMyPost**](PostV1alpha1UcAPI.md#CreateMyPost) | **Post** /apis/uc.api.content.halo.run/v1alpha1/posts | 
[**GetMyPost**](PostV1alpha1UcAPI.md#GetMyPost) | **Get** /apis/uc.api.content.halo.run/v1alpha1/posts/{name} | 
[**GetMyPostDraft**](PostV1alpha1UcAPI.md#GetMyPostDraft) | **Get** /apis/uc.api.content.halo.run/v1alpha1/posts/{name}/draft | 
[**ListMyPosts**](PostV1alpha1UcAPI.md#ListMyPosts) | **Get** /apis/uc.api.content.halo.run/v1alpha1/posts | 
[**PublishMyPost**](PostV1alpha1UcAPI.md#PublishMyPost) | **Put** /apis/uc.api.content.halo.run/v1alpha1/posts/{name}/publish | 
[**RecycleMyPost**](PostV1alpha1UcAPI.md#RecycleMyPost) | **Delete** /apis/uc.api.content.halo.run/v1alpha1/posts/{name}/recycle | 
[**UnpublishMyPost**](PostV1alpha1UcAPI.md#UnpublishMyPost) | **Put** /apis/uc.api.content.halo.run/v1alpha1/posts/{name}/unpublish | 
[**UpdateMyPost**](PostV1alpha1UcAPI.md#UpdateMyPost) | **Put** /apis/uc.api.content.halo.run/v1alpha1/posts/{name} | 
[**UpdateMyPostDraft**](PostV1alpha1UcAPI.md#UpdateMyPostDraft) | **Put** /apis/uc.api.content.halo.run/v1alpha1/posts/{name}/draft | 



## CreateMyPost

> Post CreateMyPost(ctx).Post(post).Execute()





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
	post := *openapiclient.NewPost("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewPostSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example")) // Post |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.CreateMyPost(context.Background()).Post(post).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.CreateMyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMyPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.CreateMyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **post** | [**Post**](Post.md) |  | 

### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPost

> Post GetMyPost(ctx, name).Execute()





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
	name := "name_example" // string | Post name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.GetMyPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.GetMyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.GetMyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Post name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPostDraft

> Snapshot GetMyPostDraft(ctx, name).Patched(patched).Execute()





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
	name := "name_example" // string | Post name
	patched := true // bool | Should include patched content and raw or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.GetMyPostDraft(context.Background(), name).Patched(patched).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.GetMyPostDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPostDraft`: Snapshot
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.GetMyPostDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Post name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPostDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ListMyPosts

> ListedPostList ListMyPosts(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).PublishPhase(publishPhase).Keyword(keyword).CategoryWithChildren(categoryWithChildren).Execute()





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
	page := int32(56) // int32 | Page number. Default is 0. (optional)
	size := int32(56) // int32 | Size number. Default is 0. (optional)
	labelSelector := []string{"Inner_example"} // []string | Label selector. e.g.: hidden!=true (optional)
	fieldSelector := []string{"Inner_example"} // []string | Field selector. e.g.: metadata.name==halo (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	publishPhase := "publishPhase_example" // string | Posts filtered by publish phase. (optional)
	keyword := "keyword_example" // string | Posts filtered by keyword. (optional)
	categoryWithChildren := "categoryWithChildren_example" // string | Posts filtered by category including sub-categories. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.ListMyPosts(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).PublishPhase(publishPhase).Keyword(keyword).CategoryWithChildren(categoryWithChildren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.ListMyPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyPosts`: ListedPostList
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.ListMyPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **publishPhase** | **string** | Posts filtered by publish phase. | 
 **keyword** | **string** | Posts filtered by keyword. | 
 **categoryWithChildren** | **string** | Posts filtered by category including sub-categories. | 

### Return type

[**ListedPostList**](ListedPostList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishMyPost

> Post PublishMyPost(ctx, name).Execute()





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
	name := "name_example" // string | Post name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.PublishMyPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.PublishMyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishMyPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.PublishMyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Post name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishMyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecycleMyPost

> Post RecycleMyPost(ctx, name).Execute()





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
	name := "name_example" // string | Post name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.RecycleMyPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.RecycleMyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecycleMyPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.RecycleMyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Post name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecycleMyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishMyPost

> Post UnpublishMyPost(ctx, name).Execute()





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
	name := "name_example" // string | Post name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.UnpublishMyPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.UnpublishMyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpublishMyPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.UnpublishMyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Post name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishMyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyPost

> Post UpdateMyPost(ctx, name).Post(post).Execute()





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
	name := "name_example" // string | Post name
	post := *openapiclient.NewPost("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewPostSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example")) // Post |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.UpdateMyPost(context.Background(), name).Post(post).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.UpdateMyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMyPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.UpdateMyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Post name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **post** | [**Post**](Post.md) |  | 

### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyPostDraft

> Snapshot UpdateMyPostDraft(ctx, name).Snapshot(snapshot).Execute()





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
	name := "name_example" // string | Post name
	snapshot := *openapiclient.NewSnapshot("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewSnapShotSpec("Owner_example", "RawType_example", *openapiclient.NewRef("Name_example"))) // Snapshot |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1UcAPI.UpdateMyPostDraft(context.Background(), name).Snapshot(snapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1UcAPI.UpdateMyPostDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMyPostDraft`: Snapshot
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1UcAPI.UpdateMyPostDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Post name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyPostDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshot** | [**Snapshot**](Snapshot.md) |  | 

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

