# ActionParameterDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MultiValued** | Pointer to **bool** |  | [optional] 
**Mandatory** | Pointer to **bool** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 

## Methods

### NewActionParameterDefinition

`func NewActionParameterDefinition() *ActionParameterDefinition`

NewActionParameterDefinition instantiates a new ActionParameterDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionParameterDefinitionWithDefaults

`func NewActionParameterDefinitionWithDefaults() *ActionParameterDefinition`

NewActionParameterDefinitionWithDefaults instantiates a new ActionParameterDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionParameterDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionParameterDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionParameterDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionParameterDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ActionParameterDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionParameterDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionParameterDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionParameterDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMultiValued

`func (o *ActionParameterDefinition) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *ActionParameterDefinition) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *ActionParameterDefinition) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *ActionParameterDefinition) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetMandatory

`func (o *ActionParameterDefinition) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ActionParameterDefinition) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ActionParameterDefinition) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *ActionParameterDefinition) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *ActionParameterDefinition) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *ActionParameterDefinition) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *ActionParameterDefinition) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *ActionParameterDefinition) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


