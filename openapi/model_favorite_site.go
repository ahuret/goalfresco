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

// checks if the FavoriteSite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FavoriteSite{}

// FavoriteSite struct for FavoriteSite
type FavoriteSite struct {
	Id string `json:"id"`
}

// NewFavoriteSite instantiates a new FavoriteSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFavoriteSite(id string) *FavoriteSite {
	this := FavoriteSite{}
	this.Id = id
	return &this
}

// NewFavoriteSiteWithDefaults instantiates a new FavoriteSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFavoriteSiteWithDefaults() *FavoriteSite {
	this := FavoriteSite{}
	return &this
}

// GetId returns the Id field value
func (o *FavoriteSite) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FavoriteSite) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FavoriteSite) SetId(v string) {
	o.Id = v
}

func (o FavoriteSite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FavoriteSite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableFavoriteSite struct {
	value *FavoriteSite
	isSet bool
}

func (v NullableFavoriteSite) Get() *FavoriteSite {
	return v.value
}

func (v *NullableFavoriteSite) Set(val *FavoriteSite) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoriteSite) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoriteSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoriteSite(val *FavoriteSite) *NullableFavoriteSite {
	return &NullableFavoriteSite{value: val, isSet: true}
}

func (v NullableFavoriteSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoriteSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


