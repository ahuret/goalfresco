# NodeChildAssociationPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Entries** | Pointer to [**[]NodeChildAssociationEntry**](NodeChildAssociationEntry.md) |  | [optional] 
**Source** | Pointer to [**Node**](Node.md) |  | [optional] 

## Methods

### NewNodeChildAssociationPagingList

`func NewNodeChildAssociationPagingList() *NodeChildAssociationPagingList`

NewNodeChildAssociationPagingList instantiates a new NodeChildAssociationPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeChildAssociationPagingListWithDefaults

`func NewNodeChildAssociationPagingListWithDefaults() *NodeChildAssociationPagingList`

NewNodeChildAssociationPagingListWithDefaults instantiates a new NodeChildAssociationPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *NodeChildAssociationPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NodeChildAssociationPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NodeChildAssociationPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *NodeChildAssociationPagingList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetEntries

`func (o *NodeChildAssociationPagingList) GetEntries() []NodeChildAssociationEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *NodeChildAssociationPagingList) GetEntriesOk() (*[]NodeChildAssociationEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *NodeChildAssociationPagingList) SetEntries(v []NodeChildAssociationEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *NodeChildAssociationPagingList) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetSource

`func (o *NodeChildAssociationPagingList) GetSource() Node`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NodeChildAssociationPagingList) GetSourceOk() (*Node, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NodeChildAssociationPagingList) SetSource(v Node)`

SetSource sets Source field to given value.

### HasSource

`func (o *NodeChildAssociationPagingList) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


