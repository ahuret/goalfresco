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

// checks if the Pagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pagination{}

// Pagination struct for Pagination
type Pagination struct {
	// The number of objects in the entries array. 
	Count *int64 `json:"count,omitempty"`
	// A boolean value which is **true** if there are more entities in the collection beyond those in this response. A true value means a request with a larger value for the **skipCount** or the **maxItems** parameter will return more entities. 
	HasMoreItems *bool `json:"hasMoreItems,omitempty"`
	// An integer describing the total number of entities in the collection. The API might not be able to determine this value, in which case this property will not be present. 
	TotalItems *int64 `json:"totalItems,omitempty"`
	// An integer describing how many entities exist in the collection before those included in this list. If there was no **skipCount** parameter then the default value is 0. 
	SkipCount *int64 `json:"skipCount,omitempty"`
	// The value of the **maxItems** parameter used to generate this list. If there was no **maxItems** parameter then the default value is 100. 
	MaxItems *int64 `json:"maxItems,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Pagination) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Pagination) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *Pagination) SetCount(v int64) {
	o.Count = &v
}

// GetHasMoreItems returns the HasMoreItems field value if set, zero value otherwise.
func (o *Pagination) GetHasMoreItems() bool {
	if o == nil || IsNil(o.HasMoreItems) {
		var ret bool
		return ret
	}
	return *o.HasMoreItems
}

// GetHasMoreItemsOk returns a tuple with the HasMoreItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetHasMoreItemsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMoreItems) {
		return nil, false
	}
	return o.HasMoreItems, true
}

// HasHasMoreItems returns a boolean if a field has been set.
func (o *Pagination) HasHasMoreItems() bool {
	if o != nil && !IsNil(o.HasMoreItems) {
		return true
	}

	return false
}

// SetHasMoreItems gets a reference to the given bool and assigns it to the HasMoreItems field.
func (o *Pagination) SetHasMoreItems(v bool) {
	o.HasMoreItems = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *Pagination) GetTotalItems() int64 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int64
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalItemsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *Pagination) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int64 and assigns it to the TotalItems field.
func (o *Pagination) SetTotalItems(v int64) {
	o.TotalItems = &v
}

// GetSkipCount returns the SkipCount field value if set, zero value otherwise.
func (o *Pagination) GetSkipCount() int64 {
	if o == nil || IsNil(o.SkipCount) {
		var ret int64
		return ret
	}
	return *o.SkipCount
}

// GetSkipCountOk returns a tuple with the SkipCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetSkipCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SkipCount) {
		return nil, false
	}
	return o.SkipCount, true
}

// HasSkipCount returns a boolean if a field has been set.
func (o *Pagination) HasSkipCount() bool {
	if o != nil && !IsNil(o.SkipCount) {
		return true
	}

	return false
}

// SetSkipCount gets a reference to the given int64 and assigns it to the SkipCount field.
func (o *Pagination) SetSkipCount(v int64) {
	o.SkipCount = &v
}

// GetMaxItems returns the MaxItems field value if set, zero value otherwise.
func (o *Pagination) GetMaxItems() int64 {
	if o == nil || IsNil(o.MaxItems) {
		var ret int64
		return ret
	}
	return *o.MaxItems
}

// GetMaxItemsOk returns a tuple with the MaxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetMaxItemsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxItems) {
		return nil, false
	}
	return o.MaxItems, true
}

// HasMaxItems returns a boolean if a field has been set.
func (o *Pagination) HasMaxItems() bool {
	if o != nil && !IsNil(o.MaxItems) {
		return true
	}

	return false
}

// SetMaxItems gets a reference to the given int64 and assigns it to the MaxItems field.
func (o *Pagination) SetMaxItems(v int64) {
	o.MaxItems = &v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.HasMoreItems) {
		toSerialize["hasMoreItems"] = o.HasMoreItems
	}
	if !IsNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	if !IsNil(o.SkipCount) {
		toSerialize["skipCount"] = o.SkipCount
	}
	if !IsNil(o.MaxItems) {
		toSerialize["maxItems"] = o.MaxItems
	}
	return toSerialize, nil
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

