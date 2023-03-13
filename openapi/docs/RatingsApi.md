# \RatingsApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRating**](RatingsApi.md#CreateRating) | **Post** /nodes/{nodeId}/ratings | Create a rating
[**DeleteRating**](RatingsApi.md#DeleteRating) | **Delete** /nodes/{nodeId}/ratings/{ratingId} | Delete a rating
[**GetRating**](RatingsApi.md#GetRating) | **Get** /nodes/{nodeId}/ratings/{ratingId} | Get a rating
[**ListRatings**](RatingsApi.md#ListRatings) | **Get** /nodes/{nodeId}/ratings | List ratings



## CreateRating

> RatingEntry CreateRating(ctx, nodeId).RatingBodyCreate(ratingBodyCreate).Fields(fields).Execute()

Create a rating



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
    ratingBodyCreate := *openapiclient.NewRatingBody("Id_example", "MyRating_example") // RatingBody | For \"myRating\" the type is specific to the rating scheme, boolean for the likes and an integer for the fiveStar.  For example, to \"like\" a file the following body would be used:  ```JSON   {     \"id\": \"likes\",     \"myRating\": true   } ``` 
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RatingsApi.CreateRating(context.Background(), nodeId).RatingBodyCreate(ratingBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RatingsApi.CreateRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRating`: RatingEntry
    fmt.Fprintf(os.Stdout, "Response from `RatingsApi.CreateRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ratingBodyCreate** | [**RatingBody**](RatingBody.md) | For \&quot;myRating\&quot; the type is specific to the rating scheme, boolean for the likes and an integer for the fiveStar.  For example, to \&quot;like\&quot; a file the following body would be used:  &#x60;&#x60;&#x60;JSON   {     \&quot;id\&quot;: \&quot;likes\&quot;,     \&quot;myRating\&quot;: true   } &#x60;&#x60;&#x60;  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**RatingEntry**](RatingEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRating

> DeleteRating(ctx, nodeId, ratingId).Execute()

Delete a rating



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
    ratingId := "ratingId_example" // string | The identifier of a rating.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RatingsApi.DeleteRating(context.Background(), nodeId, ratingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RatingsApi.DeleteRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**ratingId** | **string** | The identifier of a rating. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRatingRequest struct via the builder pattern


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


## GetRating

> RatingEntry GetRating(ctx, nodeId, ratingId).Fields(fields).Execute()

Get a rating



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
    ratingId := "ratingId_example" // string | The identifier of a rating.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RatingsApi.GetRating(context.Background(), nodeId, ratingId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RatingsApi.GetRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRating`: RatingEntry
    fmt.Fprintf(os.Stdout, "Response from `RatingsApi.GetRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 
**ratingId** | **string** | The identifier of a rating. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**RatingEntry**](RatingEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRatings

> RatingPaging ListRatings(ctx, nodeId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()

List ratings



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
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RatingsApi.ListRatings(context.Background(), nodeId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RatingsApi.ListRatings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRatings`: RatingPaging
    fmt.Fprintf(os.Stdout, "Response from `RatingsApi.ListRatings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The identifier of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**RatingPaging**](RatingPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

