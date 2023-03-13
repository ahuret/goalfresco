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

// checks if the ErrorError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorError{}

// ErrorError struct for ErrorError
type ErrorError struct {
	ErrorKey *string `json:"errorKey,omitempty"`
	StatusCode int32 `json:"statusCode"`
	BriefSummary string `json:"briefSummary"`
	StackTrace string `json:"stackTrace"`
	DescriptionURL string `json:"descriptionURL"`
	LogId *string `json:"logId,omitempty"`
}

// NewErrorError instantiates a new ErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorError(statusCode int32, briefSummary string, stackTrace string, descriptionURL string) *ErrorError {
	this := ErrorError{}
	this.StatusCode = statusCode
	this.BriefSummary = briefSummary
	this.StackTrace = stackTrace
	this.DescriptionURL = descriptionURL
	return &this
}

// NewErrorErrorWithDefaults instantiates a new ErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorErrorWithDefaults() *ErrorError {
	this := ErrorError{}
	return &this
}

// GetErrorKey returns the ErrorKey field value if set, zero value otherwise.
func (o *ErrorError) GetErrorKey() string {
	if o == nil || IsNil(o.ErrorKey) {
		var ret string
		return ret
	}
	return *o.ErrorKey
}

// GetErrorKeyOk returns a tuple with the ErrorKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorError) GetErrorKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorKey) {
		return nil, false
	}
	return o.ErrorKey, true
}

// HasErrorKey returns a boolean if a field has been set.
func (o *ErrorError) HasErrorKey() bool {
	if o != nil && !IsNil(o.ErrorKey) {
		return true
	}

	return false
}

// SetErrorKey gets a reference to the given string and assigns it to the ErrorKey field.
func (o *ErrorError) SetErrorKey(v string) {
	o.ErrorKey = &v
}

// GetStatusCode returns the StatusCode field value
func (o *ErrorError) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *ErrorError) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetBriefSummary returns the BriefSummary field value
func (o *ErrorError) GetBriefSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BriefSummary
}

// GetBriefSummaryOk returns a tuple with the BriefSummary field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetBriefSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BriefSummary, true
}

// SetBriefSummary sets field value
func (o *ErrorError) SetBriefSummary(v string) {
	o.BriefSummary = v
}

// GetStackTrace returns the StackTrace field value
func (o *ErrorError) GetStackTrace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetStackTraceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackTrace, true
}

// SetStackTrace sets field value
func (o *ErrorError) SetStackTrace(v string) {
	o.StackTrace = v
}

// GetDescriptionURL returns the DescriptionURL field value
func (o *ErrorError) GetDescriptionURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DescriptionURL
}

// GetDescriptionURLOk returns a tuple with the DescriptionURL field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetDescriptionURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptionURL, true
}

// SetDescriptionURL sets field value
func (o *ErrorError) SetDescriptionURL(v string) {
	o.DescriptionURL = v
}

// GetLogId returns the LogId field value if set, zero value otherwise.
func (o *ErrorError) GetLogId() string {
	if o == nil || IsNil(o.LogId) {
		var ret string
		return ret
	}
	return *o.LogId
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorError) GetLogIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogId) {
		return nil, false
	}
	return o.LogId, true
}

// HasLogId returns a boolean if a field has been set.
func (o *ErrorError) HasLogId() bool {
	if o != nil && !IsNil(o.LogId) {
		return true
	}

	return false
}

// SetLogId gets a reference to the given string and assigns it to the LogId field.
func (o *ErrorError) SetLogId(v string) {
	o.LogId = &v
}

func (o ErrorError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorKey) {
		toSerialize["errorKey"] = o.ErrorKey
	}
	toSerialize["statusCode"] = o.StatusCode
	toSerialize["briefSummary"] = o.BriefSummary
	toSerialize["stackTrace"] = o.StackTrace
	toSerialize["descriptionURL"] = o.DescriptionURL
	if !IsNil(o.LogId) {
		toSerialize["logId"] = o.LogId
	}
	return toSerialize, nil
}

type NullableErrorError struct {
	value *ErrorError
	isSet bool
}

func (v NullableErrorError) Get() *ErrorError {
	return v.value
}

func (v *NullableErrorError) Set(val *ErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorError(val *ErrorError) *NullableErrorError {
	return &NullableErrorError{value: val, isSet: true}
}

func (v NullableErrorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


