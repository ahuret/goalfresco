# SiteMembershipRequestBodyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** | Optional client name used when sending an email to the end user, defaults to \&quot;share\&quot; if not provided. **Note:** The client must be registered before this API can send an email. **Note:** This is available in Alfresco 7.0.0 and newer versions.  | [optional] 

## Methods

### NewSiteMembershipRequestBodyCreate

`func NewSiteMembershipRequestBodyCreate(id string, ) *SiteMembershipRequestBodyCreate`

NewSiteMembershipRequestBodyCreate instantiates a new SiteMembershipRequestBodyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteMembershipRequestBodyCreateWithDefaults

`func NewSiteMembershipRequestBodyCreateWithDefaults() *SiteMembershipRequestBodyCreate`

NewSiteMembershipRequestBodyCreateWithDefaults instantiates a new SiteMembershipRequestBodyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SiteMembershipRequestBodyCreate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SiteMembershipRequestBodyCreate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SiteMembershipRequestBodyCreate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SiteMembershipRequestBodyCreate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetId

`func (o *SiteMembershipRequestBodyCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteMembershipRequestBodyCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteMembershipRequestBodyCreate) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *SiteMembershipRequestBodyCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SiteMembershipRequestBodyCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SiteMembershipRequestBodyCreate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SiteMembershipRequestBodyCreate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetClient

`func (o *SiteMembershipRequestBodyCreate) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SiteMembershipRequestBodyCreate) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SiteMembershipRequestBodyCreate) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *SiteMembershipRequestBodyCreate) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


