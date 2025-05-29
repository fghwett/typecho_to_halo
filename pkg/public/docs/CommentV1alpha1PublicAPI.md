# \CommentV1alpha1PublicAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComment1**](CommentV1alpha1PublicAPI.md#CreateComment1) | **Post** /apis/api.halo.run/v1alpha1/comments | 
[**CreateReply1**](CommentV1alpha1PublicAPI.md#CreateReply1) | **Post** /apis/api.halo.run/v1alpha1/comments/{name}/reply | 
[**GetComment**](CommentV1alpha1PublicAPI.md#GetComment) | **Get** /apis/api.halo.run/v1alpha1/comments/{name} | 
[**ListCommentReplies**](CommentV1alpha1PublicAPI.md#ListCommentReplies) | **Get** /apis/api.halo.run/v1alpha1/comments/{name}/reply | 
[**ListComments1**](CommentV1alpha1PublicAPI.md#ListComments1) | **Get** /apis/api.halo.run/v1alpha1/comments | 



## CreateComment1

> Comment CreateComment1(ctx).CommentRequest(commentRequest).Execute()





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
	commentRequest := *openapiclient.NewCommentRequest("Content_example", "Raw_example", *openapiclient.NewRef("Name_example")) // CommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentV1alpha1PublicAPI.CreateComment1(context.Background()).CommentRequest(commentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentV1alpha1PublicAPI.CreateComment1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment1`: Comment
	fmt.Fprintf(os.Stdout, "Response from `CommentV1alpha1PublicAPI.CreateComment1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateComment1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentRequest** | [**CommentRequest**](CommentRequest.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReply1

> Reply CreateReply1(ctx, name).ReplyRequest(replyRequest).Execute()





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
	name := "name_example" // string | 
	replyRequest := *openapiclient.NewReplyRequest("Content_example", "Raw_example") // ReplyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentV1alpha1PublicAPI.CreateReply1(context.Background(), name).ReplyRequest(replyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentV1alpha1PublicAPI.CreateReply1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReply1`: Reply
	fmt.Fprintf(os.Stdout, "Response from `CommentV1alpha1PublicAPI.CreateReply1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReply1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replyRequest** | [**ReplyRequest**](ReplyRequest.md) |  | 

### Return type

[**Reply**](Reply.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment

> CommentVoList GetComment(ctx, name).Execute()





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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentV1alpha1PublicAPI.GetComment(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentV1alpha1PublicAPI.GetComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment`: CommentVoList
	fmt.Fprintf(os.Stdout, "Response from `CommentV1alpha1PublicAPI.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommentVoList**](CommentVoList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommentReplies

> ReplyVoList ListCommentReplies(ctx, name).Page(page).Size(size).Execute()





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
	name := "name_example" // string | 
	page := int32(56) // int32 | Page number. Default is 0. (optional)
	size := int32(56) // int32 | Size number. Default is 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentV1alpha1PublicAPI.ListCommentReplies(context.Background(), name).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentV1alpha1PublicAPI.ListCommentReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommentReplies`: ReplyVoList
	fmt.Fprintf(os.Stdout, "Response from `CommentV1alpha1PublicAPI.ListCommentReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommentRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 

### Return type

[**ReplyVoList**](ReplyVoList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComments1

> CommentWithReplyVoList ListComments1(ctx).Version(version).Kind(kind).Name(name).Page(page).Size(size).Sort(sort).Group(group).WithReplies(withReplies).ReplySize(replySize).Execute()





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
	version := "version_example" // string | The comment subject version.
	kind := "kind_example" // string | The comment subject kind.
	name := "name_example" // string | The comment subject name.
	page := int32(56) // int32 | Page number. Default is 0. (optional)
	size := int32(56) // int32 | Size number. Default is 0. (optional)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	group := "group_example" // string | The comment subject group. (optional)
	withReplies := true // bool | Whether to include replies. Default is false. (optional)
	replySize := int32(56) // int32 | Reply size of the comment, default is 10, only works when withReplies is true. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentV1alpha1PublicAPI.ListComments1(context.Background()).Version(version).Kind(kind).Name(name).Page(page).Size(size).Sort(sort).Group(group).WithReplies(withReplies).ReplySize(replySize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentV1alpha1PublicAPI.ListComments1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListComments1`: CommentWithReplyVoList
	fmt.Fprintf(os.Stdout, "Response from `CommentV1alpha1PublicAPI.ListComments1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListComments1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The comment subject version. | 
 **kind** | **string** | The comment subject kind. | 
 **name** | **string** | The comment subject name. | 
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **group** | **string** | The comment subject group. | 
 **withReplies** | **bool** | Whether to include replies. Default is false. | 
 **replySize** | **int32** | Reply size of the comment, default is 10, only works when withReplies is true. | 

### Return type

[**CommentWithReplyVoList**](CommentWithReplyVoList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

