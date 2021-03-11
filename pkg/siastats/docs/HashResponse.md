# HashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**MasterHash** | Pointer to **string** |  | [optional] 
**BalanceSc** | Pointer to **float64** |  | [optional] 
**ReceivedSc** | Pointer to **float64** |  | [optional] 
**SentSc** | Pointer to **float64** |  | [optional] 
**BalanceSf** | Pointer to **float32** |  | [optional] 
**TotalTxCount** | Pointer to **int32** |  | [optional] 
**FirstSeen** | Pointer to **int32** |  | [optional] 
**Last100Transactions** | Pointer to [**[]Transaction**](Transaction.md) |  | [optional] 

## Methods

### NewHashResponse

`func NewHashResponse() *HashResponse`

NewHashResponse instantiates a new HashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashResponseWithDefaults

`func NewHashResponseWithDefaults() *HashResponse`

NewHashResponseWithDefaults instantiates a new HashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HashResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HashResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HashResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HashResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMasterHash

`func (o *HashResponse) GetMasterHash() string`

GetMasterHash returns the MasterHash field if non-nil, zero value otherwise.

### GetMasterHashOk

`func (o *HashResponse) GetMasterHashOk() (*string, bool)`

GetMasterHashOk returns a tuple with the MasterHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterHash

`func (o *HashResponse) SetMasterHash(v string)`

SetMasterHash sets MasterHash field to given value.

### HasMasterHash

`func (o *HashResponse) HasMasterHash() bool`

HasMasterHash returns a boolean if a field has been set.

### GetBalanceSc

`func (o *HashResponse) GetBalanceSc() float64`

GetBalanceSc returns the BalanceSc field if non-nil, zero value otherwise.

### GetBalanceScOk

`func (o *HashResponse) GetBalanceScOk() (*float64, bool)`

GetBalanceScOk returns a tuple with the BalanceSc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSc

`func (o *HashResponse) SetBalanceSc(v float64)`

SetBalanceSc sets BalanceSc field to given value.

### HasBalanceSc

`func (o *HashResponse) HasBalanceSc() bool`

HasBalanceSc returns a boolean if a field has been set.

### GetReceivedSc

`func (o *HashResponse) GetReceivedSc() float64`

GetReceivedSc returns the ReceivedSc field if non-nil, zero value otherwise.

### GetReceivedScOk

`func (o *HashResponse) GetReceivedScOk() (*float64, bool)`

GetReceivedScOk returns a tuple with the ReceivedSc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedSc

`func (o *HashResponse) SetReceivedSc(v float64)`

SetReceivedSc sets ReceivedSc field to given value.

### HasReceivedSc

`func (o *HashResponse) HasReceivedSc() bool`

HasReceivedSc returns a boolean if a field has been set.

### GetSentSc

`func (o *HashResponse) GetSentSc() float64`

GetSentSc returns the SentSc field if non-nil, zero value otherwise.

### GetSentScOk

`func (o *HashResponse) GetSentScOk() (*float64, bool)`

GetSentScOk returns a tuple with the SentSc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentSc

`func (o *HashResponse) SetSentSc(v float64)`

SetSentSc sets SentSc field to given value.

### HasSentSc

`func (o *HashResponse) HasSentSc() bool`

HasSentSc returns a boolean if a field has been set.

### GetBalanceSf

`func (o *HashResponse) GetBalanceSf() float32`

GetBalanceSf returns the BalanceSf field if non-nil, zero value otherwise.

### GetBalanceSfOk

`func (o *HashResponse) GetBalanceSfOk() (*float32, bool)`

GetBalanceSfOk returns a tuple with the BalanceSf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSf

`func (o *HashResponse) SetBalanceSf(v float32)`

SetBalanceSf sets BalanceSf field to given value.

### HasBalanceSf

`func (o *HashResponse) HasBalanceSf() bool`

HasBalanceSf returns a boolean if a field has been set.

### GetTotalTxCount

`func (o *HashResponse) GetTotalTxCount() int32`

GetTotalTxCount returns the TotalTxCount field if non-nil, zero value otherwise.

### GetTotalTxCountOk

`func (o *HashResponse) GetTotalTxCountOk() (*int32, bool)`

GetTotalTxCountOk returns a tuple with the TotalTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTxCount

`func (o *HashResponse) SetTotalTxCount(v int32)`

SetTotalTxCount sets TotalTxCount field to given value.

### HasTotalTxCount

`func (o *HashResponse) HasTotalTxCount() bool`

HasTotalTxCount returns a boolean if a field has been set.

### GetFirstSeen

`func (o *HashResponse) GetFirstSeen() int32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *HashResponse) GetFirstSeenOk() (*int32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *HashResponse) SetFirstSeen(v int32)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *HashResponse) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLast100Transactions

`func (o *HashResponse) GetLast100Transactions() []Transaction`

GetLast100Transactions returns the Last100Transactions field if non-nil, zero value otherwise.

### GetLast100TransactionsOk

`func (o *HashResponse) GetLast100TransactionsOk() (*[]Transaction, bool)`

GetLast100TransactionsOk returns a tuple with the Last100Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast100Transactions

`func (o *HashResponse) SetLast100Transactions(v []Transaction)`

SetLast100Transactions sets Last100Transactions field to given value.

### HasLast100Transactions

`func (o *HashResponse) HasLast100Transactions() bool`

HasLast100Transactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


