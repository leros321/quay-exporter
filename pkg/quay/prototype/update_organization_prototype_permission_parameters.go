// Code generated by go-swagger; DO NOT EDIT.

package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/leros321/quay-exporter/pkg/models"
)

// NewUpdateOrganizationPrototypePermissionParams creates a new UpdateOrganizationPrototypePermissionParams object
// with the default values initialized.
func NewUpdateOrganizationPrototypePermissionParams() *UpdateOrganizationPrototypePermissionParams {
	var ()
	return &UpdateOrganizationPrototypePermissionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationPrototypePermissionParamsWithTimeout creates a new UpdateOrganizationPrototypePermissionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrganizationPrototypePermissionParamsWithTimeout(timeout time.Duration) *UpdateOrganizationPrototypePermissionParams {
	var ()
	return &UpdateOrganizationPrototypePermissionParams{

		timeout: timeout,
	}
}

// NewUpdateOrganizationPrototypePermissionParamsWithContext creates a new UpdateOrganizationPrototypePermissionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrganizationPrototypePermissionParamsWithContext(ctx context.Context) *UpdateOrganizationPrototypePermissionParams {
	var ()
	return &UpdateOrganizationPrototypePermissionParams{

		Context: ctx,
	}
}

// NewUpdateOrganizationPrototypePermissionParamsWithHTTPClient creates a new UpdateOrganizationPrototypePermissionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrganizationPrototypePermissionParamsWithHTTPClient(client *http.Client) *UpdateOrganizationPrototypePermissionParams {
	var ()
	return &UpdateOrganizationPrototypePermissionParams{
		HTTPClient: client,
	}
}

/*UpdateOrganizationPrototypePermissionParams contains all the parameters to send to the API endpoint
for the update organization prototype permission operation typically these are written to a http.Request
*/
type UpdateOrganizationPrototypePermissionParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.PrototypeUpdate
	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Prototypeid
	  The ID of the prototype

	*/
	Prototypeid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithTimeout(timeout time.Duration) *UpdateOrganizationPrototypePermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithContext(ctx context.Context) *UpdateOrganizationPrototypePermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithHTTPClient(client *http.Client) *UpdateOrganizationPrototypePermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithBody(body *models.PrototypeUpdate) *UpdateOrganizationPrototypePermissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) SetBody(body *models.PrototypeUpdate) {
	o.Body = body
}

// WithOrgname adds the orgname to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithOrgname(orgname string) *UpdateOrganizationPrototypePermissionParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WithPrototypeid adds the prototypeid to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) WithPrototypeid(prototypeid string) *UpdateOrganizationPrototypePermissionParams {
	o.SetPrototypeid(prototypeid)
	return o
}

// SetPrototypeid adds the prototypeid to the update organization prototype permission params
func (o *UpdateOrganizationPrototypePermissionParams) SetPrototypeid(prototypeid string) {
	o.Prototypeid = prototypeid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationPrototypePermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	// path param prototypeid
	if err := r.SetPathParam("prototypeid", o.Prototypeid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
