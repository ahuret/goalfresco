# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Content** | **string** |  | 
**CreatedBy** | [**Person**](Person.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Edited** | **bool** |  | 
**ModifiedBy** | [**Person**](Person.md) |  | 
**ModifiedAt** | **time.Time** |  | 
**CanEdit** | **bool** |  | 
**CanDelete** | **bool** |  | 

## Methods

### NewComment

`func NewComment(id string, title string, content string, createdBy Person, createdAt time.Time, edited bool, modifiedBy Person, modifiedAt time.Time, canEdit bool, canDelete bool, ) *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Comment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Comment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Comment) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *Comment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Comment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Comment) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedBy

`func (o *Comment) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Comment) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Comment) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Comment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Comment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Comment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEdited

`func (o *Comment) GetEdited() bool`

GetEdited returns the Edited field if non-nil, zero value otherwise.

### GetEditedOk

`func (o *Comment) GetEditedOk() (*bool, bool)`

GetEditedOk returns a tuple with the Edited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdited

`func (o *Comment) SetEdited(v bool)`

SetEdited sets Edited field to given value.


### GetModifiedBy

`func (o *Comment) GetModifiedBy() Person`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Comment) GetModifiedByOk() (*Person, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Comment) SetModifiedBy(v Person)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifiedAt

`func (o *Comment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Comment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Comment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetCanEdit

`func (o *Comment) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *Comment) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *Comment) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.


### GetCanDelete

`func (o *Comment) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *Comment) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *Comment) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


