// Code generated by go-swagger; DO NOT EDIT.

package rules

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

// NewUpdateAlertRuleParams creates a new UpdateAlertRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAlertRuleParams() *UpdateAlertRuleParams {
	return &UpdateAlertRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAlertRuleParamsWithTimeout creates a new UpdateAlertRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateAlertRuleParamsWithTimeout(timeout time.Duration) *UpdateAlertRuleParams {
	return &UpdateAlertRuleParams{
		timeout: timeout,
	}
}

// NewUpdateAlertRuleParamsWithContext creates a new UpdateAlertRuleParams object
// with the ability to set a context for a request.
func NewUpdateAlertRuleParamsWithContext(ctx context.Context) *UpdateAlertRuleParams {
	return &UpdateAlertRuleParams{
		Context: ctx,
	}
}

// NewUpdateAlertRuleParamsWithHTTPClient creates a new UpdateAlertRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAlertRuleParamsWithHTTPClient(client *http.Client) *UpdateAlertRuleParams {
	return &UpdateAlertRuleParams{
		HTTPClient: client,
	}
}

/* UpdateAlertRuleParams contains all the parameters to send to the API endpoint
   for the update alert rule operation.

   Typically these are written to a http.Request.
*/
type UpdateAlertRuleParams struct {

	// Body.
	Body UpdateAlertRuleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAlertRuleParams) WithDefaults() *UpdateAlertRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAlertRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update alert rule params
func (o *UpdateAlertRuleParams) WithTimeout(timeout time.Duration) *UpdateAlertRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update alert rule params
func (o *UpdateAlertRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update alert rule params
func (o *UpdateAlertRuleParams) WithContext(ctx context.Context) *UpdateAlertRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update alert rule params
func (o *UpdateAlertRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update alert rule params
func (o *UpdateAlertRuleParams) WithHTTPClient(client *http.Client) *UpdateAlertRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update alert rule params
func (o *UpdateAlertRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update alert rule params
func (o *UpdateAlertRuleParams) WithBody(body UpdateAlertRuleBody) *UpdateAlertRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update alert rule params
func (o *UpdateAlertRuleParams) SetBody(body UpdateAlertRuleBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAlertRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
