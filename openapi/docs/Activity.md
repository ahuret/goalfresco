# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostPersonId** | **string** | The id of the person who performed the activity | 
**Id** | **int64** | The unique id of the activity | 
**SiteId** | Pointer to **string** | The unique id of the site on which the activity was performed | [optional] 
**PostedAt** | Pointer to **time.Time** | The date time at which the activity was performed | [optional] 
**FeedPersonId** | **string** | The feed on which this activity was posted | 
**ActivitySummary** | Pointer to **map[string]string** | An object summarizing the activity | [optional] 
**ActivityType** | **string** | The type of the activity posted | 

## Methods

### NewActivity

`func NewActivity(postPersonId string, id int64, feedPersonId string, activityType string, ) *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostPersonId

`func (o *Activity) GetPostPersonId() string`

GetPostPersonId returns the PostPersonId field if non-nil, zero value otherwise.

### GetPostPersonIdOk

`func (o *Activity) GetPostPersonIdOk() (*string, bool)`

GetPostPersonIdOk returns a tuple with the PostPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPersonId

`func (o *Activity) SetPostPersonId(v string)`

SetPostPersonId sets PostPersonId field to given value.


### GetId

`func (o *Activity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Activity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Activity) SetId(v int64)`

SetId sets Id field to given value.


### GetSiteId

`func (o *Activity) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Activity) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Activity) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Activity) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetPostedAt

`func (o *Activity) GetPostedAt() time.Time`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *Activity) GetPostedAtOk() (*time.Time, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *Activity) SetPostedAt(v time.Time)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *Activity) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetFeedPersonId

`func (o *Activity) GetFeedPersonId() string`

GetFeedPersonId returns the FeedPersonId field if non-nil, zero value otherwise.

### GetFeedPersonIdOk

`func (o *Activity) GetFeedPersonIdOk() (*string, bool)`

GetFeedPersonIdOk returns a tuple with the FeedPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPersonId

`func (o *Activity) SetFeedPersonId(v string)`

SetFeedPersonId sets FeedPersonId field to given value.


### GetActivitySummary

`func (o *Activity) GetActivitySummary() map[string]string`

GetActivitySummary returns the ActivitySummary field if non-nil, zero value otherwise.

### GetActivitySummaryOk

`func (o *Activity) GetActivitySummaryOk() (*map[string]string, bool)`

GetActivitySummaryOk returns a tuple with the ActivitySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySummary

`func (o *Activity) SetActivitySummary(v map[string]string)`

SetActivitySummary sets ActivitySummary field to given value.

### HasActivitySummary

`func (o *Activity) HasActivitySummary() bool`

HasActivitySummary returns a boolean if a field has been set.

### GetActivityType

`func (o *Activity) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *Activity) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *Activity) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


