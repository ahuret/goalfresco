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

// checks if the PersonNetworkEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonNetworkEntry{}

// PersonNetworkEntry struct for PersonNetworkEntry
type PersonNetworkEntry struct {
	Entry PersonNetwork `json:"entry"`
}

// NewPersonNetworkEntry instantiates a new PersonNetworkEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonNetworkEntry(entry PersonNetwork) *PersonNetworkEntry {
	this := PersonNetworkEntry{}
	this.Entry = entry
	return &this
}

// NewPersonNetworkEntryWithDefaults instantiates a new PersonNetworkEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonNetworkEntryWithDefaults() *PersonNetworkEntry {
	this := PersonNetworkEntry{}
	return &this
}

// GetEntry returns the Entry field value
func (o *PersonNetworkEntry) GetEntry() PersonNetwork {
	if o == nil {
		var ret PersonNetwork
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *PersonNetworkEntry) GetEntryOk() (*PersonNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *PersonNetworkEntry) SetEntry(v PersonNetwork) {
	o.Entry = v
}

func (o PersonNetworkEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonNetworkEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	return toSerialize, nil
}

type NullablePersonNetworkEntry struct {
	value *PersonNetworkEntry
	isSet bool
}

func (v NullablePersonNetworkEntry) Get() *PersonNetworkEntry {
	return v.value
}

func (v *NullablePersonNetworkEntry) Set(val *PersonNetworkEntry) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonNetworkEntry) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonNetworkEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonNetworkEntry(val *PersonNetworkEntry) *NullablePersonNetworkEntry {
	return &NullablePersonNetworkEntry{value: val, isSet: true}
}

func (v NullablePersonNetworkEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonNetworkEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


