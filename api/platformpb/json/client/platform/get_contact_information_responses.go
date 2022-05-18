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

// GetContactInformationReader is a Reader for the GetContactInformation structure.
type GetContactInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContactInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContactInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetContactInformationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetContactInformationOK creates a GetContactInformationOK with default headers values
func NewGetContactInformationOK() *GetContactInformationOK {
	return &GetContactInformationOK{}
}

/* GetContactInformationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetContactInformationOK struct {
	Payload *GetContactInformationOKBody
}

func (o *GetContactInformationOK) Error() string {
	return fmt.Sprintf("[POST /v1/Platform/GetContactInformation][%d] getContactInformationOk  %+v", 200, o.Payload)
}
func (o *GetContactInformationOK) GetPayload() *GetContactInformationOKBody {
	return o.Payload
}

func (o *GetContactInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetContactInformationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContactInformationDefault creates a GetContactInformationDefault with default headers values
func NewGetContactInformationDefault(code int) *GetContactInformationDefault {
	return &GetContactInformationDefault{
		_statusCode: code,
	}
}

/* GetContactInformationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetContactInformationDefault struct {
	_statusCode int

	Payload *GetContactInformationDefaultBody
}

// Code gets the status code for the get contact information default response
func (o *GetContactInformationDefault) Code() int {
	return o._statusCode
}

func (o *GetContactInformationDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Platform/GetContactInformation][%d] GetContactInformation default  %+v", o._statusCode, o.Payload)
}
func (o *GetContactInformationDefault) GetPayload() *GetContactInformationDefaultBody {
	return o.Payload
}

func (o *GetContactInformationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetContactInformationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetContactInformationDefaultBody get contact information default body
swagger:model GetContactInformationDefaultBody
*/
type GetContactInformationDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetContactInformationDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get contact information default body
func (o *GetContactInformationDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContactInformationDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetContactInformation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetContactInformation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get contact information default body based on the context it is used
func (o *GetContactInformationDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContactInformationDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetContactInformation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetContactInformation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetContactInformationDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContactInformationDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetContactInformationDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetContactInformationDefaultBodyDetailsItems0 get contact information default body details items0
swagger:model GetContactInformationDefaultBodyDetailsItems0
*/
type GetContactInformationDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get contact information default body details items0
func (o *GetContactInformationDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get contact information default body details items0 based on context it is used
func (o *GetContactInformationDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetContactInformationDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContactInformationDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetContactInformationDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetContactInformationOKBody get contact information OK body
swagger:model GetContactInformationOKBody
*/
type GetContactInformationOKBody struct {

	// URL to open a new support ticket.
	NewTicketURL string `json:"new_ticket_url,omitempty"`

	// customer success
	CustomerSuccess *GetContactInformationOKBodyCustomerSuccess `json:"customer_success,omitempty"`
}

// Validate validates this get contact information OK body
func (o *GetContactInformationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomerSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContactInformationOKBody) validateCustomerSuccess(formats strfmt.Registry) error {
	if swag.IsZero(o.CustomerSuccess) { // not required
		return nil
	}

	if o.CustomerSuccess != nil {
		if err := o.CustomerSuccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getContactInformationOk" + "." + "customer_success")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getContactInformationOk" + "." + "customer_success")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get contact information OK body based on the context it is used
func (o *GetContactInformationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCustomerSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetContactInformationOKBody) contextValidateCustomerSuccess(ctx context.Context, formats strfmt.Registry) error {

	if o.CustomerSuccess != nil {
		if err := o.CustomerSuccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getContactInformationOk" + "." + "customer_success")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getContactInformationOk" + "." + "customer_success")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetContactInformationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContactInformationOKBody) UnmarshalBinary(b []byte) error {
	var res GetContactInformationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetContactInformationOKBodyCustomerSuccess CustomerSuccess contains the contanct details of the customer success employee assigned to a customer's account.
swagger:model GetContactInformationOKBodyCustomerSuccess
*/
type GetContactInformationOKBodyCustomerSuccess struct {

	// name
	Name string `json:"name,omitempty"`

	// email
	Email string `json:"email,omitempty"`
}

// Validate validates this get contact information OK body customer success
func (o *GetContactInformationOKBodyCustomerSuccess) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get contact information OK body customer success based on context it is used
func (o *GetContactInformationOKBodyCustomerSuccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetContactInformationOKBodyCustomerSuccess) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContactInformationOKBodyCustomerSuccess) UnmarshalBinary(b []byte) error {
	var res GetContactInformationOKBodyCustomerSuccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
