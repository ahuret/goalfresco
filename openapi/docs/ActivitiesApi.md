# \ActivitiesApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListActivitiesForPerson**](ActivitiesApi.md#ListActivitiesForPerson) | **Get** /people/{personId}/activities | List activities



## ListActivitiesForPerson

> ActivityPaging ListActivitiesForPerson(ctx, personId).SkipCount(skipCount).MaxItems(maxItems).Who(who).SiteId(siteId).Fields(fields).Execute()

List activities



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
    personId := "personId_example" // string | The identifier of a person.
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    who := "who_example" // string | A filter to include the user's activities only `me`, other user's activities only `others`'  (optional)
    siteId := "siteId_example" // string | Include only activity feed entries relating to this site. (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivitiesApi.ListActivitiesForPerson(context.Background(), personId).SkipCount(skipCount).MaxItems(maxItems).Who(who).SiteId(siteId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesApi.ListActivitiesForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActivitiesForPerson`: ActivityPaging
    fmt.Fprintf(os.Stdout, "Response from `ActivitiesApi.ListActivitiesForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActivitiesForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **who** | **string** | A filter to include the user&#39;s activities only &#x60;me&#x60;, other user&#39;s activities only &#x60;others&#x60;&#39;  | 
 **siteId** | **string** | Include only activity feed entries relating to this site. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**ActivityPaging**](ActivityPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

