# NodeAssociation

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
**Association** | Pointer to [**AssociationInfo**](AssociationInfo.md) |  | [optional] 

## Methods

### NewNodeAssociation

`func NewNodeAssociation(id string, name string, nodeType string, isFolder bool, isFile bool, modifiedAt time.Time, modifiedByUser UserInfo, createdAt time.Time, createdByUser UserInfo, ) *NodeAssociation`

NewNodeAssociation instantiates a new NodeAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAssociationWithDefaults

`func NewNodeAssociationWithDefaults() *NodeAssociation`

NewNodeAssociationWithDefaults instantiates a new NodeAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NodeAssociation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeAssociation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeAssociation) SetName(v string)`

SetName sets Name field to given value.


### GetNodeType

`func (o *NodeAssociation) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *NodeAssociation) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *NodeAssociation) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetIsFolder

`func (o *NodeAssociation) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *NodeAssociation) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *NodeAssociation) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.


### GetIsFile

`func (o *NodeAssociation) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *NodeAssociation) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *NodeAssociation) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.


### GetIsLocked

`func (o *NodeAssociation) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *NodeAssociation) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *NodeAssociation) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *NodeAssociation) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetModifiedAt

`func (o *NodeAssociation) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NodeAssociation) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NodeAssociation) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedByUser

`func (o *NodeAssociation) GetModifiedByUser() UserInfo`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *NodeAssociation) GetModifiedByUserOk() (*UserInfo, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *NodeAssociation) SetModifiedByUser(v UserInfo)`

SetModifiedByUser sets ModifiedByUser field to given value.


### GetCreatedAt

`func (o *NodeAssociation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodeAssociation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodeAssociation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByUser

`func (o *NodeAssociation) GetCreatedByUser() UserInfo`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *NodeAssociation) GetCreatedByUserOk() (*UserInfo, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *NodeAssociation) SetCreatedByUser(v UserInfo)`

SetCreatedByUser sets CreatedByUser field to given value.


### GetParentId

`func (o *NodeAssociation) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *NodeAssociation) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *NodeAssociation) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *NodeAssociation) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetIsLink

`func (o *NodeAssociation) GetIsLink() bool`

GetIsLink returns the IsLink field if non-nil, zero value otherwise.

### GetIsLinkOk

`func (o *NodeAssociation) GetIsLinkOk() (*bool, bool)`

GetIsLinkOk returns a tuple with the IsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLink

`func (o *NodeAssociation) SetIsLink(v bool)`

SetIsLink sets IsLink field to given value.

### HasIsLink

`func (o *NodeAssociation) HasIsLink() bool`

HasIsLink returns a boolean if a field has been set.

### GetIsFavorite

`func (o *NodeAssociation) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *NodeAssociation) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *NodeAssociation) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *NodeAssociation) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetContent

`func (o *NodeAssociation) GetContent() ContentInfo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NodeAssociation) GetContentOk() (*ContentInfo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NodeAssociation) SetContent(v ContentInfo)`

SetContent sets Content field to given value.

### HasContent

`func (o *NodeAssociation) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAspectNames

`func (o *NodeAssociation) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *NodeAssociation) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *NodeAssociation) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *NodeAssociation) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *NodeAssociation) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodeAssociation) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodeAssociation) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *NodeAssociation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAllowableOperations

`func (o *NodeAssociation) GetAllowableOperations() []string`

GetAllowableOperations returns the AllowableOperations field if non-nil, zero value otherwise.

### GetAllowableOperationsOk

`func (o *NodeAssociation) GetAllowableOperationsOk() (*[]string, bool)`

GetAllowableOperationsOk returns a tuple with the AllowableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableOperations

`func (o *NodeAssociation) SetAllowableOperations(v []string)`

SetAllowableOperations sets AllowableOperations field to given value.

### HasAllowableOperations

`func (o *NodeAssociation) HasAllowableOperations() bool`

HasAllowableOperations returns a boolean if a field has been set.

### GetPath

`func (o *NodeAssociation) GetPath() PathInfo`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NodeAssociation) GetPathOk() (*PathInfo, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NodeAssociation) SetPath(v PathInfo)`

SetPath sets Path field to given value.

### HasPath

`func (o *NodeAssociation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermissions

`func (o *NodeAssociation) GetPermissions() PermissionsInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NodeAssociation) GetPermissionsOk() (*PermissionsInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NodeAssociation) SetPermissions(v PermissionsInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *NodeAssociation) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDefinition

`func (o *NodeAssociation) GetDefinition() Definition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NodeAssociation) GetDefinitionOk() (*Definition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NodeAssociation) SetDefinition(v Definition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *NodeAssociation) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetAssociation

`func (o *NodeAssociation) GetAssociation() AssociationInfo`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *NodeAssociation) GetAssociationOk() (*AssociationInfo, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *NodeAssociation) SetAssociation(v AssociationInfo)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *NodeAssociation) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


