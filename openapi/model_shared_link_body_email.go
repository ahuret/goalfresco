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

// checks if the SharedLinkBodyEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedLinkBodyEmail{}

// SharedLinkBodyEmail struct for SharedLinkBodyEmail
type SharedLinkBodyEmail struct {
	Client *string `json:"client,omitempty"`
	Message *string `json:"message,omitempty"`
	Locale *string `json:"locale,omitempty"`
	RecipientEmails []string `json:"recipientEmails,omitempty"`
}

// NewSharedLinkBodyEmail instantiates a new SharedLinkBodyEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedLinkBodyEmail() *SharedLinkBodyEmail {
	this := SharedLinkBodyEmail{}
	return &this
}

// NewSharedLinkBodyEmailWithDefaults instantiates a new SharedLinkBodyEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedLinkBodyEmailWithDefaults() *SharedLinkBodyEmail {
	this := SharedLinkBodyEmail{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *SharedLinkBodyEmail) GetClient() string {
	if o == nil || IsNil(o.Client) {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedLinkBodyEmail) GetClientOk() (*string, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *SharedLinkBodyEmail) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *SharedLinkBodyEmail) SetClient(v string) {
	o.Client = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SharedLinkBodyEmail) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedLinkBodyEmail) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SharedLinkBodyEmail) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SharedLinkBodyEmail) SetMessage(v string) {
	o.Message = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *SharedLinkBodyEmail) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedLinkBodyEmail) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *SharedLinkBodyEmail) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *SharedLinkBodyEmail) SetLocale(v string) {
	o.Locale = &v
}

// GetRecipientEmails returns the RecipientEmails field value if set, zero value otherwise.
func (o *SharedLinkBodyEmail) GetRecipientEmails() []string {
	if o == nil || IsNil(o.RecipientEmails) {
		var ret []string
		return ret
	}
	return o.RecipientEmails
}

// GetRecipientEmailsOk returns a tuple with the RecipientEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedLinkBodyEmail) GetRecipientEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.RecipientEmails) {
		return nil, false
	}
	return o.RecipientEmails, true
}

// HasRecipientEmails returns a boolean if a field has been set.
func (o *SharedLinkBodyEmail) HasRecipientEmails() bool {
	if o != nil && !IsNil(o.RecipientEmails) {
		return true
	}

	return false
}

// SetRecipientEmails gets a reference to the given []string and assigns it to the RecipientEmails field.
func (o *SharedLinkBodyEmail) SetRecipientEmails(v []string) {
	o.RecipientEmails = v
}

func (o SharedLinkBodyEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedLinkBodyEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.RecipientEmails) {
		toSerialize["recipientEmails"] = o.RecipientEmails
	}
	return toSerialize, nil
}

type NullableSharedLinkBodyEmail struct {
	value *SharedLinkBodyEmail
	isSet bool
}

func (v NullableSharedLinkBodyEmail) Get() *SharedLinkBodyEmail {
	return v.value
}

func (v *NullableSharedLinkBodyEmail) Set(val *SharedLinkBodyEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedLinkBodyEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedLinkBodyEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedLinkBodyEmail(val *SharedLinkBodyEmail) *NullableSharedLinkBodyEmail {
	return &NullableSharedLinkBodyEmail{value: val, isSet: true}
}

func (v NullableSharedLinkBodyEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedLinkBodyEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

