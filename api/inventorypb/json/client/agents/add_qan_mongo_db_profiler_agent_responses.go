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

// AddQANMongoDBProfilerAgentReader is a Reader for the AddQANMongoDBProfilerAgent structure.
type AddQANMongoDBProfilerAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddQANMongoDBProfilerAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddQANMongoDBProfilerAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddQANMongoDBProfilerAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddQANMongoDBProfilerAgentOK creates a AddQANMongoDBProfilerAgentOK with default headers values
func NewAddQANMongoDBProfilerAgentOK() *AddQANMongoDBProfilerAgentOK {
	return &AddQANMongoDBProfilerAgentOK{}
}

/* AddQANMongoDBProfilerAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddQANMongoDBProfilerAgentOK struct {
	Payload *AddQANMongoDBProfilerAgentOKBody
}

func (o *AddQANMongoDBProfilerAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANMongoDBProfilerAgent][%d] addQanMongoDbProfilerAgentOk  %+v", 200, o.Payload)
}
func (o *AddQANMongoDBProfilerAgentOK) GetPayload() *AddQANMongoDBProfilerAgentOKBody {
	return o.Payload
}

func (o *AddQANMongoDBProfilerAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANMongoDBProfilerAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddQANMongoDBProfilerAgentDefault creates a AddQANMongoDBProfilerAgentDefault with default headers values
func NewAddQANMongoDBProfilerAgentDefault(code int) *AddQANMongoDBProfilerAgentDefault {
	return &AddQANMongoDBProfilerAgentDefault{
		_statusCode: code,
	}
}

/* AddQANMongoDBProfilerAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddQANMongoDBProfilerAgentDefault struct {
	_statusCode int

	Payload *AddQANMongoDBProfilerAgentDefaultBody
}

// Code gets the status code for the add QAN mongo DB profiler agent default response
func (o *AddQANMongoDBProfilerAgentDefault) Code() int {
	return o._statusCode
}

func (o *AddQANMongoDBProfilerAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANMongoDBProfilerAgent][%d] AddQANMongoDBProfilerAgent default  %+v", o._statusCode, o.Payload)
}
func (o *AddQANMongoDBProfilerAgentDefault) GetPayload() *AddQANMongoDBProfilerAgentDefaultBody {
	return o.Payload
}

func (o *AddQANMongoDBProfilerAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANMongoDBProfilerAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddQANMongoDBProfilerAgentBody add QAN mongo DB profiler agent body
swagger:model AddQANMongoDBProfilerAgentBody
*/
type AddQANMongoDBProfilerAgentBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for getting profile data.
	Username string `json:"username,omitempty"`

	// MongoDB password for getting profile data.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Client certificate and key.
	TLSCertificateKey string `json:"tls_certificate_key,omitempty"`

	// Password for decrypting tls_certificate_key.
	TLSCertificateKeyFilePassword string `json:"tls_certificate_key_file_password,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Authentication mechanism.
	// See https://docs.mongodb.com/manual/reference/connection-string/#mongodb-urioption-urioption.authMechanism
	// for details.
	AuthenticationMechanism string `json:"authentication_mechanism,omitempty"`

	// Authentication database.
	AuthenticationDatabase string `json:"authentication_database,omitempty"`
}

// Validate validates this add QAN mongo DB profiler agent body
func (o *AddQANMongoDBProfilerAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add QAN mongo DB profiler agent body based on context it is used
func (o *AddQANMongoDBProfilerAgentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentBody) UnmarshalBinary(b []byte) error {
	var res AddQANMongoDBProfilerAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMongoDBProfilerAgentDefaultBody add QAN mongo DB profiler agent default body
swagger:model AddQANMongoDBProfilerAgentDefaultBody
*/
type AddQANMongoDBProfilerAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add QAN mongo DB profiler agent default body
func (o *AddQANMongoDBProfilerAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANMongoDBProfilerAgentDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddQANMongoDBProfilerAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddQANMongoDBProfilerAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add QAN mongo DB profiler agent default body based on the context it is used
func (o *AddQANMongoDBProfilerAgentDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANMongoDBProfilerAgentDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddQANMongoDBProfilerAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddQANMongoDBProfilerAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddQANMongoDBProfilerAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0 add QAN mongo DB profiler agent default body details items0
swagger:model AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0
*/
type AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0 struct {

	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add QAN mongo DB profiler agent default body details items0
func (o *AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add QAN mongo DB profiler agent default body details items0 based on context it is used
func (o *AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddQANMongoDBProfilerAgentDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMongoDBProfilerAgentOKBody add QAN mongo DB profiler agent OK body
swagger:model AddQANMongoDBProfilerAgentOKBody
*/
type AddQANMongoDBProfilerAgentOKBody struct {

	// qan mongodb profiler agent
	QANMongodbProfilerAgent *AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent `json:"qan_mongodb_profiler_agent,omitempty"`
}

// Validate validates this add QAN mongo DB profiler agent OK body
func (o *AddQANMongoDBProfilerAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANMongodbProfilerAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANMongoDBProfilerAgentOKBody) validateQANMongodbProfilerAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.QANMongodbProfilerAgent) { // not required
		return nil
	}

	if o.QANMongodbProfilerAgent != nil {
		if err := o.QANMongodbProfilerAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addQanMongoDbProfilerAgentOk" + "." + "qan_mongodb_profiler_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addQanMongoDbProfilerAgentOk" + "." + "qan_mongodb_profiler_agent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add QAN mongo DB profiler agent OK body based on the context it is used
func (o *AddQANMongoDBProfilerAgentOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateQANMongodbProfilerAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANMongoDBProfilerAgentOKBody) contextValidateQANMongodbProfilerAgent(ctx context.Context, formats strfmt.Registry) error {

	if o.QANMongodbProfilerAgent != nil {
		if err := o.QANMongodbProfilerAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addQanMongoDbProfilerAgentOk" + "." + "qan_mongodb_profiler_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addQanMongoDbProfilerAgentOk" + "." + "qan_mongodb_profiler_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentOKBody) UnmarshalBinary(b []byte) error {
	var res AddQANMongoDBProfilerAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent QANMongoDBProfilerAgent runs within pmm-agent and sends MongoDB Query Analytics data to the PMM Server.
swagger:model AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent
*/
type AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for getting profiler data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

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
}

// Validate validates this add QAN mongo DB profiler agent OK body QAN mongodb profiler agent
func (o *AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum = append(addQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTARTING captures enum value "STARTING"
	AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTARTING string = "STARTING"

	// AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusRUNNING captures enum value "RUNNING"
	AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusRUNNING string = "RUNNING"

	// AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusWAITING captures enum value "WAITING"
	AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusWAITING string = "WAITING"

	// AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTOPPING captures enum value "STOPPING"
	AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusSTOPPING string = "STOPPING"

	// AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusDONE captures enum value "DONE"
	AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusDONE string = "DONE"

	// AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusUNKNOWN captures enum value "UNKNOWN"
	AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgentStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addQanMongoDbProfilerAgentOkBodyQanMongodbProfilerAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addQanMongoDbProfilerAgentOk"+"."+"qan_mongodb_profiler_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add QAN mongo DB profiler agent OK body QAN mongodb profiler agent based on context it is used
func (o *AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent) UnmarshalBinary(b []byte) error {
	var res AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
