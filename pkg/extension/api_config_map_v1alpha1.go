/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package extension_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


type ConfigMapV1alpha1API interface {

	/*
	CreateConfigMap Method for CreateConfigMap

	Create ConfigMap

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateConfigMapRequest
	*/
	CreateConfigMap(ctx context.Context) ApiCreateConfigMapRequest

	// CreateConfigMapExecute executes the request
	//  @return ConfigMap
	CreateConfigMapExecute(r ApiCreateConfigMapRequest) (*ConfigMap, *http.Response, error)

	/*
	DeleteConfigMap Method for DeleteConfigMap

	Delete ConfigMap

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of configmap
	@return ApiDeleteConfigMapRequest
	*/
	DeleteConfigMap(ctx context.Context, name string) ApiDeleteConfigMapRequest

	// DeleteConfigMapExecute executes the request
	DeleteConfigMapExecute(r ApiDeleteConfigMapRequest) (*http.Response, error)

	/*
	GetConfigMap Method for GetConfigMap

	Get ConfigMap

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of configmap
	@return ApiGetConfigMapRequest
	*/
	GetConfigMap(ctx context.Context, name string) ApiGetConfigMapRequest

	// GetConfigMapExecute executes the request
	//  @return ConfigMap
	GetConfigMapExecute(r ApiGetConfigMapRequest) (*ConfigMap, *http.Response, error)

	/*
	ListConfigMap Method for ListConfigMap

	List ConfigMap

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListConfigMapRequest
	*/
	ListConfigMap(ctx context.Context) ApiListConfigMapRequest

	// ListConfigMapExecute executes the request
	//  @return ConfigMapList
	ListConfigMapExecute(r ApiListConfigMapRequest) (*ConfigMapList, *http.Response, error)

	/*
	PatchConfigMap Method for PatchConfigMap

	Patch ConfigMap

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of configmap
	@return ApiPatchConfigMapRequest
	*/
	PatchConfigMap(ctx context.Context, name string) ApiPatchConfigMapRequest

	// PatchConfigMapExecute executes the request
	//  @return ConfigMap
	PatchConfigMapExecute(r ApiPatchConfigMapRequest) (*ConfigMap, *http.Response, error)

	/*
	UpdateConfigMap Method for UpdateConfigMap

	Update ConfigMap

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of configmap
	@return ApiUpdateConfigMapRequest
	*/
	UpdateConfigMap(ctx context.Context, name string) ApiUpdateConfigMapRequest

	// UpdateConfigMapExecute executes the request
	//  @return ConfigMap
	UpdateConfigMapExecute(r ApiUpdateConfigMapRequest) (*ConfigMap, *http.Response, error)
}

// ConfigMapV1alpha1APIService ConfigMapV1alpha1API service
type ConfigMapV1alpha1APIService service

type ApiCreateConfigMapRequest struct {
	ctx context.Context
	ApiService ConfigMapV1alpha1API
	configMap *ConfigMap
}

// Fresh configmap
func (r ApiCreateConfigMapRequest) ConfigMap(configMap ConfigMap) ApiCreateConfigMapRequest {
	r.configMap = &configMap
	return r
}

func (r ApiCreateConfigMapRequest) Execute() (*ConfigMap, *http.Response, error) {
	return r.ApiService.CreateConfigMapExecute(r)
}

/*
CreateConfigMap Method for CreateConfigMap

Create ConfigMap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateConfigMapRequest
*/
func (a *ConfigMapV1alpha1APIService) CreateConfigMap(ctx context.Context) ApiCreateConfigMapRequest {
	return ApiCreateConfigMapRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigMap
func (a *ConfigMapV1alpha1APIService) CreateConfigMapExecute(r ApiCreateConfigMapRequest) (*ConfigMap, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigMap
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigMapV1alpha1APIService.CreateConfigMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1alpha1/configmaps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.configMap
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteConfigMapRequest struct {
	ctx context.Context
	ApiService ConfigMapV1alpha1API
	name string
}

func (r ApiDeleteConfigMapRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteConfigMapExecute(r)
}

/*
DeleteConfigMap Method for DeleteConfigMap

Delete ConfigMap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of configmap
 @return ApiDeleteConfigMapRequest
*/
func (a *ConfigMapV1alpha1APIService) DeleteConfigMap(ctx context.Context, name string) ApiDeleteConfigMapRequest {
	return ApiDeleteConfigMapRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
func (a *ConfigMapV1alpha1APIService) DeleteConfigMapExecute(r ApiDeleteConfigMapRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigMapV1alpha1APIService.DeleteConfigMap")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1alpha1/configmaps/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetConfigMapRequest struct {
	ctx context.Context
	ApiService ConfigMapV1alpha1API
	name string
}

func (r ApiGetConfigMapRequest) Execute() (*ConfigMap, *http.Response, error) {
	return r.ApiService.GetConfigMapExecute(r)
}

/*
GetConfigMap Method for GetConfigMap

Get ConfigMap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of configmap
 @return ApiGetConfigMapRequest
*/
func (a *ConfigMapV1alpha1APIService) GetConfigMap(ctx context.Context, name string) ApiGetConfigMapRequest {
	return ApiGetConfigMapRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return ConfigMap
func (a *ConfigMapV1alpha1APIService) GetConfigMapExecute(r ApiGetConfigMapRequest) (*ConfigMap, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigMap
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigMapV1alpha1APIService.GetConfigMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1alpha1/configmaps/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListConfigMapRequest struct {
	ctx context.Context
	ApiService ConfigMapV1alpha1API
	page *int32
	size *int32
	labelSelector *[]string
	fieldSelector *[]string
	sort *[]string
}

// Page number. Default is 0.
func (r ApiListConfigMapRequest) Page(page int32) ApiListConfigMapRequest {
	r.page = &page
	return r
}

// Size number. Default is 0.
func (r ApiListConfigMapRequest) Size(size int32) ApiListConfigMapRequest {
	r.size = &size
	return r
}

// Label selector. e.g.: hidden!&#x3D;true
func (r ApiListConfigMapRequest) LabelSelector(labelSelector []string) ApiListConfigMapRequest {
	r.labelSelector = &labelSelector
	return r
}

// Field selector. e.g.: metadata.name&#x3D;&#x3D;halo
func (r ApiListConfigMapRequest) FieldSelector(fieldSelector []string) ApiListConfigMapRequest {
	r.fieldSelector = &fieldSelector
	return r
}

// Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
func (r ApiListConfigMapRequest) Sort(sort []string) ApiListConfigMapRequest {
	r.sort = &sort
	return r
}

func (r ApiListConfigMapRequest) Execute() (*ConfigMapList, *http.Response, error) {
	return r.ApiService.ListConfigMapExecute(r)
}

/*
ListConfigMap Method for ListConfigMap

List ConfigMap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListConfigMapRequest
*/
func (a *ConfigMapV1alpha1APIService) ListConfigMap(ctx context.Context) ApiListConfigMapRequest {
	return ApiListConfigMapRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigMapList
func (a *ConfigMapV1alpha1APIService) ListConfigMapExecute(r ApiListConfigMapRequest) (*ConfigMapList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigMapList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigMapV1alpha1APIService.ListConfigMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1alpha1/configmaps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.labelSelector != nil {
		t := *r.labelSelector
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "labelSelector", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "labelSelector", t, "form", "multi")
		}
	}
	if r.fieldSelector != nil {
		t := *r.fieldSelector
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fieldSelector", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fieldSelector", t, "form", "multi")
		}
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sort", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sort", t, "form", "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchConfigMapRequest struct {
	ctx context.Context
	ApiService ConfigMapV1alpha1API
	name string
	jsonPatchInner *[]JsonPatchInner
}

func (r ApiPatchConfigMapRequest) JsonPatchInner(jsonPatchInner []JsonPatchInner) ApiPatchConfigMapRequest {
	r.jsonPatchInner = &jsonPatchInner
	return r
}

func (r ApiPatchConfigMapRequest) Execute() (*ConfigMap, *http.Response, error) {
	return r.ApiService.PatchConfigMapExecute(r)
}

/*
PatchConfigMap Method for PatchConfigMap

Patch ConfigMap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of configmap
 @return ApiPatchConfigMapRequest
*/
func (a *ConfigMapV1alpha1APIService) PatchConfigMap(ctx context.Context, name string) ApiPatchConfigMapRequest {
	return ApiPatchConfigMapRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return ConfigMap
func (a *ConfigMapV1alpha1APIService) PatchConfigMapExecute(r ApiPatchConfigMapRequest) (*ConfigMap, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigMap
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigMapV1alpha1APIService.PatchConfigMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1alpha1/configmaps/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.jsonPatchInner
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateConfigMapRequest struct {
	ctx context.Context
	ApiService ConfigMapV1alpha1API
	name string
	configMap *ConfigMap
}

// Updated configmap
func (r ApiUpdateConfigMapRequest) ConfigMap(configMap ConfigMap) ApiUpdateConfigMapRequest {
	r.configMap = &configMap
	return r
}

func (r ApiUpdateConfigMapRequest) Execute() (*ConfigMap, *http.Response, error) {
	return r.ApiService.UpdateConfigMapExecute(r)
}

/*
UpdateConfigMap Method for UpdateConfigMap

Update ConfigMap

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of configmap
 @return ApiUpdateConfigMapRequest
*/
func (a *ConfigMapV1alpha1APIService) UpdateConfigMap(ctx context.Context, name string) ApiUpdateConfigMapRequest {
	return ApiUpdateConfigMapRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return ConfigMap
func (a *ConfigMapV1alpha1APIService) UpdateConfigMapExecute(r ApiUpdateConfigMapRequest) (*ConfigMap, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigMap
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigMapV1alpha1APIService.UpdateConfigMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1alpha1/configmaps/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.configMap
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
