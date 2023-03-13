# ErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorKey** | Pointer to **string** |  | [optional] 
**StatusCode** | **int32** |  | 
**BriefSummary** | **string** |  | 
**StackTrace** | **string** |  | 
**DescriptionURL** | **string** |  | 
**LogId** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorError

`func NewErrorError(statusCode int32, briefSummary string, stackTrace string, descriptionURL string, ) *ErrorError`

NewErrorError instantiates a new ErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorErrorWithDefaults

`func NewErrorErrorWithDefaults() *ErrorError`

NewErrorErrorWithDefaults instantiates a new ErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorKey

`func (o *ErrorError) GetErrorKey() string`

GetErrorKey returns the ErrorKey field if non-nil, zero value otherwise.

### GetErrorKeyOk

`func (o *ErrorError) GetErrorKeyOk() (*string, bool)`

GetErrorKeyOk returns a tuple with the ErrorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorKey

`func (o *ErrorError) SetErrorKey(v string)`

SetErrorKey sets ErrorKey field to given value.

### HasErrorKey

`func (o *ErrorError) HasErrorKey() bool`

HasErrorKey returns a boolean if a field has been set.

### GetStatusCode

`func (o *ErrorError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ErrorError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ErrorError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetBriefSummary

`func (o *ErrorError) GetBriefSummary() string`

GetBriefSummary returns the BriefSummary field if non-nil, zero value otherwise.

### GetBriefSummaryOk

`func (o *ErrorError) GetBriefSummaryOk() (*string, bool)`

GetBriefSummaryOk returns a tuple with the BriefSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBriefSummary

`func (o *ErrorError) SetBriefSummary(v string)`

SetBriefSummary sets BriefSummary field to given value.


### GetStackTrace

`func (o *ErrorError) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ErrorError) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ErrorError) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.


### GetDescriptionURL

`func (o *ErrorError) GetDescriptionURL() string`

GetDescriptionURL returns the DescriptionURL field if non-nil, zero value otherwise.

### GetDescriptionURLOk

`func (o *ErrorError) GetDescriptionURLOk() (*string, bool)`

GetDescriptionURLOk returns a tuple with the DescriptionURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionURL

`func (o *ErrorError) SetDescriptionURL(v string)`

SetDescriptionURL sets DescriptionURL field to given value.


### GetLogId

`func (o *ErrorError) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *ErrorError) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *ErrorError) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *ErrorError) HasLogId() bool`

HasLogId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


