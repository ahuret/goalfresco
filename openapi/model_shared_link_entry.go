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

// checks if the SharedLinkEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedLinkEntry{}

// SharedLinkEntry struct for SharedLinkEntry
type SharedLinkEntry struct {
	Entry SharedLink `json:"entry"`
}

// NewSharedLinkEntry instantiates a new SharedLinkEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedLinkEntry(entry SharedLink) *SharedLinkEntry {
	this := SharedLinkEntry{}
	this.Entry = entry
	return &this
}

// NewSharedLinkEntryWithDefaults instantiates a new SharedLinkEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedLinkEntryWithDefaults() *SharedLinkEntry {
	this := SharedLinkEntry{}
	return &this
}

// GetEntry returns the Entry field value
func (o *SharedLinkEntry) GetEntry() SharedLink {
	if o == nil {
		var ret SharedLink
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *SharedLinkEntry) GetEntryOk() (*SharedLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *SharedLinkEntry) SetEntry(v SharedLink) {
	o.Entry = v
}

func (o SharedLinkEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedLinkEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	return toSerialize, nil
}

type NullableSharedLinkEntry struct {
	value *SharedLinkEntry
	isSet bool
}

func (v NullableSharedLinkEntry) Get() *SharedLinkEntry {
	return v.value
}

func (v *NullableSharedLinkEntry) Set(val *SharedLinkEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedLinkEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedLinkEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedLinkEntry(val *SharedLinkEntry) *NullableSharedLinkEntry {
	return &NullableSharedLinkEntry{value: val, isSet: true}
}

func (v NullableSharedLinkEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedLinkEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

