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

// checks if the PathElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathElement{}

// PathElement struct for PathElement
type PathElement struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	NodeType *string `json:"nodeType,omitempty"`
	AspectNames []string `json:"aspectNames,omitempty"`
}

// NewPathElement instantiates a new PathElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathElement() *PathElement {
	this := PathElement{}
	return &this
}

// NewPathElementWithDefaults instantiates a new PathElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathElementWithDefaults() *PathElement {
	this := PathElement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PathElement) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathElement) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PathElement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PathElement) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PathElement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathElement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PathElement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PathElement) SetName(v string) {
	o.Name = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *PathElement) GetNodeType() string {
	if o == nil || IsNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathElement) GetNodeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NodeType) {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *PathElement) HasNodeType() bool {
	if o != nil && !IsNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *PathElement) SetNodeType(v string) {
	o.NodeType = &v
}

// GetAspectNames returns the AspectNames field value if set, zero value otherwise.
func (o *PathElement) GetAspectNames() []string {
	if o == nil || IsNil(o.AspectNames) {
		var ret []string
		return ret
	}
	return o.AspectNames
}

// GetAspectNamesOk returns a tuple with the AspectNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathElement) GetAspectNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AspectNames) {
		return nil, false
	}
	return o.AspectNames, true
}

// HasAspectNames returns a boolean if a field has been set.
func (o *PathElement) HasAspectNames() bool {
	if o != nil && !IsNil(o.AspectNames) {
		return true
	}

	return false
}

// SetAspectNames gets a reference to the given []string and assigns it to the AspectNames field.
func (o *PathElement) SetAspectNames(v []string) {
	o.AspectNames = v
}

func (o PathElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NodeType) {
		toSerialize["nodeType"] = o.NodeType
	}
	if !IsNil(o.AspectNames) {
		toSerialize["aspectNames"] = o.AspectNames
	}
	return toSerialize, nil
}

type NullablePathElement struct {
	value *PathElement
	isSet bool
}

func (v NullablePathElement) Get() *PathElement {
	return v.value
}

func (v *NullablePathElement) Set(val *PathElement) {
	v.value = val
	v.isSet = true
}

func (v NullablePathElement) IsSet() bool {
	return v.isSet
}

func (v *NullablePathElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathElement(val *PathElement) *NullablePathElement {
	return &NullablePathElement{value: val, isSet: true}
}

func (v NullablePathElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


