// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Float Float is a nullable float64.
//
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
//
// swagger:model Float
type Float struct {

	// float64
	Float64 float64 `json:"Float64,omitempty"`

	// valid
	Valid bool `json:"Valid,omitempty"`
}

// Validate validates this float
func (m *Float) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this float based on context it is used
func (m *Float) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Float) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Float) UnmarshalBinary(b []byte) error {
	var res Float
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}