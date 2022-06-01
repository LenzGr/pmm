// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// NewTestLocationConfigParams creates a new TestLocationConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestLocationConfigParams() *TestLocationConfigParams {
	return &TestLocationConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestLocationConfigParamsWithTimeout creates a new TestLocationConfigParams object
// with the ability to set a timeout on a request.
func NewTestLocationConfigParamsWithTimeout(timeout time.Duration) *TestLocationConfigParams {
	return &TestLocationConfigParams{
		timeout: timeout,
	}
}

// NewTestLocationConfigParamsWithContext creates a new TestLocationConfigParams object
// with the ability to set a context for a request.
func NewTestLocationConfigParamsWithContext(ctx context.Context) *TestLocationConfigParams {
	return &TestLocationConfigParams{
		Context: ctx,
	}
}

// NewTestLocationConfigParamsWithHTTPClient creates a new TestLocationConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestLocationConfigParamsWithHTTPClient(client *http.Client) *TestLocationConfigParams {
	return &TestLocationConfigParams{
		HTTPClient: client,
	}
}

/* TestLocationConfigParams contains all the parameters to send to the API endpoint
   for the test location config operation.

   Typically these are written to a http.Request.
*/
type TestLocationConfigParams struct {
	// Body.
	Body TestLocationConfigBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test location config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestLocationConfigParams) WithDefaults() *TestLocationConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test location config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestLocationConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test location config params
func (o *TestLocationConfigParams) WithTimeout(timeout time.Duration) *TestLocationConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test location config params
func (o *TestLocationConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test location config params
func (o *TestLocationConfigParams) WithContext(ctx context.Context) *TestLocationConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test location config params
func (o *TestLocationConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test location config params
func (o *TestLocationConfigParams) WithHTTPClient(client *http.Client) *TestLocationConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test location config params
func (o *TestLocationConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test location config params
func (o *TestLocationConfigParams) WithBody(body TestLocationConfigBody) *TestLocationConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test location config params
func (o *TestLocationConfigParams) SetBody(body TestLocationConfigBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestLocationConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
