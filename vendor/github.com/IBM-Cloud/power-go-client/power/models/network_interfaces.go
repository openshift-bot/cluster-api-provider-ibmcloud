// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkInterfaces network interfaces
//
// swagger:model NetworkInterfaces
type NetworkInterfaces struct {

	// Network Interfaces
	// Required: true
	Interfaces []*NetworkInterface `json:"interfaces"`
}

// Validate validates this network interfaces
func (m *NetworkInterfaces) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInterfaces) validateInterfaces(formats strfmt.Registry) error {

	if err := validate.Required("interfaces", "body", m.Interfaces); err != nil {
		return err
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network interfaces based on the context it is used
func (m *NetworkInterfaces) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInterfaces) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {

			if swag.IsZero(m.Interfaces[i]) { // not required
				return nil
			}

			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfaces) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfaces) UnmarshalBinary(b []byte) error {
	var res NetworkInterfaces
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
