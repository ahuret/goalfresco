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

// checks if the SiteMembershipRejectionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteMembershipRejectionBody{}

// SiteMembershipRejectionBody struct for SiteMembershipRejectionBody
type SiteMembershipRejectionBody struct {
	Comment *string `json:"comment,omitempty"`
}

// NewSiteMembershipRejectionBody instantiates a new SiteMembershipRejectionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteMembershipRejectionBody() *SiteMembershipRejectionBody {
	this := SiteMembershipRejectionBody{}
	return &this
}

// NewSiteMembershipRejectionBodyWithDefaults instantiates a new SiteMembershipRejectionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteMembershipRejectionBodyWithDefaults() *SiteMembershipRejectionBody {
	this := SiteMembershipRejectionBody{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SiteMembershipRejectionBody) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteMembershipRejectionBody) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SiteMembershipRejectionBody) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SiteMembershipRejectionBody) SetComment(v string) {
	o.Comment = &v
}

func (o SiteMembershipRejectionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteMembershipRejectionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableSiteMembershipRejectionBody struct {
	value *SiteMembershipRejectionBody
	isSet bool
}

func (v NullableSiteMembershipRejectionBody) Get() *SiteMembershipRejectionBody {
	return v.value
}

func (v *NullableSiteMembershipRejectionBody) Set(val *SiteMembershipRejectionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMembershipRejectionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMembershipRejectionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMembershipRejectionBody(val *SiteMembershipRejectionBody) *NullableSiteMembershipRejectionBody {
	return &NullableSiteMembershipRejectionBody{value: val, isSet: true}
}

func (v NullableSiteMembershipRejectionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMembershipRejectionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


