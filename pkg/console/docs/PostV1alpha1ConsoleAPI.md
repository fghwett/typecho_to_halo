# \PostV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePostContent**](PostV1alpha1ConsoleAPI.md#DeletePostContent) | **Delete** /apis/api.console.halo.run/v1alpha1/posts/{name}/content | 
[**DraftPost**](PostV1alpha1ConsoleAPI.md#DraftPost) | **Post** /apis/api.console.halo.run/v1alpha1/posts | 
[**FetchPostContent**](PostV1alpha1ConsoleAPI.md#FetchPostContent) | **Get** /apis/api.console.halo.run/v1alpha1/posts/{name}/content | 
[**FetchPostHeadContent**](PostV1alpha1ConsoleAPI.md#FetchPostHeadContent) | **Get** /apis/api.console.halo.run/v1alpha1/posts/{name}/head-content | 
[**FetchPostReleaseContent**](PostV1alpha1ConsoleAPI.md#FetchPostReleaseContent) | **Get** /apis/api.console.halo.run/v1alpha1/posts/{name}/release-content | 
[**ListPostSnapshots**](PostV1alpha1ConsoleAPI.md#ListPostSnapshots) | **Get** /apis/api.console.halo.run/v1alpha1/posts/{name}/snapshot | 
[**ListPosts**](PostV1alpha1ConsoleAPI.md#ListPosts) | **Get** /apis/api.console.halo.run/v1alpha1/posts | 
[**PublishPost**](PostV1alpha1ConsoleAPI.md#PublishPost) | **Put** /apis/api.console.halo.run/v1alpha1/posts/{name}/publish | 
[**RecyclePost**](PostV1alpha1ConsoleAPI.md#RecyclePost) | **Put** /apis/api.console.halo.run/v1alpha1/posts/{name}/recycle | 
[**RevertToSpecifiedSnapshotForPost**](PostV1alpha1ConsoleAPI.md#RevertToSpecifiedSnapshotForPost) | **Put** /apis/api.console.halo.run/v1alpha1/posts/{name}/revert-content | 
[**UnpublishPost**](PostV1alpha1ConsoleAPI.md#UnpublishPost) | **Put** /apis/api.console.halo.run/v1alpha1/posts/{name}/unpublish | 
[**UpdateDraftPost**](PostV1alpha1ConsoleAPI.md#UpdateDraftPost) | **Put** /apis/api.console.halo.run/v1alpha1/posts/{name} | 
[**UpdatePostContent**](PostV1alpha1ConsoleAPI.md#UpdatePostContent) | **Put** /apis/api.console.halo.run/v1alpha1/posts/{name}/content | 



## DeletePostContent

> ContentWrapper DeletePostContent(ctx, name).SnapshotName(snapshotName).Execute()





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
	name := "name_example" // string | 
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.DeletePostContent(context.Background(), name).SnapshotName(snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.DeletePostContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePostContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.DeletePostContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotName** | **string** |  | 

### Return type

[**ContentWrapper**](ContentWrapper.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DraftPost

> Post DraftPost(ctx).PostRequest(postRequest).Execute()





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
	postRequest := *openapiclient.NewPostRequest(*openapiclient.NewPost("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewPostSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example"))) // PostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.DraftPost(context.Background()).PostRequest(postRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.DraftPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DraftPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.DraftPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDraftPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postRequest** | [**PostRequest**](PostRequest.md) |  | 

### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPostContent

> ContentWrapper FetchPostContent(ctx, name).SnapshotName(snapshotName).Execute()





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
	name := "name_example" // string | 
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.FetchPostContent(context.Background(), name).SnapshotName(snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.FetchPostContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPostContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.FetchPostContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPostContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotName** | **string** |  | 

### Return type

[**ContentWrapper**](ContentWrapper.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPostHeadContent

> ContentWrapper FetchPostHeadContent(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.FetchPostHeadContent(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.FetchPostHeadContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPostHeadContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.FetchPostHeadContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPostHeadContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentWrapper**](ContentWrapper.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPostReleaseContent

> ContentWrapper FetchPostReleaseContent(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.FetchPostReleaseContent(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.FetchPostReleaseContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPostReleaseContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.FetchPostReleaseContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPostReleaseContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentWrapper**](ContentWrapper.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostSnapshots

> []ListedSnapshotDto ListPostSnapshots(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.ListPostSnapshots(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.ListPostSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostSnapshots`: []ListedSnapshotDto
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.ListPostSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPostSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListedSnapshotDto**](ListedSnapshotDto.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPosts

> ListedPostList ListPosts(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).PublishPhase(publishPhase).Keyword(keyword).CategoryWithChildren(categoryWithChildren).Execute()





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
	publishPhase := "publishPhase_example" // string | Posts filtered by publish phase. (optional)
	keyword := "keyword_example" // string | Posts filtered by keyword. (optional)
	categoryWithChildren := "categoryWithChildren_example" // string | Posts filtered by category including sub-categories. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.ListPosts(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).PublishPhase(publishPhase).Keyword(keyword).CategoryWithChildren(categoryWithChildren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.ListPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPosts`: ListedPostList
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.ListPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostsRequest struct via the builder pattern


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


## PublishPost

> Post PublishPost(ctx, name).HeadSnapshot(headSnapshot).Async(async).Execute()





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
	name := "name_example" // string | 
	headSnapshot := "headSnapshot_example" // string | Head snapshot name of content. (optional)
	async := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.PublishPost(context.Background(), name).HeadSnapshot(headSnapshot).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.PublishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.PublishPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **headSnapshot** | **string** | Head snapshot name of content. | 
 **async** | **bool** |  | 

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


## RecyclePost

> RecyclePost(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostV1alpha1ConsoleAPI.RecyclePost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.RecyclePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRecyclePostRequest struct via the builder pattern


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


## RevertToSpecifiedSnapshotForPost

> Post RevertToSpecifiedSnapshotForPost(ctx, name).RevertSnapshotForPostParam(revertSnapshotForPostParam).Execute()





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
	name := "name_example" // string | 
	revertSnapshotForPostParam := *openapiclient.NewRevertSnapshotForPostParam("SnapshotName_example") // RevertSnapshotForPostParam | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.RevertToSpecifiedSnapshotForPost(context.Background(), name).RevertSnapshotForPostParam(revertSnapshotForPostParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.RevertToSpecifiedSnapshotForPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertToSpecifiedSnapshotForPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.RevertToSpecifiedSnapshotForPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertToSpecifiedSnapshotForPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revertSnapshotForPostParam** | [**RevertSnapshotForPostParam**](RevertSnapshotForPostParam.md) |  | 

### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishPost

> Post UnpublishPost(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.UnpublishPost(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.UnpublishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpublishPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.UnpublishPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishPostRequest struct via the builder pattern


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


## UpdateDraftPost

> Post UpdateDraftPost(ctx, name).PostRequest(postRequest).Execute()





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
	name := "name_example" // string | 
	postRequest := *openapiclient.NewPostRequest(*openapiclient.NewPost("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewPostSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example"))) // PostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.UpdateDraftPost(context.Background(), name).PostRequest(postRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.UpdateDraftPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraftPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.UpdateDraftPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postRequest** | [**PostRequest**](PostRequest.md) |  | 

### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostContent

> Post UpdatePostContent(ctx, name).Content(content).Execute()





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
	name := "name_example" // string | 
	content := *openapiclient.NewContent() // Content | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostV1alpha1ConsoleAPI.UpdatePostContent(context.Background(), name).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostV1alpha1ConsoleAPI.UpdatePostContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePostContent`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostV1alpha1ConsoleAPI.UpdatePostContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | [**Content**](Content.md) |  | 

### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

