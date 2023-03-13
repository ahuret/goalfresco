# SharedLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name must not contain spaces or the following special characters: * \&quot; &lt; &gt; \\ / ? : and |. The character . must not be used at the end of the name.  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedByUser** | Pointer to [**UserInfo**](UserInfo.md) |  | [optional] 
**SharedByUser** | Pointer to [**UserInfo**](UserInfo.md) |  | [optional] 
**Content** | Pointer to [**ContentInfo**](ContentInfo.md) |  | [optional] 
**AllowableOperations** | Pointer to **[]string** | The allowable operations for the Quickshare link itself. See allowableOperationsOnTarget for the allowable operations pertaining to the linked content node.  | [optional] 
**AllowableOperationsOnTarget** | Pointer to **[]string** | The allowable operations for the content node being shared.  | [optional] 
**IsFavorite** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | A subset of the target node&#39;s properties, system properties and properties already available in the SharedLink are excluded.  | [optional] 
**AspectNames** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to [**PathInfo**](PathInfo.md) |  | [optional] 

## Methods

### NewSharedLink

`func NewSharedLink() *SharedLink`

NewSharedLink instantiates a new SharedLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedLinkWithDefaults

`func NewSharedLinkWithDefaults() *SharedLink`

NewSharedLinkWithDefaults instantiates a new SharedLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SharedLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SharedLink) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SharedLink) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SharedLink) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SharedLink) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetNodeId

`func (o *SharedLink) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SharedLink) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SharedLink) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *SharedLink) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetName

`func (o *SharedLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *SharedLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SharedLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SharedLink) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SharedLink) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *SharedLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SharedLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SharedLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SharedLink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SharedLink) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SharedLink) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SharedLink) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SharedLink) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedByUser

`func (o *SharedLink) GetModifiedByUser() UserInfo`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *SharedLink) GetModifiedByUserOk() (*UserInfo, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *SharedLink) SetModifiedByUser(v UserInfo)`

SetModifiedByUser sets ModifiedByUser field to given value.

### HasModifiedByUser

`func (o *SharedLink) HasModifiedByUser() bool`

HasModifiedByUser returns a boolean if a field has been set.

### GetSharedByUser

`func (o *SharedLink) GetSharedByUser() UserInfo`

GetSharedByUser returns the SharedByUser field if non-nil, zero value otherwise.

### GetSharedByUserOk

`func (o *SharedLink) GetSharedByUserOk() (*UserInfo, bool)`

GetSharedByUserOk returns a tuple with the SharedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedByUser

`func (o *SharedLink) SetSharedByUser(v UserInfo)`

SetSharedByUser sets SharedByUser field to given value.

### HasSharedByUser

`func (o *SharedLink) HasSharedByUser() bool`

HasSharedByUser returns a boolean if a field has been set.

### GetContent

`func (o *SharedLink) GetContent() ContentInfo`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SharedLink) GetContentOk() (*ContentInfo, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SharedLink) SetContent(v ContentInfo)`

SetContent sets Content field to given value.

### HasContent

`func (o *SharedLink) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAllowableOperations

`func (o *SharedLink) GetAllowableOperations() []string`

GetAllowableOperations returns the AllowableOperations field if non-nil, zero value otherwise.

### GetAllowableOperationsOk

`func (o *SharedLink) GetAllowableOperationsOk() (*[]string, bool)`

GetAllowableOperationsOk returns a tuple with the AllowableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableOperations

`func (o *SharedLink) SetAllowableOperations(v []string)`

SetAllowableOperations sets AllowableOperations field to given value.

### HasAllowableOperations

`func (o *SharedLink) HasAllowableOperations() bool`

HasAllowableOperations returns a boolean if a field has been set.

### GetAllowableOperationsOnTarget

`func (o *SharedLink) GetAllowableOperationsOnTarget() []string`

GetAllowableOperationsOnTarget returns the AllowableOperationsOnTarget field if non-nil, zero value otherwise.

### GetAllowableOperationsOnTargetOk

`func (o *SharedLink) GetAllowableOperationsOnTargetOk() (*[]string, bool)`

GetAllowableOperationsOnTargetOk returns a tuple with the AllowableOperationsOnTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableOperationsOnTarget

`func (o *SharedLink) SetAllowableOperationsOnTarget(v []string)`

SetAllowableOperationsOnTarget sets AllowableOperationsOnTarget field to given value.

### HasAllowableOperationsOnTarget

`func (o *SharedLink) HasAllowableOperationsOnTarget() bool`

HasAllowableOperationsOnTarget returns a boolean if a field has been set.

### GetIsFavorite

`func (o *SharedLink) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *SharedLink) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *SharedLink) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *SharedLink) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetProperties

`func (o *SharedLink) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SharedLink) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SharedLink) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SharedLink) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAspectNames

`func (o *SharedLink) GetAspectNames() []string`

GetAspectNames returns the AspectNames field if non-nil, zero value otherwise.

### GetAspectNamesOk

`func (o *SharedLink) GetAspectNamesOk() (*[]string, bool)`

GetAspectNamesOk returns a tuple with the AspectNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectNames

`func (o *SharedLink) SetAspectNames(v []string)`

SetAspectNames sets AspectNames field to given value.

### HasAspectNames

`func (o *SharedLink) HasAspectNames() bool`

HasAspectNames returns a boolean if a field has been set.

### GetPath

`func (o *SharedLink) GetPath() PathInfo`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SharedLink) GetPathOk() (*PathInfo, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SharedLink) SetPath(v PathInfo)`

SetPath sets Path field to given value.

### HasPath

`func (o *SharedLink) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


