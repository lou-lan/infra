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

// InfraAPIKey struct for InfraAPIKey
type InfraAPIKey struct {
	Id      string `json:"id"`
	Created int64  `json:"created"`
	Name    string `json:"name"`
}

// NewInfraAPIKey instantiates a new InfraAPIKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraAPIKey(id string, created int64, name string) *InfraAPIKey {
	this := InfraAPIKey{}
	this.Id = id
	this.Created = created
	this.Name = name
	return &this
}

// NewInfraAPIKeyWithDefaults instantiates a new InfraAPIKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraAPIKeyWithDefaults() *InfraAPIKey {
	this := InfraAPIKey{}
	return &this
}

// GetId returns the Id field value
func (o *InfraAPIKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InfraAPIKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InfraAPIKey) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *InfraAPIKey) GetCreated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *InfraAPIKey) GetCreatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *InfraAPIKey) SetCreated(v int64) {
	o.Created = v
}

// GetName returns the Name field value
func (o *InfraAPIKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InfraAPIKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InfraAPIKey) SetName(v string) {
	o.Name = v
}

func (o InfraAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInfraAPIKey struct {
	value *InfraAPIKey
	isSet bool
}

func (v NullableInfraAPIKey) Get() *InfraAPIKey {
	return v.value
}

func (v *NullableInfraAPIKey) Set(val *InfraAPIKey) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraAPIKey) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraAPIKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraAPIKey(val *InfraAPIKey) *NullableInfraAPIKey {
	return &NullableInfraAPIKey{value: val, isSet: true}
}

func (v NullableInfraAPIKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraAPIKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}