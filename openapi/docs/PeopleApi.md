# \PeopleApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePerson**](PeopleApi.md#CreatePerson) | **Post** /people | Create person
[**DeleteAvatarImage**](PeopleApi.md#DeleteAvatarImage) | **Delete** /people/{personId}/avatar | Delete avatar image
[**GetAvatarImage**](PeopleApi.md#GetAvatarImage) | **Get** /people/{personId}/avatar | Get avatar image
[**GetPerson**](PeopleApi.md#GetPerson) | **Get** /people/{personId} | Get a person
[**ListPeople**](PeopleApi.md#ListPeople) | **Get** /people | List people
[**RequestPasswordReset**](PeopleApi.md#RequestPasswordReset) | **Post** /people/{personId}/request-password-reset | Request password reset
[**ResetPassword**](PeopleApi.md#ResetPassword) | **Post** /people/{personId}/reset-password | Reset password
[**UpdateAvatarImage**](PeopleApi.md#UpdateAvatarImage) | **Put** /people/{personId}/avatar | Update avatar image
[**UpdatePerson**](PeopleApi.md#UpdatePerson) | **Put** /people/{personId} | Update person



## CreatePerson

> PersonEntry CreatePerson(ctx).PersonBodyCreate(personBodyCreate).Fields(fields).Execute()

Create person



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
    personBodyCreate := *openapiclient.NewPersonBodyCreate("Id_example", "FirstName_example", "Email_example", "Password_example") // PersonBodyCreate | The person details.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.CreatePerson(context.Background()).PersonBodyCreate(personBodyCreate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.CreatePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePerson`: PersonEntry
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.CreatePerson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personBodyCreate** | [**PersonBodyCreate**](PersonBodyCreate.md) | The person details. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**PersonEntry**](PersonEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvatarImage

> DeleteAvatarImage(ctx, personId).Execute()

Delete avatar image



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PeopleApi.DeleteAvatarImage(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.DeleteAvatarImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvatarImageRequest struct via the builder pattern


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


## GetAvatarImage

> *os.File GetAvatarImage(ctx, personId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Placeholder(placeholder).Execute()

Get avatar image



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
    personId := "personId_example" // string | The identifier of a person.
    attachment := true // bool | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  (optional) (default to true)
    ifModifiedSince := time.Now() // time.Time | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, `Wed, 09 Mar 2016 16:56:34 GMT`.  (optional)
    placeholder := true // bool | If **true** and there is no avatar for this **personId** then the placeholder image is returned, rather than a 404 response.  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.GetAvatarImage(context.Background(), personId).Attachment(attachment).IfModifiedSince(ifModifiedSince).Placeholder(placeholder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GetAvatarImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvatarImage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GetAvatarImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachment** | **bool** | **true** enables a web browser to download the file as an attachment. **false** means a web browser may preview the file in a new tab or window, but not download the file.  You can only set this parameter to **false** if the content type of the file is in the supported list; for example, certain image files and PDF files.  If the content type is not supported for preview, then a value of **false**  is ignored, and the attachment will be returned in the response.  | [default to true]
 **ifModifiedSince** | **time.Time** | Only returns the content if it has been modified since the date provided. Use the date format defined by HTTP. For example, &#x60;Wed, 09 Mar 2016 16:56:34 GMT&#x60;.  | 
 **placeholder** | **bool** | If **true** and there is no avatar for this **personId** then the placeholder image is returned, rather than a 404 response.  | [default to true]

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


## GetPerson

> PersonEntry GetPerson(ctx, personId).Fields(fields).Execute()

Get a person



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
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.GetPerson(context.Background(), personId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.GetPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPerson`: PersonEntry
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.GetPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**PersonEntry**](PersonEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPeople

> PersonPaging ListPeople(ctx).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Include(include).Fields(fields).Execute()

List people



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
    include := []string{"Inner_example"} // []string | Returns additional information about the person. The following optional fields can be requested: * properties * aspectNames * capabilities  (optional)
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.ListPeople(context.Background()).SkipCount(skipCount).MaxItems(maxItems).OrderBy(orderBy).Include(include).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.ListPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPeople`: PersonPaging
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.ListPeople`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPeopleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipCount** | **int32** | The number of entities that exist in the collection before those included in this list. If not supplied then the default value is 0.  | [default to 0]
 **maxItems** | **int32** | The maximum number of items to return in the list. If not supplied then the default value is 100.  | [default to 100]
 **orderBy** | **[]string** | A string to control the order of the entities returned in a list. You can use the **orderBy** parameter to sort the list by one or more fields.  Each field has a default sort order, which is normally ascending order. Read the API method implementation notes above to check if any fields used in this method have a descending default search order.  To sort the entities in a specific order, you can use the **ASC** and **DESC** keywords for any field.  | 
 **include** | **[]string** | Returns additional information about the person. The following optional fields can be requested: * properties * aspectNames * capabilities  | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

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


## RequestPasswordReset

> RequestPasswordReset(ctx, personId).ClientBody(clientBody).Execute()

Request password reset



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
    clientBody := *openapiclient.NewClientBody("Client_example") // ClientBody | The client name to send email with app-specific url.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PeopleApi.RequestPasswordReset(context.Background(), personId).ClientBody(clientBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.RequestPasswordReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientBody** | [**ClientBody**](ClientBody.md) | The client name to send email with app-specific url. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword(ctx, personId).PasswordResetBody(passwordResetBody).Execute()

Reset password



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
    passwordResetBody := *openapiclient.NewPasswordResetBody("Password_example", "Id_example", "Key_example") // PasswordResetBody | The reset password details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PeopleApi.ResetPassword(context.Background(), personId).PasswordResetBody(passwordResetBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.ResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordResetBody** | [**PasswordResetBody**](PasswordResetBody.md) | The reset password details | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAvatarImage

> UpdateAvatarImage(ctx, personId).ContentBodyUpdate(contentBodyUpdate).Execute()

Update avatar image



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
    contentBodyUpdate := os.NewFile(1234, "some_file") // *os.File | The binary content

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PeopleApi.UpdateAvatarImage(context.Background(), personId).ContentBodyUpdate(contentBodyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.UpdateAvatarImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvatarImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentBodyUpdate** | ***os.File** | The binary content | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePerson

> PersonEntry UpdatePerson(ctx, personId).PersonBodyUpdate(personBodyUpdate).Fields(fields).Execute()

Update person



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
    personBodyUpdate := *openapiclient.NewPersonBodyUpdate() // PersonBodyUpdate | The person details.
    fields := []string{"Inner_example"} // []string | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.UpdatePerson(context.Background(), personId).PersonBodyUpdate(personBodyUpdate).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.UpdatePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePerson`: PersonEntry
    fmt.Fprintf(os.Stdout, "Response from `PeopleApi.UpdatePerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | The identifier of a person. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **personBodyUpdate** | [**PersonBodyUpdate**](PersonBodyUpdate.md) | The person details. | 
 **fields** | **[]string** | A list of field names.  You can use this parameter to restrict the fields returned within a response if, for example, you want to save on overall bandwidth.  The list applies to a returned individual entity or entries within a collection.  If the API method also supports the **include** parameter, then the fields specified in the **include** parameter are returned in addition to those specified in the **fields** parameter.  | 

### Return type

[**PersonEntry**](PersonEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

