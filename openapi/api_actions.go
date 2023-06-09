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


// ActionsApiService ActionsApi service
type ActionsApiService service

type ApiActionDetailsRequest struct {
	ctx context.Context
	ApiService *ActionsApiService
	actionDefinitionId string
}

func (r ApiActionDetailsRequest) Execute() (*ActionDefinitionEntry, *http.Response, error) {
	return r.ApiService.ActionDetailsExecute(r)
}

/*
ActionDetails Retrieve the details of an action definition

**Note:** this endpoint is available in Alfresco 5.2 and newer versions.

Retrieve the details of the action denoted by **actionDefinitionId**.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actionDefinitionId The identifier of an action definition.
 @return ApiActionDetailsRequest
*/
func (a *ActionsApiService) ActionDetails(ctx context.Context, actionDefinitionId string) ApiActionDetailsRequest {
	return ApiActionDetailsRequest{
		ApiService: a,
		ctx: ctx,
		actionDefinitionId: actionDefinitionId,
	}
}

// Execute executes the request
//  @return ActionDefinitionEntry
func (a *ActionsApiService) ActionDetailsExecute(r ApiActionDetailsRequest) (*ActionDefinitionEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActionDefinitionEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.ActionDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/action-definitions/{actionDefinitionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actionDefinitionId"+"}", url.PathEscape(parameterValueToString(r.actionDefinitionId, "actionDefinitionId")), -1)

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

type ApiActionExecRequest struct {
	ctx context.Context
	ApiService *ActionsApiService
	actionBodyExec *ActionBodyExec
}

// Action execution details
func (r ApiActionExecRequest) ActionBodyExec(actionBodyExec ActionBodyExec) ApiActionExecRequest {
	r.actionBodyExec = &actionBodyExec
	return r
}

func (r ApiActionExecRequest) Execute() (*ActionExecResultEntry, *http.Response, error) {
	return r.ApiService.ActionExecExecute(r)
}

/*
ActionExec Execute an action

**Note:** this endpoint is available in Alfresco 5.2 and newer versions.

Executes an action

An action may be executed against a node specified by **targetId**. For example:

```
{
  "actionDefinitionId": "copy",
  "targetId": "4c4b3c43-f18b-43ff-af84-751f16f1ddfd",
  "params": {
  	"destination-folder": "34219f79-66fa-4ebf-b371-118598af898c"
  }
}
```

Performing a POST with the request body shown above will result in the node identified by ```targetId```
being copied to the destination folder specified in the ```params``` object by the key ```destination-folder```.

**targetId** is optional, however, currently **targetId** must be a valid node ID.
In the future, actions may be executed against different entity types or
executed without the need for the context of an entity.


Parameters supplied to the action within the ```params``` object will be converted to the expected type,
where possible using the DefaultTypeConverter class. In addition:

* Node IDs may be supplied in their short form (implicit workspace://SpacesStore prefix)
* Aspect names may be supplied using their short form, e.g. cm:versionable or cm:auditable

In this example, we add the aspect ```cm:versionable``` to a node using the QName resolution mentioned above:

```
{
  "actionDefinitionId": "add-features",
  "targetId": "16349e3f-2977-44d1-93f2-73c12b8083b5",
  "params": {
  	"aspect-name": "cm:versionable"
  }
}
```

The ```actionDefinitionId``` is the ```id``` of an action definition as returned by
the _list actions_ operations (e.g. GET /action-definitions).

The action will be executed **asynchronously** with a `202` HTTP response signifying that
the request has been accepted successfully. The response body contains the unique ID of the action
pending execution. The ID may be used, for example to correlate an execution with output in the server logs.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiActionExecRequest
*/
func (a *ActionsApiService) ActionExec(ctx context.Context) ApiActionExecRequest {
	return ApiActionExecRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActionExecResultEntry
func (a *ActionsApiService) ActionExecExecute(r ApiActionExecRequest) (*ActionExecResultEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActionExecResultEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.ActionExec")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/action-executions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.actionBodyExec == nil {
		return localVarReturnValue, nil, reportError("actionBodyExec is required and must be specified")
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
	localVarPostBody = r.actionBodyExec
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

type ApiListActionsRequest struct {
	ctx context.Context
	ApiService *ActionsApiService
	skipCount *int32
	maxItems *int32
	orderBy *[]string
	fields *[]string
}

// The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0. 
func (r ApiListActionsRequest) SkipCount(skipCount int32) ApiListActionsRequest {
	r.skipCount = &skipCount
	return r
}

// The maximum number of items to return in the list. If not supplied then the default value is 100. 
func (r ApiListActionsRequest) MaxItems(maxItems int32) ApiListActionsRequest {
	r.maxItems = &maxItems
	return r
}

// A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field. 
func (r ApiListActionsRequest) OrderBy(orderBy []string) ApiListActionsRequest {
	r.orderBy = &orderBy
	return r
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiListActionsRequest) Fields(fields []string) ApiListActionsRequest {
	r.fields = &fields
	return r
}

func (r ApiListActionsRequest) Execute() (*ActionDefinitionList, *http.Response, error) {
	return r.ApiService.ListActionsExecute(r)
}

/*
ListActions Retrieve list of available actions

**Note:** this endpoint is available in Alfresco 5.2.2 and newer versions.

Gets a list of all available actions

The default sort order for the returned list is for actions to be sorted by ascending name.
You can override the default by using the **orderBy** parameter.

You can use any of the following fields to order the results:
* name
* title


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListActionsRequest
*/
func (a *ActionsApiService) ListActions(ctx context.Context) ApiListActionsRequest {
	return ApiListActionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActionDefinitionList
func (a *ActionsApiService) ListActionsExecute(r ApiListActionsRequest) (*ActionDefinitionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActionDefinitionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.ListActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/action-definitions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skipCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipCount", r.skipCount, "")
	}
	if r.maxItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxItems", r.maxItems, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "csv")
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

type ApiNodeActionsRequest struct {
	ctx context.Context
	ApiService *ActionsApiService
	nodeId string
	skipCount *int32
	maxItems *int32
	orderBy *[]string
	fields *[]string
}

// The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0. 
func (r ApiNodeActionsRequest) SkipCount(skipCount int32) ApiNodeActionsRequest {
	r.skipCount = &skipCount
	return r
}

// The maximum number of items to return in the list. If not supplied then the default value is 100. 
func (r ApiNodeActionsRequest) MaxItems(maxItems int32) ApiNodeActionsRequest {
	r.maxItems = &maxItems
	return r
}

// A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field. 
func (r ApiNodeActionsRequest) OrderBy(orderBy []string) ApiNodeActionsRequest {
	r.orderBy = &orderBy
	return r
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiNodeActionsRequest) Fields(fields []string) ApiNodeActionsRequest {
	r.fields = &fields
	return r
}

func (r ApiNodeActionsRequest) Execute() (*ActionDefinitionList, *http.Response, error) {
	return r.ApiService.NodeActionsExecute(r)
}

/*
NodeActions Retrieve actions for a node

**Note:** this endpoint is available in Alfresco 5.2 and newer versions.

Retrieve the list of actions that may be executed against the given **nodeId**.

The default sort order for the returned list is for actions to be sorted by ascending name.
You can override the default by using the **orderBy** parameter.

You can use any of the following fields to order the results:
* name
* title


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeId The identifier of a node.
 @return ApiNodeActionsRequest
*/
func (a *ActionsApiService) NodeActions(ctx context.Context, nodeId string) ApiNodeActionsRequest {
	return ApiNodeActionsRequest{
		ApiService: a,
		ctx: ctx,
		nodeId: nodeId,
	}
}

// Execute executes the request
//  @return ActionDefinitionList
func (a *ActionsApiService) NodeActionsExecute(r ApiNodeActionsRequest) (*ActionDefinitionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActionDefinitionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActionsApiService.NodeActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeId}/action-definitions"
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
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "csv")
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
