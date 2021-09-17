// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/agent/models"
)

// NewOperationsPurgeParams creates a new OperationsPurgeParams object
// with the default values initialized.
func NewOperationsPurgeParams() *OperationsPurgeParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsPurgeParams{
		Async: asyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewOperationsPurgeParamsWithTimeout creates a new OperationsPurgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperationsPurgeParamsWithTimeout(timeout time.Duration) *OperationsPurgeParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsPurgeParams{
		Async: asyncDefault,

		timeout: timeout,
	}
}

// NewOperationsPurgeParamsWithContext creates a new OperationsPurgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperationsPurgeParamsWithContext(ctx context.Context) *OperationsPurgeParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsPurgeParams{
		Async: asyncDefault,

		Context: ctx,
	}
}

// NewOperationsPurgeParamsWithHTTPClient creates a new OperationsPurgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperationsPurgeParamsWithHTTPClient(client *http.Client) *OperationsPurgeParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsPurgeParams{
		Async:      asyncDefault,
		HTTPClient: client,
	}
}

/*OperationsPurgeParams contains all the parameters to send to the API endpoint
for the operations purge operation typically these are written to a http.Request
*/
type OperationsPurgeParams struct {

	/*Async
	  Async request

	*/
	Async bool
	/*Group
	  Place this operation under this stat group

	*/
	Group string
	/*Purge
	  purge

	*/
	Purge *models.RemotePath

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operations purge params
func (o *OperationsPurgeParams) WithTimeout(timeout time.Duration) *OperationsPurgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operations purge params
func (o *OperationsPurgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operations purge params
func (o *OperationsPurgeParams) WithContext(ctx context.Context) *OperationsPurgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operations purge params
func (o *OperationsPurgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operations purge params
func (o *OperationsPurgeParams) WithHTTPClient(client *http.Client) *OperationsPurgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operations purge params
func (o *OperationsPurgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsync adds the async to the operations purge params
func (o *OperationsPurgeParams) WithAsync(async bool) *OperationsPurgeParams {
	o.SetAsync(async)
	return o
}

// SetAsync adds the async to the operations purge params
func (o *OperationsPurgeParams) SetAsync(async bool) {
	o.Async = async
}

// WithGroup adds the group to the operations purge params
func (o *OperationsPurgeParams) WithGroup(group string) *OperationsPurgeParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the operations purge params
func (o *OperationsPurgeParams) SetGroup(group string) {
	o.Group = group
}

// WithPurge adds the purge to the operations purge params
func (o *OperationsPurgeParams) WithPurge(purge *models.RemotePath) *OperationsPurgeParams {
	o.SetPurge(purge)
	return o
}

// SetPurge adds the purge to the operations purge params
func (o *OperationsPurgeParams) SetPurge(purge *models.RemotePath) {
	o.Purge = purge
}

// WriteToRequest writes these params to a swagger request
func (o *OperationsPurgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param _async
	qrAsync := o.Async
	qAsync := swag.FormatBool(qrAsync)
	if qAsync != "" {
		if err := r.SetQueryParam("_async", qAsync); err != nil {
			return err
		}
	}

	// query param _group
	qrGroup := o.Group
	qGroup := qrGroup
	if qGroup != "" {
		if err := r.SetQueryParam("_group", qGroup); err != nil {
			return err
		}
	}

	if o.Purge != nil {
		if err := r.SetBodyParam(o.Purge); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
