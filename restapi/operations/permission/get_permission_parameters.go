// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetPermissionParams creates a new GetPermissionParams object
//
// There are no default values defined in the spec.
func NewGetPermissionParams() GetPermissionParams {

	return GetPermissionParams{}
}

// GetPermissionParams contains all the bound params for the get permission operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetPermission
type GetPermissionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPermissionParams() beforehand.
func (o *GetPermissionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
