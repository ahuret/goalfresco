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

// checks if the CommentBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentBody{}

// CommentBody struct for CommentBody
type CommentBody struct {
	Content string `json:"content"`
}

// NewCommentBody instantiates a new CommentBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentBody(content string) *CommentBody {
	this := CommentBody{}
	this.Content = content
	return &this
}

// NewCommentBodyWithDefaults instantiates a new CommentBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentBodyWithDefaults() *CommentBody {
	this := CommentBody{}
	return &this
}

// GetContent returns the Content field value
func (o *CommentBody) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CommentBody) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CommentBody) SetContent(v string) {
	o.Content = v
}

func (o CommentBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

type NullableCommentBody struct {
	value *CommentBody
	isSet bool
}

func (v NullableCommentBody) Get() *CommentBody {
	return v.value
}

func (v *NullableCommentBody) Set(val *CommentBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentBody(val *CommentBody) *NullableCommentBody {
	return &NullableCommentBody{value: val, isSet: true}
}

func (v NullableCommentBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


