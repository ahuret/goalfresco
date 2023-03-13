# \AuditApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuditEntriesForAuditApp**](AuditApi.md#DeleteAuditEntriesForAuditApp) | **Delete** /audit-applications/{auditApplicationId}/audit-entries | Permanently delete audit entries for an audit application
[**DeleteAuditEntry**](AuditApi.md#DeleteAuditEntry) | **Delete** /audit-applications/{auditApplicationId}/audit-entries/{auditEntryId} | Permanently delete an audit entry
[**GetAuditApp**](AuditApi.md#GetAuditApp) | **Get** /audit-applications/{auditApplicationId} | Get audit application info
[**GetAuditEntry**](AuditApi.md#GetAuditEntry) | **Get** /audit-applications/{auditApplicationId}/audit-entries/{auditEntryId} | Get audit entry
[**ListAuditApps**](AuditApi.md#ListAuditApps) | **Get** /audit-applications | List audit applications
[**ListAuditEntriesForAuditApp**](AuditApi.md#ListAuditEntriesForAuditApp) | **Get** /audit-applications/{auditApplicationId}/audit-entries | List audit entries for an audit application
[**ListAuditEntriesForNode**](AuditApi.md#ListAuditEntriesForNode) | **Get** /nodes/{nodeId}/audit-entries | List audit entries for a node
[**UpdateAuditApp**](AuditApi.md#UpdateAuditApp) | **Put** /audit-applications/{auditApplicationId} | Update audit application info



## DeleteAuditEntriesForAuditApp

> DeleteAuditEntriesForAuditApp(ctx, auditApplicationId).Where(where).Execute()

Permanently delete audit entries for an audit application



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
    auditApplicationId := "auditApplicationId_example" // string | The identifier of an audit application.
    where := "where_example" // string | Audit entries to permanently delete for an audit application, given an inclusive time period or range of ids. For example:  *   ```where=(createdAt BETWEEN ('2017-06-02T12:13:51.593+01:00' , '2017-06-04T10:05:16.536+01:00')``` *   ```where=(id BETWEEN ('1234', '4321')``` 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuditApi.DeleteAuditEntriesForAuditApp(context.Background(), auditApplicationId).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.DeleteAuditEntriesForAuditApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditApplicationId** | **string** | The identifier of an audit application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuditEntriesForAuditAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **string** | Audit entries to permanently delete for an audit application, given an inclusive time period or range of ids. For example:  *   &#x60;&#x60;&#x60;where&#x3D;(createdAt BETWEEN (&#39;2017-06-02T12:13:51.593+01:00&#39; , &#39;2017-06-04T10:05:16.536+01:00&#39;)&#x60;&#x60;&#x60; *   &#x60;&#x60;&#x60;where&#x3D;(id BETWEEN (&#39;1234&#39;, &#39;4321&#39;)&#x60;&#x60;&#x60;  | 

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


## DeleteAuditEntry

> DeleteAuditEntry(ctx, auditApplicationId, auditEntryId).Execute()

Permanently delete an audit entry



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
    auditApplicationId := "auditApplicationId_example" // string | The identifier of an audit application.
    auditEntryId := "auditEntryId_example" // string | The identifier of an audit entry.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuditApi.DeleteAuditEntry(context.Background(), auditApplicationId, auditEntryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.DeleteAuditEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditApplicationId** | **string** | The identifier of an audit application. | 
**auditEntryId** | **string** | The identifier of an audit entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuditEntryRequest struct via the builder pattern


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


## GetAuditApp

> AuditApp GetAuditApp(ctx, auditApplicationId).Fields(fields).Include(include).Execute()

Get audit application info



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
    auditApplicationId := "auditApplicationId_example" // string | The identifier of an audit application.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)
    include := []string{"Inner_example"} // []string | Also include the current minimum and/or maximum audit entry ids for the application. The following optional fields can be requested: * max * min  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.GetAuditApp(context.Background(), auditApplicationId).Fields(fields).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.GetAuditApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditApp`: AuditApp
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.GetAuditApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditApplicationId** | **string** | The identifier of an audit application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 
 **include** | **[]string** | Also include the current minimum and/or maximum audit entry ids for the application. The following optional fields can be requested: * max * min  | 

### Return type

[**AuditApp**](AuditApp.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditEntry

> AuditEntryEntry GetAuditEntry(ctx, auditApplicationId, auditEntryId).Fields(fields).Execute()

Get audit entry



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
    auditApplicationId := "auditApplicationId_example" // string | The identifier of an audit application.
    auditEntryId := "auditEntryId_example" // string | The identifier of an audit entry.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.GetAuditEntry(context.Background(), auditApplicationId, auditEntryId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.GetAuditEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditEntry`: AuditEntryEntry
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.GetAuditEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditApplicationId** | **string** | The identifier of an audit application. | 
**auditEntryId** | **string** | The identifier of an audit entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**AuditEntryEntry**](AuditEntryEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuditApps

> AuditAppPaging ListAuditApps(ctx).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()

List audit applications



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
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.ListAuditApps(context.Background()).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.ListAuditApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditApps`: AuditAppPaging
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.ListAuditApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**AuditAppPaging**](AuditAppPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuditEntriesForAuditApp

> AuditEntryPaging ListAuditEntriesForAuditApp(ctx, auditApplicationId).SkipCount(skipCount).OmitTotalItems(omitTotalItems).OrderBy(orderBy).MaxItems(maxItems).Where(where).Include(include).Fields(fields).Execute()

List audit entries for an audit application



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
    auditApplicationId := "auditApplicationId_example" // string | The identifier of an audit application.
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    omitTotalItems := true // bool | A boolean to control if the response provides the total numbers of items in the collection. If not supplied then the default value is false.  (optional) (default to false)
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    where := "where_example" // string | Optionally filter the list. Here are some examples:  *   ```where=(createdByUser='jbloggs')```  *   ```where=(id BETWEEN ('1234', '4321')```  *   ```where=(createdAt BETWEEN ('2017-06-02T12:13:51.593+01:00' , '2017-06-04T10:05:16.536+01:00')```  *   ```where=(createdByUser='jbloggs' and createdAt BETWEEN ('2017-06-02T12:13:51.593+01:00' , '2017-06-04T10:05:16.536+01:00')```  *   ```where=(valuesKey='/alfresco-access/login/user')```  *   ```where=(valuesKey='/alfresco-access/transaction/action' and valuesValue='DELETE')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the audit entry. The following optional fields can be requested: * values  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.ListAuditEntriesForAuditApp(context.Background(), auditApplicationId).SkipCount(skipCount).OmitTotalItems(omitTotalItems).OrderBy(orderBy).MaxItems(maxItems).Where(where).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.ListAuditEntriesForAuditApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditEntriesForAuditApp`: AuditEntryPaging
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.ListAuditEntriesForAuditApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditApplicationId** | **string** | The identifier of an audit application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuditEntriesForAuditAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **omitTotalItems** | **bool** | A boolean to control if the response provides the total numbers of items in the collection. If not supplied then the default value is false.  | [default to false]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **where** | **string** | Optionally filter the list. Here are some examples:  *   &#x60;&#x60;&#x60;where&#x3D;(createdByUser&#x3D;&#39;jbloggs&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(id BETWEEN (&#39;1234&#39;, &#39;4321&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(createdAt BETWEEN (&#39;2017-06-02T12:13:51.593+01:00&#39; , &#39;2017-06-04T10:05:16.536+01:00&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(createdByUser&#x3D;&#39;jbloggs&#39; and createdAt BETWEEN (&#39;2017-06-02T12:13:51.593+01:00&#39; , &#39;2017-06-04T10:05:16.536+01:00&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(valuesKey&#x3D;&#39;/alfresco-access/login/user&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(valuesKey&#x3D;&#39;/alfresco-access/transaction/action&#39; and valuesValue&#x3D;&#39;DELETE&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the audit entry. The following optional fields can be requested: * values  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**AuditEntryPaging**](AuditEntryPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuditEntriesForNode

> AuditEntryPaging ListAuditEntriesForNode(ctx, nodeId).SkipCount(skipCount).OrderBy(orderBy).MaxItems(maxItems).Where(where).Include(include).Fields(fields).Execute()

List audit entries for a node



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
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    where := "where_example" // string | Optionally filter the list. Here are some examples:  *   ```where=(createdByUser='-me-')```  *   ```where=(createdAt BETWEEN ('2017-06-02T12:13:51.593+01:00' , '2017-06-04T10:05:16.536+01:00')```  *   ```where=(createdByUser='jbloggs' and createdAt BETWEEN ('2017-06-02T12:13:51.593+01:00' , '2017-06-04T10:05:16.536+01:00')```  (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about the audit entry. The following optional fields can be requested: * values  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.ListAuditEntriesForNode(context.Background(), nodeId).SkipCount(skipCount).OrderBy(orderBy).MaxItems(maxItems).Where(where).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.ListAuditEntriesForNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditEntriesForNode`: AuditEntryPaging
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.ListAuditEntriesForNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuditEntriesForNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **where** | **string** | Optionally filter the list. Here are some examples:  *   &#x60;&#x60;&#x60;where&#x3D;(createdByUser&#x3D;&#39;-me-&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(createdAt BETWEEN (&#39;2017-06-02T12:13:51.593+01:00&#39; , &#39;2017-06-04T10:05:16.536+01:00&#39;)&#x60;&#x60;&#x60;  *   &#x60;&#x60;&#x60;where&#x3D;(createdByUser&#x3D;&#39;jbloggs&#39; and createdAt BETWEEN (&#39;2017-06-02T12:13:51.593+01:00&#39; , &#39;2017-06-04T10:05:16.536+01:00&#39;)&#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about the audit entry. The following optional fields can be requested: * values  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**AuditEntryPaging**](AuditEntryPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuditApp

> AuditApp UpdateAuditApp(ctx, auditApplicationId).AuditAppBodyUpdate(auditAppBodyUpdate).Fields(fields).Execute()

Update audit application info



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
    auditApplicationId := "auditApplicationId_example" // string | The identifier of an audit application.
    auditAppBodyUpdate := *openapiclient.NewAuditBodyUpdate() // AuditBodyUpdate | The audit application to update.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.UpdateAuditApp(context.Background(), auditApplicationId).AuditAppBodyUpdate(auditAppBodyUpdate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.UpdateAuditApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuditApp`: AuditApp
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.UpdateAuditApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditApplicationId** | **string** | The identifier of an audit application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuditAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditAppBodyUpdate** | [**AuditBodyUpdate**](AuditBodyUpdate.md) | The audit application to update. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**AuditApp**](AuditApp.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

