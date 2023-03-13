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

// checks if the PermissionsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsInfo{}

// PermissionsInfo struct for PermissionsInfo
type PermissionsInfo struct {
	IsInheritanceEnabled *bool `json:"isInheritanceEnabled,omitempty"`
	Inherited []PermissionElement `json:"inherited,omitempty"`
	LocallySet []PermissionElement `json:"locallySet,omitempty"`
	Settable []string `json:"settable,omitempty"`
}

// NewPermissionsInfo instantiates a new PermissionsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsInfo() *PermissionsInfo {
	this := PermissionsInfo{}
	return &this
}

// NewPermissionsInfoWithDefaults instantiates a new PermissionsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsInfoWithDefaults() *PermissionsInfo {
	this := PermissionsInfo{}
	return &this
}

// GetIsInheritanceEnabled returns the IsInheritanceEnabled field value if set, zero value otherwise.
func (o *PermissionsInfo) GetIsInheritanceEnabled() bool {
	if o == nil || IsNil(o.IsInheritanceEnabled) {
		var ret bool
		return ret
	}
	return *o.IsInheritanceEnabled
}

// GetIsInheritanceEnabledOk returns a tuple with the IsInheritanceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsInfo) GetIsInheritanceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInheritanceEnabled) {
		return nil, false
	}
	return o.IsInheritanceEnabled, true
}

// HasIsInheritanceEnabled returns a boolean if a field has been set.
func (o *PermissionsInfo) HasIsInheritanceEnabled() bool {
	if o != nil && !IsNil(o.IsInheritanceEnabled) {
		return true
	}

	return false
}

// SetIsInheritanceEnabled gets a reference to the given bool and assigns it to the IsInheritanceEnabled field.
func (o *PermissionsInfo) SetIsInheritanceEnabled(v bool) {
	o.IsInheritanceEnabled = &v
}

// GetInherited returns the Inherited field value if set, zero value otherwise.
func (o *PermissionsInfo) GetInherited() []PermissionElement {
	if o == nil || IsNil(o.Inherited) {
		var ret []PermissionElement
		return ret
	}
	return o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsInfo) GetInheritedOk() ([]PermissionElement, bool) {
	if o == nil || IsNil(o.Inherited) {
		return nil, false
	}
	return o.Inherited, true
}

// HasInherited returns a boolean if a field has been set.
func (o *PermissionsInfo) HasInherited() bool {
	if o != nil && !IsNil(o.Inherited) {
		return true
	}

	return false
}

// SetInherited gets a reference to the given []PermissionElement and assigns it to the Inherited field.
func (o *PermissionsInfo) SetInherited(v []PermissionElement) {
	o.Inherited = v
}

// GetLocallySet returns the LocallySet field value if set, zero value otherwise.
func (o *PermissionsInfo) GetLocallySet() []PermissionElement {
	if o == nil || IsNil(o.LocallySet) {
		var ret []PermissionElement
		return ret
	}
	return o.LocallySet
}

// GetLocallySetOk returns a tuple with the LocallySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsInfo) GetLocallySetOk() ([]PermissionElement, bool) {
	if o == nil || IsNil(o.LocallySet) {
		return nil, false
	}
	return o.LocallySet, true
}

// HasLocallySet returns a boolean if a field has been set.
func (o *PermissionsInfo) HasLocallySet() bool {
	if o != nil && !IsNil(o.LocallySet) {
		return true
	}

	return false
}

// SetLocallySet gets a reference to the given []PermissionElement and assigns it to the LocallySet field.
func (o *PermissionsInfo) SetLocallySet(v []PermissionElement) {
	o.LocallySet = v
}

// GetSettable returns the Settable field value if set, zero value otherwise.
func (o *PermissionsInfo) GetSettable() []string {
	if o == nil || IsNil(o.Settable) {
		var ret []string
		return ret
	}
	return o.Settable
}

// GetSettableOk returns a tuple with the Settable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsInfo) GetSettableOk() ([]string, bool) {
	if o == nil || IsNil(o.Settable) {
		return nil, false
	}
	return o.Settable, true
}

// HasSettable returns a boolean if a field has been set.
func (o *PermissionsInfo) HasSettable() bool {
	if o != nil && !IsNil(o.Settable) {
		return true
	}

	return false
}

// SetSettable gets a reference to the given []string and assigns it to the Settable field.
func (o *PermissionsInfo) SetSettable(v []string) {
	o.Settable = v
}

func (o PermissionsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsInheritanceEnabled) {
		toSerialize["isInheritanceEnabled"] = o.IsInheritanceEnabled
	}
	if !IsNil(o.Inherited) {
		toSerialize["inherited"] = o.Inherited
	}
	if !IsNil(o.LocallySet) {
		toSerialize["locallySet"] = o.LocallySet
	}
	if !IsNil(o.Settable) {
		toSerialize["settable"] = o.Settable
	}
	return toSerialize, nil
}

type NullablePermissionsInfo struct {
	value *PermissionsInfo
	isSet bool
}

func (v NullablePermissionsInfo) Get() *PermissionsInfo {
	return v.value
}

func (v *NullablePermissionsInfo) Set(val *PermissionsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsInfo(val *PermissionsInfo) *NullablePermissionsInfo {
	return &NullablePermissionsInfo{value: val, isSet: true}
}

func (v NullablePermissionsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

