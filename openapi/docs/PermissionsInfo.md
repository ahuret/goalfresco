# PermissionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInheritanceEnabled** | Pointer to **bool** |  | [optional] 
**Inherited** | Pointer to [**[]PermissionElement**](PermissionElement.md) |  | [optional] 
**LocallySet** | Pointer to [**[]PermissionElement**](PermissionElement.md) |  | [optional] 
**Settable** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPermissionsInfo

`func NewPermissionsInfo() *PermissionsInfo`

NewPermissionsInfo instantiates a new PermissionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsInfoWithDefaults

`func NewPermissionsInfoWithDefaults() *PermissionsInfo`

NewPermissionsInfoWithDefaults instantiates a new PermissionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsInheritanceEnabled

`func (o *PermissionsInfo) GetIsInheritanceEnabled() bool`

GetIsInheritanceEnabled returns the IsInheritanceEnabled field if non-nil, zero value otherwise.

### GetIsInheritanceEnabledOk

`func (o *PermissionsInfo) GetIsInheritanceEnabledOk() (*bool, bool)`

GetIsInheritanceEnabledOk returns a tuple with the IsInheritanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInheritanceEnabled

`func (o *PermissionsInfo) SetIsInheritanceEnabled(v bool)`

SetIsInheritanceEnabled sets IsInheritanceEnabled field to given value.

### HasIsInheritanceEnabled

`func (o *PermissionsInfo) HasIsInheritanceEnabled() bool`

HasIsInheritanceEnabled returns a boolean if a field has been set.

### GetInherited

`func (o *PermissionsInfo) GetInherited() []PermissionElement`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *PermissionsInfo) GetInheritedOk() (*[]PermissionElement, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *PermissionsInfo) SetInherited(v []PermissionElement)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *PermissionsInfo) HasInherited() bool`

HasInherited returns a boolean if a field has been set.

### GetLocallySet

`func (o *PermissionsInfo) GetLocallySet() []PermissionElement`

GetLocallySet returns the LocallySet field if non-nil, zero value otherwise.

### GetLocallySetOk

`func (o *PermissionsInfo) GetLocallySetOk() (*[]PermissionElement, bool)`

GetLocallySetOk returns a tuple with the LocallySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocallySet

`func (o *PermissionsInfo) SetLocallySet(v []PermissionElement)`

SetLocallySet sets LocallySet field to given value.

### HasLocallySet

`func (o *PermissionsInfo) HasLocallySet() bool`

HasLocallySet returns a boolean if a field has been set.

### GetSettable

`func (o *PermissionsInfo) GetSettable() []string`

GetSettable returns the Settable field if non-nil, zero value otherwise.

### GetSettableOk

`func (o *PermissionsInfo) GetSettableOk() (*[]string, bool)`

GetSettableOk returns a tuple with the Settable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettable

`func (o *PermissionsInfo) SetSettable(v []string)`

SetSettable sets Settable field to given value.

### HasSettable

`func (o *PermissionsInfo) HasSettable() bool`

HasSettable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


