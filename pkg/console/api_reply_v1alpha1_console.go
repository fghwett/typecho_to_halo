/*
Halo

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.20.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package console_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


type ReplyV1alpha1ConsoleAPI interface {

	/*
	ListReplies Method for ListReplies

	List replies.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListRepliesRequest
	*/
	ListReplies(ctx context.Context) ApiListRepliesRequest

	// ListRepliesExecute executes the request
	//  @return ListedReplyList
	ListRepliesExecute(r ApiListRepliesRequest) (*ListedReplyList, *http.Response, error)
}

// ReplyV1alpha1ConsoleAPIService ReplyV1alpha1ConsoleAPI service
type ReplyV1alpha1ConsoleAPIService service

type ApiListRepliesRequest struct {
	ctx context.Context
	ApiService ReplyV1alpha1ConsoleAPI
	commentName *string
	page *int32
	size *int32
	labelSelector *[]string
	fieldSelector *[]string
	sort *[]string
}

// Replies filtered by commentName.
func (r ApiListRepliesRequest) CommentName(commentName string) ApiListRepliesRequest {
	r.commentName = &commentName
	return r
}

// Page number. Default is 0.
func (r ApiListRepliesRequest) Page(page int32) ApiListRepliesRequest {
	r.page = &page
	return r
}

// Size number. Default is 0.
func (r ApiListRepliesRequest) Size(size int32) ApiListRepliesRequest {
	r.size = &size
	return r
}

// Label selector. e.g.: hidden!&#x3D;true
func (r ApiListRepliesRequest) LabelSelector(labelSelector []string) ApiListRepliesRequest {
	r.labelSelector = &labelSelector
	return r
}

// Field selector. e.g.: metadata.name&#x3D;&#x3D;halo
func (r ApiListRepliesRequest) FieldSelector(fieldSelector []string) ApiListRepliesRequest {
	r.fieldSelector = &fieldSelector
	return r
}

// Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
func (r ApiListRepliesRequest) Sort(sort []string) ApiListRepliesRequest {
	r.sort = &sort
	return r
}

func (r ApiListRepliesRequest) Execute() (*ListedReplyList, *http.Response, error) {
	return r.ApiService.ListRepliesExecute(r)
}

/*
ListReplies Method for ListReplies

List replies.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListRepliesRequest
*/
func (a *ReplyV1alpha1ConsoleAPIService) ListReplies(ctx context.Context) ApiListRepliesRequest {
	return ApiListRepliesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListedReplyList
func (a *ReplyV1alpha1ConsoleAPIService) ListRepliesExecute(r ApiListRepliesRequest) (*ListedReplyList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListedReplyList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReplyV1alpha1ConsoleAPIService.ListReplies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apis/api.console.halo.run/v1alpha1/replies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commentName == nil {
		return localVarReturnValue, nil, reportError("commentName is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarQueryParams, "commentName", r.commentName, "form", "")
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
			var v ListedReplyList
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
