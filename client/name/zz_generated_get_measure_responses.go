// Code generated by go-swagger; DO NOT EDIT.

package name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetMeasureReader is a Reader for the GetMeasure structure.
type GetMeasureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeasureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMeasureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMeasureOK creates a GetMeasureOK with default headers values
func NewGetMeasureOK() *GetMeasureOK {
	return &GetMeasureOK{}
}

/*
GetMeasureOK describes a response with status code 200, with default header values.

OK
*/
type GetMeasureOK struct {
	Payload string
}

func (o *GetMeasureOK) Error() string {
	return fmt.Sprintf("[GET /name/measure][%d] getMeasureOK  %+v", 200, o.Payload)
}
func (o *GetMeasureOK) GetPayload() string {
	return o.Payload
}

func (o *GetMeasureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
