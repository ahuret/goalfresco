# DeletedNodesPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Entries** | Pointer to [**[]DeletedNodeEntry**](DeletedNodeEntry.md) |  | [optional] 

## Methods

### NewDeletedNodesPagingList

`func NewDeletedNodesPagingList() *DeletedNodesPagingList`

NewDeletedNodesPagingList instantiates a new DeletedNodesPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedNodesPagingListWithDefaults

`func NewDeletedNodesPagingListWithDefaults() *DeletedNodesPagingList`

NewDeletedNodesPagingListWithDefaults instantiates a new DeletedNodesPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *DeletedNodesPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DeletedNodesPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DeletedNodesPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DeletedNodesPagingList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetEntries

`func (o *DeletedNodesPagingList) GetEntries() []DeletedNodeEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *DeletedNodesPagingList) GetEntriesOk() (*[]DeletedNodeEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *DeletedNodesPagingList) SetEntries(v []DeletedNodeEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *DeletedNodesPagingList) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


