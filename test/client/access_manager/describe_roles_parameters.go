// Code generated by go-swagger; DO NOT EDIT.

package access_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeRolesParams creates a new DescribeRolesParams object
// with the default values initialized.
func NewDescribeRolesParams() *DescribeRolesParams {
	var ()
	return &DescribeRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRolesParamsWithTimeout creates a new DescribeRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeRolesParamsWithTimeout(timeout time.Duration) *DescribeRolesParams {
	var ()
	return &DescribeRolesParams{

		timeout: timeout,
	}
}

// NewDescribeRolesParamsWithContext creates a new DescribeRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeRolesParamsWithContext(ctx context.Context) *DescribeRolesParams {
	var ()
	return &DescribeRolesParams{

		Context: ctx,
	}
}

// NewDescribeRolesParamsWithHTTPClient creates a new DescribeRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeRolesParamsWithHTTPClient(client *http.Client) *DescribeRolesParams {
	var ()
	return &DescribeRolesParams{
		HTTPClient: client,
	}
}

/*DescribeRolesParams contains all the parameters to send to the API endpoint
for the describe roles operation typically these are written to a http.Request
*/
type DescribeRolesParams struct {

	/*Portal*/
	Portal []string
	/*RoleID*/
	RoleID []string
	/*RoleName*/
	RoleName []string
	/*UserID*/
	UserID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe roles params
func (o *DescribeRolesParams) WithTimeout(timeout time.Duration) *DescribeRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe roles params
func (o *DescribeRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe roles params
func (o *DescribeRolesParams) WithContext(ctx context.Context) *DescribeRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe roles params
func (o *DescribeRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe roles params
func (o *DescribeRolesParams) WithHTTPClient(client *http.Client) *DescribeRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe roles params
func (o *DescribeRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPortal adds the portal to the describe roles params
func (o *DescribeRolesParams) WithPortal(portal []string) *DescribeRolesParams {
	o.SetPortal(portal)
	return o
}

// SetPortal adds the portal to the describe roles params
func (o *DescribeRolesParams) SetPortal(portal []string) {
	o.Portal = portal
}

// WithRoleID adds the roleID to the describe roles params
func (o *DescribeRolesParams) WithRoleID(roleID []string) *DescribeRolesParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the describe roles params
func (o *DescribeRolesParams) SetRoleID(roleID []string) {
	o.RoleID = roleID
}

// WithRoleName adds the roleName to the describe roles params
func (o *DescribeRolesParams) WithRoleName(roleName []string) *DescribeRolesParams {
	o.SetRoleName(roleName)
	return o
}

// SetRoleName adds the roleName to the describe roles params
func (o *DescribeRolesParams) SetRoleName(roleName []string) {
	o.RoleName = roleName
}

// WithUserID adds the userID to the describe roles params
func (o *DescribeRolesParams) WithUserID(userID []string) *DescribeRolesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the describe roles params
func (o *DescribeRolesParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesPortal := o.Portal

	joinedPortal := swag.JoinByFormat(valuesPortal, "multi")
	// query array param portal
	if err := r.SetQueryParam("portal", joinedPortal...); err != nil {
		return err
	}

	valuesRoleID := o.RoleID

	joinedRoleID := swag.JoinByFormat(valuesRoleID, "multi")
	// query array param role_id
	if err := r.SetQueryParam("role_id", joinedRoleID...); err != nil {
		return err
	}

	valuesRoleName := o.RoleName

	joinedRoleName := swag.JoinByFormat(valuesRoleName, "multi")
	// query array param role_name
	if err := r.SetQueryParam("role_name", joinedRoleName...); err != nil {
		return err
	}

	valuesUserID := o.UserID

	joinedUserID := swag.JoinByFormat(valuesUserID, "multi")
	// query array param user_id
	if err := r.SetQueryParam("user_id", joinedUserID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}