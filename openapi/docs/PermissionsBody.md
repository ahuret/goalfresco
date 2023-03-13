# PermissionsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInheritanceEnabled** | Pointer to **bool** |  | [optional] 
**LocallySet** | Pointer to [**[]PermissionElement**](PermissionElement.md) |  | [optional] 

## Methods

### NewPermissionsBody

`func NewPermissionsBody() *PermissionsBody`

NewPermissionsBody instantiates a new PermissionsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsBodyWithDefaults

`func NewPermissionsBodyWithDefaults() *PermissionsBody`

NewPermissionsBodyWithDefaults instantiates a new PermissionsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsInheritanceEnabled

`func (o *PermissionsBody) GetIsInheritanceEnabled() bool`

GetIsInheritanceEnabled returns the IsInheritanceEnabled field if non-nil, zero value otherwise.

### GetIsInheritanceEnabledOk

`func (o *PermissionsBody) GetIsInheritanceEnabledOk() (*bool, bool)`

GetIsInheritanceEnabledOk returns a tuple with the IsInheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInheritanceEnabled

`func (o *PermissionsBody) SetIsInheritanceEnabled(v bool)`

SetIsInheritanceEnabled sets IsInheritanceEnabled field to given value.

### HasIsInheritanceEnabled

`func (o *PermissionsBody) HasIsInheritanceEnabled() bool`

HasIsInheritanceEnabled returns a boolean if a field has been set.

### GetLocallySet

`func (o *PermissionsBody) GetLocallySet() []PermissionElement`

GetLocallySet returns the LocallySet field if non-nil, zero value otherwise.

### GetLocallySetOk

`func (o *PermissionsBody) GetLocallySetOk() (*[]PermissionElement, bool)`

GetLocallySetOk returns a tuple with the LocallySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallySet

`func (o *PermissionsBody) SetLocallySet(v []PermissionElement)`

SetLocallySet sets LocallySet field to given value.

### HasLocallySet

`func (o *PermissionsBody) HasLocallySet() bool`

HasLocallySet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


