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

// checks if the GroupMemberEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMemberEntry{}

// GroupMemberEntry struct for GroupMemberEntry
type GroupMemberEntry struct {
	Entry GroupMember `json:"entry"`
}

// NewGroupMemberEntry instantiates a new GroupMemberEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMemberEntry(entry GroupMember) *GroupMemberEntry {
	this := GroupMemberEntry{}
	this.Entry = entry
	return &this
}

// NewGroupMemberEntryWithDefaults instantiates a new GroupMemberEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMemberEntryWithDefaults() *GroupMemberEntry {
	this := GroupMemberEntry{}
	return &this
}

// GetEntry returns the Entry field value
func (o *GroupMemberEntry) GetEntry() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *GroupMemberEntry) GetEntryOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *GroupMemberEntry) SetEntry(v GroupMember) {
	o.Entry = v
}

func (o GroupMemberEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMemberEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	return toSerialize, nil
}

type NullableGroupMemberEntry struct {
	value *GroupMemberEntry
	isSet bool
}

func (v NullableGroupMemberEntry) Get() *GroupMemberEntry {
	return v.value
}

func (v *NullableGroupMemberEntry) Set(val *GroupMemberEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMemberEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMemberEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMemberEntry(val *GroupMemberEntry) *NullableGroupMemberEntry {
	return &NullableGroupMemberEntry{value: val, isSet: true}
}

func (v NullableGroupMemberEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMemberEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


