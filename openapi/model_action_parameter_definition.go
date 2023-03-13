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

// checks if the ActionParameterDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionParameterDefinition{}

// ActionParameterDefinition struct for ActionParameterDefinition
type ActionParameterDefinition struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	MultiValued *bool `json:"multiValued,omitempty"`
	Mandatory *bool `json:"mandatory,omitempty"`
	DisplayLabel *string `json:"displayLabel,omitempty"`
}

// NewActionParameterDefinition instantiates a new ActionParameterDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionParameterDefinition() *ActionParameterDefinition {
	this := ActionParameterDefinition{}
	return &this
}

// NewActionParameterDefinitionWithDefaults instantiates a new ActionParameterDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionParameterDefinitionWithDefaults() *ActionParameterDefinition {
	this := ActionParameterDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActionParameterDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionParameterDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActionParameterDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActionParameterDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActionParameterDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionParameterDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActionParameterDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActionParameterDefinition) SetType(v string) {
	o.Type = &v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *ActionParameterDefinition) GetMultiValued() bool {
	if o == nil || IsNil(o.MultiValued) {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionParameterDefinition) GetMultiValuedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiValued) {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *ActionParameterDefinition) HasMultiValued() bool {
	if o != nil && !IsNil(o.MultiValued) {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *ActionParameterDefinition) SetMultiValued(v bool) {
	o.MultiValued = &v
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise.
func (o *ActionParameterDefinition) GetMandatory() bool {
	if o == nil || IsNil(o.Mandatory) {
		var ret bool
		return ret
	}
	return *o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionParameterDefinition) GetMandatoryOk() (*bool, bool) {
	if o == nil || IsNil(o.Mandatory) {
		return nil, false
	}
	return o.Mandatory, true
}

// HasMandatory returns a boolean if a field has been set.
func (o *ActionParameterDefinition) HasMandatory() bool {
	if o != nil && !IsNil(o.Mandatory) {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given bool and assigns it to the Mandatory field.
func (o *ActionParameterDefinition) SetMandatory(v bool) {
	o.Mandatory = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *ActionParameterDefinition) GetDisplayLabel() string {
	if o == nil || IsNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionParameterDefinition) GetDisplayLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *ActionParameterDefinition) HasDisplayLabel() bool {
	if o != nil && !IsNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *ActionParameterDefinition) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

func (o ActionParameterDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionParameterDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.MultiValued) {
		toSerialize["multiValued"] = o.MultiValued
	}
	if !IsNil(o.Mandatory) {
		toSerialize["mandatory"] = o.Mandatory
	}
	if !IsNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	return toSerialize, nil
}

type NullableActionParameterDefinition struct {
	value *ActionParameterDefinition
	isSet bool
}

func (v NullableActionParameterDefinition) Get() *ActionParameterDefinition {
	return v.value
}

func (v *NullableActionParameterDefinition) Set(val *ActionParameterDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableActionParameterDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableActionParameterDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionParameterDefinition(val *ActionParameterDefinition) *NullableActionParameterDefinition {
	return &NullableActionParameterDefinition{value: val, isSet: true}
}

func (v NullableActionParameterDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionParameterDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

