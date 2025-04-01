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

// AssetsUpdateRequest assets update request
//
// swagger:model assets.UpdateRequest
type AssetsUpdateRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// environments
	Environments map[string]AssetEnvironment `json:"environments,omitempty"`

	// external links
	ExternalLinks []*AssetExternalLink `json:"external_links"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// providers
	Providers []string `json:"providers"`

	// schema
	Schema interface{} `json:"schema,omitempty"`

	// sources
	Sources []*AssetAssetSource `json:"sources"`

	// tags
	Tags []string `json:"tags"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this assets update request
func (m *AssetsUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetsUpdateRequest) validateEnvironments(formats strfmt.Registry) error {
	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for k := range m.Environments {

		if err := validate.Required("environments"+"."+k, "body", m.Environments[k]); err != nil {
			return err
		}
		if val, ok := m.Environments[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssetsUpdateRequest) validateExternalLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalLinks) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalLinks); i++ {
		if swag.IsZero(m.ExternalLinks[i]) { // not required
			continue
		}

		if m.ExternalLinks[i] != nil {
			if err := m.ExternalLinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssetsUpdateRequest) validateSources(formats strfmt.Registry) error {
	if swag.IsZero(m.Sources) { // not required
		return nil
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this assets update request based on the context it is used
func (m *AssetsUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvironments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetsUpdateRequest) contextValidateEnvironments(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Environments {

		if val, ok := m.Environments[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *AssetsUpdateRequest) contextValidateExternalLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalLinks); i++ {

		if m.ExternalLinks[i] != nil {

			if swag.IsZero(m.ExternalLinks[i]) { // not required
				return nil
			}

			if err := m.ExternalLinks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssetsUpdateRequest) contextValidateSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sources); i++ {

		if m.Sources[i] != nil {

			if swag.IsZero(m.Sources[i]) { // not required
				return nil
			}

			if err := m.Sources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetsUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetsUpdateRequest) UnmarshalBinary(b []byte) error {
	var res AssetsUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
