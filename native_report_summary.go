// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NativeReportSummary The summary for the native report
// swagger:model NativeReportSummary
type NativeReportSummary struct {

	// The complete percent of the scanning which value is between 0 and 100
	CompletePercent int64 `json:"complete_percent,omitempty"`

	// The seconds spent for generating the report
	Duration int64 `json:"duration,omitempty"`

	// The end time of the scan process that generating report
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// id of the native scan report
	ReportID string `json:"report_id,omitempty"`

	// The status of the report generating process
	ScanStatus string `json:"scan_status,omitempty"`

	// scanner
	Scanner *Scanner `json:"scanner,omitempty"`

	// The overall severity
	Severity string `json:"severity,omitempty"`

	// The start time of the scan process that generating report
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// summary
	Summary *VulnerabilitySummary `json:"summary,omitempty"`
}

// Validate validates this native report summary
func (m *NativeReportSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeReportSummary) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NativeReportSummary) validateScanner(formats strfmt.Registry) error {

	if swag.IsZero(m.Scanner) { // not required
		return nil
	}

	if m.Scanner != nil {
		if err := m.Scanner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner")
			}
			return err
		}
	}

	return nil
}

func (m *NativeReportSummary) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NativeReportSummary) validateSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeReportSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeReportSummary) UnmarshalBinary(b []byte) error {
	var res NativeReportSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
