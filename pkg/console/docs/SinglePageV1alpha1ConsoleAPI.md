# \SinglePageV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSinglePageContent**](SinglePageV1alpha1ConsoleAPI.md#DeleteSinglePageContent) | **Delete** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/content | 
[**DraftSinglePage**](SinglePageV1alpha1ConsoleAPI.md#DraftSinglePage) | **Post** /apis/api.console.halo.run/v1alpha1/singlepages | 
[**FetchSinglePageContent**](SinglePageV1alpha1ConsoleAPI.md#FetchSinglePageContent) | **Get** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/content | 
[**FetchSinglePageHeadContent**](SinglePageV1alpha1ConsoleAPI.md#FetchSinglePageHeadContent) | **Get** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/head-content | 
[**FetchSinglePageReleaseContent**](SinglePageV1alpha1ConsoleAPI.md#FetchSinglePageReleaseContent) | **Get** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/release-content | 
[**ListSinglePageSnapshots**](SinglePageV1alpha1ConsoleAPI.md#ListSinglePageSnapshots) | **Get** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/snapshot | 
[**ListSinglePages**](SinglePageV1alpha1ConsoleAPI.md#ListSinglePages) | **Get** /apis/api.console.halo.run/v1alpha1/singlepages | 
[**PublishSinglePage**](SinglePageV1alpha1ConsoleAPI.md#PublishSinglePage) | **Put** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/publish | 
[**RevertToSpecifiedSnapshotForSinglePage**](SinglePageV1alpha1ConsoleAPI.md#RevertToSpecifiedSnapshotForSinglePage) | **Put** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/revert-content | 
[**UpdateDraftSinglePage**](SinglePageV1alpha1ConsoleAPI.md#UpdateDraftSinglePage) | **Put** /apis/api.console.halo.run/v1alpha1/singlepages/{name} | 
[**UpdateSinglePageContent**](SinglePageV1alpha1ConsoleAPI.md#UpdateSinglePageContent) | **Put** /apis/api.console.halo.run/v1alpha1/singlepages/{name}/content | 



## DeleteSinglePageContent

> ContentWrapper DeleteSinglePageContent(ctx, name).SnapshotName(snapshotName).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.DeleteSinglePageContent(context.Background(), name).SnapshotName(snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.DeleteSinglePageContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSinglePageContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.DeleteSinglePageContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSinglePageContentRequest struct via the builder pattern


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


## DraftSinglePage

> SinglePage DraftSinglePage(ctx).SinglePageRequest(singlePageRequest).Execute()





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
	singlePageRequest := *openapiclient.NewSinglePageRequest(*openapiclient.NewContentUpdateParam(), *openapiclient.NewSinglePage("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewSinglePageSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example"))) // SinglePageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.DraftSinglePage(context.Background()).SinglePageRequest(singlePageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.DraftSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DraftSinglePage`: SinglePage
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.DraftSinglePage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDraftSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singlePageRequest** | [**SinglePageRequest**](SinglePageRequest.md) |  | 

### Return type

[**SinglePage**](SinglePage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSinglePageContent

> ContentWrapper FetchSinglePageContent(ctx, name).SnapshotName(snapshotName).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.FetchSinglePageContent(context.Background(), name).SnapshotName(snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.FetchSinglePageContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSinglePageContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.FetchSinglePageContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSinglePageContentRequest struct via the builder pattern


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


## FetchSinglePageHeadContent

> ContentWrapper FetchSinglePageHeadContent(ctx, name).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.FetchSinglePageHeadContent(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.FetchSinglePageHeadContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSinglePageHeadContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.FetchSinglePageHeadContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSinglePageHeadContentRequest struct via the builder pattern


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


## FetchSinglePageReleaseContent

> ContentWrapper FetchSinglePageReleaseContent(ctx, name).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.FetchSinglePageReleaseContent(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.FetchSinglePageReleaseContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchSinglePageReleaseContent`: ContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.FetchSinglePageReleaseContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSinglePageReleaseContentRequest struct via the builder pattern


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


## ListSinglePageSnapshots

> []ListedSnapshotDto ListSinglePageSnapshots(ctx, name).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.ListSinglePageSnapshots(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.ListSinglePageSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSinglePageSnapshots`: []ListedSnapshotDto
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.ListSinglePageSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSinglePageSnapshotsRequest struct via the builder pattern


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


## ListSinglePages

> ListedSinglePageList ListSinglePages(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Contributor(contributor).PublishPhase(publishPhase).Visible(visible).Keyword(keyword).Execute()





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
	contributor := []string{"Inner_example"} // []string | SinglePages filtered by contributor. (optional)
	publishPhase := "publishPhase_example" // string | SinglePages filtered by publish phase. (optional)
	visible := "visible_example" // string | SinglePages filtered by visibility. (optional)
	keyword := "keyword_example" // string | SinglePages filtered by keyword. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.ListSinglePages(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Contributor(contributor).PublishPhase(publishPhase).Visible(visible).Keyword(keyword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.ListSinglePages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSinglePages`: ListedSinglePageList
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.ListSinglePages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSinglePagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **contributor** | **[]string** | SinglePages filtered by contributor. | 
 **publishPhase** | **string** | SinglePages filtered by publish phase. | 
 **visible** | **string** | SinglePages filtered by visibility. | 
 **keyword** | **string** | SinglePages filtered by keyword. | 

### Return type

[**ListedSinglePageList**](ListedSinglePageList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishSinglePage

> SinglePage PublishSinglePage(ctx, name).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.PublishSinglePage(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.PublishSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishSinglePage`: SinglePage
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.PublishSinglePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SinglePage**](SinglePage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertToSpecifiedSnapshotForSinglePage

> Post RevertToSpecifiedSnapshotForSinglePage(ctx, name).RevertSnapshotForSingleParam(revertSnapshotForSingleParam).Execute()





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
	revertSnapshotForSingleParam := *openapiclient.NewRevertSnapshotForSingleParam("SnapshotName_example") // RevertSnapshotForSingleParam | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.RevertToSpecifiedSnapshotForSinglePage(context.Background(), name).RevertSnapshotForSingleParam(revertSnapshotForSingleParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.RevertToSpecifiedSnapshotForSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevertToSpecifiedSnapshotForSinglePage`: Post
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.RevertToSpecifiedSnapshotForSinglePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertToSpecifiedSnapshotForSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revertSnapshotForSingleParam** | [**RevertSnapshotForSingleParam**](RevertSnapshotForSingleParam.md) |  | 

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


## UpdateDraftSinglePage

> SinglePage UpdateDraftSinglePage(ctx, name).SinglePageRequest(singlePageRequest).Execute()





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
	singlePageRequest := *openapiclient.NewSinglePageRequest(*openapiclient.NewContentUpdateParam(), *openapiclient.NewSinglePage("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example"), *openapiclient.NewSinglePageSpec(false, false, *openapiclient.NewExcerpt(false), false, int32(123), false, "Slug_example", "Title_example", "Visible_example"))) // SinglePageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.UpdateDraftSinglePage(context.Background(), name).SinglePageRequest(singlePageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.UpdateDraftSinglePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraftSinglePage`: SinglePage
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.UpdateDraftSinglePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftSinglePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singlePageRequest** | [**SinglePageRequest**](SinglePageRequest.md) |  | 

### Return type

[**SinglePage**](SinglePage.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSinglePageContent

> Post UpdateSinglePageContent(ctx, name).Content(content).Execute()





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
	resp, r, err := apiClient.SinglePageV1alpha1ConsoleAPI.UpdateSinglePageContent(context.Background(), name).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinglePageV1alpha1ConsoleAPI.UpdateSinglePageContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSinglePageContent`: Post
	fmt.Fprintf(os.Stdout, "Response from `SinglePageV1alpha1ConsoleAPI.UpdateSinglePageContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSinglePageContentRequest struct via the builder pattern


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

