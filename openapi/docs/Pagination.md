# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of objects in the entries array.  | [optional] 
**HasMoreItems** | Pointer to **bool** | A boolean value which is **true** if there are more entities in the collection beyond those in this response. A true value means a request with a larger value for the **skipCount** or the **maxItems** parameter will return more entities.  | [optional] 
**TotalItems** | Pointer to **int64** | An integer describing the total number of entities in the collection. The API might not be able to determine this value, in which case this property will not be present.  | [optional] 
**SkipCount** | Pointer to **int64** | An integer describing how many entities exist in the collection before those included in this list. If there was no **skipCount** parameter then the default value is 0.  | [optional] 
**MaxItems** | Pointer to **int64** | The value of the **maxItems** parameter used to generate this list. If there was no **maxItems** parameter then the default value is 100.  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Pagination) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Pagination) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Pagination) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Pagination) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHasMoreItems

`func (o *Pagination) GetHasMoreItems() bool`

GetHasMoreItems returns the HasMoreItems field if non-nil, zero value otherwise.

### GetHasMoreItemsOk

`func (o *Pagination) GetHasMoreItemsOk() (*bool, bool)`

GetHasMoreItemsOk returns a tuple with the HasMoreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreItems

`func (o *Pagination) SetHasMoreItems(v bool)`

SetHasMoreItems sets HasMoreItems field to given value.

### HasHasMoreItems

`func (o *Pagination) HasHasMoreItems() bool`

HasHasMoreItems returns a boolean if a field has been set.

### GetTotalItems

`func (o *Pagination) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *Pagination) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *Pagination) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *Pagination) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetSkipCount

`func (o *Pagination) GetSkipCount() int64`

GetSkipCount returns the SkipCount field if non-nil, zero value otherwise.

### GetSkipCountOk

`func (o *Pagination) GetSkipCountOk() (*int64, bool)`

GetSkipCountOk returns a tuple with the SkipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCount

`func (o *Pagination) SetSkipCount(v int64)`

SetSkipCount sets SkipCount field to given value.

### HasSkipCount

`func (o *Pagination) HasSkipCount() bool`

HasSkipCount returns a boolean if a field has been set.

### GetMaxItems

`func (o *Pagination) GetMaxItems() int64`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *Pagination) GetMaxItemsOk() (*int64, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *Pagination) SetMaxItems(v int64)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *Pagination) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


