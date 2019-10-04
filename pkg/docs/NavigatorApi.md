# \NavigatorApi

All URIs are relative to *https://siastats.info:3500/navigator-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Hash**](NavigatorApi.md#Hash) | **Get** /hash/{hash} | Returns the information about any hash (address, Tx, contract...) or block height on the blockchain
[**Status**](NavigatorApi.md#Status) | **Get** /status | Global status and blockchain sync situation of SiaStats nodes



## Hash

> []HashResponse Hash(ctx, hash)
Returns the information about any hash (address, Tx, contract...) or block height on the blockchain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string**| hash or block height | 

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

> []Status Status(ctx, )
Global status and blockchain sync situation of SiaStats nodes

### Required Parameters

This endpoint does not need any parameter.

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

