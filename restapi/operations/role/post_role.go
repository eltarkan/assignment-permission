// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostRoleHandlerFunc turns a function with the right signature into a post role handler
type PostRoleHandlerFunc func(PostRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRoleHandlerFunc) Handle(params PostRoleParams) middleware.Responder {
	return fn(params)
}

// PostRoleHandler interface for that can handle valid post role params
type PostRoleHandler interface {
	Handle(PostRoleParams) middleware.Responder
}

// NewPostRole creates a new http.Handler for the post role operation
func NewPostRole(ctx *middleware.Context, handler PostRoleHandler) *PostRole {
	return &PostRole{Context: ctx, Handler: handler}
}

/*
	PostRole swagger:route POST /role Role postRole

# Create Role

Create Role
*/
type PostRole struct {
	Context *middleware.Context
	Handler PostRoleHandler
}

func (o *PostRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostRoleParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
