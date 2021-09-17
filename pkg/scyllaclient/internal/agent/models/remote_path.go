// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RemotePath remote path
// swagger:model RemotePath
type RemotePath struct {

	// A remote name string eg. drive:
	Fs string `json:"fs,omitempty"`

	// A path within that remote eg. dir
	Remote string `json:"remote,omitempty"`
}

// Validate validates this remote path
func (m *RemotePath) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemotePath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemotePath) UnmarshalBinary(b []byte) error {
	var res RemotePath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
