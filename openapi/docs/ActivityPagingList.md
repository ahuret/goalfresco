# ActivityPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Entries** | [**[]ActivityEntry**](ActivityEntry.md) |  | 

## Methods

### NewActivityPagingList

`func NewActivityPagingList(pagination Pagination, entries []ActivityEntry, ) *ActivityPagingList`

NewActivityPagingList instantiates a new ActivityPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityPagingListWithDefaults

`func NewActivityPagingListWithDefaults() *ActivityPagingList`

NewActivityPagingListWithDefaults instantiates a new ActivityPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ActivityPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ActivityPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ActivityPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetEntries

`func (o *ActivityPagingList) GetEntries() []ActivityEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ActivityPagingList) GetEntriesOk() (*[]ActivityEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ActivityPagingList) SetEntries(v []ActivityEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


