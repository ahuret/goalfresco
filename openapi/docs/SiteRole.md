# SiteRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | [**Site**](Site.md) |  | 
**Id** | **string** |  | 
**Guid** | **string** |  | 
**Role** | **string** |  | 

## Methods

### NewSiteRole

`func NewSiteRole(site Site, id string, guid string, role string, ) *SiteRole`

NewSiteRole instantiates a new SiteRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteRoleWithDefaults

`func NewSiteRoleWithDefaults() *SiteRole`

NewSiteRoleWithDefaults instantiates a new SiteRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *SiteRole) GetSite() Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SiteRole) GetSiteOk() (*Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *SiteRole) SetSite(v Site)`

SetSite sets Site field to given value.


### GetId

`func (o *SiteRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteRole) SetId(v string)`

SetId sets Id field to given value.


### GetGuid

`func (o *SiteRole) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SiteRole) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SiteRole) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetRole

`func (o *SiteRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SiteRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SiteRole) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


