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

// checks if the SitePaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SitePaging{}

// SitePaging struct for SitePaging
type SitePaging struct {
	List SitePagingList `json:"list"`
}

// NewSitePaging instantiates a new SitePaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSitePaging(list SitePagingList) *SitePaging {
	this := SitePaging{}
	this.List = list
	return &this
}

// NewSitePagingWithDefaults instantiates a new SitePaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSitePagingWithDefaults() *SitePaging {
	this := SitePaging{}
	return &this
}

// GetList returns the List field value
func (o *SitePaging) GetList() SitePagingList {
	if o == nil {
		var ret SitePagingList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *SitePaging) GetListOk() (*SitePagingList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *SitePaging) SetList(v SitePagingList) {
	o.List = v
}

func (o SitePaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SitePaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

type NullableSitePaging struct {
	value *SitePaging
	isSet bool
}

func (v NullableSitePaging) Get() *SitePaging {
	return v.value
}

func (v *NullableSitePaging) Set(val *SitePaging) {
	v.value = val
	v.isSet = true
}

func (v NullableSitePaging) IsSet() bool {
	return v.isSet
}

func (v *NullableSitePaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSitePaging(val *SitePaging) *NullableSitePaging {
	return &NullableSitePaging{value: val, isSet: true}
}

func (v NullableSitePaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSitePaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

