/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MessageAttemptExhaustedEventData struct for MessageAttemptExhaustedEventData
type MessageAttemptExhaustedEventData struct {
	AppId string `json:"appId"`
	// Optional unique identifier for the application
	AppUid *string `json:"appUid,omitempty"`
	MsgId string `json:"msgId"`
	EndpointId string `json:"endpointId"`
	LastAttempt MessageAttemptFailedEvent `json:"lastAttempt"`
}

// NewMessageAttemptExhaustedEventData instantiates a new MessageAttemptExhaustedEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttemptExhaustedEventData(appId string, msgId string, endpointId string, lastAttempt MessageAttemptFailedEvent) *MessageAttemptExhaustedEventData {
	this := MessageAttemptExhaustedEventData{}
	this.AppId = appId
	this.MsgId = msgId
	this.EndpointId = endpointId
	this.LastAttempt = lastAttempt
	return &this
}

// NewMessageAttemptExhaustedEventDataWithDefaults instantiates a new MessageAttemptExhaustedEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttemptExhaustedEventDataWithDefaults() *MessageAttemptExhaustedEventData {
	this := MessageAttemptExhaustedEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *MessageAttemptExhaustedEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MessageAttemptExhaustedEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise.
func (o *MessageAttemptExhaustedEventData) GetAppUid() string {
	if o == nil || o.AppUid == nil {
		var ret string
		return ret
	}
	return *o.AppUid
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetAppUidOk() (*string, bool) {
	if o == nil || o.AppUid == nil {
		return nil, false
	}
	return o.AppUid, true
}

// HasAppUid returns a boolean if a field has been set.
func (o *MessageAttemptExhaustedEventData) HasAppUid() bool {
	if o != nil && o.AppUid != nil {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given string and assigns it to the AppUid field.
func (o *MessageAttemptExhaustedEventData) SetAppUid(v string) {
	o.AppUid = &v
}

// GetMsgId returns the MsgId field value
func (o *MessageAttemptExhaustedEventData) GetMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetMsgIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MsgId, true
}

// SetMsgId sets field value
func (o *MessageAttemptExhaustedEventData) SetMsgId(v string) {
	o.MsgId = v
}

// GetEndpointId returns the EndpointId field value
func (o *MessageAttemptExhaustedEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *MessageAttemptExhaustedEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetLastAttempt returns the LastAttempt field value
func (o *MessageAttemptExhaustedEventData) GetLastAttempt() MessageAttemptFailedEvent {
	if o == nil {
		var ret MessageAttemptFailedEvent
		return ret
	}

	return o.LastAttempt
}

// GetLastAttemptOk returns a tuple with the LastAttempt field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetLastAttemptOk() (*MessageAttemptFailedEvent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastAttempt, true
}

// SetLastAttempt sets field value
func (o *MessageAttemptExhaustedEventData) SetLastAttempt(v MessageAttemptFailedEvent) {
	o.LastAttempt = v
}

func (o MessageAttemptExhaustedEventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if o.AppUid != nil {
		toSerialize["appUid"] = o.AppUid
	}
	if true {
		toSerialize["msgId"] = o.MsgId
	}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	if true {
		toSerialize["lastAttempt"] = o.LastAttempt
	}
	return json.Marshal(toSerialize)
}

type NullableMessageAttemptExhaustedEventData struct {
	value *MessageAttemptExhaustedEventData
	isSet bool
}

func (v NullableMessageAttemptExhaustedEventData) Get() *MessageAttemptExhaustedEventData {
	return v.value
}

func (v *NullableMessageAttemptExhaustedEventData) Set(val *MessageAttemptExhaustedEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttemptExhaustedEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttemptExhaustedEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttemptExhaustedEventData(val *MessageAttemptExhaustedEventData) *NullableMessageAttemptExhaustedEventData {
	return &NullableMessageAttemptExhaustedEventData{value: val, isSet: true}
}

func (v NullableMessageAttemptExhaustedEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttemptExhaustedEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


