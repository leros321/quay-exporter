// Code generated by go-swagger; DO NOT EDIT.

package permission

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

// NewChangeUserPermissionsParams creates a new ChangeUserPermissionsParams object
// with the default values initialized.
func NewChangeUserPermissionsParams() *ChangeUserPermissionsParams {
	var ()
	return &ChangeUserPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeUserPermissionsParamsWithTimeout creates a new ChangeUserPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeUserPermissionsParamsWithTimeout(timeout time.Duration) *ChangeUserPermissionsParams {
	var ()
	return &ChangeUserPermissionsParams{

		timeout: timeout,
	}
}

// NewChangeUserPermissionsParamsWithContext creates a new ChangeUserPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeUserPermissionsParamsWithContext(ctx context.Context) *ChangeUserPermissionsParams {
	var ()
	return &ChangeUserPermissionsParams{

		Context: ctx,
	}
}

// NewChangeUserPermissionsParamsWithHTTPClient creates a new ChangeUserPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeUserPermissionsParamsWithHTTPClient(client *http.Client) *ChangeUserPermissionsParams {
	var ()
	return &ChangeUserPermissionsParams{
		HTTPClient: client,
	}
}

/*ChangeUserPermissionsParams contains all the parameters to send to the API endpoint
for the change user permissions operation typically these are written to a http.Request
*/
type ChangeUserPermissionsParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.UserPermission
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Username
	  The username of the user to which the permission applies

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change user permissions params
func (o *ChangeUserPermissionsParams) WithTimeout(timeout time.Duration) *ChangeUserPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change user permissions params
func (o *ChangeUserPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change user permissions params
func (o *ChangeUserPermissionsParams) WithContext(ctx context.Context) *ChangeUserPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change user permissions params
func (o *ChangeUserPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change user permissions params
func (o *ChangeUserPermissionsParams) WithHTTPClient(client *http.Client) *ChangeUserPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change user permissions params
func (o *ChangeUserPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change user permissions params
func (o *ChangeUserPermissionsParams) WithBody(body *models.UserPermission) *ChangeUserPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change user permissions params
func (o *ChangeUserPermissionsParams) SetBody(body *models.UserPermission) {
	o.Body = body
}

// WithRepository adds the repository to the change user permissions params
func (o *ChangeUserPermissionsParams) WithRepository(repository string) *ChangeUserPermissionsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the change user permissions params
func (o *ChangeUserPermissionsParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithUsername adds the username to the change user permissions params
func (o *ChangeUserPermissionsParams) WithUsername(username string) *ChangeUserPermissionsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the change user permissions params
func (o *ChangeUserPermissionsParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeUserPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
