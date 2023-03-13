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

// checks if the GroupPagingList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupPagingList{}

// GroupPagingList struct for GroupPagingList
type GroupPagingList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Entries []GroupEntry `json:"entries,omitempty"`
}

// NewGroupPagingList instantiates a new GroupPagingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPagingList() *GroupPagingList {
	this := GroupPagingList{}
	return &this
}

// NewGroupPagingListWithDefaults instantiates a new GroupPagingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPagingListWithDefaults() *GroupPagingList {
	this := GroupPagingList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GroupPagingList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPagingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GroupPagingList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GroupPagingList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *GroupPagingList) GetEntries() []GroupEntry {
	if o == nil || IsNil(o.Entries) {
		var ret []GroupEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPagingList) GetEntriesOk() ([]GroupEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *GroupPagingList) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []GroupEntry and assigns it to the Entries field.
func (o *GroupPagingList) SetEntries(v []GroupEntry) {
	o.Entries = v
}

func (o GroupPagingList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupPagingList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableGroupPagingList struct {
	value *GroupPagingList
	isSet bool
}

func (v NullableGroupPagingList) Get() *GroupPagingList {
	return v.value
}

func (v *NullableGroupPagingList) Set(val *GroupPagingList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPagingList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPagingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPagingList(val *GroupPagingList) *NullableGroupPagingList {
	return &NullableGroupPagingList{value: val, isSet: true}
}

func (v NullableGroupPagingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPagingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

