# NodeAssociationPagingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Entries** | Pointer to [**[]NodeAssociationEntry**](NodeAssociationEntry.md) |  | [optional] 
**Source** | Pointer to [**Node**](Node.md) |  | [optional] 

## Methods

### NewNodeAssociationPagingList

`func NewNodeAssociationPagingList() *NodeAssociationPagingList`

NewNodeAssociationPagingList instantiates a new NodeAssociationPagingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAssociationPagingListWithDefaults

`func NewNodeAssociationPagingListWithDefaults() *NodeAssociationPagingList`

NewNodeAssociationPagingListWithDefaults instantiates a new NodeAssociationPagingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *NodeAssociationPagingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NodeAssociationPagingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NodeAssociationPagingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *NodeAssociationPagingList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetEntries

`func (o *NodeAssociationPagingList) GetEntries() []NodeAssociationEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *NodeAssociationPagingList) GetEntriesOk() (*[]NodeAssociationEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *NodeAssociationPagingList) SetEntries(v []NodeAssociationEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *NodeAssociationPagingList) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetSource

`func (o *NodeAssociationPagingList) GetSource() Node`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NodeAssociationPagingList) GetSourceOk() (*Node, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NodeAssociationPagingList) SetSource(v Node)`

SetSource sets Source field to given value.

### HasSource

`func (o *NodeAssociationPagingList) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


