# NodeBodyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character . must not be used at the end of the name.  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Permissions** | Pointer to [**PermissionsBody**](PermissionsBody.md) |  | [optional] 

## Methods

### NewNodeBodyUpdate

`func NewNodeBodyUpdate() *NodeBodyUpdate`

NewNodeBodyUpdate instantiates a new NodeBodyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBodyUpdateWithDefaults

`func NewNodeBodyUpdateWithDefaults() *NodeBodyUpdate`

NewNodeBodyUpdateWithDefaults instantiates a new NodeBodyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeBodyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeBodyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeBodyUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeBodyUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeType

`func (o *NodeBodyUpdate) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *NodeBodyUpdate) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *NodeBodyUpdate) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *NodeBodyUpdate) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetAspectNames

`func (o *NodeBodyUpdate) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *NodeBodyUpdate) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *NodeBodyUpdate) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *NodeBodyUpdate) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *NodeBodyUpdate) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodeBodyUpdate) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodeBodyUpdate) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *NodeBodyUpdate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPermissions

`func (o *NodeBodyUpdate) GetPermissions() PermissionsBody`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NodeBodyUpdate) GetPermissionsOk() (*PermissionsBody, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NodeBodyUpdate) SetPermissions(v PermissionsBody)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *NodeBodyUpdate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


