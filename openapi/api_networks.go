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


// NetworksApiService NetworksApi service
type NetworksApiService service

type ApiGetNetworkRequest struct {
	ctx context.Context
	ApiService *NetworksApiService
	networkId string
	fields *[]string
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiGetNetworkRequest) Fields(fields []string) ApiGetNetworkRequest {
	r.fields = &fields
	return r
}

func (r ApiGetNetworkRequest) Execute() (*PersonNetworkEntry, *http.Response, error) {
	return r.ApiService.GetNetworkExecute(r)
}

/*
GetNetwork Get a network

Gets information for a network **networkId**.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId The identifier of a network.
 @return ApiGetNetworkRequest
*/
func (a *NetworksApiService) GetNetwork(ctx context.Context, networkId string) ApiGetNetworkRequest {
	return ApiGetNetworkRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return PersonNetworkEntry
func (a *NetworksApiService) GetNetworkExecute(r ApiGetNetworkRequest) (*PersonNetworkEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PersonNetworkEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksApiService.GetNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetNetworkForPersonRequest struct {
	ctx context.Context
	ApiService *NetworksApiService
	personId string
	networkId string
	fields *[]string
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiGetNetworkForPersonRequest) Fields(fields []string) ApiGetNetworkForPersonRequest {
	r.fields = &fields
	return r
}

func (r ApiGetNetworkForPersonRequest) Execute() (*PersonNetworkEntry, *http.Response, error) {
	return r.ApiService.GetNetworkForPersonExecute(r)
}

/*
GetNetworkForPerson Get network information

Gets network information on a single network specified by **networkId** for **personId**.

You can use the `-me-` string in place of `<personId>` to specify the currently authenticated user.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param personId The identifier of a person.
 @param networkId The identifier of a network.
 @return ApiGetNetworkForPersonRequest
*/
func (a *NetworksApiService) GetNetworkForPerson(ctx context.Context, personId string, networkId string) ApiGetNetworkForPersonRequest {
	return ApiGetNetworkForPersonRequest{
		ApiService: a,
		ctx: ctx,
		personId: personId,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return PersonNetworkEntry
func (a *NetworksApiService) GetNetworkForPersonExecute(r ApiGetNetworkForPersonRequest) (*PersonNetworkEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PersonNetworkEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksApiService.GetNetworkForPerson")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/people/{personId}/networks/{networkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListNetworksForPersonRequest struct {
	ctx context.Context
	ApiService *NetworksApiService
	personId string
	skipCount *int32
	maxItems *int32
	fields *[]string
}

// The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0. 
func (r ApiListNetworksForPersonRequest) SkipCount(skipCount int32) ApiListNetworksForPersonRequest {
	r.skipCount = &skipCount
	return r
}

// The maximum number of items to return in the list. If not supplied then the default value is 100. 
func (r ApiListNetworksForPersonRequest) MaxItems(maxItems int32) ApiListNetworksForPersonRequest {
	r.maxItems = &maxItems
	return r
}

// A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter. 
func (r ApiListNetworksForPersonRequest) Fields(fields []string) ApiListNetworksForPersonRequest {
	r.fields = &fields
	return r
}

func (r ApiListNetworksForPersonRequest) Execute() (*PersonNetworkPaging, *http.Response, error) {
	return r.ApiService.ListNetworksForPersonExecute(r)
}

/*
ListNetworksForPerson List network membership

Gets a list of network memberships for person **personId**.

You can use the `-me-` string in place of `<personId>` to specify the currently authenticated user.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param personId The identifier of a person.
 @return ApiListNetworksForPersonRequest
*/
func (a *NetworksApiService) ListNetworksForPerson(ctx context.Context, personId string) ApiListNetworksForPersonRequest {
	return ApiListNetworksForPersonRequest{
		ApiService: a,
		ctx: ctx,
		personId: personId,
	}
}

// Execute executes the request
//  @return PersonNetworkPaging
func (a *NetworksApiService) ListNetworksForPersonExecute(r ApiListNetworksForPersonRequest) (*PersonNetworkPaging, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PersonNetworkPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksApiService.ListNetworksForPerson")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/people/{personId}/networks"
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

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
