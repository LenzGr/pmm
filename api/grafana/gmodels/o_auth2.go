// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OAuth2 OAuth2 is the oauth2 client configuration.
//
// swagger:model OAuth2
type OAuth2 struct {

	// client ID
	ClientID string `json:"client_id,omitempty"`

	// client secret file
	ClientSecretFile string `json:"client_secret_file,omitempty"`

	// endpoint params
	EndpointParams map[string]string `json:"endpoint_params,omitempty"`

	// scopes
	Scopes []string `json:"scopes"`

	// TLS config
	TLSConfig *TLSConfig `json:"TLSConfig,omitempty"`

	// token URL
	TokenURL string `json:"token_url,omitempty"`

	// client secret
	ClientSecret Secret `json:"client_secret,omitempty"`
}

// Validate validates this o auth2
func (m *OAuth2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuth2) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

func (m *OAuth2) validateClientSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientSecret) { // not required
		return nil
	}

	if err := m.ClientSecret.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("client_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("client_secret")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o auth2 based on the context it is used
func (m *OAuth2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuth2) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TLSConfig != nil {
		if err := m.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TLSConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TLSConfig")
			}
			return err
		}
	}

	return nil
}

func (m *OAuth2) contextValidateClientSecret(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ClientSecret.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("client_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("client_secret")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OAuth2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuth2) UnmarshalBinary(b []byte) error {
	var res OAuth2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}