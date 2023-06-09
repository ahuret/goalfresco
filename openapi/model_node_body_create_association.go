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

// checks if the NodeBodyCreateAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeBodyCreateAssociation{}

// NodeBodyCreateAssociation struct for NodeBodyCreateAssociation
type NodeBodyCreateAssociation struct {
	AssocType *string `json:"assocType,omitempty"`
}

// NewNodeBodyCreateAssociation instantiates a new NodeBodyCreateAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeBodyCreateAssociation() *NodeBodyCreateAssociation {
	this := NodeBodyCreateAssociation{}
	return &this
}

// NewNodeBodyCreateAssociationWithDefaults instantiates a new NodeBodyCreateAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeBodyCreateAssociationWithDefaults() *NodeBodyCreateAssociation {
	this := NodeBodyCreateAssociation{}
	return &this
}

// GetAssocType returns the AssocType field value if set, zero value otherwise.
func (o *NodeBodyCreateAssociation) GetAssocType() string {
	if o == nil || IsNil(o.AssocType) {
		var ret string
		return ret
	}
	return *o.AssocType
}

// GetAssocTypeOk returns a tuple with the AssocType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBodyCreateAssociation) GetAssocTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssocType) {
		return nil, false
	}
	return o.AssocType, true
}

// HasAssocType returns a boolean if a field has been set.
func (o *NodeBodyCreateAssociation) HasAssocType() bool {
	if o != nil && !IsNil(o.AssocType) {
		return true
	}

	return false
}

// SetAssocType gets a reference to the given string and assigns it to the AssocType field.
func (o *NodeBodyCreateAssociation) SetAssocType(v string) {
	o.AssocType = &v
}

func (o NodeBodyCreateAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeBodyCreateAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssocType) {
		toSerialize["assocType"] = o.AssocType
	}
	return toSerialize, nil
}

type NullableNodeBodyCreateAssociation struct {
	value *NodeBodyCreateAssociation
	isSet bool
}

func (v NullableNodeBodyCreateAssociation) Get() *NodeBodyCreateAssociation {
	return v.value
}

func (v *NullableNodeBodyCreateAssociation) Set(val *NodeBodyCreateAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeBodyCreateAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeBodyCreateAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeBodyCreateAssociation(val *NodeBodyCreateAssociation) *NullableNodeBodyCreateAssociation {
	return &NullableNodeBodyCreateAssociation{value: val, isSet: true}
}

func (v NullableNodeBodyCreateAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeBodyCreateAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


