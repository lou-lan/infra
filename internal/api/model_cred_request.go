/*
Infra API

Infra REST API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CredRequest struct for CredRequest
type CredRequest struct {
	Destination *string `json:"destination,omitempty" validate:"required"`
}

// NewCredRequest instantiates a new CredRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredRequest() *CredRequest {
	this := CredRequest{}
	return &this
}

// NewCredRequestWithDefaults instantiates a new CredRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredRequestWithDefaults() *CredRequest {
	this := CredRequest{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *CredRequest) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredRequest) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *CredRequest) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *CredRequest) SetDestination(v string) {
	o.Destination = &v
}

func (o CredRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableCredRequest struct {
	value *CredRequest
	isSet bool
}

func (v NullableCredRequest) Get() *CredRequest {
	return v.value
}

func (v *NullableCredRequest) Set(val *CredRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCredRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCredRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredRequest(val *CredRequest) *NullableCredRequest {
	return &NullableCredRequest{value: val, isSet: true}
}

func (v NullableCredRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
