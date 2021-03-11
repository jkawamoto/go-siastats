# \NavigatorApi

All URIs are relative to *https://siastats.info:3500/navigator-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Hash**](NavigatorApi.md#Hash) | **Get** /hash/{hash} | Returns the information about any hash (address, Tx, contract...) or block height on the blockchain
[**Status**](NavigatorApi.md#Status) | **Get** /status | Global status and blockchain sync situation of SiaStats nodes



## Hash

> []HashResponse Hash(ctx, hash).Execute()

Returns the information about any hash (address, Tx, contract...) or block height on the blockchain

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
    hash := "hash_example" // string | hash or block height

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NavigatorApi.Hash(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NavigatorApi.Hash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Hash`: []HashResponse
    fmt.Fprintf(os.Stdout, "Response from `NavigatorApi.Hash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | hash or block height | 

### Other Parameters

Other parameters are passed through a pointer to a apiHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HashResponse**](HashResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> []Status Status(ctx).Execute()

Global status and blockchain sync situation of SiaStats nodes

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NavigatorApi.Status(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NavigatorApi.Status``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status`: []Status
    fmt.Fprintf(os.Stdout, "Response from `NavigatorApi.Status`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


### Return type

[**[]Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

