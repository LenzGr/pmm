// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateTeamCommand update team command
//
// swagger:model UpdateTeamCommand
type UpdateTeamCommand struct {

	// email
	Email string `json:"Email,omitempty"`

	// Id
	ID int64 `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this update team command
func (m *UpdateTeamCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update team command based on context it is used
func (m *UpdateTeamCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTeamCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTeamCommand) UnmarshalBinary(b []byte) error {
	var res UpdateTeamCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}