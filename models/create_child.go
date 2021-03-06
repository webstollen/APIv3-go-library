// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateChild create child
// swagger:model createChild
type CreateChild struct {

	// Company name to use to create the child account
	// Required: true
	CompanyName *string `json:"companyName"`

	// Email address to create the child account
	// Required: true
	Email *strfmt.Email `json:"email"`

	// First name to use to create the child account
	// Required: true
	FirstName *string `json:"firstName"`

	// Last name to use to create the child account
	// Required: true
	LastName *string `json:"lastName"`

	// Password for the child account to login
	// Required: true
	Password *strfmt.Password `json:"password"`
}

// Validate validates this create child
func (m *CreateChild) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanyName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateChild) validateCompanyName(formats strfmt.Registry) error {

	if err := validate.Required("companyName", "body", m.CompanyName); err != nil {
		return err
	}

	return nil
}

func (m *CreateChild) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateChild) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *CreateChild) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *CreateChild) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateChild) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateChild) UnmarshalBinary(b []byte) error {
	var res CreateChild
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
