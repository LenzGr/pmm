// Code generated by go-swagger; DO NOT EDIT.

package templates

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

// NewUpdateTemplateParams creates a new UpdateTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTemplateParams() *UpdateTemplateParams {
	return &UpdateTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTemplateParamsWithTimeout creates a new UpdateTemplateParams object
// with the ability to set a timeout on a request.
func NewUpdateTemplateParamsWithTimeout(timeout time.Duration) *UpdateTemplateParams {
	return &UpdateTemplateParams{
		timeout: timeout,
	}
}

// NewUpdateTemplateParamsWithContext creates a new UpdateTemplateParams object
// with the ability to set a context for a request.
func NewUpdateTemplateParamsWithContext(ctx context.Context) *UpdateTemplateParams {
	return &UpdateTemplateParams{
		Context: ctx,
	}
}

// NewUpdateTemplateParamsWithHTTPClient creates a new UpdateTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTemplateParamsWithHTTPClient(client *http.Client) *UpdateTemplateParams {
	return &UpdateTemplateParams{
		HTTPClient: client,
	}
}

/* UpdateTemplateParams contains all the parameters to send to the API endpoint
   for the update template operation.

   Typically these are written to a http.Request.
*/
type UpdateTemplateParams struct {
	// Body.
	Body UpdateTemplateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTemplateParams) WithDefaults() *UpdateTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update template params
func (o *UpdateTemplateParams) WithTimeout(timeout time.Duration) *UpdateTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update template params
func (o *UpdateTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update template params
func (o *UpdateTemplateParams) WithContext(ctx context.Context) *UpdateTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update template params
func (o *UpdateTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update template params
func (o *UpdateTemplateParams) WithHTTPClient(client *http.Client) *UpdateTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update template params
func (o *UpdateTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update template params
func (o *UpdateTemplateParams) WithBody(body UpdateTemplateBody) *UpdateTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update template params
func (o *UpdateTemplateParams) SetBody(body UpdateTemplateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
