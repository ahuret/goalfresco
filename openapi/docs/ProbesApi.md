# \ProbesApi

All URIs are relative to *http://localhost/alfresco/api/-default-/public/alfresco/versions/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProbe**](ProbesApi.md#GetProbe) | **Get** /probes/{probeId} | Check readiness and liveness of the repository



## GetProbe

> ProbeEntry GetProbe(ctx, probeId).Execute()

Check readiness and liveness of the repository



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
    probeId := "probeId_example" // string | The name of the probe: * -ready- * -live- 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProbesApi.GetProbe(context.Background(), probeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProbesApi.GetProbe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProbe`: ProbeEntry
    fmt.Fprintf(os.Stdout, "Response from `ProbesApi.GetProbe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | **string** | The name of the probe: * -ready- * -live-  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProbeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProbeEntry**](ProbeEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

