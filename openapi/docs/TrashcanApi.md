# \TrashcanApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeletedNode**](TrashcanApi.md#DeleteDeletedNode) | **Delete** /deleted-nodes/{nodeId} | Permanently delete a deleted node
[**GetArchivedNodeRendition**](TrashcanApi.md#GetArchivedNodeRendition) | **Get** /deleted-nodes/{nodeId}/renditions/{renditionId} | Get rendition information for a deleted node
[**GetArchivedNodeRenditionContent**](TrashcanApi.md#GetArchivedNodeRenditionContent) | **Get** /deleted-nodes/{nodeId}/renditions/{renditionId}/content | Get rendition content of a deleted node
[**GetDeletedNode**](TrashcanApi.md#GetDeletedNode) | **Get** /deleted-nodes/{nodeId} | Get a deleted node
[**GetDeletedNodeContent**](TrashcanApi.md#GetDeletedNodeContent) | **Get** /deleted-nodes/{nodeId}/content | Get deleted node content
[**ListDeletedNodeRenditions**](TrashcanApi.md#ListDeletedNodeRenditions) | **Get** /deleted-nodes/{nodeId}/renditions | List renditions for a deleted node
[**ListDeletedNodes**](TrashcanApi.md#ListDeletedNodes) | **Get** /deleted-nodes | List deleted nodes
[**RestoreDeletedNode**](TrashcanApi.md#RestoreDeletedNode) | **Post** /deleted-nodes/{nodeId}/restore | Restore a deleted node



## DeleteDeletedNode

> DeleteDeletedNode(ctx, nodeId).Execute()

Permanently delete a deleted node



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TrashcanApi.DeleteDeletedNode(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.DeleteDeletedNode``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDeletedNodeRequest struct via the builder pattern


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


## GetArchivedNodeRendition

> RenditionEntry GetArchivedNodeRendition(ctx, nodeId, renditionId).Execute()

Get rendition information for a deleted node



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
    resp, r, err := apiClient.TrashcanApi.GetArchivedNodeRendition(context.Background(), nodeId, renditionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.GetArchivedNodeRendition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchivedNodeRendition`: RenditionEntry
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.GetArchivedNodeRendition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedNodeRenditionRequest struct via the builder pattern


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


## GetArchivedNodeRenditionContent

> *os.File GetArchivedNodeRenditionContent(ctx, nodeId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Placeholder(placeholder).Execute()

Get rendition content of a deleted node



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
    resp, r, err := apiClient.TrashcanApi.GetArchivedNodeRenditionContent(context.Background(), nodeId, renditionId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Placeholder(placeholder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.GetArchivedNodeRenditionContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchivedNodeRenditionContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.GetArchivedNodeRenditionContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**renditionId** | **string** | The name of a thumbnail rendition, for example *doclib*, or *pdf*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedNodeRenditionContentRequest struct via the builder pattern


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


## GetDeletedNode

> DeletedNodeEntry GetDeletedNode(ctx, nodeId).Include(include).Execute()

Get a deleted node



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
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrashcanApi.GetDeletedNode(context.Background(), nodeId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.GetDeletedNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeletedNode`: DeletedNodeEntry
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.GetDeletedNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeletedNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 

### Return type

[**DeletedNodeEntry**](DeletedNodeEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeletedNodeContent

> *os.File GetDeletedNodeContent(ctx, nodeId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()

Get deleted node content



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
    attachment := true // bool | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  (optional) (default to true)
    ifModifiedSince := time.Now() // time.Time | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, `Wed, 09 Mar 2016 16:56:34 GMT`.  (optional)
    range_ := "range__example" // string | The Range header indicates the part of a document that the server should return. Single part request supported, for example: bytes=1-10.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrashcanApi.GetDeletedNodeContent(context.Background(), nodeId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.GetDeletedNodeContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeletedNodeContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.GetDeletedNodeContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeletedNodeContentRequest struct via the builder pattern


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


## ListDeletedNodeRenditions

> RenditionPaging ListDeletedNodeRenditions(ctx, nodeId).Where(where).Execute()

List renditions for a deleted node



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
    resp, r, err := apiClient.TrashcanApi.ListDeletedNodeRenditions(context.Background(), nodeId).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.ListDeletedNodeRenditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeletedNodeRenditions`: RenditionPaging
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.ListDeletedNodeRenditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeletedNodeRenditionsRequest struct via the builder pattern


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


## ListDeletedNodes

> DeletedNodesPaging ListDeletedNodes(ctx).SkipCount(skipCount).MaxItems(maxItems).Include(include).Execute()

List deleted nodes



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
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * association * isLink * isFavorite * isLocked * path * properties * permissions  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrashcanApi.ListDeletedNodes(context.Background()).SkipCount(skipCount).MaxItems(maxItems).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.ListDeletedNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeletedNodes`: DeletedNodesPaging
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.ListDeletedNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeletedNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * association * isLink * isFavorite * isLocked * path * properties * permissions  | 

### Return type

[**DeletedNodesPaging**](DeletedNodesPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDeletedNode

> NodeEntry RestoreDeletedNode(ctx, nodeId).Fields(fields).DeletedNodeBodyRestore(deletedNodeBodyRestore).Execute()

Restore a deleted node



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
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)
    deletedNodeBodyRestore := *openapiclient.NewDeletedNodeBodyRestore() // DeletedNodeBodyRestore | The targetParentId if the node is restored to a new location. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrashcanApi.RestoreDeletedNode(context.Background(), nodeId).Fields(fields).DeletedNodeBodyRestore(deletedNodeBodyRestore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrashcanApi.RestoreDeletedNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreDeletedNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `TrashcanApi.RestoreDeletedNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDeletedNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 
 **deletedNodeBodyRestore** | [**DeletedNodeBodyRestore**](DeletedNodeBodyRestore.md) | The targetParentId if the node is restored to a new location. | 

### Return type

[**NodeEntry**](NodeEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

