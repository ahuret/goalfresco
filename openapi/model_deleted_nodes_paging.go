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

// checks if the DeletedNodesPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletedNodesPaging{}

// DeletedNodesPaging struct for DeletedNodesPaging
type DeletedNodesPaging struct {
	List *DeletedNodesPagingList `json:"list,omitempty"`
}

// NewDeletedNodesPaging instantiates a new DeletedNodesPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedNodesPaging() *DeletedNodesPaging {
	this := DeletedNodesPaging{}
	return &this
}

// NewDeletedNodesPagingWithDefaults instantiates a new DeletedNodesPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedNodesPagingWithDefaults() *DeletedNodesPaging {
	this := DeletedNodesPaging{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *DeletedNodesPaging) GetList() DeletedNodesPagingList {
	if o == nil || IsNil(o.List) {
		var ret DeletedNodesPagingList
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedNodesPaging) GetListOk() (*DeletedNodesPagingList, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *DeletedNodesPaging) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given DeletedNodesPagingList and assigns it to the List field.
func (o *DeletedNodesPaging) SetList(v DeletedNodesPagingList) {
	o.List = &v
}

func (o DeletedNodesPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletedNodesPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableDeletedNodesPaging struct {
	value *DeletedNodesPaging
	isSet bool
}

func (v NullableDeletedNodesPaging) Get() *DeletedNodesPaging {
	return v.value
}

func (v *NullableDeletedNodesPaging) Set(val *DeletedNodesPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedNodesPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletedNodesPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedNodesPaging(val *DeletedNodesPaging) *NullableDeletedNodesPaging {
	return &NullableDeletedNodesPaging{value: val, isSet: true}
}

func (v NullableDeletedNodesPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedNodesPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


