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

// NewRemoveLocationParams creates a new RemoveLocationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveLocationParams() *RemoveLocationParams {
	return &RemoveLocationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveLocationParamsWithTimeout creates a new RemoveLocationParams object
// with the ability to set a timeout on a request.
func NewRemoveLocationParamsWithTimeout(timeout time.Duration) *RemoveLocationParams {
	return &RemoveLocationParams{
		timeout: timeout,
	}
}

// NewRemoveLocationParamsWithContext creates a new RemoveLocationParams object
// with the ability to set a context for a request.
func NewRemoveLocationParamsWithContext(ctx context.Context) *RemoveLocationParams {
	return &RemoveLocationParams{
		Context: ctx,
	}
}

// NewRemoveLocationParamsWithHTTPClient creates a new RemoveLocationParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveLocationParamsWithHTTPClient(client *http.Client) *RemoveLocationParams {
	return &RemoveLocationParams{
		HTTPClient: client,
	}
}

/* RemoveLocationParams contains all the parameters to send to the API endpoint
   for the remove location operation.

   Typically these are written to a http.Request.
*/
type RemoveLocationParams struct {

	// Body.
	Body RemoveLocationBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove location params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveLocationParams) WithDefaults() *RemoveLocationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove location params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveLocationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove location params
func (o *RemoveLocationParams) WithTimeout(timeout time.Duration) *RemoveLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove location params
func (o *RemoveLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove location params
func (o *RemoveLocationParams) WithContext(ctx context.Context) *RemoveLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove location params
func (o *RemoveLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove location params
func (o *RemoveLocationParams) WithHTTPClient(client *http.Client) *RemoveLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove location params
func (o *RemoveLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove location params
func (o *RemoveLocationParams) WithBody(body RemoveLocationBody) *RemoveLocationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove location params
func (o *RemoveLocationParams) SetBody(body RemoveLocationBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
