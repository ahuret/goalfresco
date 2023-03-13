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

// checks if the GroupMemberPagingList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMemberPagingList{}

// GroupMemberPagingList struct for GroupMemberPagingList
type GroupMemberPagingList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Entries []GroupMemberEntry `json:"entries,omitempty"`
}

// NewGroupMemberPagingList instantiates a new GroupMemberPagingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMemberPagingList() *GroupMemberPagingList {
	this := GroupMemberPagingList{}
	return &this
}

// NewGroupMemberPagingListWithDefaults instantiates a new GroupMemberPagingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMemberPagingListWithDefaults() *GroupMemberPagingList {
	this := GroupMemberPagingList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GroupMemberPagingList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMemberPagingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GroupMemberPagingList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GroupMemberPagingList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *GroupMemberPagingList) GetEntries() []GroupMemberEntry {
	if o == nil || IsNil(o.Entries) {
		var ret []GroupMemberEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMemberPagingList) GetEntriesOk() ([]GroupMemberEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *GroupMemberPagingList) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []GroupMemberEntry and assigns it to the Entries field.
func (o *GroupMemberPagingList) SetEntries(v []GroupMemberEntry) {
	o.Entries = v
}

func (o GroupMemberPagingList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMemberPagingList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	return toSerialize, nil
}

type NullableGroupMemberPagingList struct {
	value *GroupMemberPagingList
	isSet bool
}

func (v NullableGroupMemberPagingList) Get() *GroupMemberPagingList {
	return v.value
}

func (v *NullableGroupMemberPagingList) Set(val *GroupMemberPagingList) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMemberPagingList) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMemberPagingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMemberPagingList(val *GroupMemberPagingList) *NullableGroupMemberPagingList {
	return &NullableGroupMemberPagingList{value: val, isSet: true}
}

func (v NullableGroupMemberPagingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMemberPagingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

