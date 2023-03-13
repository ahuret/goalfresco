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

// checks if the AuditAppPagingList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditAppPagingList{}

// AuditAppPagingList struct for AuditAppPagingList
type AuditAppPagingList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Entries []AuditAppEntry `json:"entries,omitempty"`
}

// NewAuditAppPagingList instantiates a new AuditAppPagingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditAppPagingList() *AuditAppPagingList {
	this := AuditAppPagingList{}
	return &this
}

// NewAuditAppPagingListWithDefaults instantiates a new AuditAppPagingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditAppPagingListWithDefaults() *AuditAppPagingList {
	this := AuditAppPagingList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AuditAppPagingList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAppPagingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AuditAppPagingList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *AuditAppPagingList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *AuditAppPagingList) GetEntries() []AuditAppEntry {
	if o == nil || IsNil(o.Entries) {
		var ret []AuditAppEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAppPagingList) GetEntriesOk() ([]AuditAppEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *AuditAppPagingList) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []AuditAppEntry and assigns it to the Entries field.
func (o *AuditAppPagingList) SetEntries(v []AuditAppEntry) {
	o.Entries = v
}

func (o AuditAppPagingList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditAppPagingList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableAuditAppPagingList struct {
	value *AuditAppPagingList
	isSet bool
}

func (v NullableAuditAppPagingList) Get() *AuditAppPagingList {
	return v.value
}

func (v *NullableAuditAppPagingList) Set(val *AuditAppPagingList) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditAppPagingList) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditAppPagingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditAppPagingList(val *AuditAppPagingList) *NullableAuditAppPagingList {
	return &NullableAuditAppPagingList{value: val, isSet: true}
}

func (v NullableAuditAppPagingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditAppPagingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


