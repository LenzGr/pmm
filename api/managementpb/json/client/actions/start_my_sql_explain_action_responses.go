// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartMySQLExplainActionReader is a Reader for the StartMySQLExplainAction structure.
type StartMySQLExplainActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMySQLExplainActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartMySQLExplainActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartMySQLExplainActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMySQLExplainActionOK creates a StartMySQLExplainActionOK with default headers values
func NewStartMySQLExplainActionOK() *StartMySQLExplainActionOK {
	return &StartMySQLExplainActionOK{}
}

/* StartMySQLExplainActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartMySQLExplainActionOK struct {
	Payload *StartMySQLExplainActionOKBody
}

func (o *StartMySQLExplainActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLExplain][%d] startMySqlExplainActionOk  %+v", 200, o.Payload)
}
func (o *StartMySQLExplainActionOK) GetPayload() *StartMySQLExplainActionOKBody {
	return o.Payload
}

func (o *StartMySQLExplainActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLExplainActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMySQLExplainActionDefault creates a StartMySQLExplainActionDefault with default headers values
func NewStartMySQLExplainActionDefault(code int) *StartMySQLExplainActionDefault {
	return &StartMySQLExplainActionDefault{
		_statusCode: code,
	}
}

/* StartMySQLExplainActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartMySQLExplainActionDefault struct {
	_statusCode int

	Payload *StartMySQLExplainActionDefaultBody
}

// Code gets the status code for the start my SQL explain action default response
func (o *StartMySQLExplainActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMySQLExplainActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLExplain][%d] StartMySQLExplainAction default  %+v", o._statusCode, o.Payload)
}
func (o *StartMySQLExplainActionDefault) GetPayload() *StartMySQLExplainActionDefaultBody {
	return o.Payload
}

func (o *StartMySQLExplainActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLExplainActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartMySQLExplainActionBody start my SQL explain action body
swagger:model StartMySQLExplainActionBody
*/
type StartMySQLExplainActionBody struct {

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// SQL query. Required.
	Query string `json:"query,omitempty"`

	// Database name. Required if it can't be deduced from the query.
	Database string `json:"database,omitempty"`
}

// Validate validates this start my SQL explain action body
func (o *StartMySQLExplainActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL explain action body based on context it is used
func (o *StartMySQLExplainActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainActionBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLExplainActionDefaultBody start my SQL explain action default body
swagger:model StartMySQLExplainActionDefaultBody
*/
type StartMySQLExplainActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartMySQLExplainActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start my SQL explain action default body
func (o *StartMySQLExplainActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMySQLExplainActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartMySQLExplainAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartMySQLExplainAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start my SQL explain action default body based on the context it is used
func (o *StartMySQLExplainActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMySQLExplainActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartMySQLExplainAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartMySQLExplainAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLExplainActionDefaultBodyDetailsItems0 start my SQL explain action default body details items0
swagger:model StartMySQLExplainActionDefaultBodyDetailsItems0
*/
type StartMySQLExplainActionDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this start my SQL explain action default body details items0
func (o *StartMySQLExplainActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL explain action default body details items0 based on context it is used
func (o *StartMySQLExplainActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLExplainActionOKBody start my SQL explain action OK body
swagger:model StartMySQLExplainActionOKBody
*/
type StartMySQLExplainActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start my SQL explain action OK body
func (o *StartMySQLExplainActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL explain action OK body based on context it is used
func (o *StartMySQLExplainActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
