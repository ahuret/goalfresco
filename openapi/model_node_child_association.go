/*
Alfresco Content Services REST API

**Core API**  Provides access to the core features of Alfresco Content Services. 

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"github.com/ahuret/goalfresco/time"
)

// checks if the NodeChildAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeChildAssociation{}

// NodeChildAssociation struct for NodeChildAssociation
type NodeChildAssociation struct {
	Id string `json:"id"`
	// The name must not contain spaces or the following special characters: * \" < > \\ / ? : and |. The character . must not be used at the end of the name. 
	Name string `json:"name"`
	NodeType string `json:"nodeType"`
	IsFolder bool `json:"isFolder"`
	IsFile bool `json:"isFile"`
	IsLocked *bool `json:"isLocked,omitempty"`
	ModifiedAt time.Time `json:"modifiedAt"`
	ModifiedByUser UserInfo `json:"modifiedByUser"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedByUser UserInfo `json:"createdByUser"`
	ParentId *string `json:"parentId,omitempty"`
	IsLink *bool `json:"isLink,omitempty"`
	IsFavorite *bool `json:"isFavorite,omitempty"`
	Content *ContentInfo `json:"content,omitempty"`
	AspectNames []string `json:"aspectNames,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	AllowableOperations []string `json:"allowableOperations,omitempty"`
	Path *PathInfo `json:"path,omitempty"`
	Permissions *PermissionsInfo `json:"permissions,omitempty"`
	Definition *Definition `json:"definition,omitempty"`
	Association *ChildAssociationInfo `json:"association,omitempty"`
}

// NewNodeChildAssociation instantiates a new NodeChildAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeChildAssociation(id string, name string, nodeType string, isFolder bool, isFile bool, modifiedAt time.Time, modifiedByUser UserInfo, createdAt time.Time, createdByUser UserInfo) *NodeChildAssociation {
	this := NodeChildAssociation{}
	this.Id = id
	this.Name = name
	this.NodeType = nodeType
	this.IsFolder = isFolder
	this.IsFile = isFile
	var isLocked bool = false
	this.IsLocked = &isLocked
	this.ModifiedAt = modifiedAt
	this.ModifiedByUser = modifiedByUser
	this.CreatedAt = createdAt
	this.CreatedByUser = createdByUser
	return &this
}

// NewNodeChildAssociationWithDefaults instantiates a new NodeChildAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeChildAssociationWithDefaults() *NodeChildAssociation {
	this := NodeChildAssociation{}
	var isLocked bool = false
	this.IsLocked = &isLocked
	return &this
}

// GetId returns the Id field value
func (o *NodeChildAssociation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NodeChildAssociation) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *NodeChildAssociation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodeChildAssociation) SetName(v string) {
	o.Name = v
}

// GetNodeType returns the NodeType field value
func (o *NodeChildAssociation) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *NodeChildAssociation) SetNodeType(v string) {
	o.NodeType = v
}

// GetIsFolder returns the IsFolder field value
func (o *NodeChildAssociation) GetIsFolder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFolder
}

// GetIsFolderOk returns a tuple with the IsFolder field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetIsFolderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFolder, true
}

// SetIsFolder sets field value
func (o *NodeChildAssociation) SetIsFolder(v bool) {
	o.IsFolder = v
}

// GetIsFile returns the IsFile field value
func (o *NodeChildAssociation) GetIsFile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFile
}

// GetIsFileOk returns a tuple with the IsFile field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetIsFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFile, true
}

// SetIsFile sets field value
func (o *NodeChildAssociation) SetIsFile(v bool) {
	o.IsFile = v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *NodeChildAssociation) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *NodeChildAssociation) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *NodeChildAssociation) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedByUser returns the ModifiedByUser field value
func (o *NodeChildAssociation) GetModifiedByUser() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.ModifiedByUser
}

// GetModifiedByUserOk returns a tuple with the ModifiedByUser field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetModifiedByUserOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedByUser, true
}

// SetModifiedByUser sets field value
func (o *NodeChildAssociation) SetModifiedByUser(v UserInfo) {
	o.ModifiedByUser = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NodeChildAssociation) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NodeChildAssociation) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedByUser returns the CreatedByUser field value
func (o *NodeChildAssociation) GetCreatedByUser() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetCreatedByUserOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// SetCreatedByUser sets field value
func (o *NodeChildAssociation) SetCreatedByUser(v UserInfo) {
	o.CreatedByUser = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *NodeChildAssociation) SetParentId(v string) {
	o.ParentId = &v
}

// GetIsLink returns the IsLink field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetIsLink() bool {
	if o == nil || IsNil(o.IsLink) {
		var ret bool
		return ret
	}
	return *o.IsLink
}

// GetIsLinkOk returns a tuple with the IsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetIsLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLink) {
		return nil, false
	}
	return o.IsLink, true
}

// HasIsLink returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasIsLink() bool {
	if o != nil && !IsNil(o.IsLink) {
		return true
	}

	return false
}

// SetIsLink gets a reference to the given bool and assigns it to the IsLink field.
func (o *NodeChildAssociation) SetIsLink(v bool) {
	o.IsLink = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasIsFavorite() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *NodeChildAssociation) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetContent() ContentInfo {
	if o == nil || IsNil(o.Content) {
		var ret ContentInfo
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetContentOk() (*ContentInfo, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given ContentInfo and assigns it to the Content field.
func (o *NodeChildAssociation) SetContent(v ContentInfo) {
	o.Content = &v
}

// GetAspectNames returns the AspectNames field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetAspectNames() []string {
	if o == nil || IsNil(o.AspectNames) {
		var ret []string
		return ret
	}
	return o.AspectNames
}

// GetAspectNamesOk returns a tuple with the AspectNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetAspectNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AspectNames) {
		return nil, false
	}
	return o.AspectNames, true
}

// HasAspectNames returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasAspectNames() bool {
	if o != nil && !IsNil(o.AspectNames) {
		return true
	}

	return false
}

// SetAspectNames gets a reference to the given []string and assigns it to the AspectNames field.
func (o *NodeChildAssociation) SetAspectNames(v []string) {
	o.AspectNames = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *NodeChildAssociation) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetAllowableOperations returns the AllowableOperations field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetAllowableOperations() []string {
	if o == nil || IsNil(o.AllowableOperations) {
		var ret []string
		return ret
	}
	return o.AllowableOperations
}

// GetAllowableOperationsOk returns a tuple with the AllowableOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetAllowableOperationsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowableOperations) {
		return nil, false
	}
	return o.AllowableOperations, true
}

// HasAllowableOperations returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasAllowableOperations() bool {
	if o != nil && !IsNil(o.AllowableOperations) {
		return true
	}

	return false
}

// SetAllowableOperations gets a reference to the given []string and assigns it to the AllowableOperations field.
func (o *NodeChildAssociation) SetAllowableOperations(v []string) {
	o.AllowableOperations = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetPath() PathInfo {
	if o == nil || IsNil(o.Path) {
		var ret PathInfo
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetPathOk() (*PathInfo, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given PathInfo and assigns it to the Path field.
func (o *NodeChildAssociation) SetPath(v PathInfo) {
	o.Path = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetPermissions() PermissionsInfo {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsInfo
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetPermissionsOk() (*PermissionsInfo, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsInfo and assigns it to the Permissions field.
func (o *NodeChildAssociation) SetPermissions(v PermissionsInfo) {
	o.Permissions = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetDefinition() Definition {
	if o == nil || IsNil(o.Definition) {
		var ret Definition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetDefinitionOk() (*Definition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given Definition and assigns it to the Definition field.
func (o *NodeChildAssociation) SetDefinition(v Definition) {
	o.Definition = &v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *NodeChildAssociation) GetAssociation() ChildAssociationInfo {
	if o == nil || IsNil(o.Association) {
		var ret ChildAssociationInfo
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeChildAssociation) GetAssociationOk() (*ChildAssociationInfo, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *NodeChildAssociation) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given ChildAssociationInfo and assigns it to the Association field.
func (o *NodeChildAssociation) SetAssociation(v ChildAssociationInfo) {
	o.Association = &v
}

func (o NodeChildAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeChildAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["nodeType"] = o.NodeType
	toSerialize["isFolder"] = o.IsFolder
	toSerialize["isFile"] = o.IsFile
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedByUser"] = o.ModifiedByUser
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdByUser"] = o.CreatedByUser
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.IsLink) {
		toSerialize["isLink"] = o.IsLink
	}
	if !IsNil(o.IsFavorite) {
		toSerialize["isFavorite"] = o.IsFavorite
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.AspectNames) {
		toSerialize["aspectNames"] = o.AspectNames
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.AllowableOperations) {
		toSerialize["allowableOperations"] = o.AllowableOperations
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	return toSerialize, nil
}

type NullableNodeChildAssociation struct {
	value *NodeChildAssociation
	isSet bool
}

func (v NullableNodeChildAssociation) Get() *NodeChildAssociation {
	return v.value
}

func (v *NullableNodeChildAssociation) Set(val *NodeChildAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeChildAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeChildAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeChildAssociation(val *NodeChildAssociation) *NullableNodeChildAssociation {
	return &NullableNodeChildAssociation{value: val, isSet: true}
}

func (v NullableNodeChildAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeChildAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


