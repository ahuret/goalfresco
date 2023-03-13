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

// checks if the SiteContainerPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteContainerPaging{}

// SiteContainerPaging struct for SiteContainerPaging
type SiteContainerPaging struct {
	List SiteContainerPagingList `json:"list"`
}

// NewSiteContainerPaging instantiates a new SiteContainerPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteContainerPaging(list SiteContainerPagingList) *SiteContainerPaging {
	this := SiteContainerPaging{}
	this.List = list
	return &this
}

// NewSiteContainerPagingWithDefaults instantiates a new SiteContainerPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteContainerPagingWithDefaults() *SiteContainerPaging {
	this := SiteContainerPaging{}
	return &this
}

// GetList returns the List field value
func (o *SiteContainerPaging) GetList() SiteContainerPagingList {
	if o == nil {
		var ret SiteContainerPagingList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *SiteContainerPaging) GetListOk() (*SiteContainerPagingList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *SiteContainerPaging) SetList(v SiteContainerPagingList) {
	o.List = v
}

func (o SiteContainerPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteContainerPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

type NullableSiteContainerPaging struct {
	value *SiteContainerPaging
	isSet bool
}

func (v NullableSiteContainerPaging) Get() *SiteContainerPaging {
	return v.value
}

func (v *NullableSiteContainerPaging) Set(val *SiteContainerPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteContainerPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteContainerPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteContainerPaging(val *SiteContainerPaging) *NullableSiteContainerPaging {
	return &NullableSiteContainerPaging{value: val, isSet: true}
}

func (v NullableSiteContainerPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteContainerPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


