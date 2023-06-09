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

// checks if the ChildAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChildAssociation{}

// ChildAssociation struct for ChildAssociation
type ChildAssociation struct {
	ChildId string `json:"childId"`
	AssocType string `json:"assocType"`
}

// NewChildAssociation instantiates a new ChildAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChildAssociation(childId string, assocType string) *ChildAssociation {
	this := ChildAssociation{}
	this.ChildId = childId
	this.AssocType = assocType
	return &this
}

// NewChildAssociationWithDefaults instantiates a new ChildAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChildAssociationWithDefaults() *ChildAssociation {
	this := ChildAssociation{}
	return &this
}

// GetChildId returns the ChildId field value
func (o *ChildAssociation) GetChildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChildId
}

// GetChildIdOk returns a tuple with the ChildId field value
// and a boolean to check if the value has been set.
func (o *ChildAssociation) GetChildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildId, true
}

// SetChildId sets field value
func (o *ChildAssociation) SetChildId(v string) {
	o.ChildId = v
}

// GetAssocType returns the AssocType field value
func (o *ChildAssociation) GetAssocType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssocType
}

// GetAssocTypeOk returns a tuple with the AssocType field value
// and a boolean to check if the value has been set.
func (o *ChildAssociation) GetAssocTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssocType, true
}

// SetAssocType sets field value
func (o *ChildAssociation) SetAssocType(v string) {
	o.AssocType = v
}

func (o ChildAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChildAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["childId"] = o.ChildId
	toSerialize["assocType"] = o.AssocType
	return toSerialize, nil
}

type NullableChildAssociation struct {
	value *ChildAssociation
	isSet bool
}

func (v NullableChildAssociation) Get() *ChildAssociation {
	return v.value
}

func (v *NullableChildAssociation) Set(val *ChildAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableChildAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableChildAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChildAssociation(val *ChildAssociation) *NullableChildAssociation {
	return &NullableChildAssociation{value: val, isSet: true}
}

func (v NullableChildAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChildAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


