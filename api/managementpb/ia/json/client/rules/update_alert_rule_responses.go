// Code generated by go-swagger; DO NOT EDIT.

package rules

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

// UpdateAlertRuleReader is a Reader for the UpdateAlertRule structure.
type UpdateAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAlertRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAlertRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAlertRuleOK creates a UpdateAlertRuleOK with default headers values
func NewUpdateAlertRuleOK() *UpdateAlertRuleOK {
	return &UpdateAlertRuleOK{}
}

/* UpdateAlertRuleOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateAlertRuleOK struct {
	Payload interface{}
}

func (o *UpdateAlertRuleOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Update][%d] updateAlertRuleOk  %+v", 200, o.Payload)
}

func (o *UpdateAlertRuleOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateAlertRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertRuleDefault creates a UpdateAlertRuleDefault with default headers values
func NewUpdateAlertRuleDefault(code int) *UpdateAlertRuleDefault {
	return &UpdateAlertRuleDefault{
		_statusCode: code,
	}
}

/* UpdateAlertRuleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateAlertRuleDefault struct {
	_statusCode int

	Payload *UpdateAlertRuleDefaultBody
}

// Code gets the status code for the update alert rule default response
func (o *UpdateAlertRuleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAlertRuleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Update][%d] UpdateAlertRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAlertRuleDefault) GetPayload() *UpdateAlertRuleDefaultBody {
	return o.Payload
}

func (o *UpdateAlertRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UpdateAlertRuleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateAlertRuleBody update alert rule body
swagger:model UpdateAlertRuleBody
*/
type UpdateAlertRuleBody struct {
	// Rule ID.
	RuleID string `json:"rule_id,omitempty"`

	// Rule name. Should be set.
	Name string `json:"name,omitempty"`

	// New rule status. Should be set.
	Disabled bool `json:"disabled,omitempty"`

	// Rule parameters. All template parameters should be set.
	Params []*UpdateAlertRuleParamsBodyParamsItems0 `json:"params"`

	// Rule duration. Should be set.
	For string `json:"for,omitempty"`

	// Severity represents severity level of the check result or alert.
	// Enum: [SEVERITY_INVALID SEVERITY_EMERGENCY SEVERITY_ALERT SEVERITY_CRITICAL SEVERITY_ERROR SEVERITY_WARNING SEVERITY_NOTICE SEVERITY_INFO SEVERITY_DEBUG]
	Severity *string `json:"severity,omitempty"`

	// All custom labels to add or remove (with empty values) to default labels from template.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Filters. Should be set.
	Filters []*UpdateAlertRuleParamsBodyFiltersItems0 `json:"filters"`

	// Channels. Should be set.
	ChannelIds []string `json:"channel_ids"`
}

// Validate validates this update alert rule body
func (o *UpdateAlertRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAlertRuleBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var updateAlertRuleBodyTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEVERITY_INVALID","SEVERITY_EMERGENCY","SEVERITY_ALERT","SEVERITY_CRITICAL","SEVERITY_ERROR","SEVERITY_WARNING","SEVERITY_NOTICE","SEVERITY_INFO","SEVERITY_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateAlertRuleBodyTypeSeverityPropEnum = append(updateAlertRuleBodyTypeSeverityPropEnum, v)
	}
}

const (

	// UpdateAlertRuleBodySeveritySEVERITYINVALID captures enum value "SEVERITY_INVALID"
	UpdateAlertRuleBodySeveritySEVERITYINVALID string = "SEVERITY_INVALID"

	// UpdateAlertRuleBodySeveritySEVERITYEMERGENCY captures enum value "SEVERITY_EMERGENCY"
	UpdateAlertRuleBodySeveritySEVERITYEMERGENCY string = "SEVERITY_EMERGENCY"

	// UpdateAlertRuleBodySeveritySEVERITYALERT captures enum value "SEVERITY_ALERT"
	UpdateAlertRuleBodySeveritySEVERITYALERT string = "SEVERITY_ALERT"

	// UpdateAlertRuleBodySeveritySEVERITYCRITICAL captures enum value "SEVERITY_CRITICAL"
	UpdateAlertRuleBodySeveritySEVERITYCRITICAL string = "SEVERITY_CRITICAL"

	// UpdateAlertRuleBodySeveritySEVERITYERROR captures enum value "SEVERITY_ERROR"
	UpdateAlertRuleBodySeveritySEVERITYERROR string = "SEVERITY_ERROR"

	// UpdateAlertRuleBodySeveritySEVERITYWARNING captures enum value "SEVERITY_WARNING"
	UpdateAlertRuleBodySeveritySEVERITYWARNING string = "SEVERITY_WARNING"

	// UpdateAlertRuleBodySeveritySEVERITYNOTICE captures enum value "SEVERITY_NOTICE"
	UpdateAlertRuleBodySeveritySEVERITYNOTICE string = "SEVERITY_NOTICE"

	// UpdateAlertRuleBodySeveritySEVERITYINFO captures enum value "SEVERITY_INFO"
	UpdateAlertRuleBodySeveritySEVERITYINFO string = "SEVERITY_INFO"

	// UpdateAlertRuleBodySeveritySEVERITYDEBUG captures enum value "SEVERITY_DEBUG"
	UpdateAlertRuleBodySeveritySEVERITYDEBUG string = "SEVERITY_DEBUG"
)

// prop value enum
func (o *UpdateAlertRuleBody) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateAlertRuleBodyTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAlertRuleBody) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(o.Severity) { // not required
		return nil
	}

	// value enum
	if err := o.validateSeverityEnum("body"+"."+"severity", "body", *o.Severity); err != nil {
		return err
	}

	return nil
}

func (o *UpdateAlertRuleBody) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(o.Filters) { // not required
		return nil
	}

	for i := 0; i < len(o.Filters); i++ {
		if swag.IsZero(o.Filters[i]) { // not required
			continue
		}

		if o.Filters[i] != nil {
			if err := o.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update alert rule body based on the context it is used
func (o *UpdateAlertRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAlertRuleBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Params); i++ {
		if o.Params[i] != nil {
			if err := o.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *UpdateAlertRuleBody) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Filters); i++ {
		if o.Filters[i] != nil {
			if err := o.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateAlertRuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAlertRuleBody) UnmarshalBinary(b []byte) error {
	var res UpdateAlertRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateAlertRuleDefaultBody update alert rule default body
swagger:model UpdateAlertRuleDefaultBody
*/
type UpdateAlertRuleDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UpdateAlertRuleDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this update alert rule default body
func (o *UpdateAlertRuleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAlertRuleDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UpdateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update alert rule default body based on the context it is used
func (o *UpdateAlertRuleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAlertRuleDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateAlertRule default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateAlertRuleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAlertRuleDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateAlertRuleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateAlertRuleDefaultBodyDetailsItems0 update alert rule default body details items0
swagger:model UpdateAlertRuleDefaultBodyDetailsItems0
*/
type UpdateAlertRuleDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this update alert rule default body details items0
func (o *UpdateAlertRuleDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update alert rule default body details items0 based on context it is used
func (o *UpdateAlertRuleDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateAlertRuleDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAlertRuleDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateAlertRuleDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateAlertRuleParamsBodyFiltersItems0 Filter repsents a single filter condition.
swagger:model UpdateAlertRuleParamsBodyFiltersItems0
*/
type UpdateAlertRuleParamsBodyFiltersItems0 struct {
	// FilterType represents filter matching type.
	//
	//  - EQUAL: =
	//  - REGEX: =~
	// Enum: [FILTER_TYPE_INVALID EQUAL REGEX]
	Type *string `json:"type,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this update alert rule params body filters items0
func (o *UpdateAlertRuleParamsBodyFiltersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateAlertRuleParamsBodyFiltersItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FILTER_TYPE_INVALID","EQUAL","REGEX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateAlertRuleParamsBodyFiltersItems0TypeTypePropEnum = append(updateAlertRuleParamsBodyFiltersItems0TypeTypePropEnum, v)
	}
}

const (

	// UpdateAlertRuleParamsBodyFiltersItems0TypeFILTERTYPEINVALID captures enum value "FILTER_TYPE_INVALID"
	UpdateAlertRuleParamsBodyFiltersItems0TypeFILTERTYPEINVALID string = "FILTER_TYPE_INVALID"

	// UpdateAlertRuleParamsBodyFiltersItems0TypeEQUAL captures enum value "EQUAL"
	UpdateAlertRuleParamsBodyFiltersItems0TypeEQUAL string = "EQUAL"

	// UpdateAlertRuleParamsBodyFiltersItems0TypeREGEX captures enum value "REGEX"
	UpdateAlertRuleParamsBodyFiltersItems0TypeREGEX string = "REGEX"
)

// prop value enum
func (o *UpdateAlertRuleParamsBodyFiltersItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateAlertRuleParamsBodyFiltersItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAlertRuleParamsBodyFiltersItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update alert rule params body filters items0 based on context it is used
func (o *UpdateAlertRuleParamsBodyFiltersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateAlertRuleParamsBodyFiltersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAlertRuleParamsBodyFiltersItems0) UnmarshalBinary(b []byte) error {
	var res UpdateAlertRuleParamsBodyFiltersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateAlertRuleParamsBodyParamsItems0 ParamValue represents a single rule parameter value for List, Change and Update APIs.
swagger:model UpdateAlertRuleParamsBodyParamsItems0
*/
type UpdateAlertRuleParamsBodyParamsItems0 struct {
	// Machine-readable name (ID) that is used in expression.
	Name string `json:"name,omitempty"`

	// ParamType represents template parameter type.
	// Enum: [PARAM_TYPE_INVALID BOOL FLOAT STRING]
	Type *string `json:"type,omitempty"`

	// Bool value.
	Bool bool `json:"bool,omitempty"`

	// Float value.
	Float float64 `json:"float,omitempty"`

	// String value.
	String string `json:"string,omitempty"`
}

// Validate validates this update alert rule params body params items0
func (o *UpdateAlertRuleParamsBodyParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateAlertRuleParamsBodyParamsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARAM_TYPE_INVALID","BOOL","FLOAT","STRING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateAlertRuleParamsBodyParamsItems0TypeTypePropEnum = append(updateAlertRuleParamsBodyParamsItems0TypeTypePropEnum, v)
	}
}

const (

	// UpdateAlertRuleParamsBodyParamsItems0TypePARAMTYPEINVALID captures enum value "PARAM_TYPE_INVALID"
	UpdateAlertRuleParamsBodyParamsItems0TypePARAMTYPEINVALID string = "PARAM_TYPE_INVALID"

	// UpdateAlertRuleParamsBodyParamsItems0TypeBOOL captures enum value "BOOL"
	UpdateAlertRuleParamsBodyParamsItems0TypeBOOL string = "BOOL"

	// UpdateAlertRuleParamsBodyParamsItems0TypeFLOAT captures enum value "FLOAT"
	UpdateAlertRuleParamsBodyParamsItems0TypeFLOAT string = "FLOAT"

	// UpdateAlertRuleParamsBodyParamsItems0TypeSTRING captures enum value "STRING"
	UpdateAlertRuleParamsBodyParamsItems0TypeSTRING string = "STRING"
)

// prop value enum
func (o *UpdateAlertRuleParamsBodyParamsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateAlertRuleParamsBodyParamsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAlertRuleParamsBodyParamsItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update alert rule params body params items0 based on context it is used
func (o *UpdateAlertRuleParamsBodyParamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateAlertRuleParamsBodyParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAlertRuleParamsBodyParamsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateAlertRuleParamsBodyParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
