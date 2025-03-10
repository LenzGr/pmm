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

// StartPTMongoDBSummaryActionReader is a Reader for the StartPTMongoDBSummaryAction structure.
type StartPTMongoDBSummaryActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPTMongoDBSummaryActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPTMongoDBSummaryActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartPTMongoDBSummaryActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPTMongoDBSummaryActionOK creates a StartPTMongoDBSummaryActionOK with default headers values
func NewStartPTMongoDBSummaryActionOK() *StartPTMongoDBSummaryActionOK {
	return &StartPTMongoDBSummaryActionOK{}
}

/* StartPTMongoDBSummaryActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartPTMongoDBSummaryActionOK struct {
	Payload *StartPTMongoDBSummaryActionOKBody
}

func (o *StartPTMongoDBSummaryActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPTMongoDBSummary][%d] startPtMongoDbSummaryActionOk  %+v", 200, o.Payload)
}

func (o *StartPTMongoDBSummaryActionOK) GetPayload() *StartPTMongoDBSummaryActionOKBody {
	return o.Payload
}

func (o *StartPTMongoDBSummaryActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartPTMongoDBSummaryActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPTMongoDBSummaryActionDefault creates a StartPTMongoDBSummaryActionDefault with default headers values
func NewStartPTMongoDBSummaryActionDefault(code int) *StartPTMongoDBSummaryActionDefault {
	return &StartPTMongoDBSummaryActionDefault{
		_statusCode: code,
	}
}

/* StartPTMongoDBSummaryActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartPTMongoDBSummaryActionDefault struct {
	_statusCode int

	Payload *StartPTMongoDBSummaryActionDefaultBody
}

// Code gets the status code for the start PT mongo DB summary action default response
func (o *StartPTMongoDBSummaryActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPTMongoDBSummaryActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPTMongoDBSummary][%d] StartPTMongoDBSummaryAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartPTMongoDBSummaryActionDefault) GetPayload() *StartPTMongoDBSummaryActionDefaultBody {
	return o.Payload
}

func (o *StartPTMongoDBSummaryActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartPTMongoDBSummaryActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartPTMongoDBSummaryActionBody Message to prepare pt-mongodb-summary data
swagger:model StartPTMongoDBSummaryActionBody
*/
type StartPTMongoDBSummaryActionBody struct {
	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this start PT mongo DB summary action body
func (o *StartPTMongoDBSummaryActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT mongo DB summary action body based on context it is used
func (o *StartPTMongoDBSummaryActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionBody) UnmarshalBinary(b []byte) error {
	var res StartPTMongoDBSummaryActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTMongoDBSummaryActionDefaultBody start PT mongo DB summary action default body
swagger:model StartPTMongoDBSummaryActionDefaultBody
*/
type StartPTMongoDBSummaryActionDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartPTMongoDBSummaryActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start PT mongo DB summary action default body
func (o *StartPTMongoDBSummaryActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPTMongoDBSummaryActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("StartPTMongoDBSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPTMongoDBSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start PT mongo DB summary action default body based on the context it is used
func (o *StartPTMongoDBSummaryActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPTMongoDBSummaryActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartPTMongoDBSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartPTMongoDBSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPTMongoDBSummaryActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTMongoDBSummaryActionDefaultBodyDetailsItems0 start PT mongo DB summary action default body details items0
swagger:model StartPTMongoDBSummaryActionDefaultBodyDetailsItems0
*/
type StartPTMongoDBSummaryActionDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this start PT mongo DB summary action default body details items0
func (o *StartPTMongoDBSummaryActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT mongo DB summary action default body details items0 based on context it is used
func (o *StartPTMongoDBSummaryActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartPTMongoDBSummaryActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTMongoDBSummaryActionOKBody Message to retrieve the prepared pt-mongodb-summary data
swagger:model StartPTMongoDBSummaryActionOKBody
*/
type StartPTMongoDBSummaryActionOKBody struct {
	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start PT mongo DB summary action OK body
func (o *StartPTMongoDBSummaryActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start PT mongo DB summary action OK body based on context it is used
func (o *StartPTMongoDBSummaryActionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTMongoDBSummaryActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPTMongoDBSummaryActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
