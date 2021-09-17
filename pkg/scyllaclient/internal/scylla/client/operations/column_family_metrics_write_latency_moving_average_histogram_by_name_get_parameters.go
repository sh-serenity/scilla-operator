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

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams() *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics write latency moving average histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics write latency moving average histogram by name get params
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
