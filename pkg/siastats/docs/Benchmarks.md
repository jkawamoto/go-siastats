# Benchmarks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinalScore** | Pointer to **int32** |  | [optional] 
**Latency** | Pointer to **int32** |  | [optional] 
**LatencyScore** | Pointer to **int32** |  | [optional] 
**Up** | Pointer to **int64** |  | [optional] 
**UpScore** | Pointer to **int32** |  | [optional] 
**Down** | Pointer to **int64** |  | [optional] 
**DownScore** | Pointer to **int32** |  | [optional] 
**ContractSuccess** | Pointer to **bool** |  | [optional] 
**BenchFailureRate** | Pointer to **int32** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**ErrorFull** | Pointer to **string** |  | [optional] 
**ScoreChange1day** | Pointer to **string** |  | [optional] 
**ScoreChange7day** | Pointer to **string** |  | [optional] 
**ScoreChange7daySmooth** | Pointer to **string** |  | [optional] 

## Methods

### NewBenchmarks

`func NewBenchmarks() *Benchmarks`

NewBenchmarks instantiates a new Benchmarks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarksWithDefaults

`func NewBenchmarksWithDefaults() *Benchmarks`

NewBenchmarksWithDefaults instantiates a new Benchmarks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalScore

`func (o *Benchmarks) GetFinalScore() int32`

GetFinalScore returns the FinalScore field if non-nil, zero value otherwise.

### GetFinalScoreOk

`func (o *Benchmarks) GetFinalScoreOk() (*int32, bool)`

GetFinalScoreOk returns a tuple with the FinalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalScore

`func (o *Benchmarks) SetFinalScore(v int32)`

SetFinalScore sets FinalScore field to given value.

### HasFinalScore

`func (o *Benchmarks) HasFinalScore() bool`

HasFinalScore returns a boolean if a field has been set.

### GetLatency

`func (o *Benchmarks) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *Benchmarks) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *Benchmarks) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *Benchmarks) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetLatencyScore

`func (o *Benchmarks) GetLatencyScore() int32`

GetLatencyScore returns the LatencyScore field if non-nil, zero value otherwise.

### GetLatencyScoreOk

`func (o *Benchmarks) GetLatencyScoreOk() (*int32, bool)`

GetLatencyScoreOk returns a tuple with the LatencyScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyScore

`func (o *Benchmarks) SetLatencyScore(v int32)`

SetLatencyScore sets LatencyScore field to given value.

### HasLatencyScore

`func (o *Benchmarks) HasLatencyScore() bool`

HasLatencyScore returns a boolean if a field has been set.

### GetUp

`func (o *Benchmarks) GetUp() int64`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *Benchmarks) GetUpOk() (*int64, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *Benchmarks) SetUp(v int64)`

SetUp sets Up field to given value.

### HasUp

`func (o *Benchmarks) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpScore

`func (o *Benchmarks) GetUpScore() int32`

GetUpScore returns the UpScore field if non-nil, zero value otherwise.

### GetUpScoreOk

`func (o *Benchmarks) GetUpScoreOk() (*int32, bool)`

GetUpScoreOk returns a tuple with the UpScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpScore

`func (o *Benchmarks) SetUpScore(v int32)`

SetUpScore sets UpScore field to given value.

### HasUpScore

`func (o *Benchmarks) HasUpScore() bool`

HasUpScore returns a boolean if a field has been set.

### GetDown

`func (o *Benchmarks) GetDown() int64`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *Benchmarks) GetDownOk() (*int64, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *Benchmarks) SetDown(v int64)`

SetDown sets Down field to given value.

### HasDown

`func (o *Benchmarks) HasDown() bool`

HasDown returns a boolean if a field has been set.

### GetDownScore

`func (o *Benchmarks) GetDownScore() int32`

GetDownScore returns the DownScore field if non-nil, zero value otherwise.

### GetDownScoreOk

`func (o *Benchmarks) GetDownScoreOk() (*int32, bool)`

GetDownScoreOk returns a tuple with the DownScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownScore

`func (o *Benchmarks) SetDownScore(v int32)`

SetDownScore sets DownScore field to given value.

### HasDownScore

`func (o *Benchmarks) HasDownScore() bool`

HasDownScore returns a boolean if a field has been set.

### GetContractSuccess

`func (o *Benchmarks) GetContractSuccess() bool`

GetContractSuccess returns the ContractSuccess field if non-nil, zero value otherwise.

### GetContractSuccessOk

`func (o *Benchmarks) GetContractSuccessOk() (*bool, bool)`

GetContractSuccessOk returns a tuple with the ContractSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractSuccess

`func (o *Benchmarks) SetContractSuccess(v bool)`

SetContractSuccess sets ContractSuccess field to given value.

### HasContractSuccess

`func (o *Benchmarks) HasContractSuccess() bool`

HasContractSuccess returns a boolean if a field has been set.

### GetBenchFailureRate

`func (o *Benchmarks) GetBenchFailureRate() int32`

GetBenchFailureRate returns the BenchFailureRate field if non-nil, zero value otherwise.

### GetBenchFailureRateOk

`func (o *Benchmarks) GetBenchFailureRateOk() (*int32, bool)`

GetBenchFailureRateOk returns a tuple with the BenchFailureRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchFailureRate

`func (o *Benchmarks) SetBenchFailureRate(v int32)`

SetBenchFailureRate sets BenchFailureRate field to given value.

### HasBenchFailureRate

`func (o *Benchmarks) HasBenchFailureRate() bool`

HasBenchFailureRate returns a boolean if a field has been set.

### GetErrorType

`func (o *Benchmarks) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *Benchmarks) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *Benchmarks) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *Benchmarks) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetErrorDescription

`func (o *Benchmarks) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *Benchmarks) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *Benchmarks) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *Benchmarks) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorFull

`func (o *Benchmarks) GetErrorFull() string`

GetErrorFull returns the ErrorFull field if non-nil, zero value otherwise.

### GetErrorFullOk

`func (o *Benchmarks) GetErrorFullOk() (*string, bool)`

GetErrorFullOk returns a tuple with the ErrorFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFull

`func (o *Benchmarks) SetErrorFull(v string)`

SetErrorFull sets ErrorFull field to given value.

### HasErrorFull

`func (o *Benchmarks) HasErrorFull() bool`

HasErrorFull returns a boolean if a field has been set.

### GetScoreChange1day

`func (o *Benchmarks) GetScoreChange1day() string`

GetScoreChange1day returns the ScoreChange1day field if non-nil, zero value otherwise.

### GetScoreChange1dayOk

`func (o *Benchmarks) GetScoreChange1dayOk() (*string, bool)`

GetScoreChange1dayOk returns a tuple with the ScoreChange1day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreChange1day

`func (o *Benchmarks) SetScoreChange1day(v string)`

SetScoreChange1day sets ScoreChange1day field to given value.

### HasScoreChange1day

`func (o *Benchmarks) HasScoreChange1day() bool`

HasScoreChange1day returns a boolean if a field has been set.

### GetScoreChange7day

`func (o *Benchmarks) GetScoreChange7day() string`

GetScoreChange7day returns the ScoreChange7day field if non-nil, zero value otherwise.

### GetScoreChange7dayOk

`func (o *Benchmarks) GetScoreChange7dayOk() (*string, bool)`

GetScoreChange7dayOk returns a tuple with the ScoreChange7day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreChange7day

`func (o *Benchmarks) SetScoreChange7day(v string)`

SetScoreChange7day sets ScoreChange7day field to given value.

### HasScoreChange7day

`func (o *Benchmarks) HasScoreChange7day() bool`

HasScoreChange7day returns a boolean if a field has been set.

### GetScoreChange7daySmooth

`func (o *Benchmarks) GetScoreChange7daySmooth() string`

GetScoreChange7daySmooth returns the ScoreChange7daySmooth field if non-nil, zero value otherwise.

### GetScoreChange7daySmoothOk

`func (o *Benchmarks) GetScoreChange7daySmoothOk() (*string, bool)`

GetScoreChange7daySmoothOk returns a tuple with the ScoreChange7daySmooth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreChange7daySmooth

`func (o *Benchmarks) SetScoreChange7daySmooth(v string)`

SetScoreChange7daySmooth sets ScoreChange7daySmooth field to given value.

### HasScoreChange7daySmooth

`func (o *Benchmarks) HasScoreChange7daySmooth() bool`

HasScoreChange7daySmooth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


