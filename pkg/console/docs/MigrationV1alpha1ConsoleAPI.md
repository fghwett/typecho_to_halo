# \MigrationV1alpha1ConsoleAPI

All URIs are relative to *http://localhost:8091*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadBackups**](MigrationV1alpha1ConsoleAPI.md#DownloadBackups) | **Get** /apis/console.api.migration.halo.run/v1alpha1/backups/{name}/files/{filename} | 
[**GetBackupFiles**](MigrationV1alpha1ConsoleAPI.md#GetBackupFiles) | **Get** /apis/console.api.migration.halo.run/v1alpha1/backup-files | 
[**RestoreBackup**](MigrationV1alpha1ConsoleAPI.md#RestoreBackup) | **Post** /apis/console.api.migration.halo.run/v1alpha1/restorations | 



## DownloadBackups

> DownloadBackups(ctx, name, filename).Execute()



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
	name := "name_example" // string | Backup name.
	filename := "filename_example" // string | Backup filename.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MigrationV1alpha1ConsoleAPI.DownloadBackups(context.Background(), name, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationV1alpha1ConsoleAPI.DownloadBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Backup name. | 
**filename** | **string** | Backup filename. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBackupsRequest struct via the builder pattern


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


## GetBackupFiles

> []BackupFile GetBackupFiles(ctx).Execute()





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
	resp, r, err := apiClient.MigrationV1alpha1ConsoleAPI.GetBackupFiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationV1alpha1ConsoleAPI.GetBackupFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupFiles`: []BackupFile
	fmt.Fprintf(os.Stdout, "Response from `MigrationV1alpha1ConsoleAPI.GetBackupFiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupFilesRequest struct via the builder pattern


### Return type

[**[]BackupFile**](BackupFile.md)

### Authorization

[basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreBackup

> RestoreBackup(ctx).BackupName(backupName).DownloadUrl(downloadUrl).File(file).Filename(filename).Execute()





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
	backupName := "backupName_example" // string | Backup metadata name. (optional)
	downloadUrl := "downloadUrl_example" // string | Remote backup HTTP URL. (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	filename := "filename_example" // string | Filename of backup file in backups root. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MigrationV1alpha1ConsoleAPI.RestoreBackup(context.Background()).BackupName(backupName).DownloadUrl(downloadUrl).File(file).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationV1alpha1ConsoleAPI.RestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupName** | **string** | Backup metadata name. | 
 **downloadUrl** | **string** | Remote backup HTTP URL. | 
 **file** | ***os.File** |  | 
 **filename** | **string** | Filename of backup file in backups root. | 

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

