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

// checks if the NodePaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodePaging{}

// NodePaging struct for NodePaging
type NodePaging struct {
	List *NodePagingList `json:"list,omitempty"`
}

// NewNodePaging instantiates a new NodePaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodePaging() *NodePaging {
	this := NodePaging{}
	return &this
}

// NewNodePagingWithDefaults instantiates a new NodePaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodePagingWithDefaults() *NodePaging {
	this := NodePaging{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *NodePaging) GetList() NodePagingList {
	if o == nil || IsNil(o.List) {
		var ret NodePagingList
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePaging) GetListOk() (*NodePagingList, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *NodePaging) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given NodePagingList and assigns it to the List field.
func (o *NodePaging) SetList(v NodePagingList) {
	o.List = &v
}

func (o NodePaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodePaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableNodePaging struct {
	value *NodePaging
	isSet bool
}

func (v NullableNodePaging) Get() *NodePaging {
	return v.value
}

func (v *NullableNodePaging) Set(val *NodePaging) {
	v.value = val
	v.isSet = true
}

func (v NullableNodePaging) IsSet() bool {
	return v.isSet
}

func (v *NullableNodePaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodePaging(val *NodePaging) *NullableNodePaging {
	return &NullableNodePaging{value: val, isSet: true}
}

func (v NullableNodePaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodePaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


