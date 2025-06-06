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


type NotificationV1alpha1API interface {

	/*
	CreateNotification Method for CreateNotification

	Create Notification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateNotificationRequest
	*/
	CreateNotification(ctx context.Context) ApiCreateNotificationRequest

	// CreateNotificationExecute executes the request
	//  @return Notification
	CreateNotificationExecute(r ApiCreateNotificationRequest) (*Notification, *http.Response, error)

	/*
	DeleteNotification Method for DeleteNotification

	Delete Notification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of notification
	@return ApiDeleteNotificationRequest
	*/
	DeleteNotification(ctx context.Context, name string) ApiDeleteNotificationRequest

	// DeleteNotificationExecute executes the request
	DeleteNotificationExecute(r ApiDeleteNotificationRequest) (*http.Response, error)

	/*
	GetNotification Method for GetNotification

	Get Notification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of notification
	@return ApiGetNotificationRequest
	*/
	GetNotification(ctx context.Context, name string) ApiGetNotificationRequest

	// GetNotificationExecute executes the request
	//  @return Notification
	GetNotificationExecute(r ApiGetNotificationRequest) (*Notification, *http.Response, error)

	/*
	ListNotification Method for ListNotification

	List Notification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListNotificationRequest
	*/
	ListNotification(ctx context.Context) ApiListNotificationRequest

	// ListNotificationExecute executes the request
	//  @return NotificationList
	ListNotificationExecute(r ApiListNotificationRequest) (*NotificationList, *http.Response, error)

	/*
	PatchNotification Method for PatchNotification

	Patch Notification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of notification
	@return ApiPatchNotificationRequest
	*/
	PatchNotification(ctx context.Context, name string) ApiPatchNotificationRequest

	// PatchNotificationExecute executes the request
	//  @return Notification
	PatchNotificationExecute(r ApiPatchNotificationRequest) (*Notification, *http.Response, error)

	/*
	UpdateNotification Method for UpdateNotification

	Update Notification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of notification
	@return ApiUpdateNotificationRequest
	*/
	UpdateNotification(ctx context.Context, name string) ApiUpdateNotificationRequest

	// UpdateNotificationExecute executes the request
	//  @return Notification
	UpdateNotificationExecute(r ApiUpdateNotificationRequest) (*Notification, *http.Response, error)
}

// NotificationV1alpha1APIService NotificationV1alpha1API service
type NotificationV1alpha1APIService service

type ApiCreateNotificationRequest struct {
	ctx context.Context
	ApiService NotificationV1alpha1API
	notification *Notification
}

// Fresh notification
func (r ApiCreateNotificationRequest) Notification(notification Notification) ApiCreateNotificationRequest {
	r.notification = &notification
	return r
}

func (r ApiCreateNotificationRequest) Execute() (*Notification, *http.Response, error) {
	return r.ApiService.CreateNotificationExecute(r)
}

/*
CreateNotification Method for CreateNotification

Create Notification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNotificationRequest
*/
func (a *NotificationV1alpha1APIService) CreateNotification(ctx context.Context) ApiCreateNotificationRequest {
	return ApiCreateNotificationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Notification
func (a *NotificationV1alpha1APIService) CreateNotificationExecute(r ApiCreateNotificationRequest) (*Notification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Notification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationV1alpha1APIService.CreateNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/notification.halo.run/v1alpha1/notifications"

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
	localVarPostBody = r.notification
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

type ApiDeleteNotificationRequest struct {
	ctx context.Context
	ApiService NotificationV1alpha1API
	name string
}

func (r ApiDeleteNotificationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNotificationExecute(r)
}

/*
DeleteNotification Method for DeleteNotification

Delete Notification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of notification
 @return ApiDeleteNotificationRequest
*/
func (a *NotificationV1alpha1APIService) DeleteNotification(ctx context.Context, name string) ApiDeleteNotificationRequest {
	return ApiDeleteNotificationRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
func (a *NotificationV1alpha1APIService) DeleteNotificationExecute(r ApiDeleteNotificationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationV1alpha1APIService.DeleteNotification")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/notification.halo.run/v1alpha1/notifications/{name}"
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

type ApiGetNotificationRequest struct {
	ctx context.Context
	ApiService NotificationV1alpha1API
	name string
}

func (r ApiGetNotificationRequest) Execute() (*Notification, *http.Response, error) {
	return r.ApiService.GetNotificationExecute(r)
}

/*
GetNotification Method for GetNotification

Get Notification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of notification
 @return ApiGetNotificationRequest
*/
func (a *NotificationV1alpha1APIService) GetNotification(ctx context.Context, name string) ApiGetNotificationRequest {
	return ApiGetNotificationRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return Notification
func (a *NotificationV1alpha1APIService) GetNotificationExecute(r ApiGetNotificationRequest) (*Notification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Notification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationV1alpha1APIService.GetNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/notification.halo.run/v1alpha1/notifications/{name}"
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

type ApiListNotificationRequest struct {
	ctx context.Context
	ApiService NotificationV1alpha1API
	page *int32
	size *int32
	labelSelector *[]string
	fieldSelector *[]string
	sort *[]string
}

// Page number. Default is 0.
func (r ApiListNotificationRequest) Page(page int32) ApiListNotificationRequest {
	r.page = &page
	return r
}

// Size number. Default is 0.
func (r ApiListNotificationRequest) Size(size int32) ApiListNotificationRequest {
	r.size = &size
	return r
}

// Label selector. e.g.: hidden!&#x3D;true
func (r ApiListNotificationRequest) LabelSelector(labelSelector []string) ApiListNotificationRequest {
	r.labelSelector = &labelSelector
	return r
}

// Field selector. e.g.: metadata.name&#x3D;&#x3D;halo
func (r ApiListNotificationRequest) FieldSelector(fieldSelector []string) ApiListNotificationRequest {
	r.fieldSelector = &fieldSelector
	return r
}

// Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
func (r ApiListNotificationRequest) Sort(sort []string) ApiListNotificationRequest {
	r.sort = &sort
	return r
}

func (r ApiListNotificationRequest) Execute() (*NotificationList, *http.Response, error) {
	return r.ApiService.ListNotificationExecute(r)
}

/*
ListNotification Method for ListNotification

List Notification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNotificationRequest
*/
func (a *NotificationV1alpha1APIService) ListNotification(ctx context.Context) ApiListNotificationRequest {
	return ApiListNotificationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NotificationList
func (a *NotificationV1alpha1APIService) ListNotificationExecute(r ApiListNotificationRequest) (*NotificationList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotificationList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationV1alpha1APIService.ListNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/notification.halo.run/v1alpha1/notifications"

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

type ApiPatchNotificationRequest struct {
	ctx context.Context
	ApiService NotificationV1alpha1API
	name string
	jsonPatchInner *[]JsonPatchInner
}

func (r ApiPatchNotificationRequest) JsonPatchInner(jsonPatchInner []JsonPatchInner) ApiPatchNotificationRequest {
	r.jsonPatchInner = &jsonPatchInner
	return r
}

func (r ApiPatchNotificationRequest) Execute() (*Notification, *http.Response, error) {
	return r.ApiService.PatchNotificationExecute(r)
}

/*
PatchNotification Method for PatchNotification

Patch Notification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of notification
 @return ApiPatchNotificationRequest
*/
func (a *NotificationV1alpha1APIService) PatchNotification(ctx context.Context, name string) ApiPatchNotificationRequest {
	return ApiPatchNotificationRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return Notification
func (a *NotificationV1alpha1APIService) PatchNotificationExecute(r ApiPatchNotificationRequest) (*Notification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Notification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationV1alpha1APIService.PatchNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/notification.halo.run/v1alpha1/notifications/{name}"
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

type ApiUpdateNotificationRequest struct {
	ctx context.Context
	ApiService NotificationV1alpha1API
	name string
	notification *Notification
}

// Updated notification
func (r ApiUpdateNotificationRequest) Notification(notification Notification) ApiUpdateNotificationRequest {
	r.notification = &notification
	return r
}

func (r ApiUpdateNotificationRequest) Execute() (*Notification, *http.Response, error) {
	return r.ApiService.UpdateNotificationExecute(r)
}

/*
UpdateNotification Method for UpdateNotification

Update Notification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Name of notification
 @return ApiUpdateNotificationRequest
*/
func (a *NotificationV1alpha1APIService) UpdateNotification(ctx context.Context, name string) ApiUpdateNotificationRequest {
	return ApiUpdateNotificationRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return Notification
func (a *NotificationV1alpha1APIService) UpdateNotificationExecute(r ApiUpdateNotificationRequest) (*Notification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Notification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationV1alpha1APIService.UpdateNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/notification.halo.run/v1alpha1/notifications/{name}"
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
	localVarPostBody = r.notification
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
