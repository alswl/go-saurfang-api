// Code generated by go-swagger; DO NOT EDIT.

package name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	ccontext "context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetMeasureParams creates a new GetMeasureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMeasureParams() *GetMeasureParams {
	return &GetMeasureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeasureParamsWithTimeout creates a new GetMeasureParams object
// with the ability to set a timeout on a request.
func NewGetMeasureParamsWithTimeout(timeout time.Duration) *GetMeasureParams {
	return &GetMeasureParams{
		timeout: timeout,
	}
}

// NewGetMeasureParamsWithContext creates a new GetMeasureParams object
// with the ability to set a context for a request.
func NewGetMeasureParamsWithContext(ctx ccontext.Context) *GetMeasureParams {
	return &GetMeasureParams{
		Context: ctx,
	}
}

// NewGetMeasureParamsWithHTTPClient creates a new GetMeasureParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMeasureParamsWithHTTPClient(client *http.Client) *GetMeasureParams {
	return &GetMeasureParams{
		HTTPClient: client,
	}
}

/*
GetMeasureParams contains all the parameters to send to the API endpoint

	for the get measure operation.

	Typically these are written to a http.Request.
*/
type GetMeasureParams struct {

	/* FirstName.

	   first name
	*/
	FirstName string

	/* LastName.

	   last name
	*/
	LastName string

	timeout    time.Duration
	Context    ccontext.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get measure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMeasureParams) WithDefaults() *GetMeasureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get measure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMeasureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get measure params
func (o *GetMeasureParams) WithTimeout(timeout time.Duration) *GetMeasureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get measure params
func (o *GetMeasureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get measure params
func (o *GetMeasureParams) WithContext(ctx ccontext.Context) *GetMeasureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get measure params
func (o *GetMeasureParams) SetContext(ctx ccontext.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get measure params
func (o *GetMeasureParams) WithHTTPClient(client *http.Client) *GetMeasureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get measure params
func (o *GetMeasureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFirstName adds the firstName to the get measure params
func (o *GetMeasureParams) WithFirstName(firstName string) *GetMeasureParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the get measure params
func (o *GetMeasureParams) SetFirstName(firstName string) {
	o.FirstName = firstName
}

// WithLastName adds the lastName to the get measure params
func (o *GetMeasureParams) WithLastName(lastName string) *GetMeasureParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the get measure params
func (o *GetMeasureParams) SetLastName(lastName string) {
	o.LastName = lastName
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeasureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param first_name
	qrFirstName := o.FirstName
	qFirstName := qrFirstName
	if qFirstName != "" {

		if err := r.SetQueryParam("first_name", qFirstName); err != nil {
			return err
		}
	}

	// query param last_name
	qrLastName := o.LastName
	qLastName := qrLastName
	if qLastName != "" {

		if err := r.SetQueryParam("last_name", qLastName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
