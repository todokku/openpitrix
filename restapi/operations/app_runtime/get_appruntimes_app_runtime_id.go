// Code generated by go-swagger; DO NOT EDIT.

package app_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAppruntimesAppRuntimeIDHandlerFunc turns a function with the right signature into a get appruntimes app runtime ID handler
type GetAppruntimesAppRuntimeIDHandlerFunc func(GetAppruntimesAppRuntimeIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAppruntimesAppRuntimeIDHandlerFunc) Handle(params GetAppruntimesAppRuntimeIDParams) middleware.Responder {
	return fn(params)
}

// GetAppruntimesAppRuntimeIDHandler interface for that can handle valid get appruntimes app runtime ID params
type GetAppruntimesAppRuntimeIDHandler interface {
	Handle(GetAppruntimesAppRuntimeIDParams) middleware.Responder
}

// NewGetAppruntimesAppRuntimeID creates a new http.Handler for the get appruntimes app runtime ID operation
func NewGetAppruntimesAppRuntimeID(ctx *middleware.Context, handler GetAppruntimesAppRuntimeIDHandler) *GetAppruntimesAppRuntimeID {
	return &GetAppruntimesAppRuntimeID{Context: ctx, Handler: handler}
}

/*GetAppruntimesAppRuntimeID swagger:route GET /appruntimes/{appRuntimeId} app-runtime getAppruntimesAppRuntimeId

Gets an app runtime

Returns a single runtime by its id

*/
type GetAppruntimesAppRuntimeID struct {
	Context *middleware.Context
	Handler GetAppruntimesAppRuntimeIDHandler
}

func (o *GetAppruntimesAppRuntimeID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAppruntimesAppRuntimeIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}