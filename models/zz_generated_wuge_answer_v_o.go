// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	ccontext "context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WugeAnswerVO wuge answer v o
//
// swagger:model WugeAnswerVO
type WugeAnswerVO struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// ge
	// Required: true
	Ge *string `json:"ge"`

	// lucky level
	// Required: true
	LuckyLevel *string `json:"luckyLevel"`

	// stroke
	// Required: true
	Stroke *int32 `json:"stroke"`
}

// Validate validates this wuge answer v o
func (m *WugeAnswerVO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuckyLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStroke(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WugeAnswerVO) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *WugeAnswerVO) validateGe(formats strfmt.Registry) error {

	if err := validate.Required("ge", "body", m.Ge); err != nil {
		return err
	}

	return nil
}

func (m *WugeAnswerVO) validateLuckyLevel(formats strfmt.Registry) error {

	if err := validate.Required("luckyLevel", "body", m.LuckyLevel); err != nil {
		return err
	}

	return nil
}

func (m *WugeAnswerVO) validateStroke(formats strfmt.Registry) error {

	if err := validate.Required("stroke", "body", m.Stroke); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wuge answer v o based on context it is used
func (m *WugeAnswerVO) ContextValidate(ctx ccontext.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WugeAnswerVO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WugeAnswerVO) UnmarshalBinary(b []byte) error {
	var res WugeAnswerVO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
