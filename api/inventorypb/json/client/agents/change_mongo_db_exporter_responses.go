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

// ChangeMongoDBExporterReader is a Reader for the ChangeMongoDBExporter structure.
type ChangeMongoDBExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeMongoDBExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeMongoDBExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeMongoDBExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeMongoDBExporterOK creates a ChangeMongoDBExporterOK with default headers values
func NewChangeMongoDBExporterOK() *ChangeMongoDBExporterOK {
	return &ChangeMongoDBExporterOK{}
}

/* ChangeMongoDBExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeMongoDBExporterOK struct {
	Payload *ChangeMongoDBExporterOKBody
}

func (o *ChangeMongoDBExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeMongoDBExporter][%d] changeMongoDbExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeMongoDBExporterOK) GetPayload() *ChangeMongoDBExporterOKBody {
	return o.Payload
}

func (o *ChangeMongoDBExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeMongoDBExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeMongoDBExporterDefault creates a ChangeMongoDBExporterDefault with default headers values
func NewChangeMongoDBExporterDefault(code int) *ChangeMongoDBExporterDefault {
	return &ChangeMongoDBExporterDefault{
		_statusCode: code,
	}
}

/* ChangeMongoDBExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeMongoDBExporterDefault struct {
	_statusCode int

	Payload *ChangeMongoDBExporterDefaultBody
}

// Code gets the status code for the change mongo DB exporter default response
func (o *ChangeMongoDBExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeMongoDBExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeMongoDBExporter][%d] ChangeMongoDBExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeMongoDBExporterDefault) GetPayload() *ChangeMongoDBExporterDefaultBody {
	return o.Payload
}

func (o *ChangeMongoDBExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeMongoDBExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeMongoDBExporterBody change mongo DB exporter body
swagger:model ChangeMongoDBExporterBody
*/
type ChangeMongoDBExporterBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeMongoDBExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change mongo DB exporter body
func (o *ChangeMongoDBExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMongoDBExporterBody) validateCommon(formats strfmt.Registry) error {
	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change mongo DB exporter body based on the context it is used
func (o *ChangeMongoDBExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMongoDBExporterBody) contextValidateCommon(ctx context.Context, formats strfmt.Registry) error {
	if o.Common != nil {
		if err := o.Common.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterDefaultBody change mongo DB exporter default body
swagger:model ChangeMongoDBExporterDefaultBody
*/
type ChangeMongoDBExporterDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeMongoDBExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change mongo DB exporter default body
func (o *ChangeMongoDBExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMongoDBExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change mongo DB exporter default body based on the context it is used
func (o *ChangeMongoDBExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMongoDBExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterDefaultBodyDetailsItems0 change mongo DB exporter default body details items0
swagger:model ChangeMongoDBExporterDefaultBodyDetailsItems0
*/
type ChangeMongoDBExporterDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this change mongo DB exporter default body details items0
func (o *ChangeMongoDBExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change mongo DB exporter default body details items0 based on context it is used
func (o *ChangeMongoDBExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterOKBody change mongo DB exporter OK body
swagger:model ChangeMongoDBExporterOKBody
*/
type ChangeMongoDBExporterOKBody struct {
	// mongodb exporter
	MongodbExporter *ChangeMongoDBExporterOKBodyMongodbExporter `json:"mongodb_exporter,omitempty"`
}

// Validate validates this change mongo DB exporter OK body
func (o *ChangeMongoDBExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMongoDBExporterOKBody) validateMongodbExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeMongoDbExporterOk" + "." + "mongodb_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeMongoDbExporterOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change mongo DB exporter OK body based on the context it is used
func (o *ChangeMongoDBExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMongodbExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMongoDBExporterOKBody) contextValidateMongodbExporter(ctx context.Context, formats strfmt.Registry) error {
	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeMongoDbExporterOk" + "." + "mongodb_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeMongoDbExporterOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterOKBodyMongodbExporter MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model ChangeMongoDBExporterOKBodyMongodbExporter
*/
type ChangeMongoDBExporterOKBodyMongodbExporter struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
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

	// List of colletions to get stats from. Can use *
	StatsCollections []string `json:"stats_collections"`

	// Collections limit. Only get Databases and collection stats if the total number of collections in the server
	// is less than this value. 0: no limit
	CollectionsLimit int32 `json:"collections_limit,omitempty"`

	// Enable All collectors.
	EnableAllCollectors bool `json:"enable_all_collectors,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// Log level for exporters
	// Enum: [auto fatal error warn info debug]
	LogLevel *string `json:"log_level,omitempty"`
}

// Validate validates this change mongo DB exporter OK body mongodb exporter
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum = append(changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusSTARTING captures enum value "STARTING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusSTARTING string = "STARTING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusRUNNING captures enum value "RUNNING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusRUNNING string = "RUNNING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusWAITING captures enum value "WAITING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusWAITING string = "WAITING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusSTOPPING string = "STOPPING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusDONE captures enum value "DONE"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusDONE string = "DONE"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusUNKNOWN captures enum value "UNKNOWN"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeMongoDBExporterOKBodyMongodbExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeMongoDbExporterOk"+"."+"mongodb_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var changeMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","fatal","error","warn","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum = append(changeMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum, v)
	}
}

const (

	// ChangeMongoDBExporterOKBodyMongodbExporterLogLevelAuto captures enum value "auto"
	ChangeMongoDBExporterOKBodyMongodbExporterLogLevelAuto string = "auto"

	// ChangeMongoDBExporterOKBodyMongodbExporterLogLevelFatal captures enum value "fatal"
	ChangeMongoDBExporterOKBodyMongodbExporterLogLevelFatal string = "fatal"

	// ChangeMongoDBExporterOKBodyMongodbExporterLogLevelError captures enum value "error"
	ChangeMongoDBExporterOKBodyMongodbExporterLogLevelError string = "error"

	// ChangeMongoDBExporterOKBodyMongodbExporterLogLevelWarn captures enum value "warn"
	ChangeMongoDBExporterOKBodyMongodbExporterLogLevelWarn string = "warn"

	// ChangeMongoDBExporterOKBodyMongodbExporterLogLevelInfo captures enum value "info"
	ChangeMongoDBExporterOKBodyMongodbExporterLogLevelInfo string = "info"

	// ChangeMongoDBExporterOKBodyMongodbExporterLogLevelDebug captures enum value "debug"
	ChangeMongoDBExporterOKBodyMongodbExporterLogLevelDebug string = "debug"
)

// prop value enum
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeMongoDbExporterOkBodyMongodbExporterTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeMongoDBExporterOKBodyMongodbExporter) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("changeMongoDbExporterOk"+"."+"mongodb_exporter"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change mongo DB exporter OK body mongodb exporter based on context it is used
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterOKBodyMongodbExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeMongoDBExporterParamsBodyCommon
*/
type ChangeMongoDBExporterParamsBodyCommon struct {
	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change mongo DB exporter params body common
func (o *ChangeMongoDBExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change mongo DB exporter params body common based on context it is used
func (o *ChangeMongoDBExporterParamsBodyCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
