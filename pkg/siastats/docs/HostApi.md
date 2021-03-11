# \HostApi

All URIs are relative to *https://siastats.info:3500/navigator-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllHosts**](HostApi.md#AllHosts) | **Post** /allhosts | Summary information and SiaStats final scores of all the hosts in the Sia network.
[**Host**](HostApi.md#Host) | **Get** /host/{id} | Information and detailed SiaStats benchmarks about a host.



## AllHosts

> []Host AllHosts(ctx).AllHostsRequest(allHostsRequest).Execute()

Summary information and SiaStats final scores of all the hosts in the Sia network.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allHostsRequest := *openapiclient.NewAllHostsRequest("Network_example", "List_example") // AllHostsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostApi.AllHosts(context.Background()).AllHostsRequest(allHostsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostApi.AllHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllHosts`: []Host
    fmt.Fprintf(os.Stdout, "Response from `HostApi.AllHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allHostsRequest** | [**AllHostsRequest**](AllHostsRequest.md) |  | 

### Return type

[**[]Host**](Host.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Host

> HostResponse Host(ctx, id).Execute()

Information and detailed SiaStats benchmarks about a host.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | SiaStats host ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostApi.Host(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostApi.Host``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Host`: HostResponse
    fmt.Fprintf(os.Stdout, "Response from `HostApi.Host`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | SiaStats host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostResponse**](HostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

