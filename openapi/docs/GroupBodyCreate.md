# GroupBodyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DisplayName** | **string** |  | 
**ParentIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupBodyCreate

`func NewGroupBodyCreate(id string, displayName string, ) *GroupBodyCreate`

NewGroupBodyCreate instantiates a new GroupBodyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupBodyCreateWithDefaults

`func NewGroupBodyCreateWithDefaults() *GroupBodyCreate`

NewGroupBodyCreateWithDefaults instantiates a new GroupBodyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupBodyCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupBodyCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupBodyCreate) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *GroupBodyCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupBodyCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupBodyCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetParentIds

`func (o *GroupBodyCreate) GetParentIds() []string`

GetParentIds returns the ParentIds field if non-nil, zero value otherwise.

### GetParentIdsOk

`func (o *GroupBodyCreate) GetParentIdsOk() (*[]string, bool)`

GetParentIdsOk returns a tuple with the ParentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIds

`func (o *GroupBodyCreate) SetParentIds(v []string)`

SetParentIds sets ParentIds field to given value.

### HasParentIds

`func (o *GroupBodyCreate) HasParentIds() bool`

HasParentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


