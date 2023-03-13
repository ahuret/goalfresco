# SiteBodyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | **string** |  | [default to "PUBLIC"]

## Methods

### NewSiteBodyCreate

`func NewSiteBodyCreate(title string, visibility string, ) *SiteBodyCreate`

NewSiteBodyCreate instantiates a new SiteBodyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteBodyCreateWithDefaults

`func NewSiteBodyCreateWithDefaults() *SiteBodyCreate`

NewSiteBodyCreateWithDefaults instantiates a new SiteBodyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteBodyCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteBodyCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteBodyCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteBodyCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *SiteBodyCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SiteBodyCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SiteBodyCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SiteBodyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteBodyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteBodyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteBodyCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *SiteBodyCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SiteBodyCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SiteBodyCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


