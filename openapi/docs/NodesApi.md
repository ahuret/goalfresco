# \NodesApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyNode**](NodesApi.md#CopyNode) | **Post** /nodes/{nodeId}/copy | Copy a node
[**CreateAssociation**](NodesApi.md#CreateAssociation) | **Post** /nodes/{nodeId}/targets | Create node association
[**CreateNode**](NodesApi.md#CreateNode) | **Post** /nodes/{nodeId}/children | Create a node
[**CreateSecondaryChildAssociation**](NodesApi.md#CreateSecondaryChildAssociation) | **Post** /nodes/{nodeId}/secondary-children | Create secondary child
[**DeleteAssociation**](NodesApi.md#DeleteAssociation) | **Delete** /nodes/{nodeId}/targets/{targetId} | Delete node association(s)
[**DeleteNode**](NodesApi.md#DeleteNode) | **Delete** /nodes/{nodeId} | Delete a node
[**DeleteSecondaryChildAssociation**](NodesApi.md#DeleteSecondaryChildAssociation) | **Delete** /nodes/{nodeId}/secondary-children/{childId} | Delete secondary child or children
[**GetNode**](NodesApi.md#GetNode) | **Get** /nodes/{nodeId} | Get a node
[**GetNodeContent**](NodesApi.md#GetNodeContent) | **Get** /nodes/{nodeId}/content | Get node content
[**ListNodeChildren**](NodesApi.md#ListNodeChildren) | **Get** /nodes/{nodeId}/children | List node children
[**ListParents**](NodesApi.md#ListParents) | **Get** /nodes/{nodeId}/parents | List parents
[**ListSecondaryChildren**](NodesApi.md#ListSecondaryChildren) | **Get** /nodes/{nodeId}/secondary-children | List secondary children
[**ListSourceAssociations**](NodesApi.md#ListSourceAssociations) | **Get** /nodes/{nodeId}/sources | List source associations
[**ListTargetAssociations**](NodesApi.md#ListTargetAssociations) | **Get** /nodes/{nodeId}/targets | List target associations
[**LockNode**](NodesApi.md#LockNode) | **Post** /nodes/{nodeId}/lock | Lock a node
[**MoveNode**](NodesApi.md#MoveNode) | **Post** /nodes/{nodeId}/move | Move a node
[**UnlockNode**](NodesApi.md#UnlockNode) | **Post** /nodes/{nodeId}/unlock | Unlock a node
[**UpdateNode**](NodesApi.md#UpdateNode) | **Put** /nodes/{nodeId} | Update a node
[**UpdateNodeContent**](NodesApi.md#UpdateNodeContent) | **Put** /nodes/{nodeId}/content | Update node content



## CopyNode

> NodeEntry CopyNode(ctx, nodeId).NodeBodyCopy(nodeBodyCopy).Include(include).Fields(fields).Execute()

Copy a node



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
    nodeBodyCopy := *openapiclient.NewNodeBodyCopy("TargetParentId_example") // NodeBodyCopy | The targetParentId and, optionally, a new name which should include the fileÂ extension.
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CopyNode(context.Background(), nodeId).NodeBodyCopy(nodeBodyCopy).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CopyNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CopyNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeBodyCopy** | [**NodeBodyCopy**](NodeBodyCopy.md) | The targetParentId and, optionally, a new name which should include the fileÂ extension. | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

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


## CreateAssociation

> AssociationEntry CreateAssociation(ctx, nodeId).AssociationBodyCreate(associationBodyCreate).Fields(fields).Execute()

Create node association



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
    nodeId := "nodeId_example" // string | The identifier of a source node.
    associationBodyCreate := *openapiclient.NewAssociationBody("TargetId_example", "AssocType_example") // AssociationBody | The target node id and assoc type.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateAssociation(context.Background(), nodeId).AssociationBodyCreate(associationBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssociation`: AssociationEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a source node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associationBodyCreate** | [**AssociationBody**](AssociationBody.md) | The target node id and assoc type. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**AssociationEntry**](AssociationEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNode

> NodeEntry CreateNode(ctx, nodeId).NodeBodyCreate(nodeBodyCreate).AutoRename(autoRename).MajorVersion(majorVersion).VersioningEnabled(versioningEnabled).Include(include).Fields(fields).Execute()

Create a node



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
    nodeId := "nodeId_example" // string | The identifier of a node. You can also use one of these well-known aliases: * -my- * -shared- * -root- 
    nodeBodyCreate := *openapiclient.NewNodeBodyCreate("Name_example", "NodeType_example") // NodeBodyCreate | The node information to create.
    autoRename := true // bool | If true, then  a name clash will cause an attempt to auto rename by finding a unique name using an integer suffix. (optional)
    majorVersion := true // bool | If true, then created node will be version *1.0 MAJOR*. If false, then created node will be version *0.1 MINOR*. (optional)
    versioningEnabled := true // bool | If true, then created node will be versioned. If false, then created node will be unversioned and auto-versioning disabled. (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateNode(context.Background(), nodeId).NodeBodyCreate(nodeBodyCreate).AutoRename(autoRename).MajorVersion(majorVersion).VersioningEnabled(versioningEnabled).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. You can also use one of these well-known aliases: * -my- * -shared- * -root-  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeBodyCreate** | [**NodeBodyCreate**](NodeBodyCreate.md) | The node information to create. | 
 **autoRename** | **bool** | If true, then  a name clash will cause an attempt to auto rename by finding a unique name using an integer suffix. | 
 **majorVersion** | **bool** | If true, then created node will be version *1.0 MAJOR*. If false, then created node will be version *0.1 MINOR*. | 
 **versioningEnabled** | **bool** | If true, then created node will be versioned. If false, then created node will be unversioned and auto-versioning disabled. | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeEntry**](NodeEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecondaryChildAssociation

> ChildAssociationEntry CreateSecondaryChildAssociation(ctx, nodeId).SecondaryChildAssociationBodyCreate(secondaryChildAssociationBodyCreate).Fields(fields).Execute()

Create secondary child



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
    nodeId := "nodeId_example" // string | The identifier of a parent node.
    secondaryChildAssociationBodyCreate := *openapiclient.NewChildAssociationBody("ChildId_example", "AssocType_example") // ChildAssociationBody | The child node id and assoc type.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateSecondaryChildAssociation(context.Background(), nodeId).SecondaryChildAssociationBodyCreate(secondaryChildAssociationBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateSecondaryChildAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecondaryChildAssociation`: ChildAssociationEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateSecondaryChildAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a parent node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecondaryChildAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secondaryChildAssociationBodyCreate** | [**ChildAssociationBody**](ChildAssociationBody.md) | The child node id and assoc type. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**ChildAssociationEntry**](ChildAssociationEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssociation

> DeleteAssociation(ctx, nodeId, targetId).AssocType(assocType).Execute()

Delete node association(s)



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
    nodeId := "nodeId_example" // string | The identifier of a source node.
    targetId := "targetId_example" // string | The identifier of a target node.
    assocType := "assocType_example" // string | Only delete associations of this type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NodesApi.DeleteAssociation(context.Background(), nodeId, targetId).AssocType(assocType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.DeleteAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a source node. | 
**targetId** | **string** | The identifier of a target node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assocType** | **string** | Only delete associations of this type. | 

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


## DeleteNode

> DeleteNode(ctx, nodeId).Permanent(permanent).Execute()

Delete a node



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
    permanent := true // bool | If **true** then the node is deleted permanently, without moving to the trashcan. Only the owner of the node or an admin can permanently delete the node.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NodesApi.DeleteNode(context.Background(), nodeId).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.DeleteNode``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permanent** | **bool** | If **true** then the node is deleted permanently, without moving to the trashcan. Only the owner of the node or an admin can permanently delete the node.  | [default to false]

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


## DeleteSecondaryChildAssociation

> DeleteSecondaryChildAssociation(ctx, nodeId, childId).AssocType(assocType).Execute()

Delete secondary child or children



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
    nodeId := "nodeId_example" // string | The identifier of a parent node.
    childId := "childId_example" // string | The identifier of a child node.
    assocType := "assocType_example" // string | Only delete associations of this type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NodesApi.DeleteSecondaryChildAssociation(context.Background(), nodeId, childId).AssocType(assocType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.DeleteSecondaryChildAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a parent node. | 
**childId** | **string** | The identifier of a child node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecondaryChildAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assocType** | **string** | Only delete associations of this type. | 

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


## GetNode

> NodeEntry GetNode(ctx, nodeId).Include(include).RelativePath(relativePath).Fields(fields).Execute()

Get a node



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
    nodeId := "nodeId_example" // string | The identifier of a node. You can also use one of these well-known aliases: * -my- * -shared- * -root- 
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)
    relativePath := "relativePath_example" // string | A path relative to the **nodeId**. If you set this, information is returned on the node resolved by this path.  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetNode(context.Background(), nodeId).Include(include).RelativePath(relativePath).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. You can also use one of these well-known aliases: * -my- * -shared- * -root-  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **relativePath** | **string** | A path relative to the **nodeId**. If you set this, information is returned on the node resolved by this path.  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeEntry**](NodeEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeContent

> *os.File GetNodeContent(ctx, nodeId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()

Get node content



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
    resp, r, err := apiClient.NodesApi.GetNodeContent(context.Background(), nodeId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetNodeContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetNodeContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeContentRequest struct via the builder pattern


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


## ListNodeChildren

> NodeChildAssociationPaging ListNodeChildren(ctx, nodeId).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Where(where).Include(include).RelativePath(relativePath).IncludeSource(includeSource).Fields(fields).Execute()

List node children



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
    nodeId := "nodeId_example" // string | The identifier of a node. You can also use one of these well-known aliases: * -my- * -shared- * -root- 
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)
    where := "where_example" // string | Optionally filter the list. Here are some examples:  *   ```where=(isFolder=true)```  *   ```where=(isFile=true)```  *   ```where=(nodeType='my:specialNodeType')```  *   ```where=(nodeType='my:specialNodeType INCLUDESUBTYPES')```  *   ```where=(isPrimary=true)```  *   ```where=(assocType='my:specialAssocType')```  *   ```where=(isPrimary=false and assocType='my:specialAssocType')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * association * isLink * isFavorite * isLocked * path * properties * permissions  (optional)
    relativePath := "relativePath_example" // string | Return information on children in the folder resolved by this path. The path is relative to **nodeId**. (optional)
    includeSource := true // bool | Also include **source** in addition to **entries** with folder information on the parent node â€“ either the specified parent **nodeId**, or as resolved by **relativePath**. (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ListNodeChildren(context.Background(), nodeId).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Where(where).Include(include).RelativePath(relativePath).IncludeSource(includeSource).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ListNodeChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodeChildren`: NodeChildAssociationPaging
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ListNodeChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. You can also use one of these well-known aliases: * -my- * -shared- * -root-  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodeChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **where** | **string** | Optionally filter the list. Here are some examples:  *   &#x60;&#x60;&#x60;where&#x3D;(isFolder&#x3D;true)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(isFile&#x3D;true)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(nodeType&#x3D;&#39;my:specialNodeType&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(nodeType&#x3D;&#39;my:specialNodeType INCLUDESUBTYPES&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(isPrimary&#x3D;true)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(assocType&#x3D;&#39;my:specialAssocType&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(isPrimary&#x3D;false and assocType&#x3D;&#39;my:specialAssocType&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * association * isLink * isFavorite * isLocked * path * properties * permissions  | 
 **relativePath** | **string** | Return information on children in the folder resolved by this path. The path is relative to **nodeId**. | 
 **includeSource** | **bool** | Also include **source** in addition to **entries** with folder information on the parent node â€“ either the specified parent **nodeId**, or as resolved by **relativePath**. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeChildAssociationPaging**](NodeChildAssociationPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParents

> NodeAssociationPaging ListParents(ctx, nodeId).Where(where).Include(include).SkipCount(skipCount).MaxItems(maxItems).IncludeSource(includeSource).Fields(fields).Execute()

List parents



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
    nodeId := "nodeId_example" // string | The identifier of a child node. You can also use one of these well-known aliases: * -my- * -shared- * -root- 
    where := "where_example" // string | Optionally filter the list by **assocType** and/or **isPrimary**. Here are some example filters:  *   ```where=(assocType='my:specialAssocType')```  *   ```where=(isPrimary=true)```  *   ```where=(isPrimary=false and assocType='my:specialAssocType')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  (optional)
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    includeSource := true // bool | Also include **source** (in addition to **entries**) with folder information on **nodeId** (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ListParents(context.Background(), nodeId).Where(where).Include(include).SkipCount(skipCount).MaxItems(maxItems).IncludeSource(includeSource).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ListParents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListParents`: NodeAssociationPaging
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ListParents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a child node. You can also use one of these well-known aliases: * -my- * -shared- * -root-  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListParentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **string** | Optionally filter the list by **assocType** and/or **isPrimary**. Here are some example filters:  *   &#x60;&#x60;&#x60;where&#x3D;(assocType&#x3D;&#39;my:specialAssocType&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(isPrimary&#x3D;true)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(isPrimary&#x3D;false and assocType&#x3D;&#39;my:specialAssocType&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  | 
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **includeSource** | **bool** | Also include **source** (in addition to **entries**) with folder information on **nodeId** | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeAssociationPaging**](NodeAssociationPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecondaryChildren

> NodeChildAssociationPaging ListSecondaryChildren(ctx, nodeId).Where(where).Include(include).SkipCount(skipCount).MaxItems(maxItems).IncludeSource(includeSource).Fields(fields).Execute()

List secondary children



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
    nodeId := "nodeId_example" // string | The identifier of a parent node. You can also use one of these well-known aliases: * -my- * -shared- * -root- 
    where := "where_example" // string | Optionally filter the list by assocType. Here's an example:  *   ```where=(assocType='my:specialAssocType')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  (optional)
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    includeSource := true // bool | Also include **source** (in addition to **entries**) with folder information on **nodeId** (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ListSecondaryChildren(context.Background(), nodeId).Where(where).Include(include).SkipCount(skipCount).MaxItems(maxItems).IncludeSource(includeSource).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ListSecondaryChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecondaryChildren`: NodeChildAssociationPaging
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ListSecondaryChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a parent node. You can also use one of these well-known aliases: * -my- * -shared- * -root-  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecondaryChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **string** | Optionally filter the list by assocType. Here&#39;s an example:  *   &#x60;&#x60;&#x60;where&#x3D;(assocType&#x3D;&#39;my:specialAssocType&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  | 
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **includeSource** | **bool** | Also include **source** (in addition to **entries**) with folder information on **nodeId** | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeChildAssociationPaging**](NodeChildAssociationPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceAssociations

> NodeAssociationPaging ListSourceAssociations(ctx, nodeId).Where(where).Include(include).Fields(fields).Execute()

List source associations



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
    nodeId := "nodeId_example" // string | The identifier of a target node.
    where := "where_example" // string | Optionally filter the list by **assocType**. Here's an example:  *   ```where=(assocType='my:specialAssocType')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ListSourceAssociations(context.Background(), nodeId).Where(where).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ListSourceAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceAssociations`: NodeAssociationPaging
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ListSourceAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a target node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **string** | Optionally filter the list by **assocType**. Here&#39;s an example:  *   &#x60;&#x60;&#x60;where&#x3D;(assocType&#x3D;&#39;my:specialAssocType&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeAssociationPaging**](NodeAssociationPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTargetAssociations

> NodeAssociationPaging ListTargetAssociations(ctx, nodeId).Where(where).Include(include).Fields(fields).Execute()

List target associations



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
    nodeId := "nodeId_example" // string | The identifier of a source node.
    where := "where_example" // string | Optionally filter the list by **assocType**. Here's an example:  *   ```where=(assocType='my:specialAssocType')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ListTargetAssociations(context.Background(), nodeId).Where(where).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ListTargetAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTargetAssociations`: NodeAssociationPaging
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ListTargetAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a source node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTargetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **string** | Optionally filter the list by **assocType**. Here&#39;s an example:  *   &#x60;&#x60;&#x60;where&#x3D;(assocType&#x3D;&#39;my:specialAssocType&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeAssociationPaging**](NodeAssociationPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockNode

> NodeEntry LockNode(ctx, nodeId).NodeBodyLock(nodeBodyLock).Include(include).Fields(fields).Execute()

Lock a node



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
    nodeBodyLock := *openapiclient.NewNodeBodyLock() // NodeBodyLock | Lock details.
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.LockNode(context.Background(), nodeId).NodeBodyLock(nodeBodyLock).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.LockNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.LockNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeBodyLock** | [**NodeBodyLock**](NodeBodyLock.md) | Lock details. | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

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


## MoveNode

> NodeEntry MoveNode(ctx, nodeId).NodeBodyMove(nodeBodyMove).Include(include).Fields(fields).Execute()

Move a node



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
    nodeBodyMove := *openapiclient.NewNodeBodyMove("TargetParentId_example") // NodeBodyMove | The targetParentId and, optionally, a new name which should include the fileÂ extension.
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.MoveNode(context.Background(), nodeId).NodeBodyMove(nodeBodyMove).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.MoveNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.MoveNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeBodyMove** | [**NodeBodyMove**](NodeBodyMove.md) | The targetParentId and, optionally, a new name which should include the fileÂ extension. | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

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


## UnlockNode

> NodeEntry UnlockNode(ctx, nodeId).Include(include).Fields(fields).Execute()

Unlock a node



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
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UnlockNode(context.Background(), nodeId).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UnlockNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UnlockNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeEntry**](NodeEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNode

> NodeEntry UpdateNode(ctx, nodeId).NodeBodyUpdate(nodeBodyUpdate).Include(include).Fields(fields).Execute()

Update a node



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
    nodeBodyUpdate := *openapiclient.NewNodeBodyUpdate() // NodeBodyUpdate | The node information to update.
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UpdateNode(context.Background(), nodeId).NodeBodyUpdate(nodeBodyUpdate).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UpdateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNode`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeBodyUpdate** | [**NodeBodyUpdate**](NodeBodyUpdate.md) | The node information to update. | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

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


## UpdateNodeContent

> NodeEntry UpdateNodeContent(ctx, nodeId).ContentBodyUpdate(contentBodyUpdate).MajorVersion(majorVersion).Comment(comment).Name(name).Include(include).Fields(fields).Execute()

Update node content



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
    contentBodyUpdate := os.NewFile(1234, "some_file") // *os.File | The binary content
    majorVersion := true // bool | If **true**, create a major version. Setting this parameter also enables versioning of this node, if it is not already versioned.  (optional) (default to false)
    comment := "comment_example" // string | Add a version comment which will appear in version history. Setting this parameter also enables versioning of this node, if it is not already versioned.  (optional)
    name := "name_example" // string | Optional new name. This should include the fileÂ extension. The name must not contain spaces or the following special characters: * \" < > \\ / ? : and |. The character `.` must not be used at the end of the name.  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UpdateNodeContent(context.Background(), nodeId).ContentBodyUpdate(contentBodyUpdate).MajorVersion(majorVersion).Comment(comment).Name(name).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UpdateNodeContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNodeContent`: NodeEntry
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UpdateNodeContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentBodyUpdate** | ***os.File** | The binary content | 
 **majorVersion** | **bool** | If **true**, create a major version. Setting this parameter also enables versioning of this node, if it is not already versioned.  | [default to false]
 **comment** | **string** | Add a version comment which will appear in version history. Setting this parameter also enables versioning of this node, if it is not already versioned.  | 
 **name** | **string** | Optional new name. This should include the fileÂ extension. The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character &#x60;.&#x60; must not be used at the end of the name.  | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * association * isLink * isFavorite * isLocked * path * permissions * definition  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodeEntry**](NodeEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

