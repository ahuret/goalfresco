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

// checks if the Property type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Property{}

// Property struct for Property
type Property struct {
	Id string `json:"id"`
	// the human-readable title
	Title *string `json:"title,omitempty"`
	// the human-readable description
	Description *string `json:"description,omitempty"`
	// the default value
	DefaultValue *string `json:"defaultValue,omitempty"`
	// the name of the property type (e.g. d:text)
	DataType *string `json:"dataType,omitempty"`
	// define if the property is multi-valued
	IsMultiValued *bool `json:"isMultiValued,omitempty"`
	// define if the property is mandatory
	IsMandatory *bool `json:"isMandatory,omitempty"`
	// define if the presence of mandatory properties is enforced
	IsMandatoryEnforced *bool `json:"isMandatoryEnforced,omitempty"`
	// define if the property is system maintained
	IsProtected *bool `json:"isProtected,omitempty"`
	// list of constraints defined for the property
	Constraints []Constraint `json:"constraints,omitempty"`
}

// NewProperty instantiates a new Property object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProperty(id string) *Property {
	this := Property{}
	this.Id = id
	return &this
}

// NewPropertyWithDefaults instantiates a new Property object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyWithDefaults() *Property {
	this := Property{}
	return &this
}

// GetId returns the Id field value
func (o *Property) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Property) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Property) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Property) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Property) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Property) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Property) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Property) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Property) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *Property) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *Property) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *Property) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *Property) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *Property) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *Property) SetDataType(v string) {
	o.DataType = &v
}

// GetIsMultiValued returns the IsMultiValued field value if set, zero value otherwise.
func (o *Property) GetIsMultiValued() bool {
	if o == nil || IsNil(o.IsMultiValued) {
		var ret bool
		return ret
	}
	return *o.IsMultiValued
}

// GetIsMultiValuedOk returns a tuple with the IsMultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetIsMultiValuedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMultiValued) {
		return nil, false
	}
	return o.IsMultiValued, true
}

// HasIsMultiValued returns a boolean if a field has been set.
func (o *Property) HasIsMultiValued() bool {
	if o != nil && !IsNil(o.IsMultiValued) {
		return true
	}

	return false
}

// SetIsMultiValued gets a reference to the given bool and assigns it to the IsMultiValued field.
func (o *Property) SetIsMultiValued(v bool) {
	o.IsMultiValued = &v
}

// GetIsMandatory returns the IsMandatory field value if set, zero value otherwise.
func (o *Property) GetIsMandatory() bool {
	if o == nil || IsNil(o.IsMandatory) {
		var ret bool
		return ret
	}
	return *o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetIsMandatoryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMandatory) {
		return nil, false
	}
	return o.IsMandatory, true
}

// HasIsMandatory returns a boolean if a field has been set.
func (o *Property) HasIsMandatory() bool {
	if o != nil && !IsNil(o.IsMandatory) {
		return true
	}

	return false
}

// SetIsMandatory gets a reference to the given bool and assigns it to the IsMandatory field.
func (o *Property) SetIsMandatory(v bool) {
	o.IsMandatory = &v
}

// GetIsMandatoryEnforced returns the IsMandatoryEnforced field value if set, zero value otherwise.
func (o *Property) GetIsMandatoryEnforced() bool {
	if o == nil || IsNil(o.IsMandatoryEnforced) {
		var ret bool
		return ret
	}
	return *o.IsMandatoryEnforced
}

// GetIsMandatoryEnforcedOk returns a tuple with the IsMandatoryEnforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetIsMandatoryEnforcedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMandatoryEnforced) {
		return nil, false
	}
	return o.IsMandatoryEnforced, true
}

// HasIsMandatoryEnforced returns a boolean if a field has been set.
func (o *Property) HasIsMandatoryEnforced() bool {
	if o != nil && !IsNil(o.IsMandatoryEnforced) {
		return true
	}

	return false
}

// SetIsMandatoryEnforced gets a reference to the given bool and assigns it to the IsMandatoryEnforced field.
func (o *Property) SetIsMandatoryEnforced(v bool) {
	o.IsMandatoryEnforced = &v
}

// GetIsProtected returns the IsProtected field value if set, zero value otherwise.
func (o *Property) GetIsProtected() bool {
	if o == nil || IsNil(o.IsProtected) {
		var ret bool
		return ret
	}
	return *o.IsProtected
}

// GetIsProtectedOk returns a tuple with the IsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetIsProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProtected) {
		return nil, false
	}
	return o.IsProtected, true
}

// HasIsProtected returns a boolean if a field has been set.
func (o *Property) HasIsProtected() bool {
	if o != nil && !IsNil(o.IsProtected) {
		return true
	}

	return false
}

// SetIsProtected gets a reference to the given bool and assigns it to the IsProtected field.
func (o *Property) SetIsProtected(v bool) {
	o.IsProtected = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *Property) GetConstraints() []Constraint {
	if o == nil || IsNil(o.Constraints) {
		var ret []Constraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetConstraintsOk() ([]Constraint, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *Property) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []Constraint and assigns it to the Constraints field.
func (o *Property) SetConstraints(v []Constraint) {
	o.Constraints = v
}

func (o Property) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Property) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.IsMultiValued) {
		toSerialize["isMultiValued"] = o.IsMultiValued
	}
	if !IsNil(o.IsMandatory) {
		toSerialize["isMandatory"] = o.IsMandatory
	}
	if !IsNil(o.IsMandatoryEnforced) {
		toSerialize["isMandatoryEnforced"] = o.IsMandatoryEnforced
	}
	if !IsNil(o.IsProtected) {
		toSerialize["isProtected"] = o.IsProtected
	}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	return toSerialize, nil
}

type NullableProperty struct {
	value *Property
	isSet bool
}

func (v NullableProperty) Get() *Property {
	return v.value
}

func (v *NullableProperty) Set(val *Property) {
	v.value = val
	v.isSet = true
}

func (v NullableProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperty(val *Property) *NullableProperty {
	return &NullableProperty{value: val, isSet: true}
}

func (v NullableProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


