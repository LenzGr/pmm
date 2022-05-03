// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// RemoveNodeReader is a Reader for the RemoveNode structure.
type RemoveNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveNodeOK creates a RemoveNodeOK with default headers values
func NewRemoveNodeOK() *RemoveNodeOK {
	return &RemoveNodeOK{}
}

/* RemoveNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveNodeOK struct {
	Payload interface{}
}

func (o *RemoveNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Remove][%d] removeNodeOk  %+v", 200, o.Payload)
}
func (o *RemoveNodeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveNodeDefault creates a RemoveNodeDefault with default headers values
func NewRemoveNodeDefault(code int) *RemoveNodeDefault {
	return &RemoveNodeDefault{
		_statusCode: code,
	}
}

/* RemoveNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveNodeDefault struct {
	_statusCode int

	Payload *RemoveNodeDefaultBody
}

// Code gets the status code for the remove node default response
func (o *RemoveNodeDefault) Code() int {
	return o._statusCode
}

func (o *RemoveNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Remove][%d] RemoveNode default  %+v", o._statusCode, o.Payload)
}
func (o *RemoveNodeDefault) GetPayload() *RemoveNodeDefaultBody {
	return o.Payload
}

func (o *RemoveNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveNodeBody remove node body
swagger:model RemoveNodeBody
*/
type RemoveNodeBody struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Remove node with all dependencies.
	Force bool `json:"force,omitempty"`
}

// Validate validates this remove node body
func (o *RemoveNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove node body based on context it is used
func (o *RemoveNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveNodeBody) UnmarshalBinary(b []byte) error {
	var res RemoveNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveNodeDefaultBody remove node default body
swagger:model RemoveNodeDefaultBody
*/
type RemoveNodeDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RemoveNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this remove node default body
func (o *RemoveNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remove node default body based on the context it is used
func (o *RemoveNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveNodeDefaultBodyDetailsItems0 remove node default body details items0
swagger:model RemoveNodeDefaultBodyDetailsItems0
*/
type RemoveNodeDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this remove node default body details items0
func (o *RemoveNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove node default body details items0 based on context it is used
func (o *RemoveNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RemoveNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
