// Code generated by go-swagger; DO NOT EDIT.

package ammodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Alert alert
//
// swagger:model alert
type Alert struct {
	// generator URL
	// Format: uri
	GeneratorURL strfmt.URI `json:"generatorURL,omitempty"`

	// labels
	// Required: true
	Labels LabelSet `json:"labels"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneratorURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) validateGeneratorURL(formats strfmt.Registry) error {
	if swag.IsZero(m.GeneratorURL) { // not required
		return nil
	}

	if err := validate.FormatOf("generatorURL", "body", "uri", m.GeneratorURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Alert) validateLabels(formats strfmt.Registry) error {
	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alert based on the context it is used
func (m *Alert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {
	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
