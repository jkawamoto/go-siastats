# \HostApi

All URIs are relative to *https://siastats.info:3500/navigator-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllHosts**](HostApi.md#AllHosts) | **Post** /allhosts | Summary information and SiaStats final scores of all the hosts in the Sia network.
[**Host**](HostApi.md#Host) | **Get** /host/{id} | Information and detailed SiaStats benchmarks about a host.



## AllHosts

> []Host AllHosts(ctx, allHostsRequest)
Summary information and SiaStats final scores of all the hosts in the Sia network.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allHostsRequest** | [**AllHostsRequest**](AllHostsRequest.md)|  | 

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

> HostResponse Host(ctx, id)
Information and detailed SiaStats benchmarks about a host.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| SiaStats host ID | 

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

