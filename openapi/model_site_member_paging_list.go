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

// checks if the SiteMemberPagingList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteMemberPagingList{}

// SiteMemberPagingList struct for SiteMemberPagingList
type SiteMemberPagingList struct {
	Pagination Pagination `json:"pagination"`
	Entries []SiteMemberEntry `json:"entries"`
}

// NewSiteMemberPagingList instantiates a new SiteMemberPagingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMemberPagingList(pagination Pagination, entries []SiteMemberEntry) *SiteMemberPagingList {
	this := SiteMemberPagingList{}
	this.Pagination = pagination
	this.Entries = entries
	return &this
}

// NewSiteMemberPagingListWithDefaults instantiates a new SiteMemberPagingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMemberPagingListWithDefaults() *SiteMemberPagingList {
	this := SiteMemberPagingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *SiteMemberPagingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *SiteMemberPagingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *SiteMemberPagingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetEntries returns the Entries field value
func (o *SiteMemberPagingList) GetEntries() []SiteMemberEntry {
	if o == nil {
		var ret []SiteMemberEntry
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *SiteMemberPagingList) GetEntriesOk() ([]SiteMemberEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *SiteMemberPagingList) SetEntries(v []SiteMemberEntry) {
	o.Entries = v
}

func (o SiteMemberPagingList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteMemberPagingList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["entries"] = o.Entries
	return toSerialize, nil
}

type NullableSiteMemberPagingList struct {
	value *SiteMemberPagingList
	isSet bool
}

func (v NullableSiteMemberPagingList) Get() *SiteMemberPagingList {
	return v.value
}

func (v *NullableSiteMemberPagingList) Set(val *SiteMemberPagingList) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMemberPagingList) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMemberPagingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMemberPagingList(val *SiteMemberPagingList) *NullableSiteMemberPagingList {
	return &NullableSiteMemberPagingList{value: val, isSet: true}
}

func (v NullableSiteMemberPagingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMemberPagingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


