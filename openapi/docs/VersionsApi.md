# \VersionsApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVersionRendition**](VersionsApi.md#CreateVersionRendition) | **Post** /nodes/{nodeId}/versions/{versionId}/renditions | Create rendition for a file version
[**DeleteVersion**](VersionsApi.md#DeleteVersion) | **Delete** /nodes/{nodeId}/versions/{versionId} | Delete a version
[**GetVersion**](VersionsApi.md#GetVersion) | **Get** /nodes/{nodeId}/versions/{versionId} | Get version information
[**GetVersionContent**](VersionsApi.md#GetVersionContent) | **Get** /nodes/{nodeId}/versions/{versionId}/content | Get version content
[**GetVersionRendition**](VersionsApi.md#GetVersionRendition) | **Get** /nodes/{nodeId}/versions/{versionId}/renditions/{renditionId} | Get rendition information for a file version
[**GetVersionRenditionContent**](VersionsApi.md#GetVersionRenditionContent) | **Get** /nodes/{nodeId}/versions/{versionId}/renditions/{renditionId}/content | Get rendition content for a file version
[**ListVersionHistory**](VersionsApi.md#ListVersionHistory) | **Get** /nodes/{nodeId}/versions | List version history
[**ListVersionRenditions**](VersionsApi.md#ListVersionRenditions) | **Get** /nodes/{nodeId}/versions/{versionId}/renditions | List renditions for a file version
[**RevertVersion**](VersionsApi.md#RevertVersion) | **Post** /nodes/{nodeId}/versions/{versionId}/revert | Revert a version



## CreateVersionRendition

> CreateVersionRendition(ctx, nodeId, versionId).RenditionBodyCreate(renditionBodyCreate).Execute()

Create rendition for a file version



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.
    renditionBodyCreate := *openapiclient.NewRenditionBodyCreate("Id_example") // RenditionBodyCreate | The rendition \"id\".

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VersionsApi.CreateVersionRendition(context.Background(), nodeId, versionId).RenditionBodyCreate(renditionBodyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.CreateVersionRendition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVersionRenditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **renditionBodyCreate** | [**RenditionBodyCreate**](RenditionBodyCreate.md) | The rendition \&quot;id\&quot;. | 

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


## DeleteVersion

> DeleteVersion(ctx, nodeId, versionId).Execute()

Delete a version



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VersionsApi.DeleteVersion(context.Background(), nodeId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.DeleteVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVersionRequest struct via the builder pattern


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


## GetVersion

> VersionEntry GetVersion(ctx, nodeId, versionId).Execute()

Get version information



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.GetVersion(context.Background(), nodeId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: VersionEntry
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.GetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionEntry**](VersionEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionContent

> *os.File GetVersionContent(ctx, nodeId, versionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()

Get version content



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.
    attachment := true // bool | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  (optional) (default to true)
    ifModifiedSince := time.Now() // time.Time | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, `Wed, 09 Mar 2016 16:56:34 GMT`.  (optional)
    range_ := "range__example" // string | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes=1-10.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.GetVersionContent(context.Background(), nodeId, versionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.GetVersionContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersionContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.GetVersionContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionContentRequest struct via the builder pattern


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


## GetVersionRendition

> RenditionEntry GetVersionRendition(ctx, nodeId, versionId, renditionId).Execute()

Get rendition information for a file version



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.
    renditionId := "renditionId_example" // string | The name of a thumbnail rendition, for example *doclib*, or *pdf*.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.GetVersionRendition(context.Background(), nodeId, versionId, renditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.GetVersionRendition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersionRendition`: RenditionEntry
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.GetVersionRendition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRenditionRequest struct via the builder pattern


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


## GetVersionRenditionContent

> *os.File GetVersionRenditionContent(ctx, nodeId, versionId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Placeholder(placeholder).Execute()

Get rendition content for a file version



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.
    renditionId := "renditionId_example" // string | The name of a thumbnail rendition, for example *doclib*, or *pdf*.
    attachment := true // bool | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  (optional) (default to true)
    ifModifiedSince := time.Now() // time.Time | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, `Wed, 09 Mar 2016 16:56:34 GMT`.  (optional)
    range_ := "range__example" // string | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes=1-10.  (optional)
    placeholder := true // bool | If **true** and there is no rendition for this **nodeId** and **renditionId**, then the placeholder image for the mime type of this rendition is returned, rather than a 404 response.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.GetVersionRenditionContent(context.Background(), nodeId, versionId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Placeholder(placeholder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.GetVersionRenditionContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersionRenditionContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.GetVersionRenditionContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRenditionContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **attachment** | **bool** | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  | [default to true]
 **ifModifiedSince** | **time.Time** | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, &#x60;Wed, 09 Mar 2016 16:56:34 GMT&#x60;.  | 
 **range_** | **string** | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes&#x3D;1-10.  | 
 **placeholder** | **bool** | If **true** and there is no rendition for this **nodeId** and **renditionId**, then the placeholder image for the mime type of this rendition is returned, rather than a 404 response.  | [default to false]

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


## ListVersionHistory

> VersionPaging ListVersionHistory(ctx, nodeId).Include(include).Fields(fields).SkipCount(skipCount).MaxItems(maxItems).Execute()

List version history



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    include := []string{"Inner_example"} // []string | Returns additional information about the version node. The following optional fields can be requested: * properties * aspectNames  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.ListVersionHistory(context.Background(), nodeId).Include(include).Fields(fields).SkipCount(skipCount).MaxItems(maxItems).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ListVersionHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersionHistory`: VersionPaging
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ListVersionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Returns additional information about the version node. The following optional fields can be requested: * properties * aspectNames  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]

### Return type

[**VersionPaging**](VersionPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersionRenditions

> RenditionPaging ListVersionRenditions(ctx, nodeId, versionId).Where(where).Execute()

List renditions for a file version



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.
    where := "where_example" // string | A string to restrict the returned objects by using a predicate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.ListVersionRenditions(context.Background(), nodeId, versionId).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.ListVersionRenditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersionRenditions`: RenditionPaging
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.ListVersionRenditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionRenditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **where** | **string** | A string to restrict the returned objects by using a predicate. | 

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


## RevertVersion

> VersionEntry RevertVersion(ctx, nodeId, versionId).RevertBody(revertBody).Fields(fields).Execute()

Revert a version



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
    nodeId := "nodeId_example" // string | The identifier of a node.
    versionId := "versionId_example" // string | The identifier of a version, ie. version label, within the version history of a node.
    revertBody := *openapiclient.NewRevertBody() // RevertBody | Optionally, specify a version comment and whether this should be a major version, or not.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsApi.RevertVersion(context.Background(), nodeId, versionId).RevertBody(revertBody).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.RevertVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertVersion`: VersionEntry
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.RevertVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**versionId** | **string** | The identifier of a version, ie. version label, within the version history of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revertBody** | [**RevertBody**](RevertBody.md) | Optionally, specify a version comment and whether this should be a major version, or not. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**VersionEntry**](VersionEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

