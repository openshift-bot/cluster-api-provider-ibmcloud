// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

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

// NewPcloudCloudinstancesImagesDeleteParams creates a new PcloudCloudinstancesImagesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudinstancesImagesDeleteParams() *PcloudCloudinstancesImagesDeleteParams {
	return &PcloudCloudinstancesImagesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesImagesDeleteParamsWithTimeout creates a new PcloudCloudinstancesImagesDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudinstancesImagesDeleteParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesImagesDeleteParams {
	return &PcloudCloudinstancesImagesDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudCloudinstancesImagesDeleteParamsWithContext creates a new PcloudCloudinstancesImagesDeleteParams object
// with the ability to set a context for a request.
func NewPcloudCloudinstancesImagesDeleteParamsWithContext(ctx context.Context) *PcloudCloudinstancesImagesDeleteParams {
	return &PcloudCloudinstancesImagesDeleteParams{
		Context: ctx,
	}
}

// NewPcloudCloudinstancesImagesDeleteParamsWithHTTPClient creates a new PcloudCloudinstancesImagesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudinstancesImagesDeleteParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesImagesDeleteParams {
	return &PcloudCloudinstancesImagesDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudCloudinstancesImagesDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud cloudinstances images delete operation.

	Typically these are written to a http.Request.
*/
type PcloudCloudinstancesImagesDeleteParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* ImageID.

	   Image ID of a image
	*/
	ImageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudinstances images delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesImagesDeleteParams) WithDefaults() *PcloudCloudinstancesImagesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudinstances images delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesImagesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesImagesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) WithContext(ctx context.Context) *PcloudCloudinstancesImagesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesImagesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesImagesDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithImageID adds the imageID to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) WithImageID(imageID string) *PcloudCloudinstancesImagesDeleteParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the pcloud cloudinstances images delete params
func (o *PcloudCloudinstancesImagesDeleteParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesImagesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param image_id
	if err := r.SetPathParam("image_id", o.ImageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
