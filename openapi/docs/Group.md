# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DisplayName** | **string** |  | 
**IsRoot** | **bool** |  | [default to true]
**ParentIds** | Pointer to **[]string** |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroup

`func NewGroup(id string, displayName string, isRoot bool, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *Group) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Group) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsRoot

`func (o *Group) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *Group) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *Group) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.


### GetParentIds

`func (o *Group) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *Group) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *Group) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.

### HasParentIds

`func (o *Group) HasParentIds() bool`

HasParentIds returns a boolean if a field has been set.

### GetZones

`func (o *Group) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *Group) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *Group) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *Group) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


