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

// NewStorageServiceKeyspaceCompactionByKeyspacePostParams creates a new StorageServiceKeyspaceCompactionByKeyspacePostParams object
// with the default values initialized.
func NewStorageServiceKeyspaceCompactionByKeyspacePostParams() *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceCompactionByKeyspacePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceKeyspaceCompactionByKeyspacePostParamsWithTimeout creates a new StorageServiceKeyspaceCompactionByKeyspacePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceKeyspaceCompactionByKeyspacePostParamsWithTimeout(timeout time.Duration) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceCompactionByKeyspacePostParams{

		timeout: timeout,
	}
}

// NewStorageServiceKeyspaceCompactionByKeyspacePostParamsWithContext creates a new StorageServiceKeyspaceCompactionByKeyspacePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceKeyspaceCompactionByKeyspacePostParamsWithContext(ctx context.Context) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceCompactionByKeyspacePostParams{

		Context: ctx,
	}
}

// NewStorageServiceKeyspaceCompactionByKeyspacePostParamsWithHTTPClient creates a new StorageServiceKeyspaceCompactionByKeyspacePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceKeyspaceCompactionByKeyspacePostParamsWithHTTPClient(client *http.Client) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceKeyspaceCompactionByKeyspacePostParams{
		HTTPClient: client,
	}
}

/*StorageServiceKeyspaceCompactionByKeyspacePostParams contains all the parameters to send to the API endpoint
for the storage service keyspace compaction by keyspace post operation typically these are written to a http.Request
*/
type StorageServiceKeyspaceCompactionByKeyspacePostParams struct {

	/*Cf
	  Comma seperated column family names

	*/
	Cf *string
	/*Keyspace
	  The keyspace to query about

	*/
	Keyspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) WithTimeout(timeout time.Duration) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) WithContext(ctx context.Context) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) WithHTTPClient(client *http.Client) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCf adds the cf to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) WithCf(cf *string) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	o.SetCf(cf)
	return o
}

// SetCf adds the cf to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) SetCf(cf *string) {
	o.Cf = cf
}

// WithKeyspace adds the keyspace to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) WithKeyspace(keyspace string) *StorageServiceKeyspaceCompactionByKeyspacePostParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the storage service keyspace compaction by keyspace post params
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceKeyspaceCompactionByKeyspacePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cf != nil {

		// query param cf
		var qrCf string
		if o.Cf != nil {
			qrCf = *o.Cf
		}
		qCf := qrCf
		if qCf != "" {
			if err := r.SetQueryParam("cf", qCf); err != nil {
				return err
			}
		}

	}

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
