# RatingBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The rating scheme type. Possible values are likes and fiveStar. | [default to "likes"]
**MyRating** | **string** | The rating. The type is specific to the rating scheme, boolean for the likes and an integer for the fiveStar | 

## Methods

### NewRatingBody

`func NewRatingBody(id string, myRating string, ) *RatingBody`

NewRatingBody instantiates a new RatingBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingBodyWithDefaults

`func NewRatingBodyWithDefaults() *RatingBody`

NewRatingBodyWithDefaults instantiates a new RatingBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RatingBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RatingBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RatingBody) SetId(v string)`

SetId sets Id field to given value.


### GetMyRating

`func (o *RatingBody) GetMyRating() string`

GetMyRating returns the MyRating field if non-nil, zero value otherwise.

### GetMyRatingOk

`func (o *RatingBody) GetMyRatingOk() (*string, bool)`

GetMyRatingOk returns a tuple with the MyRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyRating

`func (o *RatingBody) SetMyRating(v string)`

SetMyRating sets MyRating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


