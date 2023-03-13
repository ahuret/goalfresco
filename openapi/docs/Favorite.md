# Favorite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetGuid** | **string** | The guid of the object that is a favorite. | 
**CreatedAt** | Pointer to **time.Time** | The time the object was made a favorite. | [optional] 
**Target** | **map[string]interface{}** |  | 
**Properties** | Pointer to **map[string]interface{}** | A subset of the target favorite properties, system properties and properties already available in the target are excluded. | [optional] 

## Methods

### NewFavorite

`func NewFavorite(targetGuid string, target map[string]interface{}, ) *Favorite`

NewFavorite instantiates a new Favorite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteWithDefaults

`func NewFavoriteWithDefaults() *Favorite`

NewFavoriteWithDefaults instantiates a new Favorite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGuid

`func (o *Favorite) GetTargetGuid() string`

GetTargetGuid returns the TargetGuid field if non-nil, zero value otherwise.

### GetTargetGuidOk

`func (o *Favorite) GetTargetGuidOk() (*string, bool)`

GetTargetGuidOk returns a tuple with the TargetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGuid

`func (o *Favorite) SetTargetGuid(v string)`

SetTargetGuid sets TargetGuid field to given value.


### GetCreatedAt

`func (o *Favorite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Favorite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Favorite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Favorite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTarget

`func (o *Favorite) GetTarget() map[string]interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Favorite) GetTargetOk() (*map[string]interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Favorite) SetTarget(v map[string]interface{})`

SetTarget sets Target field to given value.


### GetProperties

`func (o *Favorite) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Favorite) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Favorite) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Favorite) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


