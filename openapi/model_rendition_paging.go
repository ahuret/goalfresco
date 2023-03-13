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

// checks if the RenditionPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenditionPaging{}

// RenditionPaging struct for RenditionPaging
type RenditionPaging struct {
	List *RenditionPagingList `json:"list,omitempty"`
}

// NewRenditionPaging instantiates a new RenditionPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenditionPaging() *RenditionPaging {
	this := RenditionPaging{}
	return &this
}

// NewRenditionPagingWithDefaults instantiates a new RenditionPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenditionPagingWithDefaults() *RenditionPaging {
	this := RenditionPaging{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *RenditionPaging) GetList() RenditionPagingList {
	if o == nil || IsNil(o.List) {
		var ret RenditionPagingList
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenditionPaging) GetListOk() (*RenditionPagingList, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *RenditionPaging) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given RenditionPagingList and assigns it to the List field.
func (o *RenditionPaging) SetList(v RenditionPagingList) {
	o.List = &v
}

func (o RenditionPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenditionPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableRenditionPaging struct {
	value *RenditionPaging
	isSet bool
}

func (v NullableRenditionPaging) Get() *RenditionPaging {
	return v.value
}

func (v *NullableRenditionPaging) Set(val *RenditionPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableRenditionPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableRenditionPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenditionPaging(val *RenditionPaging) *NullableRenditionPaging {
	return &NullableRenditionPaging{value: val, isSet: true}
}

func (v NullableRenditionPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenditionPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


