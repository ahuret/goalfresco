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

// checks if the GroupMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMember{}

// GroupMember struct for GroupMember
type GroupMember struct {
	Id string `json:"id"`
	DisplayName string `json:"displayName"`
	MemberType string `json:"memberType"`
}

// NewGroupMember instantiates a new GroupMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMember(id string, displayName string, memberType string) *GroupMember {
	this := GroupMember{}
	this.Id = id
	this.DisplayName = displayName
	this.MemberType = memberType
	return &this
}

// NewGroupMemberWithDefaults instantiates a new GroupMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMemberWithDefaults() *GroupMember {
	this := GroupMember{}
	return &this
}

// GetId returns the Id field value
func (o *GroupMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupMember) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *GroupMember) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GroupMember) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GroupMember) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetMemberType returns the MemberType field value
func (o *GroupMember) GetMemberType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberType
}

// GetMemberTypeOk returns a tuple with the MemberType field value
// and a boolean to check if the value has been set.
func (o *GroupMember) GetMemberTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberType, true
}

// SetMemberType sets field value
func (o *GroupMember) SetMemberType(v string) {
	o.MemberType = v
}

func (o GroupMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["displayName"] = o.DisplayName
	toSerialize["memberType"] = o.MemberType
	return toSerialize, nil
}

type NullableGroupMember struct {
	value *GroupMember
	isSet bool
}

func (v NullableGroupMember) Get() *GroupMember {
	return v.value
}

func (v *NullableGroupMember) Set(val *GroupMember) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMember) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMember(val *GroupMember) *NullableGroupMember {
	return &NullableGroupMember{value: val, isSet: true}
}

func (v NullableGroupMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


