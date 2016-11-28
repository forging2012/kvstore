package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Error the error model is a model for all the error responses coming from kvstore
//
// swagger:model error
type Error struct {

	// cause
	Cause *Error `json:"cause,omitempty"`

	// The error code
	// Required: true
	Code *int64 `json:"code"`

	// link to help page explaining the error in more detail
	HelpURL strfmt.URI `json:"helpUrl,omitempty"`

	// The error message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateCause(formats strfmt.Registry) error {

	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {

		if err := m.Cause.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Error) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}
