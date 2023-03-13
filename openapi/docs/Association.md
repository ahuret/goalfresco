# Association

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **string** |  | 
**AssocType** | **string** |  | 

## Methods

### NewAssociation

`func NewAssociation(targetId string, assocType string, ) *Association`

NewAssociation instantiates a new Association object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationWithDefaults

`func NewAssociationWithDefaults() *Association`

NewAssociationWithDefaults instantiates a new Association object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *Association) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Association) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Association) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetAssocType

`func (o *Association) GetAssocType() string`

GetAssocType returns the AssocType field if non-nil, zero value otherwise.

### GetAssocTypeOk

`func (o *Association) GetAssocTypeOk() (*string, bool)`

GetAssocTypeOk returns a tuple with the AssocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocType

`func (o *Association) SetAssocType(v string)`

SetAssocType sets AssocType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


