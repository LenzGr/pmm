// Code generated by go-swagger; DO NOT EDIT.

package psmdb_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdatePSMDBClusterReader is a Reader for the UpdatePSMDBCluster structure.
type UpdatePSMDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePSMDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePSMDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdatePSMDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePSMDBClusterOK creates a UpdatePSMDBClusterOK with default headers values
func NewUpdatePSMDBClusterOK() *UpdatePSMDBClusterOK {
	return &UpdatePSMDBClusterOK{}
}

/*UpdatePSMDBClusterOK handles this case with default header values.

A successful response.
*/
type UpdatePSMDBClusterOK struct {
	Payload interface{}
}

func (o *UpdatePSMDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBCluster/Update][%d] updatePsmdbClusterOk  %+v", 200, o.Payload)
}

func (o *UpdatePSMDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdatePSMDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePSMDBClusterDefault creates a UpdatePSMDBClusterDefault with default headers values
func NewUpdatePSMDBClusterDefault(code int) *UpdatePSMDBClusterDefault {
	return &UpdatePSMDBClusterDefault{
		_statusCode: code,
	}
}

/*UpdatePSMDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type UpdatePSMDBClusterDefault struct {
	_statusCode int

	Payload *UpdatePSMDBClusterDefaultBody
}

// Code gets the status code for the update PSMDB cluster default response
func (o *UpdatePSMDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePSMDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBCluster/Update][%d] UpdatePSMDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *UpdatePSMDBClusterDefault) GetPayload() *UpdatePSMDBClusterDefaultBody {
	return o.Payload
}

func (o *UpdatePSMDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdatePSMDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdatePSMDBClusterBody update PSMDB cluster body
swagger:model UpdatePSMDBClusterBody
*/
type UpdatePSMDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// PSMDB cluster name.
	Name string `json:"name,omitempty"`

	// params
	Params *UpdatePSMDBClusterParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this update PSMDB cluster body
func (o *UpdatePSMDBClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePSMDBClusterBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePSMDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePSMDBClusterBody) UnmarshalBinary(b []byte) error {
	var res UpdatePSMDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePSMDBClusterDefaultBody update PSMDB cluster default body
swagger:model UpdatePSMDBClusterDefaultBody
*/
type UpdatePSMDBClusterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this update PSMDB cluster default body
func (o *UpdatePSMDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePSMDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("UpdatePSMDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePSMDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePSMDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdatePSMDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePSMDBClusterParamsBodyParams UpdatePSMDBClusterParams represents PSMDB cluster parameters that can be updated.
swagger:model UpdatePSMDBClusterParamsBodyParams
*/
type UpdatePSMDBClusterParamsBodyParams struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// Suspend cluster `pause: true`.
	Suspend bool `json:"suspend,omitempty"`

	// Resume cluster `pause: false`.
	Resume bool `json:"resume,omitempty"`

	// PSMDB image to use. If it's the same image but with different version tag, upgrade of database cluster to version
	// in given tag is triggered. If entirely different image is given, error is returned.
	Image string `json:"image,omitempty"`

	// replicaset
	Replicaset *UpdatePSMDBClusterParamsBodyParamsReplicaset `json:"replicaset,omitempty"`
}

// Validate validates this update PSMDB cluster params body params
func (o *UpdatePSMDBClusterParamsBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReplicaset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePSMDBClusterParamsBodyParams) validateReplicaset(formats strfmt.Registry) error {

	if swag.IsZero(o.Replicaset) { // not required
		return nil
	}

	if o.Replicaset != nil {
		if err := o.Replicaset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePSMDBClusterParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePSMDBClusterParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res UpdatePSMDBClusterParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePSMDBClusterParamsBodyParamsReplicaset ReplicaSet container parameters.
swagger:model UpdatePSMDBClusterParamsBodyParamsReplicaset
*/
type UpdatePSMDBClusterParamsBodyParamsReplicaset struct {

	// compute resources
	ComputeResources *UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this update PSMDB cluster params body params replicaset
func (o *UpdatePSMDBClusterParamsBodyParamsReplicaset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePSMDBClusterParamsBodyParamsReplicaset) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "replicaset" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePSMDBClusterParamsBodyParamsReplicaset) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePSMDBClusterParamsBodyParamsReplicaset) UnmarshalBinary(b []byte) error {
	var res UpdatePSMDBClusterParamsBodyParamsReplicaset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources
*/
type UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this update PSMDB cluster params body params replicaset compute resources
func (o *UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources) UnmarshalBinary(b []byte) error {
	var res UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
