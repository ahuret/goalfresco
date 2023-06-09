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

// checks if the Association type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Association{}

// Association struct for Association
type Association struct {
	TargetId string `json:"targetId"`
	AssocType string `json:"assocType"`
}

// NewAssociation instantiates a new Association object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociation(targetId string, assocType string) *Association {
	this := Association{}
	this.TargetId = targetId
	this.AssocType = assocType
	return &this
}

// NewAssociationWithDefaults instantiates a new Association object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationWithDefaults() *Association {
	this := Association{}
	return &this
}

// GetTargetId returns the TargetId field value
func (o *Association) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *Association) GetTargetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *Association) SetTargetId(v string) {
	o.TargetId = v
}

// GetAssocType returns the AssocType field value
func (o *Association) GetAssocType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssocType
}

// GetAssocTypeOk returns a tuple with the AssocType field value
// and a boolean to check if the value has been set.
func (o *Association) GetAssocTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssocType, true
}

// SetAssocType sets field value
func (o *Association) SetAssocType(v string) {
	o.AssocType = v
}

func (o Association) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Association) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetId"] = o.TargetId
	toSerialize["assocType"] = o.AssocType
	return toSerialize, nil
}

type NullableAssociation struct {
	value *Association
	isSet bool
}

func (v NullableAssociation) Get() *Association {
	return v.value
}

func (v *NullableAssociation) Set(val *Association) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociation(val *Association) *NullableAssociation {
	return &NullableAssociation{value: val, isSet: true}
}

func (v NullableAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


