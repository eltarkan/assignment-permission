// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPermissionHandlerFunc turns a function with the right signature into a get permission handler
type GetPermissionHandlerFunc func(GetPermissionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPermissionHandlerFunc) Handle(params GetPermissionParams) middleware.Responder {
	return fn(params)
}

// GetPermissionHandler interface for that can handle valid get permission params
type GetPermissionHandler interface {
	Handle(GetPermissionParams) middleware.Responder
}

// NewGetPermission creates a new http.Handler for the get permission operation
func NewGetPermission(ctx *middleware.Context, handler GetPermissionHandler) *GetPermission {
	return &GetPermission{Context: ctx, Handler: handler}
}

/*
	GetPermission swagger:route GET /permission Permission getPermission

# Get Permission

Get Permission
*/
type GetPermission struct {
	Context *middleware.Context
	Handler GetPermissionHandler
}

func (o *GetPermission) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPermissionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
