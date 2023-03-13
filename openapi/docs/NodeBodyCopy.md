# NodeBodyCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetParentId** | **string** |  | 
**Name** | Pointer to **string** | The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character . must not be used at the end of the name.  | [optional] 

## Methods

### NewNodeBodyCopy

`func NewNodeBodyCopy(targetParentId string, ) *NodeBodyCopy`

NewNodeBodyCopy instantiates a new NodeBodyCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBodyCopyWithDefaults

`func NewNodeBodyCopyWithDefaults() *NodeBodyCopy`

NewNodeBodyCopyWithDefaults instantiates a new NodeBodyCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetParentId

`func (o *NodeBodyCopy) GetTargetParentId() string`

GetTargetParentId returns the TargetParentId field if non-nil, zero value otherwise.

### GetTargetParentIdOk

`func (o *NodeBodyCopy) GetTargetParentIdOk() (*string, bool)`

GetTargetParentIdOk returns a tuple with the TargetParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParentId

`func (o *NodeBodyCopy) SetTargetParentId(v string)`

SetTargetParentId sets TargetParentId field to given value.


### GetName

`func (o *NodeBodyCopy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeBodyCopy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeBodyCopy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeBodyCopy) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


