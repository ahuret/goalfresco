# SiteGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**GroupMember**](GroupMember.md) |  | 
**Role** | **string** |  | 

## Methods

### NewSiteGroup

`func NewSiteGroup(id string, group GroupMember, role string, ) *SiteGroup`

NewSiteGroup instantiates a new SiteGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteGroupWithDefaults

`func NewSiteGroupWithDefaults() *SiteGroup`

NewSiteGroupWithDefaults instantiates a new SiteGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteGroup) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *SiteGroup) GetGroup() GroupMember`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SiteGroup) GetGroupOk() (*GroupMember, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SiteGroup) SetGroup(v GroupMember)`

SetGroup sets Group field to given value.


### GetRole

`func (o *SiteGroup) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SiteGroup) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SiteGroup) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


