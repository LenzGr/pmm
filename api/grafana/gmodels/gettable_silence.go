// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GettableSilence GettableSilence gettable silence
//
// swagger:model gettableSilence
type GettableSilence struct {

	// comment
	// Required: true
	Comment *string `json:"comment"`

	// created by
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// ends at
	// Required: true
	// Format: date-time
	EndsAt *strfmt.DateTime `json:"endsAt"`

	// id
	// Required: true
	ID *string `json:"id"`

	// starts at
	// Required: true
	// Format: date-time
	StartsAt *strfmt.DateTime `json:"startsAt"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// matchers
	// Required: true
	Matchers Matchers `json:"matchers"`

	// status
	// Required: true
	Status *SilenceStatus `json:"status"`
}

// Validate validates this gettable silence
func (m *GettableSilence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableSilence) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *GettableSilence) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *GettableSilence) validateEndsAt(formats strfmt.Registry) error {

	if err := validate.Required("endsAt", "body", m.EndsAt); err != nil {
		return err
	}

	if err := validate.FormatOf("endsAt", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GettableSilence) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GettableSilence) validateStartsAt(formats strfmt.Registry) error {

	if err := validate.Required("startsAt", "body", m.StartsAt); err != nil {
		return err
	}

	if err := validate.FormatOf("startsAt", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GettableSilence) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GettableSilence) validateMatchers(formats strfmt.Registry) error {

	if err := validate.Required("matchers", "body", m.Matchers); err != nil {
		return err
	}

	if err := m.Matchers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("matchers")
		}
		return err
	}

	return nil
}

func (m *GettableSilence) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gettable silence based on the context it is used
func (m *GettableSilence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableSilence) contextValidateMatchers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Matchers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("matchers")
		}
		return err
	}

	return nil
}

func (m *GettableSilence) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GettableSilence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GettableSilence) UnmarshalBinary(b []byte) error {
	var res GettableSilence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}