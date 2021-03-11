# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consensusblock** | Pointer to **int32** |  | [optional] 
**Lastblock** | Pointer to **int32** |  | [optional] 
**Heartbeat** | Pointer to **int64** |  | [optional] 
**Peers** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsensusblock

`func (o *Status) GetConsensusblock() int32`

GetConsensusblock returns the Consensusblock field if non-nil, zero value otherwise.

### GetConsensusblockOk

`func (o *Status) GetConsensusblockOk() (*int32, bool)`

GetConsensusblockOk returns a tuple with the Consensusblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsensusblock

`func (o *Status) SetConsensusblock(v int32)`

SetConsensusblock sets Consensusblock field to given value.

### HasConsensusblock

`func (o *Status) HasConsensusblock() bool`

HasConsensusblock returns a boolean if a field has been set.

### GetLastblock

`func (o *Status) GetLastblock() int32`

GetLastblock returns the Lastblock field if non-nil, zero value otherwise.

### GetLastblockOk

`func (o *Status) GetLastblockOk() (*int32, bool)`

GetLastblockOk returns a tuple with the Lastblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastblock

`func (o *Status) SetLastblock(v int32)`

SetLastblock sets Lastblock field to given value.

### HasLastblock

`func (o *Status) HasLastblock() bool`

HasLastblock returns a boolean if a field has been set.

### GetHeartbeat

`func (o *Status) GetHeartbeat() int64`

GetHeartbeat returns the Heartbeat field if non-nil, zero value otherwise.

### GetHeartbeatOk

`func (o *Status) GetHeartbeatOk() (*int64, bool)`

GetHeartbeatOk returns a tuple with the Heartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeat

`func (o *Status) SetHeartbeat(v int64)`

SetHeartbeat sets Heartbeat field to given value.

### HasHeartbeat

`func (o *Status) HasHeartbeat() bool`

HasHeartbeat returns a boolean if a field has been set.

### GetPeers

`func (o *Status) GetPeers() int32`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *Status) GetPeersOk() (*int32, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *Status) SetPeers(v int32)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *Status) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetVersion

`func (o *Status) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Status) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Status) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Status) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


