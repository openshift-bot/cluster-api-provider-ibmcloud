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

// SnapshotCreateResponse snapshot create response
//
// swagger:model SnapshotCreateResponse
type SnapshotCreateResponse struct {

	// ID of the PVM instance snapshot
	// Required: true
	SnapshotID *string `json:"snapshotID"`
}

// Validate validates this snapshot create response
func (m *SnapshotCreateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotCreateResponse) validateSnapshotID(formats strfmt.Registry) error {

	if err := validate.Required("snapshotID", "body", m.SnapshotID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this snapshot create response based on context it is used
func (m *SnapshotCreateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotCreateResponse) UnmarshalBinary(b []byte) error {
	var res SnapshotCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
