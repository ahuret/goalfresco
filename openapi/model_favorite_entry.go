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

// checks if the FavoriteEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FavoriteEntry{}

// FavoriteEntry struct for FavoriteEntry
type FavoriteEntry struct {
	Entry Favorite `json:"entry"`
}

// NewFavoriteEntry instantiates a new FavoriteEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFavoriteEntry(entry Favorite) *FavoriteEntry {
	this := FavoriteEntry{}
	this.Entry = entry
	return &this
}

// NewFavoriteEntryWithDefaults instantiates a new FavoriteEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFavoriteEntryWithDefaults() *FavoriteEntry {
	this := FavoriteEntry{}
	return &this
}

// GetEntry returns the Entry field value
func (o *FavoriteEntry) GetEntry() Favorite {
	if o == nil {
		var ret Favorite
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *FavoriteEntry) GetEntryOk() (*Favorite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *FavoriteEntry) SetEntry(v Favorite) {
	o.Entry = v
}

func (o FavoriteEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FavoriteEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	return toSerialize, nil
}

type NullableFavoriteEntry struct {
	value *FavoriteEntry
	isSet bool
}

func (v NullableFavoriteEntry) Get() *FavoriteEntry {
	return v.value
}

func (v *NullableFavoriteEntry) Set(val *FavoriteEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoriteEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoriteEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoriteEntry(val *FavoriteEntry) *NullableFavoriteEntry {
	return &NullableFavoriteEntry{value: val, isSet: true}
}

func (v NullableFavoriteEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoriteEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

