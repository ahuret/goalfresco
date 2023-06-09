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

// checks if the VersionPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionPaging{}

// VersionPaging struct for VersionPaging
type VersionPaging struct {
	List *VersionPagingList `json:"list,omitempty"`
}

// NewVersionPaging instantiates a new VersionPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionPaging() *VersionPaging {
	this := VersionPaging{}
	return &this
}

// NewVersionPagingWithDefaults instantiates a new VersionPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionPagingWithDefaults() *VersionPaging {
	this := VersionPaging{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *VersionPaging) GetList() VersionPagingList {
	if o == nil || IsNil(o.List) {
		var ret VersionPagingList
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionPaging) GetListOk() (*VersionPagingList, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *VersionPaging) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given VersionPagingList and assigns it to the List field.
func (o *VersionPaging) SetList(v VersionPagingList) {
	o.List = &v
}

func (o VersionPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableVersionPaging struct {
	value *VersionPaging
	isSet bool
}

func (v NullableVersionPaging) Get() *VersionPaging {
	return v.value
}

func (v *NullableVersionPaging) Set(val *VersionPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionPaging(val *VersionPaging) *NullableVersionPaging {
	return &NullableVersionPaging{value: val, isSet: true}
}

func (v NullableVersionPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


