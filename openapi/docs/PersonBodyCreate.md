# PersonBodyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**SkypeId** | Pointer to **string** |  | [optional] 
**GoogleId** | Pointer to **string** |  | [optional] 
**InstantMessageId** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**Mobile** | Pointer to **string** |  | [optional] 
**Telephone** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**EmailNotificationsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Password** | **string** |  | 
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPersonBodyCreate

`func NewPersonBodyCreate(id string, firstName string, email string, password string, ) *PersonBodyCreate`

NewPersonBodyCreate instantiates a new PersonBodyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonBodyCreateWithDefaults

`func NewPersonBodyCreateWithDefaults() *PersonBodyCreate`

NewPersonBodyCreateWithDefaults instantiates a new PersonBodyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonBodyCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonBodyCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonBodyCreate) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *PersonBodyCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PersonBodyCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PersonBodyCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *PersonBodyCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PersonBodyCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PersonBodyCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PersonBodyCreate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDescription

`func (o *PersonBodyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PersonBodyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PersonBodyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PersonBodyCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *PersonBodyCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PersonBodyCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PersonBodyCreate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetSkypeId

`func (o *PersonBodyCreate) GetSkypeId() string`

GetSkypeId returns the SkypeId field if non-nil, zero value otherwise.

### GetSkypeIdOk

`func (o *PersonBodyCreate) GetSkypeIdOk() (*string, bool)`

GetSkypeIdOk returns a tuple with the SkypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkypeId

`func (o *PersonBodyCreate) SetSkypeId(v string)`

SetSkypeId sets SkypeId field to given value.

### HasSkypeId

`func (o *PersonBodyCreate) HasSkypeId() bool`

HasSkypeId returns a boolean if a field has been set.

### GetGoogleId

`func (o *PersonBodyCreate) GetGoogleId() string`

GetGoogleId returns the GoogleId field if non-nil, zero value otherwise.

### GetGoogleIdOk

`func (o *PersonBodyCreate) GetGoogleIdOk() (*string, bool)`

GetGoogleIdOk returns a tuple with the GoogleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleId

`func (o *PersonBodyCreate) SetGoogleId(v string)`

SetGoogleId sets GoogleId field to given value.

### HasGoogleId

`func (o *PersonBodyCreate) HasGoogleId() bool`

HasGoogleId returns a boolean if a field has been set.

### GetInstantMessageId

`func (o *PersonBodyCreate) GetInstantMessageId() string`

GetInstantMessageId returns the InstantMessageId field if non-nil, zero value otherwise.

### GetInstantMessageIdOk

`func (o *PersonBodyCreate) GetInstantMessageIdOk() (*string, bool)`

GetInstantMessageIdOk returns a tuple with the InstantMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantMessageId

`func (o *PersonBodyCreate) SetInstantMessageId(v string)`

SetInstantMessageId sets InstantMessageId field to given value.

### HasInstantMessageId

`func (o *PersonBodyCreate) HasInstantMessageId() bool`

HasInstantMessageId returns a boolean if a field has been set.

### GetJobTitle

`func (o *PersonBodyCreate) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *PersonBodyCreate) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *PersonBodyCreate) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *PersonBodyCreate) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLocation

`func (o *PersonBodyCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PersonBodyCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PersonBodyCreate) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PersonBodyCreate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCompany

`func (o *PersonBodyCreate) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PersonBodyCreate) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PersonBodyCreate) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PersonBodyCreate) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetMobile

`func (o *PersonBodyCreate) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *PersonBodyCreate) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *PersonBodyCreate) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *PersonBodyCreate) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetTelephone

`func (o *PersonBodyCreate) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *PersonBodyCreate) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *PersonBodyCreate) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *PersonBodyCreate) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetUserStatus

`func (o *PersonBodyCreate) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *PersonBodyCreate) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *PersonBodyCreate) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *PersonBodyCreate) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *PersonBodyCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PersonBodyCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PersonBodyCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PersonBodyCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEmailNotificationsEnabled

`func (o *PersonBodyCreate) GetEmailNotificationsEnabled() bool`

GetEmailNotificationsEnabled returns the EmailNotificationsEnabled field if non-nil, zero value otherwise.

### GetEmailNotificationsEnabledOk

`func (o *PersonBodyCreate) GetEmailNotificationsEnabledOk() (*bool, bool)`

GetEmailNotificationsEnabledOk returns a tuple with the EmailNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationsEnabled

`func (o *PersonBodyCreate) SetEmailNotificationsEnabled(v bool)`

SetEmailNotificationsEnabled sets EmailNotificationsEnabled field to given value.

### HasEmailNotificationsEnabled

`func (o *PersonBodyCreate) HasEmailNotificationsEnabled() bool`

HasEmailNotificationsEnabled returns a boolean if a field has been set.

### GetPassword

`func (o *PersonBodyCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PersonBodyCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PersonBodyCreate) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAspectNames

`func (o *PersonBodyCreate) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *PersonBodyCreate) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *PersonBodyCreate) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *PersonBodyCreate) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *PersonBodyCreate) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PersonBodyCreate) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PersonBodyCreate) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PersonBodyCreate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


