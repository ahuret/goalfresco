# PersonNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This network&#39;s unique id | 
**HomeNetwork** | Pointer to **bool** | Is this the home network? | [optional] 
**IsEnabled** | **bool** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**PaidNetwork** | Pointer to **bool** |  | [optional] 
**SubscriptionLevel** | Pointer to **string** |  | [optional] 
**Quotas** | Pointer to [**[]NetworkQuota**](NetworkQuota.md) |  | [optional] 

## Methods

### NewPersonNetwork

`func NewPersonNetwork(id string, isEnabled bool, ) *PersonNetwork`

NewPersonNetwork instantiates a new PersonNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonNetworkWithDefaults

`func NewPersonNetworkWithDefaults() *PersonNetwork`

NewPersonNetworkWithDefaults instantiates a new PersonNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetHomeNetwork

`func (o *PersonNetwork) GetHomeNetwork() bool`

GetHomeNetwork returns the HomeNetwork field if non-nil, zero value otherwise.

### GetHomeNetworkOk

`func (o *PersonNetwork) GetHomeNetworkOk() (*bool, bool)`

GetHomeNetworkOk returns a tuple with the HomeNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeNetwork

`func (o *PersonNetwork) SetHomeNetwork(v bool)`

SetHomeNetwork sets HomeNetwork field to given value.

### HasHomeNetwork

`func (o *PersonNetwork) HasHomeNetwork() bool`

HasHomeNetwork returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PersonNetwork) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PersonNetwork) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PersonNetwork) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetCreatedAt

`func (o *PersonNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PersonNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PersonNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PersonNetwork) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPaidNetwork

`func (o *PersonNetwork) GetPaidNetwork() bool`

GetPaidNetwork returns the PaidNetwork field if non-nil, zero value otherwise.

### GetPaidNetworkOk

`func (o *PersonNetwork) GetPaidNetworkOk() (*bool, bool)`

GetPaidNetworkOk returns a tuple with the PaidNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidNetwork

`func (o *PersonNetwork) SetPaidNetwork(v bool)`

SetPaidNetwork sets PaidNetwork field to given value.

### HasPaidNetwork

`func (o *PersonNetwork) HasPaidNetwork() bool`

HasPaidNetwork returns a boolean if a field has been set.

### GetSubscriptionLevel

`func (o *PersonNetwork) GetSubscriptionLevel() string`

GetSubscriptionLevel returns the SubscriptionLevel field if non-nil, zero value otherwise.

### GetSubscriptionLevelOk

`func (o *PersonNetwork) GetSubscriptionLevelOk() (*string, bool)`

GetSubscriptionLevelOk returns a tuple with the SubscriptionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLevel

`func (o *PersonNetwork) SetSubscriptionLevel(v string)`

SetSubscriptionLevel sets SubscriptionLevel field to given value.

### HasSubscriptionLevel

`func (o *PersonNetwork) HasSubscriptionLevel() bool`

HasSubscriptionLevel returns a boolean if a field has been set.

### GetQuotas

`func (o *PersonNetwork) GetQuotas() []NetworkQuota`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *PersonNetwork) GetQuotasOk() (*[]NetworkQuota, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *PersonNetwork) SetQuotas(v []NetworkQuota)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *PersonNetwork) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


