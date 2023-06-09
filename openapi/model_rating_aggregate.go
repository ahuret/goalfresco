/*
Alfresco Content Services REST API

**Core API**  Provides access to the core features of Alfresco Content Services. 

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RatingAggregate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatingAggregate{}

// RatingAggregate struct for RatingAggregate
type RatingAggregate struct {
	NumberOfRatings int32 `json:"numberOfRatings"`
	Average *int32 `json:"average,omitempty"`
}

// NewRatingAggregate instantiates a new RatingAggregate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatingAggregate(numberOfRatings int32) *RatingAggregate {
	this := RatingAggregate{}
	this.NumberOfRatings = numberOfRatings
	return &this
}

// NewRatingAggregateWithDefaults instantiates a new RatingAggregate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingAggregateWithDefaults() *RatingAggregate {
	this := RatingAggregate{}
	return &this
}

// GetNumberOfRatings returns the NumberOfRatings field value
func (o *RatingAggregate) GetNumberOfRatings() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfRatings
}

// GetNumberOfRatingsOk returns a tuple with the NumberOfRatings field value
// and a boolean to check if the value has been set.
func (o *RatingAggregate) GetNumberOfRatingsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfRatings, true
}

// SetNumberOfRatings sets field value
func (o *RatingAggregate) SetNumberOfRatings(v int32) {
	o.NumberOfRatings = v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *RatingAggregate) GetAverage() int32 {
	if o == nil || IsNil(o.Average) {
		var ret int32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingAggregate) GetAverageOk() (*int32, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *RatingAggregate) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given int32 and assigns it to the Average field.
func (o *RatingAggregate) SetAverage(v int32) {
	o.Average = &v
}

func (o RatingAggregate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatingAggregate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numberOfRatings"] = o.NumberOfRatings
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return toSerialize, nil
}

type NullableRatingAggregate struct {
	value *RatingAggregate
	isSet bool
}

func (v NullableRatingAggregate) Get() *RatingAggregate {
	return v.value
}

func (v *NullableRatingAggregate) Set(val *RatingAggregate) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingAggregate) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingAggregate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingAggregate(val *RatingAggregate) *NullableRatingAggregate {
	return &NullableRatingAggregate{value: val, isSet: true}
}

func (v NullableRatingAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatingAggregate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


