# PermissionElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorityId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccessStatus** | Pointer to **string** |  | [optional] [default to "ALLOWED"]

## Methods

### NewPermissionElement

`func NewPermissionElement() *PermissionElement`

NewPermissionElement instantiates a new PermissionElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionElementWithDefaults

`func NewPermissionElementWithDefaults() *PermissionElement`

NewPermissionElementWithDefaults instantiates a new PermissionElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorityId

`func (o *PermissionElement) GetAuthorityId() string`

GetAuthorityId returns the AuthorityId field if non-nil, zero value otherwise.

### GetAuthorityIdOk

`func (o *PermissionElement) GetAuthorityIdOk() (*string, bool)`

GetAuthorityIdOk returns a tuple with the AuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityId

`func (o *PermissionElement) SetAuthorityId(v string)`

SetAuthorityId sets AuthorityId field to given value.

### HasAuthorityId

`func (o *PermissionElement) HasAuthorityId() bool`

HasAuthorityId returns a boolean if a field has been set.

### GetName

`func (o *PermissionElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccessStatus

`func (o *PermissionElement) GetAccessStatus() string`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *PermissionElement) GetAccessStatusOk() (*string, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *PermissionElement) SetAccessStatus(v string)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *PermissionElement) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


