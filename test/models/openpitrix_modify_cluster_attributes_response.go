// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyClusterAttributesResponse openpitrix modify cluster attributes response
// swagger:model openpitrixModifyClusterAttributesResponse
type OpenpitrixModifyClusterAttributesResponse struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`
}

// Validate validates this openpitrix modify cluster attributes response
func (m *OpenpitrixModifyClusterAttributesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyClusterAttributesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyClusterAttributesResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyClusterAttributesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
