# FavoritePagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Entries** | [**[]FavoriteEntry**](FavoriteEntry.md) |  | 

## Methods

### NewFavoritePagingList

`func NewFavoritePagingList(pagination Pagination, entries []FavoriteEntry, ) *FavoritePagingList`

NewFavoritePagingList instantiates a new FavoritePagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoritePagingListWithDefaults

`func NewFavoritePagingListWithDefaults() *FavoritePagingList`

NewFavoritePagingListWithDefaults instantiates a new FavoritePagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *FavoritePagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FavoritePagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FavoritePagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetEntries

`func (o *FavoritePagingList) GetEntries() []FavoriteEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *FavoritePagingList) GetEntriesOk() (*[]FavoriteEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *FavoritePagingList) SetEntries(v []FavoriteEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


