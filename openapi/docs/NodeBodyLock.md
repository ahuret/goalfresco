# NodeBodyLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeToExpire** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "ALLOW_OWNER_CHANGES"]
**Lifetime** | Pointer to **string** |  | [optional] [default to "PERSISTENT"]

## Methods

### NewNodeBodyLock

`func NewNodeBodyLock() *NodeBodyLock`

NewNodeBodyLock instantiates a new NodeBodyLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBodyLockWithDefaults

`func NewNodeBodyLockWithDefaults() *NodeBodyLock`

NewNodeBodyLockWithDefaults instantiates a new NodeBodyLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeToExpire

`func (o *NodeBodyLock) GetTimeToExpire() int32`

GetTimeToExpire returns the TimeToExpire field if non-nil, zero value otherwise.

### GetTimeToExpireOk

`func (o *NodeBodyLock) GetTimeToExpireOk() (*int32, bool)`

GetTimeToExpireOk returns a tuple with the TimeToExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToExpire

`func (o *NodeBodyLock) SetTimeToExpire(v int32)`

SetTimeToExpire sets TimeToExpire field to given value.

### HasTimeToExpire

`func (o *NodeBodyLock) HasTimeToExpire() bool`

HasTimeToExpire returns a boolean if a field has been set.

### GetType

`func (o *NodeBodyLock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeBodyLock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeBodyLock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeBodyLock) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLifetime

`func (o *NodeBodyLock) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *NodeBodyLock) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *NodeBodyLock) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *NodeBodyLock) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


