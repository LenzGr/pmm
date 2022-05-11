// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/grafana/gmodels"
)

// GetDatasourceByNameReader is a Reader for the GetDatasourceByName structure.
type GetDatasourceByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatasourceByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatasourceByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDatasourceByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDatasourceByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDatasourceByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatasourceByNameOK creates a GetDatasourceByNameOK with default headers values
func NewGetDatasourceByNameOK() *GetDatasourceByNameOK {
	return &GetDatasourceByNameOK{}
}

/* GetDatasourceByNameOK describes a response with status code 200, with default header values.

GetDatasourceByNameOK get datasource by name o k
*/
type GetDatasourceByNameOK struct {
	Payload *gmodels.DataSource
}

func (o *GetDatasourceByNameOK) Error() string {
	return fmt.Sprintf("[GET /api/datasources/name/{datasource_name}][%d] getDatasourceByNameOK  %+v", 200, o.Payload)
}
func (o *GetDatasourceByNameOK) GetPayload() *gmodels.DataSource {
	return o.Payload
}

func (o *GetDatasourceByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gmodels.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasourceByNameUnauthorized creates a GetDatasourceByNameUnauthorized with default headers values
func NewGetDatasourceByNameUnauthorized() *GetDatasourceByNameUnauthorized {
	return &GetDatasourceByNameUnauthorized{}
}

/* GetDatasourceByNameUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetDatasourceByNameUnauthorized struct {
	Payload *gmodels.ErrorResponseBody
}

func (o *GetDatasourceByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/datasources/name/{datasource_name}][%d] getDatasourceByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDatasourceByNameUnauthorized) GetPayload() *gmodels.ErrorResponseBody {
	return o.Payload
}

func (o *GetDatasourceByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gmodels.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasourceByNameForbidden creates a GetDatasourceByNameForbidden with default headers values
func NewGetDatasourceByNameForbidden() *GetDatasourceByNameForbidden {
	return &GetDatasourceByNameForbidden{}
}

/* GetDatasourceByNameForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetDatasourceByNameForbidden struct {
	Payload *gmodels.ErrorResponseBody
}

func (o *GetDatasourceByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /api/datasources/name/{datasource_name}][%d] getDatasourceByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetDatasourceByNameForbidden) GetPayload() *gmodels.ErrorResponseBody {
	return o.Payload
}

func (o *GetDatasourceByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gmodels.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasourceByNameInternalServerError creates a GetDatasourceByNameInternalServerError with default headers values
func NewGetDatasourceByNameInternalServerError() *GetDatasourceByNameInternalServerError {
	return &GetDatasourceByNameInternalServerError{}
}

/* GetDatasourceByNameInternalServerError describes a response with status code 500, with default header values.

GetDatasourceByNameInternalServerError get datasource by name internal server error
*/
type GetDatasourceByNameInternalServerError struct {
}

func (o *GetDatasourceByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/datasources/name/{datasource_name}][%d] getDatasourceByNameInternalServerError ", 500)
}

func (o *GetDatasourceByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
