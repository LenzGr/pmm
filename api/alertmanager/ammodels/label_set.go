// Code generated by go-swagger; DO NOT EDIT.

package ammodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// LabelSet label set
//
// swagger:model labelSet
type LabelSet map[string]string

// Validate validates this label set
func (m LabelSet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this label set based on context it is used
func (m LabelSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
