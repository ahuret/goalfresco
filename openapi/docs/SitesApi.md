# \SitesApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveSiteMembershipRequest**](SitesApi.md#ApproveSiteMembershipRequest) | **Post** /sites/{siteId}/site-membership-requests/{inviteeId}/approve | Approve a site membership request
[**CreateSite**](SitesApi.md#CreateSite) | **Post** /sites | Create a site
[**CreateSiteGroupMembership**](SitesApi.md#CreateSiteGroupMembership) | **Post** /sites/{siteId}/group-members | Create a site membership for group
[**CreateSiteMembership**](SitesApi.md#CreateSiteMembership) | **Post** /sites/{siteId}/members | Create a site membership
[**CreateSiteMembershipRequestForPerson**](SitesApi.md#CreateSiteMembershipRequestForPerson) | **Post** /people/{personId}/site-membership-requests | Create a site membership request
[**DeleteSite**](SitesApi.md#DeleteSite) | **Delete** /sites/{siteId} | Delete a site
[**DeleteSiteGroupMembership**](SitesApi.md#DeleteSiteGroupMembership) | **Delete** /sites/{siteId}/group-members/{groupId} | Delete a group membership for site
[**DeleteSiteMembership**](SitesApi.md#DeleteSiteMembership) | **Delete** /sites/{siteId}/members/{personId} | Delete a site membership
[**DeleteSiteMembershipForPerson**](SitesApi.md#DeleteSiteMembershipForPerson) | **Delete** /people/{personId}/sites/{siteId} | Delete a site membership
[**DeleteSiteMembershipRequestForPerson**](SitesApi.md#DeleteSiteMembershipRequestForPerson) | **Delete** /people/{personId}/site-membership-requests/{siteId} | Delete a site membership request
[**GetSite**](SitesApi.md#GetSite) | **Get** /sites/{siteId} | Get a site
[**GetSiteContainer**](SitesApi.md#GetSiteContainer) | **Get** /sites/{siteId}/containers/{containerId} | Get a site container
[**GetSiteGroupMembership**](SitesApi.md#GetSiteGroupMembership) | **Get** /sites/{siteId}/group-members/{groupId} | Get information about site membership of group
[**GetSiteMembership**](SitesApi.md#GetSiteMembership) | **Get** /sites/{siteId}/members/{personId} | Get a site membership
[**GetSiteMembershipForPerson**](SitesApi.md#GetSiteMembershipForPerson) | **Get** /people/{personId}/sites/{siteId} | Get a site membership
[**GetSiteMembershipRequestForPerson**](SitesApi.md#GetSiteMembershipRequestForPerson) | **Get** /people/{personId}/site-membership-requests/{siteId} | Get a site membership request
[**GetSiteMembershipRequests**](SitesApi.md#GetSiteMembershipRequests) | **Get** /site-membership-requests | Get site membership requests
[**ListSiteContainers**](SitesApi.md#ListSiteContainers) | **Get** /sites/{siteId}/containers | List site containers
[**ListSiteGroups**](SitesApi.md#ListSiteGroups) | **Get** /sites/{siteId}/group-members | List group membership for site
[**ListSiteMembershipRequestsForPerson**](SitesApi.md#ListSiteMembershipRequestsForPerson) | **Get** /people/{personId}/site-membership-requests | List site membership requests
[**ListSiteMemberships**](SitesApi.md#ListSiteMemberships) | **Get** /sites/{siteId}/members | List site memberships
[**ListSiteMembershipsForPerson**](SitesApi.md#ListSiteMembershipsForPerson) | **Get** /people/{personId}/sites | List site memberships
[**ListSites**](SitesApi.md#ListSites) | **Get** /sites | List sites
[**RejectSiteMembershipRequest**](SitesApi.md#RejectSiteMembershipRequest) | **Post** /sites/{siteId}/site-membership-requests/{inviteeId}/reject | Reject a site membership request
[**UpdateSite**](SitesApi.md#UpdateSite) | **Put** /sites/{siteId} | Update a site
[**UpdateSiteGroupMembership**](SitesApi.md#UpdateSiteGroupMembership) | **Put** /sites/{siteId}/group-members/{groupId} | Update site membership of group
[**UpdateSiteMembership**](SitesApi.md#UpdateSiteMembership) | **Put** /sites/{siteId}/members/{personId} | Update a site membership
[**UpdateSiteMembershipRequestForPerson**](SitesApi.md#UpdateSiteMembershipRequestForPerson) | **Put** /people/{personId}/site-membership-requests/{siteId} | Update a site membership request



## ApproveSiteMembershipRequest

> ApproveSiteMembershipRequest(ctx, siteId, inviteeId).SiteMembershipApprovalBody(siteMembershipApprovalBody).Execute()

Approve a site membership request



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
    siteId := "siteId_example" // string | The identifier of a site.
    inviteeId := "inviteeId_example" // string | The invitee user name.
    siteMembershipApprovalBody := *openapiclient.NewSiteMembershipApprovalBody() // SiteMembershipApprovalBody | Accepting a request to join, optionally, allows assignment of a role to the user.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesApi.ApproveSiteMembershipRequest(context.Background(), siteId, inviteeId).SiteMembershipApprovalBody(siteMembershipApprovalBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.ApproveSiteMembershipRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**inviteeId** | **string** | The invitee user name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveSiteMembershipRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteMembershipApprovalBody** | [**SiteMembershipApprovalBody**](SiteMembershipApprovalBody.md) | Accepting a request to join, optionally, allows assignment of a role to the user.  | 

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


## CreateSite

> SiteEntry CreateSite(ctx).SiteBodyCreate(siteBodyCreate).SkipConfiguration(skipConfiguration).SkipAddToFavorites(skipAddToFavorites).Fields(fields).Execute()

Create a site



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
    siteBodyCreate := *openapiclient.NewSiteBodyCreate("Title_example", "Visibility_example") // SiteBodyCreate | The site details
    skipConfiguration := true // bool | Flag to indicate whether the Share-specific (surf) configuration files for the site should not be created. (optional) (default to false)
    skipAddToFavorites := true // bool | Flag to indicate whether the site should not be added to the user's site favorites. (optional) (default to false)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.CreateSite(context.Background()).SiteBodyCreate(siteBodyCreate).SkipConfiguration(skipConfiguration).SkipAddToFavorites(skipAddToFavorites).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.CreateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSite`: SiteEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.CreateSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteBodyCreate** | [**SiteBodyCreate**](SiteBodyCreate.md) | The site details | 
 **skipConfiguration** | **bool** | Flag to indicate whether the Share-specific (surf) configuration files for the site should not be created. | [default to false]
 **skipAddToFavorites** | **bool** | Flag to indicate whether the site should not be added to the user&#39;s site favorites. | [default to false]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteEntry**](SiteEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteGroupMembership

> SiteGroupEntry CreateSiteGroupMembership(ctx, siteId).SiteMembershipBodyCreate(siteMembershipBodyCreate).Fields(fields).Execute()

Create a site membership for group



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
    siteId := "siteId_example" // string | The identifier of a site.
    siteMembershipBodyCreate := *openapiclient.NewSiteMembershipBodyCreate("Role_example", "Id_example") // SiteMembershipBodyCreate | The group to add and their role
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.CreateSiteGroupMembership(context.Background(), siteId).SiteMembershipBodyCreate(siteMembershipBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.CreateSiteGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiteGroupMembership`: SiteGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.CreateSiteGroupMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteGroupMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteMembershipBodyCreate** | [**SiteMembershipBodyCreate**](SiteMembershipBodyCreate.md) | The group to add and their role | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteGroupEntry**](SiteGroupEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteMembership

> SiteMemberEntry CreateSiteMembership(ctx, siteId).SiteMembershipBodyCreate(siteMembershipBodyCreate).Fields(fields).Execute()

Create a site membership



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
    siteId := "siteId_example" // string | The identifier of a site.
    siteMembershipBodyCreate := *openapiclient.NewSiteMembershipBodyCreate("Role_example", "Id_example") // SiteMembershipBodyCreate | The person to add and their role
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.CreateSiteMembership(context.Background(), siteId).SiteMembershipBodyCreate(siteMembershipBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.CreateSiteMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiteMembership`: SiteMemberEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.CreateSiteMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteMembershipBodyCreate** | [**SiteMembershipBodyCreate**](SiteMembershipBodyCreate.md) | The person to add and their role | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMemberEntry**](SiteMemberEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteMembershipRequestForPerson

> SiteMembershipRequestEntry CreateSiteMembershipRequestForPerson(ctx, personId).SiteMembershipRequestBodyCreate(siteMembershipRequestBodyCreate).Fields(fields).Execute()

Create a site membership request



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
    siteMembershipRequestBodyCreate := *openapiclient.NewSiteMembershipRequestBodyCreate("Id_example") // SiteMembershipRequestBodyCreate | Site membership request details
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.CreateSiteMembershipRequestForPerson(context.Background(), personId).SiteMembershipRequestBodyCreate(siteMembershipRequestBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.CreateSiteMembershipRequestForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSiteMembershipRequestForPerson`: SiteMembershipRequestEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.CreateSiteMembershipRequestForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteMembershipRequestForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteMembershipRequestBodyCreate** | [**SiteMembershipRequestBodyCreate**](SiteMembershipRequestBodyCreate.md) | Site membership request details | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMembershipRequestEntry**](SiteMembershipRequestEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSite

> DeleteSite(ctx, siteId).Permanent(permanent).Execute()

Delete a site



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
    siteId := "siteId_example" // string | The identifier of a site.
    permanent := true // bool | Flag to indicate whether the site should be permanently deleted i.e. bypass the trashcan. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesApi.DeleteSite(context.Background(), siteId).Permanent(permanent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.DeleteSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permanent** | **bool** | Flag to indicate whether the site should be permanently deleted i.e. bypass the trashcan. | [default to false]

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


## DeleteSiteGroupMembership

> DeleteSiteGroupMembership(ctx, siteId, groupId).Execute()

Delete a group membership for site



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
    siteId := "siteId_example" // string | The identifier of a site.
    groupId := "groupId_example" // string | The identifier of a group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesApi.DeleteSiteGroupMembership(context.Background(), siteId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.DeleteSiteGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**groupId** | **string** | The identifier of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteGroupMembershipRequest struct via the builder pattern


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


## DeleteSiteMembership

> DeleteSiteMembership(ctx, siteId, personId).Execute()

Delete a site membership



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
    siteId := "siteId_example" // string | The identifier of a site.
    personId := "personId_example" // string | The identifier of a person.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesApi.DeleteSiteMembership(context.Background(), siteId, personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.DeleteSiteMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteMembershipRequest struct via the builder pattern


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


## DeleteSiteMembershipForPerson

> DeleteSiteMembershipForPerson(ctx, personId, siteId).Execute()

Delete a site membership



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
    r, err := apiClient.SitesApi.DeleteSiteMembershipForPerson(context.Background(), personId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.DeleteSiteMembershipForPerson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteMembershipForPersonRequest struct via the builder pattern


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


## DeleteSiteMembershipRequestForPerson

> DeleteSiteMembershipRequestForPerson(ctx, personId, siteId).Execute()

Delete a site membership request



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
    r, err := apiClient.SitesApi.DeleteSiteMembershipRequestForPerson(context.Background(), personId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.DeleteSiteMembershipRequestForPerson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteMembershipRequestForPersonRequest struct via the builder pattern


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


## GetSite

> SiteEntry GetSite(ctx, siteId).Relations(relations).Fields(fields).Execute()

Get a site



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
    siteId := "siteId_example" // string | The identifier of a site.
    relations := []string{"Inner_example"} // []string | Use the relations parameter to include one or more related entities in a single response. (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.GetSite(context.Background(), siteId).Relations(relations).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSite`: SiteEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relations** | **[]string** | Use the relations parameter to include one or more related entities in a single response. | 
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


## GetSiteContainer

> SiteContainerEntry GetSiteContainer(ctx, siteId, containerId).Fields(fields).Execute()

Get a site container



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
    siteId := "siteId_example" // string | The identifier of a site.
    containerId := "containerId_example" // string | The unique identifier of a site container.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.GetSiteContainer(context.Background(), siteId, containerId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSiteContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteContainer`: SiteContainerEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSiteContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**containerId** | **string** | The unique identifier of a site container. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteContainerEntry**](SiteContainerEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteGroupMembership

> SiteGroupEntry GetSiteGroupMembership(ctx, siteId, groupId).Fields(fields).Execute()

Get information about site membership of group



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
    siteId := "siteId_example" // string | The identifier of a site.
    groupId := "groupId_example" // string | The identifier of a group.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.GetSiteGroupMembership(context.Background(), siteId, groupId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSiteGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteGroupMembership`: SiteGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSiteGroupMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**groupId** | **string** | The identifier of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteGroupMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteGroupEntry**](SiteGroupEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteMembership

> SiteMemberEntry GetSiteMembership(ctx, siteId, personId).Fields(fields).Execute()

Get a site membership



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
    siteId := "siteId_example" // string | The identifier of a site.
    personId := "personId_example" // string | The identifier of a person.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.GetSiteMembership(context.Background(), siteId, personId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSiteMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteMembership`: SiteMemberEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSiteMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMemberEntry**](SiteMemberEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteMembershipForPerson

> SiteRoleEntry GetSiteMembershipForPerson(ctx, personId, siteId).Execute()

Get a site membership



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
    resp, r, err := apiClient.SitesApi.GetSiteMembershipForPerson(context.Background(), personId, siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSiteMembershipForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteMembershipForPerson`: SiteRoleEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSiteMembershipForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMembershipForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SiteRoleEntry**](SiteRoleEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteMembershipRequestForPerson

> SiteMembershipRequestEntry GetSiteMembershipRequestForPerson(ctx, personId, siteId).Fields(fields).Execute()

Get a site membership request



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
    resp, r, err := apiClient.SitesApi.GetSiteMembershipRequestForPerson(context.Background(), personId, siteId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSiteMembershipRequestForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteMembershipRequestForPerson`: SiteMembershipRequestEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSiteMembershipRequestForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMembershipRequestForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMembershipRequestEntry**](SiteMembershipRequestEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteMembershipRequests

> SiteMembershipRequestWithPersonPaging GetSiteMembershipRequests(ctx).SkipCount(skipCount).MaxItems(maxItems).Where(where).Fields(fields).Execute()

Get site membership requests



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
    where := "where_example" // string | A string to restrict the returned objects by using a predicate. (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.GetSiteMembershipRequests(context.Background()).SkipCount(skipCount).MaxItems(maxItems).Where(where).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSiteMembershipRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteMembershipRequests`: SiteMembershipRequestWithPersonPaging
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSiteMembershipRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMembershipRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **where** | **string** | A string to restrict the returned objects by using a predicate. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMembershipRequestWithPersonPaging**](SiteMembershipRequestWithPersonPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteContainers

> SiteContainerPaging ListSiteContainers(ctx, siteId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()

List site containers



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
    siteId := "siteId_example" // string | The identifier of a site.
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.ListSiteContainers(context.Background(), siteId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.ListSiteContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSiteContainers`: SiteContainerPaging
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.ListSiteContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteContainerPaging**](SiteContainerPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteGroups

> SiteGroupPaging ListSiteGroups(ctx, siteId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()

List group membership for site



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
    siteId := "siteId_example" // string | The identifier of a site.
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.ListSiteGroups(context.Background(), siteId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.ListSiteGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSiteGroups`: SiteGroupPaging
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.ListSiteGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteGroupPaging**](SiteGroupPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteMembershipRequestsForPerson

> SiteMembershipRequestPaging ListSiteMembershipRequestsForPerson(ctx, personId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()

List site membership requests



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
    resp, r, err := apiClient.SitesApi.ListSiteMembershipRequestsForPerson(context.Background(), personId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.ListSiteMembershipRequestsForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSiteMembershipRequestsForPerson`: SiteMembershipRequestPaging
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.ListSiteMembershipRequestsForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteMembershipRequestsForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMembershipRequestPaging**](SiteMembershipRequestPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteMemberships

> SiteMemberPaging ListSiteMemberships(ctx, siteId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Where(where).Execute()

List site memberships



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
    siteId := "siteId_example" // string | The identifier of a site.
    skipCount := int32(56) // int32 | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  (optional) (default to 0)
    maxItems := int32(56) // int32 | The maximum number of items to return in the list. If not supplied then the default value is 100.  (optional) (default to 100)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)
    where := "where_example" // string | Optionally filter the list. *   ```where=(isMemberOfGroup=false|true)```  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.ListSiteMemberships(context.Background(), siteId).SkipCount(skipCount).MaxItems(maxItems).Fields(fields).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.ListSiteMemberships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSiteMemberships`: SiteMemberPaging
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.ListSiteMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 
 **where** | **string** | Optionally filter the list. *   &#x60;&#x60;&#x60;where&#x3D;(isMemberOfGroup&#x3D;false|true)&#x60;&#x60;&#x60;  | 

### Return type

[**SiteMemberPaging**](SiteMemberPaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteMembershipsForPerson

> SiteRolePaging ListSiteMembershipsForPerson(ctx, personId).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Relations(relations).Fields(fields).Where(where).Execute()

List site memberships



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
    relations := []string{"Inner_example"} // []string | Use the relations parameter to include one or more related entities in a single response. (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)
    where := "where_example" // string | A string to restrict the returned objects by using a predicate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.ListSiteMembershipsForPerson(context.Background(), personId).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Relations(relations).Fields(fields).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.ListSiteMembershipsForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSiteMembershipsForPerson`: SiteRolePaging
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.ListSiteMembershipsForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteMembershipsForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **relations** | **[]string** | Use the relations parameter to include one or more related entities in a single response. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 
 **where** | **string** | A string to restrict the returned objects by using a predicate. | 

### Return type

[**SiteRolePaging**](SiteRolePaging.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSites

> SitePaging ListSites(ctx).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Relations(relations).Fields(fields).Where(where).Execute()

List sites



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
    orderBy := []string{"Inner_example"} // []string | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  (optional)
    relations := []string{"Inner_example"} // []string | Use the relations parameter to include one or more related entities in a single response. (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)
    where := "where_example" // string | A string to restrict the returned objects by using a predicate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.ListSites(context.Background()).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Relations(relations).Fields(fields).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.ListSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSites`: SitePaging
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.ListSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **relations** | **[]string** | Use the relations parameter to include one or more related entities in a single response. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 
 **where** | **string** | A string to restrict the returned objects by using a predicate. | 

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


## RejectSiteMembershipRequest

> RejectSiteMembershipRequest(ctx, siteId, inviteeId).SiteMembershipRejectionBody(siteMembershipRejectionBody).Execute()

Reject a site membership request



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
    siteId := "siteId_example" // string | The identifier of a site.
    inviteeId := "inviteeId_example" // string | The invitee user name.
    siteMembershipRejectionBody := *openapiclient.NewSiteMembershipRejectionBody() // SiteMembershipRejectionBody | Rejecting a request to join, optionally, allows the inclusion of comment.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SitesApi.RejectSiteMembershipRequest(context.Background(), siteId, inviteeId).SiteMembershipRejectionBody(siteMembershipRejectionBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.RejectSiteMembershipRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**inviteeId** | **string** | The invitee user name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectSiteMembershipRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteMembershipRejectionBody** | [**SiteMembershipRejectionBody**](SiteMembershipRejectionBody.md) | Rejecting a request to join, optionally, allows the inclusion of comment.  | 

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


## UpdateSite

> SiteEntry UpdateSite(ctx, siteId).SiteBodyUpdate(siteBodyUpdate).Fields(fields).Execute()

Update a site



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
    siteId := "siteId_example" // string | The identifier of a site.
    siteBodyUpdate := *openapiclient.NewSiteBodyUpdate() // SiteBodyUpdate | The site information to update.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.UpdateSite(context.Background(), siteId).SiteBodyUpdate(siteBodyUpdate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.UpdateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSite`: SiteEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.UpdateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteBodyUpdate** | [**SiteBodyUpdate**](SiteBodyUpdate.md) | The site information to update. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteEntry**](SiteEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteGroupMembership

> SiteGroupEntry UpdateSiteGroupMembership(ctx, siteId, groupId).SiteMembershipBodyUpdate(siteMembershipBodyUpdate).Fields(fields).Execute()

Update site membership of group



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
    siteId := "siteId_example" // string | The identifier of a site.
    groupId := "groupId_example" // string | The identifier of a group.
    siteMembershipBodyUpdate := *openapiclient.NewSiteMembershipBodyUpdate("Role_example") // SiteMembershipBodyUpdate | The groupId new role
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.UpdateSiteGroupMembership(context.Background(), siteId, groupId).SiteMembershipBodyUpdate(siteMembershipBodyUpdate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.UpdateSiteGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiteGroupMembership`: SiteGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.UpdateSiteGroupMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**groupId** | **string** | The identifier of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteGroupMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteMembershipBodyUpdate** | [**SiteMembershipBodyUpdate**](SiteMembershipBodyUpdate.md) | The groupId new role | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteGroupEntry**](SiteGroupEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteMembership

> SiteMemberEntry UpdateSiteMembership(ctx, siteId, personId).SiteMembershipBodyUpdate(siteMembershipBodyUpdate).Fields(fields).Execute()

Update a site membership



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
    siteId := "siteId_example" // string | The identifier of a site.
    personId := "personId_example" // string | The identifier of a person.
    siteMembershipBodyUpdate := *openapiclient.NewSiteMembershipBodyUpdate("Role_example") // SiteMembershipBodyUpdate | The persons new role
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.UpdateSiteMembership(context.Background(), siteId, personId).SiteMembershipBodyUpdate(siteMembershipBodyUpdate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.UpdateSiteMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiteMembership`: SiteMemberEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.UpdateSiteMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | The identifier of a site. | 
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteMembershipBodyUpdate** | [**SiteMembershipBodyUpdate**](SiteMembershipBodyUpdate.md) | The persons new role | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMemberEntry**](SiteMemberEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteMembershipRequestForPerson

> SiteMembershipRequestEntry UpdateSiteMembershipRequestForPerson(ctx, personId, siteId).SiteMembershipRequestBodyUpdate(siteMembershipRequestBodyUpdate).Fields(fields).Execute()

Update a site membership request



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
    siteMembershipRequestBodyUpdate := *openapiclient.NewSiteMembershipRequestBodyUpdate() // SiteMembershipRequestBodyUpdate | The new message to display
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.UpdateSiteMembershipRequestForPerson(context.Background(), personId, siteId).SiteMembershipRequestBodyUpdate(siteMembershipRequestBodyUpdate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.UpdateSiteMembershipRequestForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSiteMembershipRequestForPerson`: SiteMembershipRequestEntry
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.UpdateSiteMembershipRequestForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 
**siteId** | **string** | The identifier of a site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteMembershipRequestForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteMembershipRequestBodyUpdate** | [**SiteMembershipRequestBodyUpdate**](SiteMembershipRequestBodyUpdate.md) | The new message to display | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**SiteMembershipRequestEntry**](SiteMembershipRequestEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

