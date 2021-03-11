# HostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Pubkey** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Lon** | Pointer to **float32** |  | [optional] 
**Lat** | Pointer to **float32** |  | [optional] 
**Isp** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CountryName** | Pointer to **string** |  | [optional] 
**MaxDuration** | Pointer to **float32** |  | [optional] 
**TotalStorage** | Pointer to **float32** |  | [optional] 
**UsedStorage** | Pointer to **float32** |  | [optional] 
**Collateral** | Pointer to **float32** |  | [optional] 
**StoragePrice** | Pointer to **float32** |  | [optional] 
**UploadPrice** | Pointer to **float32** |  | [optional] 
**DownloadPrice** | Pointer to **float32** |  | [optional] 
**MinContractPrice** | Pointer to **float32** |  | [optional] 
**FirstSeenBlock** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**SectorSize** | Pointer to **int32** |  | [optional] 
**WindowSize** | Pointer to **int32** |  | [optional] 
**SeenBy** | Pointer to **[]string** |  | [optional] 
**Benchmarks** | Pointer to [**Benchmarks**](Benchmarks.md) |  | [optional] 
**ComparisonMatrix** | Pointer to [**ComparisonMatrix**](ComparisonMatrix.md) |  | [optional] 
**Ticks** | Pointer to [**Ticks**](Ticks.md) |  | [optional] 
**PlotBands** | Pointer to [**[]PlotBand**](PlotBand.md) |  | [optional] 
**RegionalScores** | Pointer to [**[]RegionalScore**](RegionalScore.md) |  | [optional] 
**Uptime1m** | Pointer to **float32** |  | [optional] 
**Uptime6m** | Pointer to **float32** |  | [optional] 

## Methods

### NewHostResponse

`func NewHostResponse() *HostResponse`

NewHostResponse instantiates a new HostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostResponseWithDefaults

`func NewHostResponseWithDefaults() *HostResponse`

NewHostResponseWithDefaults instantiates a new HostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestStatus

`func (o *HostResponse) GetRequestStatus() string`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *HostResponse) GetRequestStatusOk() (*string, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *HostResponse) SetRequestStatus(v string)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *HostResponse) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetId

`func (o *HostResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HostResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOnline

`func (o *HostResponse) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *HostResponse) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *HostResponse) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *HostResponse) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetRank

`func (o *HostResponse) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *HostResponse) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *HostResponse) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *HostResponse) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetNetwork

`func (o *HostResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HostResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HostResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *HostResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPubkey

`func (o *HostResponse) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *HostResponse) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *HostResponse) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *HostResponse) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetIp

`func (o *HostResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HostResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HostResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HostResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLon

`func (o *HostResponse) GetLon() float32`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *HostResponse) GetLonOk() (*float32, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *HostResponse) SetLon(v float32)`

SetLon sets Lon field to given value.

### HasLon

`func (o *HostResponse) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetLat

`func (o *HostResponse) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *HostResponse) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *HostResponse) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *HostResponse) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetIsp

`func (o *HostResponse) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *HostResponse) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *HostResponse) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *HostResponse) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetCountryCode

`func (o *HostResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *HostResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *HostResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *HostResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *HostResponse) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *HostResponse) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *HostResponse) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *HostResponse) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetMaxDuration

`func (o *HostResponse) GetMaxDuration() float32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *HostResponse) GetMaxDurationOk() (*float32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *HostResponse) SetMaxDuration(v float32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *HostResponse) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetTotalStorage

`func (o *HostResponse) GetTotalStorage() float32`

GetTotalStorage returns the TotalStorage field if non-nil, zero value otherwise.

### GetTotalStorageOk

`func (o *HostResponse) GetTotalStorageOk() (*float32, bool)`

GetTotalStorageOk returns a tuple with the TotalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorage

`func (o *HostResponse) SetTotalStorage(v float32)`

SetTotalStorage sets TotalStorage field to given value.

### HasTotalStorage

`func (o *HostResponse) HasTotalStorage() bool`

HasTotalStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *HostResponse) GetUsedStorage() float32`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *HostResponse) GetUsedStorageOk() (*float32, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *HostResponse) SetUsedStorage(v float32)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *HostResponse) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetCollateral

`func (o *HostResponse) GetCollateral() float32`

GetCollateral returns the Collateral field if non-nil, zero value otherwise.

### GetCollateralOk

`func (o *HostResponse) GetCollateralOk() (*float32, bool)`

GetCollateralOk returns a tuple with the Collateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateral

`func (o *HostResponse) SetCollateral(v float32)`

SetCollateral sets Collateral field to given value.

### HasCollateral

`func (o *HostResponse) HasCollateral() bool`

HasCollateral returns a boolean if a field has been set.

### GetStoragePrice

`func (o *HostResponse) GetStoragePrice() float32`

GetStoragePrice returns the StoragePrice field if non-nil, zero value otherwise.

### GetStoragePriceOk

`func (o *HostResponse) GetStoragePriceOk() (*float32, bool)`

GetStoragePriceOk returns a tuple with the StoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePrice

`func (o *HostResponse) SetStoragePrice(v float32)`

SetStoragePrice sets StoragePrice field to given value.

### HasStoragePrice

`func (o *HostResponse) HasStoragePrice() bool`

HasStoragePrice returns a boolean if a field has been set.

### GetUploadPrice

`func (o *HostResponse) GetUploadPrice() float32`

GetUploadPrice returns the UploadPrice field if non-nil, zero value otherwise.

### GetUploadPriceOk

`func (o *HostResponse) GetUploadPriceOk() (*float32, bool)`

GetUploadPriceOk returns a tuple with the UploadPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadPrice

`func (o *HostResponse) SetUploadPrice(v float32)`

SetUploadPrice sets UploadPrice field to given value.

### HasUploadPrice

`func (o *HostResponse) HasUploadPrice() bool`

HasUploadPrice returns a boolean if a field has been set.

### GetDownloadPrice

`func (o *HostResponse) GetDownloadPrice() float32`

GetDownloadPrice returns the DownloadPrice field if non-nil, zero value otherwise.

### GetDownloadPriceOk

`func (o *HostResponse) GetDownloadPriceOk() (*float32, bool)`

GetDownloadPriceOk returns a tuple with the DownloadPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPrice

`func (o *HostResponse) SetDownloadPrice(v float32)`

SetDownloadPrice sets DownloadPrice field to given value.

### HasDownloadPrice

`func (o *HostResponse) HasDownloadPrice() bool`

HasDownloadPrice returns a boolean if a field has been set.

### GetMinContractPrice

`func (o *HostResponse) GetMinContractPrice() float32`

GetMinContractPrice returns the MinContractPrice field if non-nil, zero value otherwise.

### GetMinContractPriceOk

`func (o *HostResponse) GetMinContractPriceOk() (*float32, bool)`

GetMinContractPriceOk returns a tuple with the MinContractPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinContractPrice

`func (o *HostResponse) SetMinContractPrice(v float32)`

SetMinContractPrice sets MinContractPrice field to given value.

### HasMinContractPrice

`func (o *HostResponse) HasMinContractPrice() bool`

HasMinContractPrice returns a boolean if a field has been set.

### GetFirstSeenBlock

`func (o *HostResponse) GetFirstSeenBlock() int32`

GetFirstSeenBlock returns the FirstSeenBlock field if non-nil, zero value otherwise.

### GetFirstSeenBlockOk

`func (o *HostResponse) GetFirstSeenBlockOk() (*int32, bool)`

GetFirstSeenBlockOk returns a tuple with the FirstSeenBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenBlock

`func (o *HostResponse) SetFirstSeenBlock(v int32)`

SetFirstSeenBlock sets FirstSeenBlock field to given value.

### HasFirstSeenBlock

`func (o *HostResponse) HasFirstSeenBlock() bool`

HasFirstSeenBlock returns a boolean if a field has been set.

### GetVersion

`func (o *HostResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HostResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HostResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HostResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSectorSize

`func (o *HostResponse) GetSectorSize() int32`

GetSectorSize returns the SectorSize field if non-nil, zero value otherwise.

### GetSectorSizeOk

`func (o *HostResponse) GetSectorSizeOk() (*int32, bool)`

GetSectorSizeOk returns a tuple with the SectorSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSize

`func (o *HostResponse) SetSectorSize(v int32)`

SetSectorSize sets SectorSize field to given value.

### HasSectorSize

`func (o *HostResponse) HasSectorSize() bool`

HasSectorSize returns a boolean if a field has been set.

### GetWindowSize

`func (o *HostResponse) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *HostResponse) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *HostResponse) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *HostResponse) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetSeenBy

`func (o *HostResponse) GetSeenBy() []string`

GetSeenBy returns the SeenBy field if non-nil, zero value otherwise.

### GetSeenByOk

`func (o *HostResponse) GetSeenByOk() (*[]string, bool)`

GetSeenByOk returns a tuple with the SeenBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenBy

`func (o *HostResponse) SetSeenBy(v []string)`

SetSeenBy sets SeenBy field to given value.

### HasSeenBy

`func (o *HostResponse) HasSeenBy() bool`

HasSeenBy returns a boolean if a field has been set.

### GetBenchmarks

`func (o *HostResponse) GetBenchmarks() Benchmarks`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *HostResponse) GetBenchmarksOk() (*Benchmarks, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *HostResponse) SetBenchmarks(v Benchmarks)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *HostResponse) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### GetComparisonMatrix

`func (o *HostResponse) GetComparisonMatrix() ComparisonMatrix`

GetComparisonMatrix returns the ComparisonMatrix field if non-nil, zero value otherwise.

### GetComparisonMatrixOk

`func (o *HostResponse) GetComparisonMatrixOk() (*ComparisonMatrix, bool)`

GetComparisonMatrixOk returns a tuple with the ComparisonMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonMatrix

`func (o *HostResponse) SetComparisonMatrix(v ComparisonMatrix)`

SetComparisonMatrix sets ComparisonMatrix field to given value.

### HasComparisonMatrix

`func (o *HostResponse) HasComparisonMatrix() bool`

HasComparisonMatrix returns a boolean if a field has been set.

### GetTicks

`func (o *HostResponse) GetTicks() Ticks`

GetTicks returns the Ticks field if non-nil, zero value otherwise.

### GetTicksOk

`func (o *HostResponse) GetTicksOk() (*Ticks, bool)`

GetTicksOk returns a tuple with the Ticks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicks

`func (o *HostResponse) SetTicks(v Ticks)`

SetTicks sets Ticks field to given value.

### HasTicks

`func (o *HostResponse) HasTicks() bool`

HasTicks returns a boolean if a field has been set.

### GetPlotBands

`func (o *HostResponse) GetPlotBands() []PlotBand`

GetPlotBands returns the PlotBands field if non-nil, zero value otherwise.

### GetPlotBandsOk

`func (o *HostResponse) GetPlotBandsOk() (*[]PlotBand, bool)`

GetPlotBandsOk returns a tuple with the PlotBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotBands

`func (o *HostResponse) SetPlotBands(v []PlotBand)`

SetPlotBands sets PlotBands field to given value.

### HasPlotBands

`func (o *HostResponse) HasPlotBands() bool`

HasPlotBands returns a boolean if a field has been set.

### GetRegionalScores

`func (o *HostResponse) GetRegionalScores() []RegionalScore`

GetRegionalScores returns the RegionalScores field if non-nil, zero value otherwise.

### GetRegionalScoresOk

`func (o *HostResponse) GetRegionalScoresOk() (*[]RegionalScore, bool)`

GetRegionalScoresOk returns a tuple with the RegionalScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalScores

`func (o *HostResponse) SetRegionalScores(v []RegionalScore)`

SetRegionalScores sets RegionalScores field to given value.

### HasRegionalScores

`func (o *HostResponse) HasRegionalScores() bool`

HasRegionalScores returns a boolean if a field has been set.

### GetUptime1m

`func (o *HostResponse) GetUptime1m() float32`

GetUptime1m returns the Uptime1m field if non-nil, zero value otherwise.

### GetUptime1mOk

`func (o *HostResponse) GetUptime1mOk() (*float32, bool)`

GetUptime1mOk returns a tuple with the Uptime1m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime1m

`func (o *HostResponse) SetUptime1m(v float32)`

SetUptime1m sets Uptime1m field to given value.

### HasUptime1m

`func (o *HostResponse) HasUptime1m() bool`

HasUptime1m returns a boolean if a field has been set.

### GetUptime6m

`func (o *HostResponse) GetUptime6m() float32`

GetUptime6m returns the Uptime6m field if non-nil, zero value otherwise.

### GetUptime6mOk

`func (o *HostResponse) GetUptime6mOk() (*float32, bool)`

GetUptime6mOk returns a tuple with the Uptime6m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime6m

`func (o *HostResponse) SetUptime6m(v float32)`

SetUptime6m sets Uptime6m field to given value.

### HasUptime6m

`func (o *HostResponse) HasUptime6m() bool`

HasUptime6m returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


