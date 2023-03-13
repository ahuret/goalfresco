# PersonPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Entries** | Pointer to [**[]PersonEntry**](PersonEntry.md) |  | [optional] 

## Methods

### NewPersonPagingList

`func NewPersonPagingList() *PersonPagingList`

NewPersonPagingList instantiates a new PersonPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonPagingListWithDefaults

`func NewPersonPagingListWithDefaults() *PersonPagingList`

NewPersonPagingListWithDefaults instantiates a new PersonPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PersonPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PersonPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PersonPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PersonPagingList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetEntries

`func (o *PersonPagingList) GetEntries() []PersonEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *PersonPagingList) GetEntriesOk() (*[]PersonEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *PersonPagingList) SetEntries(v []PersonEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *PersonPagingList) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


