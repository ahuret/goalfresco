# PersonBodyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**SkypeId** | Pointer to **string** |  | [optional] 
**GoogleId** | Pointer to **string** |  | [optional] 
**InstantMessageId** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**Mobile** | Pointer to **string** |  | [optional] 
**Telephone** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EmailNotificationsEnabled** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**OldPassword** | Pointer to **string** |  | [optional] 
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPersonBodyUpdate

`func NewPersonBodyUpdate() *PersonBodyUpdate`

NewPersonBodyUpdate instantiates a new PersonBodyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonBodyUpdateWithDefaults

`func NewPersonBodyUpdateWithDefaults() *PersonBodyUpdate`

NewPersonBodyUpdateWithDefaults instantiates a new PersonBodyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *PersonBodyUpdate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PersonBodyUpdate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PersonBodyUpdate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PersonBodyUpdate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PersonBodyUpdate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PersonBodyUpdate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PersonBodyUpdate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PersonBodyUpdate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDescription

`func (o *PersonBodyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PersonBodyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PersonBodyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PersonBodyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *PersonBodyUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PersonBodyUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PersonBodyUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PersonBodyUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSkypeId

`func (o *PersonBodyUpdate) GetSkypeId() string`

GetSkypeId returns the SkypeId field if non-nil, zero value otherwise.

### GetSkypeIdOk

`func (o *PersonBodyUpdate) GetSkypeIdOk() (*string, bool)`

GetSkypeIdOk returns a tuple with the SkypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkypeId

`func (o *PersonBodyUpdate) SetSkypeId(v string)`

SetSkypeId sets SkypeId field to given value.

### HasSkypeId

`func (o *PersonBodyUpdate) HasSkypeId() bool`

HasSkypeId returns a boolean if a field has been set.

### GetGoogleId

`func (o *PersonBodyUpdate) GetGoogleId() string`

GetGoogleId returns the GoogleId field if non-nil, zero value otherwise.

### GetGoogleIdOk

`func (o *PersonBodyUpdate) GetGoogleIdOk() (*string, bool)`

GetGoogleIdOk returns a tuple with the GoogleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleId

`func (o *PersonBodyUpdate) SetGoogleId(v string)`

SetGoogleId sets GoogleId field to given value.

### HasGoogleId

`func (o *PersonBodyUpdate) HasGoogleId() bool`

HasGoogleId returns a boolean if a field has been set.

### GetInstantMessageId

`func (o *PersonBodyUpdate) GetInstantMessageId() string`

GetInstantMessageId returns the InstantMessageId field if non-nil, zero value otherwise.

### GetInstantMessageIdOk

`func (o *PersonBodyUpdate) GetInstantMessageIdOk() (*string, bool)`

GetInstantMessageIdOk returns a tuple with the InstantMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantMessageId

`func (o *PersonBodyUpdate) SetInstantMessageId(v string)`

SetInstantMessageId sets InstantMessageId field to given value.

### HasInstantMessageId

`func (o *PersonBodyUpdate) HasInstantMessageId() bool`

HasInstantMessageId returns a boolean if a field has been set.

### GetJobTitle

`func (o *PersonBodyUpdate) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *PersonBodyUpdate) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *PersonBodyUpdate) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *PersonBodyUpdate) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLocation

`func (o *PersonBodyUpdate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PersonBodyUpdate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PersonBodyUpdate) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PersonBodyUpdate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCompany

`func (o *PersonBodyUpdate) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PersonBodyUpdate) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PersonBodyUpdate) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PersonBodyUpdate) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetMobile

`func (o *PersonBodyUpdate) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *PersonBodyUpdate) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *PersonBodyUpdate) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *PersonBodyUpdate) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetTelephone

`func (o *PersonBodyUpdate) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *PersonBodyUpdate) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *PersonBodyUpdate) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *PersonBodyUpdate) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetUserStatus

`func (o *PersonBodyUpdate) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *PersonBodyUpdate) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *PersonBodyUpdate) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *PersonBodyUpdate) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *PersonBodyUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PersonBodyUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PersonBodyUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PersonBodyUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEmailNotificationsEnabled

`func (o *PersonBodyUpdate) GetEmailNotificationsEnabled() bool`

GetEmailNotificationsEnabled returns the EmailNotificationsEnabled field if non-nil, zero value otherwise.

### GetEmailNotificationsEnabledOk

`func (o *PersonBodyUpdate) GetEmailNotificationsEnabledOk() (*bool, bool)`

GetEmailNotificationsEnabledOk returns a tuple with the EmailNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationsEnabled

`func (o *PersonBodyUpdate) SetEmailNotificationsEnabled(v bool)`

SetEmailNotificationsEnabled sets EmailNotificationsEnabled field to given value.

### HasEmailNotificationsEnabled

`func (o *PersonBodyUpdate) HasEmailNotificationsEnabled() bool`

HasEmailNotificationsEnabled returns a boolean if a field has been set.

### GetPassword

`func (o *PersonBodyUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PersonBodyUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PersonBodyUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PersonBodyUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOldPassword

`func (o *PersonBodyUpdate) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *PersonBodyUpdate) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *PersonBodyUpdate) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *PersonBodyUpdate) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetAspectNames

`func (o *PersonBodyUpdate) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *PersonBodyUpdate) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *PersonBodyUpdate) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *PersonBodyUpdate) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *PersonBodyUpdate) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PersonBodyUpdate) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PersonBodyUpdate) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PersonBodyUpdate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


