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

// checks if the SiteRolePaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteRolePaging{}

// SiteRolePaging struct for SiteRolePaging
type SiteRolePaging struct {
	List SiteRolePagingList `json:"list"`
}

// NewSiteRolePaging instantiates a new SiteRolePaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteRolePaging(list SiteRolePagingList) *SiteRolePaging {
	this := SiteRolePaging{}
	this.List = list
	return &this
}

// NewSiteRolePagingWithDefaults instantiates a new SiteRolePaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteRolePagingWithDefaults() *SiteRolePaging {
	this := SiteRolePaging{}
	return &this
}

// GetList returns the List field value
func (o *SiteRolePaging) GetList() SiteRolePagingList {
	if o == nil {
		var ret SiteRolePagingList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *SiteRolePaging) GetListOk() (*SiteRolePagingList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *SiteRolePaging) SetList(v SiteRolePagingList) {
	o.List = v
}

func (o SiteRolePaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteRolePaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

type NullableSiteRolePaging struct {
	value *SiteRolePaging
	isSet bool
}

func (v NullableSiteRolePaging) Get() *SiteRolePaging {
	return v.value
}

func (v *NullableSiteRolePaging) Set(val *SiteRolePaging) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteRolePaging) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteRolePaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteRolePaging(val *SiteRolePaging) *NullableSiteRolePaging {
	return &NullableSiteRolePaging{value: val, isSet: true}
}

func (v NullableSiteRolePaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteRolePaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


