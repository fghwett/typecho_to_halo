# \TwoFactorAuthV1alpha1UcAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurerTotp**](TwoFactorAuthV1alpha1UcAPI.md#ConfigurerTotp) | **Post** /apis/uc.api.security.halo.run/v1alpha1/authentications/two-factor/totp | 
[**DeleteTotp**](TwoFactorAuthV1alpha1UcAPI.md#DeleteTotp) | **Delete** /apis/uc.api.security.halo.run/v1alpha1/authentications/two-factor/totp/- | 
[**DisableTwoFactor**](TwoFactorAuthV1alpha1UcAPI.md#DisableTwoFactor) | **Put** /apis/uc.api.security.halo.run/v1alpha1/authentications/two-factor/settings/disabled | 
[**EnableTwoFactor**](TwoFactorAuthV1alpha1UcAPI.md#EnableTwoFactor) | **Put** /apis/uc.api.security.halo.run/v1alpha1/authentications/two-factor/settings/enabled | 
[**GetTotpAuthLink**](TwoFactorAuthV1alpha1UcAPI.md#GetTotpAuthLink) | **Get** /apis/uc.api.security.halo.run/v1alpha1/authentications/two-factor/totp/auth-link | 
[**GetTwoFactorAuthenticationSettings**](TwoFactorAuthV1alpha1UcAPI.md#GetTwoFactorAuthenticationSettings) | **Get** /apis/uc.api.security.halo.run/v1alpha1/authentications/two-factor/settings | 



## ConfigurerTotp

> TwoFactorAuthSettings ConfigurerTotp(ctx).TotpRequest(totpRequest).Execute()





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
	totpRequest := *openapiclient.NewTotpRequest("Code_example", "Password_example", "Secret_example") // TotpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwoFactorAuthV1alpha1UcAPI.ConfigurerTotp(context.Background()).TotpRequest(totpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthV1alpha1UcAPI.ConfigurerTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigurerTotp`: TwoFactorAuthSettings
	fmt.Fprintf(os.Stdout, "Response from `TwoFactorAuthV1alpha1UcAPI.ConfigurerTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigurerTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totpRequest** | [**TotpRequest**](TotpRequest.md) |  | 

### Return type

[**TwoFactorAuthSettings**](TwoFactorAuthSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTotp

> TwoFactorAuthSettings DeleteTotp(ctx).PasswordRequest(passwordRequest).Execute()



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
	passwordRequest := *openapiclient.NewPasswordRequest("Password_example") // PasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwoFactorAuthV1alpha1UcAPI.DeleteTotp(context.Background()).PasswordRequest(passwordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthV1alpha1UcAPI.DeleteTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTotp`: TwoFactorAuthSettings
	fmt.Fprintf(os.Stdout, "Response from `TwoFactorAuthV1alpha1UcAPI.DeleteTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordRequest** | [**PasswordRequest**](PasswordRequest.md) |  | 

### Return type

[**TwoFactorAuthSettings**](TwoFactorAuthSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableTwoFactor

> TwoFactorAuthSettings DisableTwoFactor(ctx).PasswordRequest(passwordRequest).Execute()





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
	passwordRequest := *openapiclient.NewPasswordRequest("Password_example") // PasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwoFactorAuthV1alpha1UcAPI.DisableTwoFactor(context.Background()).PasswordRequest(passwordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthV1alpha1UcAPI.DisableTwoFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableTwoFactor`: TwoFactorAuthSettings
	fmt.Fprintf(os.Stdout, "Response from `TwoFactorAuthV1alpha1UcAPI.DisableTwoFactor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableTwoFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordRequest** | [**PasswordRequest**](PasswordRequest.md) |  | 

### Return type

[**TwoFactorAuthSettings**](TwoFactorAuthSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTwoFactor

> TwoFactorAuthSettings EnableTwoFactor(ctx).PasswordRequest(passwordRequest).Execute()





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
	passwordRequest := *openapiclient.NewPasswordRequest("Password_example") // PasswordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwoFactorAuthV1alpha1UcAPI.EnableTwoFactor(context.Background()).PasswordRequest(passwordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthV1alpha1UcAPI.EnableTwoFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableTwoFactor`: TwoFactorAuthSettings
	fmt.Fprintf(os.Stdout, "Response from `TwoFactorAuthV1alpha1UcAPI.EnableTwoFactor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableTwoFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordRequest** | [**PasswordRequest**](PasswordRequest.md) |  | 

### Return type

[**TwoFactorAuthSettings**](TwoFactorAuthSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotpAuthLink

> TotpAuthLinkResponse GetTotpAuthLink(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwoFactorAuthV1alpha1UcAPI.GetTotpAuthLink(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthV1alpha1UcAPI.GetTotpAuthLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotpAuthLink`: TotpAuthLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TwoFactorAuthV1alpha1UcAPI.GetTotpAuthLink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTotpAuthLinkRequest struct via the builder pattern


### Return type

[**TotpAuthLinkResponse**](TotpAuthLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTwoFactorAuthenticationSettings

> TwoFactorAuthSettings GetTwoFactorAuthenticationSettings(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TwoFactorAuthV1alpha1UcAPI.GetTwoFactorAuthenticationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TwoFactorAuthV1alpha1UcAPI.GetTwoFactorAuthenticationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTwoFactorAuthenticationSettings`: TwoFactorAuthSettings
	fmt.Fprintf(os.Stdout, "Response from `TwoFactorAuthV1alpha1UcAPI.GetTwoFactorAuthenticationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTwoFactorAuthenticationSettingsRequest struct via the builder pattern


### Return type

[**TwoFactorAuthSettings**](TwoFactorAuthSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

