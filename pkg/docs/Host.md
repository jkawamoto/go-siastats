# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Pubkey** | Pointer to **string** |  | [optional] 
**CurrentIp** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**TotalStorage** | Pointer to **float32** |  | [optional] 
**AcceptingContracts** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**UsedStorage** | Pointer to **float32** |  | [optional] 
**Collateral** | Pointer to **int32** |  | [optional] 
**ContractPrice** | Pointer to **int32** |  | [optional] 
**StoragePrice** | Pointer to **int32** |  | [optional] 
**UploadPrice** | Pointer to **int32** |  | [optional] 
**DownloadPrice** | Pointer to **int32** |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**FinalScore** | Pointer to **int32** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 

## Methods

### NewHost

`func NewHost() *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Host) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Host) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOnline

`func (o *Host) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *Host) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *Host) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *Host) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetPubkey

`func (o *Host) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *Host) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *Host) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *Host) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetCurrentIp

`func (o *Host) GetCurrentIp() string`

GetCurrentIp returns the CurrentIp field if non-nil, zero value otherwise.

### GetCurrentIpOk

`func (o *Host) GetCurrentIpOk() (*string, bool)`

GetCurrentIpOk returns a tuple with the CurrentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIp

`func (o *Host) SetCurrentIp(v string)`

SetCurrentIp sets CurrentIp field to given value.

### HasCurrentIp

`func (o *Host) HasCurrentIp() bool`

HasCurrentIp returns a boolean if a field has been set.

### GetCountryCode

`func (o *Host) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Host) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Host) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Host) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetTotalStorage

`func (o *Host) GetTotalStorage() float32`

GetTotalStorage returns the TotalStorage field if non-nil, zero value otherwise.

### GetTotalStorageOk

`func (o *Host) GetTotalStorageOk() (*float32, bool)`

GetTotalStorageOk returns a tuple with the TotalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorage

`func (o *Host) SetTotalStorage(v float32)`

SetTotalStorage sets TotalStorage field to given value.

### HasTotalStorage

`func (o *Host) HasTotalStorage() bool`

HasTotalStorage returns a boolean if a field has been set.

### GetAcceptingContracts

`func (o *Host) GetAcceptingContracts() bool`

GetAcceptingContracts returns the AcceptingContracts field if non-nil, zero value otherwise.

### GetAcceptingContractsOk

`func (o *Host) GetAcceptingContractsOk() (*bool, bool)`

GetAcceptingContractsOk returns a tuple with the AcceptingContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingContracts

`func (o *Host) SetAcceptingContracts(v bool)`

SetAcceptingContracts sets AcceptingContracts field to given value.

### HasAcceptingContracts

`func (o *Host) HasAcceptingContracts() bool`

HasAcceptingContracts returns a boolean if a field has been set.

### GetVersion

`func (o *Host) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Host) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Host) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Host) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUsedStorage

`func (o *Host) GetUsedStorage() float32`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *Host) GetUsedStorageOk() (*float32, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *Host) SetUsedStorage(v float32)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *Host) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetCollateral

`func (o *Host) GetCollateral() int32`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *Host) GetCollateralOk() (*int32, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *Host) SetCollateral(v int32)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *Host) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetContractPrice

`func (o *Host) GetContractPrice() int32`

GetContractPrice returns the ContractPrice field if non-nil, zero value otherwise.

### GetContractPriceOk

`func (o *Host) GetContractPriceOk() (*int32, bool)`

GetContractPriceOk returns a tuple with the ContractPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractPrice

`func (o *Host) SetContractPrice(v int32)`

SetContractPrice sets ContractPrice field to given value.

### HasContractPrice

`func (o *Host) HasContractPrice() bool`

HasContractPrice returns a boolean if a field has been set.

### GetStoragePrice

`func (o *Host) GetStoragePrice() int32`

GetStoragePrice returns the StoragePrice field if non-nil, zero value otherwise.

### GetStoragePriceOk

`func (o *Host) GetStoragePriceOk() (*int32, bool)`

GetStoragePriceOk returns a tuple with the StoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePrice

`func (o *Host) SetStoragePrice(v int32)`

SetStoragePrice sets StoragePrice field to given value.

### HasStoragePrice

`func (o *Host) HasStoragePrice() bool`

HasStoragePrice returns a boolean if a field has been set.

### GetUploadPrice

`func (o *Host) GetUploadPrice() int32`

GetUploadPrice returns the UploadPrice field if non-nil, zero value otherwise.

### GetUploadPriceOk

`func (o *Host) GetUploadPriceOk() (*int32, bool)`

GetUploadPriceOk returns a tuple with the UploadPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadPrice

`func (o *Host) SetUploadPrice(v int32)`

SetUploadPrice sets UploadPrice field to given value.

### HasUploadPrice

`func (o *Host) HasUploadPrice() bool`

HasUploadPrice returns a boolean if a field has been set.

### GetDownloadPrice

`func (o *Host) GetDownloadPrice() int32`

GetDownloadPrice returns the DownloadPrice field if non-nil, zero value otherwise.

### GetDownloadPriceOk

`func (o *Host) GetDownloadPriceOk() (*int32, bool)`

GetDownloadPriceOk returns a tuple with the DownloadPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPrice

`func (o *Host) SetDownloadPrice(v int32)`

SetDownloadPrice sets DownloadPrice field to given value.

### HasDownloadPrice

`func (o *Host) HasDownloadPrice() bool`

HasDownloadPrice returns a boolean if a field has been set.

### GetRank

`func (o *Host) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Host) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Host) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *Host) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetFinalScore

`func (o *Host) GetFinalScore() int32`

GetFinalScore returns the FinalScore field if non-nil, zero value otherwise.

### GetFinalScoreOk

`func (o *Host) GetFinalScoreOk() (*int32, bool)`

GetFinalScoreOk returns a tuple with the FinalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalScore

`func (o *Host) SetFinalScore(v int32)`

SetFinalScore sets FinalScore field to given value.

### HasFinalScore

`func (o *Host) HasFinalScore() bool`

HasFinalScore returns a boolean if a field has been set.

### GetErrorType

`func (o *Host) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *Host) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *Host) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *Host) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


