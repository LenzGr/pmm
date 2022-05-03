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

// AddRDSExporterReader is a Reader for the AddRDSExporter structure.
type AddRDSExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRDSExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRDSExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddRDSExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRDSExporterOK creates a AddRDSExporterOK with default headers values
func NewAddRDSExporterOK() *AddRDSExporterOK {
	return &AddRDSExporterOK{}
}

/* AddRDSExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddRDSExporterOK struct {
	Payload *AddRDSExporterOKBody
}

func (o *AddRDSExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddRDSExporter][%d] addRdsExporterOk  %+v", 200, o.Payload)
}
func (o *AddRDSExporterOK) GetPayload() *AddRDSExporterOKBody {
	return o.Payload
}

func (o *AddRDSExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRDSExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRDSExporterDefault creates a AddRDSExporterDefault with default headers values
func NewAddRDSExporterDefault(code int) *AddRDSExporterDefault {
	return &AddRDSExporterDefault{
		_statusCode: code,
	}
}

/* AddRDSExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddRDSExporterDefault struct {
	_statusCode int

	Payload *AddRDSExporterDefaultBody
}

// Code gets the status code for the add RDS exporter default response
func (o *AddRDSExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddRDSExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddRDSExporter][%d] AddRDSExporter default  %+v", o._statusCode, o.Payload)
}
func (o *AddRDSExporterDefault) GetPayload() *AddRDSExporterDefaultBody {
	return o.Payload
}

func (o *AddRDSExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRDSExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddRDSExporterBody add RDS exporter body
swagger:model AddRDSExporterBody
*/
type AddRDSExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Node identifier.
	NodeID string `json:"node_id,omitempty"`

	// AWS Access Key.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

	// AWS Secret Key.
	AWSSecretKey string `json:"aws_secret_key,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Disable basic metrics.
	DisableBasicMetrics bool `json:"disable_basic_metrics,omitempty"`

	// Disable enhanced metrics.
	DisableEnhancedMetrics bool `json:"disable_enhanced_metrics,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`
}

// Validate validates this add RDS exporter body
func (o *AddRDSExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add RDS exporter body based on context it is used
func (o *AddRDSExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterBody) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSExporterDefaultBody add RDS exporter default body
swagger:model AddRDSExporterDefaultBody
*/
type AddRDSExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddRDSExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add RDS exporter default body
func (o *AddRDSExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRDSExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddRDSExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddRDSExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add RDS exporter default body based on the context it is used
func (o *AddRDSExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRDSExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddRDSExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddRDSExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSExporterDefaultBodyDetailsItems0 add RDS exporter default body details items0
swagger:model AddRDSExporterDefaultBodyDetailsItems0
*/
type AddRDSExporterDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add RDS exporter default body details items0
func (o *AddRDSExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add RDS exporter default body details items0 based on context it is used
func (o *AddRDSExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSExporterOKBody add RDS exporter OK body
swagger:model AddRDSExporterOKBody
*/
type AddRDSExporterOKBody struct {

	// rds exporter
	RDSExporter *AddRDSExporterOKBodyRDSExporter `json:"rds_exporter,omitempty"`
}

// Validate validates this add RDS exporter OK body
func (o *AddRDSExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRDSExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRDSExporterOKBody) validateRDSExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.RDSExporter) { // not required
		return nil
	}

	if o.RDSExporter != nil {
		if err := o.RDSExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsExporterOk" + "." + "rds_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRdsExporterOk" + "." + "rds_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add RDS exporter OK body based on the context it is used
func (o *AddRDSExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRDSExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRDSExporterOKBody) contextValidateRDSExporter(ctx context.Context, formats strfmt.Registry) error {

	if o.RDSExporter != nil {
		if err := o.RDSExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsExporterOk" + "." + "rds_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addRdsExporterOk" + "." + "rds_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSExporterOKBodyRDSExporter RDSExporter runs on Generic or Container Node and exposes RemoteRDS Node metrics.
swagger:model AddRDSExporterOKBodyRDSExporter
*/
type AddRDSExporterOKBodyRDSExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Node identifier.
	NodeID string `json:"node_id,omitempty"`

	// AWS Access Key.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

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

	// Listen port for scraping metrics (the same for several configurations).
	ListenPort int64 `json:"listen_port,omitempty"`

	// Basic metrics are disabled.
	BasicMetricsDisabled bool `json:"basic_metrics_disabled,omitempty"`

	// Enhanced metrics are disabled.
	EnhancedMetricsDisabled bool `json:"enhanced_metrics_disabled,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`
}

// Validate validates this add RDS exporter OK body RDS exporter
func (o *AddRDSExporterOKBodyRDSExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addRdsExporterOkBodyRdsExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addRdsExporterOkBodyRdsExporterTypeStatusPropEnum = append(addRdsExporterOkBodyRdsExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddRDSExporterOKBodyRDSExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddRDSExporterOKBodyRDSExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddRDSExporterOKBodyRDSExporterStatusSTARTING captures enum value "STARTING"
	AddRDSExporterOKBodyRDSExporterStatusSTARTING string = "STARTING"

	// AddRDSExporterOKBodyRDSExporterStatusRUNNING captures enum value "RUNNING"
	AddRDSExporterOKBodyRDSExporterStatusRUNNING string = "RUNNING"

	// AddRDSExporterOKBodyRDSExporterStatusWAITING captures enum value "WAITING"
	AddRDSExporterOKBodyRDSExporterStatusWAITING string = "WAITING"

	// AddRDSExporterOKBodyRDSExporterStatusSTOPPING captures enum value "STOPPING"
	AddRDSExporterOKBodyRDSExporterStatusSTOPPING string = "STOPPING"

	// AddRDSExporterOKBodyRDSExporterStatusDONE captures enum value "DONE"
	AddRDSExporterOKBodyRDSExporterStatusDONE string = "DONE"

	// AddRDSExporterOKBodyRDSExporterStatusUNKNOWN captures enum value "UNKNOWN"
	AddRDSExporterOKBodyRDSExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddRDSExporterOKBodyRDSExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addRdsExporterOkBodyRdsExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddRDSExporterOKBodyRDSExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addRdsExporterOk"+"."+"rds_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add RDS exporter OK body RDS exporter based on context it is used
func (o *AddRDSExporterOKBodyRDSExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSExporterOKBodyRDSExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSExporterOKBodyRDSExporter) UnmarshalBinary(b []byte) error {
	var res AddRDSExporterOKBodyRDSExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
