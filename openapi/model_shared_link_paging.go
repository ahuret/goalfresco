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

// checks if the SharedLinkPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedLinkPaging{}

// SharedLinkPaging struct for SharedLinkPaging
type SharedLinkPaging struct {
	List SharedLinkPagingList `json:"list"`
}

// NewSharedLinkPaging instantiates a new SharedLinkPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedLinkPaging(list SharedLinkPagingList) *SharedLinkPaging {
	this := SharedLinkPaging{}
	this.List = list
	return &this
}

// NewSharedLinkPagingWithDefaults instantiates a new SharedLinkPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedLinkPagingWithDefaults() *SharedLinkPaging {
	this := SharedLinkPaging{}
	return &this
}

// GetList returns the List field value
func (o *SharedLinkPaging) GetList() SharedLinkPagingList {
	if o == nil {
		var ret SharedLinkPagingList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *SharedLinkPaging) GetListOk() (*SharedLinkPagingList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *SharedLinkPaging) SetList(v SharedLinkPagingList) {
	o.List = v
}

func (o SharedLinkPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedLinkPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

type NullableSharedLinkPaging struct {
	value *SharedLinkPaging
	isSet bool
}

func (v NullableSharedLinkPaging) Get() *SharedLinkPaging {
	return v.value
}

func (v *NullableSharedLinkPaging) Set(val *SharedLinkPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedLinkPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedLinkPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedLinkPaging(val *SharedLinkPaging) *NullableSharedLinkPaging {
	return &NullableSharedLinkPaging{value: val, isSet: true}
}

func (v NullableSharedLinkPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedLinkPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

