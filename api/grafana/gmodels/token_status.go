// Code generated by go-swagger; DO NOT EDIT.

package gmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// TokenStatus token status
//
// swagger:model TokenStatus
type TokenStatus int64

// Validate validates this token status
func (m TokenStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this token status based on context it is used
func (m TokenStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}