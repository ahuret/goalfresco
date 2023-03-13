# AuditApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**MaxEntryId** | Pointer to **int32** |  | [optional] 
**MinEntryId** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuditApp

`func NewAuditApp(id string, ) *AuditApp`

NewAuditApp instantiates a new AuditApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditAppWithDefaults

`func NewAuditAppWithDefaults() *AuditApp`

NewAuditAppWithDefaults instantiates a new AuditApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditApp) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AuditApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AuditApp) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AuditApp) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AuditApp) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AuditApp) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMaxEntryId

`func (o *AuditApp) GetMaxEntryId() int32`

GetMaxEntryId returns the MaxEntryId field if non-nil, zero value otherwise.

### GetMaxEntryIdOk

`func (o *AuditApp) GetMaxEntryIdOk() (*int32, bool)`

GetMaxEntryIdOk returns a tuple with the MaxEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntryId

`func (o *AuditApp) SetMaxEntryId(v int32)`

SetMaxEntryId sets MaxEntryId field to given value.

### HasMaxEntryId

`func (o *AuditApp) HasMaxEntryId() bool`

HasMaxEntryId returns a boolean if a field has been set.

### GetMinEntryId

`func (o *AuditApp) GetMinEntryId() int32`

GetMinEntryId returns the MinEntryId field if non-nil, zero value otherwise.

### GetMinEntryIdOk

`func (o *AuditApp) GetMinEntryIdOk() (*int32, bool)`

GetMinEntryIdOk returns a tuple with the MinEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEntryId

`func (o *AuditApp) SetMinEntryId(v int32)`

SetMinEntryId sets MinEntryId field to given value.

### HasMinEntryId

`func (o *AuditApp) HasMinEntryId() bool`

HasMinEntryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


