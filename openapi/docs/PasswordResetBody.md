# PasswordResetBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | the new password | 
**Id** | **string** | the workflow id provided in the reset password email | 
**Key** | **string** | the workflow key provided in the reset password email | 

## Methods

### NewPasswordResetBody

`func NewPasswordResetBody(password string, id string, key string, ) *PasswordResetBody`

NewPasswordResetBody instantiates a new PasswordResetBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordResetBodyWithDefaults

`func NewPasswordResetBodyWithDefaults() *PasswordResetBody`

NewPasswordResetBodyWithDefaults instantiates a new PasswordResetBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PasswordResetBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordResetBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordResetBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetId

`func (o *PasswordResetBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordResetBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordResetBody) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *PasswordResetBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PasswordResetBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PasswordResetBody) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


