// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddDataSourceCommand Also acts as api DTO
//
// swagger:model AddDataSourceCommand
type AddDataSourceCommand struct {

	// basic auth
	BasicAuth bool `json:"basicAuth,omitempty"`

	// basic auth password
	BasicAuthPassword string `json:"basicAuthPassword,omitempty"`

	// basic auth user
	BasicAuthUser string `json:"basicAuthUser,omitempty"`

	// database
	Database string `json:"database,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// secure Json data
	SecureJSONData map[string]string `json:"secureJsonData,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// Uid
	UID string `json:"uid,omitempty"`

	// Url
	URL string `json:"url,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// with credentials
	WithCredentials bool `json:"withCredentials,omitempty"`

	// access
	Access DsAccess `json:"access,omitempty"`

	// json data
	JSONData JSON `json:"jsonData,omitempty"`
}

// Validate validates this add data source command
func (m *AddDataSourceCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddDataSourceCommand) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if err := m.Access.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

// ContextValidate validate this add data source command based on the context it is used
func (m *AddDataSourceCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddDataSourceCommand) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Access.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddDataSourceCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddDataSourceCommand) UnmarshalBinary(b []byte) error {
	var res AddDataSourceCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}