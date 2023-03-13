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

// checks if the NodeBodyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeBodyUpdate{}

// NodeBodyUpdate struct for NodeBodyUpdate
type NodeBodyUpdate struct {
	// The name must not contain spaces or the following special characters: * \" < > \\ / ? : and |. The character . must not be used at the end of the name. 
	Name *string `json:"name,omitempty"`
	NodeType *string `json:"nodeType,omitempty"`
	AspectNames []string `json:"aspectNames,omitempty"`
	Properties *map[string]string `json:"properties,omitempty"`
	Permissions *PermissionsBody `json:"permissions,omitempty"`
}

// NewNodeBodyUpdate instantiates a new NodeBodyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeBodyUpdate() *NodeBodyUpdate {
	this := NodeBodyUpdate{}
	return &this
}

// NewNodeBodyUpdateWithDefaults instantiates a new NodeBodyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeBodyUpdateWithDefaults() *NodeBodyUpdate {
	this := NodeBodyUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NodeBodyUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBodyUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NodeBodyUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NodeBodyUpdate) SetName(v string) {
	o.Name = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *NodeBodyUpdate) GetNodeType() string {
	if o == nil || IsNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBodyUpdate) GetNodeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NodeType) {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *NodeBodyUpdate) HasNodeType() bool {
	if o != nil && !IsNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *NodeBodyUpdate) SetNodeType(v string) {
	o.NodeType = &v
}

// GetAspectNames returns the AspectNames field value if set, zero value otherwise.
func (o *NodeBodyUpdate) GetAspectNames() []string {
	if o == nil || IsNil(o.AspectNames) {
		var ret []string
		return ret
	}
	return o.AspectNames
}

// GetAspectNamesOk returns a tuple with the AspectNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBodyUpdate) GetAspectNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AspectNames) {
		return nil, false
	}
	return o.AspectNames, true
}

// HasAspectNames returns a boolean if a field has been set.
func (o *NodeBodyUpdate) HasAspectNames() bool {
	if o != nil && !IsNil(o.AspectNames) {
		return true
	}

	return false
}

// SetAspectNames gets a reference to the given []string and assigns it to the AspectNames field.
func (o *NodeBodyUpdate) SetAspectNames(v []string) {
	o.AspectNames = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *NodeBodyUpdate) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBodyUpdate) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *NodeBodyUpdate) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *NodeBodyUpdate) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *NodeBodyUpdate) GetPermissions() PermissionsBody {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsBody
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBodyUpdate) GetPermissionsOk() (*PermissionsBody, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *NodeBodyUpdate) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsBody and assigns it to the Permissions field.
func (o *NodeBodyUpdate) SetPermissions(v PermissionsBody) {
	o.Permissions = &v
}

func (o NodeBodyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeBodyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NodeType) {
		toSerialize["nodeType"] = o.NodeType
	}
	if !IsNil(o.AspectNames) {
		toSerialize["aspectNames"] = o.AspectNames
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableNodeBodyUpdate struct {
	value *NodeBodyUpdate
	isSet bool
}

func (v NullableNodeBodyUpdate) Get() *NodeBodyUpdate {
	return v.value
}

func (v *NullableNodeBodyUpdate) Set(val *NodeBodyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeBodyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeBodyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeBodyUpdate(val *NodeBodyUpdate) *NullableNodeBodyUpdate {
	return &NullableNodeBodyUpdate{value: val, isSet: true}
}

func (v NullableNodeBodyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeBodyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


