# DeletedNodeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivedByUser** | [**UserInfo**](UserInfo.md) |  | 
**ArchivedAt** | **time.Time** |  | 

## Methods

### NewDeletedNodeAllOf

`func NewDeletedNodeAllOf(archivedByUser UserInfo, archivedAt time.Time, ) *DeletedNodeAllOf`

NewDeletedNodeAllOf instantiates a new DeletedNodeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedNodeAllOfWithDefaults

`func NewDeletedNodeAllOfWithDefaults() *DeletedNodeAllOf`

NewDeletedNodeAllOfWithDefaults instantiates a new DeletedNodeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivedByUser

`func (o *DeletedNodeAllOf) GetArchivedByUser() UserInfo`

GetArchivedByUser returns the ArchivedByUser field if non-nil, zero value otherwise.

### GetArchivedByUserOk

`func (o *DeletedNodeAllOf) GetArchivedByUserOk() (*UserInfo, bool)`

GetArchivedByUserOk returns a tuple with the ArchivedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedByUser

`func (o *DeletedNodeAllOf) SetArchivedByUser(v UserInfo)`

SetArchivedByUser sets ArchivedByUser field to given value.


### GetArchivedAt

`func (o *DeletedNodeAllOf) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *DeletedNodeAllOf) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *DeletedNodeAllOf) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


