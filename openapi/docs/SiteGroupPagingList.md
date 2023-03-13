# SiteGroupPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Entries** | [**[]SiteGroupEntry**](SiteGroupEntry.md) |  | 

## Methods

### NewSiteGroupPagingList

`func NewSiteGroupPagingList(pagination Pagination, entries []SiteGroupEntry, ) *SiteGroupPagingList`

NewSiteGroupPagingList instantiates a new SiteGroupPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteGroupPagingListWithDefaults

`func NewSiteGroupPagingListWithDefaults() *SiteGroupPagingList`

NewSiteGroupPagingListWithDefaults instantiates a new SiteGroupPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SiteGroupPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SiteGroupPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SiteGroupPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetEntries

`func (o *SiteGroupPagingList) GetEntries() []SiteGroupEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *SiteGroupPagingList) GetEntriesOk() (*[]SiteGroupEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *SiteGroupPagingList) SetEntries(v []SiteGroupEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


