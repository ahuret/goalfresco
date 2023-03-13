# PathInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]PathElement**](PathElement.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsComplete** | Pointer to **bool** |  | [optional] 

## Methods

### NewPathInfo

`func NewPathInfo() *PathInfo`

NewPathInfo instantiates a new PathInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathInfoWithDefaults

`func NewPathInfoWithDefaults() *PathInfo`

NewPathInfoWithDefaults instantiates a new PathInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *PathInfo) GetElements() []PathElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *PathInfo) GetElementsOk() (*[]PathElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *PathInfo) SetElements(v []PathElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *PathInfo) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetName

`func (o *PathInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PathInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PathInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PathInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsComplete

`func (o *PathInfo) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *PathInfo) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *PathInfo) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *PathInfo) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


