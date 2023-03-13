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

// checks if the SiteBodyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteBodyUpdate{}

// SiteBodyUpdate struct for SiteBodyUpdate
type SiteBodyUpdate struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// NewSiteBodyUpdate instantiates a new SiteBodyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteBodyUpdate() *SiteBodyUpdate {
	this := SiteBodyUpdate{}
	return &this
}

// NewSiteBodyUpdateWithDefaults instantiates a new SiteBodyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteBodyUpdateWithDefaults() *SiteBodyUpdate {
	this := SiteBodyUpdate{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SiteBodyUpdate) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteBodyUpdate) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SiteBodyUpdate) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SiteBodyUpdate) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteBodyUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteBodyUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteBodyUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteBodyUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *SiteBodyUpdate) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteBodyUpdate) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *SiteBodyUpdate) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *SiteBodyUpdate) SetVisibility(v string) {
	o.Visibility = &v
}

func (o SiteBodyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteBodyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableSiteBodyUpdate struct {
	value *SiteBodyUpdate
	isSet bool
}

func (v NullableSiteBodyUpdate) Get() *SiteBodyUpdate {
	return v.value
}

func (v *NullableSiteBodyUpdate) Set(val *SiteBodyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteBodyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteBodyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteBodyUpdate(val *SiteBodyUpdate) *NullableSiteBodyUpdate {
	return &NullableSiteBodyUpdate{value: val, isSet: true}
}

func (v NullableSiteBodyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteBodyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

