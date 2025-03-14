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

// StartMySQLShowCreateTableActionReader is a Reader for the StartMySQLShowCreateTableAction structure.
type StartMySQLShowCreateTableActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMySQLShowCreateTableActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartMySQLShowCreateTableActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartMySQLShowCreateTableActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMySQLShowCreateTableActionOK creates a StartMySQLShowCreateTableActionOK with default headers values
func NewStartMySQLShowCreateTableActionOK() *StartMySQLShowCreateTableActionOK {
	return &StartMySQLShowCreateTableActionOK{}
}

/* StartMySQLShowCreateTableActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartMySQLShowCreateTableActionOK struct {
	Payload *StartMySQLShowCreateTableActionOKBody
}

func (o *StartMySQLShowCreateTableActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLShowCreateTable][%d] startMySqlShowCreateTableActionOk  %+v", 200, o.Payload)
}

func (o *StartMySQLShowCreateTableActionOK) GetPayload() *StartMySQLShowCreateTableActionOKBody {
	return o.Payload
}

func (o *StartMySQLShowCreateTableActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartMySQLShowCreateTableActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMySQLShowCreateTableActionDefault creates a StartMySQLShowCreateTableActionDefault with default headers values
func NewStartMySQLShowCreateTableActionDefault(code int) *StartMySQLShowCreateTableActionDefault {
	return &StartMySQLShowCreateTableActionDefault{
		_statusCode: code,
	}
}

/* StartMySQLShowCreateTableActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartMySQLShowCreateTableActionDefault struct {
	_statusCode int

	Payload *StartMySQLShowCreateTableActionDefaultBody
}

// Code gets the status code for the start my SQL show create table action default response
func (o *StartMySQLShowCreateTableActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMySQLShowCreateTableActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLShowCreateTable][%d] StartMySQLShowCreateTableAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartMySQLShowCreateTableActionDefault) GetPayload() *StartMySQLShowCreateTableActionDefaultBody {
	return o.Payload
}

func (o *StartMySQLShowCreateTableActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartMySQLShowCreateTableActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartMySQLShowCreateTableActionBody start my SQL show create table action body
swagger:model StartMySQLShowCreateTableActionBody
*/
type StartMySQLShowCreateTableActionBody struct {
	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Table name. Required. May additionally contain a database name.
	TableName string `json:"table_name,omitempty"`

	// Database name. Required if not given in the table_name field.
	Database string `json:"database,omitempty"`
}

// Validate validates this start my SQL show create table action body
func (o *StartMySQLShowCreateTableActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL show create table action body based on context it is used
func (o *StartMySQLShowCreateTableActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowCreateTableActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLShowCreateTableActionDefaultBody start my SQL show create table action default body
swagger:model StartMySQLShowCreateTableActionDefaultBody
*/
type StartMySQLShowCreateTableActionDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartMySQLShowCreateTableActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start my SQL show create table action default body
func (o *StartMySQLShowCreateTableActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMySQLShowCreateTableActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartMySQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartMySQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start my SQL show create table action default body based on the context it is used
func (o *StartMySQLShowCreateTableActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMySQLShowCreateTableActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartMySQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartMySQLShowCreateTableAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowCreateTableActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLShowCreateTableActionDefaultBodyDetailsItems0 start my SQL show create table action default body details items0
swagger:model StartMySQLShowCreateTableActionDefaultBodyDetailsItems0
*/
type StartMySQLShowCreateTableActionDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this start my SQL show create table action default body details items0
func (o *StartMySQLShowCreateTableActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL show create table action default body details items0 based on context it is used
func (o *StartMySQLShowCreateTableActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowCreateTableActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLShowCreateTableActionOKBody start my SQL show create table action OK body
swagger:model StartMySQLShowCreateTableActionOKBody
*/
type StartMySQLShowCreateTableActionOKBody struct {
	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start my SQL show create table action OK body
func (o *StartMySQLShowCreateTableActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start my SQL show create table action OK body based on context it is used
func (o *StartMySQLShowCreateTableActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowCreateTableActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowCreateTableActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
