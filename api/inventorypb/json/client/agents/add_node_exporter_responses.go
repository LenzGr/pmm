// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

// AddNodeExporterReader is a Reader for the AddNodeExporter structure.
type AddNodeExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNodeExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddNodeExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddNodeExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddNodeExporterOK creates a AddNodeExporterOK with default headers values
func NewAddNodeExporterOK() *AddNodeExporterOK {
	return &AddNodeExporterOK{}
}

/*AddNodeExporterOK handles this case with default header values.

A successful response.
*/
type AddNodeExporterOK struct {
	Payload *AddNodeExporterOKBody
}

func (o *AddNodeExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddNodeExporter][%d] addNodeExporterOk  %+v", 200, o.Payload)
}

func (o *AddNodeExporterOK) GetPayload() *AddNodeExporterOKBody {
	return o.Payload
}

func (o *AddNodeExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddNodeExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNodeExporterDefault creates a AddNodeExporterDefault with default headers values
func NewAddNodeExporterDefault(code int) *AddNodeExporterDefault {
	return &AddNodeExporterDefault{
		_statusCode: code,
	}
}

/*AddNodeExporterDefault handles this case with default header values.

An unexpected error response.
*/
type AddNodeExporterDefault struct {
	_statusCode int

	Payload *AddNodeExporterDefaultBody
}

// Code gets the status code for the add node exporter default response
func (o *AddNodeExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddNodeExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddNodeExporter][%d] AddNodeExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddNodeExporterDefault) GetPayload() *AddNodeExporterDefaultBody {
	return o.Payload
}

func (o *AddNodeExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddNodeExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddNodeExporterBody add node exporter body
swagger:model AddNodeExporterBody
*/
type AddNodeExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`
}

// Validate validates this add node exporter body
func (o *AddNodeExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeExporterBody) UnmarshalBinary(b []byte) error {
	var res AddNodeExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddNodeExporterDefaultBody add node exporter default body
swagger:model AddNodeExporterDefaultBody
*/
type AddNodeExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add node exporter default body
func (o *AddNodeExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddNodeExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddNodeExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddNodeExporterOKBody add node exporter OK body
swagger:model AddNodeExporterOKBody
*/
type AddNodeExporterOKBody struct {

	// node exporter
	NodeExporter *AddNodeExporterOKBodyNodeExporter `json:"node_exporter,omitempty"`
}

// Validate validates this add node exporter OK body
func (o *AddNodeExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddNodeExporterOKBody) validateNodeExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeExporter) { // not required
		return nil
	}

	if o.NodeExporter != nil {
		if err := o.NodeExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addNodeExporterOk" + "." + "node_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddNodeExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddNodeExporterOKBodyNodeExporter NodeExporter runs on Generic or Container Node and exposes its metrics.
swagger:model AddNodeExporterOKBodyNodeExporter
*/
type AddNodeExporterOKBodyNodeExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

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

// Validate validates this add node exporter OK body node exporter
func (o *AddNodeExporterOKBodyNodeExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addNodeExporterOkBodyNodeExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addNodeExporterOkBodyNodeExporterTypeStatusPropEnum = append(addNodeExporterOkBodyNodeExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddNodeExporterOKBodyNodeExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddNodeExporterOKBodyNodeExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddNodeExporterOKBodyNodeExporterStatusSTARTING captures enum value "STARTING"
	AddNodeExporterOKBodyNodeExporterStatusSTARTING string = "STARTING"

	// AddNodeExporterOKBodyNodeExporterStatusRUNNING captures enum value "RUNNING"
	AddNodeExporterOKBodyNodeExporterStatusRUNNING string = "RUNNING"

	// AddNodeExporterOKBodyNodeExporterStatusWAITING captures enum value "WAITING"
	AddNodeExporterOKBodyNodeExporterStatusWAITING string = "WAITING"

	// AddNodeExporterOKBodyNodeExporterStatusSTOPPING captures enum value "STOPPING"
	AddNodeExporterOKBodyNodeExporterStatusSTOPPING string = "STOPPING"

	// AddNodeExporterOKBodyNodeExporterStatusDONE captures enum value "DONE"
	AddNodeExporterOKBodyNodeExporterStatusDONE string = "DONE"

	// AddNodeExporterOKBodyNodeExporterStatusUNKNOWN captures enum value "UNKNOWN"
	AddNodeExporterOKBodyNodeExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddNodeExporterOKBodyNodeExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addNodeExporterOkBodyNodeExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddNodeExporterOKBodyNodeExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addNodeExporterOk"+"."+"node_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddNodeExporterOKBodyNodeExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddNodeExporterOKBodyNodeExporter) UnmarshalBinary(b []byte) error {
	var res AddNodeExporterOKBodyNodeExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
