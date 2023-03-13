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

// checks if the CommentPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentPaging{}

// CommentPaging struct for CommentPaging
type CommentPaging struct {
	List CommentPagingList `json:"list"`
}

// NewCommentPaging instantiates a new CommentPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentPaging(list CommentPagingList) *CommentPaging {
	this := CommentPaging{}
	this.List = list
	return &this
}

// NewCommentPagingWithDefaults instantiates a new CommentPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentPagingWithDefaults() *CommentPaging {
	this := CommentPaging{}
	return &this
}

// GetList returns the List field value
func (o *CommentPaging) GetList() CommentPagingList {
	if o == nil {
		var ret CommentPagingList
		return ret
	}

	return o.List
}

// GetListOk returns a tuple with the List field value
// and a boolean to check if the value has been set.
func (o *CommentPaging) GetListOk() (*CommentPagingList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.List, true
}

// SetList sets field value
func (o *CommentPaging) SetList(v CommentPagingList) {
	o.List = v
}

func (o CommentPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list"] = o.List
	return toSerialize, nil
}

type NullableCommentPaging struct {
	value *CommentPaging
	isSet bool
}

func (v NullableCommentPaging) Get() *CommentPaging {
	return v.value
}

func (v *NullableCommentPaging) Set(val *CommentPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentPaging(val *CommentPaging) *NullableCommentPaging {
	return &NullableCommentPaging{value: val, isSet: true}
}

func (v NullableCommentPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

