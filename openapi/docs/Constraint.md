# Constraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to **string** | the type of the constraint | [optional] 
**Title** | Pointer to **string** | the human-readable constraint title | [optional] 
**Description** | Pointer to **string** | the human-readable constraint description | [optional] 
**Parameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewConstraint

`func NewConstraint(id string, ) *Constraint`

NewConstraint instantiates a new Constraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintWithDefaults

`func NewConstraintWithDefaults() *Constraint`

NewConstraintWithDefaults instantiates a new Constraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Constraint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Constraint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Constraint) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Constraint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Constraint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Constraint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Constraint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *Constraint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Constraint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Constraint) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Constraint) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Constraint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Constraint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Constraint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Constraint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *Constraint) GetParameters() map[string]map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Constraint) GetParametersOk() (*map[string]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Constraint) SetParameters(v map[string]map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Constraint) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


