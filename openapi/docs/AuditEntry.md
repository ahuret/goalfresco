# AuditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AuditApplicationId** | **string** |  | 
**CreatedByUser** | [**UserInfo**](UserInfo.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Values** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAuditEntry

`func NewAuditEntry(id string, auditApplicationId string, createdByUser UserInfo, createdAt time.Time, ) *AuditEntry`

NewAuditEntry instantiates a new AuditEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEntryWithDefaults

`func NewAuditEntryWithDefaults() *AuditEntry`

NewAuditEntryWithDefaults instantiates a new AuditEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditEntry) SetId(v string)`

SetId sets Id field to given value.


### GetAuditApplicationId

`func (o *AuditEntry) GetAuditApplicationId() string`

GetAuditApplicationId returns the AuditApplicationId field if non-nil, zero value otherwise.

### GetAuditApplicationIdOk

`func (o *AuditEntry) GetAuditApplicationIdOk() (*string, bool)`

GetAuditApplicationIdOk returns a tuple with the AuditApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditApplicationId

`func (o *AuditEntry) SetAuditApplicationId(v string)`

SetAuditApplicationId sets AuditApplicationId field to given value.


### GetCreatedByUser

`func (o *AuditEntry) GetCreatedByUser() UserInfo`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *AuditEntry) GetCreatedByUserOk() (*UserInfo, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *AuditEntry) SetCreatedByUser(v UserInfo)`

SetCreatedByUser sets CreatedByUser field to given value.


### GetCreatedAt

`func (o *AuditEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetValues

`func (o *AuditEntry) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AuditEntry) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AuditEntry) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *AuditEntry) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


