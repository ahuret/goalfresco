# SiteMembershipRequestWithPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Site** | [**Site**](Site.md) |  | 
**Person** | [**Person**](Person.md) |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewSiteMembershipRequestWithPerson

`func NewSiteMembershipRequestWithPerson(id string, createdAt time.Time, site Site, person Person, ) *SiteMembershipRequestWithPerson`

NewSiteMembershipRequestWithPerson instantiates a new SiteMembershipRequestWithPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMembershipRequestWithPersonWithDefaults

`func NewSiteMembershipRequestWithPersonWithDefaults() *SiteMembershipRequestWithPerson`

NewSiteMembershipRequestWithPersonWithDefaults instantiates a new SiteMembershipRequestWithPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteMembershipRequestWithPerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteMembershipRequestWithPerson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteMembershipRequestWithPerson) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SiteMembershipRequestWithPerson) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SiteMembershipRequestWithPerson) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SiteMembershipRequestWithPerson) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSite

`func (o *SiteMembershipRequestWithPerson) GetSite() Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SiteMembershipRequestWithPerson) GetSiteOk() (*Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *SiteMembershipRequestWithPerson) SetSite(v Site)`

SetSite sets Site field to given value.


### GetPerson

`func (o *SiteMembershipRequestWithPerson) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *SiteMembershipRequestWithPerson) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *SiteMembershipRequestWithPerson) SetPerson(v Person)`

SetPerson sets Person field to given value.


### GetMessage

`func (o *SiteMembershipRequestWithPerson) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SiteMembershipRequestWithPerson) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SiteMembershipRequestWithPerson) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SiteMembershipRequestWithPerson) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


