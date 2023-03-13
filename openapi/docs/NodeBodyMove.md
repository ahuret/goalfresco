# NodeBodyMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetParentId** | **string** |  | 
**Name** | Pointer to **string** | The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character . must not be used at the end of the name.  | [optional] 

## Methods

### NewNodeBodyMove

`func NewNodeBodyMove(targetParentId string, ) *NodeBodyMove`

NewNodeBodyMove instantiates a new NodeBodyMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBodyMoveWithDefaults

`func NewNodeBodyMoveWithDefaults() *NodeBodyMove`

NewNodeBodyMoveWithDefaults instantiates a new NodeBodyMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetParentId

`func (o *NodeBodyMove) GetTargetParentId() string`

GetTargetParentId returns the TargetParentId field if non-nil, zero value otherwise.

### GetTargetParentIdOk

`func (o *NodeBodyMove) GetTargetParentIdOk() (*string, bool)`

GetTargetParentIdOk returns a tuple with the TargetParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParentId

`func (o *NodeBodyMove) SetTargetParentId(v string)`

SetTargetParentId sets TargetParentId field to given value.


### GetName

`func (o *NodeBodyMove) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeBodyMove) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeBodyMove) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeBodyMove) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


