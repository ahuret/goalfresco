# NodeChildAssociation

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
**Association** | Pointer to [**ChildAssociationInfo**](ChildAssociationInfo.md) |  | [optional] 

## Methods

### NewNodeChildAssociation

`func NewNodeChildAssociation(id string, name string, nodeType string, isFolder bool, isFile bool, modifiedAt time.Time, modifiedByUser UserInfo, createdAt time.Time, createdByUser UserInfo, ) *NodeChildAssociation`

NewNodeChildAssociation instantiates a new NodeChildAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeChildAssociationWithDefaults

`func NewNodeChildAssociationWithDefaults() *NodeChildAssociation`

NewNodeChildAssociationWithDefaults instantiates a new NodeChildAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeChildAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeChildAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeChildAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NodeChildAssociation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeChildAssociation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeChildAssociation) SetName(v string)`

SetName sets Name field to given value.


### GetNodeType

`func (o *NodeChildAssociation) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *NodeChildAssociation) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *NodeChildAssociation) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetIsFolder

`func (o *NodeChildAssociation) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *NodeChildAssociation) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *NodeChildAssociation) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.


### GetIsFile

`func (o *NodeChildAssociation) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *NodeChildAssociation) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *NodeChildAssociation) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.


### GetIsLocked

`func (o *NodeChildAssociation) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *NodeChildAssociation) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *NodeChildAssociation) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *NodeChildAssociation) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetModifiedAt

`func (o *NodeChildAssociation) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NodeChildAssociation) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NodeChildAssociation) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedByUser

`func (o *NodeChildAssociation) GetModifiedByUser() UserInfo`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *NodeChildAssociation) GetModifiedByUserOk() (*UserInfo, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *NodeChildAssociation) SetModifiedByUser(v UserInfo)`

SetModifiedByUser sets ModifiedByUser field to given value.


### GetCreatedAt

`func (o *NodeChildAssociation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodeChildAssociation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodeChildAssociation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByUser

`func (o *NodeChildAssociation) GetCreatedByUser() UserInfo`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *NodeChildAssociation) GetCreatedByUserOk() (*UserInfo, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *NodeChildAssociation) SetCreatedByUser(v UserInfo)`

SetCreatedByUser sets CreatedByUser field to given value.


### GetParentId

`func (o *NodeChildAssociation) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *NodeChildAssociation) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *NodeChildAssociation) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *NodeChildAssociation) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetIsLink

`func (o *NodeChildAssociation) GetIsLink() bool`

GetIsLink returns the IsLink field if non-nil, zero value otherwise.

### GetIsLinkOk

`func (o *NodeChildAssociation) GetIsLinkOk() (*bool, bool)`

GetIsLinkOk returns a tuple with the IsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLink

`func (o *NodeChildAssociation) SetIsLink(v bool)`

SetIsLink sets IsLink field to given value.

### HasIsLink

`func (o *NodeChildAssociation) HasIsLink() bool`

HasIsLink returns a boolean if a field has been set.

### GetIsFavorite

`func (o *NodeChildAssociation) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *NodeChildAssociation) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *NodeChildAssociation) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *NodeChildAssociation) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetContent

`func (o *NodeChildAssociation) GetContent() ContentInfo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NodeChildAssociation) GetContentOk() (*ContentInfo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NodeChildAssociation) SetContent(v ContentInfo)`

SetContent sets Content field to given value.

### HasContent

`func (o *NodeChildAssociation) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAspectNames

`func (o *NodeChildAssociation) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *NodeChildAssociation) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *NodeChildAssociation) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *NodeChildAssociation) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *NodeChildAssociation) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodeChildAssociation) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodeChildAssociation) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *NodeChildAssociation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAllowableOperations

`func (o *NodeChildAssociation) GetAllowableOperations() []string`

GetAllowableOperations returns the AllowableOperations field if non-nil, zero value otherwise.

### GetAllowableOperationsOk

`func (o *NodeChildAssociation) GetAllowableOperationsOk() (*[]string, bool)`

GetAllowableOperationsOk returns a tuple with the AllowableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableOperations

`func (o *NodeChildAssociation) SetAllowableOperations(v []string)`

SetAllowableOperations sets AllowableOperations field to given value.

### HasAllowableOperations

`func (o *NodeChildAssociation) HasAllowableOperations() bool`

HasAllowableOperations returns a boolean if a field has been set.

### GetPath

`func (o *NodeChildAssociation) GetPath() PathInfo`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NodeChildAssociation) GetPathOk() (*PathInfo, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NodeChildAssociation) SetPath(v PathInfo)`

SetPath sets Path field to given value.

### HasPath

`func (o *NodeChildAssociation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPermissions

`func (o *NodeChildAssociation) GetPermissions() PermissionsInfo`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NodeChildAssociation) GetPermissionsOk() (*PermissionsInfo, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NodeChildAssociation) SetPermissions(v PermissionsInfo)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *NodeChildAssociation) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDefinition

`func (o *NodeChildAssociation) GetDefinition() Definition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NodeChildAssociation) GetDefinitionOk() (*Definition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NodeChildAssociation) SetDefinition(v Definition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *NodeChildAssociation) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetAssociation

`func (o *NodeChildAssociation) GetAssociation() ChildAssociationInfo`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *NodeChildAssociation) GetAssociationOk() (*ChildAssociationInfo, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *NodeChildAssociation) SetAssociation(v ChildAssociationInfo)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *NodeChildAssociation) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


