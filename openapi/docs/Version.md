# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VersionComment** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character . must not be used at the end of the name.  | 
**NodeType** | **string** |  | 
**IsFolder** | **bool** |  | 
**IsFile** | **bool** |  | 
**ModifiedAt** | **time.Time** |  | 
**ModifiedByUser** | [**UserInfo**](UserInfo.md) |  | 
**Content** | Pointer to [**ContentInfo**](ContentInfo.md) |  | [optional] 
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVersion

`func NewVersion(id string, name string, nodeType string, isFolder bool, isFile bool, modifiedAt time.Time, modifiedByUser UserInfo, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Version) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Version) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Version) SetId(v string)`

SetId sets Id field to given value.


### GetVersionComment

`func (o *Version) GetVersionComment() string`

GetVersionComment returns the VersionComment field if non-nil, zero value otherwise.

### GetVersionCommentOk

`func (o *Version) GetVersionCommentOk() (*string, bool)`

GetVersionCommentOk returns a tuple with the VersionComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionComment

`func (o *Version) SetVersionComment(v string)`

SetVersionComment sets VersionComment field to given value.

### HasVersionComment

`func (o *Version) HasVersionComment() bool`

HasVersionComment returns a boolean if a field has been set.

### GetName

`func (o *Version) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Version) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Version) SetName(v string)`

SetName sets Name field to given value.


### GetNodeType

`func (o *Version) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *Version) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *Version) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetIsFolder

`func (o *Version) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *Version) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *Version) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.


### GetIsFile

`func (o *Version) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *Version) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *Version) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.


### GetModifiedAt

`func (o *Version) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Version) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Version) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedByUser

`func (o *Version) GetModifiedByUser() UserInfo`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *Version) GetModifiedByUserOk() (*UserInfo, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *Version) SetModifiedByUser(v UserInfo)`

SetModifiedByUser sets ModifiedByUser field to given value.


### GetContent

`func (o *Version) GetContent() ContentInfo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Version) GetContentOk() (*ContentInfo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Version) SetContent(v ContentInfo)`

SetContent sets Content field to given value.

### HasContent

`func (o *Version) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAspectNames

`func (o *Version) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *Version) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *Version) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *Version) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *Version) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Version) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Version) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Version) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


