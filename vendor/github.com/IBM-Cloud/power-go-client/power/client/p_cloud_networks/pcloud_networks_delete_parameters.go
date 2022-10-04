// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPcloudNetworksDeleteParams creates a new PcloudNetworksDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudNetworksDeleteParams() *PcloudNetworksDeleteParams {
	return &PcloudNetworksDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksDeleteParamsWithTimeout creates a new PcloudNetworksDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudNetworksDeleteParamsWithTimeout(timeout time.Duration) *PcloudNetworksDeleteParams {
	return &PcloudNetworksDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudNetworksDeleteParamsWithContext creates a new PcloudNetworksDeleteParams object
// with the ability to set a context for a request.
func NewPcloudNetworksDeleteParamsWithContext(ctx context.Context) *PcloudNetworksDeleteParams {
	return &PcloudNetworksDeleteParams{
		Context: ctx,
	}
}

// NewPcloudNetworksDeleteParamsWithHTTPClient creates a new PcloudNetworksDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudNetworksDeleteParamsWithHTTPClient(client *http.Client) *PcloudNetworksDeleteParams {
	return &PcloudNetworksDeleteParams{
		HTTPClient: client,
	}
}

/* PcloudNetworksDeleteParams contains all the parameters to send to the API endpoint
   for the pcloud networks delete operation.

   Typically these are written to a http.Request.
*/
type PcloudNetworksDeleteParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud networks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksDeleteParams) WithDefaults() *PcloudNetworksDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud networks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) WithTimeout(timeout time.Duration) *PcloudNetworksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) WithContext(ctx context.Context) *PcloudNetworksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) WithHTTPClient(client *http.Client) *PcloudNetworksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) WithNetworkID(networkID string) *PcloudNetworksDeleteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud networks delete params
func (o *PcloudNetworksDeleteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
