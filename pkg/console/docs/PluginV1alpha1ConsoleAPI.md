# \PluginV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePluginRunningState**](PluginV1alpha1ConsoleAPI.md#ChangePluginRunningState) | **Put** /apis/api.console.halo.run/v1alpha1/plugins/{name}/plugin-state | 
[**FetchCssBundle**](PluginV1alpha1ConsoleAPI.md#FetchCssBundle) | **Get** /apis/api.console.halo.run/v1alpha1/plugins/-/bundle.css | 
[**FetchJsBundle**](PluginV1alpha1ConsoleAPI.md#FetchJsBundle) | **Get** /apis/api.console.halo.run/v1alpha1/plugins/-/bundle.js | 
[**FetchPluginConfig**](PluginV1alpha1ConsoleAPI.md#FetchPluginConfig) | **Get** /apis/api.console.halo.run/v1alpha1/plugins/{name}/config | 
[**FetchPluginJsonConfig**](PluginV1alpha1ConsoleAPI.md#FetchPluginJsonConfig) | **Get** /apis/api.console.halo.run/v1alpha1/plugins/{name}/json-config | 
[**FetchPluginSetting**](PluginV1alpha1ConsoleAPI.md#FetchPluginSetting) | **Get** /apis/api.console.halo.run/v1alpha1/plugins/{name}/setting | 
[**InstallPlugin**](PluginV1alpha1ConsoleAPI.md#InstallPlugin) | **Post** /apis/api.console.halo.run/v1alpha1/plugins/install | 
[**InstallPluginFromUri**](PluginV1alpha1ConsoleAPI.md#InstallPluginFromUri) | **Post** /apis/api.console.halo.run/v1alpha1/plugins/-/install-from-uri | 
[**ListPlugins**](PluginV1alpha1ConsoleAPI.md#ListPlugins) | **Get** /apis/api.console.halo.run/v1alpha1/plugins | 
[**ReloadPlugin**](PluginV1alpha1ConsoleAPI.md#ReloadPlugin) | **Put** /apis/api.console.halo.run/v1alpha1/plugins/{name}/reload | 
[**ResetPluginConfig**](PluginV1alpha1ConsoleAPI.md#ResetPluginConfig) | **Put** /apis/api.console.halo.run/v1alpha1/plugins/{name}/reset-config | 
[**UpdatePluginConfig**](PluginV1alpha1ConsoleAPI.md#UpdatePluginConfig) | **Put** /apis/api.console.halo.run/v1alpha1/plugins/{name}/config | 
[**UpdatePluginJsonConfig**](PluginV1alpha1ConsoleAPI.md#UpdatePluginJsonConfig) | **Put** /apis/api.console.halo.run/v1alpha1/plugins/{name}/json-config | 
[**UpgradePlugin**](PluginV1alpha1ConsoleAPI.md#UpgradePlugin) | **Post** /apis/api.console.halo.run/v1alpha1/plugins/{name}/upgrade | 
[**UpgradePluginFromUri**](PluginV1alpha1ConsoleAPI.md#UpgradePluginFromUri) | **Post** /apis/api.console.halo.run/v1alpha1/plugins/{name}/upgrade-from-uri | 



## ChangePluginRunningState

> Plugin ChangePluginRunningState(ctx, name).PluginRunningStateRequest(pluginRunningStateRequest).Execute()





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
	pluginRunningStateRequest := *openapiclient.NewPluginRunningStateRequest() // PluginRunningStateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.ChangePluginRunningState(context.Background(), name).PluginRunningStateRequest(pluginRunningStateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.ChangePluginRunningState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePluginRunningState`: Plugin
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.ChangePluginRunningState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePluginRunningStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pluginRunningStateRequest** | [**PluginRunningStateRequest**](PluginRunningStateRequest.md) |  | 

### Return type

[**Plugin**](Plugin.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCssBundle

> string FetchCssBundle(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.FetchCssBundle(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.FetchCssBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCssBundle`: string
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.FetchCssBundle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCssBundleRequest struct via the builder pattern


### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchJsBundle

> string FetchJsBundle(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.FetchJsBundle(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.FetchJsBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchJsBundle`: string
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.FetchJsBundle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchJsBundleRequest struct via the builder pattern


### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPluginConfig

> ConfigMap FetchPluginConfig(ctx, name).Execute()





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
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.FetchPluginConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.FetchPluginConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPluginConfig`: ConfigMap
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.FetchPluginConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPluginConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigMap**](ConfigMap.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPluginJsonConfig

> map[string]interface{} FetchPluginJsonConfig(ctx, name).Execute()





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
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.FetchPluginJsonConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.FetchPluginJsonConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPluginJsonConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.FetchPluginJsonConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPluginJsonConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPluginSetting

> Setting FetchPluginSetting(ctx, name).Execute()





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
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.FetchPluginSetting(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.FetchPluginSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchPluginSetting`: Setting
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.FetchPluginSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPluginSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Setting**](Setting.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallPlugin

> Plugin InstallPlugin(ctx).File(file).PresetName(presetName).Source(source).Execute()





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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	presetName := "presetName_example" // string | Plugin preset name. We will find the plugin from plugin presets (optional)
	source := "source_example" // string | Install source. Default is file. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.InstallPlugin(context.Background()).File(file).PresetName(presetName).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.InstallPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallPlugin`: Plugin
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.InstallPlugin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstallPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **presetName** | **string** | Plugin preset name. We will find the plugin from plugin presets | 
 **source** | **string** | Install source. Default is file. | 

### Return type

[**Plugin**](Plugin.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallPluginFromUri

> Plugin InstallPluginFromUri(ctx).InstallFromUriRequest(installFromUriRequest).Execute()





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
	installFromUriRequest := *openapiclient.NewInstallFromUriRequest("Uri_example") // InstallFromUriRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.InstallPluginFromUri(context.Background()).InstallFromUriRequest(installFromUriRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.InstallPluginFromUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallPluginFromUri`: Plugin
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.InstallPluginFromUri`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstallPluginFromUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installFromUriRequest** | [**InstallFromUriRequest**](InstallFromUriRequest.md) |  | 

### Return type

[**Plugin**](Plugin.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlugins

> PluginList ListPlugins(ctx).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Keyword(keyword).Enabled(enabled).Execute()





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
	keyword := "keyword_example" // string | Keyword of plugin name or description (optional)
	enabled := true // bool | Whether the plugin is enabled (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.ListPlugins(context.Background()).Page(page).Size(size).LabelSelector(labelSelector).FieldSelector(fieldSelector).Sort(sort).Keyword(keyword).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.ListPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlugins`: PluginList
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.ListPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Default is 0. | 
 **size** | **int32** | Size number. Default is 0. | 
 **labelSelector** | **[]string** | Label selector. e.g.: hidden!&#x3D;true | 
 **fieldSelector** | **[]string** | Field selector. e.g.: metadata.name&#x3D;&#x3D;halo | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **keyword** | **string** | Keyword of plugin name or description | 
 **enabled** | **bool** | Whether the plugin is enabled | 

### Return type

[**PluginList**](PluginList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadPlugin

> Plugin ReloadPlugin(ctx, name).Execute()





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
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.ReloadPlugin(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.ReloadPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReloadPlugin`: Plugin
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.ReloadPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReloadPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plugin**](Plugin.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPluginConfig

> ConfigMap ResetPluginConfig(ctx, name).Execute()





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
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.ResetPluginConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.ResetPluginConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetPluginConfig`: ConfigMap
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.ResetPluginConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPluginConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigMap**](ConfigMap.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePluginConfig

> ConfigMap UpdatePluginConfig(ctx, name).ConfigMap(configMap).Execute()





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
	configMap := *openapiclient.NewConfigMap("ApiVersion_example", "Kind_example", *openapiclient.NewMetadata("Name_example")) // ConfigMap | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.UpdatePluginConfig(context.Background(), name).ConfigMap(configMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.UpdatePluginConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePluginConfig`: ConfigMap
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.UpdatePluginConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePluginConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configMap** | [**ConfigMap**](ConfigMap.md) |  | 

### Return type

[**ConfigMap**](ConfigMap.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePluginJsonConfig

> UpdatePluginJsonConfig(ctx, name).Body(body).Execute()





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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginV1alpha1ConsoleAPI.UpdatePluginJsonConfig(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.UpdatePluginJsonConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePluginJsonConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePlugin

> UpgradePlugin(ctx, name).File(file).PresetName(presetName).Source(source).Execute()





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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	presetName := "presetName_example" // string | Plugin preset name. We will find the plugin from plugin presets (optional)
	source := "source_example" // string | Install source. Default is file. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginV1alpha1ConsoleAPI.UpgradePlugin(context.Background(), name).File(file).PresetName(presetName).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.UpgradePlugin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpgradePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **presetName** | **string** | Plugin preset name. We will find the plugin from plugin presets | 
 **source** | **string** | Install source. Default is file. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePluginFromUri

> Plugin UpgradePluginFromUri(ctx, name).UpgradeFromUriRequest(upgradeFromUriRequest).Execute()





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
	upgradeFromUriRequest := *openapiclient.NewUpgradeFromUriRequest("Uri_example") // UpgradeFromUriRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginV1alpha1ConsoleAPI.UpgradePluginFromUri(context.Background(), name).UpgradeFromUriRequest(upgradeFromUriRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginV1alpha1ConsoleAPI.UpgradePluginFromUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradePluginFromUri`: Plugin
	fmt.Fprintf(os.Stdout, "Response from `PluginV1alpha1ConsoleAPI.UpgradePluginFromUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradePluginFromUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeFromUriRequest** | [**UpgradeFromUriRequest**](UpgradeFromUriRequest.md) |  | 

### Return type

[**Plugin**](Plugin.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

