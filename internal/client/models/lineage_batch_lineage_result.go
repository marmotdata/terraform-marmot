// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LineageBatchLineageResult lineage batch lineage result
//
// swagger:model lineage.BatchLineageResult
type LineageBatchLineageResult struct {

	// edge
	Edge *LineageLineageEdge `json:"edge,omitempty"`

	// "created", "duplicate", or "existing"
	Status string `json:"status,omitempty"`
}

// Validate validates this lineage batch lineage result
func (m *LineageBatchLineageResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineageBatchLineageResult) validateEdge(formats strfmt.Registry) error {
	if swag.IsZero(m.Edge) { // not required
		return nil
	}

	if m.Edge != nil {
		if err := m.Edge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this lineage batch lineage result based on the context it is used
func (m *LineageBatchLineageResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineageBatchLineageResult) contextValidateEdge(ctx context.Context, formats strfmt.Registry) error {

	if m.Edge != nil {

		if swag.IsZero(m.Edge) { // not required
			return nil
		}

		if err := m.Edge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LineageBatchLineageResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LineageBatchLineageResult) UnmarshalBinary(b []byte) error {
	var res LineageBatchLineageResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
