// Code generated by go-swagger; DO NOT EDIT.

package pxc_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPXCClusterCredentialsParams creates a new GetPXCClusterCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPXCClusterCredentialsParams() *GetPXCClusterCredentialsParams {
	return &GetPXCClusterCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPXCClusterCredentialsParamsWithTimeout creates a new GetPXCClusterCredentialsParams object
// with the ability to set a timeout on a request.
func NewGetPXCClusterCredentialsParamsWithTimeout(timeout time.Duration) *GetPXCClusterCredentialsParams {
	return &GetPXCClusterCredentialsParams{
		timeout: timeout,
	}
}

// NewGetPXCClusterCredentialsParamsWithContext creates a new GetPXCClusterCredentialsParams object
// with the ability to set a context for a request.
func NewGetPXCClusterCredentialsParamsWithContext(ctx context.Context) *GetPXCClusterCredentialsParams {
	return &GetPXCClusterCredentialsParams{
		Context: ctx,
	}
}

// NewGetPXCClusterCredentialsParamsWithHTTPClient creates a new GetPXCClusterCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPXCClusterCredentialsParamsWithHTTPClient(client *http.Client) *GetPXCClusterCredentialsParams {
	return &GetPXCClusterCredentialsParams{
		HTTPClient: client,
	}
}

/* GetPXCClusterCredentialsParams contains all the parameters to send to the API endpoint
   for the get PXC cluster credentials operation.

   Typically these are written to a http.Request.
*/
type GetPXCClusterCredentialsParams struct {
	// Body.
	Body GetPXCClusterCredentialsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get PXC cluster credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPXCClusterCredentialsParams) WithDefaults() *GetPXCClusterCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get PXC cluster credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPXCClusterCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) WithTimeout(timeout time.Duration) *GetPXCClusterCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) WithContext(ctx context.Context) *GetPXCClusterCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) WithHTTPClient(client *http.Client) *GetPXCClusterCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) WithBody(body GetPXCClusterCredentialsBody) *GetPXCClusterCredentialsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get PXC cluster credentials params
func (o *GetPXCClusterCredentialsParams) SetBody(body GetPXCClusterCredentialsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetPXCClusterCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
