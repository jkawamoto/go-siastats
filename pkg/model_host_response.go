/*
 * siastats.info API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2019-10-03
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package siastats

import (
	"encoding/json"
)

// HostResponse struct for HostResponse
type HostResponse struct {
	RequestStatus    *string           `json:"requestStatus,omitempty"`
	Id               *int32            `json:"id,omitempty"`
	Online           *bool             `json:"online,omitempty"`
	Rank             *int32            `json:"rank,omitempty"`
	Network          *string           `json:"network,omitempty"`
	Pubkey           *string           `json:"pubkey,omitempty"`
	Ip               *string           `json:"ip,omitempty"`
	Lon              *float32          `json:"lon,omitempty"`
	Lat              *float32          `json:"lat,omitempty"`
	Isp              *string           `json:"isp,omitempty"`
	CountryCode      *string           `json:"countryCode,omitempty"`
	CountryName      *string           `json:"countryName,omitempty"`
	MaxDuration      *float32          `json:"maxDuration,omitempty"`
	TotalStorage     *float32          `json:"totalStorage,omitempty"`
	UsedStorage      *float32          `json:"usedStorage,omitempty"`
	Collateral       *float32          `json:"collateral,omitempty"`
	StoragePrice     *float32          `json:"storagePrice,omitempty"`
	UploadPrice      *float32          `json:"uploadPrice,omitempty"`
	DownloadPrice    *float32          `json:"downloadPrice,omitempty"`
	MinContractPrice *float32          `json:"minContractPrice,omitempty"`
	FirstSeenBlock   *int32            `json:"firstSeenBlock,omitempty"`
	Version          *string           `json:"version,omitempty"`
	SectorSize       *int32            `json:"sectorSize,omitempty"`
	WindowSize       *int32            `json:"windowSize,omitempty"`
	SeenBy           *[]string         `json:"seenBy,omitempty"`
	Benchmarks       *Benchmarks       `json:"benchmarks,omitempty"`
	ComparisonMatrix *ComparisonMatrix `json:"comparisonMatrix,omitempty"`
	Ticks            *Ticks            `json:"ticks,omitempty"`
	PlotBands        *[]PlotBand       `json:"plotBands,omitempty"`
	RegionalScores   *[]RegionalScore  `json:"regionalScores,omitempty"`
	Uptime1m         *float32          `json:"uptime1m,omitempty"`
	Uptime6m         *float32          `json:"uptime6m,omitempty"`
}

// NewHostResponse instantiates a new HostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostResponse() *HostResponse {
	this := HostResponse{}
	return &this
}

// NewHostResponseWithDefaults instantiates a new HostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostResponseWithDefaults() *HostResponse {
	this := HostResponse{}
	return &this
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *HostResponse) GetRequestStatus() string {
	if o == nil || o.RequestStatus == nil {
		var ret string
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetRequestStatusOk() (*string, bool) {
	if o == nil || o.RequestStatus == nil {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *HostResponse) HasRequestStatus() bool {
	if o != nil && o.RequestStatus != nil {
		return true
	}

	return false
}

// SetRequestStatus gets a reference to the given string and assigns it to the RequestStatus field.
func (o *HostResponse) SetRequestStatus(v string) {
	o.RequestStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostResponse) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *HostResponse) SetId(v int32) {
	o.Id = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *HostResponse) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *HostResponse) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *HostResponse) SetOnline(v bool) {
	o.Online = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *HostResponse) GetRank() int32 {
	if o == nil || o.Rank == nil {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetRankOk() (*int32, bool) {
	if o == nil || o.Rank == nil {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *HostResponse) HasRank() bool {
	if o != nil && o.Rank != nil {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *HostResponse) SetRank(v int32) {
	o.Rank = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *HostResponse) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *HostResponse) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *HostResponse) SetNetwork(v string) {
	o.Network = &v
}

// GetPubkey returns the Pubkey field value if set, zero value otherwise.
func (o *HostResponse) GetPubkey() string {
	if o == nil || o.Pubkey == nil {
		var ret string
		return ret
	}
	return *o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetPubkeyOk() (*string, bool) {
	if o == nil || o.Pubkey == nil {
		return nil, false
	}
	return o.Pubkey, true
}

// HasPubkey returns a boolean if a field has been set.
func (o *HostResponse) HasPubkey() bool {
	if o != nil && o.Pubkey != nil {
		return true
	}

	return false
}

// SetPubkey gets a reference to the given string and assigns it to the Pubkey field.
func (o *HostResponse) SetPubkey(v string) {
	o.Pubkey = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *HostResponse) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *HostResponse) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *HostResponse) SetIp(v string) {
	o.Ip = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *HostResponse) GetLon() float32 {
	if o == nil || o.Lon == nil {
		var ret float32
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetLonOk() (*float32, bool) {
	if o == nil || o.Lon == nil {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *HostResponse) HasLon() bool {
	if o != nil && o.Lon != nil {
		return true
	}

	return false
}

// SetLon gets a reference to the given float32 and assigns it to the Lon field.
func (o *HostResponse) SetLon(v float32) {
	o.Lon = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *HostResponse) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *HostResponse) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *HostResponse) SetLat(v float32) {
	o.Lat = &v
}

// GetIsp returns the Isp field value if set, zero value otherwise.
func (o *HostResponse) GetIsp() string {
	if o == nil || o.Isp == nil {
		var ret string
		return ret
	}
	return *o.Isp
}

// GetIspOk returns a tuple with the Isp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetIspOk() (*string, bool) {
	if o == nil || o.Isp == nil {
		return nil, false
	}
	return o.Isp, true
}

// HasIsp returns a boolean if a field has been set.
func (o *HostResponse) HasIsp() bool {
	if o != nil && o.Isp != nil {
		return true
	}

	return false
}

// SetIsp gets a reference to the given string and assigns it to the Isp field.
func (o *HostResponse) SetIsp(v string) {
	o.Isp = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *HostResponse) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *HostResponse) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *HostResponse) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *HostResponse) GetCountryName() string {
	if o == nil || o.CountryName == nil {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetCountryNameOk() (*string, bool) {
	if o == nil || o.CountryName == nil {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *HostResponse) HasCountryName() bool {
	if o != nil && o.CountryName != nil {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *HostResponse) SetCountryName(v string) {
	o.CountryName = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *HostResponse) GetMaxDuration() float32 {
	if o == nil || o.MaxDuration == nil {
		var ret float32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetMaxDurationOk() (*float32, bool) {
	if o == nil || o.MaxDuration == nil {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *HostResponse) HasMaxDuration() bool {
	if o != nil && o.MaxDuration != nil {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given float32 and assigns it to the MaxDuration field.
func (o *HostResponse) SetMaxDuration(v float32) {
	o.MaxDuration = &v
}

// GetTotalStorage returns the TotalStorage field value if set, zero value otherwise.
func (o *HostResponse) GetTotalStorage() float32 {
	if o == nil || o.TotalStorage == nil {
		var ret float32
		return ret
	}
	return *o.TotalStorage
}

// GetTotalStorageOk returns a tuple with the TotalStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetTotalStorageOk() (*float32, bool) {
	if o == nil || o.TotalStorage == nil {
		return nil, false
	}
	return o.TotalStorage, true
}

// HasTotalStorage returns a boolean if a field has been set.
func (o *HostResponse) HasTotalStorage() bool {
	if o != nil && o.TotalStorage != nil {
		return true
	}

	return false
}

// SetTotalStorage gets a reference to the given float32 and assigns it to the TotalStorage field.
func (o *HostResponse) SetTotalStorage(v float32) {
	o.TotalStorage = &v
}

// GetUsedStorage returns the UsedStorage field value if set, zero value otherwise.
func (o *HostResponse) GetUsedStorage() float32 {
	if o == nil || o.UsedStorage == nil {
		var ret float32
		return ret
	}
	return *o.UsedStorage
}

// GetUsedStorageOk returns a tuple with the UsedStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetUsedStorageOk() (*float32, bool) {
	if o == nil || o.UsedStorage == nil {
		return nil, false
	}
	return o.UsedStorage, true
}

// HasUsedStorage returns a boolean if a field has been set.
func (o *HostResponse) HasUsedStorage() bool {
	if o != nil && o.UsedStorage != nil {
		return true
	}

	return false
}

// SetUsedStorage gets a reference to the given float32 and assigns it to the UsedStorage field.
func (o *HostResponse) SetUsedStorage(v float32) {
	o.UsedStorage = &v
}

// GetCollateral returns the Collateral field value if set, zero value otherwise.
func (o *HostResponse) GetCollateral() float32 {
	if o == nil || o.Collateral == nil {
		var ret float32
		return ret
	}
	return *o.Collateral
}

// GetCollateralOk returns a tuple with the Collateral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetCollateralOk() (*float32, bool) {
	if o == nil || o.Collateral == nil {
		return nil, false
	}
	return o.Collateral, true
}

// HasCollateral returns a boolean if a field has been set.
func (o *HostResponse) HasCollateral() bool {
	if o != nil && o.Collateral != nil {
		return true
	}

	return false
}

// SetCollateral gets a reference to the given float32 and assigns it to the Collateral field.
func (o *HostResponse) SetCollateral(v float32) {
	o.Collateral = &v
}

// GetStoragePrice returns the StoragePrice field value if set, zero value otherwise.
func (o *HostResponse) GetStoragePrice() float32 {
	if o == nil || o.StoragePrice == nil {
		var ret float32
		return ret
	}
	return *o.StoragePrice
}

// GetStoragePriceOk returns a tuple with the StoragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetStoragePriceOk() (*float32, bool) {
	if o == nil || o.StoragePrice == nil {
		return nil, false
	}
	return o.StoragePrice, true
}

// HasStoragePrice returns a boolean if a field has been set.
func (o *HostResponse) HasStoragePrice() bool {
	if o != nil && o.StoragePrice != nil {
		return true
	}

	return false
}

// SetStoragePrice gets a reference to the given float32 and assigns it to the StoragePrice field.
func (o *HostResponse) SetStoragePrice(v float32) {
	o.StoragePrice = &v
}

// GetUploadPrice returns the UploadPrice field value if set, zero value otherwise.
func (o *HostResponse) GetUploadPrice() float32 {
	if o == nil || o.UploadPrice == nil {
		var ret float32
		return ret
	}
	return *o.UploadPrice
}

// GetUploadPriceOk returns a tuple with the UploadPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetUploadPriceOk() (*float32, bool) {
	if o == nil || o.UploadPrice == nil {
		return nil, false
	}
	return o.UploadPrice, true
}

// HasUploadPrice returns a boolean if a field has been set.
func (o *HostResponse) HasUploadPrice() bool {
	if o != nil && o.UploadPrice != nil {
		return true
	}

	return false
}

// SetUploadPrice gets a reference to the given float32 and assigns it to the UploadPrice field.
func (o *HostResponse) SetUploadPrice(v float32) {
	o.UploadPrice = &v
}

// GetDownloadPrice returns the DownloadPrice field value if set, zero value otherwise.
func (o *HostResponse) GetDownloadPrice() float32 {
	if o == nil || o.DownloadPrice == nil {
		var ret float32
		return ret
	}
	return *o.DownloadPrice
}

// GetDownloadPriceOk returns a tuple with the DownloadPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetDownloadPriceOk() (*float32, bool) {
	if o == nil || o.DownloadPrice == nil {
		return nil, false
	}
	return o.DownloadPrice, true
}

// HasDownloadPrice returns a boolean if a field has been set.
func (o *HostResponse) HasDownloadPrice() bool {
	if o != nil && o.DownloadPrice != nil {
		return true
	}

	return false
}

// SetDownloadPrice gets a reference to the given float32 and assigns it to the DownloadPrice field.
func (o *HostResponse) SetDownloadPrice(v float32) {
	o.DownloadPrice = &v
}

// GetMinContractPrice returns the MinContractPrice field value if set, zero value otherwise.
func (o *HostResponse) GetMinContractPrice() float32 {
	if o == nil || o.MinContractPrice == nil {
		var ret float32
		return ret
	}
	return *o.MinContractPrice
}

// GetMinContractPriceOk returns a tuple with the MinContractPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetMinContractPriceOk() (*float32, bool) {
	if o == nil || o.MinContractPrice == nil {
		return nil, false
	}
	return o.MinContractPrice, true
}

// HasMinContractPrice returns a boolean if a field has been set.
func (o *HostResponse) HasMinContractPrice() bool {
	if o != nil && o.MinContractPrice != nil {
		return true
	}

	return false
}

// SetMinContractPrice gets a reference to the given float32 and assigns it to the MinContractPrice field.
func (o *HostResponse) SetMinContractPrice(v float32) {
	o.MinContractPrice = &v
}

// GetFirstSeenBlock returns the FirstSeenBlock field value if set, zero value otherwise.
func (o *HostResponse) GetFirstSeenBlock() int32 {
	if o == nil || o.FirstSeenBlock == nil {
		var ret int32
		return ret
	}
	return *o.FirstSeenBlock
}

// GetFirstSeenBlockOk returns a tuple with the FirstSeenBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetFirstSeenBlockOk() (*int32, bool) {
	if o == nil || o.FirstSeenBlock == nil {
		return nil, false
	}
	return o.FirstSeenBlock, true
}

// HasFirstSeenBlock returns a boolean if a field has been set.
func (o *HostResponse) HasFirstSeenBlock() bool {
	if o != nil && o.FirstSeenBlock != nil {
		return true
	}

	return false
}

// SetFirstSeenBlock gets a reference to the given int32 and assigns it to the FirstSeenBlock field.
func (o *HostResponse) SetFirstSeenBlock(v int32) {
	o.FirstSeenBlock = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HostResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HostResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HostResponse) SetVersion(v string) {
	o.Version = &v
}

// GetSectorSize returns the SectorSize field value if set, zero value otherwise.
func (o *HostResponse) GetSectorSize() int32 {
	if o == nil || o.SectorSize == nil {
		var ret int32
		return ret
	}
	return *o.SectorSize
}

// GetSectorSizeOk returns a tuple with the SectorSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetSectorSizeOk() (*int32, bool) {
	if o == nil || o.SectorSize == nil {
		return nil, false
	}
	return o.SectorSize, true
}

// HasSectorSize returns a boolean if a field has been set.
func (o *HostResponse) HasSectorSize() bool {
	if o != nil && o.SectorSize != nil {
		return true
	}

	return false
}

// SetSectorSize gets a reference to the given int32 and assigns it to the SectorSize field.
func (o *HostResponse) SetSectorSize(v int32) {
	o.SectorSize = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *HostResponse) GetWindowSize() int32 {
	if o == nil || o.WindowSize == nil {
		var ret int32
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetWindowSizeOk() (*int32, bool) {
	if o == nil || o.WindowSize == nil {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *HostResponse) HasWindowSize() bool {
	if o != nil && o.WindowSize != nil {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int32 and assigns it to the WindowSize field.
func (o *HostResponse) SetWindowSize(v int32) {
	o.WindowSize = &v
}

// GetSeenBy returns the SeenBy field value if set, zero value otherwise.
func (o *HostResponse) GetSeenBy() []string {
	if o == nil || o.SeenBy == nil {
		var ret []string
		return ret
	}
	return *o.SeenBy
}

// GetSeenByOk returns a tuple with the SeenBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetSeenByOk() (*[]string, bool) {
	if o == nil || o.SeenBy == nil {
		return nil, false
	}
	return o.SeenBy, true
}

// HasSeenBy returns a boolean if a field has been set.
func (o *HostResponse) HasSeenBy() bool {
	if o != nil && o.SeenBy != nil {
		return true
	}

	return false
}

// SetSeenBy gets a reference to the given []string and assigns it to the SeenBy field.
func (o *HostResponse) SetSeenBy(v []string) {
	o.SeenBy = &v
}

// GetBenchmarks returns the Benchmarks field value if set, zero value otherwise.
func (o *HostResponse) GetBenchmarks() Benchmarks {
	if o == nil || o.Benchmarks == nil {
		var ret Benchmarks
		return ret
	}
	return *o.Benchmarks
}

// GetBenchmarksOk returns a tuple with the Benchmarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetBenchmarksOk() (*Benchmarks, bool) {
	if o == nil || o.Benchmarks == nil {
		return nil, false
	}
	return o.Benchmarks, true
}

// HasBenchmarks returns a boolean if a field has been set.
func (o *HostResponse) HasBenchmarks() bool {
	if o != nil && o.Benchmarks != nil {
		return true
	}

	return false
}

// SetBenchmarks gets a reference to the given Benchmarks and assigns it to the Benchmarks field.
func (o *HostResponse) SetBenchmarks(v Benchmarks) {
	o.Benchmarks = &v
}

// GetComparisonMatrix returns the ComparisonMatrix field value if set, zero value otherwise.
func (o *HostResponse) GetComparisonMatrix() ComparisonMatrix {
	if o == nil || o.ComparisonMatrix == nil {
		var ret ComparisonMatrix
		return ret
	}
	return *o.ComparisonMatrix
}

// GetComparisonMatrixOk returns a tuple with the ComparisonMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetComparisonMatrixOk() (*ComparisonMatrix, bool) {
	if o == nil || o.ComparisonMatrix == nil {
		return nil, false
	}
	return o.ComparisonMatrix, true
}

// HasComparisonMatrix returns a boolean if a field has been set.
func (o *HostResponse) HasComparisonMatrix() bool {
	if o != nil && o.ComparisonMatrix != nil {
		return true
	}

	return false
}

// SetComparisonMatrix gets a reference to the given ComparisonMatrix and assigns it to the ComparisonMatrix field.
func (o *HostResponse) SetComparisonMatrix(v ComparisonMatrix) {
	o.ComparisonMatrix = &v
}

// GetTicks returns the Ticks field value if set, zero value otherwise.
func (o *HostResponse) GetTicks() Ticks {
	if o == nil || o.Ticks == nil {
		var ret Ticks
		return ret
	}
	return *o.Ticks
}

// GetTicksOk returns a tuple with the Ticks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetTicksOk() (*Ticks, bool) {
	if o == nil || o.Ticks == nil {
		return nil, false
	}
	return o.Ticks, true
}

// HasTicks returns a boolean if a field has been set.
func (o *HostResponse) HasTicks() bool {
	if o != nil && o.Ticks != nil {
		return true
	}

	return false
}

// SetTicks gets a reference to the given Ticks and assigns it to the Ticks field.
func (o *HostResponse) SetTicks(v Ticks) {
	o.Ticks = &v
}

// GetPlotBands returns the PlotBands field value if set, zero value otherwise.
func (o *HostResponse) GetPlotBands() []PlotBand {
	if o == nil || o.PlotBands == nil {
		var ret []PlotBand
		return ret
	}
	return *o.PlotBands
}

// GetPlotBandsOk returns a tuple with the PlotBands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetPlotBandsOk() (*[]PlotBand, bool) {
	if o == nil || o.PlotBands == nil {
		return nil, false
	}
	return o.PlotBands, true
}

// HasPlotBands returns a boolean if a field has been set.
func (o *HostResponse) HasPlotBands() bool {
	if o != nil && o.PlotBands != nil {
		return true
	}

	return false
}

// SetPlotBands gets a reference to the given []PlotBand and assigns it to the PlotBands field.
func (o *HostResponse) SetPlotBands(v []PlotBand) {
	o.PlotBands = &v
}

// GetRegionalScores returns the RegionalScores field value if set, zero value otherwise.
func (o *HostResponse) GetRegionalScores() []RegionalScore {
	if o == nil || o.RegionalScores == nil {
		var ret []RegionalScore
		return ret
	}
	return *o.RegionalScores
}

// GetRegionalScoresOk returns a tuple with the RegionalScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetRegionalScoresOk() (*[]RegionalScore, bool) {
	if o == nil || o.RegionalScores == nil {
		return nil, false
	}
	return o.RegionalScores, true
}

// HasRegionalScores returns a boolean if a field has been set.
func (o *HostResponse) HasRegionalScores() bool {
	if o != nil && o.RegionalScores != nil {
		return true
	}

	return false
}

// SetRegionalScores gets a reference to the given []RegionalScore and assigns it to the RegionalScores field.
func (o *HostResponse) SetRegionalScores(v []RegionalScore) {
	o.RegionalScores = &v
}

// GetUptime1m returns the Uptime1m field value if set, zero value otherwise.
func (o *HostResponse) GetUptime1m() float32 {
	if o == nil || o.Uptime1m == nil {
		var ret float32
		return ret
	}
	return *o.Uptime1m
}

// GetUptime1mOk returns a tuple with the Uptime1m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetUptime1mOk() (*float32, bool) {
	if o == nil || o.Uptime1m == nil {
		return nil, false
	}
	return o.Uptime1m, true
}

// HasUptime1m returns a boolean if a field has been set.
func (o *HostResponse) HasUptime1m() bool {
	if o != nil && o.Uptime1m != nil {
		return true
	}

	return false
}

// SetUptime1m gets a reference to the given float32 and assigns it to the Uptime1m field.
func (o *HostResponse) SetUptime1m(v float32) {
	o.Uptime1m = &v
}

// GetUptime6m returns the Uptime6m field value if set, zero value otherwise.
func (o *HostResponse) GetUptime6m() float32 {
	if o == nil || o.Uptime6m == nil {
		var ret float32
		return ret
	}
	return *o.Uptime6m
}

// GetUptime6mOk returns a tuple with the Uptime6m field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResponse) GetUptime6mOk() (*float32, bool) {
	if o == nil || o.Uptime6m == nil {
		return nil, false
	}
	return o.Uptime6m, true
}

// HasUptime6m returns a boolean if a field has been set.
func (o *HostResponse) HasUptime6m() bool {
	if o != nil && o.Uptime6m != nil {
		return true
	}

	return false
}

// SetUptime6m gets a reference to the given float32 and assigns it to the Uptime6m field.
func (o *HostResponse) SetUptime6m(v float32) {
	o.Uptime6m = &v
}

func (o HostResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestStatus != nil {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.Rank != nil {
		toSerialize["rank"] = o.Rank
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Pubkey != nil {
		toSerialize["pubkey"] = o.Pubkey
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Lon != nil {
		toSerialize["lon"] = o.Lon
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Isp != nil {
		toSerialize["isp"] = o.Isp
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.CountryName != nil {
		toSerialize["countryName"] = o.CountryName
	}
	if o.MaxDuration != nil {
		toSerialize["maxDuration"] = o.MaxDuration
	}
	if o.TotalStorage != nil {
		toSerialize["totalStorage"] = o.TotalStorage
	}
	if o.UsedStorage != nil {
		toSerialize["usedStorage"] = o.UsedStorage
	}
	if o.Collateral != nil {
		toSerialize["collateral"] = o.Collateral
	}
	if o.StoragePrice != nil {
		toSerialize["storagePrice"] = o.StoragePrice
	}
	if o.UploadPrice != nil {
		toSerialize["uploadPrice"] = o.UploadPrice
	}
	if o.DownloadPrice != nil {
		toSerialize["downloadPrice"] = o.DownloadPrice
	}
	if o.MinContractPrice != nil {
		toSerialize["minContractPrice"] = o.MinContractPrice
	}
	if o.FirstSeenBlock != nil {
		toSerialize["firstSeenBlock"] = o.FirstSeenBlock
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.SectorSize != nil {
		toSerialize["sectorSize"] = o.SectorSize
	}
	if o.WindowSize != nil {
		toSerialize["windowSize"] = o.WindowSize
	}
	if o.SeenBy != nil {
		toSerialize["seenBy"] = o.SeenBy
	}
	if o.Benchmarks != nil {
		toSerialize["benchmarks"] = o.Benchmarks
	}
	if o.ComparisonMatrix != nil {
		toSerialize["comparisonMatrix"] = o.ComparisonMatrix
	}
	if o.Ticks != nil {
		toSerialize["ticks"] = o.Ticks
	}
	if o.PlotBands != nil {
		toSerialize["plotBands"] = o.PlotBands
	}
	if o.RegionalScores != nil {
		toSerialize["regionalScores"] = o.RegionalScores
	}
	if o.Uptime1m != nil {
		toSerialize["uptime1m"] = o.Uptime1m
	}
	if o.Uptime6m != nil {
		toSerialize["uptime6m"] = o.Uptime6m
	}
	return json.Marshal(toSerialize)
}

type NullableHostResponse struct {
	value *HostResponse
	isSet bool
}

func (v NullableHostResponse) Get() *HostResponse {
	return v.value
}

func (v *NullableHostResponse) Set(val *HostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostResponse(val *HostResponse) *NullableHostResponse {
	return &NullableHostResponse{value: val, isSet: true}
}

func (v NullableHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
