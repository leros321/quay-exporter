// Code generated by go-swagger; DO NOT EDIT.

package robot

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
)

// NewCreateOrgRobotParams creates a new CreateOrgRobotParams object
// with the default values initialized.
func NewCreateOrgRobotParams() *CreateOrgRobotParams {
	var ()
	return &CreateOrgRobotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrgRobotParamsWithTimeout creates a new CreateOrgRobotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrgRobotParamsWithTimeout(timeout time.Duration) *CreateOrgRobotParams {
	var ()
	return &CreateOrgRobotParams{

		timeout: timeout,
	}
}

// NewCreateOrgRobotParamsWithContext creates a new CreateOrgRobotParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrgRobotParamsWithContext(ctx context.Context) *CreateOrgRobotParams {
	var ()
	return &CreateOrgRobotParams{

		Context: ctx,
	}
}

// NewCreateOrgRobotParamsWithHTTPClient creates a new CreateOrgRobotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrgRobotParamsWithHTTPClient(client *http.Client) *CreateOrgRobotParams {
	var ()
	return &CreateOrgRobotParams{
		HTTPClient: client,
	}
}

/*CreateOrgRobotParams contains all the parameters to send to the API endpoint
for the create org robot operation typically these are written to a http.Request
*/
type CreateOrgRobotParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*RobotShortname
	  The short name for the robot, without any user or organization prefix

	*/
	RobotShortname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create org robot params
func (o *CreateOrgRobotParams) WithTimeout(timeout time.Duration) *CreateOrgRobotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create org robot params
func (o *CreateOrgRobotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create org robot params
func (o *CreateOrgRobotParams) WithContext(ctx context.Context) *CreateOrgRobotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create org robot params
func (o *CreateOrgRobotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create org robot params
func (o *CreateOrgRobotParams) WithHTTPClient(client *http.Client) *CreateOrgRobotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create org robot params
func (o *CreateOrgRobotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgname adds the orgname to the create org robot params
func (o *CreateOrgRobotParams) WithOrgname(orgname string) *CreateOrgRobotParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the create org robot params
func (o *CreateOrgRobotParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WithRobotShortname adds the robotShortname to the create org robot params
func (o *CreateOrgRobotParams) WithRobotShortname(robotShortname string) *CreateOrgRobotParams {
	o.SetRobotShortname(robotShortname)
	return o
}

// SetRobotShortname adds the robotShortname to the create org robot params
func (o *CreateOrgRobotParams) SetRobotShortname(robotShortname string) {
	o.RobotShortname = robotShortname
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrgRobotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	// path param robot_shortname
	if err := r.SetPathParam("robot_shortname", o.RobotShortname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}