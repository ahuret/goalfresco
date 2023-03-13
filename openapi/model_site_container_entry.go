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

// checks if the SiteContainerEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteContainerEntry{}

// SiteContainerEntry struct for SiteContainerEntry
type SiteContainerEntry struct {
	Entry SiteContainer `json:"entry"`
}

// NewSiteContainerEntry instantiates a new SiteContainerEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteContainerEntry(entry SiteContainer) *SiteContainerEntry {
	this := SiteContainerEntry{}
	this.Entry = entry
	return &this
}

// NewSiteContainerEntryWithDefaults instantiates a new SiteContainerEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteContainerEntryWithDefaults() *SiteContainerEntry {
	this := SiteContainerEntry{}
	return &this
}

// GetEntry returns the Entry field value
func (o *SiteContainerEntry) GetEntry() SiteContainer {
	if o == nil {
		var ret SiteContainer
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *SiteContainerEntry) GetEntryOk() (*SiteContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *SiteContainerEntry) SetEntry(v SiteContainer) {
	o.Entry = v
}

func (o SiteContainerEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteContainerEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	return toSerialize, nil
}

type NullableSiteContainerEntry struct {
	value *SiteContainerEntry
	isSet bool
}

func (v NullableSiteContainerEntry) Get() *SiteContainerEntry {
	return v.value
}

func (v *NullableSiteContainerEntry) Set(val *SiteContainerEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteContainerEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteContainerEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteContainerEntry(val *SiteContainerEntry) *NullableSiteContainerEntry {
	return &NullableSiteContainerEntry{value: val, isSet: true}
}

func (v NullableSiteContainerEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteContainerEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

