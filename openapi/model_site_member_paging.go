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

// checks if the SiteMemberPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteMemberPaging{}

// SiteMemberPaging struct for SiteMemberPaging
type SiteMemberPaging struct {
	List SiteMemberPagingList `json:"list"`
}

// NewSiteMemberPaging instantiates a new SiteMemberPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMemberPaging(list SiteMemberPagingList) *SiteMemberPaging {
	this := SiteMemberPaging{}
	this.List = list
	return &this
}

// NewSiteMemberPagingWithDefaults instantiates a new SiteMemberPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMemberPagingWithDefaults() *SiteMemberPaging {
	this := SiteMemberPaging{}
	return &this
}

// GetList returns the List field value
func (o *SiteMemberPaging) GetList() SiteMemberPagingList {
	if o == nil {
		var ret SiteMemberPagingList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *SiteMemberPaging) GetListOk() (*SiteMemberPagingList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *SiteMemberPaging) SetList(v SiteMemberPagingList) {
	o.List = v
}

func (o SiteMemberPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteMemberPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

type NullableSiteMemberPaging struct {
	value *SiteMemberPaging
	isSet bool
}

func (v NullableSiteMemberPaging) Get() *SiteMemberPaging {
	return v.value
}

func (v *NullableSiteMemberPaging) Set(val *SiteMemberPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMemberPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMemberPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMemberPaging(val *SiteMemberPaging) *NullableSiteMemberPaging {
	return &NullableSiteMemberPaging{value: val, isSet: true}
}

func (v NullableSiteMemberPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMemberPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


