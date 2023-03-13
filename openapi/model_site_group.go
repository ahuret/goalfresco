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

// checks if the SiteGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteGroup{}

// SiteGroup struct for SiteGroup
type SiteGroup struct {
	Id string `json:"id"`
	Group GroupMember `json:"group"`
	Role string `json:"role"`
}

// NewSiteGroup instantiates a new SiteGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteGroup(id string, group GroupMember, role string) *SiteGroup {
	this := SiteGroup{}
	this.Id = id
	this.Group = group
	this.Role = role
	return &this
}

// NewSiteGroupWithDefaults instantiates a new SiteGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteGroupWithDefaults() *SiteGroup {
	this := SiteGroup{}
	return &this
}

// GetId returns the Id field value
func (o *SiteGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SiteGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SiteGroup) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *SiteGroup) GetGroup() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SiteGroup) GetGroupOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SiteGroup) SetGroup(v GroupMember) {
	o.Group = v
}

// GetRole returns the Role field value
func (o *SiteGroup) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SiteGroup) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SiteGroup) SetRole(v string) {
	o.Role = v
}

func (o SiteGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableSiteGroup struct {
	value *SiteGroup
	isSet bool
}

func (v NullableSiteGroup) Get() *SiteGroup {
	return v.value
}

func (v *NullableSiteGroup) Set(val *SiteGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteGroup(val *SiteGroup) *NullableSiteGroup {
	return &NullableSiteGroup{value: val, isSet: true}
}

func (v NullableSiteGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


