# \QueriesApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindNodes**](QueriesApi.md#FindNodes) | **Get** /queries/nodes | Find nodes
[**FindPeople**](QueriesApi.md#FindPeople) | **Get** /queries/people | Find people
[**FindSites**](QueriesApi.md#FindSites) | **Get** /queries/sites | Find sites



## FindNodes

> NodePaging FindNodes(ctx).Term(term).RootNodeId(rootNodeId).SkipCount(skipCount).MaxItems(maxItems).NodeType(nodeType).Include(include).OrderBy(orderBy).Fields(fields).Execute()

Find nodes



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
    term := "term_example" // string | The term to search for.
    rootNodeId := "rootNodeId_example" // string | The id of the node to start the search from.  Supports the aliases -my-, -root- and -shared-.  (optional)
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    nodeType := "nodeType_example" // string | Restrict the returned results to only those of the given node type and its sub-types  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  (optional)
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.FindNodes(context.Background()).Term(term).RootNodeId(rootNodeId).SkipCount(skipCount).MaxItems(maxItems).NodeType(nodeType).Include(include).OrderBy(orderBy).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.FindNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindNodes`: NodePaging
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.FindNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** | The term to search for. | 
 **rootNodeId** | **string** | The id of the node to start the search from.  Supports the aliases -my-, -root- and -shared-.  | 
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **nodeType** | **string** | Restrict the returned results to only those of the given node type and its sub-types  | 
 **include** | **[]string** | Returns additional information about the node. The following optional fields can be requested: * allowableOperations * aspectNames * isLink * isFavorite * isLocked * path * properties  | 
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**NodePaging**](NodePaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPeople

> PersonPaging FindPeople(ctx).Term(term).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).OrderBy(orderBy).Execute()

Find people



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
    term := "term_example" // string | The term to search for. 
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.FindPeople(context.Background()).Term(term).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.FindPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPeople`: PersonPaging
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.FindPeople`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindPeopleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** | The term to search for.  | 
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 

### Return type

[**PersonPaging**](PersonPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSites

> SitePaging FindSites(ctx).Term(term).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Fields(fields).Execute()

Find sites



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
    term := "term_example" // string | The term to search for.
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueriesApi.FindSites(context.Background()).Term(term).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueriesApi.FindSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSites`: SitePaging
    fmt.Fprintf(os.Stdout, "Response from `QueriesApi.FindSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** | The term to search for. | 
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SitePaging**](SitePaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

