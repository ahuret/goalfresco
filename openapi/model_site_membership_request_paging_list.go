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

// checks if the SiteMembershipRequestPagingList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteMembershipRequestPagingList{}

// SiteMembershipRequestPagingList struct for SiteMembershipRequestPagingList
type SiteMembershipRequestPagingList struct {
	Pagination Pagination `json:"pagination"`
	Entries []SiteMembershipRequestEntry `json:"entries"`
}

// NewSiteMembershipRequestPagingList instantiates a new SiteMembershipRequestPagingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMembershipRequestPagingList(pagination Pagination, entries []SiteMembershipRequestEntry) *SiteMembershipRequestPagingList {
	this := SiteMembershipRequestPagingList{}
	this.Pagination = pagination
	this.Entries = entries
	return &this
}

// NewSiteMembershipRequestPagingListWithDefaults instantiates a new SiteMembershipRequestPagingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMembershipRequestPagingListWithDefaults() *SiteMembershipRequestPagingList {
	this := SiteMembershipRequestPagingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *SiteMembershipRequestPagingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *SiteMembershipRequestPagingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *SiteMembershipRequestPagingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetEntries returns the Entries field value
func (o *SiteMembershipRequestPagingList) GetEntries() []SiteMembershipRequestEntry {
	if o == nil {
		var ret []SiteMembershipRequestEntry
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *SiteMembershipRequestPagingList) GetEntriesOk() ([]SiteMembershipRequestEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *SiteMembershipRequestPagingList) SetEntries(v []SiteMembershipRequestEntry) {
	o.Entries = v
}

func (o SiteMembershipRequestPagingList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteMembershipRequestPagingList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["entries"] = o.Entries
	return toSerialize, nil
}

type NullableSiteMembershipRequestPagingList struct {
	value *SiteMembershipRequestPagingList
	isSet bool
}

func (v NullableSiteMembershipRequestPagingList) Get() *SiteMembershipRequestPagingList {
	return v.value
}

func (v *NullableSiteMembershipRequestPagingList) Set(val *SiteMembershipRequestPagingList) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMembershipRequestPagingList) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMembershipRequestPagingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMembershipRequestPagingList(val *SiteMembershipRequestPagingList) *NullableSiteMembershipRequestPagingList {
	return &NullableSiteMembershipRequestPagingList{value: val, isSet: true}
}

func (v NullableSiteMembershipRequestPagingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMembershipRequestPagingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

