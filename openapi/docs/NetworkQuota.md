# NetworkQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Limit** | **int64** |  | 
**Usage** | **int64** |  | 

## Methods

### NewNetworkQuota

`func NewNetworkQuota(id string, limit int64, usage int64, ) *NetworkQuota`

NewNetworkQuota instantiates a new NetworkQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkQuotaWithDefaults

`func NewNetworkQuotaWithDefaults() *NetworkQuota`

NewNetworkQuotaWithDefaults instantiates a new NetworkQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkQuota) SetId(v string)`

SetId sets Id field to given value.


### GetLimit

`func (o *NetworkQuota) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NetworkQuota) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NetworkQuota) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetUsage

`func (o *NetworkQuota) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *NetworkQuota) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *NetworkQuota) SetUsage(v int64)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


