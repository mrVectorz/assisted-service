// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CancelInstallationHandlerFunc turns a function with the right signature into a cancel installation handler
type CancelInstallationHandlerFunc func(CancelInstallationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CancelInstallationHandlerFunc) Handle(params CancelInstallationParams) middleware.Responder {
	return fn(params)
}

// CancelInstallationHandler interface for that can handle valid cancel installation params
type CancelInstallationHandler interface {
	Handle(CancelInstallationParams) middleware.Responder
}

// NewCancelInstallation creates a new http.Handler for the cancel installation operation
func NewCancelInstallation(ctx *middleware.Context, handler CancelInstallationHandler) *CancelInstallation {
	return &CancelInstallation{Context: ctx, Handler: handler}
}

/*CancelInstallation swagger:route POST /clusters/{cluster_id}/actions/cancel installer cancelInstallation

Cancels an ongoing installation.

*/
type CancelInstallation struct {
	Context *middleware.Context
	Handler CancelInstallationHandler
}

func (o *CancelInstallation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCancelInstallationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}