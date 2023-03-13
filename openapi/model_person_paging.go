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

// checks if the PersonPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonPaging{}

// PersonPaging struct for PersonPaging
type PersonPaging struct {
	List *PersonPagingList `json:"list,omitempty"`
}

// NewPersonPaging instantiates a new PersonPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonPaging() *PersonPaging {
	this := PersonPaging{}
	return &this
}

// NewPersonPagingWithDefaults instantiates a new PersonPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonPagingWithDefaults() *PersonPaging {
	this := PersonPaging{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *PersonPaging) GetList() PersonPagingList {
	if o == nil || IsNil(o.List) {
		var ret PersonPagingList
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonPaging) GetListOk() (*PersonPagingList, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *PersonPaging) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given PersonPagingList and assigns it to the List field.
func (o *PersonPaging) SetList(v PersonPagingList) {
	o.List = &v
}

func (o PersonPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullablePersonPaging struct {
	value *PersonPaging
	isSet bool
}

func (v NullablePersonPaging) Get() *PersonPaging {
	return v.value
}

func (v *NullablePersonPaging) Set(val *PersonPaging) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonPaging) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonPaging(val *PersonPaging) *NullablePersonPaging {
	return &NullablePersonPaging{value: val, isSet: true}
}

func (v NullablePersonPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


