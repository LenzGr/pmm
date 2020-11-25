// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// DeleteAlertReader is a Reader for the DeleteAlert structure.
type DeleteAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteAlertDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAlertOK creates a DeleteAlertOK with default headers values
func NewDeleteAlertOK() *DeleteAlertOK {
	return &DeleteAlertOK{}
}

/*DeleteAlertOK handles this case with default header values.

A successful response.
*/
type DeleteAlertOK struct {
	Payload interface{}
}

func (o *DeleteAlertOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Alerts/Delete][%d] deleteAlertOk  %+v", 200, o.Payload)
}

func (o *DeleteAlertOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefault creates a DeleteAlertDefault with default headers values
func NewDeleteAlertDefault(code int) *DeleteAlertDefault {
	return &DeleteAlertDefault{
		_statusCode: code,
	}
}

/*DeleteAlertDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteAlertDefault struct {
	_statusCode int

	Payload *DeleteAlertDefaultBody
}

// Code gets the status code for the delete alert default response
func (o *DeleteAlertDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAlertDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Alerts/Delete][%d] DeleteAlert default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAlertDefault) GetPayload() *DeleteAlertDefaultBody {
	return o.Payload
}

func (o *DeleteAlertDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteAlertDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteAlertBody delete alert body
swagger:model DeleteAlertBody
*/
type DeleteAlertBody struct {

	// ID.
	AlertID string `json:"alert_id,omitempty"`
}

// Validate validates this delete alert body
func (o *DeleteAlertBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAlertBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAlertBody) UnmarshalBinary(b []byte) error {
	var res DeleteAlertBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteAlertDefaultBody delete alert default body
swagger:model DeleteAlertDefaultBody
*/
type DeleteAlertDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this delete alert default body
func (o *DeleteAlertDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteAlertDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("DeleteAlert default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteAlertDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteAlertDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteAlertDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}