# VersionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewVersionEntry

`func NewVersionEntry() *VersionEntry`

NewVersionEntry instantiates a new VersionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionEntryWithDefaults

`func NewVersionEntryWithDefaults() *VersionEntry`

NewVersionEntryWithDefaults instantiates a new VersionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *VersionEntry) GetEntry() Version`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *VersionEntry) GetEntryOk() (*Version, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *VersionEntry) SetEntry(v Version)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *VersionEntry) HasEntry() bool`

HasEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


