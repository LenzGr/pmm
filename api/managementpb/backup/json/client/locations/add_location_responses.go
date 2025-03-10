// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// AddLocationReader is a Reader for the AddLocation structure.
type AddLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddLocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddLocationOK creates a AddLocationOK with default headers values
func NewAddLocationOK() *AddLocationOK {
	return &AddLocationOK{}
}

/* AddLocationOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddLocationOK struct {
	Payload *AddLocationOKBody
}

func (o *AddLocationOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/Add][%d] addLocationOk  %+v", 200, o.Payload)
}

func (o *AddLocationOK) GetPayload() *AddLocationOKBody {
	return o.Payload
}

func (o *AddLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddLocationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLocationDefault creates a AddLocationDefault with default headers values
func NewAddLocationDefault(code int) *AddLocationDefault {
	return &AddLocationDefault{
		_statusCode: code,
	}
}

/* AddLocationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddLocationDefault struct {
	_statusCode int

	Payload *AddLocationDefaultBody
}

// Code gets the status code for the add location default response
func (o *AddLocationDefault) Code() int {
	return o._statusCode
}

func (o *AddLocationDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/Add][%d] AddLocation default  %+v", o._statusCode, o.Payload)
}

func (o *AddLocationDefault) GetPayload() *AddLocationDefaultBody {
	return o.Payload
}

func (o *AddLocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddLocationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddLocationBody add location body
swagger:model AddLocationBody
*/
type AddLocationBody struct {
	// Location name
	Name string `json:"name,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// pmm client config
	PMMClientConfig *AddLocationParamsBodyPMMClientConfig `json:"pmm_client_config,omitempty"`

	// pmm server config
	PMMServerConfig *AddLocationParamsBodyPMMServerConfig `json:"pmm_server_config,omitempty"`

	// s3 config
	S3Config *AddLocationParamsBodyS3Config `json:"s3_config,omitempty"`
}

// Validate validates this add location body
func (o *AddLocationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePMMClientConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateS3Config(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddLocationBody) validatePMMClientConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.PMMClientConfig) { // not required
		return nil
	}

	if o.PMMClientConfig != nil {
		if err := o.PMMClientConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_client_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_client_config")
			}
			return err
		}
	}

	return nil
}

func (o *AddLocationBody) validatePMMServerConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.PMMServerConfig) { // not required
		return nil
	}

	if o.PMMServerConfig != nil {
		if err := o.PMMServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_server_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_server_config")
			}
			return err
		}
	}

	return nil
}

func (o *AddLocationBody) validateS3Config(formats strfmt.Registry) error {
	if swag.IsZero(o.S3Config) { // not required
		return nil
	}

	if o.S3Config != nil {
		if err := o.S3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "s3_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "s3_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add location body based on the context it is used
func (o *AddLocationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePMMClientConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePMMServerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateS3Config(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddLocationBody) contextValidatePMMClientConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.PMMClientConfig != nil {
		if err := o.PMMClientConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_client_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_client_config")
			}
			return err
		}
	}

	return nil
}

func (o *AddLocationBody) contextValidatePMMServerConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.PMMServerConfig != nil {
		if err := o.PMMServerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_server_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pmm_server_config")
			}
			return err
		}
	}

	return nil
}

func (o *AddLocationBody) contextValidateS3Config(ctx context.Context, formats strfmt.Registry) error {
	if o.S3Config != nil {
		if err := o.S3Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "s3_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "s3_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationBody) UnmarshalBinary(b []byte) error {
	var res AddLocationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddLocationDefaultBody add location default body
swagger:model AddLocationDefaultBody
*/
type AddLocationDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddLocationDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add location default body
func (o *AddLocationDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddLocationDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddLocation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddLocation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add location default body based on the context it is used
func (o *AddLocationDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddLocationDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddLocation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddLocation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddLocationDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddLocationDefaultBodyDetailsItems0 add location default body details items0
swagger:model AddLocationDefaultBodyDetailsItems0
*/
type AddLocationDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this add location default body details items0
func (o *AddLocationDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add location default body details items0 based on context it is used
func (o *AddLocationDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddLocationDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddLocationOKBody add location OK body
swagger:model AddLocationOKBody
*/
type AddLocationOKBody struct {
	// Machine-readable ID.
	LocationID string `json:"location_id,omitempty"`
}

// Validate validates this add location OK body
func (o *AddLocationOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add location OK body based on context it is used
func (o *AddLocationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationOKBody) UnmarshalBinary(b []byte) error {
	var res AddLocationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddLocationParamsBodyPMMClientConfig PMMClientLocationConfig represents file system config inside pmm-client.
swagger:model AddLocationParamsBodyPMMClientConfig
*/
type AddLocationParamsBodyPMMClientConfig struct {
	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this add location params body PMM client config
func (o *AddLocationParamsBodyPMMClientConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add location params body PMM client config based on context it is used
func (o *AddLocationParamsBodyPMMClientConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationParamsBodyPMMClientConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationParamsBodyPMMClientConfig) UnmarshalBinary(b []byte) error {
	var res AddLocationParamsBodyPMMClientConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddLocationParamsBodyPMMServerConfig PMMServerLocationConfig represents file system config inside pmm-server.
swagger:model AddLocationParamsBodyPMMServerConfig
*/
type AddLocationParamsBodyPMMServerConfig struct {
	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this add location params body PMM server config
func (o *AddLocationParamsBodyPMMServerConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add location params body PMM server config based on context it is used
func (o *AddLocationParamsBodyPMMServerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationParamsBodyPMMServerConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationParamsBodyPMMServerConfig) UnmarshalBinary(b []byte) error {
	var res AddLocationParamsBodyPMMServerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddLocationParamsBodyS3Config S3LocationConfig represents S3 bucket configuration.
swagger:model AddLocationParamsBodyS3Config
*/
type AddLocationParamsBodyS3Config struct {
	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// access key
	AccessKey string `json:"access_key,omitempty"`

	// secret key
	SecretKey string `json:"secret_key,omitempty"`

	// bucket name
	BucketName string `json:"bucket_name,omitempty"`
}

// Validate validates this add location params body s3 config
func (o *AddLocationParamsBodyS3Config) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add location params body s3 config based on context it is used
func (o *AddLocationParamsBodyS3Config) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationParamsBodyS3Config) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationParamsBodyS3Config) UnmarshalBinary(b []byte) error {
	var res AddLocationParamsBodyS3Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
