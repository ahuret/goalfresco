/*
Alfresco Content Services REST API

**Core API**  Provides access to the core features of Alfresco Content Services. 

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the NetworkQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkQuota{}

// NetworkQuota Limits and usage of each quota. A network will have quotas for File space, the number of sites in the network, the number of people in the network, and the number of network administrators 
type NetworkQuota struct {
	Id string `json:"id"`
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
}

// NewNetworkQuota instantiates a new NetworkQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkQuota(id string, limit int64, usage int64) *NetworkQuota {
	this := NetworkQuota{}
	this.Id = id
	this.Limit = limit
	this.Usage = usage
	return &this
}

// NewNetworkQuotaWithDefaults instantiates a new NetworkQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkQuotaWithDefaults() *NetworkQuota {
	this := NetworkQuota{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkQuota) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkQuota) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkQuota) SetId(v string) {
	o.Id = v
}

// GetLimit returns the Limit field value
func (o *NetworkQuota) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *NetworkQuota) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *NetworkQuota) SetLimit(v int64) {
	o.Limit = v
}

// GetUsage returns the Usage field value
func (o *NetworkQuota) GetUsage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *NetworkQuota) GetUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *NetworkQuota) SetUsage(v int64) {
	o.Usage = v
}

func (o NetworkQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["limit"] = o.Limit
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

type NullableNetworkQuota struct {
	value *NetworkQuota
	isSet bool
}

func (v NullableNetworkQuota) Get() *NetworkQuota {
	return v.value
}

func (v *NullableNetworkQuota) Set(val *NetworkQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkQuota(val *NetworkQuota) *NullableNetworkQuota {
	return &NullableNetworkQuota{value: val, isSet: true}
}

func (v NullableNetworkQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


