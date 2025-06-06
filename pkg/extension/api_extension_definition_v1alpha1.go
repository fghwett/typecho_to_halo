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


type ExtensionDefinitionV1alpha1API interface {

	/*
	CreateExtensionDefinition Method for CreateExtensionDefinition

	Create ExtensionDefinition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateExtensionDefinitionRequest
	*/
	CreateExtensionDefinition(ctx context.Context) ApiCreateExtensionDefinitionRequest

	// CreateExtensionDefinitionExecute executes the request
	//  @return ExtensionDefinition
	CreateExtensionDefinitionExecute(r ApiCreateExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error)

	/*
	DeleteExtensionDefinition Method for DeleteExtensionDefinition

	Delete ExtensionDefinition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of extensiondefinition
	@return ApiDeleteExtensionDefinitionRequest
	*/
	DeleteExtensionDefinition(ctx context.Context, name string) ApiDeleteExtensionDefinitionRequest

	// DeleteExtensionDefinitionExecute executes the request
	DeleteExtensionDefinitionExecute(r ApiDeleteExtensionDefinitionRequest) (*http.Response, error)

	/*
	GetExtensionDefinition Method for GetExtensionDefinition

	Get ExtensionDefinition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of extensiondefinition
	@return ApiGetExtensionDefinitionRequest
	*/
	GetExtensionDefinition(ctx context.Context, name string) ApiGetExtensionDefinitionRequest

	// GetExtensionDefinitionExecute executes the request
	//  @return ExtensionDefinition
	GetExtensionDefinitionExecute(r ApiGetExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error)

	/*
	ListExtensionDefinition Method for ListExtensionDefinition

	List ExtensionDefinition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListExtensionDefinitionRequest
	*/
	ListExtensionDefinition(ctx context.Context) ApiListExtensionDefinitionRequest

	// ListExtensionDefinitionExecute executes the request
	//  @return ExtensionDefinitionList
	ListExtensionDefinitionExecute(r ApiListExtensionDefinitionRequest) (*ExtensionDefinitionList, *http.Response, error)

	/*
	PatchExtensionDefinition Method for PatchExtensionDefinition

	Patch ExtensionDefinition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of extensiondefinition
	@return ApiPatchExtensionDefinitionRequest
	*/
	PatchExtensionDefinition(ctx context.Context, name string) ApiPatchExtensionDefinitionRequest

	// PatchExtensionDefinitionExecute executes the request
	//  @return ExtensionDefinition
	PatchExtensionDefinitionExecute(r ApiPatchExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error)

	/*
	UpdateExtensionDefinition Method for UpdateExtensionDefinition

	Update ExtensionDefinition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of extensiondefinition
	@return ApiUpdateExtensionDefinitionRequest
	*/
	UpdateExtensionDefinition(ctx context.Context, name string) ApiUpdateExtensionDefinitionRequest

	// UpdateExtensionDefinitionExecute executes the request
	//  @return ExtensionDefinition
	UpdateExtensionDefinitionExecute(r ApiUpdateExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error)
}

// ExtensionDefinitionV1alpha1APIService ExtensionDefinitionV1alpha1API service
type ExtensionDefinitionV1alpha1APIService service

type ApiCreateExtensionDefinitionRequest struct {
	ctx context.Context
	ApiService ExtensionDefinitionV1alpha1API
	extensionDefinition *ExtensionDefinition
}

// Fresh extensiondefinition
func (r ApiCreateExtensionDefinitionRequest) ExtensionDefinition(extensionDefinition ExtensionDefinition) ApiCreateExtensionDefinitionRequest {
	r.extensionDefinition = &extensionDefinition
	return r
}

func (r ApiCreateExtensionDefinitionRequest) Execute() (*ExtensionDefinition, *http.Response, error) {
	return r.ApiService.CreateExtensionDefinitionExecute(r)
}

/*
CreateExtensionDefinition Method for CreateExtensionDefinition

Create ExtensionDefinition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateExtensionDefinitionRequest
*/
func (a *ExtensionDefinitionV1alpha1APIService) CreateExtensionDefinition(ctx context.Context) ApiCreateExtensionDefinitionRequest {
	return ApiCreateExtensionDefinitionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExtensionDefinition
func (a *ExtensionDefinitionV1alpha1APIService) CreateExtensionDefinitionExecute(r ApiCreateExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtensionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionDefinitionV1alpha1APIService.CreateExtensionDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/plugin.halo.run/v1alpha1/extensiondefinitions"

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
	localVarPostBody = r.extensionDefinition
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

type ApiDeleteExtensionDefinitionRequest struct {
	ctx context.Context
	ApiService ExtensionDefinitionV1alpha1API
	name string
}

func (r ApiDeleteExtensionDefinitionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExtensionDefinitionExecute(r)
}

/*
DeleteExtensionDefinition Method for DeleteExtensionDefinition

Delete ExtensionDefinition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of extensiondefinition
 @return ApiDeleteExtensionDefinitionRequest
*/
func (a *ExtensionDefinitionV1alpha1APIService) DeleteExtensionDefinition(ctx context.Context, name string) ApiDeleteExtensionDefinitionRequest {
	return ApiDeleteExtensionDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
func (a *ExtensionDefinitionV1alpha1APIService) DeleteExtensionDefinitionExecute(r ApiDeleteExtensionDefinitionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionDefinitionV1alpha1APIService.DeleteExtensionDefinition")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name}"
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

type ApiGetExtensionDefinitionRequest struct {
	ctx context.Context
	ApiService ExtensionDefinitionV1alpha1API
	name string
}

func (r ApiGetExtensionDefinitionRequest) Execute() (*ExtensionDefinition, *http.Response, error) {
	return r.ApiService.GetExtensionDefinitionExecute(r)
}

/*
GetExtensionDefinition Method for GetExtensionDefinition

Get ExtensionDefinition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of extensiondefinition
 @return ApiGetExtensionDefinitionRequest
*/
func (a *ExtensionDefinitionV1alpha1APIService) GetExtensionDefinition(ctx context.Context, name string) ApiGetExtensionDefinitionRequest {
	return ApiGetExtensionDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return ExtensionDefinition
func (a *ExtensionDefinitionV1alpha1APIService) GetExtensionDefinitionExecute(r ApiGetExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtensionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionDefinitionV1alpha1APIService.GetExtensionDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name}"
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

type ApiListExtensionDefinitionRequest struct {
	ctx context.Context
	ApiService ExtensionDefinitionV1alpha1API
	page *int32
	size *int32
	labelSelector *[]string
	fieldSelector *[]string
	sort *[]string
}

// Page number. Default is 0.
func (r ApiListExtensionDefinitionRequest) Page(page int32) ApiListExtensionDefinitionRequest {
	r.page = &page
	return r
}

// Size number. Default is 0.
func (r ApiListExtensionDefinitionRequest) Size(size int32) ApiListExtensionDefinitionRequest {
	r.size = &size
	return r
}

// Label selector. e.g.: hidden!&#x3D;true
func (r ApiListExtensionDefinitionRequest) LabelSelector(labelSelector []string) ApiListExtensionDefinitionRequest {
	r.labelSelector = &labelSelector
	return r
}

// Field selector. e.g.: metadata.name&#x3D;&#x3D;halo
func (r ApiListExtensionDefinitionRequest) FieldSelector(fieldSelector []string) ApiListExtensionDefinitionRequest {
	r.fieldSelector = &fieldSelector
	return r
}

// Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
func (r ApiListExtensionDefinitionRequest) Sort(sort []string) ApiListExtensionDefinitionRequest {
	r.sort = &sort
	return r
}

func (r ApiListExtensionDefinitionRequest) Execute() (*ExtensionDefinitionList, *http.Response, error) {
	return r.ApiService.ListExtensionDefinitionExecute(r)
}

/*
ListExtensionDefinition Method for ListExtensionDefinition

List ExtensionDefinition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListExtensionDefinitionRequest
*/
func (a *ExtensionDefinitionV1alpha1APIService) ListExtensionDefinition(ctx context.Context) ApiListExtensionDefinitionRequest {
	return ApiListExtensionDefinitionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExtensionDefinitionList
func (a *ExtensionDefinitionV1alpha1APIService) ListExtensionDefinitionExecute(r ApiListExtensionDefinitionRequest) (*ExtensionDefinitionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtensionDefinitionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionDefinitionV1alpha1APIService.ListExtensionDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/plugin.halo.run/v1alpha1/extensiondefinitions"

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

type ApiPatchExtensionDefinitionRequest struct {
	ctx context.Context
	ApiService ExtensionDefinitionV1alpha1API
	name string
	jsonPatchInner *[]JsonPatchInner
}

func (r ApiPatchExtensionDefinitionRequest) JsonPatchInner(jsonPatchInner []JsonPatchInner) ApiPatchExtensionDefinitionRequest {
	r.jsonPatchInner = &jsonPatchInner
	return r
}

func (r ApiPatchExtensionDefinitionRequest) Execute() (*ExtensionDefinition, *http.Response, error) {
	return r.ApiService.PatchExtensionDefinitionExecute(r)
}

/*
PatchExtensionDefinition Method for PatchExtensionDefinition

Patch ExtensionDefinition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of extensiondefinition
 @return ApiPatchExtensionDefinitionRequest
*/
func (a *ExtensionDefinitionV1alpha1APIService) PatchExtensionDefinition(ctx context.Context, name string) ApiPatchExtensionDefinitionRequest {
	return ApiPatchExtensionDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return ExtensionDefinition
func (a *ExtensionDefinitionV1alpha1APIService) PatchExtensionDefinitionExecute(r ApiPatchExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtensionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionDefinitionV1alpha1APIService.PatchExtensionDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name}"
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

type ApiUpdateExtensionDefinitionRequest struct {
	ctx context.Context
	ApiService ExtensionDefinitionV1alpha1API
	name string
	extensionDefinition *ExtensionDefinition
}

// Updated extensiondefinition
func (r ApiUpdateExtensionDefinitionRequest) ExtensionDefinition(extensionDefinition ExtensionDefinition) ApiUpdateExtensionDefinitionRequest {
	r.extensionDefinition = &extensionDefinition
	return r
}

func (r ApiUpdateExtensionDefinitionRequest) Execute() (*ExtensionDefinition, *http.Response, error) {
	return r.ApiService.UpdateExtensionDefinitionExecute(r)
}

/*
UpdateExtensionDefinition Method for UpdateExtensionDefinition

Update ExtensionDefinition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of extensiondefinition
 @return ApiUpdateExtensionDefinitionRequest
*/
func (a *ExtensionDefinitionV1alpha1APIService) UpdateExtensionDefinition(ctx context.Context, name string) ApiUpdateExtensionDefinitionRequest {
	return ApiUpdateExtensionDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return ExtensionDefinition
func (a *ExtensionDefinitionV1alpha1APIService) UpdateExtensionDefinitionExecute(r ApiUpdateExtensionDefinitionRequest) (*ExtensionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtensionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionDefinitionV1alpha1APIService.UpdateExtensionDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/plugin.halo.run/v1alpha1/extensiondefinitions/{name}"
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
	localVarPostBody = r.extensionDefinition
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
