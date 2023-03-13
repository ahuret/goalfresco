# NodeBodyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character . must not be used at the end of the name.  | 
**NodeType** | **string** |  | 
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Permissions** | Pointer to [**PermissionsBody**](PermissionsBody.md) |  | [optional] 
**Definition** | Pointer to [**Definition**](Definition.md) |  | [optional] 
**RelativePath** | Pointer to **string** |  | [optional] 
**Association** | Pointer to [**NodeBodyCreateAssociation**](NodeBodyCreateAssociation.md) |  | [optional] 
**SecondaryChildren** | Pointer to [**[]ChildAssociationBody**](ChildAssociationBody.md) |  | [optional] 
**Targets** | Pointer to [**[]AssociationBody**](AssociationBody.md) |  | [optional] 

## Methods

### NewNodeBodyCreate

`func NewNodeBodyCreate(name string, nodeType string, ) *NodeBodyCreate`

NewNodeBodyCreate instantiates a new NodeBodyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBodyCreateWithDefaults

`func NewNodeBodyCreateWithDefaults() *NodeBodyCreate`

NewNodeBodyCreateWithDefaults instantiates a new NodeBodyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeBodyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeBodyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeBodyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetNodeType

`func (o *NodeBodyCreate) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *NodeBodyCreate) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *NodeBodyCreate) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetAspectNames

`func (o *NodeBodyCreate) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *NodeBodyCreate) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *NodeBodyCreate) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *NodeBodyCreate) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *NodeBodyCreate) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodeBodyCreate) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodeBodyCreate) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *NodeBodyCreate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPermissions

`func (o *NodeBodyCreate) GetPermissions() PermissionsBody`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NodeBodyCreate) GetPermissionsOk() (*PermissionsBody, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NodeBodyCreate) SetPermissions(v PermissionsBody)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *NodeBodyCreate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDefinition

`func (o *NodeBodyCreate) GetDefinition() Definition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NodeBodyCreate) GetDefinitionOk() (*Definition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NodeBodyCreate) SetDefinition(v Definition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *NodeBodyCreate) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetRelativePath

`func (o *NodeBodyCreate) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *NodeBodyCreate) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *NodeBodyCreate) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *NodeBodyCreate) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetAssociation

`func (o *NodeBodyCreate) GetAssociation() NodeBodyCreateAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *NodeBodyCreate) GetAssociationOk() (*NodeBodyCreateAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *NodeBodyCreate) SetAssociation(v NodeBodyCreateAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *NodeBodyCreate) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetSecondaryChildren

`func (o *NodeBodyCreate) GetSecondaryChildren() []ChildAssociationBody`

GetSecondaryChildren returns the SecondaryChildren field if non-nil, zero value otherwise.

### GetSecondaryChildrenOk

`func (o *NodeBodyCreate) GetSecondaryChildrenOk() (*[]ChildAssociationBody, bool)`

GetSecondaryChildrenOk returns a tuple with the SecondaryChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryChildren

`func (o *NodeBodyCreate) SetSecondaryChildren(v []ChildAssociationBody)`

SetSecondaryChildren sets SecondaryChildren field to given value.

### HasSecondaryChildren

`func (o *NodeBodyCreate) HasSecondaryChildren() bool`

HasSecondaryChildren returns a boolean if a field has been set.

### GetTargets

`func (o *NodeBodyCreate) GetTargets() []AssociationBody`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *NodeBodyCreate) GetTargetsOk() (*[]AssociationBody, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *NodeBodyCreate) SetTargets(v []AssociationBody)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *NodeBodyCreate) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


