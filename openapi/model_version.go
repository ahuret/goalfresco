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

// checks if the Version type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Version{}

// Version struct for Version
type Version struct {
	Id string `json:"id"`
	VersionComment *string `json:"versionComment,omitempty"`
	// The name must not contain spaces or the following special characters: * \" < > \\ / ? : and |. The character . must not be used at the end of the name. 
	Name string `json:"name"`
	NodeType string `json:"nodeType"`
	IsFolder bool `json:"isFolder"`
	IsFile bool `json:"isFile"`
	ModifiedAt time.Time `json:"modifiedAt"`
	ModifiedByUser UserInfo `json:"modifiedByUser"`
	Content *ContentInfo `json:"content,omitempty"`
	AspectNames []string `json:"aspectNames,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion(id string, name string, nodeType string, isFolder bool, isFile bool, modifiedAt time.Time, modifiedByUser UserInfo) *Version {
	this := Version{}
	this.Id = id
	this.Name = name
	this.NodeType = nodeType
	this.IsFolder = isFolder
	this.IsFile = isFile
	this.ModifiedAt = modifiedAt
	this.ModifiedByUser = modifiedByUser
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetId returns the Id field value
func (o *Version) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Version) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Version) SetId(v string) {
	o.Id = v
}

// GetVersionComment returns the VersionComment field value if set, zero value otherwise.
func (o *Version) GetVersionComment() string {
	if o == nil || IsNil(o.VersionComment) {
		var ret string
		return ret
	}
	return *o.VersionComment
}

// GetVersionCommentOk returns a tuple with the VersionComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetVersionCommentOk() (*string, bool) {
	if o == nil || IsNil(o.VersionComment) {
		return nil, false
	}
	return o.VersionComment, true
}

// HasVersionComment returns a boolean if a field has been set.
func (o *Version) HasVersionComment() bool {
	if o != nil && !IsNil(o.VersionComment) {
		return true
	}

	return false
}

// SetVersionComment gets a reference to the given string and assigns it to the VersionComment field.
func (o *Version) SetVersionComment(v string) {
	o.VersionComment = &v
}

// GetName returns the Name field value
func (o *Version) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Version) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Version) SetName(v string) {
	o.Name = v
}

// GetNodeType returns the NodeType field value
func (o *Version) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *Version) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *Version) SetNodeType(v string) {
	o.NodeType = v
}

// GetIsFolder returns the IsFolder field value
func (o *Version) GetIsFolder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFolder
}

// GetIsFolderOk returns a tuple with the IsFolder field value
// and a boolean to check if the value has been set.
func (o *Version) GetIsFolderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFolder, true
}

// SetIsFolder sets field value
func (o *Version) SetIsFolder(v bool) {
	o.IsFolder = v
}

// GetIsFile returns the IsFile field value
func (o *Version) GetIsFile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFile
}

// GetIsFileOk returns a tuple with the IsFile field value
// and a boolean to check if the value has been set.
func (o *Version) GetIsFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFile, true
}

// SetIsFile sets field value
func (o *Version) SetIsFile(v bool) {
	o.IsFile = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *Version) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *Version) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *Version) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedByUser returns the ModifiedByUser field value
func (o *Version) GetModifiedByUser() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.ModifiedByUser
}

// GetModifiedByUserOk returns a tuple with the ModifiedByUser field value
// and a boolean to check if the value has been set.
func (o *Version) GetModifiedByUserOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedByUser, true
}

// SetModifiedByUser sets field value
func (o *Version) SetModifiedByUser(v UserInfo) {
	o.ModifiedByUser = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Version) GetContent() ContentInfo {
	if o == nil || IsNil(o.Content) {
		var ret ContentInfo
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetContentOk() (*ContentInfo, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Version) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given ContentInfo and assigns it to the Content field.
func (o *Version) SetContent(v ContentInfo) {
	o.Content = &v
}

// GetAspectNames returns the AspectNames field value if set, zero value otherwise.
func (o *Version) GetAspectNames() []string {
	if o == nil || IsNil(o.AspectNames) {
		var ret []string
		return ret
	}
	return o.AspectNames
}

// GetAspectNamesOk returns a tuple with the AspectNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetAspectNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AspectNames) {
		return nil, false
	}
	return o.AspectNames, true
}

// HasAspectNames returns a boolean if a field has been set.
func (o *Version) HasAspectNames() bool {
	if o != nil && !IsNil(o.AspectNames) {
		return true
	}

	return false
}

// SetAspectNames gets a reference to the given []string and assigns it to the AspectNames field.
func (o *Version) SetAspectNames(v []string) {
	o.AspectNames = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Version) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Version) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *Version) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Version) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.VersionComment) {
		toSerialize["versionComment"] = o.VersionComment
	}
	toSerialize["name"] = o.Name
	toSerialize["nodeType"] = o.NodeType
	toSerialize["isFolder"] = o.IsFolder
	toSerialize["isFile"] = o.IsFile
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedByUser"] = o.ModifiedByUser
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.AspectNames) {
		toSerialize["aspectNames"] = o.AspectNames
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


