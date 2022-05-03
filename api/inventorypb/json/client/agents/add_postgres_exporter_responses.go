// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddPostgresExporterReader is a Reader for the AddPostgresExporter structure.
type AddPostgresExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPostgresExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPostgresExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddPostgresExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPostgresExporterOK creates a AddPostgresExporterOK with default headers values
func NewAddPostgresExporterOK() *AddPostgresExporterOK {
	return &AddPostgresExporterOK{}
}

/* AddPostgresExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddPostgresExporterOK struct {
	Payload *AddPostgresExporterOKBody
}

func (o *AddPostgresExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddPostgresExporter][%d] addPostgresExporterOk  %+v", 200, o.Payload)
}
func (o *AddPostgresExporterOK) GetPayload() *AddPostgresExporterOKBody {
	return o.Payload
}

func (o *AddPostgresExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgresExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPostgresExporterDefault creates a AddPostgresExporterDefault with default headers values
func NewAddPostgresExporterDefault(code int) *AddPostgresExporterDefault {
	return &AddPostgresExporterDefault{
		_statusCode: code,
	}
}

/* AddPostgresExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddPostgresExporterDefault struct {
	_statusCode int

	Payload *AddPostgresExporterDefaultBody
}

// Code gets the status code for the add postgres exporter default response
func (o *AddPostgresExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddPostgresExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddPostgresExporter][%d] AddPostgresExporter default  %+v", o._statusCode, o.Payload)
}
func (o *AddPostgresExporterDefault) GetPayload() *AddPostgresExporterDefaultBody {
	return o.Payload
}

func (o *AddPostgresExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgresExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddPostgresExporterBody add postgres exporter body
swagger:model AddPostgresExporterBody
*/
type AddPostgresExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// PostgreSQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// TLS CA certificate.
	TLSCa string `json:"tls_ca,omitempty"`

	// TLS Certifcate.
	TLSCert string `json:"tls_cert,omitempty"`

	// TLS Certificate Key.
	TLSKey string `json:"tls_key,omitempty"`

	// Custom password for exporter endpoint /metrics.
	AgentPassword string `json:"agent_password,omitempty"`
}

// Validate validates this add postgres exporter body
func (o *AddPostgresExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add postgres exporter body based on context it is used
func (o *AddPostgresExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgresExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgresExporterBody) UnmarshalBinary(b []byte) error {
	var res AddPostgresExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgresExporterDefaultBody add postgres exporter default body
swagger:model AddPostgresExporterDefaultBody
*/
type AddPostgresExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddPostgresExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add postgres exporter default body
func (o *AddPostgresExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgresExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddPostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddPostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add postgres exporter default body based on the context it is used
func (o *AddPostgresExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgresExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddPostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddPostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgresExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgresExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddPostgresExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgresExporterDefaultBodyDetailsItems0 add postgres exporter default body details items0
swagger:model AddPostgresExporterDefaultBodyDetailsItems0
*/
type AddPostgresExporterDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add postgres exporter default body details items0
func (o *AddPostgresExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add postgres exporter default body details items0 based on context it is used
func (o *AddPostgresExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgresExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgresExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddPostgresExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgresExporterOKBody add postgres exporter OK body
swagger:model AddPostgresExporterOKBody
*/
type AddPostgresExporterOKBody struct {

	// postgres exporter
	PostgresExporter *AddPostgresExporterOKBodyPostgresExporter `json:"postgres_exporter,omitempty"`
}

// Validate validates this add postgres exporter OK body
func (o *AddPostgresExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePostgresExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgresExporterOKBody) validatePostgresExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.PostgresExporter) { // not required
		return nil
	}

	if o.PostgresExporter != nil {
		if err := o.PostgresExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgresExporterOk" + "." + "postgres_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addPostgresExporterOk" + "." + "postgres_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add postgres exporter OK body based on the context it is used
func (o *AddPostgresExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePostgresExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgresExporterOKBody) contextValidatePostgresExporter(ctx context.Context, formats strfmt.Registry) error {

	if o.PostgresExporter != nil {
		if err := o.PostgresExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgresExporterOk" + "." + "postgres_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addPostgresExporterOk" + "." + "postgres_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgresExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgresExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddPostgresExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgresExporterOKBodyPostgresExporter PostgresExporter runs on Generic or Container Node and exposes PostgreSQL Service metrics.
swagger:model AddPostgresExporterOKBodyPostgresExporter
*/
type AddPostgresExporterOKBodyPostgresExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	DisabledCollectors []string `json:"disabled_collectors"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this add postgres exporter OK body postgres exporter
func (o *AddPostgresExporterOKBodyPostgresExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgresExporterOkBodyPostgresExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgresExporterOkBodyPostgresExporterTypeStatusPropEnum = append(addPostgresExporterOkBodyPostgresExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddPostgresExporterOKBodyPostgresExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddPostgresExporterOKBodyPostgresExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddPostgresExporterOKBodyPostgresExporterStatusSTARTING captures enum value "STARTING"
	AddPostgresExporterOKBodyPostgresExporterStatusSTARTING string = "STARTING"

	// AddPostgresExporterOKBodyPostgresExporterStatusRUNNING captures enum value "RUNNING"
	AddPostgresExporterOKBodyPostgresExporterStatusRUNNING string = "RUNNING"

	// AddPostgresExporterOKBodyPostgresExporterStatusWAITING captures enum value "WAITING"
	AddPostgresExporterOKBodyPostgresExporterStatusWAITING string = "WAITING"

	// AddPostgresExporterOKBodyPostgresExporterStatusSTOPPING captures enum value "STOPPING"
	AddPostgresExporterOKBodyPostgresExporterStatusSTOPPING string = "STOPPING"

	// AddPostgresExporterOKBodyPostgresExporterStatusDONE captures enum value "DONE"
	AddPostgresExporterOKBodyPostgresExporterStatusDONE string = "DONE"

	// AddPostgresExporterOKBodyPostgresExporterStatusUNKNOWN captures enum value "UNKNOWN"
	AddPostgresExporterOKBodyPostgresExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddPostgresExporterOKBodyPostgresExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addPostgresExporterOkBodyPostgresExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgresExporterOKBodyPostgresExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addPostgresExporterOk"+"."+"postgres_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add postgres exporter OK body postgres exporter based on context it is used
func (o *AddPostgresExporterOKBodyPostgresExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgresExporterOKBodyPostgresExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgresExporterOKBodyPostgresExporter) UnmarshalBinary(b []byte) error {
	var res AddPostgresExporterOKBodyPostgresExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
