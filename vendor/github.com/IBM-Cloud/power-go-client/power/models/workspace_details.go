// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkspaceDetails workspace details
//
// swagger:model WorkspaceDetails
type WorkspaceDetails struct {

	// Workspace creation date
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// The Workspace crn
	// Required: true
	Crn *string `json:"crn"`

	// Link to Workspace Resource
	Href string `json:"href,omitempty"`

	// The Workspace Network Security Groups information
	NetworkSecurityGroups *WorkspaceNetworkSecurityGroupsDetails `json:"networkSecurityGroups,omitempty"`

	// The Workspace Power Edge Router information
	PowerEdgeRouter *WorkspacePowerEdgeRouterDetails `json:"powerEdgeRouter,omitempty"`
}

// Validate validates this workspace details
func (m *WorkspaceDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerEdgeRouter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceDetails) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceDetails) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceDetails) validateNetworkSecurityGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkSecurityGroups) { // not required
		return nil
	}

	if m.NetworkSecurityGroups != nil {
		if err := m.NetworkSecurityGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSecurityGroups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkSecurityGroups")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceDetails) validatePowerEdgeRouter(formats strfmt.Registry) error {
	if swag.IsZero(m.PowerEdgeRouter) { // not required
		return nil
	}

	if m.PowerEdgeRouter != nil {
		if err := m.PowerEdgeRouter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("powerEdgeRouter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("powerEdgeRouter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this workspace details based on the context it is used
func (m *WorkspaceDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkSecurityGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerEdgeRouter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceDetails) contextValidateNetworkSecurityGroups(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkSecurityGroups != nil {

		if swag.IsZero(m.NetworkSecurityGroups) { // not required
			return nil
		}

		if err := m.NetworkSecurityGroups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkSecurityGroups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkSecurityGroups")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceDetails) contextValidatePowerEdgeRouter(ctx context.Context, formats strfmt.Registry) error {

	if m.PowerEdgeRouter != nil {

		if swag.IsZero(m.PowerEdgeRouter) { // not required
			return nil
		}

		if err := m.PowerEdgeRouter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("powerEdgeRouter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("powerEdgeRouter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceDetails) UnmarshalBinary(b []byte) error {
	var res WorkspaceDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}