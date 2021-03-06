// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateEmailCampaignRecipients update email campaign recipients
// swagger:model updateEmailCampaignRecipients
type UpdateEmailCampaignRecipients struct {

	// List ids which have to be excluded from a campaign
	ExclusionListIds []int64 `json:"exclusionListIds"`

	// Lists Ids to send the campaign to. REQUIRED if already not present in campaign and scheduledAt is not empty
	ListIds []int64 `json:"listIds"`
}

// Validate validates this update email campaign recipients
func (m *UpdateEmailCampaignRecipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEmailCampaignRecipients) validateExclusionListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ExclusionListIds) { // not required
		return nil
	}

	return nil
}

func (m *UpdateEmailCampaignRecipients) validateListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ListIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateEmailCampaignRecipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateEmailCampaignRecipients) UnmarshalBinary(b []byte) error {
	var res UpdateEmailCampaignRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
