# SiteRolePagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Entries** | [**[]SiteRoleEntry**](SiteRoleEntry.md) |  | 

## Methods

### NewSiteRolePagingList

`func NewSiteRolePagingList(pagination Pagination, entries []SiteRoleEntry, ) *SiteRolePagingList`

NewSiteRolePagingList instantiates a new SiteRolePagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteRolePagingListWithDefaults

`func NewSiteRolePagingListWithDefaults() *SiteRolePagingList`

NewSiteRolePagingListWithDefaults instantiates a new SiteRolePagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SiteRolePagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SiteRolePagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SiteRolePagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetEntries

`func (o *SiteRolePagingList) GetEntries() []SiteRoleEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *SiteRolePagingList) GetEntriesOk() (*[]SiteRoleEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *SiteRolePagingList) SetEntries(v []SiteRoleEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


