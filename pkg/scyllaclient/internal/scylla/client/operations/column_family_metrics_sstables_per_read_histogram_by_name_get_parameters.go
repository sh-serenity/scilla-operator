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

	strfmt "github.com/go-openapi/strfmt"
)

// NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams creates a new ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams() *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsSstablesPerReadHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics sstables per read histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics sstables per read histogram by name get params
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsSstablesPerReadHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
