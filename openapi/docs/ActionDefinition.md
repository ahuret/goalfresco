# ActionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the action definition â€” used for example when executing an action | 
**Name** | Pointer to **string** | name of the action definition, e.g. \&quot;move\&quot; | [optional] 
**Title** | Pointer to **string** | title of the action definition, e.g. \&quot;Move\&quot; | [optional] 
**Description** | Pointer to **string** | describes the action definition, e.g. \&quot;This will move the matched item to another space.\&quot; | [optional] 
**ApplicableTypes** | **[]string** | QNames of the types this action applies to | 
**TrackStatus** | **bool** | whether the basic action definition supports action tracking or not | 
**ParameterDefinitions** | Pointer to [**[]ActionParameterDefinition**](ActionParameterDefinition.md) |  | [optional] 

## Methods

### NewActionDefinition

`func NewActionDefinition(id string, applicableTypes []string, trackStatus bool, ) *ActionDefinition`

NewActionDefinition instantiates a new ActionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDefinitionWithDefaults

`func NewActionDefinitionWithDefaults() *ActionDefinition`

NewActionDefinitionWithDefaults instantiates a new ActionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ActionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ActionDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActionDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActionDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActionDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ActionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplicableTypes

`func (o *ActionDefinition) GetApplicableTypes() []string`

GetApplicableTypes returns the ApplicableTypes field if non-nil, zero value otherwise.

### GetApplicableTypesOk

`func (o *ActionDefinition) GetApplicableTypesOk() (*[]string, bool)`

GetApplicableTypesOk returns a tuple with the ApplicableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableTypes

`func (o *ActionDefinition) SetApplicableTypes(v []string)`

SetApplicableTypes sets ApplicableTypes field to given value.


### GetTrackStatus

`func (o *ActionDefinition) GetTrackStatus() bool`

GetTrackStatus returns the TrackStatus field if non-nil, zero value otherwise.

### GetTrackStatusOk

`func (o *ActionDefinition) GetTrackStatusOk() (*bool, bool)`

GetTrackStatusOk returns a tuple with the TrackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackStatus

`func (o *ActionDefinition) SetTrackStatus(v bool)`

SetTrackStatus sets TrackStatus field to given value.


### GetParameterDefinitions

`func (o *ActionDefinition) GetParameterDefinitions() []ActionParameterDefinition`

GetParameterDefinitions returns the ParameterDefinitions field if non-nil, zero value otherwise.

### GetParameterDefinitionsOk

`func (o *ActionDefinition) GetParameterDefinitionsOk() (*[]ActionParameterDefinition, bool)`

GetParameterDefinitionsOk returns a tuple with the ParameterDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterDefinitions

`func (o *ActionDefinition) SetParameterDefinitions(v []ActionParameterDefinition)`

SetParameterDefinitions sets ParameterDefinitions field to given value.

### HasParameterDefinitions

`func (o *ActionDefinition) HasParameterDefinitions() bool`

HasParameterDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


