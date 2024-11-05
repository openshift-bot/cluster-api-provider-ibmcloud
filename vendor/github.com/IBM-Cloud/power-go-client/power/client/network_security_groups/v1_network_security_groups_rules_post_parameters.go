// Code generated by go-swagger; DO NOT EDIT.

package network_security_groups

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewV1NetworkSecurityGroupsRulesPostParams creates a new V1NetworkSecurityGroupsRulesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1NetworkSecurityGroupsRulesPostParams() *V1NetworkSecurityGroupsRulesPostParams {
	return &V1NetworkSecurityGroupsRulesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1NetworkSecurityGroupsRulesPostParamsWithTimeout creates a new V1NetworkSecurityGroupsRulesPostParams object
// with the ability to set a timeout on a request.
func NewV1NetworkSecurityGroupsRulesPostParamsWithTimeout(timeout time.Duration) *V1NetworkSecurityGroupsRulesPostParams {
	return &V1NetworkSecurityGroupsRulesPostParams{
		timeout: timeout,
	}
}

// NewV1NetworkSecurityGroupsRulesPostParamsWithContext creates a new V1NetworkSecurityGroupsRulesPostParams object
// with the ability to set a context for a request.
func NewV1NetworkSecurityGroupsRulesPostParamsWithContext(ctx context.Context) *V1NetworkSecurityGroupsRulesPostParams {
	return &V1NetworkSecurityGroupsRulesPostParams{
		Context: ctx,
	}
}

// NewV1NetworkSecurityGroupsRulesPostParamsWithHTTPClient creates a new V1NetworkSecurityGroupsRulesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1NetworkSecurityGroupsRulesPostParamsWithHTTPClient(client *http.Client) *V1NetworkSecurityGroupsRulesPostParams {
	return &V1NetworkSecurityGroupsRulesPostParams{
		HTTPClient: client,
	}
}

/*
V1NetworkSecurityGroupsRulesPostParams contains all the parameters to send to the API endpoint

	for the v1 network security groups rules post operation.

	Typically these are written to a http.Request.
*/
type V1NetworkSecurityGroupsRulesPostParams struct {

	/* Body.

	   Parameters for adding a rule to a Network Security Group
	*/
	Body *models.NetworkSecurityGroupAddRule

	/* NetworkSecurityGroupID.

	   Network Security Group ID
	*/
	NetworkSecurityGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 network security groups rules post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworkSecurityGroupsRulesPostParams) WithDefaults() *V1NetworkSecurityGroupsRulesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 network security groups rules post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworkSecurityGroupsRulesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) WithTimeout(timeout time.Duration) *V1NetworkSecurityGroupsRulesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) WithContext(ctx context.Context) *V1NetworkSecurityGroupsRulesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) WithHTTPClient(client *http.Client) *V1NetworkSecurityGroupsRulesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) WithBody(body *models.NetworkSecurityGroupAddRule) *V1NetworkSecurityGroupsRulesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) SetBody(body *models.NetworkSecurityGroupAddRule) {
	o.Body = body
}

// WithNetworkSecurityGroupID adds the networkSecurityGroupID to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) WithNetworkSecurityGroupID(networkSecurityGroupID string) *V1NetworkSecurityGroupsRulesPostParams {
	o.SetNetworkSecurityGroupID(networkSecurityGroupID)
	return o
}

// SetNetworkSecurityGroupID adds the networkSecurityGroupId to the v1 network security groups rules post params
func (o *V1NetworkSecurityGroupsRulesPostParams) SetNetworkSecurityGroupID(networkSecurityGroupID string) {
	o.NetworkSecurityGroupID = networkSecurityGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *V1NetworkSecurityGroupsRulesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param network_security_group_id
	if err := r.SetPathParam("network_security_group_id", o.NetworkSecurityGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
