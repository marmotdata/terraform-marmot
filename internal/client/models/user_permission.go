// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserPermission user permission
//
// swagger:model user.Permission
type UserPermission struct {

	// action
	Action string `json:"action,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource type
	ResourceType string `json:"resource_type,omitempty"`
}

// Validate validates this user permission
func (m *UserPermission) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user permission based on context it is used
func (m *UserPermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPermission) UnmarshalBinary(b []byte) error {
	var res UserPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
