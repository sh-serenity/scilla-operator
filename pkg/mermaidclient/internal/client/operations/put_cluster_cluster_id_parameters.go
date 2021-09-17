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

	models "github.com/scylladb/scylla-operator/pkg/mermaidclient/internal/models"
)

// NewPutClusterClusterIDParams creates a new PutClusterClusterIDParams object
// with the default values initialized.
func NewPutClusterClusterIDParams() *PutClusterClusterIDParams {
	var ()
	return &PutClusterClusterIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClusterClusterIDParamsWithTimeout creates a new PutClusterClusterIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClusterClusterIDParamsWithTimeout(timeout time.Duration) *PutClusterClusterIDParams {
	var ()
	return &PutClusterClusterIDParams{

		timeout: timeout,
	}
}

// NewPutClusterClusterIDParamsWithContext creates a new PutClusterClusterIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClusterClusterIDParamsWithContext(ctx context.Context) *PutClusterClusterIDParams {
	var ()
	return &PutClusterClusterIDParams{

		Context: ctx,
	}
}

// NewPutClusterClusterIDParamsWithHTTPClient creates a new PutClusterClusterIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClusterClusterIDParamsWithHTTPClient(client *http.Client) *PutClusterClusterIDParams {
	var ()
	return &PutClusterClusterIDParams{
		HTTPClient: client,
	}
}

/*PutClusterClusterIDParams contains all the parameters to send to the API endpoint
for the put cluster cluster ID operation typically these are written to a http.Request
*/
type PutClusterClusterIDParams struct {

	/*Cluster*/
	Cluster *models.Cluster
	/*ClusterID*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) WithTimeout(timeout time.Duration) *PutClusterClusterIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) WithContext(ctx context.Context) *PutClusterClusterIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) WithHTTPClient(client *http.Client) *PutClusterClusterIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) WithCluster(cluster *models.Cluster) *PutClusterClusterIDParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) SetCluster(cluster *models.Cluster) {
	o.Cluster = cluster
}

// WithClusterID adds the clusterID to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) WithClusterID(clusterID string) *PutClusterClusterIDParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the put cluster cluster ID params
func (o *PutClusterClusterIDParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *PutClusterClusterIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {
		if err := r.SetBodyParam(o.Cluster); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
