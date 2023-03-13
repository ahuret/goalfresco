# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AvatarId** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**SkypeId** | Pointer to **string** |  | [optional] 
**GoogleId** | Pointer to **string** |  | [optional] 
**InstantMessageId** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**Mobile** | Pointer to **string** |  | [optional] 
**Telephone** | Pointer to **string** |  | [optional] 
**StatusUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | [default to true]
**EmailNotificationsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Capabilities** | Pointer to [**Capabilities**](Capabilities.md) |  | [optional] 

## Methods

### NewPerson

`func NewPerson(id string, firstName string, email string, enabled bool, ) *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Person) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Person) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Person) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *Person) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Person) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Person) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Person) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Person) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Person) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Person) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDisplayName

`func (o *Person) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Person) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Person) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Person) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *Person) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Person) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Person) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Person) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvatarId

`func (o *Person) GetAvatarId() string`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *Person) GetAvatarIdOk() (*string, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *Person) SetAvatarId(v string)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *Person) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetEmail

`func (o *Person) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Person) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Person) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetSkypeId

`func (o *Person) GetSkypeId() string`

GetSkypeId returns the SkypeId field if non-nil, zero value otherwise.

### GetSkypeIdOk

`func (o *Person) GetSkypeIdOk() (*string, bool)`

GetSkypeIdOk returns a tuple with the SkypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkypeId

`func (o *Person) SetSkypeId(v string)`

SetSkypeId sets SkypeId field to given value.

### HasSkypeId

`func (o *Person) HasSkypeId() bool`

HasSkypeId returns a boolean if a field has been set.

### GetGoogleId

`func (o *Person) GetGoogleId() string`

GetGoogleId returns the GoogleId field if non-nil, zero value otherwise.

### GetGoogleIdOk

`func (o *Person) GetGoogleIdOk() (*string, bool)`

GetGoogleIdOk returns a tuple with the GoogleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleId

`func (o *Person) SetGoogleId(v string)`

SetGoogleId sets GoogleId field to given value.

### HasGoogleId

`func (o *Person) HasGoogleId() bool`

HasGoogleId returns a boolean if a field has been set.

### GetInstantMessageId

`func (o *Person) GetInstantMessageId() string`

GetInstantMessageId returns the InstantMessageId field if non-nil, zero value otherwise.

### GetInstantMessageIdOk

`func (o *Person) GetInstantMessageIdOk() (*string, bool)`

GetInstantMessageIdOk returns a tuple with the InstantMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantMessageId

`func (o *Person) SetInstantMessageId(v string)`

SetInstantMessageId sets InstantMessageId field to given value.

### HasInstantMessageId

`func (o *Person) HasInstantMessageId() bool`

HasInstantMessageId returns a boolean if a field has been set.

### GetJobTitle

`func (o *Person) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Person) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Person) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Person) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLocation

`func (o *Person) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Person) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Person) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Person) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCompany

`func (o *Person) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Person) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Person) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Person) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetMobile

`func (o *Person) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *Person) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *Person) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *Person) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetTelephone

`func (o *Person) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *Person) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *Person) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *Person) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *Person) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *Person) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *Person) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *Person) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetUserStatus

`func (o *Person) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *Person) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *Person) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *Person) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *Person) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Person) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Person) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEmailNotificationsEnabled

`func (o *Person) GetEmailNotificationsEnabled() bool`

GetEmailNotificationsEnabled returns the EmailNotificationsEnabled field if non-nil, zero value otherwise.

### GetEmailNotificationsEnabledOk

`func (o *Person) GetEmailNotificationsEnabledOk() (*bool, bool)`

GetEmailNotificationsEnabledOk returns a tuple with the EmailNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationsEnabled

`func (o *Person) SetEmailNotificationsEnabled(v bool)`

SetEmailNotificationsEnabled sets EmailNotificationsEnabled field to given value.

### HasEmailNotificationsEnabled

`func (o *Person) HasEmailNotificationsEnabled() bool`

HasEmailNotificationsEnabled returns a boolean if a field has been set.

### GetAspectNames

`func (o *Person) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *Person) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *Person) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *Person) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetProperties

`func (o *Person) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Person) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Person) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Person) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCapabilities

`func (o *Person) GetCapabilities() Capabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Person) GetCapabilitiesOk() (*Capabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Person) SetCapabilities(v Capabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Person) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


