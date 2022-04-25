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

// AddRemoteAzureDatabaseNodeReader is a Reader for the AddRemoteAzureDatabaseNode structure.
type AddRemoteAzureDatabaseNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRemoteAzureDatabaseNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRemoteAzureDatabaseNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddRemoteAzureDatabaseNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRemoteAzureDatabaseNodeOK creates a AddRemoteAzureDatabaseNodeOK with default headers values
func NewAddRemoteAzureDatabaseNodeOK() *AddRemoteAzureDatabaseNodeOK {
	return &AddRemoteAzureDatabaseNodeOK{}
}

/* AddRemoteAzureDatabaseNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddRemoteAzureDatabaseNodeOK struct {
	Payload *AddRemoteAzureDatabaseNodeOKBody
}

func (o *AddRemoteAzureDatabaseNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddRemoteAzureDatabase][%d] addRemoteAzureDatabaseNodeOk  %+v", 200, o.Payload)
}
func (o *AddRemoteAzureDatabaseNodeOK) GetPayload() *AddRemoteAzureDatabaseNodeOKBody {
	return o.Payload
}

func (o *AddRemoteAzureDatabaseNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRemoteAzureDatabaseNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRemoteAzureDatabaseNodeDefault creates a AddRemoteAzureDatabaseNodeDefault with default headers values
func NewAddRemoteAzureDatabaseNodeDefault(code int) *AddRemoteAzureDatabaseNodeDefault {
	return &AddRemoteAzureDatabaseNodeDefault{
		_statusCode: code,
	}
}

/* AddRemoteAzureDatabaseNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddRemoteAzureDatabaseNodeDefault struct {
	_statusCode int

	Payload *AddRemoteAzureDatabaseNodeDefaultBody
}

// Code gets the status code for the add remote azure database node default response
func (o *AddRemoteAzureDatabaseNodeDefault) Code() int {
	return o._statusCode
}

func (o *AddRemoteAzureDatabaseNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddRemoteAzureDatabase][%d] AddRemoteAzureDatabaseNode default  %+v", o._statusCode, o.Payload)
}
func (o *AddRemoteAzureDatabaseNodeDefault) GetPayload() *AddRemoteAzureDatabaseNodeDefaultBody {
	return o.Payload
}

func (o *AddRemoteAzureDatabaseNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRemoteAzureDatabaseNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddRemoteAzureDatabaseNodeBody add remote azure database node body
swagger:model AddRemoteAzureDatabaseNodeBody
*/
type AddRemoteAzureDatabaseNodeBody struct {

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add remote azure database node body
func (o *AddRemoteAzureDatabaseNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add remote azure database node body based on context it is used
func (o *AddRemoteAzureDatabaseNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteAzureDatabaseNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteAzureDatabaseNodeDefaultBody add remote azure database node default body
swagger:model AddRemoteAzureDatabaseNodeDefaultBody
*/
type AddRemoteAzureDatabaseNodeDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add remote azure database node default body
func (o *AddRemoteAzureDatabaseNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRemoteAzureDatabaseNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddRemoteAzureDatabaseNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddRemoteAzureDatabaseNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add remote azure database node default body based on the context it is used
func (o *AddRemoteAzureDatabaseNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRemoteAzureDatabaseNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddRemoteAzureDatabaseNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddRemoteAzureDatabaseNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteAzureDatabaseNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0 add remote azure database node default body details items0
swagger:model AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0
*/
type AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this add remote azure database node default body details items0
func (o *AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add remote azure database node default body details items0 based on context it is used
func (o *AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddRemoteAzureDatabaseNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteAzureDatabaseNodeOKBody add remote azure database node OK body
swagger:model AddRemoteAzureDatabaseNodeOKBody
*/
type AddRemoteAzureDatabaseNodeOKBody struct {

	// remote azure database
	RemoteAzureDatabase *AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase `json:"remote_azure_database,omitempty"`
}

// Validate validates this add remote azure database node OK body
func (o *AddRemoteAzureDatabaseNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRemoteAzureDatabase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRemoteAzureDatabaseNodeOKBody) validateRemoteAzureDatabase(formats strfmt.Registry) error {
	if swag.IsZero(o.RemoteAzureDatabase) { // not required
		return nil
	}

	if o.RemoteAzureDatabase != nil {
		if err := o.RemoteAzureDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRemoteAzureDatabaseNodeOk" + "." + "remote_azure_database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRemoteAzureDatabaseNodeOk" + "." + "remote_azure_database")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add remote azure database node OK body based on the context it is used
func (o *AddRemoteAzureDatabaseNodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRemoteAzureDatabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRemoteAzureDatabaseNodeOKBody) contextValidateRemoteAzureDatabase(ctx context.Context, formats strfmt.Registry) error {

	if o.RemoteAzureDatabase != nil {
		if err := o.RemoteAzureDatabase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRemoteAzureDatabaseNodeOk" + "." + "remote_azure_database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRemoteAzureDatabaseNodeOk" + "." + "remote_azure_database")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeOKBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteAzureDatabaseNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase RemoteAzureDatabaseNode represents remote AzureDatabase Node. Agents can't run on Remote AzureDatabase Nodes.
swagger:model AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase
*/
type AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add remote azure database node OK body remote azure database
func (o *AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add remote azure database node OK body remote azure database based on context it is used
func (o *AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase) UnmarshalBinary(b []byte) error {
	var res AddRemoteAzureDatabaseNodeOKBodyRemoteAzureDatabase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
