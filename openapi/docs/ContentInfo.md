# ContentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MimeType** | **string** |  | 
**MimeTypeName** | Pointer to **string** |  | [optional] 
**SizeInBytes** | Pointer to **int64** |  | [optional] 
**Encoding** | Pointer to **string** |  | [optional] 

## Methods

### NewContentInfo

`func NewContentInfo(mimeType string, ) *ContentInfo`

NewContentInfo instantiates a new ContentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentInfoWithDefaults

`func NewContentInfoWithDefaults() *ContentInfo`

NewContentInfoWithDefaults instantiates a new ContentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMimeType

`func (o *ContentInfo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ContentInfo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ContentInfo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetMimeTypeName

`func (o *ContentInfo) GetMimeTypeName() string`

GetMimeTypeName returns the MimeTypeName field if non-nil, zero value otherwise.

### GetMimeTypeNameOk

`func (o *ContentInfo) GetMimeTypeNameOk() (*string, bool)`

GetMimeTypeNameOk returns a tuple with the MimeTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypeName

`func (o *ContentInfo) SetMimeTypeName(v string)`

SetMimeTypeName sets MimeTypeName field to given value.

### HasMimeTypeName

`func (o *ContentInfo) HasMimeTypeName() bool`

HasMimeTypeName returns a boolean if a field has been set.

### GetSizeInBytes

`func (o *ContentInfo) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *ContentInfo) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *ContentInfo) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *ContentInfo) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.

### GetEncoding

`func (o *ContentInfo) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ContentInfo) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ContentInfo) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *ContentInfo) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


