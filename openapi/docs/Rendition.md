# Rendition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**ContentInfo**](ContentInfo.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewRendition

`func NewRendition() *Rendition`

NewRendition instantiates a new Rendition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenditionWithDefaults

`func NewRenditionWithDefaults() *Rendition`

NewRenditionWithDefaults instantiates a new Rendition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rendition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rendition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rendition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Rendition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *Rendition) GetContent() ContentInfo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Rendition) GetContentOk() (*ContentInfo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Rendition) SetContent(v ContentInfo)`

SetContent sets Content field to given value.

### HasContent

`func (o *Rendition) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStatus

`func (o *Rendition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Rendition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Rendition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Rendition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


