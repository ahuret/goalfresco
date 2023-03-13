# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character . must not be used at the end of the name.  | 
**NodeType** | **string** |  | 
**IsFolder** | **bool** |  | 
**IsFile** | **bool** |  | 
**IsLocked** | Pointer to **bool** |  | [optional] [default to false]
**ModifiedAt** | **time.Time** |  | 
**ModifiedByUser** | [**UserInfo**](UserInfo.md) |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedByUser** | [**UserInfo**](UserInfo.md) |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**IsLink** | Pointer to **bool** |  | [optional] 
**IsFavorite** | Pointer to **bool** |  | [optional] 
**Content** | Pointer to [**ContentInfo**](ContentInfo.md) |  | [optional] 
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**AllowableOperations** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to [**PathInfo**](PathInfo.md) |  | [optional] 
**Permissions** | Pointer to [**PermissionsInfo**](PermissionsInfo.md) |  | [optional] 
**Definition** | Pointer to [**Definition**](Definition.md) |  | [optional] 

## Methods

### NewNode

`func NewNode(id string, name string, nodeType string, isFolder bool, isFile bool, modifiedAt time.Time, modifiedByUser UserInfo, createdAt time.Time, createdByUser UserInfo, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.


### GetNodeType

`func (o *Node) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *Node) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *Node) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetIsFolder

`func (o *Node) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *Node) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *Node) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.


### GetIsFile

`func (o *Node) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *Node) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *Node) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.


### GetIsLocked

`func (o *Node) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Node) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Node) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *Node) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Node) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Node) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Node) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedByUser

`func (o *Node) GetModifiedByUser() UserInfo`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *Node) GetModifiedByUserOk() (*UserInfo, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *Node) SetModifiedByUser(v UserInfo)`

SetModifiedByUser sets ModifiedByUser field to given value.


### GetCreatedAt

`func (o *Node) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Node) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Node) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByUser

`func (o *Node) GetCreatedByUser() UserInfo`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *Node) GetCreatedByUserOk() (*UserInfo, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *Node) SetCreatedByUser(v UserInfo)`

SetCreatedByUser sets CreatedByUser field to given value.


### GetParentId

`func (o *Node) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Node) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Node) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Node) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetIsLink

`func (o *Node) GetIsLink() bool`

GetIsLink returns the IsLink field if non-nil, zero value otherwise.

### GetIsLinkOk

`func (o *Node) GetIsLinkOk() (*bool, bool)`

GetIsLinkOk returns a tuple with the IsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLink

`func (o *Node) SetIsLink(v bool)`

SetIsLink sets IsLink field to given value.

### HasIsLink

`func (o *Node) HasIsLink() bool`

HasIsLink returns a boolean if a field has been set.

### GetIsFavorite

`func (o *Node) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *Node) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *Node) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *Node) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetContent

`func (o *Node) GetContent() ContentInfo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Node) GetContentOk() (*ContentInfo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Node) SetContent(v ContentInfo)`

SetContent sets Content field to given value.

### HasContent

`func (o *Node) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAspectNames

`func (o *Node) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *Node) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *Node) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *Node) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *Node) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Node) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Node) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Node) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAllowableOperations

`func (o *Node) GetAllowableOperations() []string`

GetAllowableOperations returns the AllowableOperations field if non-nil, zero value otherwise.

### GetAllowableOperationsOk

`func (o *Node) GetAllowableOperationsOk() (*[]string, bool)`

GetAllowableOperationsOk returns a tuple with the AllowableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableOperations

`func (o *Node) SetAllowableOperations(v []string)`

SetAllowableOperations sets AllowableOperations field to given value.

### HasAllowableOperations

`func (o *Node) HasAllowableOperations() bool`

HasAllowableOperations returns a boolean if a field has been set.

### GetPath

`func (o *Node) GetPath() PathInfo`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Node) GetPathOk() (*PathInfo, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Node) SetPath(v PathInfo)`

SetPath sets Path field to given value.

### HasPath

`func (o *Node) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermissions

`func (o *Node) GetPermissions() PermissionsInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Node) GetPermissionsOk() (*PermissionsInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Node) SetPermissions(v PermissionsInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Node) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDefinition

`func (o *Node) GetDefinition() Definition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *Node) GetDefinitionOk() (*Definition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *Node) SetDefinition(v Definition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *Node) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


