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

// SharedProcessorPool shared processor pool
//
// swagger:model SharedProcessorPool
type SharedProcessorPool struct {

	// The amount of allocated processor cores for the Shared Processor Pool
	// Required: true
	AllocatedCores *float64 `json:"allocatedCores"`

	// The amount of available processor cores for the Shared Processor Pool
	// Required: true
	AvailableCores *float64 `json:"availableCores"`

	// The host group the host belongs to
	HostGroup string `json:"hostGroup,omitempty"`

	// The ID of the host where the Shared Processor Pool resides
	HostID int64 `json:"hostID,omitempty"`

	// The id of the Shared Processor Pool
	// Required: true
	ID *string `json:"id"`

	// The name of the Shared Processor Pool
	// Required: true
	Name *string `json:"name"`

	// The amount of reserved processor cores for the Shared Processor Pool
	// Required: true
	ReservedCores *int64 `json:"reservedCores"`

	// list of Shared Processor Pool Placement Groups
	SharedProcessorPoolPlacementGroups []*SharedProcessorPoolPlacementGroup `json:"sharedProcessorPoolPlacementGroups"`

	// The status of the Shared Processor Pool
	Status string `json:"status,omitempty"`

	// The status details of the Shared Processor Pool
	StatusDetail string `json:"statusDetail,omitempty"`
}

// Validate validates this shared processor pool
func (m *SharedProcessorPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailableCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReservedCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedProcessorPoolPlacementGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SharedProcessorPool) validateAllocatedCores(formats strfmt.Registry) error {

	if err := validate.Required("allocatedCores", "body", m.AllocatedCores); err != nil {
		return err
	}

	return nil
}

func (m *SharedProcessorPool) validateAvailableCores(formats strfmt.Registry) error {

	if err := validate.Required("availableCores", "body", m.AvailableCores); err != nil {
		return err
	}

	return nil
}

func (m *SharedProcessorPool) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SharedProcessorPool) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SharedProcessorPool) validateReservedCores(formats strfmt.Registry) error {

	if err := validate.Required("reservedCores", "body", m.ReservedCores); err != nil {
		return err
	}

	return nil
}

func (m *SharedProcessorPool) validateSharedProcessorPoolPlacementGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.SharedProcessorPoolPlacementGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.SharedProcessorPoolPlacementGroups); i++ {
		if swag.IsZero(m.SharedProcessorPoolPlacementGroups[i]) { // not required
			continue
		}

		if m.SharedProcessorPoolPlacementGroups[i] != nil {
			if err := m.SharedProcessorPoolPlacementGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sharedProcessorPoolPlacementGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sharedProcessorPoolPlacementGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this shared processor pool based on the context it is used
func (m *SharedProcessorPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSharedProcessorPoolPlacementGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SharedProcessorPool) contextValidateSharedProcessorPoolPlacementGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SharedProcessorPoolPlacementGroups); i++ {

		if m.SharedProcessorPoolPlacementGroups[i] != nil {

			if swag.IsZero(m.SharedProcessorPoolPlacementGroups[i]) { // not required
				return nil
			}

			if err := m.SharedProcessorPoolPlacementGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sharedProcessorPoolPlacementGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sharedProcessorPoolPlacementGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SharedProcessorPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharedProcessorPool) UnmarshalBinary(b []byte) error {
	var res SharedProcessorPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
