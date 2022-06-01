// Code generated by go-swagger; DO NOT EDIT.

package templates

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

// UpdateTemplateReader is a Reader for the UpdateTemplate structure.
type UpdateTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTemplateOK creates a UpdateTemplateOK with default headers values
func NewUpdateTemplateOK() *UpdateTemplateOK {
	return &UpdateTemplateOK{}
}

/* UpdateTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateTemplateOK struct {
	Payload interface{}
}

func (o *UpdateTemplateOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Templates/Update][%d] updateTemplateOk  %+v", 200, o.Payload)
}

func (o *UpdateTemplateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTemplateDefault creates a UpdateTemplateDefault with default headers values
func NewUpdateTemplateDefault(code int) *UpdateTemplateDefault {
	return &UpdateTemplateDefault{
		_statusCode: code,
	}
}

/* UpdateTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateTemplateDefault struct {
	_statusCode int

	Payload *UpdateTemplateDefaultBody
}

// Code gets the status code for the update template default response
func (o *UpdateTemplateDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Templates/Update][%d] UpdateTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTemplateDefault) GetPayload() *UpdateTemplateDefaultBody {
	return o.Payload
}

func (o *UpdateTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UpdateTemplateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateTemplateBody update template body
swagger:model UpdateTemplateBody
*/
type UpdateTemplateBody struct {
	// Machine-readable name (ID).
	Name string `json:"name,omitempty"`

	// YAML (or JSON) template file content.
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this update template body
func (o *UpdateTemplateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update template body based on context it is used
func (o *UpdateTemplateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateTemplateBody) UnmarshalBinary(b []byte) error {
	var res UpdateTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateTemplateDefaultBody update template default body
swagger:model UpdateTemplateDefaultBody
*/
type UpdateTemplateDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UpdateTemplateDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this update template default body
func (o *UpdateTemplateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateTemplateDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UpdateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update template default body based on the context it is used
func (o *UpdateTemplateDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateTemplateDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateTemplateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateTemplateDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateTemplateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateTemplateDefaultBodyDetailsItems0 update template default body details items0
swagger:model UpdateTemplateDefaultBodyDetailsItems0
*/
type UpdateTemplateDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this update template default body details items0
func (o *UpdateTemplateDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update template default body details items0 based on context it is used
func (o *UpdateTemplateDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateTemplateDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateTemplateDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateTemplateDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
