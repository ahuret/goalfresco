# DirectAccessUrlBodyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**ValidFor** | Pointer to **int32** | The length of time in seconds that the url is valid for.  | [optional] 

## Methods

### NewDirectAccessUrlBodyCreate

`func NewDirectAccessUrlBodyCreate() *DirectAccessUrlBodyCreate`

NewDirectAccessUrlBodyCreate instantiates a new DirectAccessUrlBodyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectAccessUrlBodyCreateWithDefaults

`func NewDirectAccessUrlBodyCreateWithDefaults() *DirectAccessUrlBodyCreate`

NewDirectAccessUrlBodyCreateWithDefaults instantiates a new DirectAccessUrlBodyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *DirectAccessUrlBodyCreate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DirectAccessUrlBodyCreate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DirectAccessUrlBodyCreate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DirectAccessUrlBodyCreate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetValidFor

`func (o *DirectAccessUrlBodyCreate) GetValidFor() int32`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *DirectAccessUrlBodyCreate) GetValidForOk() (*int32, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *DirectAccessUrlBodyCreate) SetValidFor(v int32)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *DirectAccessUrlBodyCreate) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


