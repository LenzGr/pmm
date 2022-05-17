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

// TimeInterval TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
// within the interval.
//
// swagger:model TimeInterval
type TimeInterval struct {

	// days of month
	DaysOfMonth []*DayOfMonthRange `json:"days_of_month"`

	// months
	Months []*MonthRange `json:"months"`

	// times
	Times []*TimeRange `json:"times"`

	// weekdays
	Weekdays []*WeekdayRange `json:"weekdays"`

	// years
	Years []*YearRange `json:"years"`
}

// Validate validates this time interval
func (m *TimeInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaysOfMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekdays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYears(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeInterval) validateDaysOfMonth(formats strfmt.Registry) error {
	if swag.IsZero(m.DaysOfMonth) { // not required
		return nil
	}

	for i := 0; i < len(m.DaysOfMonth); i++ {
		if swag.IsZero(m.DaysOfMonth[i]) { // not required
			continue
		}

		if m.DaysOfMonth[i] != nil {
			if err := m.DaysOfMonth[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("days_of_month" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("days_of_month" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) validateMonths(formats strfmt.Registry) error {
	if swag.IsZero(m.Months) { // not required
		return nil
	}

	for i := 0; i < len(m.Months); i++ {
		if swag.IsZero(m.Months[i]) { // not required
			continue
		}

		if m.Months[i] != nil {
			if err := m.Months[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("months" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("months" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) validateTimes(formats strfmt.Registry) error {
	if swag.IsZero(m.Times) { // not required
		return nil
	}

	for i := 0; i < len(m.Times); i++ {
		if swag.IsZero(m.Times[i]) { // not required
			continue
		}

		if m.Times[i] != nil {
			if err := m.Times[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("times" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("times" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) validateWeekdays(formats strfmt.Registry) error {
	if swag.IsZero(m.Weekdays) { // not required
		return nil
	}

	for i := 0; i < len(m.Weekdays); i++ {
		if swag.IsZero(m.Weekdays[i]) { // not required
			continue
		}

		if m.Weekdays[i] != nil {
			if err := m.Weekdays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weekdays" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weekdays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) validateYears(formats strfmt.Registry) error {
	if swag.IsZero(m.Years) { // not required
		return nil
	}

	for i := 0; i < len(m.Years); i++ {
		if swag.IsZero(m.Years[i]) { // not required
			continue
		}

		if m.Years[i] != nil {
			if err := m.Years[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("years" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("years" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this time interval based on the context it is used
func (m *TimeInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDaysOfMonth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeekdays(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateYears(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeInterval) contextValidateDaysOfMonth(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DaysOfMonth); i++ {

		if m.DaysOfMonth[i] != nil {
			if err := m.DaysOfMonth[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("days_of_month" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("days_of_month" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) contextValidateMonths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Months); i++ {

		if m.Months[i] != nil {
			if err := m.Months[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("months" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("months" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) contextValidateTimes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Times); i++ {

		if m.Times[i] != nil {
			if err := m.Times[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("times" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("times" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) contextValidateWeekdays(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Weekdays); i++ {

		if m.Weekdays[i] != nil {
			if err := m.Weekdays[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weekdays" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weekdays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TimeInterval) contextValidateYears(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Years); i++ {

		if m.Years[i] != nil {
			if err := m.Years[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("years" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("years" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeInterval) UnmarshalBinary(b []byte) error {
	var res TimeInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}