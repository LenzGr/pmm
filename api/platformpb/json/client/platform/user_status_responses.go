// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// UserStatusReader is a Reader for the UserStatus structure.
type UserStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserStatusOK creates a UserStatusOK with default headers values
func NewUserStatusOK() *UserStatusOK {
	return &UserStatusOK{}
}

/* UserStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type UserStatusOK struct {
	Payload *UserStatusOKBody
}

func (o *UserStatusOK) Error() string {
	return fmt.Sprintf("[POST /v1/Platform/UserStatus][%d] userStatusOk  %+v", 200, o.Payload)
}

func (o *UserStatusOK) GetPayload() *UserStatusOKBody {
	return o.Payload
}

func (o *UserStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UserStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserStatusDefault creates a UserStatusDefault with default headers values
func NewUserStatusDefault(code int) *UserStatusDefault {
	return &UserStatusDefault{
		_statusCode: code,
	}
}

/* UserStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UserStatusDefault struct {
	_statusCode int

	Payload *UserStatusDefaultBody
}

// Code gets the status code for the user status default response
func (o *UserStatusDefault) Code() int {
	return o._statusCode
}

func (o *UserStatusDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Platform/UserStatus][%d] UserStatus default  %+v", o._statusCode, o.Payload)
}

func (o *UserStatusDefault) GetPayload() *UserStatusDefaultBody {
	return o.Payload
}

func (o *UserStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UserStatusDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UserStatusDefaultBody user status default body
swagger:model UserStatusDefaultBody
*/
type UserStatusDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UserStatusDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this user status default body
func (o *UserStatusDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserStatusDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UserStatus default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UserStatus default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this user status default body based on the context it is used
func (o *UserStatusDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserStatusDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UserStatus default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UserStatus default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserStatusDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserStatusDefaultBody) UnmarshalBinary(b []byte) error {
	var res UserStatusDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UserStatusDefaultBodyDetailsItems0 user status default body details items0
swagger:model UserStatusDefaultBodyDetailsItems0
*/
type UserStatusDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this user status default body details items0
func (o *UserStatusDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user status default body details items0 based on context it is used
func (o *UserStatusDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UserStatusDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserStatusDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UserStatusDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UserStatusOKBody user status OK body
swagger:model UserStatusOKBody
*/
type UserStatusOKBody struct {
	// is platform user
	IsPlatformUser bool `json:"is_platform_user,omitempty"`
}

// Validate validates this user status OK body
func (o *UserStatusOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user status OK body based on context it is used
func (o *UserStatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UserStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserStatusOKBody) UnmarshalBinary(b []byte) error {
	var res UserStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
