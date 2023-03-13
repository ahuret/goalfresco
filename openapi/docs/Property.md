# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | Pointer to **string** | the human-readable title | [optional] 
**Description** | Pointer to **string** | the human-readable description | [optional] 
**DefaultValue** | Pointer to **string** | the default value | [optional] 
**DataType** | Pointer to **string** | the name of the property type (e.g. d:text) | [optional] 
**IsMultiValued** | Pointer to **bool** | define if the property is multi-valued | [optional] 
**IsMandatory** | Pointer to **bool** | define if the property is mandatory | [optional] 
**IsMandatoryEnforced** | Pointer to **bool** | define if the presence of mandatory properties is enforced | [optional] 
**IsProtected** | Pointer to **bool** | define if the property is system maintained | [optional] 
**Constraints** | Pointer to [**[]Constraint**](Constraint.md) | list of constraints defined for the property | [optional] 

## Methods

### NewProperty

`func NewProperty(id string, ) *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Property) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Property) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Property) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Property) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Property) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Property) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Property) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Property) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Property) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Property) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Property) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *Property) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *Property) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *Property) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *Property) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDataType

`func (o *Property) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Property) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Property) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *Property) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetIsMultiValued

`func (o *Property) GetIsMultiValued() bool`

GetIsMultiValued returns the IsMultiValued field if non-nil, zero value otherwise.

### GetIsMultiValuedOk

`func (o *Property) GetIsMultiValuedOk() (*bool, bool)`

GetIsMultiValuedOk returns a tuple with the IsMultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiValued

`func (o *Property) SetIsMultiValued(v bool)`

SetIsMultiValued sets IsMultiValued field to given value.

### HasIsMultiValued

`func (o *Property) HasIsMultiValued() bool`

HasIsMultiValued returns a boolean if a field has been set.

### GetIsMandatory

`func (o *Property) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *Property) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *Property) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *Property) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.

### GetIsMandatoryEnforced

`func (o *Property) GetIsMandatoryEnforced() bool`

GetIsMandatoryEnforced returns the IsMandatoryEnforced field if non-nil, zero value otherwise.

### GetIsMandatoryEnforcedOk

`func (o *Property) GetIsMandatoryEnforcedOk() (*bool, bool)`

GetIsMandatoryEnforcedOk returns a tuple with the IsMandatoryEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatoryEnforced

`func (o *Property) SetIsMandatoryEnforced(v bool)`

SetIsMandatoryEnforced sets IsMandatoryEnforced field to given value.

### HasIsMandatoryEnforced

`func (o *Property) HasIsMandatoryEnforced() bool`

HasIsMandatoryEnforced returns a boolean if a field has been set.

### GetIsProtected

`func (o *Property) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *Property) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *Property) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *Property) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetConstraints

`func (o *Property) GetConstraints() []Constraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Property) GetConstraintsOk() (*[]Constraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Property) SetConstraints(v []Constraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Property) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


