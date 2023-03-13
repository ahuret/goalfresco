# PersonNetworkPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Entries** | [**[]PersonNetworkEntry**](PersonNetworkEntry.md) |  | 

## Methods

### NewPersonNetworkPagingList

`func NewPersonNetworkPagingList(pagination Pagination, entries []PersonNetworkEntry, ) *PersonNetworkPagingList`

NewPersonNetworkPagingList instantiates a new PersonNetworkPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonNetworkPagingListWithDefaults

`func NewPersonNetworkPagingListWithDefaults() *PersonNetworkPagingList`

NewPersonNetworkPagingListWithDefaults instantiates a new PersonNetworkPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PersonNetworkPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PersonNetworkPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PersonNetworkPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetEntries

`func (o *PersonNetworkPagingList) GetEntries() []PersonNetworkEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *PersonNetworkPagingList) GetEntriesOk() (*[]PersonNetworkEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *PersonNetworkPagingList) SetEntries(v []PersonNetworkEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


