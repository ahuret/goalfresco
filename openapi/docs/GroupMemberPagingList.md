# GroupMemberPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Entries** | Pointer to [**[]GroupMemberEntry**](GroupMemberEntry.md) |  | [optional] 

## Methods

### NewGroupMemberPagingList

`func NewGroupMemberPagingList() *GroupMemberPagingList`

NewGroupMemberPagingList instantiates a new GroupMemberPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberPagingListWithDefaults

`func NewGroupMemberPagingListWithDefaults() *GroupMemberPagingList`

NewGroupMemberPagingListWithDefaults instantiates a new GroupMemberPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GroupMemberPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GroupMemberPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GroupMemberPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GroupMemberPagingList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetEntries

`func (o *GroupMemberPagingList) GetEntries() []GroupMemberEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *GroupMemberPagingList) GetEntriesOk() (*[]GroupMemberEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *GroupMemberPagingList) SetEntries(v []GroupMemberEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *GroupMemberPagingList) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


