// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud tenants ssh keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new p cloud tenants ssh keys API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new p cloud tenants ssh keys API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for p cloud tenants ssh keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudTenantsSshkeysDelete(params *PcloudTenantsSshkeysDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysDeleteOK, error)

	PcloudTenantsSshkeysGet(params *PcloudTenantsSshkeysGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysGetOK, error)

	PcloudTenantsSshkeysGetall(params *PcloudTenantsSshkeysGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysGetallOK, error)

	PcloudTenantsSshkeysPost(params *PcloudTenantsSshkeysPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysPostOK, *PcloudTenantsSshkeysPostCreated, error)

	PcloudTenantsSshkeysPut(params *PcloudTenantsSshkeysPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PcloudTenantsSshkeysDelete deletes a tenant s SSH key
*/
func (a *Client) PcloudTenantsSshkeysDelete(params *PcloudTenantsSshkeysDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudTenantsSshkeysDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.tenants.sshkeys.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudTenantsSshkeysDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudTenantsSshkeysDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.tenants.sshkeys.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudTenantsSshkeysGet gets a tenant s SSH key by name
*/
func (a *Client) PcloudTenantsSshkeysGet(params *PcloudTenantsSshkeysGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudTenantsSshkeysGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.tenants.sshkeys.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudTenantsSshkeysGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudTenantsSshkeysGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.tenants.sshkeys.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudTenantsSshkeysGetall lists a tenant s SSH keys
*/
func (a *Client) PcloudTenantsSshkeysGetall(params *PcloudTenantsSshkeysGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudTenantsSshkeysGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.tenants.sshkeys.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/tenants/{tenant_id}/sshkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudTenantsSshkeysGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudTenantsSshkeysGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.tenants.sshkeys.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudTenantsSshkeysPost adds a new SSH key to the tenant
*/
func (a *Client) PcloudTenantsSshkeysPost(params *PcloudTenantsSshkeysPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysPostOK, *PcloudTenantsSshkeysPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudTenantsSshkeysPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.tenants.sshkeys.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/tenants/{tenant_id}/sshkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudTenantsSshkeysPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudTenantsSshkeysPostOK:
		return value, nil, nil
	case *PcloudTenantsSshkeysPostCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for p_cloud_tenants_ssh_keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudTenantsSshkeysPut updates an SSH key
*/
func (a *Client) PcloudTenantsSshkeysPut(params *PcloudTenantsSshkeysPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudTenantsSshkeysPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudTenantsSshkeysPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.tenants.sshkeys.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudTenantsSshkeysPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudTenantsSshkeysPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.tenants.sshkeys.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
