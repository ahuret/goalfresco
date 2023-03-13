# SharedLinkBodyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** |  | 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSharedLinkBodyCreate

`func NewSharedLinkBodyCreate(nodeId string, ) *SharedLinkBodyCreate`

NewSharedLinkBodyCreate instantiates a new SharedLinkBodyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedLinkBodyCreateWithDefaults

`func NewSharedLinkBodyCreateWithDefaults() *SharedLinkBodyCreate`

NewSharedLinkBodyCreateWithDefaults instantiates a new SharedLinkBodyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *SharedLinkBodyCreate) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SharedLinkBodyCreate) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SharedLinkBodyCreate) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetExpiresAt

`func (o *SharedLinkBodyCreate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SharedLinkBodyCreate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SharedLinkBodyCreate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SharedLinkBodyCreate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


