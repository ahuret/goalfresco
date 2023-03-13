# SiteMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Site** | [**Site**](Site.md) |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewSiteMembershipRequest

`func NewSiteMembershipRequest(id string, createdAt time.Time, site Site, ) *SiteMembershipRequest`

NewSiteMembershipRequest instantiates a new SiteMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMembershipRequestWithDefaults

`func NewSiteMembershipRequestWithDefaults() *SiteMembershipRequest`

NewSiteMembershipRequestWithDefaults instantiates a new SiteMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteMembershipRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteMembershipRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteMembershipRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SiteMembershipRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SiteMembershipRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SiteMembershipRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSite

`func (o *SiteMembershipRequest) GetSite() Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SiteMembershipRequest) GetSiteOk() (*Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *SiteMembershipRequest) SetSite(v Site)`

SetSite sets Site field to given value.


### GetMessage

`func (o *SiteMembershipRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SiteMembershipRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SiteMembershipRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SiteMembershipRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


