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

// checks if the PathInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathInfo{}

// PathInfo struct for PathInfo
type PathInfo struct {
	Elements []PathElement `json:"elements,omitempty"`
	Name *string `json:"name,omitempty"`
	IsComplete *bool `json:"isComplete,omitempty"`
}

// NewPathInfo instantiates a new PathInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathInfo() *PathInfo {
	this := PathInfo{}
	return &this
}

// NewPathInfoWithDefaults instantiates a new PathInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathInfoWithDefaults() *PathInfo {
	this := PathInfo{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *PathInfo) GetElements() []PathElement {
	if o == nil || IsNil(o.Elements) {
		var ret []PathElement
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathInfo) GetElementsOk() ([]PathElement, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *PathInfo) HasElements() bool {
	if o != nil && !IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []PathElement and assigns it to the Elements field.
func (o *PathInfo) SetElements(v []PathElement) {
	o.Elements = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PathInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PathInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PathInfo) SetName(v string) {
	o.Name = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *PathInfo) GetIsComplete() bool {
	if o == nil || IsNil(o.IsComplete) {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathInfo) GetIsCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsComplete) {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *PathInfo) HasIsComplete() bool {
	if o != nil && !IsNil(o.IsComplete) {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *PathInfo) SetIsComplete(v bool) {
	o.IsComplete = &v
}

func (o PathInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsComplete) {
		toSerialize["isComplete"] = o.IsComplete
	}
	return toSerialize, nil
}

type NullablePathInfo struct {
	value *PathInfo
	isSet bool
}

func (v NullablePathInfo) Get() *PathInfo {
	return v.value
}

func (v *NullablePathInfo) Set(val *PathInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePathInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePathInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathInfo(val *PathInfo) *NullablePathInfo {
	return &NullablePathInfo{value: val, isSet: true}
}

func (v NullablePathInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


