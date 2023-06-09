/*
Alfresco Content Services REST API

**Core API**  Provides access to the core features of Alfresco Content Services. 

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CommentsApiService CommentsApi service
type CommentsApiService service

type ApiCreateCommentRequest struct {
	ctx context.Context
	ApiService *CommentsApiService
	nodeId string
	commentBodyCreate *CommentBody
	fields *[]string
}

// The comment text. Note that you can also provide a list of comments.
func (r ApiCreateCommentRequest) CommentBodyCreate(commentBodyCreate CommentBody) ApiCreateCommentRequest {
	r.commentBodyCreate = &commentBodyCreate
	return r
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiCreateCommentRequest) Fields(fields []string) ApiCreateCommentRequest {
	r.fields = &fields
	return r
}

func (r ApiCreateCommentRequest) Execute() (*CommentEntry, *http.Response, error) {
	return r.ApiService.CreateCommentExecute(r)
}

/*
CreateComment Create a comment

Creates a comment on node **nodeId**. You specify the comment in a JSON body like this:

```JSON
{
  "content": "This is a comment"
}
```

**Note:** You can create more than one comment by
specifying a list of comments in the JSON body like this:

```JSON
[
  {
    "content": "This is a comment"
  },
  {
    "content": "This is another comment"
  }
]
```
If you specify a list as input, then a paginated list rather than an entry is returned in the response body. For example:

```JSON
{
  "list": {
    "pagination": {
      "count": 2,
      "hasMoreItems": false,
      "totalItems": 2,
      "skipCount": 0,
      "maxItems": 100
    },
    "entries": [
      {
        "entry": {
          ...
        }
      },
      {
        "entry": {
          ...
        }
      }
    ]
  }
}
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId The identifier of a node.
 @return ApiCreateCommentRequest
*/
func (a *CommentsApiService) CreateComment(ctx context.Context, nodeId string) ApiCreateCommentRequest {
	return ApiCreateCommentRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return CommentEntry
func (a *CommentsApiService) CreateCommentExecute(r ApiCreateCommentRequest) (*CommentEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommentEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.CreateComment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}/comments"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commentBodyCreate == nil {
		return localVarReturnValue, nil, reportError("commentBodyCreate is required and must be specified")
	}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.commentBodyCreate
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
			var v Error
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

type ApiDeleteCommentRequest struct {
	ctx context.Context
	ApiService *CommentsApiService
	nodeId string
	commentId string
}

func (r ApiDeleteCommentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCommentExecute(r)
}

/*
DeleteComment Delete a comment

Deletes the comment **commentId** from node **nodeId**.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId The identifier of a node.
 @param commentId The identifier of a comment.
 @return ApiDeleteCommentRequest
*/
func (a *CommentsApiService) DeleteComment(ctx context.Context, nodeId string, commentId string) ApiDeleteCommentRequest {
	return ApiDeleteCommentRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
		commentId: commentId,
	}
}

// Execute executes the request
func (a *CommentsApiService) DeleteCommentExecute(r ApiDeleteCommentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.DeleteComment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}/comments/{commentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListCommentsRequest struct {
	ctx context.Context
	ApiService *CommentsApiService
	nodeId string
	skipCount *int32
	maxItems *int32
	fields *[]string
}

// The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0. 
func (r ApiListCommentsRequest) SkipCount(skipCount int32) ApiListCommentsRequest {
	r.skipCount = &skipCount
	return r
}

// The maximum number of items to return in the list. If not supplied then the default value is 100. 
func (r ApiListCommentsRequest) MaxItems(maxItems int32) ApiListCommentsRequest {
	r.maxItems = &maxItems
	return r
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiListCommentsRequest) Fields(fields []string) ApiListCommentsRequest {
	r.fields = &fields
	return r
}

func (r ApiListCommentsRequest) Execute() (*CommentPaging, *http.Response, error) {
	return r.ApiService.ListCommentsExecute(r)
}

/*
ListComments List comments

Gets a list of comments for the node **nodeId**, sorted chronologically with the newest comment first.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId The identifier of a node.
 @return ApiListCommentsRequest
*/
func (a *CommentsApiService) ListComments(ctx context.Context, nodeId string) ApiListCommentsRequest {
	return ApiListCommentsRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return CommentPaging
func (a *CommentsApiService) ListCommentsExecute(r ApiListCommentsRequest) (*CommentPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommentPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.ListComments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}/comments"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skipCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipCount", r.skipCount, "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxItems", r.maxItems, "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
			var v Error
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

type ApiUpdateCommentRequest struct {
	ctx context.Context
	ApiService *CommentsApiService
	nodeId string
	commentId string
	commentBodyUpdate *CommentBody
	fields *[]string
}

// The JSON representing the comment to be updated.
func (r ApiUpdateCommentRequest) CommentBodyUpdate(commentBodyUpdate CommentBody) ApiUpdateCommentRequest {
	r.commentBodyUpdate = &commentBodyUpdate
	return r
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiUpdateCommentRequest) Fields(fields []string) ApiUpdateCommentRequest {
	r.fields = &fields
	return r
}

func (r ApiUpdateCommentRequest) Execute() (*CommentEntry, *http.Response, error) {
	return r.ApiService.UpdateCommentExecute(r)
}

/*
UpdateComment Update a comment

Updates an existing comment **commentId** on node **nodeId**.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId The identifier of a node.
 @param commentId The identifier of a comment.
 @return ApiUpdateCommentRequest
*/
func (a *CommentsApiService) UpdateComment(ctx context.Context, nodeId string, commentId string) ApiUpdateCommentRequest {
	return ApiUpdateCommentRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
		commentId: commentId,
	}
}

// Execute executes the request
//  @return CommentEntry
func (a *CommentsApiService) UpdateCommentExecute(r ApiUpdateCommentRequest) (*CommentEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommentEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.UpdateComment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}/comments/{commentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", url.PathEscape(parameterValueToString(r.nodeId, "nodeId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commentBodyUpdate == nil {
		return localVarReturnValue, nil, reportError("commentBodyUpdate is required and must be specified")
	}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.commentBodyUpdate
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
			var v Error
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
