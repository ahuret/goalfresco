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

// checks if the DownloadEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadEntry{}

// DownloadEntry struct for DownloadEntry
type DownloadEntry struct {
	Entry Download `json:"entry"`
}

// NewDownloadEntry instantiates a new DownloadEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadEntry(entry Download) *DownloadEntry {
	this := DownloadEntry{}
	this.Entry = entry
	return &this
}

// NewDownloadEntryWithDefaults instantiates a new DownloadEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadEntryWithDefaults() *DownloadEntry {
	this := DownloadEntry{}
	return &this
}

// GetEntry returns the Entry field value
func (o *DownloadEntry) GetEntry() Download {
	if o == nil {
		var ret Download
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *DownloadEntry) GetEntryOk() (*Download, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *DownloadEntry) SetEntry(v Download) {
	o.Entry = v
}

func (o DownloadEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	return toSerialize, nil
}

type NullableDownloadEntry struct {
	value *DownloadEntry
	isSet bool
}

func (v NullableDownloadEntry) Get() *DownloadEntry {
	return v.value
}

func (v *NullableDownloadEntry) Set(val *DownloadEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadEntry(val *DownloadEntry) *NullableDownloadEntry {
	return &NullableDownloadEntry{value: val, isSet: true}
}

func (v NullableDownloadEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

