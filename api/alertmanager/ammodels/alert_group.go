// Code generated by go-swagger; DO NOT EDIT.

package ammodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertGroup alert group
//
// swagger:model alertGroup
type AlertGroup struct {

	// alerts
	// Required: true
	Alerts []*GettableAlert `json:"alerts"`

	// labels
	// Required: true
	Labels LabelSet `json:"labels"`

	// receiver
	// Required: true
	Receiver *Receiver `json:"receiver"`
}

// Validate validates this alert group
func (m *AlertGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertGroup) validateAlerts(formats strfmt.Registry) error {

	if err := validate.Required("alerts", "body", m.Alerts); err != nil {
		return err
	}

	for i := 0; i < len(m.Alerts); i++ {
		if swag.IsZero(m.Alerts[i]) { // not required
			continue
		}

		if m.Alerts[i] != nil {
			if err := m.Alerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertGroup) validateLabels(formats strfmt.Registry) error {

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

func (m *AlertGroup) validateReceiver(formats strfmt.Registry) error {

	if err := validate.Required("receiver", "body", m.Receiver); err != nil {
		return err
	}

	if m.Receiver != nil {
		if err := m.Receiver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receiver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("receiver")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alert group based on the context it is used
func (m *AlertGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReceiver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertGroup) contextValidateAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Alerts); i++ {

		if m.Alerts[i] != nil {
			if err := m.Alerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertGroup) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AlertGroup) contextValidateReceiver(ctx context.Context, formats strfmt.Registry) error {

	if m.Receiver != nil {
		if err := m.Receiver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receiver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("receiver")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertGroup) UnmarshalBinary(b []byte) error {
	var res AlertGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
