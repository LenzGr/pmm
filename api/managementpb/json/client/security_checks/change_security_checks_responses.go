// Code generated by go-swagger; DO NOT EDIT.

package security_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeSecurityChecksReader is a Reader for the ChangeSecurityChecks structure.
type ChangeSecurityChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeSecurityChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeSecurityChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeSecurityChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeSecurityChecksOK creates a ChangeSecurityChecksOK with default headers values
func NewChangeSecurityChecksOK() *ChangeSecurityChecksOK {
	return &ChangeSecurityChecksOK{}
}

/* ChangeSecurityChecksOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeSecurityChecksOK struct {
	Payload interface{}
}

func (o *ChangeSecurityChecksOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/Change][%d] changeSecurityChecksOk  %+v", 200, o.Payload)
}

func (o *ChangeSecurityChecksOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeSecurityChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeSecurityChecksDefault creates a ChangeSecurityChecksDefault with default headers values
func NewChangeSecurityChecksDefault(code int) *ChangeSecurityChecksDefault {
	return &ChangeSecurityChecksDefault{
		_statusCode: code,
	}
}

/* ChangeSecurityChecksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeSecurityChecksDefault struct {
	_statusCode int

	Payload *ChangeSecurityChecksDefaultBody
}

// Code gets the status code for the change security checks default response
func (o *ChangeSecurityChecksDefault) Code() int {
	return o._statusCode
}

func (o *ChangeSecurityChecksDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/Change][%d] ChangeSecurityChecks default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeSecurityChecksDefault) GetPayload() *ChangeSecurityChecksDefaultBody {
	return o.Payload
}

func (o *ChangeSecurityChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeSecurityChecksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeSecurityChecksBody change security checks body
swagger:model ChangeSecurityChecksBody
*/
type ChangeSecurityChecksBody struct {
	// params
	Params []*ChangeSecurityChecksParamsBodyParamsItems0 `json:"params"`
}

// Validate validates this change security checks body
func (o *ChangeSecurityChecksBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSecurityChecksBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change security checks body based on the context it is used
func (o *ChangeSecurityChecksBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSecurityChecksBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Params); i++ {
		if o.Params[i] != nil {
			if err := o.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSecurityChecksBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSecurityChecksBody) UnmarshalBinary(b []byte) error {
	var res ChangeSecurityChecksBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSecurityChecksDefaultBody change security checks default body
swagger:model ChangeSecurityChecksDefaultBody
*/
type ChangeSecurityChecksDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeSecurityChecksDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change security checks default body
func (o *ChangeSecurityChecksDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSecurityChecksDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change security checks default body based on the context it is used
func (o *ChangeSecurityChecksDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSecurityChecksDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSecurityChecksDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSecurityChecksDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeSecurityChecksDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSecurityChecksDefaultBodyDetailsItems0 change security checks default body details items0
swagger:model ChangeSecurityChecksDefaultBodyDetailsItems0
*/
type ChangeSecurityChecksDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this change security checks default body details items0
func (o *ChangeSecurityChecksDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change security checks default body details items0 based on context it is used
func (o *ChangeSecurityChecksDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSecurityChecksDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSecurityChecksDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeSecurityChecksDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSecurityChecksParamsBodyParamsItems0 ChangeSecurityCheckParams specifies a single check parameters.
swagger:model ChangeSecurityChecksParamsBodyParamsItems0
*/
type ChangeSecurityChecksParamsBodyParamsItems0 struct {
	// The name of the check to change.
	Name string `json:"name,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`

	// disable
	Disable bool `json:"disable,omitempty"`

	// SecurityCheckInterval represents possible execution interval values for checks.
	// Enum: [SECURITY_CHECK_INTERVAL_INVALID STANDARD FREQUENT RARE]
	Interval *string `json:"interval,omitempty"`
}

// Validate validates this change security checks params body params items0
func (o *ChangeSecurityChecksParamsBodyParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeSecurityChecksParamsBodyParamsItems0TypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SECURITY_CHECK_INTERVAL_INVALID","STANDARD","FREQUENT","RARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeSecurityChecksParamsBodyParamsItems0TypeIntervalPropEnum = append(changeSecurityChecksParamsBodyParamsItems0TypeIntervalPropEnum, v)
	}
}

const (

	// ChangeSecurityChecksParamsBodyParamsItems0IntervalSECURITYCHECKINTERVALINVALID captures enum value "SECURITY_CHECK_INTERVAL_INVALID"
	ChangeSecurityChecksParamsBodyParamsItems0IntervalSECURITYCHECKINTERVALINVALID string = "SECURITY_CHECK_INTERVAL_INVALID"

	// ChangeSecurityChecksParamsBodyParamsItems0IntervalSTANDARD captures enum value "STANDARD"
	ChangeSecurityChecksParamsBodyParamsItems0IntervalSTANDARD string = "STANDARD"

	// ChangeSecurityChecksParamsBodyParamsItems0IntervalFREQUENT captures enum value "FREQUENT"
	ChangeSecurityChecksParamsBodyParamsItems0IntervalFREQUENT string = "FREQUENT"

	// ChangeSecurityChecksParamsBodyParamsItems0IntervalRARE captures enum value "RARE"
	ChangeSecurityChecksParamsBodyParamsItems0IntervalRARE string = "RARE"
)

// prop value enum
func (o *ChangeSecurityChecksParamsBodyParamsItems0) validateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeSecurityChecksParamsBodyParamsItems0TypeIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeSecurityChecksParamsBodyParamsItems0) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(o.Interval) { // not required
		return nil
	}

	// value enum
	if err := o.validateIntervalEnum("interval", "body", *o.Interval); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change security checks params body params items0 based on context it is used
func (o *ChangeSecurityChecksParamsBodyParamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSecurityChecksParamsBodyParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSecurityChecksParamsBodyParamsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeSecurityChecksParamsBodyParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
