# DeletedNode

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
**ArchivedByUser** | [**UserInfo**](UserInfo.md) |  | 
**ArchivedAt** | **time.Time** |  | 

## Methods

### NewDeletedNode

`func NewDeletedNode(id string, name string, nodeType string, isFolder bool, isFile bool, modifiedAt time.Time, modifiedByUser UserInfo, createdAt time.Time, createdByUser UserInfo, archivedByUser UserInfo, archivedAt time.Time, ) *DeletedNode`

NewDeletedNode instantiates a new DeletedNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedNodeWithDefaults

`func NewDeletedNodeWithDefaults() *DeletedNode`

NewDeletedNodeWithDefaults instantiates a new DeletedNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeletedNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeletedNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeletedNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeletedNode) SetName(v string)`

SetName sets Name field to given value.


### GetNodeType

`func (o *DeletedNode) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *DeletedNode) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *DeletedNode) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetIsFolder

`func (o *DeletedNode) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *DeletedNode) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *DeletedNode) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.


### GetIsFile

`func (o *DeletedNode) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *DeletedNode) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *DeletedNode) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.


### GetIsLocked

`func (o *DeletedNode) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *DeletedNode) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *DeletedNode) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *DeletedNode) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetModifiedAt

`func (o *DeletedNode) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DeletedNode) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DeletedNode) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedByUser

`func (o *DeletedNode) GetModifiedByUser() UserInfo`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *DeletedNode) GetModifiedByUserOk() (*UserInfo, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *DeletedNode) SetModifiedByUser(v UserInfo)`

SetModifiedByUser sets ModifiedByUser field to given value.


### GetCreatedAt

`func (o *DeletedNode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeletedNode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeletedNode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByUser

`func (o *DeletedNode) GetCreatedByUser() UserInfo`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *DeletedNode) GetCreatedByUserOk() (*UserInfo, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *DeletedNode) SetCreatedByUser(v UserInfo)`

SetCreatedByUser sets CreatedByUser field to given value.


### GetParentId

`func (o *DeletedNode) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DeletedNode) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DeletedNode) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DeletedNode) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetIsLink

`func (o *DeletedNode) GetIsLink() bool`

GetIsLink returns the IsLink field if non-nil, zero value otherwise.

### GetIsLinkOk

`func (o *DeletedNode) GetIsLinkOk() (*bool, bool)`

GetIsLinkOk returns a tuple with the IsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLink

`func (o *DeletedNode) SetIsLink(v bool)`

SetIsLink sets IsLink field to given value.

### HasIsLink

`func (o *DeletedNode) HasIsLink() bool`

HasIsLink returns a boolean if a field has been set.

### GetIsFavorite

`func (o *DeletedNode) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *DeletedNode) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *DeletedNode) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *DeletedNode) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetContent

`func (o *DeletedNode) GetContent() ContentInfo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DeletedNode) GetContentOk() (*ContentInfo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DeletedNode) SetContent(v ContentInfo)`

SetContent sets Content field to given value.

### HasContent

`func (o *DeletedNode) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAspectNames

`func (o *DeletedNode) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *DeletedNode) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *DeletedNode) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *DeletedNode) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *DeletedNode) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeletedNode) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeletedNode) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeletedNode) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAllowableOperations

`func (o *DeletedNode) GetAllowableOperations() []string`

GetAllowableOperations returns the AllowableOperations field if non-nil, zero value otherwise.

### GetAllowableOperationsOk

`func (o *DeletedNode) GetAllowableOperationsOk() (*[]string, bool)`

GetAllowableOperationsOk returns a tuple with the AllowableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableOperations

`func (o *DeletedNode) SetAllowableOperations(v []string)`

SetAllowableOperations sets AllowableOperations field to given value.

### HasAllowableOperations

`func (o *DeletedNode) HasAllowableOperations() bool`

HasAllowableOperations returns a boolean if a field has been set.

### GetPath

`func (o *DeletedNode) GetPath() PathInfo`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeletedNode) GetPathOk() (*PathInfo, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeletedNode) SetPath(v PathInfo)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeletedNode) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermissions

`func (o *DeletedNode) GetPermissions() PermissionsInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DeletedNode) GetPermissionsOk() (*PermissionsInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DeletedNode) SetPermissions(v PermissionsInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DeletedNode) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDefinition

`func (o *DeletedNode) GetDefinition() Definition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DeletedNode) GetDefinitionOk() (*Definition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DeletedNode) SetDefinition(v Definition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DeletedNode) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetArchivedByUser

`func (o *DeletedNode) GetArchivedByUser() UserInfo`

GetArchivedByUser returns the ArchivedByUser field if non-nil, zero value otherwise.

### GetArchivedByUserOk

`func (o *DeletedNode) GetArchivedByUserOk() (*UserInfo, bool)`

GetArchivedByUserOk returns a tuple with the ArchivedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedByUser

`func (o *DeletedNode) SetArchivedByUser(v UserInfo)`

SetArchivedByUser sets ArchivedByUser field to given value.


### GetArchivedAt

`func (o *DeletedNode) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *DeletedNode) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *DeletedNode) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


