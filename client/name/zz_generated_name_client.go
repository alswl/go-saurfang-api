// Code generated by go-swagger; DO NOT EDIT.

package name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new name API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for name API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetMeasure(params *GetMeasureParams, opts ...ClientOption) (*GetMeasureOK, error)

	GetScore(params *GetScoreParams, opts ...ClientOption) (*GetScoreOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetMeasure gets measure

get measure
*/
func (a *Client) GetMeasure(params *GetMeasureParams, opts ...ClientOption) (*GetMeasureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMeasureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMeasure",
		Method:             "GET",
		PathPattern:        "/name/measure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMeasureReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMeasureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMeasure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScore gets score

get score
*/
func (a *Client) GetScore(params *GetScoreParams, opts ...ClientOption) (*GetScoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getScore",
		Method:             "GET",
		PathPattern:        "/name/score",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetScoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}