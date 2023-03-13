# Rating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Aggregate** | [**RatingAggregate**](RatingAggregate.md) |  | 
**RatedAt** | Pointer to **time.Time** |  | [optional] 
**MyRating** | Pointer to **string** | The rating. The type is specific to the rating scheme, boolean for the likes and an integer for the fiveStar. | [optional] 

## Methods

### NewRating

`func NewRating(id string, aggregate RatingAggregate, ) *Rating`

NewRating instantiates a new Rating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingWithDefaults

`func NewRatingWithDefaults() *Rating`

NewRatingWithDefaults instantiates a new Rating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rating) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rating) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rating) SetId(v string)`

SetId sets Id field to given value.


### GetAggregate

`func (o *Rating) GetAggregate() RatingAggregate`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *Rating) GetAggregateOk() (*RatingAggregate, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *Rating) SetAggregate(v RatingAggregate)`

SetAggregate sets Aggregate field to given value.


### GetRatedAt

`func (o *Rating) GetRatedAt() time.Time`

GetRatedAt returns the RatedAt field if non-nil, zero value otherwise.

### GetRatedAtOk

`func (o *Rating) GetRatedAtOk() (*time.Time, bool)`

GetRatedAtOk returns a tuple with the RatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatedAt

`func (o *Rating) SetRatedAt(v time.Time)`

SetRatedAt sets RatedAt field to given value.

### HasRatedAt

`func (o *Rating) HasRatedAt() bool`

HasRatedAt returns a boolean if a field has been set.

### GetMyRating

`func (o *Rating) GetMyRating() string`

GetMyRating returns the MyRating field if non-nil, zero value otherwise.

### GetMyRatingOk

`func (o *Rating) GetMyRatingOk() (*string, bool)`

GetMyRatingOk returns a tuple with the MyRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyRating

`func (o *Rating) SetMyRating(v string)`

SetMyRating sets MyRating field to given value.

### HasMyRating

`func (o *Rating) HasMyRating() bool`

HasMyRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


