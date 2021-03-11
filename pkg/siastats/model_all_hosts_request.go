/*
 * siastats.info API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2021-03-10
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package siastats

import (
	"encoding/json"
)

// AllHostsRequest struct for AllHostsRequest
type AllHostsRequest struct {
	Network string `json:"network"`
	List    string `json:"list"`
}

// NewAllHostsRequest instantiates a new AllHostsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllHostsRequest(network string, list string) *AllHostsRequest {
	this := AllHostsRequest{}
	this.Network = network
	this.List = list
	return &this
}

// NewAllHostsRequestWithDefaults instantiates a new AllHostsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllHostsRequestWithDefaults() *AllHostsRequest {
	this := AllHostsRequest{}
	return &this
}

// GetNetwork returns the Network field value
func (o *AllHostsRequest) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *AllHostsRequest) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *AllHostsRequest) SetNetwork(v string) {
	o.Network = v
}

// GetList returns the List field value
func (o *AllHostsRequest) GetList() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *AllHostsRequest) GetListOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *AllHostsRequest) SetList(v string) {
	o.List = v
}

func (o AllHostsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["list"] = o.List
	}
	return json.Marshal(toSerialize)
}

type NullableAllHostsRequest struct {
	value *AllHostsRequest
	isSet bool
}

func (v NullableAllHostsRequest) Get() *AllHostsRequest {
	return v.value
}

func (v *NullableAllHostsRequest) Set(val *AllHostsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAllHostsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAllHostsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllHostsRequest(val *AllHostsRequest) *NullableAllHostsRequest {
	return &NullableAllHostsRequest{value: val, isSet: true}
}

func (v NullableAllHostsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllHostsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}