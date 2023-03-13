# RatingAggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfRatings** | **int32** |  | 
**Average** | Pointer to **int32** |  | [optional] 

## Methods

### NewRatingAggregate

`func NewRatingAggregate(numberOfRatings int32, ) *RatingAggregate`

NewRatingAggregate instantiates a new RatingAggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingAggregateWithDefaults

`func NewRatingAggregateWithDefaults() *RatingAggregate`

NewRatingAggregateWithDefaults instantiates a new RatingAggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfRatings

`func (o *RatingAggregate) GetNumberOfRatings() int32`

GetNumberOfRatings returns the NumberOfRatings field if non-nil, zero value otherwise.

### GetNumberOfRatingsOk

`func (o *RatingAggregate) GetNumberOfRatingsOk() (*int32, bool)`

GetNumberOfRatingsOk returns a tuple with the NumberOfRatings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRatings

`func (o *RatingAggregate) SetNumberOfRatings(v int32)`

SetNumberOfRatings sets NumberOfRatings field to given value.


### GetAverage

`func (o *RatingAggregate) GetAverage() int32`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *RatingAggregate) GetAverageOk() (*int32, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *RatingAggregate) SetAverage(v int32)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *RatingAggregate) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


