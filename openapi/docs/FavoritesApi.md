# \FavoritesApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFavorite**](FavoritesApi.md#CreateFavorite) | **Post** /people/{personId}/favorites | Create a favorite
[**CreateSiteFavorite**](FavoritesApi.md#CreateSiteFavorite) | **Post** /people/{personId}/favorite-sites | Create a site favorite
[**DeleteFavorite**](FavoritesApi.md#DeleteFavorite) | **Delete** /people/{personId}/favorites/{favoriteId} | Delete a favorite
[**DeleteSiteFavorite**](FavoritesApi.md#DeleteSiteFavorite) | **Delete** /people/{personId}/favorite-sites/{siteId} | Delete a site favorite
[**GetFavorite**](FavoritesApi.md#GetFavorite) | **Get** /people/{personId}/favorites/{favoriteId} | Get a favorite
[**GetFavoriteSite**](FavoritesApi.md#GetFavoriteSite) | **Get** /people/{personId}/favorite-sites/{siteId} | Get a favorite site
[**ListFavoriteSitesForPerson**](FavoritesApi.md#ListFavoriteSitesForPerson) | **Get** /people/{personId}/favorite-sites | List favorite sites
[**ListFavorites**](FavoritesApi.md#ListFavorites) | **Get** /people/{personId}/favorites | List favorites



## CreateFavorite

> FavoriteEntry CreateFavorite(ctx, personId).FavoriteBodyCreate(favoriteBodyCreate).Include(include).Fields(fields).Execute()

Create a favorite



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
    favoriteBodyCreate := *openapiclient.NewFavoriteBodyCreate(map[string]interface{}(123)) // FavoriteBodyCreate | An object identifying the entity to be favorited.  The object consists of a single property which is an object with the name `site`, `file`, or `folder`. The content of that object is the `guid` of the target entity.  For example, to favorite a file the following body would be used:  ```JSON {    \"target\": {       \"file\": {          \"guid\": \"abcde-01234-....\"       }    } } ``` 
    include := []string{"Inner_example"} // []string | Returns additional information about favorites, the following optional fields can be requested: * path (note, this only applies to files and folders) * properties  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FavoritesApi.CreateFavorite(context.Background(), personId).FavoriteBodyCreate(favoriteBodyCreate).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.CreateFavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFavorite`: FavoriteEntry
    fmt.Fprintf(os.Stdout, "Response from `FavoritesApi.CreateFavorite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **favoriteBodyCreate** | [**FavoriteBodyCreate**](FavoriteBodyCreate.md) | An object identifying the entity to be favorited.  The object consists of a single property which is an object with the name &#x60;site&#x60;, &#x60;file&#x60;, or &#x60;folder&#x60;. The content of that object is the &#x60;guid&#x60; of the target entity.  For example, to favorite a file the following body would be used:  &#x60;&#x60;&#x60;JSON {    \&quot;target\&quot;: {       \&quot;file\&quot;: {          \&quot;guid\&quot;: \&quot;abcde-01234-....\&quot;       }    } } &#x60;&#x60;&#x60;  | 
 **include** | **[]string** | Returns additional information about favorites, the following optional fields can be requested: * path (note, this only applies to files and folders) * properties  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**FavoriteEntry**](FavoriteEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteFavorite

> FavoriteSiteEntry CreateSiteFavorite(ctx, personId).FavoriteSiteBodyCreate(favoriteSiteBodyCreate).Fields(fields).Execute()

Create a site favorite



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
    favoriteSiteBodyCreate := *openapiclient.NewFavoriteSiteBodyCreate("Id_example") // FavoriteSiteBodyCreate | The id of the site to favorite.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FavoritesApi.CreateSiteFavorite(context.Background(), personId).FavoriteSiteBodyCreate(favoriteSiteBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.CreateSiteFavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiteFavorite`: FavoriteSiteEntry
    fmt.Fprintf(os.Stdout, "Response from `FavoritesApi.CreateSiteFavorite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **favoriteSiteBodyCreate** | [**FavoriteSiteBodyCreate**](FavoriteSiteBodyCreate.md) | The id of the site to favorite. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**FavoriteSiteEntry**](FavoriteSiteEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFavorite

> DeleteFavorite(ctx, personId, favoriteId).Execute()

Delete a favorite



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
    favoriteId := "favoriteId_example" // string | The identifier of a favorite.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FavoritesApi.DeleteFavorite(context.Background(), personId, favoriteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.DeleteFavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 
**favoriteId** | **string** | The identifier of a favorite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFavoriteRequest struct via the builder pattern


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


## DeleteSiteFavorite

> DeleteSiteFavorite(ctx, personId, siteId).Execute()

Delete a site favorite



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
    siteId := "siteId_example" // string | The identifier of a site.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FavoritesApi.DeleteSiteFavorite(context.Background(), personId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.DeleteSiteFavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteFavoriteRequest struct via the builder pattern


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


## GetFavorite

> FavoriteEntry GetFavorite(ctx, personId, favoriteId).Include(include).Fields(fields).Execute()

Get a favorite



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
    favoriteId := "favoriteId_example" // string | The identifier of a favorite.
    include := []string{"Inner_example"} // []string | Returns additional information about favorites, the following optional fields can be requested: * path (note, this only applies to files and folders) * properties  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FavoritesApi.GetFavorite(context.Background(), personId, favoriteId).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.GetFavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFavorite`: FavoriteEntry
    fmt.Fprintf(os.Stdout, "Response from `FavoritesApi.GetFavorite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 
**favoriteId** | **string** | The identifier of a favorite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **[]string** | Returns additional information about favorites, the following optional fields can be requested: * path (note, this only applies to files and folders) * properties  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**FavoriteEntry**](FavoriteEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoriteSite

> SiteEntry GetFavoriteSite(ctx, personId, siteId).Fields(fields).Execute()

Get a favorite site



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
    siteId := "siteId_example" // string | The identifier of a site.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FavoritesApi.GetFavoriteSite(context.Background(), personId, siteId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.GetFavoriteSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFavoriteSite`: SiteEntry
    fmt.Fprintf(os.Stdout, "Response from `FavoritesApi.GetFavoriteSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoriteSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteEntry**](SiteEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFavoriteSitesForPerson

> SitePaging ListFavoriteSitesForPerson(ctx, personId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()

List favorite sites



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
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FavoritesApi.ListFavoriteSitesForPerson(context.Background(), personId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.ListFavoriteSitesForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFavoriteSitesForPerson`: SitePaging
    fmt.Fprintf(os.Stdout, "Response from `FavoritesApi.ListFavoriteSitesForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFavoriteSitesForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
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


## ListFavorites

> FavoritePaging ListFavorites(ctx, personId).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Where(where).Include(include).Fields(fields).Execute()

List favorites



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
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)
    where := "where_example" // string | A string to restrict the returned objects by using a predicate. (optional)
    include := []string{"Inner_example"} // []string | Returns additional information about favorites, the following optional fields can be requested: * path (note, this only applies to files and folders) * properties  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FavoritesApi.ListFavorites(context.Background(), personId).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Where(where).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FavoritesApi.ListFavorites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFavorites`: FavoritePaging
    fmt.Fprintf(os.Stdout, "Response from `FavoritesApi.ListFavorites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **where** | **string** | A string to restrict the returned objects by using a predicate. | 
 **include** | **[]string** | Returns additional information about favorites, the following optional fields can be requested: * path (note, this only applies to files and folders) * properties  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**FavoritePaging**](FavoritePaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

