# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Guid** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | **string** |  | 
**Preset** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewSite

`func NewSite(id string, guid string, title string, visibility string, ) *Site`

NewSite instantiates a new Site object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWithDefaults

`func NewSiteWithDefaults() *Site`

NewSiteWithDefaults instantiates a new Site object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Site) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Site) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Site) SetId(v string)`

SetId sets Id field to given value.


### GetGuid

`func (o *Site) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Site) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Site) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetTitle

`func (o *Site) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Site) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Site) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Site) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Site) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Site) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Site) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *Site) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Site) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Site) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetPreset

`func (o *Site) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *Site) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *Site) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *Site) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetRole

`func (o *Site) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Site) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Site) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Site) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


