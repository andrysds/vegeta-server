// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReportResponse report response
// swagger:model ReportResponse
type ReportResponse struct {

	// The attack ID used to fetch status and reports.
	ID string `json:"id,omitempty"`

	// report
	Report *Report `json:"report,omitempty"`
}

// Validate validates this report response
func (m *ReportResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportResponse) validateReport(formats strfmt.Registry) error {

	if swag.IsZero(m.Report) { // not required
		return nil
	}

	if m.Report != nil {
		if err := m.Report.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportResponse) UnmarshalBinary(b []byte) error {
	var res ReportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}