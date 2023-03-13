# ActionBodyExec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDefinitionId** | **string** |  | 
**TargetId** | Pointer to **string** | The entity upon which to execute the action, typically a node ID or similar. | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewActionBodyExec

`func NewActionBodyExec(actionDefinitionId string, ) *ActionBodyExec`

NewActionBodyExec instantiates a new ActionBodyExec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionBodyExecWithDefaults

`func NewActionBodyExecWithDefaults() *ActionBodyExec`

NewActionBodyExecWithDefaults instantiates a new ActionBodyExec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDefinitionId

`func (o *ActionBodyExec) GetActionDefinitionId() string`

GetActionDefinitionId returns the ActionDefinitionId field if non-nil, zero value otherwise.

### GetActionDefinitionIdOk

`func (o *ActionBodyExec) GetActionDefinitionIdOk() (*string, bool)`

GetActionDefinitionIdOk returns a tuple with the ActionDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDefinitionId

`func (o *ActionBodyExec) SetActionDefinitionId(v string)`

SetActionDefinitionId sets ActionDefinitionId field to given value.


### GetTargetId

`func (o *ActionBodyExec) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ActionBodyExec) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ActionBodyExec) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ActionBodyExec) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetParams

`func (o *ActionBodyExec) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ActionBodyExec) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ActionBodyExec) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ActionBodyExec) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


