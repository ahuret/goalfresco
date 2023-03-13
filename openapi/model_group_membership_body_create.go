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

// checks if the GroupMembershipBodyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMembershipBodyCreate{}

// GroupMembershipBodyCreate struct for GroupMembershipBodyCreate
type GroupMembershipBodyCreate struct {
	Id string `json:"id"`
	MemberType string `json:"memberType"`
}

// NewGroupMembershipBodyCreate instantiates a new GroupMembershipBodyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMembershipBodyCreate(id string, memberType string) *GroupMembershipBodyCreate {
	this := GroupMembershipBodyCreate{}
	this.Id = id
	this.MemberType = memberType
	return &this
}

// NewGroupMembershipBodyCreateWithDefaults instantiates a new GroupMembershipBodyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMembershipBodyCreateWithDefaults() *GroupMembershipBodyCreate {
	this := GroupMembershipBodyCreate{}
	return &this
}

// GetId returns the Id field value
func (o *GroupMembershipBodyCreate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupMembershipBodyCreate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupMembershipBodyCreate) SetId(v string) {
	o.Id = v
}

// GetMemberType returns the MemberType field value
func (o *GroupMembershipBodyCreate) GetMemberType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MemberType
}

// GetMemberTypeOk returns a tuple with the MemberType field value
// and a boolean to check if the value has been set.
func (o *GroupMembershipBodyCreate) GetMemberTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberType, true
}

// SetMemberType sets field value
func (o *GroupMembershipBodyCreate) SetMemberType(v string) {
	o.MemberType = v
}

func (o GroupMembershipBodyCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMembershipBodyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["memberType"] = o.MemberType
	return toSerialize, nil
}

type NullableGroupMembershipBodyCreate struct {
	value *GroupMembershipBodyCreate
	isSet bool
}

func (v NullableGroupMembershipBodyCreate) Get() *GroupMembershipBodyCreate {
	return v.value
}

func (v *NullableGroupMembershipBodyCreate) Set(val *GroupMembershipBodyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMembershipBodyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMembershipBodyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMembershipBodyCreate(val *GroupMembershipBodyCreate) *NullableGroupMembershipBodyCreate {
	return &NullableGroupMembershipBodyCreate{value: val, isSet: true}
}

func (v NullableGroupMembershipBodyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMembershipBodyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


