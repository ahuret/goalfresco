# Download

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesAdded** | Pointer to **int32** | number of files added so far in the zip | [optional] 
**BytesAdded** | Pointer to **int32** | number of bytes added so far in the zip | [optional] 
**Id** | Pointer to **string** | the id of the download node | [optional] 
**TotalFiles** | Pointer to **int32** | the total number of files to be added in the zip | [optional] 
**TotalBytes** | Pointer to **int32** | the total number of bytes to be added in the zip | [optional] 
**Status** | Pointer to **string** | the current status of the download node creation | [optional] [default to "PENDING"]

## Methods

### NewDownload

`func NewDownload() *Download`

NewDownload instantiates a new Download object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadWithDefaults

`func NewDownloadWithDefaults() *Download`

NewDownloadWithDefaults instantiates a new Download object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesAdded

`func (o *Download) GetFilesAdded() int32`

GetFilesAdded returns the FilesAdded field if non-nil, zero value otherwise.

### GetFilesAddedOk

`func (o *Download) GetFilesAddedOk() (*int32, bool)`

GetFilesAddedOk returns a tuple with the FilesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesAdded

`func (o *Download) SetFilesAdded(v int32)`

SetFilesAdded sets FilesAdded field to given value.

### HasFilesAdded

`func (o *Download) HasFilesAdded() bool`

HasFilesAdded returns a boolean if a field has been set.

### GetBytesAdded

`func (o *Download) GetBytesAdded() int32`

GetBytesAdded returns the BytesAdded field if non-nil, zero value otherwise.

### GetBytesAddedOk

`func (o *Download) GetBytesAddedOk() (*int32, bool)`

GetBytesAddedOk returns a tuple with the BytesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesAdded

`func (o *Download) SetBytesAdded(v int32)`

SetBytesAdded sets BytesAdded field to given value.

### HasBytesAdded

`func (o *Download) HasBytesAdded() bool`

HasBytesAdded returns a boolean if a field has been set.

### GetId

`func (o *Download) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Download) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Download) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Download) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTotalFiles

`func (o *Download) GetTotalFiles() int32`

GetTotalFiles returns the TotalFiles field if non-nil, zero value otherwise.

### GetTotalFilesOk

`func (o *Download) GetTotalFilesOk() (*int32, bool)`

GetTotalFilesOk returns a tuple with the TotalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFiles

`func (o *Download) SetTotalFiles(v int32)`

SetTotalFiles sets TotalFiles field to given value.

### HasTotalFiles

`func (o *Download) HasTotalFiles() bool`

HasTotalFiles returns a boolean if a field has been set.

### GetTotalBytes

`func (o *Download) GetTotalBytes() int32`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *Download) GetTotalBytesOk() (*int32, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *Download) SetTotalBytes(v int32)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *Download) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetStatus

`func (o *Download) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Download) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Download) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Download) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


