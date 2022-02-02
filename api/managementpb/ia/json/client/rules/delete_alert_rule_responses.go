// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteAlertRuleReader is a Reader for the DeleteAlertRule structure.
type DeleteAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAlertRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAlertRuleOK creates a DeleteAlertRuleOK with default headers values
func NewDeleteAlertRuleOK() *DeleteAlertRuleOK {
	return &DeleteAlertRuleOK{}
}

/*DeleteAlertRuleOK handles this case with default header values.

A successful response.
*/
type DeleteAlertRuleOK struct {
	Payload interface{}
}

func (o *DeleteAlertRuleOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Delete][%d] deleteAlertRuleOk  %+v", 200, o.Payload)
}

func (o *DeleteAlertRuleOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAlertRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertRuleDefault creates a DeleteAlertRuleDefault with default headers values
func NewDeleteAlertRuleDefault(code int) *DeleteAlertRuleDefault {
	return &DeleteAlertRuleDefault{
		_statusCode: code,
	}
}

/*DeleteAlertRuleDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteAlertRuleDefault struct {
	_statusCode int

	Payload *DeleteAlertRuleDefaultBody
}

// Code gets the status code for the delete alert rule default response
func (o *DeleteAlertRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAlertRuleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Delete][%d] DeleteAlertRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAlertRuleDefault) GetPayload() *DeleteAlertRuleDefaultBody {
	return o.Payload
}

func (o *DeleteAlertRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAlertRuleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteAlertRuleBody delete alert rule body
swagger:model DeleteAlertRuleBody
*/
type DeleteAlertRuleBody struct {

	// Rule ID.
	RuleID string `json:"rule_id,omitempty"`
}

// Validate validates this delete alert rule body
func (o *DeleteAlertRuleBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAlertRuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAlertRuleBody) UnmarshalBinary(b []byte) error {
	var res DeleteAlertRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteAlertRuleDefaultBody delete alert rule default body
swagger:model DeleteAlertRuleDefaultBody
*/
type DeleteAlertRuleDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this delete alert rule default body
func (o *DeleteAlertRuleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAlertRuleDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("DeleteAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAlertRuleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAlertRuleDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteAlertRuleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
