/*
 * Decision Service Fleet Manager
 *
 * Decision Service Fleet Manager is a Rest API to manage decision instances and connectors.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package decisclient

import (
	"encoding/json"
	"time"
)

// DecisionRequestAllOf struct for DecisionRequestAllOf
type DecisionRequestAllOf struct {
	Status        *string                       `json:"status,omitempty"`
	StatusMessage *string                       `json:"status_message,omitempty"`
	Description   *string                       `json:"description,omitempty"`
	Name          *string                       `json:"name,omitempty"`
	Url           *string                       `json:"url,omitempty"`
	Model         *DecisionRequestAllOfModel    `json:"model,omitempty"`
	Eventing      *DecisionRequestAllOfEventing `json:"eventing,omitempty"`
	SubmittedAt   *time.Time                    `json:"submitted_at,omitempty"`
	PublishedAt   *time.Time                    `json:"published_at,omitempty"`
	Configuration *map[string]string            `json:"configuration,omitempty"`
	Tags          *map[string]string            `json:"tags,omitempty"`
	Version       *int                          `json:"version,omitempty"`
}

// NewDecisionRequestAllOf instantiates a new DecisionRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionRequestAllOf() *DecisionRequestAllOf {
	this := DecisionRequestAllOf{}
	return &this
}

// NewDecisionRequestAllOfWithDefaults instantiates a new DecisionRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionRequestAllOfWithDefaults() *DecisionRequestAllOf {
	this := DecisionRequestAllOf{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DecisionRequestAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *DecisionRequestAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DecisionRequestAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DecisionRequestAllOf) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DecisionRequestAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetModel() DecisionRequestAllOfModel {
	if o == nil || o.Model == nil {
		var ret DecisionRequestAllOfModel
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetModelOk() (*DecisionRequestAllOfModel, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given DecisionRequestAllOfModel and assigns it to the Model field.
func (o *DecisionRequestAllOf) SetModel(v DecisionRequestAllOfModel) {
	o.Model = &v
}

// GetEventing returns the Eventing field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetEventing() DecisionRequestAllOfEventing {
	if o == nil || o.Eventing == nil {
		var ret DecisionRequestAllOfEventing
		return ret
	}
	return *o.Eventing
}

// GetEventingOk returns a tuple with the Eventing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetEventingOk() (*DecisionRequestAllOfEventing, bool) {
	if o == nil || o.Eventing == nil {
		return nil, false
	}
	return o.Eventing, true
}

// HasEventing returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasEventing() bool {
	if o != nil && o.Eventing != nil {
		return true
	}

	return false
}

// SetEventing gets a reference to the given DecisionRequestAllOfEventing and assigns it to the Eventing field.
func (o *DecisionRequestAllOf) SetEventing(v DecisionRequestAllOfEventing) {
	o.Eventing = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetSubmittedAt() time.Time {
	if o == nil || o.SubmittedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil || o.SubmittedAt == nil {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasSubmittedAt() bool {
	if o != nil && o.SubmittedAt != nil {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given time.Time and assigns it to the SubmittedAt field.
func (o *DecisionRequestAllOf) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetPublishedAt() time.Time {
	if o == nil || o.PublishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || o.PublishedAt == nil {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasPublishedAt() bool {
	if o != nil && o.PublishedAt != nil {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *DecisionRequestAllOf) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetConfiguration() map[string]string {
	if o == nil || o.Configuration == nil {
		var ret map[string]string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *DecisionRequestAllOf) SetConfiguration(v map[string]string) {
	o.Configuration = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *DecisionRequestAllOf) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DecisionRequestAllOf) GetVersion() int {
	if o == nil || o.Version == nil {
		var ret int
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOf) GetVersionOk() (*int, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DecisionRequestAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int and assigns it to the Version field.
func (o *DecisionRequestAllOf) SetVersion(v int) {
	o.Version = &v
}

func (o DecisionRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["status_message"] = o.StatusMessage
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Eventing != nil {
		toSerialize["eventing"] = o.Eventing
	}
	if o.SubmittedAt != nil {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if o.PublishedAt != nil {
		toSerialize["published_at"] = o.PublishedAt
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDecisionRequestAllOf struct {
	value *DecisionRequestAllOf
	isSet bool
}

func (v NullableDecisionRequestAllOf) Get() *DecisionRequestAllOf {
	return v.value
}

func (v *NullableDecisionRequestAllOf) Set(val *DecisionRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionRequestAllOf(val *DecisionRequestAllOf) *NullableDecisionRequestAllOf {
	return &NullableDecisionRequestAllOf{value: val, isSet: true}
}

func (v NullableDecisionRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
