// Code generated by go-swagger; DO NOT EDIT.

package security_checks

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

// GetFailedChecksReader is a Reader for the GetFailedChecks structure.
type GetFailedChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFailedChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFailedChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFailedChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFailedChecksOK creates a GetFailedChecksOK with default headers values
func NewGetFailedChecksOK() *GetFailedChecksOK {
	return &GetFailedChecksOK{}
}

/*GetFailedChecksOK handles this case with default header values.

A successful response.
*/
type GetFailedChecksOK struct {
	Payload *GetFailedChecksOKBody
}

func (o *GetFailedChecksOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/FailedChecks][%d] getFailedChecksOk  %+v", 200, o.Payload)
}

func (o *GetFailedChecksOK) GetPayload() *GetFailedChecksOKBody {
	return o.Payload
}

func (o *GetFailedChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFailedChecksOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFailedChecksDefault creates a GetFailedChecksDefault with default headers values
func NewGetFailedChecksDefault(code int) *GetFailedChecksDefault {
	return &GetFailedChecksDefault{
		_statusCode: code,
	}
}

/*GetFailedChecksDefault handles this case with default header values.

An unexpected error response.
*/
type GetFailedChecksDefault struct {
	_statusCode int

	Payload *GetFailedChecksDefaultBody
}

// Code gets the status code for the get failed checks default response
func (o *GetFailedChecksDefault) Code() int {
	return o._statusCode
}

func (o *GetFailedChecksDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/FailedChecks][%d] GetFailedChecks default  %+v", o._statusCode, o.Payload)
}

func (o *GetFailedChecksDefault) GetPayload() *GetFailedChecksDefaultBody {
	return o.Payload
}

func (o *GetFailedChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFailedChecksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFailedChecksBody get failed checks body
swagger:model GetFailedChecksBody
*/
type GetFailedChecksBody struct {

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// page params
	PageParams *GetFailedChecksParamsBodyPageParams `json:"page_params,omitempty"`
}

// Validate validates this get failed checks body
func (o *GetFailedChecksBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePageParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFailedChecksBody) validatePageParams(formats strfmt.Registry) error {

	if swag.IsZero(o.PageParams) { // not required
		return nil
	}

	if o.PageParams != nil {
		if err := o.PageParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "page_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFailedChecksBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFailedChecksBody) UnmarshalBinary(b []byte) error {
	var res GetFailedChecksBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFailedChecksDefaultBody get failed checks default body
swagger:model GetFailedChecksDefaultBody
*/
type GetFailedChecksDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get failed checks default body
func (o *GetFailedChecksDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFailedChecksDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetFailedChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFailedChecksDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFailedChecksDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetFailedChecksDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFailedChecksOKBody get failed checks OK body
swagger:model GetFailedChecksOKBody
*/
type GetFailedChecksOKBody struct {

	// results
	Results []*ResultsItems0 `json:"results"`

	// page totals
	PageTotals *GetFailedChecksOKBodyPageTotals `json:"page_totals,omitempty"`
}

// Validate validates this get failed checks OK body
func (o *GetFailedChecksOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePageTotals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFailedChecksOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getFailedChecksOk" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetFailedChecksOKBody) validatePageTotals(formats strfmt.Registry) error {

	if swag.IsZero(o.PageTotals) { // not required
		return nil
	}

	if o.PageTotals != nil {
		if err := o.PageTotals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getFailedChecksOk" + "." + "page_totals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFailedChecksOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFailedChecksOKBody) UnmarshalBinary(b []byte) error {
	var res GetFailedChecksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFailedChecksOKBodyPageTotals PageTotals represents total values for pagination.
swagger:model GetFailedChecksOKBodyPageTotals
*/
type GetFailedChecksOKBodyPageTotals struct {

	// Total number of results.
	TotalItems int32 `json:"total_items,omitempty"`

	// Total number of pages.
	TotalPages int32 `json:"total_pages,omitempty"`
}

// Validate validates this get failed checks OK body page totals
func (o *GetFailedChecksOKBodyPageTotals) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFailedChecksOKBodyPageTotals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFailedChecksOKBodyPageTotals) UnmarshalBinary(b []byte) error {
	var res GetFailedChecksOKBodyPageTotals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFailedChecksParamsBodyPageParams PageParams represents page request parameters for pagination.
swagger:model GetFailedChecksParamsBodyPageParams
*/
type GetFailedChecksParamsBodyPageParams struct {

	// Maximum number of results per page.
	PageSize int32 `json:"page_size,omitempty"`

	// Index of the requested page, starts from 0.
	Index int32 `json:"index,omitempty"`
}

// Validate validates this get failed checks params body page params
func (o *GetFailedChecksParamsBodyPageParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFailedChecksParamsBodyPageParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFailedChecksParamsBodyPageParams) UnmarshalBinary(b []byte) error {
	var res GetFailedChecksParamsBodyPageParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ResultsItems0 CheckResult represents the check results for a given service.
swagger:model ResultsItems0
*/
type ResultsItems0 struct {

	// summary
	Summary string `json:"summary,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Severity represents severity level of the check result or alert.
	// Enum: [SEVERITY_INVALID SEVERITY_EMERGENCY SEVERITY_ALERT SEVERITY_CRITICAL SEVERITY_ERROR SEVERITY_WARNING SEVERITY_NOTICE SEVERITY_INFO SEVERITY_DEBUG]
	Severity *string `json:"severity,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// URL containing information on how to resolve an issue detected by an STT check.
	ReadMoreURL string `json:"read_more_url,omitempty"`

	// Name of the monitored service on which the check ran.
	ServiceName string `json:"service_name,omitempty"`

	// ID of the monitored service on which the check ran.
	ServiceID string `json:"service_id,omitempty"`

	// Name of the check that failed
	CheckName string `json:"check_name,omitempty"`

	// ID of the check result as stored in AlertManager
	AlertID string `json:"alert_id,omitempty"`

	// Silence status of the check result
	Silenced bool `json:"silenced,omitempty"`
}

// Validate validates this results items0
func (o *ResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resultsItems0TypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEVERITY_INVALID","SEVERITY_EMERGENCY","SEVERITY_ALERT","SEVERITY_CRITICAL","SEVERITY_ERROR","SEVERITY_WARNING","SEVERITY_NOTICE","SEVERITY_INFO","SEVERITY_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resultsItems0TypeSeverityPropEnum = append(resultsItems0TypeSeverityPropEnum, v)
	}
}

const (

	// ResultsItems0SeveritySEVERITYINVALID captures enum value "SEVERITY_INVALID"
	ResultsItems0SeveritySEVERITYINVALID string = "SEVERITY_INVALID"

	// ResultsItems0SeveritySEVERITYEMERGENCY captures enum value "SEVERITY_EMERGENCY"
	ResultsItems0SeveritySEVERITYEMERGENCY string = "SEVERITY_EMERGENCY"

	// ResultsItems0SeveritySEVERITYALERT captures enum value "SEVERITY_ALERT"
	ResultsItems0SeveritySEVERITYALERT string = "SEVERITY_ALERT"

	// ResultsItems0SeveritySEVERITYCRITICAL captures enum value "SEVERITY_CRITICAL"
	ResultsItems0SeveritySEVERITYCRITICAL string = "SEVERITY_CRITICAL"

	// ResultsItems0SeveritySEVERITYERROR captures enum value "SEVERITY_ERROR"
	ResultsItems0SeveritySEVERITYERROR string = "SEVERITY_ERROR"

	// ResultsItems0SeveritySEVERITYWARNING captures enum value "SEVERITY_WARNING"
	ResultsItems0SeveritySEVERITYWARNING string = "SEVERITY_WARNING"

	// ResultsItems0SeveritySEVERITYNOTICE captures enum value "SEVERITY_NOTICE"
	ResultsItems0SeveritySEVERITYNOTICE string = "SEVERITY_NOTICE"

	// ResultsItems0SeveritySEVERITYINFO captures enum value "SEVERITY_INFO"
	ResultsItems0SeveritySEVERITYINFO string = "SEVERITY_INFO"

	// ResultsItems0SeveritySEVERITYDEBUG captures enum value "SEVERITY_DEBUG"
	ResultsItems0SeveritySEVERITYDEBUG string = "SEVERITY_DEBUG"
)

// prop value enum
func (o *ResultsItems0) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resultsItems0TypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ResultsItems0) validateSeverity(formats strfmt.Registry) error {

	if swag.IsZero(o.Severity) { // not required
		return nil
	}

	// value enum
	if err := o.validateSeverityEnum("severity", "body", *o.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResultsItems0) UnmarshalBinary(b []byte) error {
	var res ResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}