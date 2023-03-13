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

// checks if the RenditionBodyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenditionBodyCreate{}

// RenditionBodyCreate struct for RenditionBodyCreate
type RenditionBodyCreate struct {
	Id string `json:"id"`
}

// NewRenditionBodyCreate instantiates a new RenditionBodyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenditionBodyCreate(id string) *RenditionBodyCreate {
	this := RenditionBodyCreate{}
	this.Id = id
	return &this
}

// NewRenditionBodyCreateWithDefaults instantiates a new RenditionBodyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenditionBodyCreateWithDefaults() *RenditionBodyCreate {
	this := RenditionBodyCreate{}
	return &this
}

// GetId returns the Id field value
func (o *RenditionBodyCreate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RenditionBodyCreate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RenditionBodyCreate) SetId(v string) {
	o.Id = v
}

func (o RenditionBodyCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenditionBodyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableRenditionBodyCreate struct {
	value *RenditionBodyCreate
	isSet bool
}

func (v NullableRenditionBodyCreate) Get() *RenditionBodyCreate {
	return v.value
}

func (v *NullableRenditionBodyCreate) Set(val *RenditionBodyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRenditionBodyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRenditionBodyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenditionBodyCreate(val *RenditionBodyCreate) *NullableRenditionBodyCreate {
	return &NullableRenditionBodyCreate{value: val, isSet: true}
}

func (v NullableRenditionBodyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenditionBodyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


