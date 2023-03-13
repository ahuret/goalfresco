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

// checks if the SiteMembershipRequestEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteMembershipRequestEntry{}

// SiteMembershipRequestEntry struct for SiteMembershipRequestEntry
type SiteMembershipRequestEntry struct {
	Entry SiteMembershipRequest `json:"entry"`
}

// NewSiteMembershipRequestEntry instantiates a new SiteMembershipRequestEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMembershipRequestEntry(entry SiteMembershipRequest) *SiteMembershipRequestEntry {
	this := SiteMembershipRequestEntry{}
	this.Entry = entry
	return &this
}

// NewSiteMembershipRequestEntryWithDefaults instantiates a new SiteMembershipRequestEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMembershipRequestEntryWithDefaults() *SiteMembershipRequestEntry {
	this := SiteMembershipRequestEntry{}
	return &this
}

// GetEntry returns the Entry field value
func (o *SiteMembershipRequestEntry) GetEntry() SiteMembershipRequest {
	if o == nil {
		var ret SiteMembershipRequest
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *SiteMembershipRequestEntry) GetEntryOk() (*SiteMembershipRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *SiteMembershipRequestEntry) SetEntry(v SiteMembershipRequest) {
	o.Entry = v
}

func (o SiteMembershipRequestEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteMembershipRequestEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	return toSerialize, nil
}

type NullableSiteMembershipRequestEntry struct {
	value *SiteMembershipRequestEntry
	isSet bool
}

func (v NullableSiteMembershipRequestEntry) Get() *SiteMembershipRequestEntry {
	return v.value
}

func (v *NullableSiteMembershipRequestEntry) Set(val *SiteMembershipRequestEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMembershipRequestEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMembershipRequestEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMembershipRequestEntry(val *SiteMembershipRequestEntry) *NullableSiteMembershipRequestEntry {
	return &NullableSiteMembershipRequestEntry{value: val, isSet: true}
}

func (v NullableSiteMembershipRequestEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMembershipRequestEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


