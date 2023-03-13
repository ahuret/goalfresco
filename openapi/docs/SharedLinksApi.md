# \SharedLinksApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedLink**](SharedLinksApi.md#CreateSharedLink) | **Post** /shared-links | Create a shared link to a file
[**DeleteSharedLink**](SharedLinksApi.md#DeleteSharedLink) | **Delete** /shared-links/{sharedId} | Deletes a shared link
[**EmailSharedLink**](SharedLinksApi.md#EmailSharedLink) | **Post** /shared-links/{sharedId}/email | Email shared link
[**GetSharedLink**](SharedLinksApi.md#GetSharedLink) | **Get** /shared-links/{sharedId} | Get a shared link
[**GetSharedLinkContent**](SharedLinksApi.md#GetSharedLinkContent) | **Get** /shared-links/{sharedId}/content | Get shared link content
[**GetSharedLinkRendition**](SharedLinksApi.md#GetSharedLinkRendition) | **Get** /shared-links/{sharedId}/renditions/{renditionId} | Get shared link rendition information
[**GetSharedLinkRenditionContent**](SharedLinksApi.md#GetSharedLinkRenditionContent) | **Get** /shared-links/{sharedId}/renditions/{renditionId}/content | Get shared link rendition content
[**ListSharedLinkRenditions**](SharedLinksApi.md#ListSharedLinkRenditions) | **Get** /shared-links/{sharedId}/renditions | List renditions for a shared link
[**ListSharedLinks**](SharedLinksApi.md#ListSharedLinks) | **Get** /shared-links | List shared links



## CreateSharedLink

> SharedLinkEntry CreateSharedLink(ctx).SharedLinkBodyCreate(sharedLinkBodyCreate).Include(include).Fields(fields).Execute()

Create a shared link to a file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedLinkBodyCreate := *openapiclient.NewSharedLinkBodyCreate("NodeId_example") // SharedLinkBodyCreate | The nodeId to create a shared link for.
    include := []string{"Inner_example"} // []string | Returns additional information about the shared link, the following optional fields can be requested: * allowableOperations * path * properties * isFavorite * aspectNames  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLinksApi.CreateSharedLink(context.Background()).SharedLinkBodyCreate(sharedLinkBodyCreate).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.CreateSharedLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSharedLink`: SharedLinkEntry
    fmt.Fprintf(os.Stdout, "Response from `SharedLinksApi.CreateSharedLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSharedLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedLinkBodyCreate** | [**SharedLinkBodyCreate**](SharedLinkBodyCreate.md) | The nodeId to create a shared link for. | 
 **include** | **[]string** | Returns additional information about the shared link, the following optional fields can be requested: * allowableOperations * path * properties * isFavorite * aspectNames  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SharedLinkEntry**](SharedLinkEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSharedLink

> DeleteSharedLink(ctx, sharedId).Execute()

Deletes a shared link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedId := "sharedId_example" // string | The identifier of a shared link to a file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SharedLinksApi.DeleteSharedLink(context.Background(), sharedId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.DeleteSharedLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedId** | **string** | The identifier of a shared link to a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailSharedLink

> EmailSharedLink(ctx, sharedId).SharedLinkBodyEmail(sharedLinkBodyEmail).Execute()

Email shared link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedId := "sharedId_example" // string | The identifier of a shared link to a file.
    sharedLinkBodyEmail := *openapiclient.NewSharedLinkBodyEmail() // SharedLinkBodyEmail | The shared link email to send.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SharedLinksApi.EmailSharedLink(context.Background(), sharedId).SharedLinkBodyEmail(sharedLinkBodyEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.EmailSharedLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedId** | **string** | The identifier of a shared link to a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailSharedLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharedLinkBodyEmail** | [**SharedLinkBodyEmail**](SharedLinkBodyEmail.md) | The shared link email to send. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedLink

> SharedLinkEntry GetSharedLink(ctx, sharedId).Fields(fields).Execute()

Get a shared link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedId := "sharedId_example" // string | The identifier of a shared link to a file.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLinksApi.GetSharedLink(context.Background(), sharedId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.GetSharedLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSharedLink`: SharedLinkEntry
    fmt.Fprintf(os.Stdout, "Response from `SharedLinksApi.GetSharedLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedId** | **string** | The identifier of a shared link to a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SharedLinkEntry**](SharedLinkEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedLinkContent

> *os.File GetSharedLinkContent(ctx, sharedId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()

Get shared link content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/ahuret/goalfresco/time"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedId := "sharedId_example" // string | The identifier of a shared link to a file.
    attachment := true // bool | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  (optional) (default to true)
    ifModifiedSince := time.Now() // time.Time | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, `Wed, 09 Mar 2016 16:56:34 GMT`.  (optional)
    range_ := "range__example" // string | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes=1-10.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLinksApi.GetSharedLinkContent(context.Background(), sharedId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.GetSharedLinkContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSharedLinkContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SharedLinksApi.GetSharedLinkContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedId** | **string** | The identifier of a shared link to a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedLinkContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachment** | **bool** | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  | [default to true]
 **ifModifiedSince** | **time.Time** | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, &#x60;Wed, 09 Mar 2016 16:56:34 GMT&#x60;.  | 
 **range_** | **string** | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes&#x3D;1-10.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedLinkRendition

> RenditionEntry GetSharedLinkRendition(ctx, sharedId, renditionId).Execute()

Get shared link rendition information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedId := "sharedId_example" // string | The identifier of a shared link to a file.
    renditionId := "renditionId_example" // string | The name of a thumbnail rendition, for example *doclib*, or *pdf*.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLinksApi.GetSharedLinkRendition(context.Background(), sharedId, renditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.GetSharedLinkRendition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSharedLinkRendition`: RenditionEntry
    fmt.Fprintf(os.Stdout, "Response from `SharedLinksApi.GetSharedLinkRendition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedId** | **string** | The identifier of a shared link to a file. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedLinkRenditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RenditionEntry**](RenditionEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedLinkRenditionContent

> *os.File GetSharedLinkRenditionContent(ctx, sharedId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()

Get shared link rendition content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/ahuret/goalfresco/time"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedId := "sharedId_example" // string | The identifier of a shared link to a file.
    renditionId := "renditionId_example" // string | The name of a thumbnail rendition, for example *doclib*, or *pdf*.
    attachment := true // bool | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  (optional) (default to true)
    ifModifiedSince := time.Now() // time.Time | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, `Wed, 09 Mar 2016 16:56:34 GMT`.  (optional)
    range_ := "range__example" // string | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes=1-10.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLinksApi.GetSharedLinkRenditionContent(context.Background(), sharedId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.GetSharedLinkRenditionContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSharedLinkRenditionContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SharedLinksApi.GetSharedLinkRenditionContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedId** | **string** | The identifier of a shared link to a file. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedLinkRenditionContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachment** | **bool** | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  | [default to true]
 **ifModifiedSince** | **time.Time** | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, &#x60;Wed, 09 Mar 2016 16:56:34 GMT&#x60;.  | 
 **range_** | **string** | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes&#x3D;1-10.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedLinkRenditions

> RenditionPaging ListSharedLinkRenditions(ctx, sharedId).Execute()

List renditions for a shared link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    sharedId := "sharedId_example" // string | The identifier of a shared link to a file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLinksApi.ListSharedLinkRenditions(context.Background(), sharedId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.ListSharedLinkRenditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSharedLinkRenditions`: RenditionPaging
    fmt.Fprintf(os.Stdout, "Response from `SharedLinksApi.ListSharedLinkRenditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedId** | **string** | The identifier of a shared link to a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSharedLinkRenditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RenditionPaging**](RenditionPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedLinks

> SharedLinkPaging ListSharedLinks(ctx).SkipCount(skipCount).MaxItems(maxItems).Where(where).Include(include).Fields(fields).Execute()

List shared links



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ahuret/goalfresco/openapi"
)

func main() {
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    where := "where_example" // string | Optionally filter the list by \"sharedByUser\" userid of person who shared the link (can also use -me-)  *   ```where=(sharedByUser='jbloggs')```  *   ```where=(sharedByUser='-me-')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the shared link, the following optional fields can be requested: * allowableOperations * path * properties * isFavorite * aspectNames  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedLinksApi.ListSharedLinks(context.Background()).SkipCount(skipCount).MaxItems(maxItems).Where(where).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLinksApi.ListSharedLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSharedLinks`: SharedLinkPaging
    fmt.Fprintf(os.Stdout, "Response from `SharedLinksApi.ListSharedLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSharedLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **where** | **string** | Optionally filter the list by \&quot;sharedByUser\&quot; userid of person who shared the link (can also use -me-)  *   &#x60;&#x60;&#x60;where&#x3D;(sharedByUser&#x3D;&#39;jbloggs&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(sharedByUser&#x3D;&#39;-me-&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the shared link, the following optional fields can be requested: * allowableOperations * path * properties * isFavorite * aspectNames  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SharedLinkPaging**](SharedLinkPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

