// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssetAvailableFilters asset available filters
//
// swagger:model asset.AvailableFilters
type AssetAvailableFilters struct {

	// providers
	Providers map[string]int64 `json:"providers,omitempty"`

	// tags
	Tags map[string]int64 `json:"tags,omitempty"`

	// types
	Types map[string]int64 `json:"types,omitempty"`
}

// Validate validates this asset available filters
func (m *AssetAvailableFilters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this asset available filters based on context it is used
func (m *AssetAvailableFilters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssetAvailableFilters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetAvailableFilters) UnmarshalBinary(b []byte) error {
	var res AssetAvailableFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
