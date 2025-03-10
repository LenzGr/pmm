// Code generated by go-swagger; DO NOT EDIT.

package components

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

// GetPXCComponentsReader is a Reader for the GetPXCComponents structure.
type GetPXCComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPXCComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPXCComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPXCComponentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPXCComponentsOK creates a GetPXCComponentsOK with default headers values
func NewGetPXCComponentsOK() *GetPXCComponentsOK {
	return &GetPXCComponentsOK{}
}

/* GetPXCComponentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetPXCComponentsOK struct {
	Payload *GetPXCComponentsOKBody
}

func (o *GetPXCComponentsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPXC][%d] getPxcComponentsOk  %+v", 200, o.Payload)
}

func (o *GetPXCComponentsOK) GetPayload() *GetPXCComponentsOKBody {
	return o.Payload
}

func (o *GetPXCComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPXCComponentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPXCComponentsDefault creates a GetPXCComponentsDefault with default headers values
func NewGetPXCComponentsDefault(code int) *GetPXCComponentsDefault {
	return &GetPXCComponentsDefault{
		_statusCode: code,
	}
}

/* GetPXCComponentsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetPXCComponentsDefault struct {
	_statusCode int

	Payload *GetPXCComponentsDefaultBody
}

// Code gets the status code for the get PXC components default response
func (o *GetPXCComponentsDefault) Code() int {
	return o._statusCode
}

func (o *GetPXCComponentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPXC][%d] GetPXCComponents default  %+v", o._statusCode, o.Payload)
}

func (o *GetPXCComponentsDefault) GetPayload() *GetPXCComponentsDefaultBody {
	return o.Payload
}

func (o *GetPXCComponentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPXCComponentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPXCComponentsBody get PXC components body
swagger:model GetPXCComponentsBody
*/
type GetPXCComponentsBody struct {
	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// Version of DB.
	DBVersion string `json:"db_version,omitempty"`
}

// Validate validates this get PXC components body
func (o *GetPXCComponentsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components body based on context it is used
func (o *GetPXCComponentsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsBody) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsDefaultBody get PXC components default body
swagger:model GetPXCComponentsDefaultBody
*/
type GetPXCComponentsDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetPXCComponentsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get PXC components default body
func (o *GetPXCComponentsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetPXCComponents default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPXCComponents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PXC components default body based on the context it is used
func (o *GetPXCComponentsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetPXCComponents default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPXCComponents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsDefaultBodyDetailsItems0 get PXC components default body details items0
swagger:model GetPXCComponentsDefaultBodyDetailsItems0
*/
type GetPXCComponentsDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this get PXC components default body details items0
func (o *GetPXCComponentsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components default body details items0 based on context it is used
func (o *GetPXCComponentsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBody get PXC components OK body
swagger:model GetPXCComponentsOKBody
*/
type GetPXCComponentsOKBody struct {
	// versions
	Versions []*GetPXCComponentsOKBodyVersionsItems0 `json:"versions"`
}

// Validate validates this get PXC components OK body
func (o *GetPXCComponentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsOKBody) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPxcComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPxcComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PXC components OK body based on the context it is used
func (o *GetPXCComponentsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsOKBody) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Versions); i++ {
		if o.Versions[i] != nil {
			if err := o.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPxcComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPxcComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0 OperatorVersion contains information about operator and components matrix.
swagger:model GetPXCComponentsOKBodyVersionsItems0
*/
type GetPXCComponentsOKBodyVersionsItems0 struct {
	// product
	Product string `json:"product,omitempty"`

	// operator
	Operator string `json:"operator,omitempty"`

	// matrix
	Matrix *GetPXCComponentsOKBodyVersionsItems0Matrix `json:"matrix,omitempty"`
}

// Validate validates this get PXC components OK body versions items0
func (o *GetPXCComponentsOKBodyVersionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMatrix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0) validateMatrix(formats strfmt.Registry) error {
	if swag.IsZero(o.Matrix) { // not required
		return nil
	}

	if o.Matrix != nil {
		if err := o.Matrix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC components OK body versions items0 based on the context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMatrix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0) contextValidateMatrix(ctx context.Context, formats strfmt.Registry) error {
	if o.Matrix != nil {
		if err := o.Matrix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0Matrix Matrix contains all available components.
swagger:model GetPXCComponentsOKBodyVersionsItems0Matrix
*/
type GetPXCComponentsOKBodyVersionsItems0Matrix struct {
	// mongod
	Mongod map[string]GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon `json:"mongod,omitempty"`

	// pxc
	PXC map[string]GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon `json:"pxc,omitempty"`

	// pmm
	PMM map[string]GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon `json:"pmm,omitempty"`

	// proxysql
	Proxysql map[string]GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon `json:"proxysql,omitempty"`

	// haproxy
	Haproxy map[string]GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon `json:"haproxy,omitempty"`

	// backup
	Backup map[string]GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon `json:"backup,omitempty"`

	// operator
	Operator map[string]GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon `json:"operator,omitempty"`

	// log collector
	LogCollector map[string]GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon `json:"log_collector,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix
func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePXC(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMM(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogCollector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validateMongod(formats strfmt.Registry) error {
	if swag.IsZero(o.Mongod) { // not required
		return nil
	}

	for k := range o.Mongod {

		if swag.IsZero(o.Mongod[k]) { // not required
			continue
		}
		if val, ok := o.Mongod[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "mongod" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "mongod" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validatePXC(formats strfmt.Registry) error {
	if swag.IsZero(o.PXC) { // not required
		return nil
	}

	for k := range o.PXC {

		if swag.IsZero(o.PXC[k]) { // not required
			continue
		}
		if val, ok := o.PXC[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "pxc" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "pxc" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validatePMM(formats strfmt.Registry) error {
	if swag.IsZero(o.PMM) { // not required
		return nil
	}

	for k := range o.PMM {

		if swag.IsZero(o.PMM[k]) { // not required
			continue
		}
		if val, ok := o.PMM[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "pmm" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "pmm" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validateProxysql(formats strfmt.Registry) error {
	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	for k := range o.Proxysql {

		if swag.IsZero(o.Proxysql[k]) { // not required
			continue
		}
		if val, ok := o.Proxysql[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "proxysql" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "proxysql" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	for k := range o.Haproxy {

		if swag.IsZero(o.Haproxy[k]) { // not required
			continue
		}
		if val, ok := o.Haproxy[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "haproxy" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "haproxy" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(o.Backup) { // not required
		return nil
	}

	for k := range o.Backup {

		if swag.IsZero(o.Backup[k]) { // not required
			continue
		}
		if val, ok := o.Backup[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "backup" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "backup" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(o.Operator) { // not required
		return nil
	}

	for k := range o.Operator {

		if swag.IsZero(o.Operator[k]) { // not required
			continue
		}
		if val, ok := o.Operator[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "operator" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "operator" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) validateLogCollector(formats strfmt.Registry) error {
	if swag.IsZero(o.LogCollector) { // not required
		return nil
	}

	for k := range o.LogCollector {

		if swag.IsZero(o.LogCollector[k]) { // not required
			continue
		}
		if val, ok := o.LogCollector[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "log_collector" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "log_collector" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PXC components OK body versions items0 matrix based on the context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMongod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePXC(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePMM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateProxysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateHaproxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLogCollector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidateMongod(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Mongod {
		if val, ok := o.Mongod[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidatePXC(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.PXC {
		if val, ok := o.PXC[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidatePMM(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.PMM {
		if val, ok := o.PMM[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidateProxysql(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Proxysql {
		if val, ok := o.Proxysql[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidateHaproxy(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Haproxy {
		if val, ok := o.Haproxy[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Backup {
		if val, ok := o.Backup[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidateOperator(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Operator {
		if val, ok := o.Operator[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) contextValidateLogCollector(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.LogCollector {
		if val, ok := o.LogCollector[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0Matrix) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0Matrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix backup anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix backup anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixBackupAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix haproxy anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix haproxy anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixHaproxyAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix log collector anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix log collector anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixLogCollectorAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix mongod anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix mongod anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixMongodAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix operator anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix operator anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixOperatorAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix PMM anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix PMM anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixPMMAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix PXC anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix PXC anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixPXCAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon Component contains information about component.
swagger:model GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon
*/
type GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PXC components OK body versions items0 matrix proxysql anon
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC components OK body versions items0 matrix proxysql anon based on context it is used
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon) UnmarshalBinary(b []byte) error {
	var res GetPXCComponentsOKBodyVersionsItems0MatrixProxysqlAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
