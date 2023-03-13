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

// checks if the GroupBodyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupBodyCreate{}

// GroupBodyCreate struct for GroupBodyCreate
type GroupBodyCreate struct {
	Id string `json:"id"`
	DisplayName string `json:"displayName"`
	ParentIds []string `json:"parentIds,omitempty"`
}

// NewGroupBodyCreate instantiates a new GroupBodyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupBodyCreate(id string, displayName string) *GroupBodyCreate {
	this := GroupBodyCreate{}
	this.Id = id
	this.DisplayName = displayName
	return &this
}

// NewGroupBodyCreateWithDefaults instantiates a new GroupBodyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupBodyCreateWithDefaults() *GroupBodyCreate {
	this := GroupBodyCreate{}
	return &this
}

// GetId returns the Id field value
func (o *GroupBodyCreate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupBodyCreate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupBodyCreate) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *GroupBodyCreate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GroupBodyCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GroupBodyCreate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetParentIds returns the ParentIds field value if set, zero value otherwise.
func (o *GroupBodyCreate) GetParentIds() []string {
	if o == nil || IsNil(o.ParentIds) {
		var ret []string
		return ret
	}
	return o.ParentIds
}

// GetParentIdsOk returns a tuple with the ParentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupBodyCreate) GetParentIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentIds) {
		return nil, false
	}
	return o.ParentIds, true
}

// HasParentIds returns a boolean if a field has been set.
func (o *GroupBodyCreate) HasParentIds() bool {
	if o != nil && !IsNil(o.ParentIds) {
		return true
	}

	return false
}

// SetParentIds gets a reference to the given []string and assigns it to the ParentIds field.
func (o *GroupBodyCreate) SetParentIds(v []string) {
	o.ParentIds = v
}

func (o GroupBodyCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupBodyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.ParentIds) {
		toSerialize["parentIds"] = o.ParentIds
	}
	return toSerialize, nil
}

type NullableGroupBodyCreate struct {
	value *GroupBodyCreate
	isSet bool
}

func (v NullableGroupBodyCreate) Get() *GroupBodyCreate {
	return v.value
}

func (v *NullableGroupBodyCreate) Set(val *GroupBodyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupBodyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupBodyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupBodyCreate(val *GroupBodyCreate) *NullableGroupBodyCreate {
	return &NullableGroupBodyCreate{value: val, isSet: true}
}

func (v NullableGroupBodyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupBodyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


