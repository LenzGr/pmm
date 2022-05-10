// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RuleGroupConfigResponse rule group config response
//
// swagger:model RuleGroupConfigResponse
type RuleGroupConfigResponse struct {

	// name
	Name string `json:"name,omitempty"`

	// rules
	Rules []*GettableExtendedRuleNode `json:"rules"`

	// source tenants
	SourceTenants []string `json:"source_tenants"`

	// interval
	// Format: duration
	Interval Duration `json:"interval,omitempty"`
}

// Validate validates this rule group config response
func (m *RuleGroupConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleGroupConfigResponse) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleGroupConfigResponse) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	if err := m.Interval.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interval")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interval")
		}
		return err
	}

	return nil
}

// ContextValidate validate this rule group config response based on the context it is used
func (m *RuleGroupConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleGroupConfigResponse) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {
			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleGroupConfigResponse) contextValidateInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Interval.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("interval")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("interval")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuleGroupConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleGroupConfigResponse) UnmarshalBinary(b []byte) error {
	var res RuleGroupConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
