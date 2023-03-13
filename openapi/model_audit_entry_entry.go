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

// checks if the AuditEntryEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditEntryEntry{}

// AuditEntryEntry struct for AuditEntryEntry
type AuditEntryEntry struct {
	Entry *AuditEntry `json:"entry,omitempty"`
}

// NewAuditEntryEntry instantiates a new AuditEntryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditEntryEntry() *AuditEntryEntry {
	this := AuditEntryEntry{}
	return &this
}

// NewAuditEntryEntryWithDefaults instantiates a new AuditEntryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditEntryEntryWithDefaults() *AuditEntryEntry {
	this := AuditEntryEntry{}
	return &this
}

// GetEntry returns the Entry field value if set, zero value otherwise.
func (o *AuditEntryEntry) GetEntry() AuditEntry {
	if o == nil || IsNil(o.Entry) {
		var ret AuditEntry
		return ret
	}
	return *o.Entry
}

// GetEntryOk returns a tuple with the Entry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEntryEntry) GetEntryOk() (*AuditEntry, bool) {
	if o == nil || IsNil(o.Entry) {
		return nil, false
	}
	return o.Entry, true
}

// HasEntry returns a boolean if a field has been set.
func (o *AuditEntryEntry) HasEntry() bool {
	if o != nil && !IsNil(o.Entry) {
		return true
	}

	return false
}

// SetEntry gets a reference to the given AuditEntry and assigns it to the Entry field.
func (o *AuditEntryEntry) SetEntry(v AuditEntry) {
	o.Entry = &v
}

func (o AuditEntryEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditEntryEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entry) {
		toSerialize["entry"] = o.Entry
	}
	return toSerialize, nil
}

type NullableAuditEntryEntry struct {
	value *AuditEntryEntry
	isSet bool
}

func (v NullableAuditEntryEntry) Get() *AuditEntryEntry {
	return v.value
}

func (v *NullableAuditEntryEntry) Set(val *AuditEntryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditEntryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditEntryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditEntryEntry(val *AuditEntryEntry) *NullableAuditEntryEntry {
	return &NullableAuditEntryEntry{value: val, isSet: true}
}

func (v NullableAuditEntryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditEntryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

