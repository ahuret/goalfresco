# SiteMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Person** | [**Person**](Person.md) |  | 
**Role** | **string** |  | 
**IsMemberOfGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewSiteMember

`func NewSiteMember(id string, person Person, role string, ) *SiteMember`

NewSiteMember instantiates a new SiteMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMemberWithDefaults

`func NewSiteMemberWithDefaults() *SiteMember`

NewSiteMemberWithDefaults instantiates a new SiteMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteMember) SetId(v string)`

SetId sets Id field to given value.


### GetPerson

`func (o *SiteMember) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *SiteMember) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *SiteMember) SetPerson(v Person)`

SetPerson sets Person field to given value.


### GetRole

`func (o *SiteMember) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SiteMember) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SiteMember) SetRole(v string)`

SetRole sets Role field to given value.


### GetIsMemberOfGroup

`func (o *SiteMember) GetIsMemberOfGroup() bool`

GetIsMemberOfGroup returns the IsMemberOfGroup field if non-nil, zero value otherwise.

### GetIsMemberOfGroupOk

`func (o *SiteMember) GetIsMemberOfGroupOk() (*bool, bool)`

GetIsMemberOfGroupOk returns a tuple with the IsMemberOfGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMemberOfGroup

`func (o *SiteMember) SetIsMemberOfGroup(v bool)`

SetIsMemberOfGroup sets IsMemberOfGroup field to given value.

### HasIsMemberOfGroup

`func (o *SiteMember) HasIsMemberOfGroup() bool`

HasIsMemberOfGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


