/*
Receipt Processor

A simple receipt processor

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ReceiptsIdPointsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiptsIdPointsGet200Response{}

// ReceiptsIdPointsGet200Response struct for ReceiptsIdPointsGet200Response
type ReceiptsIdPointsGet200Response struct {
	Points *int64 `json:"points,omitempty"`
}

// NewReceiptsIdPointsGet200Response instantiates a new ReceiptsIdPointsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptsIdPointsGet200Response() *ReceiptsIdPointsGet200Response {
	this := ReceiptsIdPointsGet200Response{}
	return &this
}

// NewReceiptsIdPointsGet200ResponseWithDefaults instantiates a new ReceiptsIdPointsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptsIdPointsGet200ResponseWithDefaults() *ReceiptsIdPointsGet200Response {
	this := ReceiptsIdPointsGet200Response{}
	return &this
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *ReceiptsIdPointsGet200Response) GetPoints() int64 {
	if o == nil || IsNil(o.Points) {
		var ret int64
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptsIdPointsGet200Response) GetPointsOk() (*int64, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *ReceiptsIdPointsGet200Response) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given int64 and assigns it to the Points field.
func (o *ReceiptsIdPointsGet200Response) SetPoints(v int64) {
	o.Points = &v
}

func (o ReceiptsIdPointsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptsIdPointsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	return toSerialize, nil
}

type NullableReceiptsIdPointsGet200Response struct {
	value *ReceiptsIdPointsGet200Response
	isSet bool
}

func (v NullableReceiptsIdPointsGet200Response) Get() *ReceiptsIdPointsGet200Response {
	return v.value
}

func (v *NullableReceiptsIdPointsGet200Response) Set(val *ReceiptsIdPointsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptsIdPointsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptsIdPointsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptsIdPointsGet200Response(val *ReceiptsIdPointsGet200Response) *NullableReceiptsIdPointsGet200Response {
	return &NullableReceiptsIdPointsGet200Response{value: val, isSet: true}
}

func (v NullableReceiptsIdPointsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptsIdPointsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


