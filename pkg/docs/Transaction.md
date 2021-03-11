# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterHash** | Pointer to **string** |  | [optional] 
**ScChange** | Pointer to **float64** |  | [optional] 
**SfChange** | Pointer to **float32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**TxType** | Pointer to **string** |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterHash

`func (o *Transaction) GetMasterHash() string`

GetMasterHash returns the MasterHash field if non-nil, zero value otherwise.

### GetMasterHashOk

`func (o *Transaction) GetMasterHashOk() (*string, bool)`

GetMasterHashOk returns a tuple with the MasterHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterHash

`func (o *Transaction) SetMasterHash(v string)`

SetMasterHash sets MasterHash field to given value.

### HasMasterHash

`func (o *Transaction) HasMasterHash() bool`

HasMasterHash returns a boolean if a field has been set.

### GetScChange

`func (o *Transaction) GetScChange() float64`

GetScChange returns the ScChange field if non-nil, zero value otherwise.

### GetScChangeOk

`func (o *Transaction) GetScChangeOk() (*float64, bool)`

GetScChangeOk returns a tuple with the ScChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScChange

`func (o *Transaction) SetScChange(v float64)`

SetScChange sets ScChange field to given value.

### HasScChange

`func (o *Transaction) HasScChange() bool`

HasScChange returns a boolean if a field has been set.

### GetSfChange

`func (o *Transaction) GetSfChange() float32`

GetSfChange returns the SfChange field if non-nil, zero value otherwise.

### GetSfChangeOk

`func (o *Transaction) GetSfChangeOk() (*float32, bool)`

GetSfChangeOk returns a tuple with the SfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfChange

`func (o *Transaction) SetSfChange(v float32)`

SetSfChange sets SfChange field to given value.

### HasSfChange

`func (o *Transaction) HasSfChange() bool`

HasSfChange returns a boolean if a field has been set.

### GetHeight

`func (o *Transaction) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Transaction) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Transaction) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Transaction) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetTimestamp

`func (o *Transaction) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Transaction) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Transaction) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Transaction) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTxType

`func (o *Transaction) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *Transaction) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *Transaction) SetTxType(v string)`

SetTxType sets TxType field to given value.

### HasTxType

`func (o *Transaction) HasTxType() bool`

HasTxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


