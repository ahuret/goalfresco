# \RenditionsApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRendition**](RenditionsApi.md#CreateRendition) | **Post** /nodes/{nodeId}/renditions | Create rendition
[**GetRendition**](RenditionsApi.md#GetRendition) | **Get** /nodes/{nodeId}/renditions/{renditionId} | Get rendition information
[**GetRenditionContent**](RenditionsApi.md#GetRenditionContent) | **Get** /nodes/{nodeId}/renditions/{renditionId}/content | Get rendition content
[**ListRenditions**](RenditionsApi.md#ListRenditions) | **Get** /nodes/{nodeId}/renditions | List renditions



## CreateRendition

> CreateRendition(ctx, nodeId).RenditionBodyCreate(renditionBodyCreate).Execute()

Create rendition



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
    renditionBodyCreate := *openapiclient.NewRenditionBodyCreate("Id_example") // RenditionBodyCreate | The rendition \"id\".

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RenditionsApi.CreateRendition(context.Background(), nodeId).RenditionBodyCreate(renditionBodyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenditionsApi.CreateRendition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRenditionRequest struct via the builder pattern


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


## GetRendition

> RenditionEntry GetRendition(ctx, nodeId, renditionId).Execute()

Get rendition information



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
    renditionId := "renditionId_example" // string | The name of a thumbnail rendition, for example *doclib*, or *pdf*.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenditionsApi.GetRendition(context.Background(), nodeId, renditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenditionsApi.GetRendition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRendition`: RenditionEntry
    fmt.Fprintf(os.Stdout, "Response from `RenditionsApi.GetRendition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenditionRequest struct via the builder pattern


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


## GetRenditionContent

> *os.File GetRenditionContent(ctx, nodeId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Placeholder(placeholder).Execute()

Get rendition content



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
    renditionId := "renditionId_example" // string | The name of a thumbnail rendition, for example *doclib*, or *pdf*.
    attachment := true // bool | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  (optional) (default to true)
    ifModifiedSince := time.Now() // time.Time | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, `Wed, 09 Mar 2016 16:56:34 GMT`.  (optional)
    range_ := "range__example" // string | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes=1-10.  (optional)
    placeholder := true // bool | If **true** and there is no rendition for this **nodeId** and **renditionId**, then the placeholder image for the mime type of this rendition is returned, rather than a 404 response.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenditionsApi.GetRenditionContent(context.Background(), nodeId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Placeholder(placeholder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenditionsApi.GetRenditionContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRenditionContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `RenditionsApi.GetRenditionContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRenditionContentRequest struct via the builder pattern


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


## ListRenditions

> RenditionPaging ListRenditions(ctx, nodeId).Where(where).Execute()

List renditions



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
    where := "where_example" // string | A string to restrict the returned objects by using a predicate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenditionsApi.ListRenditions(context.Background(), nodeId).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenditionsApi.ListRenditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRenditions`: RenditionPaging
    fmt.Fprintf(os.Stdout, "Response from `RenditionsApi.ListRenditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRenditionsRequest struct via the builder pattern


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

