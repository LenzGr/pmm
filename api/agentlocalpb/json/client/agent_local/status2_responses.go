// Code generated by go-swagger; DO NOT EDIT.

package agent_local

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

// Status2Reader is a Reader for the Status2 structure.
type Status2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Status2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatus2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatus2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatus2OK creates a Status2OK with default headers values
func NewStatus2OK() *Status2OK {
	return &Status2OK{}
}

/* Status2OK describes a response with status code 200, with default header values.

A successful response.
*/
type Status2OK struct {
	Payload *Status2OKBody
}

func (o *Status2OK) Error() string {
	return fmt.Sprintf("[GET /local/Status][%d] status2Ok  %+v", 200, o.Payload)
}
func (o *Status2OK) GetPayload() *Status2OKBody {
	return o.Payload
}

func (o *Status2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(Status2OKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatus2Default creates a Status2Default with default headers values
func NewStatus2Default(code int) *Status2Default {
	return &Status2Default{
		_statusCode: code,
	}
}

/* Status2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type Status2Default struct {
	_statusCode int

	Payload *Status2DefaultBody
}

// Code gets the status code for the status2 default response
func (o *Status2Default) Code() int {
	return o._statusCode
}

func (o *Status2Default) Error() string {
	return fmt.Sprintf("[GET /local/Status][%d] Status2 default  %+v", o._statusCode, o.Payload)
}
func (o *Status2Default) GetPayload() *Status2DefaultBody {
	return o.Payload
}

func (o *Status2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(Status2DefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*Status2DefaultBody status2 default body
swagger:model Status2DefaultBody
*/
type Status2DefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*Status2DefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this status2 default body
func (o *Status2DefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *Status2DefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this status2 default body based on the context it is used
func (o *Status2DefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *Status2DefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status2 default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *Status2DefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2DefaultBody) UnmarshalBinary(b []byte) error {
	var res Status2DefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*Status2DefaultBodyDetailsItems0 status2 default body details items0
swagger:model Status2DefaultBodyDetailsItems0
*/
type Status2DefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this status2 default body details items0
func (o *Status2DefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status2 default body details items0 based on context it is used
func (o *Status2DefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *Status2DefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2DefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res Status2DefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*Status2OKBody status2 OK body
swagger:model Status2OKBody
*/
type Status2OKBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// runs on node id
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// agents info
	AgentsInfo []*Status2OKBodyAgentsInfoItems0 `json:"agents_info"`

	// Config file path if pmm-agent was started with one.
	ConfigFilepath string `json:"config_filepath,omitempty"`

	// PMM Agent version.
	AgentVersion string `json:"agent_version,omitempty"`

	// server info
	ServerInfo *Status2OKBodyServerInfo `json:"server_info,omitempty"`
}

// Validate validates this status2 OK body
func (o *Status2OKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgentsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServerInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *Status2OKBody) validateAgentsInfo(formats strfmt.Registry) error {
	if swag.IsZero(o.AgentsInfo) { // not required
		return nil
	}

	for i := 0; i < len(o.AgentsInfo); i++ {
		if swag.IsZero(o.AgentsInfo[i]) { // not required
			continue
		}

		if o.AgentsInfo[i] != nil {
			if err := o.AgentsInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *Status2OKBody) validateServerInfo(formats strfmt.Registry) error {
	if swag.IsZero(o.ServerInfo) { // not required
		return nil
	}

	if o.ServerInfo != nil {
		if err := o.ServerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status2Ok" + "." + "server_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status2Ok" + "." + "server_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status2 OK body based on the context it is used
func (o *Status2OKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAgentsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateServerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *Status2OKBody) contextValidateAgentsInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AgentsInfo); i++ {

		if o.AgentsInfo[i] != nil {
			if err := o.AgentsInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status2Ok" + "." + "agents_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *Status2OKBody) contextValidateServerInfo(ctx context.Context, formats strfmt.Registry) error {

	if o.ServerInfo != nil {
		if err := o.ServerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status2Ok" + "." + "server_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status2Ok" + "." + "server_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *Status2OKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2OKBody) UnmarshalBinary(b []byte) error {
	var res Status2OKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*Status2OKBodyAgentsInfoItems0 AgentInfo contains information about Agent managed by pmm-agent.
swagger:model Status2OKBodyAgentsInfoItems0
*/
type Status2OKBodyAgentsInfoItems0 struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// AgentType describes supported Agent types.
	// Enum: [AGENT_TYPE_INVALID PMM_AGENT VM_AGENT NODE_EXPORTER MYSQLD_EXPORTER MONGODB_EXPORTER POSTGRES_EXPORTER PROXYSQL_EXPORTER QAN_MYSQL_PERFSCHEMA_AGENT QAN_MYSQL_SLOWLOG_AGENT QAN_MONGODB_PROFILER_AGENT QAN_POSTGRESQL_PGSTATEMENTS_AGENT QAN_POSTGRESQL_PGSTATMONITOR_AGENT RDS_EXPORTER EXTERNAL_EXPORTER AZURE_DATABASE_EXPORTER]
	AgentType *string `json:"agent_type,omitempty"`

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

	// The current listen port of this Agent (exporter or vmagent).
	// Zero for other Agent types, or if unknown or not yet supported.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this status2 OK body agents info items0
func (o *Status2OKBodyAgentsInfoItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgentType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_TYPE_INVALID","PMM_AGENT","VM_AGENT","NODE_EXPORTER","MYSQLD_EXPORTER","MONGODB_EXPORTER","POSTGRES_EXPORTER","PROXYSQL_EXPORTER","QAN_MYSQL_PERFSCHEMA_AGENT","QAN_MYSQL_SLOWLOG_AGENT","QAN_MONGODB_PROFILER_AGENT","QAN_POSTGRESQL_PGSTATEMENTS_AGENT","QAN_POSTGRESQL_PGSTATMONITOR_AGENT","RDS_EXPORTER","EXTERNAL_EXPORTER","AZURE_DATABASE_EXPORTER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum = append(status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum, v)
	}
}

const (

	// Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEINVALID captures enum value "AGENT_TYPE_INVALID"
	Status2OKBodyAgentsInfoItems0AgentTypeAGENTTYPEINVALID string = "AGENT_TYPE_INVALID"

	// Status2OKBodyAgentsInfoItems0AgentTypePMMAGENT captures enum value "PMM_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypePMMAGENT string = "PMM_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeVMAGENT captures enum value "VM_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeVMAGENT string = "VM_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeNODEEXPORTER captures enum value "NODE_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeNODEEXPORTER string = "NODE_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeMYSQLDEXPORTER captures enum value "MYSQLD_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeMYSQLDEXPORTER string = "MYSQLD_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeMONGODBEXPORTER captures enum value "MONGODB_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeMONGODBEXPORTER string = "MONGODB_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypePOSTGRESEXPORTER captures enum value "POSTGRES_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypePOSTGRESEXPORTER string = "POSTGRES_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypePROXYSQLEXPORTER captures enum value "PROXYSQL_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypePROXYSQLEXPORTER string = "PROXYSQL_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeQANMYSQLPERFSCHEMAAGENT captures enum value "QAN_MYSQL_PERFSCHEMA_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeQANMYSQLPERFSCHEMAAGENT string = "QAN_MYSQL_PERFSCHEMA_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeQANMYSQLSLOWLOGAGENT captures enum value "QAN_MYSQL_SLOWLOG_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeQANMYSQLSLOWLOGAGENT string = "QAN_MYSQL_SLOWLOG_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeQANMONGODBPROFILERAGENT captures enum value "QAN_MONGODB_PROFILER_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeQANMONGODBPROFILERAGENT string = "QAN_MONGODB_PROFILER_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATEMENTSAGENT captures enum value "QAN_POSTGRESQL_PGSTATEMENTS_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATEMENTSAGENT string = "QAN_POSTGRESQL_PGSTATEMENTS_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATMONITORAGENT captures enum value "QAN_POSTGRESQL_PGSTATMONITOR_AGENT"
	Status2OKBodyAgentsInfoItems0AgentTypeQANPOSTGRESQLPGSTATMONITORAGENT string = "QAN_POSTGRESQL_PGSTATMONITOR_AGENT"

	// Status2OKBodyAgentsInfoItems0AgentTypeRDSEXPORTER captures enum value "RDS_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeRDSEXPORTER string = "RDS_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeEXTERNALEXPORTER captures enum value "EXTERNAL_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeEXTERNALEXPORTER string = "EXTERNAL_EXPORTER"

	// Status2OKBodyAgentsInfoItems0AgentTypeAZUREDATABASEEXPORTER captures enum value "AZURE_DATABASE_EXPORTER"
	Status2OKBodyAgentsInfoItems0AgentTypeAZUREDATABASEEXPORTER string = "AZURE_DATABASE_EXPORTER"
)

// prop value enum
func (o *Status2OKBodyAgentsInfoItems0) validateAgentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, status2OkBodyAgentsInfoItems0TypeAgentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *Status2OKBodyAgentsInfoItems0) validateAgentType(formats strfmt.Registry) error {
	if swag.IsZero(o.AgentType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAgentTypeEnum("agent_type", "body", *o.AgentType); err != nil {
		return err
	}

	return nil
}

var status2OkBodyAgentsInfoItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		status2OkBodyAgentsInfoItems0TypeStatusPropEnum = append(status2OkBodyAgentsInfoItems0TypeStatusPropEnum, v)
	}
}

const (

	// Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	Status2OKBodyAgentsInfoItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// Status2OKBodyAgentsInfoItems0StatusSTARTING captures enum value "STARTING"
	Status2OKBodyAgentsInfoItems0StatusSTARTING string = "STARTING"

	// Status2OKBodyAgentsInfoItems0StatusRUNNING captures enum value "RUNNING"
	Status2OKBodyAgentsInfoItems0StatusRUNNING string = "RUNNING"

	// Status2OKBodyAgentsInfoItems0StatusWAITING captures enum value "WAITING"
	Status2OKBodyAgentsInfoItems0StatusWAITING string = "WAITING"

	// Status2OKBodyAgentsInfoItems0StatusSTOPPING captures enum value "STOPPING"
	Status2OKBodyAgentsInfoItems0StatusSTOPPING string = "STOPPING"

	// Status2OKBodyAgentsInfoItems0StatusDONE captures enum value "DONE"
	Status2OKBodyAgentsInfoItems0StatusDONE string = "DONE"

	// Status2OKBodyAgentsInfoItems0StatusUNKNOWN captures enum value "UNKNOWN"
	Status2OKBodyAgentsInfoItems0StatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *Status2OKBodyAgentsInfoItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, status2OkBodyAgentsInfoItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *Status2OKBodyAgentsInfoItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this status2 OK body agents info items0 based on context it is used
func (o *Status2OKBodyAgentsInfoItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *Status2OKBodyAgentsInfoItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2OKBodyAgentsInfoItems0) UnmarshalBinary(b []byte) error {
	var res Status2OKBodyAgentsInfoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*Status2OKBodyServerInfo ServerInfo contains information about the PMM Server.
swagger:model Status2OKBodyServerInfo
*/
type Status2OKBodyServerInfo struct {

	// PMM Server URL in a form https://HOST:PORT/.
	URL string `json:"url,omitempty"`

	// PMM Server's TLS certificate validation should be skipped if true.
	InsecureTLS bool `json:"insecure_tls,omitempty"`

	// True if pmm-agent is currently connected to the server.
	Connected bool `json:"connected,omitempty"`

	// PMM Server version (if agent is connected).
	Version string `json:"version,omitempty"`

	// Ping time from pmm-agent to pmm-managed (if agent is connected).
	Latency string `json:"latency,omitempty"`

	// Clock drift from PMM Server (if agent is connected).
	ClockDrift string `json:"clock_drift,omitempty"`
}

// Validate validates this status2 OK body server info
func (o *Status2OKBodyServerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status2 OK body server info based on context it is used
func (o *Status2OKBodyServerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *Status2OKBodyServerInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Status2OKBodyServerInfo) UnmarshalBinary(b []byte) error {
	var res Status2OKBodyServerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
