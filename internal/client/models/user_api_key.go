// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserAPIKey user API key
//
// swagger:model user.APIKey
type UserAPIKey struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// expires at
	ExpiresAt string `json:"expires_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// last used at
	LastUsedAt string `json:"last_used_at,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this user API key
func (m *UserAPIKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user API key based on context it is used
func (m *UserAPIKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserAPIKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAPIKey) UnmarshalBinary(b []byte) error {
	var res UserAPIKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
