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

// checks if the DeletedNodeBodyRestore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletedNodeBodyRestore{}

// DeletedNodeBodyRestore struct for DeletedNodeBodyRestore
type DeletedNodeBodyRestore struct {
	TargetParentId *string `json:"targetParentId,omitempty"`
	AssocType *string `json:"assocType,omitempty"`
}

// NewDeletedNodeBodyRestore instantiates a new DeletedNodeBodyRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedNodeBodyRestore() *DeletedNodeBodyRestore {
	this := DeletedNodeBodyRestore{}
	return &this
}

// NewDeletedNodeBodyRestoreWithDefaults instantiates a new DeletedNodeBodyRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedNodeBodyRestoreWithDefaults() *DeletedNodeBodyRestore {
	this := DeletedNodeBodyRestore{}
	return &this
}

// GetTargetParentId returns the TargetParentId field value if set, zero value otherwise.
func (o *DeletedNodeBodyRestore) GetTargetParentId() string {
	if o == nil || IsNil(o.TargetParentId) {
		var ret string
		return ret
	}
	return *o.TargetParentId
}

// GetTargetParentIdOk returns a tuple with the TargetParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedNodeBodyRestore) GetTargetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetParentId) {
		return nil, false
	}
	return o.TargetParentId, true
}

// HasTargetParentId returns a boolean if a field has been set.
func (o *DeletedNodeBodyRestore) HasTargetParentId() bool {
	if o != nil && !IsNil(o.TargetParentId) {
		return true
	}

	return false
}

// SetTargetParentId gets a reference to the given string and assigns it to the TargetParentId field.
func (o *DeletedNodeBodyRestore) SetTargetParentId(v string) {
	o.TargetParentId = &v
}

// GetAssocType returns the AssocType field value if set, zero value otherwise.
func (o *DeletedNodeBodyRestore) GetAssocType() string {
	if o == nil || IsNil(o.AssocType) {
		var ret string
		return ret
	}
	return *o.AssocType
}

// GetAssocTypeOk returns a tuple with the AssocType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedNodeBodyRestore) GetAssocTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssocType) {
		return nil, false
	}
	return o.AssocType, true
}

// HasAssocType returns a boolean if a field has been set.
func (o *DeletedNodeBodyRestore) HasAssocType() bool {
	if o != nil && !IsNil(o.AssocType) {
		return true
	}

	return false
}

// SetAssocType gets a reference to the given string and assigns it to the AssocType field.
func (o *DeletedNodeBodyRestore) SetAssocType(v string) {
	o.AssocType = &v
}

func (o DeletedNodeBodyRestore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletedNodeBodyRestore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetParentId) {
		toSerialize["targetParentId"] = o.TargetParentId
	}
	if !IsNil(o.AssocType) {
		toSerialize["assocType"] = o.AssocType
	}
	return toSerialize, nil
}

type NullableDeletedNodeBodyRestore struct {
	value *DeletedNodeBodyRestore
	isSet bool
}

func (v NullableDeletedNodeBodyRestore) Get() *DeletedNodeBodyRestore {
	return v.value
}

func (v *NullableDeletedNodeBodyRestore) Set(val *DeletedNodeBodyRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedNodeBodyRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletedNodeBodyRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedNodeBodyRestore(val *DeletedNodeBodyRestore) *NullableDeletedNodeBodyRestore {
	return &NullableDeletedNodeBodyRestore{value: val, isSet: true}
}

func (v NullableDeletedNodeBodyRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedNodeBodyRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

