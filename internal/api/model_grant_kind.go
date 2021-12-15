/*
Infra API

Infra REST API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// GrantKind the model 'GrantKind'
type GrantKind string

// List of GrantKind
const (
	GRANTKIND_KUBERNETES GrantKind = "kubernetes"
)

var allowedGrantKindEnumValues = []GrantKind{
	"kubernetes",
}

func (v *GrantKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrantKind(value)
	for _, existing := range allowedGrantKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrantKind", value)
}

// NewGrantKindFromValue returns a pointer to a valid GrantKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrantKindFromValue(v string) (*GrantKind, error) {
	ev := GrantKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrantKind: valid values are %v", v, allowedGrantKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrantKind) IsValid() bool {
	for _, existing := range allowedGrantKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GrantKind value
func (v GrantKind) Ptr() *GrantKind {
	return &v
}

type NullableGrantKind struct {
	value *GrantKind
	isSet bool
}

func (v NullableGrantKind) Get() *GrantKind {
	return v.value
}

func (v *NullableGrantKind) Set(val *GrantKind) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantKind) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantKind(val *GrantKind) *NullableGrantKind {
	return &NullableGrantKind{value: val, isSet: true}
}

func (v NullableGrantKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}