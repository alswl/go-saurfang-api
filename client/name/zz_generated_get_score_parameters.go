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

// NewGetScoreParams creates a new GetScoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScoreParams() *GetScoreParams {
	return &GetScoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScoreParamsWithTimeout creates a new GetScoreParams object
// with the ability to set a timeout on a request.
func NewGetScoreParamsWithTimeout(timeout time.Duration) *GetScoreParams {
	return &GetScoreParams{
		timeout: timeout,
	}
}

// NewGetScoreParamsWithContext creates a new GetScoreParams object
// with the ability to set a context for a request.
func NewGetScoreParamsWithContext(ctx ccontext.Context) *GetScoreParams {
	return &GetScoreParams{
		Context: ctx,
	}
}

// NewGetScoreParamsWithHTTPClient creates a new GetScoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScoreParamsWithHTTPClient(client *http.Client) *GetScoreParams {
	return &GetScoreParams{
		HTTPClient: client,
	}
}

/*
GetScoreParams contains all the parameters to send to the API endpoint

	for the get score operation.

	Typically these are written to a http.Request.
*/
type GetScoreParams struct {

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

// WithDefaults hydrates default values in the get score params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScoreParams) WithDefaults() *GetScoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get score params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get score params
func (o *GetScoreParams) WithTimeout(timeout time.Duration) *GetScoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get score params
func (o *GetScoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get score params
func (o *GetScoreParams) WithContext(ctx ccontext.Context) *GetScoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get score params
func (o *GetScoreParams) SetContext(ctx ccontext.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get score params
func (o *GetScoreParams) WithHTTPClient(client *http.Client) *GetScoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get score params
func (o *GetScoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFirstName adds the firstName to the get score params
func (o *GetScoreParams) WithFirstName(firstName string) *GetScoreParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the get score params
func (o *GetScoreParams) SetFirstName(firstName string) {
	o.FirstName = firstName
}

// WithLastName adds the lastName to the get score params
func (o *GetScoreParams) WithLastName(lastName string) *GetScoreParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the get score params
func (o *GetScoreParams) SetLastName(lastName string) {
	o.LastName = lastName
}

// WriteToRequest writes these params to a swagger request
func (o *GetScoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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